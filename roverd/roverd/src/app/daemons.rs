use anyhow::Context;
use std::{
    fs::{self, Permissions},
    os::unix::fs::PermissionsExt,
    path::PathBuf,
    process::Stdio,
    sync::Arc,
    time::{Duration, SystemTime, UNIX_EPOCH},
};

use openapi::models::ProcessStatus;
use rovervalidate::{config::Validate, service::Service};
use tokio::{
    process::Command,
    sync::{broadcast, RwLock},
    time::sleep,
};
use tracing::{error, info, warn};

use crate::util::*;
use crate::{command::ParsedCommand, constants::*};
use crate::{error::Error, time_now};

use super::bootspec::{Input, Stream};
use super::{
    bootspec::{BootSpec, BootSpecTuning},
    process::Process,
};

#[derive(Debug, Clone)]
pub struct DaemonManager {
    pub shutdown_tx: broadcast::Sender<()>,
    restart_coutner: Arc<RwLock<usize>>,
}

/// Here are two hard coded daemons, ideally we make this extensible by specifying
/// a separate "daemon pipeline" in the rover file. Don't need to expose setters
/// via the API, but getters would be nice. For now, this is hardcoded and error prone,
/// but it also will not change for the forseeable future.
impl DaemonManager {
    pub async fn new() -> Result<Self, Error> {
        let shutdown_tx = broadcast::channel::<()>(1).0;

        // First make sure the daemons are installed, this can fail which will
        // put roverd in a non operational state.
        let display_download = async move {
            match download_and_install_service(&DISPLAY_FETCH_URL.to_string(), true).await {
                Ok(fq) => Ok(fq),
                Err(e) => {
                    warn!("was not able to get latest daemon at {}", DISPLAY_FETCH_URL);
                    warn!("{:?}", e);
                    find_latest_daemon("vu-ase", "display")
                }
            }
        };

        let battery_download = async move {
            match download_and_install_service(&BATTERY_FETCH_URL.to_string(), true).await {
                Ok(fq) => Ok(fq),
                Err(e) => {
                    warn!("was not able to get latest daemon at {}", BATTERY_FETCH_URL);
                    warn!("{:?}", e);
                    find_latest_daemon("vu-ase", "battery")
                }
            }
        };

        let (display_result, battery_result) = tokio::join!(display_download, battery_download);

        let display_fq = display_result?;
        let battery_fq = battery_result?;

        let display_service_file = std::fs::read_to_string(display_fq.path()).map_err(|_| {
            Error::ServiceNotFound(format!("could not find {} on disk", display_fq.path()))
        })?;
        let battery_service_file = std::fs::read_to_string(battery_fq.path()).map_err(|_| {
            Error::ServiceNotFound(format!("could not find {} on disk", battery_fq.path()))
        })?;

        let display_service: Service = serde_yaml::from_str(&display_service_file)
            .with_context(|| format!("failed to parse {}", display_service_file))?;
        let battery_service: Service = serde_yaml::from_str(&battery_service_file)
            .with_context(|| format!("failed to parse {}", battery_service_file))?;

        let display_service = display_service.validate()?;
        let battery_service = battery_service.validate()?;

        let voltage_stream = Stream {
            name: BATTERY_STREAM_NAME.to_string(),
            address: format!("{}:{}", DATA_ADDRESS, BATTERY_PORT),
        };

        let display_name = display_service.0.get_pipeline_name();
        let battery_name = battery_service.0.get_pipeline_name();

        let display_bootspec = BootSpec {
            name: display_name.clone(),
            version: display_service.0.version.clone(),
            inputs: vec![Input {
                service: battery_name.clone(),
                streams: vec![voltage_stream.clone()],
            }],
            outputs: vec![],
            configuration: vec![],
            tuning: BootSpecTuning {
                enabled: false,
                address: format!("{}:{}", DATA_ADDRESS, START_PORT), // transceiver address
            },
        };

        let battery_bootspec = BootSpec {
            name: battery_name.clone(),
            version: battery_service.0.version.clone(),
            inputs: vec![],
            outputs: vec![voltage_stream],
            configuration: vec![],
            tuning: BootSpecTuning {
                enabled: false,
                address: format!("{}:{}", DATA_ADDRESS, START_PORT), // transceiver address
            },
        };

        let display_injected_env = serde_json::to_string(&display_bootspec)?;
        let battery_injected_env = serde_json::to_string(&battery_bootspec)?;

        let procs = vec![
            Process {
                fq: display_fq.clone(),
                command: display_service.0.commands.run.clone(),
                last_pid: None,
                last_exit_code: 0,
                name: display_name,
                status: ProcessStatus::Stopped,
                log_file: PathBuf::from(display_fq.log_file()),
                injected_env: display_injected_env.clone(),
                faults: 0,
                start_time: time_now!() as i64,
            },
            Process {
                fq: battery_fq.clone(),
                command: battery_service.0.commands.run.clone(),
                last_pid: None,
                last_exit_code: 0,
                name: battery_name,
                status: ProcessStatus::Stopped,
                log_file: PathBuf::from(battery_fq.log_file()),
                injected_env: battery_injected_env.clone(),
                faults: 0,
                start_time: time_now!() as i64,
            },
        ];

        let daemon_manager = DaemonManager {
            shutdown_tx,
            restart_coutner: Arc::new(RwLock::new(0)),
        };
        daemon_manager.start_daemons(procs).await?;

        Ok(daemon_manager)
    }

    #[allow(unreachable_code)]
    pub async fn start_daemons(&self, procs: Vec<Process>) -> Result<(), Error> {
        for proc in procs {
            let parsed_command = ParsedCommand::try_from(&proc.command)?;
            let mut shutdown_rx = self.shutdown_tx.subscribe();

            let restart_counter_clone = self.restart_coutner.clone();

            let log_file = create_log_file(&proc.log_file)?;
            let stdout = Stdio::from(
                log_file
                    .try_clone()
                    .with_context(|| format!("failed to clone log file {:?}", log_file))?,
            );
            let stderr = Stdio::from(log_file);
            let full_path = format!("{}/{}", proc.fq.dir(), &parsed_command.program);
            fs::set_permissions(&full_path, Permissions::from_mode(0o755))
                .with_context(|| format!("failed to set 755 permissions to {}", full_path))?;
            let mut command = Command::new(&parsed_command.program);
            command
                .args(&parsed_command.arguments)
                .env(ENV_KEY, &proc.injected_env)
                .current_dir(proc.fq.dir())
                .stdout(stdout)
                .stderr(stderr);

            tokio::spawn(async move {
                loop {
                    tokio::select! {
                        _ = shutdown_rx.recv() => {
                            info!("shutdown signal received in daemon {}", proc.name);
                            break;
                        }
                        result = async {
                        match command.spawn() {
                            Ok(mut child) => {
                                info!("daemon '{}' started", proc.name);
                                match child.wait().await {
                                    Ok(status) => {
                                        info!(
                                            "daemon '{}' exited with status: {}",
                                            proc.name, status
                                        )
                                    }
                                    Err(e) => info!("daemon '{}' error: {}", proc.name, e),
                                }
                            }
                            Err(e) => {
                                error!("could not start daemon '{}': {}", proc.name, e);
                            }
                        }
                        Ok::<(), Error>(())
                        } => {
                            if let Err(e) = result {
                                error!("error in daemon '{}': {:?}", proc.name, e);
                            }

                            // Check shutdown signal before restarting
                            if shutdown_rx.try_recv().is_ok() {
                                info!("shutdown signal received for daemon '{}', not restarting", proc.name);
                                break;
                            }

                            // Sleep so that we don't retry immediately
                            sleep(Duration::from_secs(3)).await;

                            // Introduce a scope here so that we give up the
                            let mut counter = restart_counter_clone.write().await;
                            if *counter > 20 {
                                info!("restarted daemons {} times, giving up", *counter);
                                break;
                            }
                            info!("restarting daemon '{}' ({})", proc.name, *counter);
                            *counter += 1;
                        }
                    }
                }
                info!("no longer attempting to start {}", proc.name);
            });
        }

        Ok(())
    }
}

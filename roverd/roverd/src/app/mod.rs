use anyhow::{anyhow, Context};
use axum_extra::extract::Multipart;
use openapi::models::*;
use process::{PipelineStats, Process, SpawnedProcess};
use rovervalidate::config::{Configuration, ValidatedConfiguration};
use rovervalidate::pipeline::interface::{Pipeline, RunnablePipeline};
use rovervalidate::service::{Service, ValidatedService};
use rovervalidate::validate::Validate;
use serde_json::Value;
use service::{Fq, FqBuf, FqBufVec, FqVec};
use state::{Dormant, Operating, RoverState};
use std::cmp;
use std::collections::{HashMap, HashSet};
use std::fs::{self, remove_file};
use std::fs::{File, OpenOptions};
use std::io::{BufRead, BufReader, Write};
use std::io::{Read, Seek, SeekFrom};
use std::marker::PhantomData;
use std::os::unix::process::ExitStatusExt;
use std::path::{Path, PathBuf};
use std::process::Stdio;
use std::sync::Arc;
use std::time::{SystemTime, UNIX_EPOCH};
use sysinfo::{CpuRefreshKind, MemoryRefreshKind, Pid, ProcessRefreshKind, RefreshKind, System};
use tokio::process::Command;
use tokio::select;
use tokio::sync::{broadcast, broadcast::Sender, Mutex, RwLock};
use tracing::{error, info, warn};

use crate::error::Error;
use crate::util::*;
use crate::{constants::*, time_now};

mod bootspec;
pub mod daemons;
pub mod info;
pub mod process;
pub mod service;
pub mod state;

/// The main struct that implements functions called from the api and holds all objects
/// in memory necessary for operation. Info member holds static information derived mostly
/// from the
#[derive(Debug, Clone)]
pub struct Roverd {
    /// Information related to the roverd daemon, contains status.
    pub info: info::Info,

    /// Run-time data structures of the Rover, interacts with the file system
    /// and spawns processes, so must be read/write locked.
    pub app: App,
}

impl Roverd {
    pub async fn new() -> Result<Self, Error> {
        let info = info::Info::new();

        let roverd = Self {
            info,
            app: App {
                processes: Arc::new(RwLock::new(vec![])),
                spawned: Arc::new(RwLock::new(vec![])),
                stats: Arc::new(RwLock::new(PipelineStats {
                    status: PipelineStatus::Startable,
                    last_start: None,
                    last_stop: None,
                    last_restart: None,
                })),
                shutdown_tx: broadcast::channel::<()>(1).0,
                built_services: Arc::new(RwLock::new(HashMap::new())),
                sysinfo: Arc::new(RwLock::new(System::new_with_specifics(
                    RefreshKind::nothing()
                        .with_processes(ProcessRefreshKind::everything())
                        .with_cpu(CpuRefreshKind::everything())
                        .with_memory(MemoryRefreshKind::everything()),
                ))),
            },
        };

        // Validate the pipeline found in the yaml file, if it is not valid
        // empty it.
        let initial_pipieline_status = match roverd.app.get_valid_pipeline().await {
            Ok(_) => PipelineStatus::Startable,
            Err(_) => PipelineStatus::Empty,
        };

        {
            let mut stats = roverd.app.stats.write().await;
            stats.status = initial_pipieline_status;

            if roverd.info.status != DaemonStatus::Operational {
                warn!("did not initialize successfully {:#?}", roverd.info);
            }
        }

        Ok(roverd)
    }

    /// Checks the pipeline state for if the rover is currently dormant.
    pub async fn try_get_dormant(&self) -> Option<RoverState<Dormant>> {
        let stats = self.app.stats.read().await;
        match stats.status {
            PipelineStatus::Started => None,
            _ => Some(RoverState {
                _ignore: PhantomData,
            }),
        }
    }

    /// Checks the pipeline state for if the rover is currently operating.
    pub async fn try_get_operating(&self) -> Option<RoverState<Operating>> {
        let stats = self.app.stats.read().await;
        match stats.status {
            PipelineStatus::Started => Some(RoverState {
                _ignore: PhantomData,
            }),
            _ => None,
        }
    }
}

impl AsRef<Roverd> for Roverd {
    fn as_ref(&self) -> &Roverd {
        self
    }
}

#[derive(Debug, Clone)]
pub struct App {
    /// Contains the "application view" of process after validation. In-between start / stop
    /// runs this vec remains unchanged.
    pub processes: Arc<RwLock<Vec<Process>>>,

    /// The "runtime" view of all processes, this contains handles to the spawned children.
    pub spawned: Arc<RwLock<Vec<SpawnedProcess>>>,

    /// Overall status of the pipeline.
    pub stats: Arc<RwLock<PipelineStats>>,

    /// Broadcast channel to send shutdown command for termination.
    pub shutdown_tx: Sender<()>,

    // Look up the last built time of a service on disk.
    pub built_services: Arc<RwLock<HashMap<FqBuf, i64>>>,

    // System information initialized once
    pub sysinfo: Arc<RwLock<System>>,
}

impl App {
    pub async fn should_invalidate(&self, fq_buf: &FqBuf) -> Result<bool, Error> {
        let mut config = get_config().await?;
        let enabled_fq = FqVec::try_from(&config.enabled)?;
        let pipeline_invalidated = enabled_fq.0.contains(&Fq::from(fq_buf));

        if pipeline_invalidated {
            config.enabled.clear();
            update_config(&config)?;
        }

        Ok(pipeline_invalidated)
    }

    /// Downloads the service from the specified url.
    pub async fn fetch_service(
        &self,
        body: &FetchPostRequest,
        _: RoverState<Dormant>,
    ) -> Result<(FqBuf, bool), Error> {
        let fq_buf = download_and_install_service(&body.url, false).await?;
        let invalidate_pipline = self.should_invalidate(&fq_buf).await?;
        Ok((fq_buf, invalidate_pipline))
    }

    pub async fn receive_upload(
        &self,
        mut body: Multipart,
        _: RoverState<Dormant>,
    ) -> Result<(FqBuf, bool), Error> {
        if let Some(field) = body
            .next_field()
            .await
            .map_err(|_| Error::ServiceUploadBadPayload)?
        {
            // Extract bytes from payload
            let data = field
                .bytes()
                .await
                .map_err(|_| Error::ServiceUploadBadPayload)?;

            // Ignore errors, since filesystem can be in any state and
            // get a clean slate of the zip file
            let _ = remove_file(ZIP_FILE);

            // Create the zip file handle
            let mut file = fs::OpenOptions::new()
                .create(true)
                .truncate(true)
                .write(true)
                .open(ZIP_FILE)
                .with_context(|| format!("failed to create file {}", ZIP_FILE))?;

            file.write_all(&data)
                .with_context(|| format!("failed to write data to {}", ZIP_FILE))?;

            let (fq_buf, directory) = extract_fq_from_zip(ZIP_FILE.to_string()).await?;

            // syncing can overwrite the current contents
            // if service_exists(&Fq::from(&fq_buf))? {
            //     return Err(Error::ServiceAlreadyExists);
            // }

            install_service(directory, &fq_buf).await?;

            let invalidate_pipline = self.should_invalidate(&fq_buf).await?;

            return Ok((fq_buf, invalidate_pipline));
        }
        Err(Error::ServiceUploadBadPayload)
    }

    /// Returns all authors from the rover directory
    pub async fn get_authors(&self) -> Result<Vec<String>, Error> {
        list_dir_contents("")
    }

    /// Returns all services given an author
    pub async fn get_services(
        &self,
        path_params: ServicesAuthorGetPathParams,
    ) -> Result<Vec<String>, Error> {
        list_dir_contents(&path_params.author.to_string())
    }

    /// Returns all versions given a service
    pub async fn get_versions(
        &self,
        path_params: ServicesAuthorServiceGetPathParams,
    ) -> Result<Vec<String>, Error> {
        list_dir_contents(format!("{}/{}", path_params.author, path_params.service).as_str())
    }

    /// Returns a valid service given a fully qualified name
    pub async fn get_service(&self, fq: FqBuf) -> Result<ValidatedService, Error> {
        let contents = fs::read_to_string(fq.path())
            .map_err(|_| Error::ServiceNotFound(format!("Could not find {} on disk", fq.path())))?;
        let service =
            serde_yaml::from_str::<rovervalidate::service::Service>(&contents)?.validate()?;

        Ok(service)
    }

    /// Writes the validated service to the corresponding service.yaml on disk.
    /// If for some reason it doesn't exist, it will error.
    pub async fn update_service(&self, service: ValidatedService) -> Result<(), Error> {
        // Check if the service.yaml exists
        let fq_path = FqBuf::from(&service).path();

        // Convert to a Path and check if it exists
        let service_path = Path::new(&fq_path);
        if !service_path.exists() {
            return Err(Error::ServiceNotFound(fq_path));
        }

        let contents = serde_yaml::to_string(&service.0)?;

        let mut file = OpenOptions::new()
            .write(true)
            .create(true)
            .truncate(true)
            .open(fq_path)
            .map_err(Error::Io)?;

        file.write_all(contents.as_bytes()).map_err(Error::Io)?;

        Ok(())
    }

    /// Deletes a service from the filesystem. Note: this only removes it from the final
    /// version directory in the "author/service/version" hierarchy, so directories may
    /// stick around
    pub async fn delete_service(
        &self,
        path_params: &ServicesAuthorServiceVersionDeletePathParams,
        _: RoverState<Dormant>,
    ) -> Result<bool, Error> {
        let delete_fq = FqBuf::from(path_params);

        // Get the current configuration from disk
        let mut config = get_config().await?;

        let mut built_services = self.built_services.write().await;

        let mut should_reset = false;
        // Return whether or not the service was enabled and if it was,
        // reset the pipeline
        let enabled_fq_vec = FqBufVec::try_from(&config.enabled)?.0;

        if enabled_fq_vec.contains(&delete_fq) {
            should_reset = true;
            config.enabled.clear();
            update_config(&config)?;

            built_services.remove(&delete_fq);
        }

        // Remove the service to delete from the filesystem
        if Path::new(&delete_fq.dir()).exists() {
            std::fs::remove_dir_all(delete_fq.dir())
                .with_context(|| format!("failed to remove {}", delete_fq.dir()))?;
        } else {
            return Err(Error::ServiceNotFound(format!(
                "wanted to delete {}, but it never existed",
                delete_fq.dir()
            )));
        }

        Ok(should_reset)
    }

    /// Performs the build command for a given service and does so by instantiating a login shell
    /// of the debix user. It also first changes into the service's directory.
    pub async fn build_service(
        &self,
        params: ServicesAuthorServiceVersionPostPathParams,
        _: RoverState<Dormant>,
    ) -> Result<(), Error> {
        let fq = FqBuf::from(&params);
        let service = self.get_service(fq.clone()).await?.0;

        let build_string = &service
            .commands
            .build
            .ok_or_else(|| Error::BuildCommandMissing)?;
        let log_file = create_log_file(&PathBuf::from(fq.build_log_file()))?;
        let stdout = Stdio::from(
            log_file
                .try_clone()
                .with_context(|| format!("failed to clone build-log file {:?}", log_file))?,
        );
        let stderr = Stdio::from(log_file);

        let mut built_services = self.built_services.write().await;
        let corrected_build_command = format!("cd {} ; {}", fq.dir(), build_string.as_str());

        // Run the build command in the login shell of the debix user (necessary for build deps)
        match Command::new("su")
            .args(["-", "debix", "-c", corrected_build_command.as_str()])
            .stdout(stdout)
            .stderr(stderr)
            .current_dir(fq.dir())
            .spawn()
        {
            Ok(mut child) => match child.wait().await {
                Ok(exit_status) => {
                    if !exit_status.success() {
                        // Build was not successful, return logs
                        let file = std::fs::File::open(fq.build_log_file())
                            .with_context(|| format!("failed to open {}", fq.build_log_file()))?;
                        let reader = BufReader::new(file);
                        let lines: Vec<String> = reader
                            .lines()
                            .collect::<Result<Vec<String>, _>>()
                            .with_context(|| {
                                format!("failed to collect lines from {}", fq.build_log_file())
                            })?;
                        return Err(Error::BuildLog(lines));
                    } else {
                        // Build was successful, save time it was built at
                        let time_now = time_now!() as i64;

                        *built_services.entry(fq).or_insert_with(|| time_now) = time_now;
                    }
                    Ok(())
                }
                Err(e) => {
                    error!("failed to wait on build command: {:?} {}", &build_string, e);
                    Err(Error::BuildCommandFailed)
                }
            },
            Err(e) => {
                error!("failed to spawn build command: {:?} {}", &build_string, e);
                Err(Error::BuildCommandFailed)
            }
        }
    }

    /// Updates the pipeline if it is valid. Updating the pipeline is only allowed when
    /// the rover is not currently running.
    pub async fn set_pipeline(
        &self,
        incoming_pipeline: Vec<PipelinePostRequestInner>,
        _: RoverState<Dormant>,
    ) -> Result<(), Error> {
        let services = FqBufVec::from(incoming_pipeline).0;

        let mut valid_services = vec![];

        for enabled in &services {
            let service_file = std::fs::read_to_string(enabled.path()).map_err(|_| {
                Error::ServiceNotFound(format!("could not find or read {}", enabled.path()))
            })?;
            let service: Service = serde_yaml::from_str(&service_file)?;
            valid_services.push(service.validate()?);
        }

        let _ = Pipeline::new(valid_services).validate()?;

        // Here we have a valid pipeline, so rover.yaml can be overwritten
        let mut config = get_config().await?;
        config.enabled.clear();

        // Services are valid since we didn't return earlier
        for service in services {
            config.enabled.push(service.path())
        }

        update_config(&config)?;

        let mut stats = self.stats.write().await;

        if config.enabled.is_empty() {
            stats.status = PipelineStatus::Empty;
        } else {
            stats.status = PipelineStatus::Startable;
        }

        Ok(())
    }

    /// Gets the current pipeline along with the list of processes if they are running.
    pub async fn get_pipeline(&self) -> Result<Vec<PipelineGet200ResponseEnabledInner>, Error> {
        let stats = self.stats.read().await;
        if stats.status == PipelineStatus::Empty {
            let config = Configuration { enabled: vec![] };
            update_config(&config)?;
        }

        let conf = get_config().await?;

        let processes = self.processes.read().await;

        let mut responses = vec![];

        for validated_service in conf.enabled.into_iter() {
            let fq = FqBuf::try_from(validated_service)?;

            let mut sysinfo = self.sysinfo.write().await;

            sysinfo.refresh_specifics(
                RefreshKind::nothing()
                    .with_processes(ProcessRefreshKind::everything())
                    .with_cpu(CpuRefreshKind::everything())
                    .with_memory(MemoryRefreshKind::everything()),
            );

            let proc = get_proc(fq.clone(), &processes);
            responses.push(match proc {
                Ok(p) => {
                    if let Some(pid) = p.last_pid {
                        let status = p.status;
                        let pid = pid as i32;
                        let mut memory = 0;
                        let mut cpu = 0;
                        let uptime = (time_now!() as i64) - p.start_time;
                        // todo: test this with a substantial payload on the debix
                        if let Some(proc_info) = sysinfo.process(Pid::from(pid as usize)) {
                            memory = (proc_info.memory() / 1000000_u64) as i32;
                            cpu = (proc_info.cpu_usage() * 100.0) as i32;
                        }
                        PipelineGet200ResponseEnabledInner {
                            process: Some(PipelineGet200ResponseEnabledInnerProcess {
                                cpu,
                                pid,
                                uptime,
                                memory,
                                status,
                            }),
                            service: PipelineGet200ResponseEnabledInnerService {
                                fq: FullyQualifiedService::from(fq),
                                faults: p.faults as i32,
                                exit: p.last_exit_code,
                            },
                        }
                    } else {
                        PipelineGet200ResponseEnabledInner {
                            process: None,
                            service: PipelineGet200ResponseEnabledInnerService {
                                fq: FullyQualifiedService::from(fq),
                                faults: p.faults as i32,
                                exit: p.last_exit_code,
                            },
                        }
                    }
                }
                Err(_) => PipelineGet200ResponseEnabledInner {
                    process: None,
                    service: PipelineGet200ResponseEnabledInnerService {
                        fq: FullyQualifiedService::from(fq),
                        faults: 0,
                        exit: 0,
                    },
                },
            });
        }

        Ok(responses)
    }

    /// We want to start a pipeline since we are given a valid pipieline, however first we must
    /// construct the ASE_SERVICE environment variable to each service. This function prepares
    /// that variable and updates the vector of processes internally, making it ready to start.
    pub async fn construct_managed_services(
        &self,
        runnable: RunnablePipeline,
    ) -> Result<(), Error> {
        // Assign the new processes state each time, so first
        // clear the existing processes and then add them again
        let mut processes = self.processes.write().await;

        let bootspecs = bootspec::BootSpecs::new(runnable.services().clone()).0;

        let mut fqs = vec![];
        let mut service_data = vec![];

        for service in runnable.services() {
            // Create bootspecs with missing inputs, since the first step is to hand out
            // ports. After knowing the ports of all outputs, we can fill in the inputs.
            // Since we are valid, a given input will _always_ have an output.
            let fq = FqBuf::from(service);
            let bootspec = bootspecs.get(&fq);
            let injected_env = serde_json::to_string(&bootspec)?;

            // Save the necessary information from each runnable service
            fqs.push(fq);
            service_data.push((service, injected_env));
        }

        // Most of the time, we will retain all processes, however when the pipeline changes
        // we need to reflect those changes
        processes.retain(|p| fqs.contains(&p.fq));
        let fqs_and_services = fqs.iter().zip(service_data.iter());

        for (fq, (service, injected_env)) in fqs_and_services {
            if let Some(proc) = processes.iter_mut().find(|p| p.fq == *fq) {
                // If the runnable service identified by its fq already exists, there's a chance
                // that the service.yaml has changed, so update only those fields
                proc.command = service.0.commands.run.clone();
                proc.injected_env = injected_env.clone();
                proc.start_time = time_now!() as i64;
            } else {
                // The runnable service has not previously been added, so add a new one
                processes.push(Process {
                    fq: fq.clone(),
                    command: service.0.commands.run.clone(),
                    last_pid: None,
                    last_exit_code: 0,
                    name: service.0.get_original_name(),
                    status: ProcessStatus::Stopped,
                    log_file: PathBuf::from(fq.log_file()),
                    injected_env: injected_env.clone(),
                    faults: 0,
                    start_time: time_now!() as i64,
                })
            }
        }

        Ok(())
    }

    /// This function is called by the API handler. It gets the runnable pipeline from disk,
    /// constructs the processes and finally spawns the processes.
    pub async fn start(&self, _: RoverState<Dormant>) -> Result<(), Error> {
        let enabled_services = get_config().await?.enabled;

        if enabled_services.is_empty() {
            return Err(Error::PipelineIsEmpty);
        }

        // Pipeline validation step
        let runnable = self.get_valid_pipeline().await?;

        // After this, self.processes will be ready
        self.construct_managed_services(runnable).await?;

        // Start the actual processes
        self.spawn_procs().await?;

        Ok(())
    }

    /// The main starting procedure of all processes.
    pub async fn spawn_procs(&self) -> Result<(), Error> {
        let mut stats = self.stats.write().await;

        match stats.status {
            PipelineStatus::Startable => (),
            PipelineStatus::Empty => return Err(Error::PipelineIsEmpty),
            PipelineStatus::Started => return Err(Error::PipelineAlreadyStarted),
        }

        let mut spawned_procs = self.spawned.write().await;
        let mut procs = self.processes.write().await;

        spawned_procs.clear();

        for p in &mut *procs {
            roverd_log(p.log_file.clone(), format!("spawned service '{}'", p.name))?;

            let log_file = create_log_file(&p.log_file)?;

            let stdout = Stdio::from(
                log_file
                    .try_clone()
                    .with_context(|| format!("failed to clone log file {:?}", log_file))?,
            );
            let stderr = Stdio::from(log_file);

            let user_run_command = format!("cd {} && {}", p.fq.dir(), &p.command);

            let mut command = Command::new("sudo");
            command
                .args(["-u", "debix", "-E", "bash", "-c", user_run_command.as_str()])
                .env(ENV_KEY, p.injected_env.clone())
                .process_group(0)
                .current_dir(p.fq.dir())
                .stdout(stdout)
                .stderr(stderr);

            info!("executing: 'sudo -u debix -E bash -c {}", user_run_command);

            match command.spawn() {
                Ok(child) => {
                    p.status = ProcessStatus::Running;
                    if let Some(id) = child.id() {
                        info!("spawned process: {:?} at {}", p.name, id);
                        p.last_pid = Some(id);
                    } else {
                        let err_msg = format!("process: {} exited immediately", p.name);
                        warn!(err_msg);
                        p.faults += 1;
                        p.last_exit_code = 1;
                        self.cancel_start(&mut stats, &mut procs, &mut spawned_procs)
                            .await;
                        return Err(Error::FailedToSpawnProcess(err_msg));
                    }
                    spawned_procs.push(SpawnedProcess {
                        fq: p.fq.clone(),
                        name: p.name.clone(),
                        child: Arc::from(Mutex::from(child)),
                    });
                }
                Err(e) => {
                    let err_msg = format!("{}", e);
                    warn!("failed to spawn process '{}': {}", p.name, &err_msg);
                    p.faults += 1;
                    p.last_exit_code = 1;
                    self.cancel_start(&mut stats, &mut procs, &mut spawned_procs)
                        .await;
                    return Err(Error::FailedToSpawnProcess(err_msg));
                }
            }
        }

        stats.status = PipelineStatus::Started;
        stats.last_start = Some(time_now!() as i64);

        for spawned in spawned_procs.clone() {
            let mut shutdown_rx = self.shutdown_tx.subscribe();
            let process_shutdown_tx = self.shutdown_tx.clone();

            let procs_clone = Arc::clone(&self.processes);
            let stats_clone = Arc::clone(&self.stats);

            tokio::spawn(async move {
                let mut child = spawned.child.lock().await;

                select! {
                    // Wait for process completion
                    result_status = child.wait() => {
                        match result_status {
                            Ok(exit_status) => {
                                // Update the pipeline's status.
                                let mut stats = stats_clone.write().await;
                                stats.status = PipelineStatus::Startable;
                                stats.last_restart = Some(time_now!() as i64);
                                info!("child {} exited with status {}", spawned.name, exit_status);




                                let exit_code = exit_status.code();
                                let mut procs_guard = procs_clone.write().await;

                                if let Some(proc) = procs_guard.iter_mut().find(|p| p.fq == spawned.fq) {
                                    proc.status = ProcessStatus::Stopped;

                                    // In general we ignore the errors of logging since
                                    // we are inside of an async closure
                                    if let Some(e) = exit_code {
                                        proc.last_exit_code = e;
                                        let _ = roverd_log(proc.log_file.clone(), format!("service exited with code: {}", e));
                                    } else if !exit_status.success() {
                                            proc.faults += 1;
                                            let _ = roverd_log(proc.log_file.clone(), format!("service exited with code: {}", exit_status.into_raw()));
                                    } else {
                                        let _ = roverd_log(proc.log_file.clone(), "service exited".to_string());
                                    }
                                }
                                process_shutdown_tx.send(()).ok();
                            }
                            Err(e) => {
                                error!("error waiting for process {}: {}", spawned.name, e);
                                process_shutdown_tx.send(()).ok();
                            }
                        }
                    }
                    _ = shutdown_rx.recv() => {
                        // We have been sent a terminate signal, so end the process

                        // Update the pipeline's status.
                        {
                            let mut stats = stats_clone.write().await;
                            stats.status = PipelineStatus::Startable;
                        }

                        if let Some(id) = child.id() {
                            info!("terminating {} pid ({})", spawned.name, id);
                            unsafe {
                                libc::kill(id as i32, libc::SIGTERM);
                            }
                        }

                        // Wait a short while before checking if child still exists
                        tokio::time::sleep(tokio::time::Duration::from_millis(1000)).await;

                        let mut procs_guard = procs_clone.write().await;
                        if let Some(proc) = procs_guard.iter_mut().find(|p| p.fq == spawned.fq) {
                            proc.status = ProcessStatus::Terminated;
                            proc.last_exit_code = 0;
                            let _ = roverd_log(proc.log_file.clone(), "terminated by roverd".to_string());
                        }

                        // If the child has not terminated, kill it.
                        match child.try_wait() {
                            Ok(None) => {
                                info!("process {} did not terminate, killing", spawned.name);
                                if let Err(e) = child.kill().await {
                                    error!("error killing process {:?}: {:?}", spawned.name, e);
                                }
                            },
                            Ok(Some(_)) => {},
                            Err(e) => {
                                error!("Error: {:?}", e);
                                if let Err(e) = child.kill().await {
                                    error!("error killing process {:?}: {:?}", spawned.name, e);
                                }
                            }
                        }
                    }
                }
            });
        }

        Ok(())
    }

    /// If one process fails during the beginning of the starting procedure, we need to
    /// kill all started children manually, set their states and clear the spawned vec.
    pub async fn cancel_start(
        &self,
        stats: &mut PipelineStats,
        processes: &mut Vec<Process>,
        spawned_procs: &mut Vec<SpawnedProcess>,
    ) {
        warn!("cancelled spawning process");
        stats.status = PipelineStatus::Startable;
        for p in &mut *processes {
            if let Some(pid) = p.last_pid {
                unsafe {
                    libc::kill(pid as i32, libc::SIGKILL);
                }
            }
            p.status = ProcessStatus::Killed
        }

        spawned_procs.clear();
    }

    /// If the pipeline is started, it will be stopped.
    pub async fn stop(&self, _: RoverState<Operating>) -> Result<(), Error> {
        let mut stats = self.stats.write().await;

        match stats.status {
            PipelineStatus::Empty => return Err(Error::PipelineIsEmpty),
            PipelineStatus::Startable => return Err(Error::NoRunningServices),
            _ => (),
        }
        stats.status = PipelineStatus::Startable;

        stats.last_stop = Some(time_now!() as i64);
        self.shutdown_tx.send(()).ok();
        let mut spawned = self.spawned.write().await;
        spawned.clear();
        Ok(())
    }

    /// Reads the config file from disk and returns a RunnablePipeline if it is valid.
    /// If it isn't valid it returns an error and resets the config.
    pub async fn get_valid_pipeline(&self) -> Result<RunnablePipeline, Error> {
        let mut config = get_config().await?;
        let mut enabled_services: Vec<ValidatedService> = vec![];

        let res = {
            for enabled in &config.enabled {
                let service_file = std::fs::read_to_string(enabled).map_err(|_| {
                    Error::ServiceNotFound(format!("could not find or read {}", enabled))
                })?;
                let service: Service = serde_yaml::from_str(&service_file)?;
                let validated = service.validate()?;
                enabled_services.push(validated);
            }

            Pipeline::new(enabled_services).validate()
        };
        match res {
            Ok(val) => Ok(val),
            Err(e) => {
                config.enabled.clear();
                update_config(&config)?;
                Err(Error::Validation(e))
            }
        }
    }

    pub async fn get_service_logs(
        &self,
        fq: FqBuf,
        num_lines: usize,
    ) -> Result<Vec<String>, Error> {
        let file = File::open(fq.log_file()).map_err(|_| Error::NoLogsFound)?;
        let mut reader = BufReader::new(file);

        // Get the file size to determine how far we can seek
        let file_size = reader.seek(SeekFrom::End(0))?;

        // If the file is empty, return an empty vector
        if file_size == 0 {
            return Ok(Vec::new());
        }

        let mut lines = Vec::new();
        let mut position = file_size;
        let mut leftover = String::new();

        // Buffer to store chunks we read
        let chunk_size = 4096;

        // Read the file backwards until we have enough lines or reach the beginning
        while lines.len() < num_lines && position > 0 {
            // Calculate how much to read in this iteration
            let read_size = cmp::min(chunk_size, position as usize);
            position -= read_size as u64;

            // Seek to the position we want to read from
            reader.seek(SeekFrom::Start(position))?;

            // Read the chunk of data
            let mut buffer = vec![0; read_size];
            reader.read_exact(&mut buffer)?;

            // Convert the bytes to a string, handling potential UTF-8 issues
            // Note: This is a simplification. For proper UTF-8 handling, more careful
            // boundary detection would be needed.
            let chunk = match String::from_utf8(buffer) {
                Ok(s) => s,
                Err(e) => String::from_utf8_lossy(&e.into_bytes()).into_owned(),
            };

            // Combine with any leftover text from previous iteration
            let combined = chunk + &leftover;

            // Split into lines
            let mut chunk_lines: Vec<_> = combined.lines().map(String::from).collect();

            // If this isn't the first chunk (beginning of file), save the first line
            // as it might be incomplete and will be combined with the previous chunk
            if position > 0 && !chunk_lines.is_empty() {
                leftover = chunk_lines.remove(0);
            }

            // Add the lines in reverse order (since we're reading backwards)
            for line in chunk_lines.into_iter().rev() {
                lines.push(line);
                if lines.len() >= num_lines {
                    break;
                }
            }
        }
        lines.reverse();
        Ok(lines)
    }

    /// Spawns a separate shell to run the update script
    pub async fn update_rover(&self, version: String, _: RoverState<Dormant>) -> Result<(), Error> {
        let mut update_cmd = Command::new("/home/debix/ase/bin/update-roverd");
        update_cmd.arg(version);

        match update_cmd.spawn() {
            Ok(_) => (),
            Err(e) => error!("unable to spawn the update command: {}", e),
        }
        Ok(())
    }

    /// Spawns a process to shutdown the rover
    pub async fn shutdown_rover(&self, _: RoverState<Dormant>) -> Result<(), Error> {
        let mut shutdown = Command::new("shutdown");
        shutdown.arg("-h").arg("now");

        match shutdown.spawn() {
            Ok(_) => (),
            Err(e) => error!("unable to run shutdown command: {}", e),
        }
        Ok(())
    }

    /// Spawns a process to emergency stop the rover
    pub async fn emergency_stop_rover(&self) -> Result<(), Error> {
        let mut emergency = Command::new("/home/debix/ase/bin/rover-reset");

        match emergency.spawn() {
            Ok(_) => (),
            Err(e) => error!("unable to run emergency command: {}", e),
        }
        Ok(())
    }

    pub async fn get_fqns(&self) -> Result<Vec<FullyQualifiedService>, Error> {
        let mut fqns = Vec::new();
        let rover_dir = Path::new(ROVER_DIR);

        // Ensure base directory exists
        if !rover_dir.exists() {
            return Ok(fqns);
        }

        // Iterate through author directories
        for author_entry in
            fs::read_dir(rover_dir).with_context(|| format!("failed to read {:?}", rover_dir))?
        {
            let author_entry = author_entry
                .with_context(|| format!("failed to unpack an entry in {:?}", rover_dir))?;
            if !author_entry
                .file_type()
                .with_context(|| {
                    format!("failed to get file metadata of {:?}", author_entry.path())
                })?
                .is_dir()
            {
                continue;
            }
            let author_path = author_entry.path();
            let author_name =
                author_path
                    .file_name()
                    .and_then(|n| n.to_str())
                    .ok_or(Error::Context(anyhow!(
                        "failed to convert path: {:?}",
                        author_path
                    )))?;

            // Iterate through service name directories
            for service_entry in fs::read_dir(&author_path)
                .with_context(|| format!("failed to unpack an entry in {:?}", author_path))?
            {
                let service_entry = service_entry
                    .with_context(|| format!("failed to unpack an entry in {:?}", author_path))?;
                if !service_entry
                    .file_type()
                    .with_context(|| {
                        format!("failed to get file metadata of {:?}", service_entry.path())
                    })?
                    .is_dir()
                {
                    continue;
                }
                let service_path = service_entry.path();
                let service_name =
                    service_path
                        .file_name()
                        .and_then(|n| n.to_str())
                        .ok_or(Error::Context(anyhow!(
                            "failed to convert path: {:?}",
                            service_path
                        )))?;

                // Iterate through version directories
                for version_entry in fs::read_dir(&service_path)
                    .with_context(|| format!("failed to unpack an entry in {:?}", service_path))?
                {
                    let version_entry = version_entry.with_context(|| {
                        format!("failed to unpack an entry in {:?}", service_path)
                    })?;
                    if !version_entry
                        .file_type()
                        .with_context(|| {
                            format!("failed to get file metadata of {:?}", version_entry.path())
                        })?
                        .is_dir()
                    {
                        continue;
                    }
                    let version = version_entry.file_name().to_string_lossy().into_owned();

                    fqns.push(FullyQualifiedService {
                        author: author_name.to_string(),
                        name: service_name.to_string(),
                        version: version.clone(),
                        r#as: get_service_as(author_name, service_name, &version),
                    });
                }
            }
        }

        Ok(fqns)
    }

    pub async fn update_service_config(
        &self,
        path_params: &ServicesAuthorServiceVersionConfigurationPostPathParams,
        config_updates: &Vec<ServicesAuthorServiceVersionConfigurationPostRequestInner>,
        _: RoverState<Dormant>,
    ) -> Result<(), Error> {
        // Convert given service from the path parameters into an Fq that we can use.
        let fq = FqBuf::from(path_params);

        // Try to get the service's configuration from the file system.
        let mut service = self.get_service(fq).await?;

        // We check every incoming "key" with the available "name" from the config
        let mut unique_keys = HashSet::new();
        for item in config_updates {
            let key = item.key.clone();

            // Check if any incoming key does not exist in the config file on disk
            if !service.0.configuration.iter().any(|item| item.name == key) {
                return Err(Error::InvalidKey(format!(
                    "Attempted to update key: '{}' which does not exist",
                    key
                )));
            }

            // If we already had this key, then there is a duplicate!
            if !unique_keys.insert(key) {
                return Err(Error::DuplicateKey(
                    "There exists a duplicate key in the config update request".to_string(),
                ));
            }
        }

        // For all the incoming configuration updates if the types are correct, we can update
        // the service object.
        for item in config_updates {
            let update_key = item.key.clone();
            let update_value = serde_json::from_str::<Value>(item.value.0.get())?;

            for config in service.0.configuration.iter_mut() {
                if update_key == config.name {
                    match &update_value {
                        Value::Number(n) => {
                            if let rovervalidate::service::Value::Double(_) = config.value {
                                if let Some(number) = n.as_f64() {
                                    // Here we are sure that the incoming json value is a Number,
                                    // and that the type stored on disk was a double, so we can update
                                    config.value = rovervalidate::service::Value::Double(number);
                                } else {
                                    return Err(Error::InvalidKeyType(format!(
                                        "Key '{}' is of type Number, however unable to cast it to an f64",
                                        update_key
                                    )));
                                }
                            } else {
                                return Err(Error::InvalidKeyType(format!(
                                    "Key '{}' is of type String, however a Number was supplied",
                                    update_key
                                )));
                            }
                        }
                        Value::String(s) => {
                            if let rovervalidate::service::Value::String(_) = config.value {
                                // Here we are sure that the incoming json value is a String,
                                // and that the type stored on disk was also a String, so we can update
                                config.value = rovervalidate::service::Value::String(s.clone())
                            } else {
                                return Err(Error::InvalidKeyType(format!(
                                    "Key '{}' is of type Number, however a String was supplied",
                                    update_key
                                )));
                            }
                        }
                        _ => {
                            return Err(Error::InvalidKeyType(format!(
                                "The supplied key '{}' was neither a Number or String",
                                update_key
                            )));
                        }
                    }
                }
            }
        }

        self.update_service(service).await?;

        Ok(())
    }
}

/// Retrieves rover.yaml file from disk, performs validation and returns object.
pub async fn get_config() -> Result<Configuration, Error> {
    if !Path::new(ROVER_CONFIG_FILE).exists() {
        // If there is no existing config, create a new file and write
        // an empty config to it.
        let empty_config = Configuration { enabled: vec![] };
        update_config(&empty_config)?;
    }

    let file_content =
        std::fs::read_to_string(ROVER_CONFIG_FILE).map_err(|_| Error::ConfigFileIO)?;

    let config: ValidatedConfiguration =
        serde_yaml::from_str::<Configuration>(&file_content)?.validate()?;

    Ok(config.0)
}

/// Returns the process which contains the fq.
pub fn get_proc(fq: FqBuf, processes: &Vec<Process>) -> Result<&Process, Error> {
    for p in processes {
        if p.fq == fq {
            return Ok(p);
        }
    }
    Err(Error::ProcessNotFound)
}

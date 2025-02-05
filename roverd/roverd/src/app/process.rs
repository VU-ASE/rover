use std::path::PathBuf;

use std::sync::Arc;

use tokio::{process::Child, sync::Mutex};

use openapi::models::*;

use crate::service::FqBuf;

#[derive(Debug, Clone)]
pub struct SpawnedProcess {
    pub fq: FqBuf,
    pub name: String,
    pub child: Arc<Mutex<Child>>,
}

/// A Process
#[derive(Debug, Clone)]
#[allow(dead_code)]
pub struct Process {
    pub fq: FqBuf,
    pub last_pid: Option<u32>,
    pub last_exit_code: i32,
    pub name: String,
    pub command: String,
    pub log_file: PathBuf,
    pub status: openapi::models::ProcessStatus,
    pub injected_env: String,
    pub faults: u32,
    pub start_time: i64,
}

#[derive(Debug)]
pub struct PipelineStats {
    pub status: PipelineStatus,
    pub last_start: Option<i64>,
    pub last_stop: Option<i64>,
    pub last_restart: Option<i64>,
}

use async_trait::async_trait;
use axum::extract::*;
use axum_extra::extract::{CookieJar, Multipart};
use bytes::Bytes;
use http::Method;
use serde::{Deserialize, Serialize};

use crate::{models, types::*};

#[derive(Debug, PartialEq, Serialize, Deserialize)]
#[must_use]
#[allow(clippy::large_enum_variant)]
pub enum LogsAuthorNameVersionGetResponse {
    /// The collection of logs
    Status200_TheCollectionOfLogs(Vec<String>),
    /// Error occurred
    Status400_ErrorOccurred(models::RoverdError),
    /// Unauthorized Access
    Status401_UnauthorizedAccess,
    /// Entity not found
    Status404_EntityNotFound,
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
#[must_use]
#[allow(clippy::large_enum_variant)]
pub enum PipelineGetResponse {
    /// Pipeline status and an array of processes
    Status200_PipelineStatusAndAnArrayOfProcesses(models::PipelineGet200Response),
    /// Error occurred
    Status400_ErrorOccurred(models::RoverdError),
    /// Unauthorized Access
    Status401_UnauthorizedAccess,
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
#[must_use]
#[allow(clippy::large_enum_variant)]
pub enum PipelinePostResponse {
    /// Operation was successful
    Status200_OperationWasSuccessful,
    /// Error occurred
    Status400_ErrorOccurred(models::RoverdError),
    /// Unauthorized Access
    Status401_UnauthorizedAccess,
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
#[must_use]
#[allow(clippy::large_enum_variant)]
pub enum PipelineStartPostResponse {
    /// Operation was successful
    Status200_OperationWasSuccessful,
    /// Error occurred
    Status400_ErrorOccurred(models::RoverdError),
    /// Unauthorized Access
    Status401_UnauthorizedAccess,
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
#[must_use]
#[allow(clippy::large_enum_variant)]
pub enum PipelineStopPostResponse {
    /// Operation was successful
    Status200_OperationWasSuccessful,
    /// Error occurred
    Status400_ErrorOccurred(models::RoverdError),
    /// Unauthorized Access
    Status401_UnauthorizedAccess,
}

/// Pipeline
#[async_trait]
#[allow(clippy::ptr_arg)]
pub trait Pipeline {
    /// Retrieve logs for any service. Logs from running or previously run services can be viewed and will be kept until rover reboot..
    ///
    /// LogsAuthorNameVersionGet - GET /logs/{author}/{name}/{version}
    async fn logs_author_name_version_get(
        &self,
        method: Method,
        host: Host,
        cookies: CookieJar,
        path_params: models::LogsAuthorNameVersionGetPathParams,
        query_params: models::LogsAuthorNameVersionGetQueryParams,
    ) -> Result<LogsAuthorNameVersionGetResponse, ()>;

    /// Retrieve pipeline status and process execution information.
    ///
    /// PipelineGet - GET /pipeline
    async fn pipeline_get(
        &self,
        method: Method,
        host: Host,
        cookies: CookieJar,
    ) -> Result<PipelineGetResponse, ()>;

    /// Set the services that are enabled in this pipeline, by specifying the fully qualified services.
    ///
    /// PipelinePost - POST /pipeline
    async fn pipeline_post(
        &self,
        method: Method,
        host: Host,
        cookies: CookieJar,
        body: Vec<models::PipelinePostRequestInner>,
    ) -> Result<PipelinePostResponse, ()>;

    /// Start the pipeline.
    ///
    /// PipelineStartPost - POST /pipeline/start
    async fn pipeline_start_post(
        &self,
        method: Method,
        host: Host,
        cookies: CookieJar,
    ) -> Result<PipelineStartPostResponse, ()>;

    /// Stop the pipeline.
    ///
    /// PipelineStopPost - POST /pipeline/stop
    async fn pipeline_stop_post(
        &self,
        method: Method,
        host: Host,
        cookies: CookieJar,
    ) -> Result<PipelineStopPostResponse, ()>;
}

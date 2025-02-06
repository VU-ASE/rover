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
    /// An error occurred
    Status400_AnErrorOccurred(models::GenericError),
    /// Entity not found
    Status404_EntityNotFound,
    /// Unauthorized access (you need to set the Authorization header with a valid username and password)
    Status401_UnauthorizedAccess,
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
#[must_use]
#[allow(clippy::large_enum_variant)]
pub enum PipelineGetResponse {
    /// Pipeline status and an array of processes
    Status200_PipelineStatusAndAnArrayOfProcesses(models::PipelineGet200Response),
    /// An error occurred
    Status400_AnErrorOccurred(models::GenericError),
    /// Unauthorized access (you need to set the Authorization header with a valid username and password)
    Status401_UnauthorizedAccess,
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
#[must_use]
#[allow(clippy::large_enum_variant)]
pub enum PipelinePostResponse {
    /// The pipeline was updated successfully
    Status200_ThePipelineWasUpdatedSuccessfully,
    /// The pipeline was not valid and could not be set
    Status400_ThePipelineWasNotValidAndCouldNotBeSet(models::PipelinePost400Response),
    /// Unauthorized access (you need to set the Authorization header with a valid username and password)
    Status401_UnauthorizedAccess,
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
#[must_use]
#[allow(clippy::large_enum_variant)]
pub enum PipelineStartPostResponse {
    /// The pipeline was started successfully. You can view its information with GET /pipeline
    Status200_ThePipelineWasStartedSuccessfully,
    /// An error occurred
    Status400_AnErrorOccurred(models::GenericError),
    /// Unauthorized access (you need to set the Authorization header with a valid username and password)
    Status401_UnauthorizedAccess,
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
#[must_use]
#[allow(clippy::large_enum_variant)]
pub enum PipelineStopPostResponse {
    /// The pipeline was stopped successfully. You can view its information with GET /pipeline
    Status200_ThePipelineWasStoppedSuccessfully,
    /// An error occurred
    Status400_AnErrorOccurred(models::GenericError),
    /// Unauthorized access (you need to set the Authorization header with a valid username and password)
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

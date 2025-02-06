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
pub enum RootGetResponse {
    /// The health and versioning information
    Status200_TheHealthAndVersioningInformation(models::StatusGet200Response),
    /// An error occurred
    Status400_AnErrorOccurred(models::GenericError),
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
#[must_use]
#[allow(clippy::large_enum_variant)]
pub enum ShutdownPostResponse {
    /// Rover shutdown successfully.
    Status200_RoverShutdownSuccessfully,
    /// An error occurred
    Status400_AnErrorOccurred(models::GenericError),
    /// Unauthorized access (you need to set the Authorization header with a valid username and password)
    Status401_UnauthorizedAccess,
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
#[must_use]
#[allow(clippy::large_enum_variant)]
pub enum StatusGetResponse {
    /// The health and versioning information
    Status200_TheHealthAndVersioningInformation(models::StatusGet200Response),
    /// An error occurred
    Status400_AnErrorOccurred(models::GenericError),
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
#[must_use]
#[allow(clippy::large_enum_variant)]
pub enum UpdatePostResponse {
    /// The roverd daemon process initiated a self-update successfully. You should expect the process to terminate and restart soon.
    Status200_TheRoverdDaemonProcessInitiatedASelf,
    /// An error occurred
    Status400_AnErrorOccurred(models::GenericError),
    /// Unauthorized access (you need to set the Authorization header with a valid username and password)
    Status401_UnauthorizedAccess,
}

/// Health
#[async_trait]
#[allow(clippy::ptr_arg)]
pub trait Health {
    /// Retrieve the health and versioning information.
    ///
    /// RootGet - GET /
    async fn root_get(
        &self,
        method: Method,
        host: Host,
        cookies: CookieJar,
    ) -> Result<RootGetResponse, ()>;

    /// Shutdown the rover..
    ///
    /// ShutdownPost - POST /shutdown
    async fn shutdown_post(
        &self,
        method: Method,
        host: Host,
        cookies: CookieJar,
    ) -> Result<ShutdownPostResponse, ()>;

    /// Retrieve the health and versioning information.
    ///
    /// StatusGet - GET /status
    async fn status_get(
        &self,
        method: Method,
        host: Host,
        cookies: CookieJar,
    ) -> Result<StatusGetResponse, ()>;

    /// Self-update the roverd daemon process.
    ///
    /// UpdatePost - POST /update
    async fn update_post(
        &self,
        method: Method,
        host: Host,
        cookies: CookieJar,
    ) -> Result<UpdatePostResponse, ()>;
}

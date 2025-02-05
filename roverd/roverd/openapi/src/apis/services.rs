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
pub enum FetchPostResponse {
    /// The service was uploaded successfully
    Status200_TheServiceWasUploadedSuccessfully(models::FetchPost200Response),
    /// An error occurred
    Status400_AnErrorOccurred(models::GenericError),
    /// Unauthorized access (you need to set the Authorization header with a valid username and password)
    Status401_UnauthorizedAccess,
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
#[must_use]
#[allow(clippy::large_enum_variant)]
pub enum FqnsGetResponse {
    /// Fully qualified services
    Status200_FullyQualifiedServices(Vec<models::FqnsGet200ResponseInner>),
    /// An error occurred
    Status400_AnErrorOccurred(models::GenericError),
    /// Unauthorized access (you need to set the Authorization header with a valid username and password)
    Status401_UnauthorizedAccess,
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
#[must_use]
#[allow(clippy::large_enum_variant)]
pub enum ServicesAuthorGetResponse {
    /// The list of services for the author
    Status200_TheListOfServicesForTheAuthor(Vec<String>),
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
pub enum ServicesAuthorServiceGetResponse {
    /// The list of versions for this author and service name
    Status200_TheListOfVersionsForThisAuthorAndServiceName(Vec<String>),
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
pub enum ServicesAuthorServiceVersionDeleteResponse {
    /// The service version was deleted successfully
    Status200_TheServiceVersionWasDeletedSuccessfully(
        models::ServicesAuthorServiceVersionDelete200Response,
    ),
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
pub enum ServicesAuthorServiceVersionGetResponse {
    /// A full description of the service at this version, with inputs, outputs and configuration
    Status200_AFullDescriptionOfTheServiceAtThisVersion(
        models::ServicesAuthorServiceVersionGet200Response,
    ),
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
pub enum ServicesAuthorServiceVersionPostResponse {
    /// The service was built successfully
    Status200_TheServiceWasBuiltSuccessfully,
    /// The build failed
    Status400_TheBuildFailed(models::ServicesAuthorServiceVersionPost400Response),
    /// Entity not found
    Status404_EntityNotFound,
    /// Unauthorized access (you need to set the Authorization header with a valid username and password)
    Status401_UnauthorizedAccess,
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
#[must_use]
#[allow(clippy::large_enum_variant)]
pub enum ServicesGetResponse {
    /// The list of authors
    Status200_TheListOfAuthors(Vec<String>),
    /// An error occurred
    Status400_AnErrorOccurred(models::GenericError),
    /// Unauthorized access (you need to set the Authorization header with a valid username and password)
    Status401_UnauthorizedAccess,
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
#[must_use]
#[allow(clippy::large_enum_variant)]
pub enum UploadPostResponse {
    /// The service was uploaded successfully
    Status200_TheServiceWasUploadedSuccessfully(models::FetchPost200Response),
    /// An error occurred
    Status400_AnErrorOccurred(models::GenericError),
    /// Unauthorized access (you need to set the Authorization header with a valid username and password)
    Status401_UnauthorizedAccess,
}

/// Services
#[async_trait]
#[allow(clippy::ptr_arg)]
pub trait Services {
    /// Fetches the zip file from the given URL and installs the service onto the filesystem.
    ///
    /// FetchPost - POST /fetch
    async fn fetch_post(
        &self,
        method: Method,
        host: Host,
        cookies: CookieJar,
        body: models::FetchPostRequest,
    ) -> Result<FetchPostResponse, ()>;

    /// Retrieve a list of all fully qualified services.
    ///
    /// FqnsGet - GET /fqns
    async fn fqns_get(
        &self,
        method: Method,
        host: Host,
        cookies: CookieJar,
    ) -> Result<FqnsGetResponse, ()>;

    /// Retrieve the list of parsable services for a specific author.
    ///
    /// ServicesAuthorGet - GET /services/{author}
    async fn services_author_get(
        &self,
        method: Method,
        host: Host,
        cookies: CookieJar,
        path_params: models::ServicesAuthorGetPathParams,
    ) -> Result<ServicesAuthorGetResponse, ()>;

    /// Retrieve the list of parsable service versions for a specific author and service.
    ///
    /// ServicesAuthorServiceGet - GET /services/{author}/{service}
    async fn services_author_service_get(
        &self,
        method: Method,
        host: Host,
        cookies: CookieJar,
        path_params: models::ServicesAuthorServiceGetPathParams,
    ) -> Result<ServicesAuthorServiceGetResponse, ()>;

    /// Delete a specific version of a service.
    ///
    /// ServicesAuthorServiceVersionDelete - DELETE /services/{author}/{service}/{version}
    async fn services_author_service_version_delete(
        &self,
        method: Method,
        host: Host,
        cookies: CookieJar,
        path_params: models::ServicesAuthorServiceVersionDeletePathParams,
    ) -> Result<ServicesAuthorServiceVersionDeleteResponse, ()>;

    /// Retrieve the status of a specific version of a service.
    ///
    /// ServicesAuthorServiceVersionGet - GET /services/{author}/{service}/{version}
    async fn services_author_service_version_get(
        &self,
        method: Method,
        host: Host,
        cookies: CookieJar,
        path_params: models::ServicesAuthorServiceVersionGetPathParams,
    ) -> Result<ServicesAuthorServiceVersionGetResponse, ()>;

    /// Build a fully qualified service version.
    ///
    /// ServicesAuthorServiceVersionPost - POST /services/{author}/{service}/{version}
    async fn services_author_service_version_post(
        &self,
        method: Method,
        host: Host,
        cookies: CookieJar,
        path_params: models::ServicesAuthorServiceVersionPostPathParams,
    ) -> Result<ServicesAuthorServiceVersionPostResponse, ()>;

    /// Retrieve the list of all authors that have parsable services. With these authors you can query further for services.
    ///
    /// ServicesGet - GET /services
    async fn services_get(
        &self,
        method: Method,
        host: Host,
        cookies: CookieJar,
    ) -> Result<ServicesGetResponse, ()>;

    /// Upload a new service or new version to the rover by uploading a ZIP file.
    ///
    /// UploadPost - POST /upload
    async fn upload_post(
        &self,
        method: Method,
        host: Host,
        cookies: CookieJar,
        body: Multipart,
    ) -> Result<UploadPostResponse, ()>;
}

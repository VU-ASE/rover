// use tracing::info;

use axum::async_trait;

use openapi::{apis::services::*, models};

use openapi::models::*;

use axum::extract::Host;
use axum::http::Method;
use axum_extra::extract::{CookieJar, Multipart};

use serde_json::value::RawValue;
use tracing::warn;

use crate::app::Roverd;
use crate::{rover_is_operating, warn_generic};
use rover_types::error::Error;
use rover_types::service::FqBuf;

#[async_trait]
impl Services for Roverd {
    /// Fetches the zip file from the given URL and installs the service onto the filesystem.
    /// `RoverState` - This function can run *only when dormant*
    /// TODO: fs_lock
    /// FetchPost - POST /fetch
    async fn fetch_post(
        &self,
        _method: Method,
        _host: Host,
        _cookies: CookieJar,
        body: models::FetchPostRequest,
    ) -> Result<FetchPostResponse, ()> {
        if let Some(rover_state) = self.try_get_dormant().await {
            let (fq_buf, invalidated_pipeline) = warn_generic!(
                self.app.fetch_service(&body, rover_state).await,
                FetchPostResponse
            );

            Ok(
                FetchPostResponse::Status200_TheServiceWasUploadedSuccessfully(
                    FetchPost200Response {
                        fq: FullyQualifiedService::from(fq_buf),
                        invalidated_pipeline,
                    },
                ),
            )
        } else {
            rover_is_operating!(FetchPostResponse)
        }
    }

    /// Upload a new service or new version to the rover by uploading a ZIP file.
    /// `RoverState` - This function can run *only when dormant*
    /// TODO: fs_lock
    /// UploadPost - POST /upload
    async fn upload_post(
        &self,
        _method: Method,
        _host: Host,
        _cookies: CookieJar,
        body: Multipart,
    ) -> Result<UploadPostResponse, ()> {
        if let Some(rover_state) = self.try_get_dormant().await {
            let (fq_buf, invalidated_pipeline) = warn_generic!(
                self.app.receive_upload(body, rover_state).await,
                UploadPostResponse
            );

            Ok(
                UploadPostResponse::Status200_TheServiceWasUploadedSuccessfully(
                    FetchPost200Response {
                        fq: FullyQualifiedService::from(fq_buf),
                        invalidated_pipeline,
                    },
                ),
            )
        } else {
            rover_is_operating!(UploadPostResponse)
        }
    }

    /// Retrieve the list of parsable service versions for a specific author and service.
    /// `RoverState` - This function can run *always*
    /// TODO: fs_lock
    /// ServicesAuthorServiceGet - GET /services/{author}/{service}
    async fn services_author_service_get(
        &self,
        _method: Method,
        _host: Host,
        _cookies: CookieJar,
        path_params: ServicesAuthorServiceGetPathParams,
    ) -> Result<ServicesAuthorServiceGetResponse, ()> {
        let versions = warn_generic!(
            self.app.get_versions(path_params).await,
            ServicesAuthorServiceGetResponse
        );

        Ok(ServicesAuthorServiceGetResponse::Status200_TheListOfVersionsForThisAuthorAndServiceName(versions))
    }

    /// Delete a specific version of a service.
    /// `RoverState` - This function can run *only when dormant*
    /// TODO: fs_lock
    /// ServicesAuthorServiceVersionDelete - DELETE /services/{author}/{service}/{version}
    async fn services_author_service_version_delete(
        &self,
        _method: Method,
        _host: Host,
        _cookies: CookieJar,
        path_params: ServicesAuthorServiceVersionDeletePathParams,
    ) -> Result<ServicesAuthorServiceVersionDeleteResponse, ()> {
        if let Some(rover_state) = self.try_get_dormant().await {
            let invalidated_pipeline = warn_generic!(
                self.app.delete_service(&path_params, rover_state).await,
                ServicesAuthorServiceVersionDeleteResponse
            );

            Ok(ServicesAuthorServiceVersionDeleteResponse::Status200_TheServiceVersionWasDeletedSuccessfully(
                ServicesAuthorServiceVersionDelete200Response {
                    invalidated_pipeline
                }
            ))
        } else {
            rover_is_operating!(ServicesAuthorServiceVersionDeleteResponse)
        }
    }

    /// Retrieve the status of a specific version of a service.
    /// `RoverState` - This function can run *always*
    /// TODO: fs_lock
    /// ServicesAuthorServiceVersionGet - GET /services/{author}/{service}/{version}
    async fn services_author_service_version_get(
        &self,
        _method: Method,
        _host: Host,
        _cookies: CookieJar,
        path_params: ServicesAuthorServiceVersionGetPathParams,
    ) -> Result<ServicesAuthorServiceVersionGetResponse, ()> {
        let fq = FqBuf::from(&path_params);

        let service = warn_generic!(
            self.app.get_service(fq.clone()).await,
            ServicesAuthorServiceVersionGetResponse
        );

        let built_services = self.app.built_services.read().await;
        let built_at = built_services.get(&fq).copied();

        let mut configuration = vec![];

        for c in service.0.configuration.iter() {
            configuration.push(models::ServicesAuthorServiceVersionGet200ResponseConfigurationInner {
                name: c.name.clone(),
                r#type: match c.value.clone() {
                    rover_validate::service::Value::Double(_) => "number".to_string(),
                    rover_validate::service::Value::String(_) => "string".to_string(),
                },
                tunable: c.tunable.unwrap_or(false),
                value: match c.value.clone() {
                    rover_validate::service::Value::Double(n) => models::ServicesAuthorServiceVersionGet200ResponseConfigurationInnerValue(RawValue::from_string(format!("{}", n)).unwrap()),
                    rover_validate::service::Value::String(s) => models::ServicesAuthorServiceVersionGet200ResponseConfigurationInnerValue(RawValue::from_string(format!("\"{}\"", s)).unwrap()),
                },
            })
        }

        Ok(
            ServicesAuthorServiceVersionGetResponse::Status200_AFullDescriptionOfTheServiceAtThisVersion(
                models::ServicesAuthorServiceVersionGet200Response {
                    inputs: service
                        .0
                        .inputs
                        .iter()
                        .map(|i| ServicesAuthorServiceVersionGet200ResponseInputsInner {
                            service: i.service.clone(),
                            streams: i.streams.clone(),
                        })
                        .collect::<Vec<_>>(),
                    built_at,
                    outputs: service.0.outputs,
                    configuration,
                },
            ),
        )
    }

    /// Build a fully qualified service version.
    /// `RoverState` - This function can run *only when dormant*
    /// TODO: fs_lock
    /// ServicesAuthorServiceVersionPost - POST /services/{author}/{service}/{version}
    async fn services_author_service_version_post(
        &self,
        _method: Method,
        _host: Host,
        _cookies: CookieJar,
        path_params: ServicesAuthorServiceVersionPostPathParams,
    ) -> Result<ServicesAuthorServiceVersionPostResponse, ()> {
        if let Some(rover_state) = self.try_get_dormant().await {
            let _ = if let Err(e) = self.app.build_service(path_params, rover_state).await {
                warn!("{:#?}", &e);
                let mut build_error_strings = vec![];
                match e {
                    Error::BuildLog(mut build_log) => {
                        build_error_strings.append(&mut build_log);
                    }
                    other_error => {
                        build_error_strings.push(format!("{:#?}", other_error));
                    }
                }
                let build_error = BuildError::new(build_error_strings);

                warn!("{:#?}", &build_error);
                // todo remove the unwraps and change to actual error
                let json_string = serde_json::to_string(&build_error).unwrap();
                let box_raw = serde_json::value::RawValue::from_string(json_string).unwrap();
                return Ok(
                    ServicesAuthorServiceVersionPostResponse::Status400_ErrorOccurred(
                        RoverdError::new("build".to_string(), RoverdErrorErrorValue(box_raw)),
                    ),
                );
            };
            Ok(ServicesAuthorServiceVersionPostResponse::Status200_OperationWasSuccessful)
        } else {
            rover_is_operating!(ServicesAuthorServiceVersionPostResponse)
        }
    }

    /// Retrieve the list of all authors that have parsable services. With these authors you can query further for services.
    /// `RoverState` - This function can run *always*
    /// TODO: fs_lock
    /// ServicesGet - GET /services
    async fn services_get(
        &self,
        _method: Method,
        _host: Host,
        _cookies: CookieJar,
    ) -> Result<ServicesGetResponse, ()> {
        let authors = warn_generic!(self.app.get_authors().await, ServicesGetResponse);
        Ok(ServicesGetResponse::Status200_TheListOfAuthors(authors))
    }

    /// Retrieve the list of parsable services for a specific author.
    /// `RoverState` - This function can run *always*
    /// TODO: fs_lock
    /// ServicesAuthorGet - GET /services/{author}
    async fn services_author_get(
        &self,
        _method: Method,
        _host: Host,
        _cookies: CookieJar,
        path_params: ServicesAuthorGetPathParams,
    ) -> Result<ServicesAuthorGetResponse, ()> {
        let services = warn_generic!(
            self.app.get_services(path_params).await,
            ServicesAuthorGetResponse
        );
        Ok(ServicesAuthorGetResponse::Status200_TheListOfServicesForTheAuthor(services))
    }

    /// Retrieve a list of all fully qualified services.
    /// `RoverState` - This function can run *always*
    /// TODO: fs_lock
    /// FqnsGet - GET /fqns
    async fn fqns_get(
        &self,
        _method: Method,
        _host: Host,
        _cookies: CookieJar,
    ) -> Result<FqnsGetResponse, ()> {
        let fqns = warn_generic!(self.app.get_fqns().await, FqnsGetResponse);
        Ok(FqnsGetResponse::Status200_FullyQualifiedServices(fqns))
    }

    /// Update service.yaml configuration values of a fully qualified service in-place.
    ///
    /// ServicesAuthorServiceVersionConfigurationPost - POST /services/{author}/{service}/{version}/configuration
    async fn services_author_service_version_configuration_post(
        &self,
        _method: Method,
        _host: Host,
        _cookies: CookieJar,
        path_params: models::ServicesAuthorServiceVersionConfigurationPostPathParams,
        body: Vec<models::ServicesAuthorServiceVersionConfigurationPostRequestInner>,
    ) -> Result<ServicesAuthorServiceVersionConfigurationPostResponse, ()> {
        if let Some(rover_state) = self.try_get_dormant().await {
            match self
                .app
                .update_service_config(&path_params, &body, rover_state)
                .await
            {
                Ok(_) => Ok(ServicesAuthorServiceVersionConfigurationPostResponse::Status200_OperationWasSuccessful),
                Err(e) => {
                    let error_msg = format!("{:#?}", e);
                    warn!("{:#?}", e);
                    let json_string = serde_json::to_string(&error_msg).unwrap();
                    let box_raw = serde_json::value::RawValue::from_string(json_string).unwrap();
                    Ok(
                        ServicesAuthorServiceVersionConfigurationPostResponse::Status400_ErrorOccurred(
                            RoverdError::new("generic".to_string(), RoverdErrorErrorValue(box_raw)),
                        ),
                    )
                },
            }
        } else {
            rover_is_operating!(ServicesAuthorServiceVersionConfigurationPostResponse)
        }
    }
}

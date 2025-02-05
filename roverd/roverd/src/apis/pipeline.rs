use axum::async_trait;
use axum::extract::Host;
use axum::http::Method;
use axum_extra::extract::CookieJar;

use openapi::apis::pipeline::*;
use openapi::models::*;

use tracing::warn;

use crate::constants::*;
use crate::{
    app::Roverd, error::Error, rover_is_dormant, rover_is_operating, service::FqBuf, warn_generic,
};

#[async_trait]
impl Pipeline for Roverd {
    /// Retrieve logs for any service. Logs from running or previously run services can be
    /// viewed and will be kept until rover reboot..
    /// `RoverState` - This function can run *always*
    /// TODO: fs_lock
    /// LogsAuthorNameVersionGet - GET /logs/{author}/{name}/{version}
    async fn logs_author_name_version_get(
        &self,
        _method: Method,
        _host: Host,
        _cookies: CookieJar,
        path_params: LogsAuthorNameVersionGetPathParams,
        query_params: LogsAuthorNameVersionGetQueryParams,
    ) -> Result<LogsAuthorNameVersionGetResponse, ()> {
        let fq = FqBuf::from(&path_params);
        let lines = query_params.lines.unwrap_or(DEFAULT_LOG_LINES) as usize;

        let logs = warn_generic!(
            self.app.get_service_logs(fq, lines).await,
            LogsAuthorNameVersionGetResponse
        );

        Ok(LogsAuthorNameVersionGetResponse::Status200_TheCollectionOfLogs(logs))
    }

    /// Retrieve pipeline status and process execution information.
    /// `RoverState` - This function can run *always*
    /// TODO: fs_lock
    /// PipelineGet - GET /pipeline
    async fn pipeline_get(
        &self,
        _method: Method,
        _host: Host,
        _cookies: CookieJar,
    ) -> Result<PipelineGetResponse, ()> {
        let enabled: Vec<PipelineGet200ResponseEnabledInner> =
            warn_generic!(self.app.get_pipeline().await, PipelineGetResponse);
        let stats = self.app.stats.read().await;
        Ok(
            PipelineGetResponse::Status200_PipelineStatusAndAnArrayOfProcesses(
                PipelineGet200Response {
                    status: stats.status,
                    last_start: stats.last_start,
                    last_stop: stats.last_stop,
                    last_restart: stats.last_restart,
                    enabled,
                },
            ),
        )
    }

    /// Set the services that are enabled in this pipeline,
    /// by specifying the fully qualified services.
    /// `RoverState` - This function can run *only when dormant*
    /// TODO: fs_lock
    /// PipelinePost - POST /pipeline
    async fn pipeline_post(
        &self,
        _method: Method,
        _host: Host,
        _cookies: CookieJar,
        body: Vec<PipelinePostRequestInner>,
    ) -> Result<PipelinePostResponse, ()> {
        if let Some(rover_state) = self.try_get_dormant().await {
            let _ = match self.app.set_pipeline(body, rover_state).await {
                Ok(a) => a,
                Err(e) => match e {
                    Error::Validation(val_errors) => {
                        let mut pipeline_errors = vec![];
                        let mut string_errors = vec![];

                        for val_error in val_errors {
                            match val_error {
                                rovervalidate::error::Error::PipelineValidationError(
                                    pipeline_validation_error,
                                ) => pipeline_errors.push(pipeline_validation_error),
                                e => string_errors.push(e.to_string()),
                            }
                        }

                        let mut unmet_streams = vec![];
                        let mut unmet_services = vec![];
                        let mut duplicate_service = vec![];

                        for i in pipeline_errors {
                            match i {
                                    rovervalidate::error::PipelineValidationError::UnmetDependencyError(unmet_dependency_error) => {
                                        match unmet_dependency_error {
                                            rovervalidate::error::UnmetDependencyError::UnmetStream(unmet_stream_error) => {
                                                unmet_streams.push(
                                                    UnmetStreamError {
                                                        source: Some(unmet_stream_error.source),
                                                        target: Some(unmet_stream_error.target),
                                                        stream: Some(unmet_stream_error.stream),
                                                    }
                                                );
                                            },
                                            rovervalidate::error::UnmetDependencyError::UnmetService(unmet_service_error) => {
                                                unmet_services.push(
                                                    UnmetServiceError {
                                                        source: Some(unmet_service_error.source),
                                                        target: Some(unmet_service_error.target)
                                                    }
                                                )
                                            },
                                        }
                                    },
                                    rovervalidate::error::PipelineValidationError::DuplicateServiceError(s) => {
                                        duplicate_service.push(openapi::models::DuplicateServiceError(s));
                                    },
                                }
                        }

                        let string_errors = if !string_errors.is_empty() {
                            Some(string_errors.concat().to_string())
                        } else {
                            None
                        };

                        let unmet_streams = if !unmet_streams.is_empty() {
                            Some(unmet_streams)
                        } else {
                            None
                        };

                        let unmet_services = if !unmet_services.is_empty() {
                            Some(unmet_services)
                        } else {
                            None
                        };

                        let duplicate_service = if !duplicate_service.is_empty() {
                            Some(duplicate_service)
                        } else {
                            None
                        };

                        return Ok(
                            PipelinePostResponse::Status400_ThePipelineWasNotValidAndCouldNotBeSet(
                                PipelinePost400Response {
                                    message: string_errors,
                                    validation_errors: PipelinePost400ResponseValidationErrors {
                                        unmet_streams,
                                        unmet_services,
                                        duplicate_service,
                                    },
                                },
                            ),
                        );
                    }
                    all_other_errors => {
                        return Ok(
                            PipelinePostResponse::Status400_ThePipelineWasNotValidAndCouldNotBeSet(
                                PipelinePost400Response {
                                    message: Some(format!("{:?}", all_other_errors)),
                                    validation_errors: PipelinePost400ResponseValidationErrors {
                                        unmet_streams: None,
                                        unmet_services: None,
                                        duplicate_service: None,
                                    },
                                },
                            ),
                        );
                    }
                },
            };

            Ok(PipelinePostResponse::Status200_ThePipelineWasUpdatedSuccessfully)
        } else {
            Ok(
                PipelinePostResponse::Status400_ThePipelineWasNotValidAndCouldNotBeSet(
                    PipelinePost400Response {
                        message: Some("".to_string()),
                        validation_errors: PipelinePost400ResponseValidationErrors {
                            unmet_streams: None,
                            unmet_services: None,
                            duplicate_service: None,
                        },
                    },
                ),
            )
        }
    }

    /// Start the pipeline.
    /// `RoverState` - This function can run *on when dormant*
    /// TODO: fs_lock
    /// PipelineStartPost - POST /pipeline/start
    async fn pipeline_start_post(
        &self,
        _method: Method,
        _host: Host,
        _cookies: CookieJar,
    ) -> Result<PipelineStartPostResponse, ()> {
        if let Some(rover_state) = self.try_get_dormant().await {
            let _ = match self.app.start(rover_state).await {
                Ok(data) => data,
                Err(e) => {
                    warn!("{:#?}", e);
                    return Ok(PipelineStartPostResponse::Status400_AnErrorOccurred(
                        GenericError {
                            message: Some(format!("{:?}", e)),
                            code: Some(1),
                        },
                    ));
                }
            };
            Ok(PipelineStartPostResponse::Status200_ThePipelineWasStartedSuccessfully)
        } else {
            rover_is_operating!(PipelineStartPostResponse);
        }
    }

    /// Stop the pipeline.
    /// `RoverState` - This function can run *only when operating*
    /// TODO: fs_lock
    /// PipelineStopPost - POST /pipeline/stop
    async fn pipeline_stop_post(
        &self,
        _method: Method,
        _host: Host,
        _cookies: CookieJar,
    ) -> Result<PipelineStopPostResponse, ()> {
        if let Some(rover_state) = self.try_get_operating().await {
            let _ = match self.app.stop(rover_state).await {
                Ok(data) => data,
                Err(e) => {
                    warn!("{:#?}", e);
                    return Ok(PipelineStopPostResponse::Status400_AnErrorOccurred(
                        GenericError {
                            message: Some(format!("{:?}", e)),
                            code: Some(1),
                        },
                    ));
                }
            };
            Ok(PipelineStopPostResponse::Status200_ThePipelineWasStoppedSuccessfully)
        } else {
            rover_is_dormant!(PipelineStopPostResponse)
        }
    }
}

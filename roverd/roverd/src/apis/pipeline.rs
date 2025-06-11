use axum::async_trait;
use axum::extract::Host;
use axum::http::Method;
use axum_extra::extract::CookieJar;

use openapi::apis::pipeline::*;
use openapi::models::*;
use tracing::warn;

use crate::{app::Roverd, rover_is_dormant, rover_is_operating, warn_generic};
use rover_constants::*;
use rover_types::error::Error;
use rover_types::service::FqBuf;

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
        let (enabled, some_service): (
            Vec<PipelineGet200ResponseEnabledInner>,
            Option<FullyQualifiedService>,
        ) = warn_generic!(self.app.get_pipeline().await, PipelineGetResponse);
        let stats = self.app.stats.read().await;

        let stopping_service = if let Some(service) = some_service {
            Some(PipelineGet200ResponseStoppingService { fq: Some(service) })
        } else {
            None
        };

        Ok(
            PipelineGetResponse::Status200_PipelineStatusAndAnArrayOfProcesses(
                PipelineGet200Response {
                    status: stats.status,
                    last_start: stats.last_start,
                    last_stop: stats.last_stop,
                    last_restart: stats.last_restart,
                    stopping_service,
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
                                rover_validate::error::Error::PipelineValidationError(
                                    pipeline_validation_error,
                                ) => pipeline_errors.push(pipeline_validation_error),
                                e => string_errors.push(e.to_string()),
                            }
                        }

                        let mut unmet_streams = vec![];
                        let mut unmet_services = vec![];
                        let mut duplicate_services = vec![];
                        let mut duplicate_aliases = vec![];
                        let mut aliases_in_use = vec![];

                        for i in pipeline_errors {
                            match i {
                                    rover_validate::error::PipelineValidationError::UnmetDependencyError(unmet_dependency_error) => {
                                        match unmet_dependency_error {
                                            rover_validate::error::UnmetDependencyError::UnmetStream(unmet_stream_error) => {
                                                unmet_streams.push(
                                                    UnmetStreamError {
                                                        source: unmet_stream_error.source,
                                                        target: unmet_stream_error.target,
                                                        stream: unmet_stream_error.stream,
                                                    }
                                                );
                                            },
                                            rover_validate::error::UnmetDependencyError::UnmetService(unmet_service_error) => {
                                                unmet_services.push(
                                                    UnmetServiceError {
                                                        source: unmet_service_error.source,
                                                        target: unmet_service_error.target,
                                                    }
                                                )
                                            },
                                        }
                                    },
                                    rover_validate::error::PipelineValidationError::DuplicateServiceError(s) => {
                                        duplicate_services.push(s);
                                    },
                                    rover_validate::error::PipelineValidationError::DuplicateAliasError(s) => {
                                        duplicate_aliases.push(s)
                                    },
                                    rover_validate::error::PipelineValidationError::AliasInUseAsNameError(s) => {
                                        aliases_in_use.push(s)
                                    },
                                }
                        }

                        let pipeline_error =
                            PipelineSetError::new(PipelineSetErrorValidationErrors {
                                unmet_streams,
                                unmet_services,
                                duplicate_services,
                                duplicate_aliases,
                                aliases_in_use,
                            });

                        // todo remove the unwraps and change to actual error
                        warn!("{:#?}", &pipeline_error);
                        let json_string = serde_json::to_string(&pipeline_error).unwrap();
                        let box_raw =
                            serde_json::value::RawValue::from_string(json_string).unwrap();
                        return Ok(PipelinePostResponse::Status400_ErrorOccurred(
                            RoverdError::new(
                                "pipeline_set".to_string(),
                                RoverdErrorErrorValue(box_raw),
                            ),
                        ));
                    }
                    some_error => {
                        let some_generic_error = GenericError::new(format!("{:#?}", some_error), 1);
                        // todo remove the unwraps and change to actual error
                        warn!("{:#?}", &some_generic_error);
                        let json_string = serde_json::to_string(&some_generic_error).unwrap();
                        let box_raw =
                            serde_json::value::RawValue::from_string(json_string).unwrap();
                        return Ok(PipelinePostResponse::Status400_ErrorOccurred(
                            RoverdError::new("generic".to_string(), RoverdErrorErrorValue(box_raw)),
                        ));
                    }
                },
            };

            Ok(PipelinePostResponse::Status200_OperationWasSuccessful)
        } else {
            rover_is_operating!(PipelinePostResponse)
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
            let _ = warn_generic!(self.app.start(rover_state).await, PipelineStartPostResponse);
            Ok(PipelineStartPostResponse::Status200_OperationWasSuccessful)
        } else {
            rover_is_operating!(PipelineStartPostResponse)
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
            let _ = warn_generic!(self.app.stop(rover_state).await, PipelineStopPostResponse);
            Ok(PipelineStopPostResponse::Status200_OperationWasSuccessful)
        } else {
            rover_is_dormant!(PipelineStopPostResponse)
        }
    }
}

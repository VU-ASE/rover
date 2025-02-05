use std::sync::Arc;

use anyhow::{anyhow, Context};
use axum::extract::{DefaultBodyLimit, Request, State};
use axum::http::{self, StatusCode};
use axum::middleware::{self, Next};
use axum::response::Response;
use base64::Engine;
use daemons::DaemonManager;
use openapi::models::DaemonStatus;
use tower_http::cors::CorsLayer;
use tracing::{error, info, warn};

mod apis;
mod app;
mod command;
mod constants;
mod error;
mod log;
mod util;

#[cfg(test)]
mod test;

use app::*;
use constants::*;
use error::Error::*;

/// Not ideal, but an error wrapper work around since middleware::from_fn_with_state expects
/// Result<Response, StatusCode>. But ideally, we want to use custom Error to propogate our custom
/// Error type as mmuch as possible.
async fn auth_wrapper(
    state: State<Roverd>,
    req: Request,
    next: Next,
) -> Result<Response, StatusCode> {
    match auth(state, req, next).await {
        Ok(response) => Ok(response),
        Err(e) => {
            warn!("Unauthorized or bad request: {:?}", e);
            match e {
                Http(status_code) => Err(status_code),
                _ => Err(StatusCode::BAD_REQUEST),
            }
        }
    }
}

/// Performs password check to hashed password stored on disk.
fn check_auth(state: &Roverd, auth_str: &str) -> Result<(), error::Error> {
    let (user, password) = auth_str
        .split_once(':')
        .ok_or(Http(StatusCode::BAD_REQUEST))?;

    let stored_hash = sha256::digest(password);

    if let Some(hash) = &state.info.password {
        if user == state.info.username && hash == &stored_hash {
            return Ok(());
        }
    }

    warn!("Unauthorized access denied");

    Err(Http(StatusCode::UNAUTHORIZED))
}

/// Main authentication logic requires authenticated requests for all endpoints
/// except for "/status".
async fn auth(
    State(state): State<Roverd>,
    req: Request,
    next: Next,
) -> Result<Response, error::Error> {
    info!("{} {}", req.method(), *req.uri());

    // the /status and / endpoints do not require authentication, all others do.
    if *req.uri() != *"/status" && *req.uri() != *"/" {
        if state.info.status == DaemonStatus::Operational {
            let auth_header = req
                .headers()
                .get(http::header::AUTHORIZATION)
                .and_then(|header| header.to_str().ok())
                .ok_or(Http(StatusCode::UNAUTHORIZED))?;

            let basic_auth: Vec<&str> = auth_header.split(' ').collect();

            if basic_auth.len() != 2 || basic_auth[0] != "Basic" {
                // warn!("request is missing basic auth header");
                return Err(Context(anyhow!(
                    "basic_auth header != Basic or auth header did not contain exactly two items"
                )));
            }

            let base64_data = basic_auth[1];

            let raw_bytes = base64::prelude::BASE64_STANDARD
                .decode(base64_data)
                .map_err(|_| Http(StatusCode::BAD_REQUEST))?;

            let auth_str =
                core::str::from_utf8(&raw_bytes).map_err(|_| Http(StatusCode::BAD_REQUEST))?;

            // Returns early if authentication fails
            check_auth(&state, auth_str)?;
        } else {
            warn!("could not handle request since roverd not operational");
            return Err(error::Error::RoverdNotOperational);
        }
    }

    // Pass the request on to the request handlers.
    let response = next.run(req).await;
    Ok(response)
}

/// Entry of program, initializes logging and constructs app state used by axum router.
#[tokio::main]
async fn main() -> Result<(), error::Error> {
    log::init();
    info!("logging initialized");

    // All app initialization happens in new()
    let mut roverd = Roverd::new().await?;

    // If our daemons are downloaded
    match DaemonManager::new().await {
        Ok(d) => {
            let daemon_manager = Arc::new(d);
            let dm_arc = Arc::clone(&daemon_manager);

            // Hand-off to axum with a max upload limit of 100MB
            let router = openapi::server::new(roverd.clone())
                .layer(middleware::from_fn_with_state(roverd, auth_wrapper))
                .layer(CorsLayer::permissive())
                .layer(DefaultBodyLimit::max(100000000));

            let listener = tokio::net::TcpListener::bind(LISTEN_ADDRESS).await.unwrap();

            info!("listening on {}", LISTEN_ADDRESS);

            axum::serve(listener, router)
                .with_graceful_shutdown(async move {
                    dm_arc.shutdown_signal().await;
                })
                .await
                .context("axum error")?;
        }
        Err(e) => {
            error!("unable to start daemons: {:?}", e);
            roverd.info.status = DaemonStatus::Unrecoverable;

            // Hand-off to axum with a max upload limit of 100MB
            let router = openapi::server::new(roverd.clone())
                .layer(middleware::from_fn_with_state(roverd, auth_wrapper))
                .layer(CorsLayer::permissive())
                .layer(DefaultBodyLimit::max(100000000));

            let listener = tokio::net::TcpListener::bind(LISTEN_ADDRESS).await.unwrap();

            axum::serve(listener, router).await.context("axum error")?;
        }
    }

    info!("roverd exiting");
    Ok(())
}

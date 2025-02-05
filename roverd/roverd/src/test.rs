use anyhow::Context;
use error::Error;
use std::path::Path;

use once_cell::sync::Lazy;
use serde_json::json;
use tokio::sync::OnceCell;

use openapi::models::*;
use reqwest::multipart;
use reqwest::{header, Client, Response};

use base64::{engine::general_purpose::STANDARD, Engine as _};

use crate::*;

static SERVER: Lazy<OnceCell<Result<(), Error>>> = Lazy::new(OnceCell::new);

async fn start_server() -> Result<(), Error> {
    log::init();
    info!("logging initialized");

    let rover_state = Roverd::new().await?;

    let router = openapi::server::new(rover_state.clone())
        .layer(middleware::from_fn_with_state(
            rover_state.clone(),
            auth_wrapper,
        ))
        .layer(CorsLayer::permissive());

    let listener = tokio::net::TcpListener::bind(LISTEN_ADDRESS).await.unwrap();

    tokio::spawn(async move {
        axum::serve(listener, router).await.unwrap();
    });

    Ok(())
}

async fn ensure_server_is_started() -> &'static Result<(), Error> {
    // TODO this doesn't seem to be working as expected, which is why
    // all tests are in one funciton.
    SERVER.get_or_init(|| async { start_server().await }).await
}

fn create_authenticated_client() -> Client {
    Client::builder()
        .default_headers({
            let mut headers = header::HeaderMap::new();
            let auth_value = STANDARD.encode("debix:debix");
            headers.insert(
                header::AUTHORIZATION,
                header::HeaderValue::from_str(&format!("Basic {}", auth_value)).unwrap(),
            );
            headers
        })
        .build()
        .expect("Failed to create authenticated client")
}

async fn check_status(client: &Client) -> Result<(), Error> {
    let response = client.get("http://localhost/status").send().await?;

    assert!(response.status().is_success());

    let response_string = response.text().await?;
    let r: StatusGet200Response = serde_json::from_str(&response_string)?;

    assert_eq!(r.status, DaemonStatus::Operational);
    assert_eq!(r.error_message, None);
    assert_eq!(r.version, String::from("0.1.0"));
    assert_eq!(r.os, String::from("Debian 12.0.0 [64-bit]"));
    assert_eq!(r.rover_id, Some(13));
    assert_eq!(r.rover_name, Some(String::from("bunny")));

    Ok(())
}

async fn delete_service(client: &Client, s: &str) -> Response {
    client.delete(s).send().await.unwrap()
}

async fn upload_file(client: &Client, s: &str) -> Result<Response, Error> {
    let file_path = Path::new(s);
    let form = multipart::Form::new()
        .file("file", file_path)
        .await
        .with_context(|| format!("failed to create multipart form from {:?}", file_path))?;

    let response = client
        .post("http://localhost:/upload") // Replace with your actual endpoint
        .multipart(form)
        .send()
        .await?;
    Ok(response)
}

async fn fetch_file(client: &Client, s: &str) -> Result<Response, Error> {
    let body = json!({
        "url": s
    });
    let response = client
        .post("http://localhost:/fetch") // Replace with your actual endpoint
        .json(&body)
        .send()
        .await?;
    Ok(response)
}

#[tokio::test]
async fn all() -> Result<(), Error> {
    ensure_server_is_started().await;
    let client = create_authenticated_client();

    check_status(&client).await?;

    // First, delete any example files, regardless if they are there or not
    let _ = delete_service(&client, "http://localhost/services/vu-ase/imaging/1.0.0").await;
    let _ = delete_service(&client, "http://localhost/services/vu-ase/controller/1.0.0").await;
    let _ = delete_service(&client, "http://localhost/services/vu-ase/actuator/1.0.0").await;

    // Upload three services, this must succeed.
    let r = upload_file(&client, "examples/actuator.zip").await?;
    assert!(r.status().is_success());
    let r = upload_file(&client, "examples/imaging.zip").await?;
    assert!(r.status().is_success());
    let r = upload_file(&client, "examples/controller.zip").await?;
    assert!(r.status().is_success());

    // Get all authors, must succeed
    let r = client.get("http://localhost/services").send().await?;
    assert!(r.status().is_success());
    let s = r.text().await?;
    let data: Vec<String> = serde_json::from_str(&s)?;
    assert_eq!(data, vec!["vu-ase"]);

    // Get all services, must succeed
    let r = client
        .get("http://localhost/services/vu-ase")
        .send()
        .await?;
    assert!(r.status().is_success());
    let s = r.text().await?;
    let data: Vec<String> = serde_json::from_str(&s)?;
    assert!(data.contains(&"actuator".to_string()));
    assert!(data.contains(&"imaging".to_string()));
    assert!(data.contains(&"controller".to_string()));

    // Get all versions, must succeed
    let r = client
        .get("http://localhost/services/vu-ase/imaging")
        .send()
        .await?;
    assert!(r.status().is_success());
    let s = r.text().await?;
    let data: Vec<String> = serde_json::from_str(&s)?;
    assert_eq!(data, vec!["1.0.0",]);

    let r = client
        .get("http://localhost/services/vu-ase/controller")
        .send()
        .await?;
    assert!(r.status().is_success());
    let s = r.text().await?;
    let data: Vec<String> = serde_json::from_str(&s)?;
    assert_eq!(data, vec!["1.0.0",]);

    let r = client
        .get("http://localhost/services/vu-ase/actuator")
        .send()
        .await?;
    assert!(r.status().is_success());
    let s = r.text().await?;
    let data: Vec<String> = serde_json::from_str(&s)?;
    assert_eq!(data, vec!["1.0.0",]);

    // Delete controller
    let r = delete_service(&client, "http://localhost/services/vu-ase/controller/1.0.0").await;
    assert!(r.status().is_success());
    let s = r.text().await?;
    let data: ServicesAuthorServiceVersionDelete200Response = serde_json::from_str(&s)?;

    // No pipeline has been set, must be false
    assert_eq!(data.invalidated_pipeline, false);

    // Fetch imaging from downloads
    let r = fetch_file(&client, "https://downloads.ase.vu.nl/api/controller/v1.2.0").await?;
    assert!(r.status().is_success());

    // Delete the the downloaded service
    let r = delete_service(&client, "http://localhost/services/vu-ase/controller/1.0.0").await;
    assert!(r.status().is_success());
    let s = r.text().await?;
    let data: ServicesAuthorServiceVersionDelete200Response = serde_json::from_str(&s)?;

    // Again, no pipeline has been set.
    assert_eq!(data.invalidated_pipeline, false);

    // Upload local controller again
    let r = upload_file(&client, "examples/controller.zip").await?;
    assert!(r.status().is_success());

    // Make pipeline

    Ok(())
}

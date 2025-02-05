use anyhow::{anyhow, Context, Result};
use reqwest::Client;
use std::fs::{File, OpenOptions};
use std::os::unix::fs::chown;
use std::time::Duration;
use std::{
    fs,
    io::{self, Read, Write},
    path::{Path, PathBuf},
};

use axum::http::StatusCode;

use rovervalidate::config::{Configuration, Validate};

use rovervalidate::service::Service;
use tracing::{info, warn};

use crate::error::Error;
use crate::service::FqBuf;

use crate::constants::*;

/// Copies all files from source to destination recursively and sets ownership of all
/// desitnation files to "debix:debix".
pub fn copy_recursively(source: impl AsRef<Path>, destination_dir: impl AsRef<Path>) -> Result<()> {
    fs::create_dir_all(&destination_dir)?;
    chown(&destination_dir, DEBIX_UID, DEBIX_GID).with_context(|| {
        format!(
            "failed to set the ownership of directory: {:?}",
            destination_dir.as_ref()
        )
    })?;

    for entry in fs::read_dir(source)? {
        let entry = entry?;
        let filetype = entry.file_type()?;
        let destination_file = destination_dir.as_ref().join(entry.file_name());

        // Make sure all files copied over have debix:debix permissions so
        // that the build command succeeds
        if filetype.is_dir() {
            copy_recursively(entry.path(), &destination_file)?;
        } else {
            fs::copy(entry.path(), &destination_file)?;
            chown(&destination_file, DEBIX_UID, DEBIX_GID).with_context(|| {
                format!(
                    "failed to set the ownership of file: {:?}",
                    &destination_file
                )
            })?;
        }
    }
    Ok(())
}

/// Extracts the contents of the zip file into the directory at destination_dir.
pub fn extract_zip(zip_file: &str, destination_dir: &str) -> Result<(), Error> {
    std::fs::create_dir_all(destination_dir)
        .with_context(|| format!("failed to create dirs {}", destination_dir))?;

    let mut file =
        fs::File::open(zip_file).with_context(|| format!("failed to open {}", zip_file))?;
    let mut bytes: Vec<u8> = Vec::new();
    file.read_to_end(&mut bytes)
        .with_context(|| format!("failed to read to end of {}", zip_file))?;

    let target = PathBuf::from(destination_dir);

    let data_cursor = io::Cursor::new(bytes);
    let mut zip = zip::ZipArchive::new(data_cursor)?;

    zip.extract(target)?;

    Ok(())
}

/// Makes sure the directories for a given service exist. If there is an
/// existing service at a given path it will delete it and prepare it such
/// that the new service can be safely moved in place.
// fn prepare_dirs(author: &str, name: &str, version: &str) -> Result<String, Error> {
fn prepare_dirs(fq: &FqBuf) -> Result<String, Error> {
    // Construct the full path
    let full_path_string = fq.dir().clone();
    let full_path = PathBuf::from(full_path_string.clone());

    // Ensure the directory exists
    std::fs::create_dir_all(full_path.clone())
        .with_context(|| format!("failed to create dirs {:?}", full_path))?;

    // If it already existed and it contained old contents, remove them.
    std::fs::remove_dir_all(full_path.as_path())
        .with_context(|| format!("failed to remove path {:?}", full_path))?;

    Ok(full_path_string)
}

/// Downloads the vu-ase service from the downloads page and creates a zip file
/// /tmp/name-version.zip.
pub async fn download_service(url: &String) -> Result<(), Error> {
    info!("Downloading: {}", url);
    let client = Client::new();

    let res = client
        .get(url)
        .timeout(Duration::from_secs(DOWNLOAD_TIMEOUT))
        .send()
        .await;

    match res {
        Ok(res) => {
            if res.status() != StatusCode::OK {
                let resp: axum::http::StatusCode = res.status();

                let fail_msg = format!("failed to download {}", url);
                match res.status() {
                    StatusCode::NOT_FOUND => {
                        return Err(Error::ServiceNotFound(format!(
                            "HTTP ({}) - {}",
                            StatusCode::NOT_FOUND,
                            &fail_msg
                        )))
                    }
                    StatusCode::BAD_REQUEST => {
                        return Err(Error::ServiceNotFound(format!(
                            "bad request ({}) - {}",
                            StatusCode::BAD_REQUEST,
                            &fail_msg
                        )))
                    }
                    StatusCode::FORBIDDEN => return Err(Error::Http(StatusCode::FORBIDDEN)),
                    _ => return Err(Error::Http(resp)),
                }
            }
            std::fs::remove_file(ZIP_FILE).ok();

            let mut file = std::fs::File::create(ZIP_FILE)
                .with_context(|| format!("failed to create {}", ZIP_FILE))?;

            let bytes = res.bytes().await?;

            file.write_all(&bytes)
                .with_context(|| format!("failed to failed to write to {}", ZIP_FILE))?;
            Ok(())
        }
        Err(err) => {
            if err.is_timeout() {
                let msg = format!("request timed out after {} seconds", DOWNLOAD_TIMEOUT);
                Err(Error::Context(anyhow!(msg)))
            } else {
                let msg = format!("request failed {:?}", err);
                Err(Error::Context(anyhow!(msg)))
            }
        }
    }
}

/// Downloads a service to /tmp and moves it into the correct place on disk.
/// There shouldn't be any directories or files in the unique path of the service,
/// however if there are, they will get deleted to make space.
pub async fn download_and_install_service(url: &String, is_daemon: bool) -> Result<FqBuf, Error> {
    download_service(url).await?;
    let mut fq = extract_fq_from_zip().await?;
    fq.is_daemon = is_daemon;
    install_service(&fq).await?;
    Ok(fq)
}

/// Attempts to unzip and read out the service in the temporary directory
/// returns the FqBuf on success.
pub async fn extract_fq_from_zip() -> Result<FqBuf, Error> {
    // Clear the destination directory, no matter if it fails
    let _ = std::fs::remove_dir_all(UNZIPPED_DIR);

    // Create directory, this must not fail
    std::fs::create_dir_all(UNZIPPED_DIR)
        .with_context(|| format!("failed to create {}", UNZIPPED_DIR))?;

    // Unpack the downloaded service and validate it.
    extract_zip(ZIP_FILE, UNZIPPED_DIR)?;

    // Read contents and
    let service_contents = std::fs::read_to_string(format!("{}/service.yaml", UNZIPPED_DIR))
        .map_err(|_| Error::ServiceYamlNotFoundInDownload)?;
    let service = serde_yaml::from_str::<Service>(&service_contents)?.validate()?;

    let fq = FqBuf::from(service);
    Ok(fq)
}

/// Expects a zipfile to be ready at ZIP_FILE, extract it and install it. Parses the service.yaml
/// and install contents into the correct location on disk.
pub async fn install_service(fq: &FqBuf) -> Result<(), Error> {
    // Deletes any existing files/dirs that are on the /author/name/version path
    // Makes sure the directories exist.
    let full_path = prepare_dirs(fq)?;

    // Copy contents into place
    copy_recursively(UNZIPPED_DIR, &full_path).with_context(|| {
        format!(
            "failed to copy contents from {} to {}",
            UNZIPPED_DIR, full_path
        )
    })?;

    chown(&full_path, DEBIX_UID, DEBIX_GID)
        .with_context(|| format!("failed to set the ownership of {}", &full_path))?;

    Ok(())
}

pub fn list_dir_contents(added_path: &str) -> Result<Vec<String>, Error> {
    let path_string = format!("{}/{}", ROVER_DIR, added_path);
    let paths = fs::read_dir(&path_string)
        .map_err(|_| Error::ServiceNotFound(format!("Could not find {} on disk", path_string)))?;
    let mut contents: Vec<String> = vec![];

    for path in paths {
        contents.push(
            path.with_context(|| "failed to unpack direntry".to_string())?
                .file_name()
                .to_os_string()
                .into_string()?,
        )
    }

    Ok(contents)
}

/// Updates the config and creates the file if it doesn't exist
pub fn update_config(config: &Configuration) -> Result<(), Error> {
    let contents = serde_yaml::to_string(&config)?;

    std::fs::create_dir_all(ROVER_CONFIG_DIR).map_err(|_| Error::ConfigFileIO)?;

    let mut file = OpenOptions::new()
        .write(true)
        .create(true)
        .truncate(true)
        .open(ROVER_CONFIG_FILE)
        .map_err(|_| Error::ConfigFileIO)?;

    file.write_all(contents.as_bytes())
        .map_err(|_| Error::ConfigFileIO)?;

    Ok(())
}

pub fn create_log_file(log_path: &PathBuf) -> Result<File, Error> {
    let path = std::path::Path::new(log_path);
    if let Some(parent_dir) = path.parent() {
        if !parent_dir.exists() {
            info!("creating parent dir of logfile: {:?}", &parent_dir);
            std::fs::create_dir_all(parent_dir)
                .with_context(|| format!("failed to create {:?}", parent_dir))?;
        }
    }

    let log_file = OpenOptions::new()
        .read(true)
        .append(true)
        .create(true)
        .open(log_path.clone())
        .with_context(|| format!("failed to create/open {:?}", log_path))?;

    Ok(log_file)
}

/// Given an array of Strings, it will return the latest.
fn get_latest_version(versions: &[String]) -> Option<String> {
    versions
        .iter()
        .filter_map(|v| semver::Version::parse(v).ok())
        .max()
        .map(|v| v.to_string())
}

/// Checks the filesystem for the latest daemon for a given author & service_name
/// returns an error if it can't find it. If it can't find one, then this fails
/// the init sequence and the rover is not operational.
pub fn find_latest_daemon(author: &str, name: &str) -> Result<FqBuf, Error> {
    // List the directory with all versions
    let daemon_path = PathBuf::from(format!("{}/{}/{}", DAEMON_DIR, author, name));

    // Collect all the entries of the daemon's directory and check which one is the
    // newest
    let mut versions = vec![];

    for entry in fs::read_dir(&daemon_path)
        .with_context(|| format!("failed to read daemon directory: {:?}", &daemon_path))?
    {
        let entry = entry.context("failed to unpack directory entry")?;
        let filetype = entry
            .file_type()
            .with_context(|| format!("could not fetch file metadata of {:?}", entry.path()))?;

        // Make sure all files copied over have debix:debix permissions so
        // that the build command succeeds
        if filetype.is_dir() {
            versions.push(entry.file_name().to_string_lossy().into_owned());
        } else {
            warn!("found non-directory in {:?}", entry.path());
        }
    }

    if let Some(latest_version) = get_latest_version(&versions) {
        Ok(FqBuf::new_daemon(author, name, &latest_version))
    } else {
        Err(Error::Context(anyhow!(
            "Could not find or parse any valid semver versions in {:?}",
            daemon_path
        )))
    }
}

#[macro_export]
macro_rules! warn_generic {
    ($expr:expr, $error_type:ty) => {{
        match $expr {
            Ok(data) => data,
            Err(e) => {
                warn!("{:#?}", e);
                return Ok(<$error_type>::Status400_AnErrorOccurred(GenericError {
                    message: Some(format!("{:?}", e)),
                    code: Some(1),
                }));
            }
        }
    }};
}

#[macro_export]
macro_rules! error_generic {
    ($expr:expr, $error_type:ty) => {{
        match $expr {
            Ok(data) => data,
            Err(e) => {
                error!("{:#?}", e);
                return Ok(<$error_type>::Status400_AnErrorOccurred(GenericError {
                    message: Some(format!("{:?}", e)),
                    code: Some(1),
                }));
            }
        }
    }};
}

#[macro_export]
macro_rules! rover_is_dormant {
    ($error_type:ty) => {{
        let msg = "unable to perform request, rover is not running";
        warn!(msg);
        return Ok(<$error_type>::Status400_AnErrorOccurred(GenericError {
            message: Some(msg.to_string()),
            code: Some(1),
        }));
    }};
}

#[macro_export]
macro_rules! rover_is_operating {
    ($error_type:ty) => {{
        let msg = "unable to perform request, rover is running";
        warn!(msg);
        return Ok(<$error_type>::Status400_AnErrorOccurred(GenericError {
            message: Some(msg.to_string()),
            code: Some(1),
        }));
    }};
}

#[macro_export]
macro_rules! time_now {
    () => {
        SystemTime::now()
            .duration_since(UNIX_EPOCH)
            .unwrap()
            .as_millis()
    };
}

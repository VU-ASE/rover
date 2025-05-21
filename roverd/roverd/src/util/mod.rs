use anyhow::{anyhow, Context, Result};
use rand::distr::{Alphanumeric, SampleString};
use reqwest::Client;
use std::fs::{File, OpenOptions};
use std::os::unix::fs::{chown, PermissionsExt};
use std::time::Duration;
use std::{
    fs,
    io::{self, Read, Write},
    path::{Path, PathBuf},
};

use axum::http::StatusCode;

use rover_validate::config::{Configuration, Validate};

use rover_validate::service::Service;
use tracing::{info, warn};

use rover_types::error::Error;
use rover_types::service::FqBuf;

use rover_constants::*;

/// Makes all files within a directory and its subdirectories executable.
fn make_files_executable<P: AsRef<Path>>(dir_path: P) -> Result<()> {
    let dir_path = dir_path.as_ref();

    // Check if path exists and is a directory
    if !dir_path.exists() || !dir_path.is_dir() {
        return Err(anyhow::anyhow!(
            "Path is not a valid directory: {:?}",
            dir_path
        ));
    }

    // Process all entries in the directory
    for entry in fs::read_dir(dir_path)
        .with_context(|| format!("Failed to read directory: {:?}", dir_path))?
    {
        let entry = entry.with_context(|| format!("Failed to access entry in: {:?}", dir_path))?;
        let path = entry.path();

        if path.is_file() {
            // Add executable bit for files
            let metadata = fs::metadata(&path)?;
            let mut permissions = metadata.permissions();
            permissions.set_mode(permissions.mode() | 0o111); // Add executable bit
            fs::set_permissions(&path, permissions)?;
        } else if path.is_dir() {
            // Recursively process subdirectories
            make_files_executable(&path)?;
        }
    }

    Ok(())
}

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

    make_files_executable(&destination_dir)?;

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
/// that the new service can be safely moved in place. All created directories
/// will be owned by the debix:debix user.
fn prepare_dirs(fq: &FqBuf) -> Result<String, Error> {
    // Construct the full path
    let full_path_string = fq.dir().clone();
    let full_path = PathBuf::from(full_path_string.clone());

    // First check if the directory exists
    let path_exists = full_path.exists();

    // If it already existed and it contained old contents, remove them.
    if path_exists {
        std::fs::remove_dir_all(full_path.as_path())
            .with_context(|| format!("failed to remove path {:?}", full_path))?;
    }

    // Ensure the directory exists by creating it (and parent directories if needed)
    std::fs::create_dir_all(full_path.clone())
        .with_context(|| format!("failed to create dirs {:?}", full_path))?;

    // Set the ownership of the directory to debix:debix
    chown(&full_path, DEBIX_UID, DEBIX_GID)
        .with_context(|| format!("failed to set the ownership of {:?}", full_path))?;

    // Also handle parent directories to ensure consistent ownership
    let mut current = full_path.clone();
    while let Some(parent) = current.parent() {
        // Stop at the root or when we hit a directory that isn't part of our service structure
        if parent.as_os_str().is_empty() || !parent.to_string_lossy().contains(ROVER_DIR) {
            break;
        }
        // Set ownership for parent directories that are part of our service structure
        chown(parent, DEBIX_UID, DEBIX_GID)
            .with_context(|| format!("failed to set the ownership of parent dir {:?}", parent))?;
        current = parent.to_path_buf();
    }

    Ok(full_path_string)
}

/// Downloads the vu-ase service from the downloads page and creates a zip file on disk,
/// returns a String of the downloaded file, for example: "/tmp/some_string"
pub async fn download_service(url: &String) -> Result<String, Error> {
    info!("Downloading: {}", url);

    // generate a random string for the filename to avoid conflicts during concurrent downloads
    let file_name = format!(
        "/tmp/{}.zip",
        Alphanumeric.sample_string(&mut rand::rng(), 16)
    );

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
            std::fs::remove_file(&file_name).ok();

            let mut file = std::fs::File::create(&file_name)
                .with_context(|| format!("failed to create {}", file_name))?;

            let bytes = res.bytes().await?;

            file.write_all(&bytes)
                .with_context(|| format!("failed to failed to write to {}", file_name))?;
            Ok(file_name)
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
    let zip_file = download_service(url).await?;
    let (mut fq, directory) = extract_fq_from_zip(zip_file).await?;
    fq.is_daemon = is_daemon;
    install_service(directory, &fq).await?;
    Ok(fq)
}

/// Given a zip file, reads out the contents into a directory with the same name of the zip file
/// since the contents of the download are a service, it will parse the service.yaml and
/// return the resulting Fq as well as the directory where it was extracted to.
pub async fn extract_fq_from_zip(zip_file: String) -> Result<(FqBuf, String), Error> {
    let unzipped_dir = match zip_file.strip_suffix(".zip") {
        Some(stripped_file_name) => stripped_file_name,
        None => &zip_file,
    };

    // Clear the destination directory, no matter if it fails
    let _ = std::fs::remove_dir_all(unzipped_dir);

    // Create directory, this must not fail
    std::fs::create_dir_all(unzipped_dir)
        .with_context(|| format!("failed to create {}", unzipped_dir))?;

    // Unpack the downloaded service and validate it.
    extract_zip(&zip_file, unzipped_dir)?;

    // Read contents and
    let service_contents = std::fs::read_to_string(format!("{}/service.yaml", unzipped_dir))
        .map_err(|_| Error::ServiceYamlNotFoundInDownload)?;
    let service = serde_yaml::from_str::<Service>(&service_contents)?.validate()?;

    let fq = FqBuf::from(service);
    Ok((fq, unzipped_dir.to_string()))
}

/// Expects a zipfile to be ready at src_dir, extract it and install it. Parses the service.yaml
/// and install contents into the correct location on disk.
pub async fn install_service(src_dir: String, fq: &FqBuf) -> Result<(), Error> {
    // Deletes any existing files/dirs that are on the /author/name/version path
    // Makes sure the directories exist.
    let full_path = prepare_dirs(fq)?;

    // Copy contents into place
    copy_recursively(&src_dir, &full_path)
        .with_context(|| format!("failed to copy contents from {} to {}", &src_dir, full_path))?;

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

pub fn roverd_log(file_path: PathBuf, msg: String) -> Result<(), Error> {
    let mut log_file = create_log_file(&file_path)?;

    let cur_time = chrono::Local::now().format("%H:%M:%S");
    if writeln!(log_file, "[roverd {}] {}", cur_time, msg).is_err() {
        warn!("could not write log_line to file: {:?}", file_path)
    };

    Ok(())
}

#[macro_export]
macro_rules! warn_generic {
    ($expr:expr, $error_type:ty) => {{
        match $expr {
            Ok(data) => data,
            Err(e) => {
                warn!("{:#?}", e);
                let generic_error = GenericError::new(format!("{:#?}", e), 1);
                // todo remove the unwraps and change to actual error
                let json_string = serde_json::to_string(&generic_error).unwrap();
                let box_raw = serde_json::value::RawValue::from_string(json_string).unwrap();
                return Ok(<$error_type>::Status400_ErrorOccurred(
                    openapi::models::RoverdError::new(
                        "generic".to_string(),
                        openapi::models::RoverdErrorErrorValue(box_raw),
                    ),
                ));
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
                let generic_error = GenericError::new(format!("{:#?}", e), 1);
                // todo remove the unwraps and change to actual error
                let json_string = serde_json::to_string(&generic_error).unwrap();
                let box_raw = serde_json::value::RawValue::from_string(json_string).unwrap();
                return Ok(<$error_type>::Status400_ErrorOccurred(RoverdError::new(
                    "generic".to_string(),
                    RoverdErrorErrorValue(box_raw),
                )));
            }
        }
    }};
}

#[macro_export]
macro_rules! rover_is_dormant {
    ($error_type:ty) => {{
        let msg = "unable to perform request, rover is not running";
        warn!(msg);

        let generic_error = GenericError::new(msg.to_string(), 1);

        // todo remove the unwraps and change to actual error
        let json_string = serde_json::to_string(&generic_error).unwrap();
        let box_raw = serde_json::value::RawValue::from_string(json_string).unwrap();
        Ok(<$error_type>::Status400_ErrorOccurred(RoverdError::new(
            "generic".to_string(),
            RoverdErrorErrorValue(box_raw),
        )))
    }};
}

#[macro_export]
macro_rules! rover_is_operating {
    ($error_type:ty) => {{
        let msg = "unable to perform request, rover is running";
        warn!(msg);

        let generic_error = GenericError::new(msg.to_string(), 1);

        // todo remove the unwraps and change to actual error
        let json_string = serde_json::to_string(&generic_error).unwrap();
        let box_raw = serde_json::value::RawValue::from_string(json_string).unwrap();
        Ok(<$error_type>::Status400_ErrorOccurred(
            openapi::models::RoverdError::new(
                "generic".to_string(),
                openapi::models::RoverdErrorErrorValue(box_raw),
            ),
        ))
    }};
}

#[macro_export]
macro_rules! time_now {
    () => {{
        SystemTime::now()
            .duration_since(UNIX_EPOCH)
            .unwrap()
            .as_millis()
    }};
}

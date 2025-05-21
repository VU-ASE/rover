use std::{fmt::Display, fs, path::Path};

use crate::error;
use crate::error::*;

use rover_validate::config::Validate;
use rover_validate::service::ValidatedService;

use rover_constants::*;

use openapi::models::*;

pub fn get_service_as<T: AsRef<str>>(author: T, name: T, version: T) -> Option<String> {
    let contents_opt = fs::read_to_string(format!(
        "{}/{}/{}/{}/service.yaml",
        ROVER_DIR,
        author.as_ref(),
        name.as_ref(),
        version.as_ref()
    ))
    .ok();

    if let Some(contents) = contents_opt {
        let service_opt = serde_yaml::from_str::<rover_validate::service::Service>(&contents).ok();
        if let Some(service) = service_opt {
            if let Ok(valid) = service.validate() {
                return valid.0.service_as;
            }
        }
    }
    None
}

#[allow(dead_code)]
pub struct FqVec<'a>(pub Vec<Fq<'a>>);

#[allow(dead_code)]
pub struct FqBufVec(pub Vec<FqBuf>);

/// Internal representation of a service, whether as a source or user service.
#[derive(Debug)]
pub struct Fq<'a> {
    pub author: &'a str,
    pub name: &'a str,
    pub version: &'a str,
}

/// Same as FqService but with Strings instead of &str.
#[derive(Debug, Eq, Hash, PartialEq)]
pub struct FqBuf {
    pub author: String,
    pub name: String,
    pub version: String,
    pub is_daemon: bool,
    pub service_as: Option<String>,
}

impl FqBuf {
    pub fn new(author: &str, name: &str, version: &str, service_as: &Option<String>) -> FqBuf {
        FqBuf {
            author: author.to_string(),
            name: name.to_string(),
            version: version.to_string(),
            service_as: service_as.clone(),
            is_daemon: false,
        }
    }

    pub fn new_daemon(author: &str, name: &str, version: &str) -> FqBuf {
        FqBuf {
            author: author.to_string(),
            name: name.to_string(),
            version: version.to_string(),
            service_as: None,
            is_daemon: true,
        }
    }
}

impl From<FqBuf> for FullyQualifiedService {
    fn from(value: FqBuf) -> Self {
        FullyQualifiedService {
            author: value.author,
            name: value.name,
            version: value.version,
            r#as: value.service_as,
        }
    }
}

impl From<FullyQualifiedService> for FqBuf {
    fn from(value: FullyQualifiedService) -> Self {
        FqBuf::new(&value.author, &value.name, &value.version, &value.r#as)
    }
}

impl From<&FullyQualifiedService> for FqBuf {
    fn from(value: &FullyQualifiedService) -> Self {
        FqBuf::new(&value.author, &value.name, &value.version, &value.r#as)
    }
}

impl Clone for FqBuf {
    fn clone(&self) -> Self {
        Self {
            author: self.author.clone(),
            name: self.name.clone(),
            version: self.version.clone(),
            service_as: self.service_as.clone(),
            is_daemon: self.is_daemon,
        }
    }
}

impl From<ValidatedService> for FqBuf {
    fn from(service: ValidatedService) -> Self {
        FqBuf {
            name: service.0.name,
            author: service.0.author,
            version: service.0.version,
            service_as: service.0.service_as,
            is_daemon: false,
        }
    }
}

impl From<&ValidatedService> for FqBuf {
    fn from(service: &ValidatedService) -> Self {
        FqBuf {
            name: service.0.name.clone(),
            author: service.0.author.clone(),
            version: service.0.version.clone(),
            service_as: service.0.service_as.clone(),
            is_daemon: false,
        }
    }
}

impl From<&ServicesAuthorServiceVersionPostPathParams> for FqBuf {
    fn from(value: &ServicesAuthorServiceVersionPostPathParams) -> Self {
        FqBuf {
            name: value.service.clone(),
            author: value.author.clone(),
            version: value.version.clone(),
            service_as: get_service_as(&value.author, &value.service, &value.version),
            is_daemon: false,
        }
    }
}

impl From<&ServicesAuthorServiceVersionDeletePathParams> for FqBuf {
    fn from(value: &ServicesAuthorServiceVersionDeletePathParams) -> Self {
        FqBuf {
            name: value.service.clone(),
            author: value.author.clone(),
            version: value.version.clone(),
            service_as: get_service_as(&value.author, &value.service, &value.version),
            is_daemon: false,
        }
    }
}

impl From<&ServicesAuthorServiceVersionGetPathParams> for FqBuf {
    fn from(value: &ServicesAuthorServiceVersionGetPathParams) -> Self {
        FqBuf {
            name: value.service.clone(),
            author: value.author.clone(),
            version: value.version.clone(),
            service_as: get_service_as(&value.author, &value.service, &value.version),
            is_daemon: false,
        }
    }
}

impl From<&PipelinePostRequestInner> for FqBuf {
    fn from(service: &PipelinePostRequestInner) -> Self {
        FqBuf {
            name: service.fq.name.clone(),
            author: service.fq.author.clone(),
            version: service.fq.version.clone(),
            service_as: get_service_as(&service.fq.author, &service.fq.name, &service.fq.version),
            is_daemon: false,
        }
    }
}

impl From<&LogsAuthorNameVersionGetPathParams> for FqBuf {
    fn from(value: &LogsAuthorNameVersionGetPathParams) -> Self {
        FqBuf {
            name: value.name.clone(),
            author: value.author.clone(),
            version: value.version.clone(),
            service_as: get_service_as(&value.author, &value.name, &value.version),
            is_daemon: false,
        }
    }
}

impl From<&ServicesAuthorServiceVersionConfigurationPostPathParams> for FqBuf {
    fn from(value: &ServicesAuthorServiceVersionConfigurationPostPathParams) -> Self {
        FqBuf {
            author: value.author.clone(),
            name: value.service.clone(),
            version: value.version.clone(),
            is_daemon: false,
            service_as: get_service_as(&value.author, &value.service, &value.version),
        }
    }
}

impl FqBuf {
    /// Returns the file path of the associated service.yaml.
    pub fn path(&self) -> String {
        if self.is_daemon {
            format!(
                "{}/{}/{}/{}/service.yaml",
                DAEMON_DIR, self.author, self.name, self.version
            )
        } else {
            format!(
                "{}/{}/{}/{}/service.yaml",
                ROVER_DIR, self.author, self.name, self.version
            )
        }
    }

    /// Returns the file path of the associated log file for the service.
    pub fn log_file(&self) -> String {
        format!(
            "{}/{}-{}-{}.log",
            LOG_DIR, self.author, self.name, self.version
        )
    }

    /// Returns the file path of the associated build-log file for the service.
    pub fn build_log_file(&self) -> String {
        format!(
            "{}/build-{}-{}-{}.log",
            BUILD_LOG_DIR, self.author, self.name, self.version
        )
    }

    /// Returns the directory of the service. This directory contains the code
    /// and service.yaml definitions.
    pub fn dir(&self) -> String {
        if self.is_daemon {
            format!(
                "{}/{}/{}/{}",
                DAEMON_DIR, self.author, self.name, self.version
            )
        } else {
            format!(
                "{}/{}/{}/{}",
                ROVER_DIR, self.author, self.name, self.version
            )
        }
    }
}

impl Display for FqBuf {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}/{}/{}", self.author, self.name, self.version)?;
        Ok(())
    }
}

impl<'a> TryFrom<&'a String> for Fq<'a> {
    type Error = error::Error;
    fn try_from(path_string: &'a String) -> Result<Self, Self::Error> {
        let path = Path::new(path_string.as_str());
        let path_vec: Vec<_> = path.components().collect();

        let num_directory_levels = path_vec.len();
        if num_directory_levels < 3 {
            return Err(Error::EnabledPathInvalid);
        }

        let values = path_vec[(path_vec.len() - 4)..(path_vec.len() - 1)]
            .iter()
            .map(|component| {
                component
                    .as_os_str()
                    .to_str()
                    .ok_or(Error::StringToFqConversion)
            })
            .collect::<Result<Vec<&str>, Error>>()?;

        Ok(Fq {
            author: values.first().ok_or(Error::StringToFqConversion)?,
            name: values.get(1).ok_or(Error::StringToFqConversion)?,
            version: values.get(2).ok_or(Error::StringToFqConversion)?,
        })
    }
}

impl TryFrom<String> for FqBuf {
    type Error = error::Error;
    fn try_from(path_string: String) -> Result<Self, Self::Error> {
        let path = Path::new(path_string.as_str());
        let path_vec: Vec<_> = path.components().collect();

        let num_directory_levels = path_vec.len();
        if num_directory_levels < 3 {
            return Err(Error::EnabledPathInvalid);
        }

        let values = path_vec[(path_vec.len() - 4)..(path_vec.len() - 1)]
            .iter()
            .map(|component| Ok(component.as_os_str().to_os_string().into_string()?))
            .collect::<Result<Vec<String>, Error>>()?;

        let author = values.first().ok_or(Error::StringToFqConversion)?.clone();
        let name = values.get(1).ok_or(Error::StringToFqConversion)?.clone();
        let version = values.get(2).ok_or(Error::StringToFqConversion)?.clone();

        Ok(FqBuf {
            author: author.clone(),
            name: name.clone(),
            version: version.clone(),
            service_as: get_service_as(author, name, version),
            is_daemon: false,
        })
    }
}

impl TryFrom<&String> for FqBuf {
    type Error = error::Error;
    fn try_from(path_string: &String) -> Result<Self, Self::Error> {
        let path = Path::new(path_string.as_str());
        let path_vec: Vec<_> = path.components().collect();

        let num_directory_levels = path_vec.len();
        if num_directory_levels < 3 {
            return Err(Error::EnabledPathInvalid);
        }

        let values = path_vec[(path_vec.len() - 4)..(path_vec.len() - 1)]
            .iter()
            .map(|component| Ok(component.as_os_str().to_os_string().into_string()?))
            .collect::<Result<Vec<String>, Error>>()?;

        let author = values.first().ok_or(Error::StringToFqConversion)?.clone();
        let name = values.get(1).ok_or(Error::StringToFqConversion)?.clone();
        let version = values.get(2).ok_or(Error::StringToFqConversion)?.clone();

        Ok(FqBuf {
            author: author.clone(),
            name: name.clone(),
            version: version.clone(),
            service_as: get_service_as(author, name, version),
            is_daemon: false,
        })
    }
}

impl<'a> TryFrom<&'a Vec<String>> for FqVec<'a> {
    type Error = error::Error;
    fn try_from(string_vec: &'a Vec<String>) -> Result<Self, Self::Error> {
        let fq_services: Vec<Fq<'a>> = string_vec
            .iter()
            .map(Fq::try_from)
            .collect::<Result<Vec<_>, _>>()?;
        Ok(FqVec(fq_services))
    }
}

impl TryFrom<Vec<String>> for FqBufVec {
    type Error = error::Error;
    fn try_from(string_vec: Vec<String>) -> Result<Self, Self::Error> {
        let fq_services: Vec<FqBuf> = string_vec
            .iter()
            .map(FqBuf::try_from)
            .collect::<Result<Vec<_>, _>>()?;
        Ok(FqBufVec(fq_services))
    }
}

impl TryFrom<&Vec<String>> for FqBufVec {
    type Error = error::Error;
    fn try_from(string_vec: &Vec<String>) -> Result<Self, Self::Error> {
        let fq_services: Vec<FqBuf> = string_vec
            .iter()
            .map(FqBuf::try_from)
            .collect::<Result<Vec<_>, _>>()?;
        Ok(FqBufVec(fq_services))
    }
}

impl<'a> From<&'a Vec<PipelinePostRequestInner>> for FqVec<'a> {
    fn from(vec: &'a Vec<PipelinePostRequestInner>) -> Self {
        let fq_services = vec.iter().map(Fq::from).collect::<Vec<_>>();
        FqVec(fq_services)
    }
}

impl From<Vec<PipelinePostRequestInner>> for FqBufVec {
    fn from(vec: Vec<PipelinePostRequestInner>) -> Self {
        let fq_services = vec.iter().map(FqBuf::from).collect::<Vec<_>>();
        FqBufVec(fq_services)
    }
}

impl Display for Fq<'_> {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}/{}/{}", self.author, self.name, self.version)?;
        Ok(())
    }
}

impl<'a> From<&'a FullyQualifiedService> for Fq<'a> {
    fn from(value: &'a FullyQualifiedService) -> Self {
        Fq {
            author: &value.author,
            name: &value.name,
            version: &value.version,
        }
    }
}

impl<'a> From<&'a PipelinePostRequestInner> for Fq<'a> {
    fn from(p: &'a PipelinePostRequestInner) -> Self {
        Fq {
            name: &p.fq.name,
            author: &p.fq.author,
            version: &p.fq.version,
        }
    }
}

impl<'a> From<&'a ServicesAuthorServiceVersionDeletePathParams> for Fq<'a> {
    fn from(param: &'a ServicesAuthorServiceVersionDeletePathParams) -> Self {
        Fq {
            name: &param.service,
            author: &param.author,
            version: &param.version,
        }
    }
}

impl<'a> From<&'a FqBuf> for Fq<'a> {
    fn from(param: &'a FqBuf) -> Self {
        Fq {
            name: &param.name,
            author: &param.author,
            version: &param.version,
        }
    }
}

impl PartialEq for Fq<'_> {
    fn eq(&self, other: &Self) -> bool {
        self.name.to_lowercase() == other.name.to_lowercase()
            && self.author.to_lowercase() == other.author.to_lowercase()
            && self.version.to_lowercase() == other.version.to_lowercase()
    }
}

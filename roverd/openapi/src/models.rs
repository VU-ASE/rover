#![allow(unused_qualifications)]

use http::HeaderValue;
use validator::Validate;

#[cfg(feature = "server")]
use crate::header;
use crate::{models, types::*};

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct LogsAuthorNameVersionGetPathParams {
    /// The author of the service.
    pub author: String,
    /// The name of the service.
    pub name: String,
    /// The version of the service.
    pub version: String,
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct LogsAuthorNameVersionGetQueryParams {
    /// The number of log lines to retrieve
    #[serde(rename = "lines")]
    #[validate(range(min = 1, max = 1000))]
    #[serde(skip_serializing_if = "Option::is_none")]
    pub lines: Option<i32>,
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct ServicesAuthorGetPathParams {
    /// The author name
    pub author: String,
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct ServicesAuthorServiceGetPathParams {
    /// The author name
    pub author: String,
    /// The service name
    pub service: String,
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct ServicesAuthorServiceVersionConfigurationPostPathParams {
    /// The author name
    pub author: String,
    /// The service name
    pub service: String,
    /// The version of the service
    pub version: String,
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct ServicesAuthorServiceVersionDeletePathParams {
    /// The author name
    pub author: String,
    /// The service name
    pub service: String,
    /// The version of the service
    pub version: String,
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct ServicesAuthorServiceVersionGetPathParams {
    /// The author name
    pub author: String,
    /// The service name
    pub service: String,
    /// The version of the service
    pub version: String,
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct ServicesAuthorServiceVersionPostPathParams {
    /// The author name
    pub author: String,
    /// The service name
    pub service: String,
    /// The version of the service
    pub version: String,
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct BuildError {
    /// The build log (one log line per item)
    #[serde(rename = "build_log")]
    pub build_log: Vec<String>,
}

impl BuildError {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(build_log: Vec<String>) -> BuildError {
        BuildError { build_log }
    }
}

/// Converts the BuildError value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for BuildError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            Some("build_log".to_string()),
            Some(
                self.build_log
                    .iter()
                    .map(|x| x.to_string())
                    .collect::<Vec<_>>()
                    .join(","),
            ),
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a BuildError value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for BuildError {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub build_log: Vec<Vec<String>>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing BuildError".to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    "build_log" => {
                        return std::result::Result::Err(
                            "Parsing a container in this style is not supported in BuildError"
                                .to_string(),
                        )
                    }
                    _ => {
                        return std::result::Result::Err(
                            "Unexpected key while parsing BuildError".to_string(),
                        )
                    }
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(BuildError {
            build_log: intermediate_rep
                .build_log
                .into_iter()
                .next()
                .ok_or_else(|| "build_log missing in BuildError".to_string())?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<BuildError> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<BuildError>> for HeaderValue {
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<BuildError>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
            std::result::Result::Ok(value) => std::result::Result::Ok(value),
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Invalid header value for BuildError - value: {} is invalid {}",
                hdr_value, e
            )),
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue> for header::IntoHeaderValue<BuildError> {
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
            std::result::Result::Ok(value) => {
                match <BuildError as std::str::FromStr>::from_str(value) {
                    std::result::Result::Ok(value) => {
                        std::result::Result::Ok(header::IntoHeaderValue(value))
                    }
                    std::result::Result::Err(err) => std::result::Result::Err(format!(
                        "Unable to convert header value '{}' into BuildError - {}",
                        value, err
                    )),
                }
            }
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Unable to convert header: {:?} to string: {}",
                hdr_value, e
            )),
        }
    }
}

/// The status of the roverd process
/// Enumeration of values.
/// Since this enum's variants do not hold data, we can easily define them as `#[repr(C)]`
/// which helps with FFI.
#[allow(non_camel_case_types)]
#[repr(C)]
#[derive(
    Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, serde::Serialize, serde::Deserialize,
)]
#[cfg_attr(feature = "conversion", derive(frunk_enum_derive::LabelledGenericEnum))]
pub enum DaemonStatus {
    #[serde(rename = "operational")]
    Operational,
    #[serde(rename = "recoverable")]
    Recoverable,
    #[serde(rename = "unrecoverable")]
    Unrecoverable,
}

impl std::fmt::Display for DaemonStatus {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match *self {
            DaemonStatus::Operational => write!(f, "operational"),
            DaemonStatus::Recoverable => write!(f, "recoverable"),
            DaemonStatus::Unrecoverable => write!(f, "unrecoverable"),
        }
    }
}

impl std::str::FromStr for DaemonStatus {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        match s {
            "operational" => std::result::Result::Ok(DaemonStatus::Operational),
            "recoverable" => std::result::Result::Ok(DaemonStatus::Recoverable),
            "unrecoverable" => std::result::Result::Ok(DaemonStatus::Unrecoverable),
            _ => std::result::Result::Err(format!("Value not valid: {}", s)),
        }
    }
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct FetchPost200Response {
    #[serde(rename = "fq")]
    pub fq: models::FullyQualifiedService,

    /// Whether the pipeline was invalidated by this service upload
    #[serde(rename = "invalidated_pipeline")]
    pub invalidated_pipeline: bool,
}

impl FetchPost200Response {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(
        fq: models::FullyQualifiedService,
        invalidated_pipeline: bool,
    ) -> FetchPost200Response {
        FetchPost200Response {
            fq,
            invalidated_pipeline,
        }
    }
}

/// Converts the FetchPost200Response value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for FetchPost200Response {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            // Skipping fq in query parameter serialization
            Some("invalidated_pipeline".to_string()),
            Some(self.invalidated_pipeline.to_string()),
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a FetchPost200Response value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for FetchPost200Response {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub fq: Vec<models::FullyQualifiedService>,
            pub invalidated_pipeline: Vec<bool>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing FetchPost200Response".to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "fq" => intermediate_rep.fq.push(
                        <models::FullyQualifiedService as std::str::FromStr>::from_str(val)
                            .map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "invalidated_pipeline" => intermediate_rep.invalidated_pipeline.push(
                        <bool as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    _ => {
                        return std::result::Result::Err(
                            "Unexpected key while parsing FetchPost200Response".to_string(),
                        )
                    }
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(FetchPost200Response {
            fq: intermediate_rep
                .fq
                .into_iter()
                .next()
                .ok_or_else(|| "fq missing in FetchPost200Response".to_string())?,
            invalidated_pipeline: intermediate_rep
                .invalidated_pipeline
                .into_iter()
                .next()
                .ok_or_else(|| {
                    "invalidated_pipeline missing in FetchPost200Response".to_string()
                })?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<FetchPost200Response> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<FetchPost200Response>> for HeaderValue {
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<FetchPost200Response>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
            std::result::Result::Ok(value) => std::result::Result::Ok(value),
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Invalid header value for FetchPost200Response - value: {} is invalid {}",
                hdr_value, e
            )),
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue> for header::IntoHeaderValue<FetchPost200Response> {
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
            std::result::Result::Ok(value) => {
                match <FetchPost200Response as std::str::FromStr>::from_str(value) {
                    std::result::Result::Ok(value) => {
                        std::result::Result::Ok(header::IntoHeaderValue(value))
                    }
                    std::result::Result::Err(err) => std::result::Result::Err(format!(
                        "Unable to convert header value '{}' into FetchPost200Response - {}",
                        value, err
                    )),
                }
            }
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Unable to convert header: {:?} to string: {}",
                hdr_value, e
            )),
        }
    }
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct FetchPostRequest {
    /// Download URL of the service to be downloaded, must include scheme
    #[serde(rename = "url")]
    pub url: String,
}

impl FetchPostRequest {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(url: String) -> FetchPostRequest {
        FetchPostRequest { url }
    }
}

/// Converts the FetchPostRequest value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for FetchPostRequest {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![Some("url".to_string()), Some(self.url.to_string())];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a FetchPostRequest value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for FetchPostRequest {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub url: Vec<String>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing FetchPostRequest".to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "url" => intermediate_rep.url.push(
                        <String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    _ => {
                        return std::result::Result::Err(
                            "Unexpected key while parsing FetchPostRequest".to_string(),
                        )
                    }
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(FetchPostRequest {
            url: intermediate_rep
                .url
                .into_iter()
                .next()
                .ok_or_else(|| "url missing in FetchPostRequest".to_string())?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<FetchPostRequest> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<FetchPostRequest>> for HeaderValue {
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<FetchPostRequest>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
            std::result::Result::Ok(value) => std::result::Result::Ok(value),
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Invalid header value for FetchPostRequest - value: {} is invalid {}",
                hdr_value, e
            )),
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue> for header::IntoHeaderValue<FetchPostRequest> {
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
            std::result::Result::Ok(value) => {
                match <FetchPostRequest as std::str::FromStr>::from_str(value) {
                    std::result::Result::Ok(value) => {
                        std::result::Result::Ok(header::IntoHeaderValue(value))
                    }
                    std::result::Result::Err(err) => std::result::Result::Err(format!(
                        "Unable to convert header value '{}' into FetchPostRequest - {}",
                        value, err
                    )),
                }
            }
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Unable to convert header: {:?} to string: {}",
                hdr_value, e
            )),
        }
    }
}

/// FullyQualifiedService
#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct FullyQualifiedService {
    /// The author of the service
    #[serde(rename = "author")]
    pub author: String,

    /// The name of the service
    #[serde(rename = "name")]
    pub name: String,

    /// The version of the service
    #[serde(rename = "version")]
    pub version: String,

    /// The (optional) alias of the name to be used in the pipeline
    #[serde(rename = "as")]
    #[serde(skip_serializing_if = "Option::is_none")]
    pub r#as: Option<String>,
}

impl FullyQualifiedService {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(author: String, name: String, version: String) -> FullyQualifiedService {
        FullyQualifiedService {
            author,
            name,
            version,
            r#as: None,
        }
    }
}

/// Converts the FullyQualifiedService value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for FullyQualifiedService {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            Some("author".to_string()),
            Some(self.author.to_string()),
            Some("name".to_string()),
            Some(self.name.to_string()),
            Some("version".to_string()),
            Some(self.version.to_string()),
            self.r#as
                .as_ref()
                .map(|r#as| ["as".to_string(), r#as.to_string()].join(",")),
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a FullyQualifiedService value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for FullyQualifiedService {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub author: Vec<String>,
            pub name: Vec<String>,
            pub version: Vec<String>,
            pub r#as: Vec<String>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing FullyQualifiedService".to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "author" => intermediate_rep.author.push(
                        <String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "name" => intermediate_rep.name.push(
                        <String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "version" => intermediate_rep.version.push(
                        <String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "as" => intermediate_rep.r#as.push(
                        <String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    _ => {
                        return std::result::Result::Err(
                            "Unexpected key while parsing FullyQualifiedService".to_string(),
                        )
                    }
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(FullyQualifiedService {
            author: intermediate_rep
                .author
                .into_iter()
                .next()
                .ok_or_else(|| "author missing in FullyQualifiedService".to_string())?,
            name: intermediate_rep
                .name
                .into_iter()
                .next()
                .ok_or_else(|| "name missing in FullyQualifiedService".to_string())?,
            version: intermediate_rep
                .version
                .into_iter()
                .next()
                .ok_or_else(|| "version missing in FullyQualifiedService".to_string())?,
            r#as: intermediate_rep.r#as.into_iter().next(),
        })
    }
}

// Methods for converting between header::IntoHeaderValue<FullyQualifiedService> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<FullyQualifiedService>> for HeaderValue {
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<FullyQualifiedService>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
            std::result::Result::Ok(value) => std::result::Result::Ok(value),
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Invalid header value for FullyQualifiedService - value: {} is invalid {}",
                hdr_value, e
            )),
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue> for header::IntoHeaderValue<FullyQualifiedService> {
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
            std::result::Result::Ok(value) => {
                match <FullyQualifiedService as std::str::FromStr>::from_str(value) {
                    std::result::Result::Ok(value) => {
                        std::result::Result::Ok(header::IntoHeaderValue(value))
                    }
                    std::result::Result::Err(err) => std::result::Result::Err(format!(
                        "Unable to convert header value '{}' into FullyQualifiedService - {}",
                        value, err
                    )),
                }
            }
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Unable to convert header: {:?} to string: {}",
                hdr_value, e
            )),
        }
    }
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct GenericError {
    /// A message describing the error
    #[serde(rename = "message")]
    pub message: String,

    /// A code describing the error (this is not an HTTP status code)
    #[serde(rename = "code")]
    pub code: i32,
}

impl GenericError {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(message: String, code: i32) -> GenericError {
        GenericError { message, code }
    }
}

/// Converts the GenericError value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for GenericError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            Some("message".to_string()),
            Some(self.message.to_string()),
            Some("code".to_string()),
            Some(self.code.to_string()),
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a GenericError value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for GenericError {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub message: Vec<String>,
            pub code: Vec<i32>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing GenericError".to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "message" => intermediate_rep.message.push(
                        <String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "code" => intermediate_rep.code.push(
                        <i32 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    _ => {
                        return std::result::Result::Err(
                            "Unexpected key while parsing GenericError".to_string(),
                        )
                    }
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(GenericError {
            message: intermediate_rep
                .message
                .into_iter()
                .next()
                .ok_or_else(|| "message missing in GenericError".to_string())?,
            code: intermediate_rep
                .code
                .into_iter()
                .next()
                .ok_or_else(|| "code missing in GenericError".to_string())?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<GenericError> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<GenericError>> for HeaderValue {
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<GenericError>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
            std::result::Result::Ok(value) => std::result::Result::Ok(value),
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Invalid header value for GenericError - value: {} is invalid {}",
                hdr_value, e
            )),
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue> for header::IntoHeaderValue<GenericError> {
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
            std::result::Result::Ok(value) => {
                match <GenericError as std::str::FromStr>::from_str(value) {
                    std::result::Result::Ok(value) => {
                        std::result::Result::Ok(header::IntoHeaderValue(value))
                    }
                    std::result::Result::Err(err) => std::result::Result::Err(format!(
                        "Unable to convert header value '{}' into GenericError - {}",
                        value, err
                    )),
                }
            }
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Unable to convert header: {:?} to string: {}",
                hdr_value, e
            )),
        }
    }
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct Get200Response {
    #[serde(rename = "status")]
    pub status: models::DaemonStatus,

    /// Error message of the daemon status
    #[serde(rename = "error_message")]
    #[serde(skip_serializing_if = "Option::is_none")]
    pub error_message: Option<String>,

    /// The version of the roverd daemon
    #[serde(rename = "version")]
    pub version: String,

    /// The number of milliseconds the roverd daemon process has been running
    #[serde(rename = "uptime")]
    pub uptime: i64,

    /// The operating system of the rover
    #[serde(rename = "os")]
    pub os: String,

    /// The system time of the rover as milliseconds since epoch
    #[serde(rename = "systime")]
    pub systime: i64,

    /// The unique identifier of the rover
    #[serde(rename = "rover_id")]
    #[serde(skip_serializing_if = "Option::is_none")]
    pub rover_id: Option<i32>,

    /// The unique name of the rover
    #[serde(rename = "rover_name")]
    #[serde(skip_serializing_if = "Option::is_none")]
    pub rover_name: Option<String>,

    #[serde(rename = "memory")]
    pub memory: models::Get200ResponseMemory,

    /// The CPU usage of the roverd process
    #[serde(rename = "cpu")]
    pub cpu: Vec<models::Get200ResponseCpuInner>,
}

impl Get200Response {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(
        status: models::DaemonStatus,
        version: String,
        uptime: i64,
        os: String,
        systime: i64,
        memory: models::Get200ResponseMemory,
        cpu: Vec<models::Get200ResponseCpuInner>,
    ) -> Get200Response {
        Get200Response {
            status,
            error_message: None,
            version,
            uptime,
            os,
            systime,
            rover_id: None,
            rover_name: None,
            memory,
            cpu,
        }
    }
}

/// Converts the Get200Response value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for Get200Response {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            // Skipping status in query parameter serialization
            self.error_message.as_ref().map(|error_message| {
                ["error_message".to_string(), error_message.to_string()].join(",")
            }),
            Some("version".to_string()),
            Some(self.version.to_string()),
            Some("uptime".to_string()),
            Some(self.uptime.to_string()),
            Some("os".to_string()),
            Some(self.os.to_string()),
            Some("systime".to_string()),
            Some(self.systime.to_string()),
            self.rover_id
                .as_ref()
                .map(|rover_id| ["rover_id".to_string(), rover_id.to_string()].join(",")),
            self.rover_name
                .as_ref()
                .map(|rover_name| ["rover_name".to_string(), rover_name.to_string()].join(",")),
            // Skipping memory in query parameter serialization

            // Skipping cpu in query parameter serialization
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a Get200Response value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for Get200Response {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub status: Vec<models::DaemonStatus>,
            pub error_message: Vec<String>,
            pub version: Vec<String>,
            pub uptime: Vec<i64>,
            pub os: Vec<String>,
            pub systime: Vec<i64>,
            pub rover_id: Vec<i32>,
            pub rover_name: Vec<String>,
            pub memory: Vec<models::Get200ResponseMemory>,
            pub cpu: Vec<Vec<models::Get200ResponseCpuInner>>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing Get200Response".to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "status" => intermediate_rep.status.push(
                        <models::DaemonStatus as std::str::FromStr>::from_str(val)
                            .map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "error_message" => intermediate_rep.error_message.push(
                        <String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "version" => intermediate_rep.version.push(
                        <String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "uptime" => intermediate_rep.uptime.push(
                        <i64 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "os" => intermediate_rep.os.push(
                        <String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "systime" => intermediate_rep.systime.push(
                        <i64 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "rover_id" => intermediate_rep.rover_id.push(
                        <i32 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "rover_name" => intermediate_rep.rover_name.push(
                        <String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "memory" => intermediate_rep.memory.push(
                        <models::Get200ResponseMemory as std::str::FromStr>::from_str(val)
                            .map_err(|x| x.to_string())?,
                    ),
                    "cpu" => {
                        return std::result::Result::Err(
                            "Parsing a container in this style is not supported in Get200Response"
                                .to_string(),
                        )
                    }
                    _ => {
                        return std::result::Result::Err(
                            "Unexpected key while parsing Get200Response".to_string(),
                        )
                    }
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(Get200Response {
            status: intermediate_rep
                .status
                .into_iter()
                .next()
                .ok_or_else(|| "status missing in Get200Response".to_string())?,
            error_message: intermediate_rep.error_message.into_iter().next(),
            version: intermediate_rep
                .version
                .into_iter()
                .next()
                .ok_or_else(|| "version missing in Get200Response".to_string())?,
            uptime: intermediate_rep
                .uptime
                .into_iter()
                .next()
                .ok_or_else(|| "uptime missing in Get200Response".to_string())?,
            os: intermediate_rep
                .os
                .into_iter()
                .next()
                .ok_or_else(|| "os missing in Get200Response".to_string())?,
            systime: intermediate_rep
                .systime
                .into_iter()
                .next()
                .ok_or_else(|| "systime missing in Get200Response".to_string())?,
            rover_id: intermediate_rep.rover_id.into_iter().next(),
            rover_name: intermediate_rep.rover_name.into_iter().next(),
            memory: intermediate_rep
                .memory
                .into_iter()
                .next()
                .ok_or_else(|| "memory missing in Get200Response".to_string())?,
            cpu: intermediate_rep
                .cpu
                .into_iter()
                .next()
                .ok_or_else(|| "cpu missing in Get200Response".to_string())?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<Get200Response> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<Get200Response>> for HeaderValue {
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<Get200Response>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
            std::result::Result::Ok(value) => std::result::Result::Ok(value),
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Invalid header value for Get200Response - value: {} is invalid {}",
                hdr_value, e
            )),
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue> for header::IntoHeaderValue<Get200Response> {
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
            std::result::Result::Ok(value) => {
                match <Get200Response as std::str::FromStr>::from_str(value) {
                    std::result::Result::Ok(value) => {
                        std::result::Result::Ok(header::IntoHeaderValue(value))
                    }
                    std::result::Result::Err(err) => std::result::Result::Err(format!(
                        "Unable to convert header value '{}' into Get200Response - {}",
                        value, err
                    )),
                }
            }
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Unable to convert header: {:?} to string: {}",
                hdr_value, e
            )),
        }
    }
}

/// CPU usage information about a specific core
#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct Get200ResponseCpuInner {
    /// The core number
    #[serde(rename = "core")]
    pub core: i32,

    /// The total amount of CPU available on the core
    #[serde(rename = "total")]
    pub total: i32,

    /// The amount of CPU used on the core
    #[serde(rename = "used")]
    pub used: i32,
}

impl Get200ResponseCpuInner {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(core: i32, total: i32, used: i32) -> Get200ResponseCpuInner {
        Get200ResponseCpuInner { core, total, used }
    }
}

/// Converts the Get200ResponseCpuInner value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for Get200ResponseCpuInner {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            Some("core".to_string()),
            Some(self.core.to_string()),
            Some("total".to_string()),
            Some(self.total.to_string()),
            Some("used".to_string()),
            Some(self.used.to_string()),
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a Get200ResponseCpuInner value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for Get200ResponseCpuInner {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub core: Vec<i32>,
            pub total: Vec<i32>,
            pub used: Vec<i32>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing Get200ResponseCpuInner".to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "core" => intermediate_rep.core.push(
                        <i32 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "total" => intermediate_rep.total.push(
                        <i32 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "used" => intermediate_rep.used.push(
                        <i32 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    _ => {
                        return std::result::Result::Err(
                            "Unexpected key while parsing Get200ResponseCpuInner".to_string(),
                        )
                    }
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(Get200ResponseCpuInner {
            core: intermediate_rep
                .core
                .into_iter()
                .next()
                .ok_or_else(|| "core missing in Get200ResponseCpuInner".to_string())?,
            total: intermediate_rep
                .total
                .into_iter()
                .next()
                .ok_or_else(|| "total missing in Get200ResponseCpuInner".to_string())?,
            used: intermediate_rep
                .used
                .into_iter()
                .next()
                .ok_or_else(|| "used missing in Get200ResponseCpuInner".to_string())?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<Get200ResponseCpuInner> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<Get200ResponseCpuInner>> for HeaderValue {
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<Get200ResponseCpuInner>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
            std::result::Result::Ok(value) => std::result::Result::Ok(value),
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Invalid header value for Get200ResponseCpuInner - value: {} is invalid {}",
                hdr_value, e
            )),
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue> for header::IntoHeaderValue<Get200ResponseCpuInner> {
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
            std::result::Result::Ok(value) => {
                match <Get200ResponseCpuInner as std::str::FromStr>::from_str(value) {
                    std::result::Result::Ok(value) => {
                        std::result::Result::Ok(header::IntoHeaderValue(value))
                    }
                    std::result::Result::Err(err) => std::result::Result::Err(format!(
                        "Unable to convert header value '{}' into Get200ResponseCpuInner - {}",
                        value, err
                    )),
                }
            }
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Unable to convert header: {:?} to string: {}",
                hdr_value, e
            )),
        }
    }
}

/// Memory usage information
#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct Get200ResponseMemory {
    /// The total amount of memory available on the rover in megabytes
    #[serde(rename = "total")]
    pub total: i32,

    /// The amount of memory used on the rover in megabytes
    #[serde(rename = "used")]
    pub used: i32,
}

impl Get200ResponseMemory {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(total: i32, used: i32) -> Get200ResponseMemory {
        Get200ResponseMemory { total, used }
    }
}

/// Converts the Get200ResponseMemory value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for Get200ResponseMemory {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            Some("total".to_string()),
            Some(self.total.to_string()),
            Some("used".to_string()),
            Some(self.used.to_string()),
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a Get200ResponseMemory value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for Get200ResponseMemory {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub total: Vec<i32>,
            pub used: Vec<i32>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing Get200ResponseMemory".to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "total" => intermediate_rep.total.push(
                        <i32 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "used" => intermediate_rep.used.push(
                        <i32 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    _ => {
                        return std::result::Result::Err(
                            "Unexpected key while parsing Get200ResponseMemory".to_string(),
                        )
                    }
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(Get200ResponseMemory {
            total: intermediate_rep
                .total
                .into_iter()
                .next()
                .ok_or_else(|| "total missing in Get200ResponseMemory".to_string())?,
            used: intermediate_rep
                .used
                .into_iter()
                .next()
                .ok_or_else(|| "used missing in Get200ResponseMemory".to_string())?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<Get200ResponseMemory> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<Get200ResponseMemory>> for HeaderValue {
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<Get200ResponseMemory>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
            std::result::Result::Ok(value) => std::result::Result::Ok(value),
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Invalid header value for Get200ResponseMemory - value: {} is invalid {}",
                hdr_value, e
            )),
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue> for header::IntoHeaderValue<Get200ResponseMemory> {
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
            std::result::Result::Ok(value) => {
                match <Get200ResponseMemory as std::str::FromStr>::from_str(value) {
                    std::result::Result::Ok(value) => {
                        std::result::Result::Ok(header::IntoHeaderValue(value))
                    }
                    std::result::Result::Err(err) => std::result::Result::Err(format!(
                        "Unable to convert header value '{}' into Get200ResponseMemory - {}",
                        value, err
                    )),
                }
            }
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Unable to convert header: {:?} to string: {}",
                hdr_value, e
            )),
        }
    }
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct PipelineGet200Response {
    #[serde(rename = "status")]
    pub status: models::PipelineStatus,

    /// Milliseconds since epoch when the pipeline was manually started
    #[serde(rename = "last_start")]
    #[serde(skip_serializing_if = "Option::is_none")]
    pub last_start: Option<i64>,

    /// Milliseconds since epoch when the pipeline was manually stopped
    #[serde(rename = "last_stop")]
    #[serde(skip_serializing_if = "Option::is_none")]
    pub last_stop: Option<i64>,

    #[serde(rename = "stopping_service")]
    #[serde(skip_serializing_if = "Option::is_none")]
    pub stopping_service: Option<models::PipelineGet200ResponseStoppingService>,

    /// Milliseconds since epoch when the pipeline was automatically restarted (on process faults)
    #[serde(rename = "last_restart")]
    #[serde(skip_serializing_if = "Option::is_none")]
    pub last_restart: Option<i64>,

    /// The list of fully qualified services that are enabled in this pipeline. If the pipeline was started, this includes a process for each service
    #[serde(rename = "enabled")]
    pub enabled: Vec<models::PipelineGet200ResponseEnabledInner>,
}

impl PipelineGet200Response {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(
        status: models::PipelineStatus,
        enabled: Vec<models::PipelineGet200ResponseEnabledInner>,
    ) -> PipelineGet200Response {
        PipelineGet200Response {
            status,
            last_start: None,
            last_stop: None,
            stopping_service: None,
            last_restart: None,
            enabled,
        }
    }
}

/// Converts the PipelineGet200Response value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for PipelineGet200Response {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            // Skipping status in query parameter serialization
            self.last_start
                .as_ref()
                .map(|last_start| ["last_start".to_string(), last_start.to_string()].join(",")),
            self.last_stop
                .as_ref()
                .map(|last_stop| ["last_stop".to_string(), last_stop.to_string()].join(",")),
            // Skipping stopping_service in query parameter serialization
            self.last_restart.as_ref().map(|last_restart| {
                ["last_restart".to_string(), last_restart.to_string()].join(",")
            }),
            // Skipping enabled in query parameter serialization
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a PipelineGet200Response value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for PipelineGet200Response {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub status: Vec<models::PipelineStatus>,
            pub last_start: Vec<i64>,
            pub last_stop: Vec<i64>,
            pub stopping_service: Vec<models::PipelineGet200ResponseStoppingService>,
            pub last_restart: Vec<i64>,
            pub enabled: Vec<Vec<models::PipelineGet200ResponseEnabledInner>>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing PipelineGet200Response".to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "status" => intermediate_rep.status.push(<models::PipelineStatus as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?),
                    #[allow(clippy::redundant_clone)]
                    "last_start" => intermediate_rep.last_start.push(<i64 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?),
                    #[allow(clippy::redundant_clone)]
                    "last_stop" => intermediate_rep.last_stop.push(<i64 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?),
                    #[allow(clippy::redundant_clone)]
                    "stopping_service" => intermediate_rep.stopping_service.push(<models::PipelineGet200ResponseStoppingService as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?),
                    #[allow(clippy::redundant_clone)]
                    "last_restart" => intermediate_rep.last_restart.push(<i64 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?),
                    "enabled" => return std::result::Result::Err("Parsing a container in this style is not supported in PipelineGet200Response".to_string()),
                    _ => return std::result::Result::Err("Unexpected key while parsing PipelineGet200Response".to_string())
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(PipelineGet200Response {
            status: intermediate_rep
                .status
                .into_iter()
                .next()
                .ok_or_else(|| "status missing in PipelineGet200Response".to_string())?,
            last_start: intermediate_rep.last_start.into_iter().next(),
            last_stop: intermediate_rep.last_stop.into_iter().next(),
            stopping_service: intermediate_rep.stopping_service.into_iter().next(),
            last_restart: intermediate_rep.last_restart.into_iter().next(),
            enabled: intermediate_rep
                .enabled
                .into_iter()
                .next()
                .ok_or_else(|| "enabled missing in PipelineGet200Response".to_string())?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<PipelineGet200Response> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<PipelineGet200Response>> for HeaderValue {
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<PipelineGet200Response>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
            std::result::Result::Ok(value) => std::result::Result::Ok(value),
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Invalid header value for PipelineGet200Response - value: {} is invalid {}",
                hdr_value, e
            )),
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue> for header::IntoHeaderValue<PipelineGet200Response> {
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
            std::result::Result::Ok(value) => {
                match <PipelineGet200Response as std::str::FromStr>::from_str(value) {
                    std::result::Result::Ok(value) => {
                        std::result::Result::Ok(header::IntoHeaderValue(value))
                    }
                    std::result::Result::Err(err) => std::result::Result::Err(format!(
                        "Unable to convert header value '{}' into PipelineGet200Response - {}",
                        value, err
                    )),
                }
            }
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Unable to convert header: {:?} to string: {}",
                hdr_value, e
            )),
        }
    }
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct PipelineGet200ResponseEnabledInner {
    #[serde(rename = "service")]
    pub service: models::PipelineGet200ResponseEnabledInnerService,

    #[serde(rename = "process")]
    #[serde(skip_serializing_if = "Option::is_none")]
    pub process: Option<models::PipelineGet200ResponseEnabledInnerProcess>,
}

impl PipelineGet200ResponseEnabledInner {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(
        service: models::PipelineGet200ResponseEnabledInnerService,
    ) -> PipelineGet200ResponseEnabledInner {
        PipelineGet200ResponseEnabledInner {
            service,
            process: None,
        }
    }
}

/// Converts the PipelineGet200ResponseEnabledInner value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for PipelineGet200ResponseEnabledInner {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            // Skipping service in query parameter serialization

            // Skipping process in query parameter serialization

        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a PipelineGet200ResponseEnabledInner value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for PipelineGet200ResponseEnabledInner {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub service: Vec<models::PipelineGet200ResponseEnabledInnerService>,
            pub process: Vec<models::PipelineGet200ResponseEnabledInnerProcess>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing PipelineGet200ResponseEnabledInner"
                            .to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "service" => intermediate_rep.service.push(<models::PipelineGet200ResponseEnabledInnerService as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?),
                    #[allow(clippy::redundant_clone)]
                    "process" => intermediate_rep.process.push(<models::PipelineGet200ResponseEnabledInnerProcess as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?),
                    _ => return std::result::Result::Err("Unexpected key while parsing PipelineGet200ResponseEnabledInner".to_string())
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(PipelineGet200ResponseEnabledInner {
            service: intermediate_rep.service.into_iter().next().ok_or_else(|| {
                "service missing in PipelineGet200ResponseEnabledInner".to_string()
            })?,
            process: intermediate_rep.process.into_iter().next(),
        })
    }
}

// Methods for converting between header::IntoHeaderValue<PipelineGet200ResponseEnabledInner> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<PipelineGet200ResponseEnabledInner>>
    for HeaderValue
{
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<PipelineGet200ResponseEnabledInner>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
             std::result::Result::Ok(value) => std::result::Result::Ok(value),
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Invalid header value for PipelineGet200ResponseEnabledInner - value: {} is invalid {}",
                     hdr_value, e))
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue>
    for header::IntoHeaderValue<PipelineGet200ResponseEnabledInner>
{
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
             std::result::Result::Ok(value) => {
                    match <PipelineGet200ResponseEnabledInner as std::str::FromStr>::from_str(value) {
                        std::result::Result::Ok(value) => std::result::Result::Ok(header::IntoHeaderValue(value)),
                        std::result::Result::Err(err) => std::result::Result::Err(
                            format!("Unable to convert header value '{}' into PipelineGet200ResponseEnabledInner - {}",
                                value, err))
                    }
             },
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Unable to convert header: {:?} to string: {}",
                     hdr_value, e))
        }
    }
}

/// The last process that was started for this service (instantiated from the service). This can be undefined if the pipeline was not started before.
#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct PipelineGet200ResponseEnabledInnerProcess {
    /// The process ID. Depending on the status, this PID might not exist anymore
    #[serde(rename = "pid")]
    pub pid: i32,

    #[serde(rename = "status")]
    pub status: models::ProcessStatus,

    /// The number of milliseconds the process has been running
    #[serde(rename = "uptime")]
    pub uptime: i64,

    /// The amount of memory used by the process in megabytes
    #[serde(rename = "memory")]
    pub memory: i32,

    /// The percentage of CPU used by the process
    #[serde(rename = "cpu")]
    pub cpu: i32,
}

impl PipelineGet200ResponseEnabledInnerProcess {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(
        pid: i32,
        status: models::ProcessStatus,
        uptime: i64,
        memory: i32,
        cpu: i32,
    ) -> PipelineGet200ResponseEnabledInnerProcess {
        PipelineGet200ResponseEnabledInnerProcess {
            pid,
            status,
            uptime,
            memory,
            cpu,
        }
    }
}

/// Converts the PipelineGet200ResponseEnabledInnerProcess value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for PipelineGet200ResponseEnabledInnerProcess {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            Some("pid".to_string()),
            Some(self.pid.to_string()),
            // Skipping status in query parameter serialization
            Some("uptime".to_string()),
            Some(self.uptime.to_string()),
            Some("memory".to_string()),
            Some(self.memory.to_string()),
            Some("cpu".to_string()),
            Some(self.cpu.to_string()),
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a PipelineGet200ResponseEnabledInnerProcess value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for PipelineGet200ResponseEnabledInnerProcess {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub pid: Vec<i32>,
            pub status: Vec<models::ProcessStatus>,
            pub uptime: Vec<i64>,
            pub memory: Vec<i32>,
            pub cpu: Vec<i32>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing PipelineGet200ResponseEnabledInnerProcess"
                            .to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "pid" => intermediate_rep.pid.push(
                        <i32 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "status" => intermediate_rep.status.push(
                        <models::ProcessStatus as std::str::FromStr>::from_str(val)
                            .map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "uptime" => intermediate_rep.uptime.push(
                        <i64 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "memory" => intermediate_rep.memory.push(
                        <i32 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "cpu" => intermediate_rep.cpu.push(
                        <i32 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    _ => return std::result::Result::Err(
                        "Unexpected key while parsing PipelineGet200ResponseEnabledInnerProcess"
                            .to_string(),
                    ),
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(PipelineGet200ResponseEnabledInnerProcess {
            pid: intermediate_rep.pid.into_iter().next().ok_or_else(|| {
                "pid missing in PipelineGet200ResponseEnabledInnerProcess".to_string()
            })?,
            status: intermediate_rep.status.into_iter().next().ok_or_else(|| {
                "status missing in PipelineGet200ResponseEnabledInnerProcess".to_string()
            })?,
            uptime: intermediate_rep.uptime.into_iter().next().ok_or_else(|| {
                "uptime missing in PipelineGet200ResponseEnabledInnerProcess".to_string()
            })?,
            memory: intermediate_rep.memory.into_iter().next().ok_or_else(|| {
                "memory missing in PipelineGet200ResponseEnabledInnerProcess".to_string()
            })?,
            cpu: intermediate_rep.cpu.into_iter().next().ok_or_else(|| {
                "cpu missing in PipelineGet200ResponseEnabledInnerProcess".to_string()
            })?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<PipelineGet200ResponseEnabledInnerProcess> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<PipelineGet200ResponseEnabledInnerProcess>>
    for HeaderValue
{
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<PipelineGet200ResponseEnabledInnerProcess>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
             std::result::Result::Ok(value) => std::result::Result::Ok(value),
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Invalid header value for PipelineGet200ResponseEnabledInnerProcess - value: {} is invalid {}",
                     hdr_value, e))
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue>
    for header::IntoHeaderValue<PipelineGet200ResponseEnabledInnerProcess>
{
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
             std::result::Result::Ok(value) => {
                    match <PipelineGet200ResponseEnabledInnerProcess as std::str::FromStr>::from_str(value) {
                        std::result::Result::Ok(value) => std::result::Result::Ok(header::IntoHeaderValue(value)),
                        std::result::Result::Err(err) => std::result::Result::Err(
                            format!("Unable to convert header value '{}' into PipelineGet200ResponseEnabledInnerProcess - {}",
                                value, err))
                    }
             },
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Unable to convert header: {:?} to string: {}",
                     hdr_value, e))
        }
    }
}

/// The fully qualified service that is enabled
#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct PipelineGet200ResponseEnabledInnerService {
    #[serde(rename = "fq")]
    pub fq: models::FullyQualifiedService,

    /// The number of faults that have occurred (causing the pipeline to restart) since pipeline.last_start
    #[serde(rename = "faults")]
    pub faults: i32,

    /// The most recent exit code returned by the process.
    #[serde(rename = "exit")]
    pub exit: i32,
}

impl PipelineGet200ResponseEnabledInnerService {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(
        fq: models::FullyQualifiedService,
        faults: i32,
        exit: i32,
    ) -> PipelineGet200ResponseEnabledInnerService {
        PipelineGet200ResponseEnabledInnerService { fq, faults, exit }
    }
}

/// Converts the PipelineGet200ResponseEnabledInnerService value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for PipelineGet200ResponseEnabledInnerService {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            // Skipping fq in query parameter serialization
            Some("faults".to_string()),
            Some(self.faults.to_string()),
            Some("exit".to_string()),
            Some(self.exit.to_string()),
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a PipelineGet200ResponseEnabledInnerService value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for PipelineGet200ResponseEnabledInnerService {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub fq: Vec<models::FullyQualifiedService>,
            pub faults: Vec<i32>,
            pub exit: Vec<i32>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing PipelineGet200ResponseEnabledInnerService"
                            .to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "fq" => intermediate_rep.fq.push(
                        <models::FullyQualifiedService as std::str::FromStr>::from_str(val)
                            .map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "faults" => intermediate_rep.faults.push(
                        <i32 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "exit" => intermediate_rep.exit.push(
                        <i32 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    _ => return std::result::Result::Err(
                        "Unexpected key while parsing PipelineGet200ResponseEnabledInnerService"
                            .to_string(),
                    ),
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(PipelineGet200ResponseEnabledInnerService {
            fq: intermediate_rep.fq.into_iter().next().ok_or_else(|| {
                "fq missing in PipelineGet200ResponseEnabledInnerService".to_string()
            })?,
            faults: intermediate_rep.faults.into_iter().next().ok_or_else(|| {
                "faults missing in PipelineGet200ResponseEnabledInnerService".to_string()
            })?,
            exit: intermediate_rep.exit.into_iter().next().ok_or_else(|| {
                "exit missing in PipelineGet200ResponseEnabledInnerService".to_string()
            })?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<PipelineGet200ResponseEnabledInnerService> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<PipelineGet200ResponseEnabledInnerService>>
    for HeaderValue
{
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<PipelineGet200ResponseEnabledInnerService>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
             std::result::Result::Ok(value) => std::result::Result::Ok(value),
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Invalid header value for PipelineGet200ResponseEnabledInnerService - value: {} is invalid {}",
                     hdr_value, e))
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue>
    for header::IntoHeaderValue<PipelineGet200ResponseEnabledInnerService>
{
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
             std::result::Result::Ok(value) => {
                    match <PipelineGet200ResponseEnabledInnerService as std::str::FromStr>::from_str(value) {
                        std::result::Result::Ok(value) => std::result::Result::Ok(header::IntoHeaderValue(value)),
                        std::result::Result::Err(err) => std::result::Result::Err(
                            format!("Unable to convert header value '{}' into PipelineGet200ResponseEnabledInnerService - {}",
                                value, err))
                    }
             },
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Unable to convert header: {:?} to string: {}",
                     hdr_value, e))
        }
    }
}

/// The service that caused the pipeline to stop
#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct PipelineGet200ResponseStoppingService {
    #[serde(rename = "fq")]
    #[serde(skip_serializing_if = "Option::is_none")]
    pub fq: Option<models::FullyQualifiedService>,
}

impl PipelineGet200ResponseStoppingService {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new() -> PipelineGet200ResponseStoppingService {
        PipelineGet200ResponseStoppingService { fq: None }
    }
}

/// Converts the PipelineGet200ResponseStoppingService value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for PipelineGet200ResponseStoppingService {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            // Skipping fq in query parameter serialization

        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a PipelineGet200ResponseStoppingService value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for PipelineGet200ResponseStoppingService {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub fq: Vec<models::FullyQualifiedService>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing PipelineGet200ResponseStoppingService"
                            .to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "fq" => intermediate_rep.fq.push(
                        <models::FullyQualifiedService as std::str::FromStr>::from_str(val)
                            .map_err(|x| x.to_string())?,
                    ),
                    _ => {
                        return std::result::Result::Err(
                            "Unexpected key while parsing PipelineGet200ResponseStoppingService"
                                .to_string(),
                        )
                    }
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(PipelineGet200ResponseStoppingService {
            fq: intermediate_rep.fq.into_iter().next(),
        })
    }
}

// Methods for converting between header::IntoHeaderValue<PipelineGet200ResponseStoppingService> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<PipelineGet200ResponseStoppingService>>
    for HeaderValue
{
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<PipelineGet200ResponseStoppingService>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
             std::result::Result::Ok(value) => std::result::Result::Ok(value),
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Invalid header value for PipelineGet200ResponseStoppingService - value: {} is invalid {}",
                     hdr_value, e))
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue>
    for header::IntoHeaderValue<PipelineGet200ResponseStoppingService>
{
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
             std::result::Result::Ok(value) => {
                    match <PipelineGet200ResponseStoppingService as std::str::FromStr>::from_str(value) {
                        std::result::Result::Ok(value) => std::result::Result::Ok(header::IntoHeaderValue(value)),
                        std::result::Result::Err(err) => std::result::Result::Err(
                            format!("Unable to convert header value '{}' into PipelineGet200ResponseStoppingService - {}",
                                value, err))
                    }
             },
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Unable to convert header: {:?} to string: {}",
                     hdr_value, e))
        }
    }
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct PipelinePostRequestInner {
    #[serde(rename = "fq")]
    pub fq: models::FullyQualifiedService,
}

impl PipelinePostRequestInner {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(fq: models::FullyQualifiedService) -> PipelinePostRequestInner {
        PipelinePostRequestInner { fq }
    }
}

/// Converts the PipelinePostRequestInner value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for PipelinePostRequestInner {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            // Skipping fq in query parameter serialization

        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a PipelinePostRequestInner value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for PipelinePostRequestInner {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub fq: Vec<models::FullyQualifiedService>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing PipelinePostRequestInner".to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "fq" => intermediate_rep.fq.push(
                        <models::FullyQualifiedService as std::str::FromStr>::from_str(val)
                            .map_err(|x| x.to_string())?,
                    ),
                    _ => {
                        return std::result::Result::Err(
                            "Unexpected key while parsing PipelinePostRequestInner".to_string(),
                        )
                    }
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(PipelinePostRequestInner {
            fq: intermediate_rep
                .fq
                .into_iter()
                .next()
                .ok_or_else(|| "fq missing in PipelinePostRequestInner".to_string())?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<PipelinePostRequestInner> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<PipelinePostRequestInner>> for HeaderValue {
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<PipelinePostRequestInner>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
            std::result::Result::Ok(value) => std::result::Result::Ok(value),
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Invalid header value for PipelinePostRequestInner - value: {} is invalid {}",
                hdr_value, e
            )),
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue> for header::IntoHeaderValue<PipelinePostRequestInner> {
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
            std::result::Result::Ok(value) => {
                match <PipelinePostRequestInner as std::str::FromStr>::from_str(value) {
                    std::result::Result::Ok(value) => {
                        std::result::Result::Ok(header::IntoHeaderValue(value))
                    }
                    std::result::Result::Err(err) => std::result::Result::Err(format!(
                        "Unable to convert header value '{}' into PipelinePostRequestInner - {}",
                        value, err
                    )),
                }
            }
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Unable to convert header: {:?} to string: {}",
                hdr_value, e
            )),
        }
    }
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct PipelineSetError {
    #[serde(rename = "validation_errors")]
    pub validation_errors: models::PipelineSetErrorValidationErrors,
}

impl PipelineSetError {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(validation_errors: models::PipelineSetErrorValidationErrors) -> PipelineSetError {
        PipelineSetError { validation_errors }
    }
}

/// Converts the PipelineSetError value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for PipelineSetError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            // Skipping validation_errors in query parameter serialization

        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a PipelineSetError value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for PipelineSetError {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub validation_errors: Vec<models::PipelineSetErrorValidationErrors>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing PipelineSetError".to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "validation_errors" => intermediate_rep.validation_errors.push(
                        <models::PipelineSetErrorValidationErrors as std::str::FromStr>::from_str(
                            val,
                        )
                        .map_err(|x| x.to_string())?,
                    ),
                    _ => {
                        return std::result::Result::Err(
                            "Unexpected key while parsing PipelineSetError".to_string(),
                        )
                    }
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(PipelineSetError {
            validation_errors: intermediate_rep
                .validation_errors
                .into_iter()
                .next()
                .ok_or_else(|| "validation_errors missing in PipelineSetError".to_string())?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<PipelineSetError> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<PipelineSetError>> for HeaderValue {
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<PipelineSetError>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
            std::result::Result::Ok(value) => std::result::Result::Ok(value),
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Invalid header value for PipelineSetError - value: {} is invalid {}",
                hdr_value, e
            )),
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue> for header::IntoHeaderValue<PipelineSetError> {
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
            std::result::Result::Ok(value) => {
                match <PipelineSetError as std::str::FromStr>::from_str(value) {
                    std::result::Result::Ok(value) => {
                        std::result::Result::Ok(header::IntoHeaderValue(value))
                    }
                    std::result::Result::Err(err) => std::result::Result::Err(format!(
                        "Unable to convert header value '{}' into PipelineSetError - {}",
                        value, err
                    )),
                }
            }
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Unable to convert header: {:?} to string: {}",
                hdr_value, e
            )),
        }
    }
}

/// The validation errors that prevent the pipeline from being set
#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct PipelineSetErrorValidationErrors {
    #[serde(rename = "unmet_streams")]
    pub unmet_streams: Vec<models::UnmetStreamError>,

    #[serde(rename = "unmet_services")]
    pub unmet_services: Vec<models::UnmetServiceError>,

    /// List of duplicate services
    #[serde(rename = "duplicate_services")]
    pub duplicate_services: Vec<String>,

    /// List of duplicate aliases
    #[serde(rename = "duplicate_aliases")]
    pub duplicate_aliases: Vec<String>,

    /// List of aliases that are already used as a name of another service
    #[serde(rename = "aliases_in_use")]
    pub aliases_in_use: Vec<String>,
}

impl PipelineSetErrorValidationErrors {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(
        unmet_streams: Vec<models::UnmetStreamError>,
        unmet_services: Vec<models::UnmetServiceError>,
        duplicate_services: Vec<String>,
        duplicate_aliases: Vec<String>,
        aliases_in_use: Vec<String>,
    ) -> PipelineSetErrorValidationErrors {
        PipelineSetErrorValidationErrors {
            unmet_streams,
            unmet_services,
            duplicate_services,
            duplicate_aliases,
            aliases_in_use,
        }
    }
}

/// Converts the PipelineSetErrorValidationErrors value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for PipelineSetErrorValidationErrors {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            // Skipping unmet_streams in query parameter serialization

            // Skipping unmet_services in query parameter serialization
            Some("duplicate_services".to_string()),
            Some(
                self.duplicate_services
                    .iter()
                    .map(|x| x.to_string())
                    .collect::<Vec<_>>()
                    .join(","),
            ),
            Some("duplicate_aliases".to_string()),
            Some(
                self.duplicate_aliases
                    .iter()
                    .map(|x| x.to_string())
                    .collect::<Vec<_>>()
                    .join(","),
            ),
            Some("aliases_in_use".to_string()),
            Some(
                self.aliases_in_use
                    .iter()
                    .map(|x| x.to_string())
                    .collect::<Vec<_>>()
                    .join(","),
            ),
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a PipelineSetErrorValidationErrors value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for PipelineSetErrorValidationErrors {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub unmet_streams: Vec<Vec<models::UnmetStreamError>>,
            pub unmet_services: Vec<Vec<models::UnmetServiceError>>,
            pub duplicate_services: Vec<Vec<String>>,
            pub duplicate_aliases: Vec<Vec<String>>,
            pub aliases_in_use: Vec<Vec<String>>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing PipelineSetErrorValidationErrors".to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    "unmet_streams" => return std::result::Result::Err("Parsing a container in this style is not supported in PipelineSetErrorValidationErrors".to_string()),
                    "unmet_services" => return std::result::Result::Err("Parsing a container in this style is not supported in PipelineSetErrorValidationErrors".to_string()),
                    "duplicate_services" => return std::result::Result::Err("Parsing a container in this style is not supported in PipelineSetErrorValidationErrors".to_string()),
                    "duplicate_aliases" => return std::result::Result::Err("Parsing a container in this style is not supported in PipelineSetErrorValidationErrors".to_string()),
                    "aliases_in_use" => return std::result::Result::Err("Parsing a container in this style is not supported in PipelineSetErrorValidationErrors".to_string()),
                    _ => return std::result::Result::Err("Unexpected key while parsing PipelineSetErrorValidationErrors".to_string())
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(PipelineSetErrorValidationErrors {
            unmet_streams: intermediate_rep
                .unmet_streams
                .into_iter()
                .next()
                .ok_or_else(|| {
                    "unmet_streams missing in PipelineSetErrorValidationErrors".to_string()
                })?,
            unmet_services: intermediate_rep
                .unmet_services
                .into_iter()
                .next()
                .ok_or_else(|| {
                    "unmet_services missing in PipelineSetErrorValidationErrors".to_string()
                })?,
            duplicate_services: intermediate_rep
                .duplicate_services
                .into_iter()
                .next()
                .ok_or_else(|| {
                    "duplicate_services missing in PipelineSetErrorValidationErrors".to_string()
                })?,
            duplicate_aliases: intermediate_rep
                .duplicate_aliases
                .into_iter()
                .next()
                .ok_or_else(|| {
                    "duplicate_aliases missing in PipelineSetErrorValidationErrors".to_string()
                })?,
            aliases_in_use: intermediate_rep
                .aliases_in_use
                .into_iter()
                .next()
                .ok_or_else(|| {
                    "aliases_in_use missing in PipelineSetErrorValidationErrors".to_string()
                })?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<PipelineSetErrorValidationErrors> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<PipelineSetErrorValidationErrors>>
    for HeaderValue
{
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<PipelineSetErrorValidationErrors>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
             std::result::Result::Ok(value) => std::result::Result::Ok(value),
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Invalid header value for PipelineSetErrorValidationErrors - value: {} is invalid {}",
                     hdr_value, e))
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue>
    for header::IntoHeaderValue<PipelineSetErrorValidationErrors>
{
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
             std::result::Result::Ok(value) => {
                    match <PipelineSetErrorValidationErrors as std::str::FromStr>::from_str(value) {
                        std::result::Result::Ok(value) => std::result::Result::Ok(header::IntoHeaderValue(value)),
                        std::result::Result::Err(err) => std::result::Result::Err(
                            format!("Unable to convert header value '{}' into PipelineSetErrorValidationErrors - {}",
                                value, err))
                    }
             },
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Unable to convert header: {:?} to string: {}",
                     hdr_value, e))
        }
    }
}

/// The status of the entire pipeline corresponding to a state machine
/// Enumeration of values.
/// Since this enum's variants do not hold data, we can easily define them as `#[repr(C)]`
/// which helps with FFI.
#[allow(non_camel_case_types)]
#[repr(C)]
#[derive(
    Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, serde::Serialize, serde::Deserialize,
)]
#[cfg_attr(feature = "conversion", derive(frunk_enum_derive::LabelledGenericEnum))]
pub enum PipelineStatus {
    #[serde(rename = "empty")]
    Empty,
    #[serde(rename = "startable")]
    Startable,
    #[serde(rename = "started")]
    Started,
}

impl std::fmt::Display for PipelineStatus {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match *self {
            PipelineStatus::Empty => write!(f, "empty"),
            PipelineStatus::Startable => write!(f, "startable"),
            PipelineStatus::Started => write!(f, "started"),
        }
    }
}

impl std::str::FromStr for PipelineStatus {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        match s {
            "empty" => std::result::Result::Ok(PipelineStatus::Empty),
            "startable" => std::result::Result::Ok(PipelineStatus::Startable),
            "started" => std::result::Result::Ok(PipelineStatus::Started),
            _ => std::result::Result::Err(format!("Value not valid: {}", s)),
        }
    }
}

/// The status of a process in the pipeline
/// Enumeration of values.
/// Since this enum's variants do not hold data, we can easily define them as `#[repr(C)]`
/// which helps with FFI.
#[allow(non_camel_case_types)]
#[repr(C)]
#[derive(
    Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, serde::Serialize, serde::Deserialize,
)]
#[cfg_attr(feature = "conversion", derive(frunk_enum_derive::LabelledGenericEnum))]
pub enum ProcessStatus {
    #[serde(rename = "running")]
    Running,
    #[serde(rename = "stopped")]
    Stopped,
    #[serde(rename = "terminated")]
    Terminated,
    #[serde(rename = "killed")]
    Killed,
}

impl std::fmt::Display for ProcessStatus {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match *self {
            ProcessStatus::Running => write!(f, "running"),
            ProcessStatus::Stopped => write!(f, "stopped"),
            ProcessStatus::Terminated => write!(f, "terminated"),
            ProcessStatus::Killed => write!(f, "killed"),
        }
    }
}

impl std::str::FromStr for ProcessStatus {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        match s {
            "running" => std::result::Result::Ok(ProcessStatus::Running),
            "stopped" => std::result::Result::Ok(ProcessStatus::Stopped),
            "terminated" => std::result::Result::Ok(ProcessStatus::Terminated),
            "killed" => std::result::Result::Ok(ProcessStatus::Killed),
            _ => std::result::Result::Err(format!("Value not valid: {}", s)),
        }
    }
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct ReferencedService {
    /// Fully qualified download url.
    #[serde(rename = "url")]
    pub url: String,
}

impl ReferencedService {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(url: String) -> ReferencedService {
        ReferencedService { url }
    }
}

/// Converts the ReferencedService value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for ReferencedService {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![Some("url".to_string()), Some(self.url.to_string())];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a ReferencedService value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for ReferencedService {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub url: Vec<String>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing ReferencedService".to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "url" => intermediate_rep.url.push(
                        <String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    _ => {
                        return std::result::Result::Err(
                            "Unexpected key while parsing ReferencedService".to_string(),
                        )
                    }
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(ReferencedService {
            url: intermediate_rep
                .url
                .into_iter()
                .next()
                .ok_or_else(|| "url missing in ReferencedService".to_string())?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<ReferencedService> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<ReferencedService>> for HeaderValue {
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<ReferencedService>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
            std::result::Result::Ok(value) => std::result::Result::Ok(value),
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Invalid header value for ReferencedService - value: {} is invalid {}",
                hdr_value, e
            )),
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue> for header::IntoHeaderValue<ReferencedService> {
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
            std::result::Result::Ok(value) => {
                match <ReferencedService as std::str::FromStr>::from_str(value) {
                    std::result::Result::Ok(value) => {
                        std::result::Result::Ok(header::IntoHeaderValue(value))
                    }
                    std::result::Result::Err(err) => std::result::Result::Err(format!(
                        "Unable to convert header value '{}' into ReferencedService - {}",
                        value, err
                    )),
                }
            }
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Unable to convert header: {:?} to string: {}",
                hdr_value, e
            )),
        }
    }
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct RoverdError {
    /// Type of error
    /// Note: inline enums are not fully supported by openapi-generator
    #[serde(rename = "errorType")]
    pub error_type: String,

    #[serde(rename = "errorValue")]
    pub error_value: models::RoverdErrorErrorValue,
}

impl RoverdError {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(error_type: String, error_value: models::RoverdErrorErrorValue) -> RoverdError {
        RoverdError {
            error_type,
            error_value,
        }
    }
}

/// Converts the RoverdError value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for RoverdError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            Some("errorType".to_string()),
            Some(self.error_type.to_string()),
            // Skipping errorValue in query parameter serialization
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a RoverdError value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for RoverdError {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub error_type: Vec<String>,
            pub error_value: Vec<models::RoverdErrorErrorValue>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing RoverdError".to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "errorType" => intermediate_rep.error_type.push(
                        <String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "errorValue" => intermediate_rep.error_value.push(
                        <models::RoverdErrorErrorValue as std::str::FromStr>::from_str(val)
                            .map_err(|x| x.to_string())?,
                    ),
                    _ => {
                        return std::result::Result::Err(
                            "Unexpected key while parsing RoverdError".to_string(),
                        )
                    }
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(RoverdError {
            error_type: intermediate_rep
                .error_type
                .into_iter()
                .next()
                .ok_or_else(|| "errorType missing in RoverdError".to_string())?,
            error_value: intermediate_rep
                .error_value
                .into_iter()
                .next()
                .ok_or_else(|| "errorValue missing in RoverdError".to_string())?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<RoverdError> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<RoverdError>> for HeaderValue {
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<RoverdError>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
            std::result::Result::Ok(value) => std::result::Result::Ok(value),
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Invalid header value for RoverdError - value: {} is invalid {}",
                hdr_value, e
            )),
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue> for header::IntoHeaderValue<RoverdError> {
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
            std::result::Result::Ok(value) => {
                match <RoverdError as std::str::FromStr>::from_str(value) {
                    std::result::Result::Ok(value) => {
                        std::result::Result::Ok(header::IntoHeaderValue(value))
                    }
                    std::result::Result::Err(err) => std::result::Result::Err(format!(
                        "Unable to convert header value '{}' into RoverdError - {}",
                        value, err
                    )),
                }
            }
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Unable to convert header: {:?} to string: {}",
                hdr_value, e
            )),
        }
    }
}

/// One of:
/// - BuildError
/// - GenericError
/// - PipelineSetError
#[derive(Debug, Clone, serde::Serialize, serde::Deserialize)]
pub struct RoverdErrorErrorValue(pub Box<serde_json::value::RawValue>);

impl validator::Validate for RoverdErrorErrorValue {
    fn validate(&self) -> std::result::Result<(), validator::ValidationErrors> {
        std::result::Result::Ok(())
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a RoverdErrorErrorValue value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for RoverdErrorErrorValue {
    type Err = serde_json::Error;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        serde_json::from_str(s)
    }
}

impl PartialEq for RoverdErrorErrorValue {
    fn eq(&self, other: &Self) -> bool {
        self.0.get() == other.0.get()
    }
}

/// The status of any given service is either enabled or disabled
/// Enumeration of values.
/// Since this enum's variants do not hold data, we can easily define them as `#[repr(C)]`
/// which helps with FFI.
#[allow(non_camel_case_types)]
#[repr(C)]
#[derive(
    Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, serde::Serialize, serde::Deserialize,
)]
#[cfg_attr(feature = "conversion", derive(frunk_enum_derive::LabelledGenericEnum))]
pub enum ServiceStatus {
    #[serde(rename = "enabled")]
    Enabled,
    #[serde(rename = "disabled")]
    Disabled,
}

impl std::fmt::Display for ServiceStatus {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match *self {
            ServiceStatus::Enabled => write!(f, "enabled"),
            ServiceStatus::Disabled => write!(f, "disabled"),
        }
    }
}

impl std::str::FromStr for ServiceStatus {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        match s {
            "enabled" => std::result::Result::Ok(ServiceStatus::Enabled),
            "disabled" => std::result::Result::Ok(ServiceStatus::Disabled),
            _ => std::result::Result::Err(format!("Value not valid: {}", s)),
        }
    }
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct ServicesAuthorServiceVersionConfigurationPostRequestInner {
    /// The unique key, corresponding to a configuration key in the service.yaml file
    #[serde(rename = "key")]
    pub key: String,

    #[serde(rename = "value")]
    pub value: models::ServicesAuthorServiceVersionConfigurationPostRequestInnerValue,
}

impl ServicesAuthorServiceVersionConfigurationPostRequestInner {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(
        key: String,
        value: models::ServicesAuthorServiceVersionConfigurationPostRequestInnerValue,
    ) -> ServicesAuthorServiceVersionConfigurationPostRequestInner {
        ServicesAuthorServiceVersionConfigurationPostRequestInner { key, value }
    }
}

/// Converts the ServicesAuthorServiceVersionConfigurationPostRequestInner value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for ServicesAuthorServiceVersionConfigurationPostRequestInner {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            Some("key".to_string()),
            Some(self.key.to_string()),
            // Skipping value in query parameter serialization
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a ServicesAuthorServiceVersionConfigurationPostRequestInner value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for ServicesAuthorServiceVersionConfigurationPostRequestInner {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub key: Vec<String>,
            pub value: Vec<models::ServicesAuthorServiceVersionConfigurationPostRequestInnerValue>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => return std::result::Result::Err("Missing value while parsing ServicesAuthorServiceVersionConfigurationPostRequestInner".to_string())
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "key" => intermediate_rep.key.push(<String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?),
                    #[allow(clippy::redundant_clone)]
                    "value" => intermediate_rep.value.push(<models::ServicesAuthorServiceVersionConfigurationPostRequestInnerValue as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?),
                    _ => return std::result::Result::Err("Unexpected key while parsing ServicesAuthorServiceVersionConfigurationPostRequestInner".to_string())
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(ServicesAuthorServiceVersionConfigurationPostRequestInner {
            key: intermediate_rep.key.into_iter().next().ok_or_else(|| {
                "key missing in ServicesAuthorServiceVersionConfigurationPostRequestInner"
                    .to_string()
            })?,
            value: intermediate_rep.value.into_iter().next().ok_or_else(|| {
                "value missing in ServicesAuthorServiceVersionConfigurationPostRequestInner"
                    .to_string()
            })?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<ServicesAuthorServiceVersionConfigurationPostRequestInner> and HeaderValue

#[cfg(feature = "server")]
impl
    std::convert::TryFrom<
        header::IntoHeaderValue<ServicesAuthorServiceVersionConfigurationPostRequestInner>,
    > for HeaderValue
{
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<
            ServicesAuthorServiceVersionConfigurationPostRequestInner,
        >,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
             std::result::Result::Ok(value) => std::result::Result::Ok(value),
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Invalid header value for ServicesAuthorServiceVersionConfigurationPostRequestInner - value: {} is invalid {}",
                     hdr_value, e))
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue>
    for header::IntoHeaderValue<ServicesAuthorServiceVersionConfigurationPostRequestInner>
{
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
             std::result::Result::Ok(value) => {
                    match <ServicesAuthorServiceVersionConfigurationPostRequestInner as std::str::FromStr>::from_str(value) {
                        std::result::Result::Ok(value) => std::result::Result::Ok(header::IntoHeaderValue(value)),
                        std::result::Result::Err(err) => std::result::Result::Err(
                            format!("Unable to convert header value '{}' into ServicesAuthorServiceVersionConfigurationPostRequestInner - {}",
                                value, err))
                    }
             },
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Unable to convert header: {:?} to string: {}",
                     hdr_value, e))
        }
    }
}

/// The value that should be set for this key. Can be either a string or a number, but must match the type in the service.yaml file
/// One of:
/// - String
/// - f64
#[derive(Debug, Clone, serde::Serialize, serde::Deserialize)]
pub struct ServicesAuthorServiceVersionConfigurationPostRequestInnerValue(
    pub Box<serde_json::value::RawValue>,
);

impl validator::Validate for ServicesAuthorServiceVersionConfigurationPostRequestInnerValue {
    fn validate(&self) -> std::result::Result<(), validator::ValidationErrors> {
        std::result::Result::Ok(())
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a ServicesAuthorServiceVersionConfigurationPostRequestInnerValue value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for ServicesAuthorServiceVersionConfigurationPostRequestInnerValue {
    type Err = serde_json::Error;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        serde_json::from_str(s)
    }
}

impl PartialEq for ServicesAuthorServiceVersionConfigurationPostRequestInnerValue {
    fn eq(&self, other: &Self) -> bool {
        self.0.get() == other.0.get()
    }
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct ServicesAuthorServiceVersionDelete200Response {
    /// Whether the pipeline was invalidated by this service deletion
    #[serde(rename = "invalidated_pipeline")]
    pub invalidated_pipeline: bool,
}

impl ServicesAuthorServiceVersionDelete200Response {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(invalidated_pipeline: bool) -> ServicesAuthorServiceVersionDelete200Response {
        ServicesAuthorServiceVersionDelete200Response {
            invalidated_pipeline,
        }
    }
}

/// Converts the ServicesAuthorServiceVersionDelete200Response value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for ServicesAuthorServiceVersionDelete200Response {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            Some("invalidated_pipeline".to_string()),
            Some(self.invalidated_pipeline.to_string()),
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a ServicesAuthorServiceVersionDelete200Response value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for ServicesAuthorServiceVersionDelete200Response {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub invalidated_pipeline: Vec<bool>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val =
                match string_iter.next() {
                    Some(x) => x,
                    None => return std::result::Result::Err(
                        "Missing value while parsing ServicesAuthorServiceVersionDelete200Response"
                            .to_string(),
                    ),
                };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "invalidated_pipeline" => intermediate_rep.invalidated_pipeline.push(<bool as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?),
                    _ => return std::result::Result::Err("Unexpected key while parsing ServicesAuthorServiceVersionDelete200Response".to_string())
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(ServicesAuthorServiceVersionDelete200Response {
            invalidated_pipeline: intermediate_rep
                .invalidated_pipeline
                .into_iter()
                .next()
                .ok_or_else(|| {
                    "invalidated_pipeline missing in ServicesAuthorServiceVersionDelete200Response"
                        .to_string()
                })?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<ServicesAuthorServiceVersionDelete200Response> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<ServicesAuthorServiceVersionDelete200Response>>
    for HeaderValue
{
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<ServicesAuthorServiceVersionDelete200Response>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
             std::result::Result::Ok(value) => std::result::Result::Ok(value),
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Invalid header value for ServicesAuthorServiceVersionDelete200Response - value: {} is invalid {}",
                     hdr_value, e))
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue>
    for header::IntoHeaderValue<ServicesAuthorServiceVersionDelete200Response>
{
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
             std::result::Result::Ok(value) => {
                    match <ServicesAuthorServiceVersionDelete200Response as std::str::FromStr>::from_str(value) {
                        std::result::Result::Ok(value) => std::result::Result::Ok(header::IntoHeaderValue(value)),
                        std::result::Result::Err(err) => std::result::Result::Err(
                            format!("Unable to convert header value '{}' into ServicesAuthorServiceVersionDelete200Response - {}",
                                value, err))
                    }
             },
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Unable to convert header: {:?} to string: {}",
                     hdr_value, e))
        }
    }
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct ServicesAuthorServiceVersionGet200Response {
    /// The time this version was last built as milliseconds since epoch, not set if the service was never built
    #[serde(rename = "built_at")]
    #[serde(skip_serializing_if = "Option::is_none")]
    pub built_at: Option<i64>,

    /// The dependencies/inputs of this service version
    #[serde(rename = "inputs")]
    pub inputs: Vec<models::ServicesAuthorServiceVersionGet200ResponseInputsInner>,

    /// The output streams of this service version
    #[serde(rename = "outputs")]
    pub outputs: Vec<String>,

    /// All configuration values of this service version and their tunability
    #[serde(rename = "configuration")]
    pub configuration: Vec<models::ServicesAuthorServiceVersionGet200ResponseConfigurationInner>,
}

impl ServicesAuthorServiceVersionGet200Response {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(
        inputs: Vec<models::ServicesAuthorServiceVersionGet200ResponseInputsInner>,
        outputs: Vec<String>,
        configuration: Vec<models::ServicesAuthorServiceVersionGet200ResponseConfigurationInner>,
    ) -> ServicesAuthorServiceVersionGet200Response {
        ServicesAuthorServiceVersionGet200Response {
            built_at: None,
            inputs,
            outputs,
            configuration,
        }
    }
}

/// Converts the ServicesAuthorServiceVersionGet200Response value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for ServicesAuthorServiceVersionGet200Response {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            self.built_at
                .as_ref()
                .map(|built_at| ["built_at".to_string(), built_at.to_string()].join(",")),
            // Skipping inputs in query parameter serialization
            Some("outputs".to_string()),
            Some(
                self.outputs
                    .iter()
                    .map(|x| x.to_string())
                    .collect::<Vec<_>>()
                    .join(","),
            ),
            // Skipping configuration in query parameter serialization
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a ServicesAuthorServiceVersionGet200Response value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for ServicesAuthorServiceVersionGet200Response {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub built_at: Vec<i64>,
            pub inputs: Vec<Vec<models::ServicesAuthorServiceVersionGet200ResponseInputsInner>>,
            pub outputs: Vec<Vec<String>>,
            pub configuration:
                Vec<Vec<models::ServicesAuthorServiceVersionGet200ResponseConfigurationInner>>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val =
                match string_iter.next() {
                    Some(x) => x,
                    None => return std::result::Result::Err(
                        "Missing value while parsing ServicesAuthorServiceVersionGet200Response"
                            .to_string(),
                    ),
                };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "built_at" => intermediate_rep.built_at.push(<i64 as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?),
                    "inputs" => return std::result::Result::Err("Parsing a container in this style is not supported in ServicesAuthorServiceVersionGet200Response".to_string()),
                    "outputs" => return std::result::Result::Err("Parsing a container in this style is not supported in ServicesAuthorServiceVersionGet200Response".to_string()),
                    "configuration" => return std::result::Result::Err("Parsing a container in this style is not supported in ServicesAuthorServiceVersionGet200Response".to_string()),
                    _ => return std::result::Result::Err("Unexpected key while parsing ServicesAuthorServiceVersionGet200Response".to_string())
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(ServicesAuthorServiceVersionGet200Response {
            built_at: intermediate_rep.built_at.into_iter().next(),
            inputs: intermediate_rep.inputs.into_iter().next().ok_or_else(|| {
                "inputs missing in ServicesAuthorServiceVersionGet200Response".to_string()
            })?,
            outputs: intermediate_rep.outputs.into_iter().next().ok_or_else(|| {
                "outputs missing in ServicesAuthorServiceVersionGet200Response".to_string()
            })?,
            configuration: intermediate_rep
                .configuration
                .into_iter()
                .next()
                .ok_or_else(|| {
                    "configuration missing in ServicesAuthorServiceVersionGet200Response"
                        .to_string()
                })?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<ServicesAuthorServiceVersionGet200Response> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<ServicesAuthorServiceVersionGet200Response>>
    for HeaderValue
{
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<ServicesAuthorServiceVersionGet200Response>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
             std::result::Result::Ok(value) => std::result::Result::Ok(value),
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Invalid header value for ServicesAuthorServiceVersionGet200Response - value: {} is invalid {}",
                     hdr_value, e))
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue>
    for header::IntoHeaderValue<ServicesAuthorServiceVersionGet200Response>
{
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
             std::result::Result::Ok(value) => {
                    match <ServicesAuthorServiceVersionGet200Response as std::str::FromStr>::from_str(value) {
                        std::result::Result::Ok(value) => std::result::Result::Ok(header::IntoHeaderValue(value)),
                        std::result::Result::Err(err) => std::result::Result::Err(
                            format!("Unable to convert header value '{}' into ServicesAuthorServiceVersionGet200Response - {}",
                                value, err))
                    }
             },
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Unable to convert header: {:?} to string: {}",
                     hdr_value, e))
        }
    }
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct ServicesAuthorServiceVersionGet200ResponseConfigurationInner {
    /// The name of the configuration value
    #[serde(rename = "name")]
    pub name: String,

    /// The type of the configuration value
    /// Note: inline enums are not fully supported by openapi-generator
    #[serde(rename = "type")]
    pub r#type: String,

    #[serde(rename = "value")]
    pub value: models::ServicesAuthorServiceVersionGet200ResponseConfigurationInnerValue,

    /// Whether this configuration value is tunable
    #[serde(rename = "tunable")]
    pub tunable: bool,
}

impl ServicesAuthorServiceVersionGet200ResponseConfigurationInner {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(
        name: String,
        r#type: String,
        value: models::ServicesAuthorServiceVersionGet200ResponseConfigurationInnerValue,
        tunable: bool,
    ) -> ServicesAuthorServiceVersionGet200ResponseConfigurationInner {
        ServicesAuthorServiceVersionGet200ResponseConfigurationInner {
            name,
            r#type,
            value,
            tunable,
        }
    }
}

/// Converts the ServicesAuthorServiceVersionGet200ResponseConfigurationInner value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for ServicesAuthorServiceVersionGet200ResponseConfigurationInner {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            Some("name".to_string()),
            Some(self.name.to_string()),
            Some("type".to_string()),
            Some(self.r#type.to_string()),
            // Skipping value in query parameter serialization
            Some("tunable".to_string()),
            Some(self.tunable.to_string()),
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a ServicesAuthorServiceVersionGet200ResponseConfigurationInner value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for ServicesAuthorServiceVersionGet200ResponseConfigurationInner {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub name: Vec<String>,
            pub r#type: Vec<String>,
            pub value:
                Vec<models::ServicesAuthorServiceVersionGet200ResponseConfigurationInnerValue>,
            pub tunable: Vec<bool>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => return std::result::Result::Err("Missing value while parsing ServicesAuthorServiceVersionGet200ResponseConfigurationInner".to_string())
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "name" => intermediate_rep.name.push(<String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?),
                    #[allow(clippy::redundant_clone)]
                    "type" => intermediate_rep.r#type.push(<String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?),
                    #[allow(clippy::redundant_clone)]
                    "value" => intermediate_rep.value.push(<models::ServicesAuthorServiceVersionGet200ResponseConfigurationInnerValue as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?),
                    #[allow(clippy::redundant_clone)]
                    "tunable" => intermediate_rep.tunable.push(<bool as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?),
                    _ => return std::result::Result::Err("Unexpected key while parsing ServicesAuthorServiceVersionGet200ResponseConfigurationInner".to_string())
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(ServicesAuthorServiceVersionGet200ResponseConfigurationInner {
            name: intermediate_rep.name.into_iter().next().ok_or_else(|| "name missing in ServicesAuthorServiceVersionGet200ResponseConfigurationInner".to_string())?,
            r#type: intermediate_rep.r#type.into_iter().next().ok_or_else(|| "type missing in ServicesAuthorServiceVersionGet200ResponseConfigurationInner".to_string())?,
            value: intermediate_rep.value.into_iter().next().ok_or_else(|| "value missing in ServicesAuthorServiceVersionGet200ResponseConfigurationInner".to_string())?,
            tunable: intermediate_rep.tunable.into_iter().next().ok_or_else(|| "tunable missing in ServicesAuthorServiceVersionGet200ResponseConfigurationInner".to_string())?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<ServicesAuthorServiceVersionGet200ResponseConfigurationInner> and HeaderValue

#[cfg(feature = "server")]
impl
    std::convert::TryFrom<
        header::IntoHeaderValue<ServicesAuthorServiceVersionGet200ResponseConfigurationInner>,
    > for HeaderValue
{
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<
            ServicesAuthorServiceVersionGet200ResponseConfigurationInner,
        >,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
             std::result::Result::Ok(value) => std::result::Result::Ok(value),
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Invalid header value for ServicesAuthorServiceVersionGet200ResponseConfigurationInner - value: {} is invalid {}",
                     hdr_value, e))
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue>
    for header::IntoHeaderValue<ServicesAuthorServiceVersionGet200ResponseConfigurationInner>
{
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
             std::result::Result::Ok(value) => {
                    match <ServicesAuthorServiceVersionGet200ResponseConfigurationInner as std::str::FromStr>::from_str(value) {
                        std::result::Result::Ok(value) => std::result::Result::Ok(header::IntoHeaderValue(value)),
                        std::result::Result::Err(err) => std::result::Result::Err(
                            format!("Unable to convert header value '{}' into ServicesAuthorServiceVersionGet200ResponseConfigurationInner - {}",
                                value, err))
                    }
             },
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Unable to convert header: {:?} to string: {}",
                     hdr_value, e))
        }
    }
}

/// The value of the configuration
/// One of:
/// - String
/// - f64
#[derive(Debug, Clone, serde::Serialize, serde::Deserialize)]
pub struct ServicesAuthorServiceVersionGet200ResponseConfigurationInnerValue(
    pub Box<serde_json::value::RawValue>,
);

impl validator::Validate for ServicesAuthorServiceVersionGet200ResponseConfigurationInnerValue {
    fn validate(&self) -> std::result::Result<(), validator::ValidationErrors> {
        std::result::Result::Ok(())
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a ServicesAuthorServiceVersionGet200ResponseConfigurationInnerValue value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for ServicesAuthorServiceVersionGet200ResponseConfigurationInnerValue {
    type Err = serde_json::Error;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        serde_json::from_str(s)
    }
}

impl PartialEq for ServicesAuthorServiceVersionGet200ResponseConfigurationInnerValue {
    fn eq(&self, other: &Self) -> bool {
        self.0.get() == other.0.get()
    }
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct ServicesAuthorServiceVersionGet200ResponseInputsInner {
    /// The name of the service dependency
    #[serde(rename = "service")]
    pub service: String,

    /// The streams of the service dependency
    #[serde(rename = "streams")]
    pub streams: Vec<String>,
}

impl ServicesAuthorServiceVersionGet200ResponseInputsInner {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(
        service: String,
        streams: Vec<String>,
    ) -> ServicesAuthorServiceVersionGet200ResponseInputsInner {
        ServicesAuthorServiceVersionGet200ResponseInputsInner { service, streams }
    }
}

/// Converts the ServicesAuthorServiceVersionGet200ResponseInputsInner value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for ServicesAuthorServiceVersionGet200ResponseInputsInner {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            Some("service".to_string()),
            Some(self.service.to_string()),
            Some("streams".to_string()),
            Some(
                self.streams
                    .iter()
                    .map(|x| x.to_string())
                    .collect::<Vec<_>>()
                    .join(","),
            ),
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a ServicesAuthorServiceVersionGet200ResponseInputsInner value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for ServicesAuthorServiceVersionGet200ResponseInputsInner {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub service: Vec<String>,
            pub streams: Vec<Vec<String>>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => return std::result::Result::Err("Missing value while parsing ServicesAuthorServiceVersionGet200ResponseInputsInner".to_string())
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "service" => intermediate_rep.service.push(<String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?),
                    "streams" => return std::result::Result::Err("Parsing a container in this style is not supported in ServicesAuthorServiceVersionGet200ResponseInputsInner".to_string()),
                    _ => return std::result::Result::Err("Unexpected key while parsing ServicesAuthorServiceVersionGet200ResponseInputsInner".to_string())
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(ServicesAuthorServiceVersionGet200ResponseInputsInner {
            service: intermediate_rep.service.into_iter().next().ok_or_else(|| {
                "service missing in ServicesAuthorServiceVersionGet200ResponseInputsInner"
                    .to_string()
            })?,
            streams: intermediate_rep.streams.into_iter().next().ok_or_else(|| {
                "streams missing in ServicesAuthorServiceVersionGet200ResponseInputsInner"
                    .to_string()
            })?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<ServicesAuthorServiceVersionGet200ResponseInputsInner> and HeaderValue

#[cfg(feature = "server")]
impl
    std::convert::TryFrom<
        header::IntoHeaderValue<ServicesAuthorServiceVersionGet200ResponseInputsInner>,
    > for HeaderValue
{
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<ServicesAuthorServiceVersionGet200ResponseInputsInner>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
             std::result::Result::Ok(value) => std::result::Result::Ok(value),
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Invalid header value for ServicesAuthorServiceVersionGet200ResponseInputsInner - value: {} is invalid {}",
                     hdr_value, e))
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue>
    for header::IntoHeaderValue<ServicesAuthorServiceVersionGet200ResponseInputsInner>
{
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
             std::result::Result::Ok(value) => {
                    match <ServicesAuthorServiceVersionGet200ResponseInputsInner as std::str::FromStr>::from_str(value) {
                        std::result::Result::Ok(value) => std::result::Result::Ok(header::IntoHeaderValue(value)),
                        std::result::Result::Err(err) => std::result::Result::Err(
                            format!("Unable to convert header value '{}' into ServicesAuthorServiceVersionGet200ResponseInputsInner - {}",
                                value, err))
                    }
             },
             std::result::Result::Err(e) => std::result::Result::Err(
                 format!("Unable to convert header: {:?} to string: {}",
                     hdr_value, e))
        }
    }
}

/// UnmetServiceError
#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct UnmetServiceError {
    #[serde(rename = "source")]
    pub source: String,

    #[serde(rename = "target")]
    pub target: String,
}

impl UnmetServiceError {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(source: String, target: String) -> UnmetServiceError {
        UnmetServiceError { source, target }
    }
}

/// Converts the UnmetServiceError value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for UnmetServiceError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            Some("source".to_string()),
            Some(self.source.to_string()),
            Some("target".to_string()),
            Some(self.target.to_string()),
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a UnmetServiceError value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for UnmetServiceError {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub source: Vec<String>,
            pub target: Vec<String>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing UnmetServiceError".to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "source" => intermediate_rep.source.push(
                        <String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "target" => intermediate_rep.target.push(
                        <String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    _ => {
                        return std::result::Result::Err(
                            "Unexpected key while parsing UnmetServiceError".to_string(),
                        )
                    }
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(UnmetServiceError {
            source: intermediate_rep
                .source
                .into_iter()
                .next()
                .ok_or_else(|| "source missing in UnmetServiceError".to_string())?,
            target: intermediate_rep
                .target
                .into_iter()
                .next()
                .ok_or_else(|| "target missing in UnmetServiceError".to_string())?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<UnmetServiceError> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<UnmetServiceError>> for HeaderValue {
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<UnmetServiceError>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
            std::result::Result::Ok(value) => std::result::Result::Ok(value),
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Invalid header value for UnmetServiceError - value: {} is invalid {}",
                hdr_value, e
            )),
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue> for header::IntoHeaderValue<UnmetServiceError> {
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
            std::result::Result::Ok(value) => {
                match <UnmetServiceError as std::str::FromStr>::from_str(value) {
                    std::result::Result::Ok(value) => {
                        std::result::Result::Ok(header::IntoHeaderValue(value))
                    }
                    std::result::Result::Err(err) => std::result::Result::Err(format!(
                        "Unable to convert header value '{}' into UnmetServiceError - {}",
                        value, err
                    )),
                }
            }
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Unable to convert header: {:?} to string: {}",
                hdr_value, e
            )),
        }
    }
}

/// UnmetStreamError
#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct UnmetStreamError {
    #[serde(rename = "source")]
    pub source: String,

    #[serde(rename = "target")]
    pub target: String,

    #[serde(rename = "stream")]
    pub stream: String,
}

impl UnmetStreamError {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(source: String, target: String, stream: String) -> UnmetStreamError {
        UnmetStreamError {
            source,
            target,
            stream,
        }
    }
}

/// Converts the UnmetStreamError value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for UnmetStreamError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> = vec![
            Some("source".to_string()),
            Some(self.source.to_string()),
            Some("target".to_string()),
            Some(self.target.to_string()),
            Some("stream".to_string()),
            Some(self.stream.to_string()),
        ];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a UnmetStreamError value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for UnmetStreamError {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub source: Vec<String>,
            pub target: Vec<String>,
            pub stream: Vec<String>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing UnmetStreamError".to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "source" => intermediate_rep.source.push(
                        <String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "target" => intermediate_rep.target.push(
                        <String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    #[allow(clippy::redundant_clone)]
                    "stream" => intermediate_rep.stream.push(
                        <String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    _ => {
                        return std::result::Result::Err(
                            "Unexpected key while parsing UnmetStreamError".to_string(),
                        )
                    }
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(UnmetStreamError {
            source: intermediate_rep
                .source
                .into_iter()
                .next()
                .ok_or_else(|| "source missing in UnmetStreamError".to_string())?,
            target: intermediate_rep
                .target
                .into_iter()
                .next()
                .ok_or_else(|| "target missing in UnmetStreamError".to_string())?,
            stream: intermediate_rep
                .stream
                .into_iter()
                .next()
                .ok_or_else(|| "stream missing in UnmetStreamError".to_string())?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<UnmetStreamError> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<UnmetStreamError>> for HeaderValue {
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<UnmetStreamError>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
            std::result::Result::Ok(value) => std::result::Result::Ok(value),
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Invalid header value for UnmetStreamError - value: {} is invalid {}",
                hdr_value, e
            )),
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue> for header::IntoHeaderValue<UnmetStreamError> {
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
            std::result::Result::Ok(value) => {
                match <UnmetStreamError as std::str::FromStr>::from_str(value) {
                    std::result::Result::Ok(value) => {
                        std::result::Result::Ok(header::IntoHeaderValue(value))
                    }
                    std::result::Result::Err(err) => std::result::Result::Err(format!(
                        "Unable to convert header value '{}' into UnmetStreamError - {}",
                        value, err
                    )),
                }
            }
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Unable to convert header: {:?} to string: {}",
                hdr_value, e
            )),
        }
    }
}

#[derive(Debug, Clone, PartialEq, serde::Serialize, serde::Deserialize, validator::Validate)]
#[cfg_attr(feature = "conversion", derive(frunk::LabelledGeneric))]
pub struct UpdatePostRequest {
    /// The roverd version to update to
    #[serde(rename = "version")]
    pub version: String,
}

impl UpdatePostRequest {
    #[allow(clippy::new_without_default, clippy::too_many_arguments)]
    pub fn new(version: String) -> UpdatePostRequest {
        UpdatePostRequest { version }
    }
}

/// Converts the UpdatePostRequest value to the Query Parameters representation (style=form, explode=false)
/// specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde serializer
impl std::fmt::Display for UpdatePostRequest {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let params: Vec<Option<String>> =
            vec![Some("version".to_string()), Some(self.version.to_string())];

        write!(
            f,
            "{}",
            params.into_iter().flatten().collect::<Vec<_>>().join(",")
        )
    }
}

/// Converts Query Parameters representation (style=form, explode=false) to a UpdatePostRequest value
/// as specified in https://swagger.io/docs/specification/serialization/
/// Should be implemented in a serde deserializer
impl std::str::FromStr for UpdatePostRequest {
    type Err = String;

    fn from_str(s: &str) -> std::result::Result<Self, Self::Err> {
        /// An intermediate representation of the struct to use for parsing.
        #[derive(Default)]
        #[allow(dead_code)]
        struct IntermediateRep {
            pub version: Vec<String>,
        }

        let mut intermediate_rep = IntermediateRep::default();

        // Parse into intermediate representation
        let mut string_iter = s.split(',');
        let mut key_result = string_iter.next();

        while key_result.is_some() {
            let val = match string_iter.next() {
                Some(x) => x,
                None => {
                    return std::result::Result::Err(
                        "Missing value while parsing UpdatePostRequest".to_string(),
                    )
                }
            };

            if let Some(key) = key_result {
                #[allow(clippy::match_single_binding)]
                match key {
                    #[allow(clippy::redundant_clone)]
                    "version" => intermediate_rep.version.push(
                        <String as std::str::FromStr>::from_str(val).map_err(|x| x.to_string())?,
                    ),
                    _ => {
                        return std::result::Result::Err(
                            "Unexpected key while parsing UpdatePostRequest".to_string(),
                        )
                    }
                }
            }

            // Get the next key
            key_result = string_iter.next();
        }

        // Use the intermediate representation to return the struct
        std::result::Result::Ok(UpdatePostRequest {
            version: intermediate_rep
                .version
                .into_iter()
                .next()
                .ok_or_else(|| "version missing in UpdatePostRequest".to_string())?,
        })
    }
}

// Methods for converting between header::IntoHeaderValue<UpdatePostRequest> and HeaderValue

#[cfg(feature = "server")]
impl std::convert::TryFrom<header::IntoHeaderValue<UpdatePostRequest>> for HeaderValue {
    type Error = String;

    fn try_from(
        hdr_value: header::IntoHeaderValue<UpdatePostRequest>,
    ) -> std::result::Result<Self, Self::Error> {
        let hdr_value = hdr_value.to_string();
        match HeaderValue::from_str(&hdr_value) {
            std::result::Result::Ok(value) => std::result::Result::Ok(value),
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Invalid header value for UpdatePostRequest - value: {} is invalid {}",
                hdr_value, e
            )),
        }
    }
}

#[cfg(feature = "server")]
impl std::convert::TryFrom<HeaderValue> for header::IntoHeaderValue<UpdatePostRequest> {
    type Error = String;

    fn try_from(hdr_value: HeaderValue) -> std::result::Result<Self, Self::Error> {
        match hdr_value.to_str() {
            std::result::Result::Ok(value) => {
                match <UpdatePostRequest as std::str::FromStr>::from_str(value) {
                    std::result::Result::Ok(value) => {
                        std::result::Result::Ok(header::IntoHeaderValue(value))
                    }
                    std::result::Result::Err(err) => std::result::Result::Err(format!(
                        "Unable to convert header value '{}' into UpdatePostRequest - {}",
                        value, err
                    )),
                }
            }
            std::result::Result::Err(e) => std::result::Result::Err(format!(
                "Unable to convert header: {:?} to string: {}",
                hdr_value, e
            )),
        }
    }
}

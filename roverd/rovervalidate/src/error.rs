use serde_derive::Serialize;
use std::fmt;

/**
 * Custom errors that make understanding all parts of the validation a breeze. They come in a nice printed format already (for CLI usage) but are also
 * serializable in a structured format to return them in the API.
 */

#[derive(Debug, PartialEq, Serialize, Clone)]
pub enum Error {
    // Isolated errors for individual files
    ParseError(FieldError),
    FieldValidationError(FieldError),
    // Errors for the state of the pipeline
    PipelineValidationError(PipelineValidationError),
    // Generic errors
    Io(String),
}

// An error in a specific field of an object (to be ued for JSON, YAML, struct validation)
// The error is recursive, to have the nested field names available in the error message (the end users will like that)
#[derive(Debug, PartialEq, Serialize, Clone)]
pub struct FieldError {
    pub path: Vec<String>, // The path to the field that failed validation
    pub message: String,   // The error message
}

#[derive(Debug, PartialEq, Serialize, Clone)]
pub enum PipelineValidationError {
    UnmetDependencyError(UnmetDependencyError),
    DuplicateServiceError(String),
    DuplicateAliasError(String),
    AliasInUseAsNameError(String),
}

#[derive(Debug, PartialEq, Serialize, Clone)]
pub enum UnmetDependencyError {
    UnmetStream(UnmetStreamError),
    UnmetService(UnmetServiceError),
}

#[derive(Debug, PartialEq, Serialize, Clone)]
pub struct UnmetStreamError {
    pub source: String, // the name of the service that has the unmet dependency
    pub target: String, // the name of the service that should have this dependency (it might or might not exist)
    pub stream: String, // the name of the stream that is missing
}

#[derive(Debug, PartialEq, Serialize, Clone)]
pub struct UnmetServiceError {
    pub source: String, // the name of the service that has the unmet dependency
    pub target: String, // the name of the service that should have this dependency but is missing
}

// Support default errors
impl From<std::io::Error> for Error {
    fn from(e: std::io::Error) -> Self {
        Error::Io(e.to_string())
    }
}

// Add user-friendly messages
impl fmt::Display for Error {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            Error::ParseError(fe) => {
                write!(f, "Parsing failed: {}", fe)
            }
            Error::FieldValidationError(fe) => {
                write!(f, "Field validation failed: {}", fe)
            }
            Error::PipelineValidationError(e) => {
                write!(f, "Pipeline validation failed: {}", e)
            }
            Error::Io(e) => {
                write!(f, "{}", e)
            }
        }
    }
}

// User-friendly messages
impl fmt::Display for FieldError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        // Iterate over the path and print it as a nested structure (indented)
        for (index, field) in self.path.iter().enumerate() {
            // Indent depends on the index
            write!(f, "\n{:indent$}{}", "", field, indent = index * 2)?;
        }

        // Print the error message with indent
        write!(
            f,
            "\n{:indent$}-> {}",
            "",
            self.message,
            indent = self.path.len() * 2
        )
    }
}
impl fmt::Display for PipelineValidationError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            PipelineValidationError::UnmetDependencyError(e) => {
                write!(f, "Unmet dependency error: {}", e)
            }
            PipelineValidationError::DuplicateServiceError(name) => {
                write!(f, "Service '{}' is defined more than once in this pipeline. Remove duplicate services.", name)
            }
            PipelineValidationError::DuplicateAliasError(duplicate_alias) => {
                write!(f, "Found duplicate alias '{}'", duplicate_alias)
            }
            PipelineValidationError::AliasInUseAsNameError(alias) => {
                write!(f, "Alias '{}' is used as a name by another service", alias)
            }
        }
    }
}
impl fmt::Display for UnmetStreamError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(
            f,
            "service '{}' depends on stream '{}' from service '{}', which is missing. Make sure to expose this stream as output in the service.yaml configuration for service '{}'.",
            self.source, self.stream, self.target, self.target
        )
    }
}
impl fmt::Display for UnmetServiceError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(
            f,
            "service '{}' depends on service '{}', which is missing. Make sure to enable the service '{}' in this pipeline.",
            self.source, self.target, self.target
        )
    }
}
impl fmt::Display for UnmetDependencyError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            UnmetDependencyError::UnmetStream(e) => e.fmt(f),
            UnmetDependencyError::UnmetService(e) => e.fmt(f),
        }
    }
}

// Compatibility with `std::error::Error`
impl std::error::Error for Error {}

// Multiple results can be returned
pub type Result<T> = std::result::Result<T, Vec<Error>>;

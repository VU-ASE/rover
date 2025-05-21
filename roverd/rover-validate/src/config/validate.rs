use super::gen;

use crate::error::{Error, Result};
use crate::validate::Validate;
use regex::Regex;

// This enforces the type-state pattern, useful for ensuring only accepting valid configurations
#[repr(transparent)]
#[derive(Debug, Clone)]
pub struct ValidatedConfiguration(pub gen::Configuration);

/**
 * This module applies validation logic to the auto-generated types for the roverd config. We cannot use the validated crate, because if we regenerate the types, we lose the validation logic.
 * But the validate_field macro gives nicer errors anyway, so it's not a big deal.
 */
impl Validate<ValidatedConfiguration> for gen::Configuration {
    // Validate all properties of the configuration in isolation
    fn validate(&self) -> Result<ValidatedConfiguration> {
        let mut errors = Vec::new();

        let pattern = Regex::new(r"^\/.*$").unwrap();
        // Validate all enabled services, iterate and use the index for every error message
        for (index, enabled) in self.enabled.iter().enumerate() {
            if enabled.is_empty() {
                errors.push(Error::FieldValidationError(crate::error::FieldError {
                    path: vec!["enabled".to_string(), index.to_string()],
                    message: "must not be empty".to_string(),
                }));
            }

            if !pattern.is_match(enabled) {
                errors.push(Error::FieldValidationError(crate::error::FieldError {
                    path: vec!["enabled".to_string(), index.to_string()],
                    message: "must be a valid path that starts with a slash".to_string(),
                }));
            }
        }

        // Check that there are no duplicate services in the enabled services
        let mut enabled_services = Vec::new();
        for (index, enabled) in self.enabled.iter().enumerate() {
            if enabled_services.contains(&enabled) {
                let msg = format!(
                    "enabled services must be unique, but {} exists more than once",
                    enabled
                );

                errors.push(Error::FieldValidationError(crate::error::FieldError {
                    path: vec!["enabled".to_string(), index.to_string()],
                    message: msg,
                }));
            } else {
                enabled_services.push(enabled);
            }
        }

        if errors.is_empty() {
            Ok(ValidatedConfiguration(self.clone()))
        } else {
            Err(errors)
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const TEST_FILES_LOCATION: &str = "./src/testfiles/roverd-yaml";

    #[test]
    fn test_valid_files() {
        let valid_path = format!("{}/valid", TEST_FILES_LOCATION);

        // Get all files in this directory
        let files = std::fs::read_dir(valid_path).unwrap();

        // Iterate over all files and validate them
        for file in files {
            let file = file.unwrap();
            let file_path = file.path();
            let file_name = file.file_name().into_string().unwrap();

            // Skip directories
            if file_path.is_dir() {
                continue;
            }

            println!("Validating file: {}", file_name);

            let file_content = std::fs::read_to_string(file_path).unwrap();
            let service: gen::Configuration = serde_yaml::from_str(&file_content).unwrap();

            // Print errors
            if let Err(errors) = service.validate() {
                for error in errors {
                    print!("{}\n", error);
                }
            }

            assert!(
                service.validate().is_ok(),
                "Validation failed for file: {}",
                file_name
            );
        }
    }

    #[test]
    fn test_invalid_files() {
        let invalid_path = format!("{}/invalid", TEST_FILES_LOCATION);

        // Get all files in this directory
        let files = std::fs::read_dir(invalid_path).unwrap();

        // Iterate over all files and validate them
        for file in files {
            let file = file.unwrap();
            let file_path = file.path();
            let file_name = file.file_name().into_string().unwrap();

            // Skip directories
            if file_path.is_dir() {
                continue;
            }

            println!("Validating file: {}", file_name);

            let file_content = std::fs::read_to_string(file_path).unwrap();
            let service: gen::Configuration = serde_yaml::from_str(&file_content).unwrap();

            assert!(
                service.validate().is_err(),
                "Validation should have failed for file: {}",
                file_name
            );
        }
    }
}

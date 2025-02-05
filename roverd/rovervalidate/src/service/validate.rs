use super::gen;

use crate::error::{Error, Result};
use crate::{validate::Validate, validate_field};
use regex::Regex;

// This enforces the type-state pattern, useful for composing a Pipeline later with strict requirements
#[repr(transparent)]
#[derive(Debug, Clone)]
pub struct ValidatedService(pub gen::Service);

/**
 * This module applies validation logic to the auto-generated types for a rover service. We cannot use the validated crate, because if we regenerate the types, we lose the validation logic.
 * But the validate_field macro gives nicer errors anyway, so it's not a big deal.
 */
impl Validate<ValidatedService> for gen::Service {
    fn validate(&self) -> Result<ValidatedService> {
        let mut errors = Vec::new();

        validate_field!(self.name, &mut errors, |name| {
            if name.is_empty() {
                return Some("must not be empty");
            }

            let pattern = Regex::new(r"^[a-z]+(-[a-z]+)*$").unwrap();
            if !pattern.is_match(name) {
                return Some("can only consist of lowercase letters and hyphens");
            }

            None
        });

        validate_field!(self.service_as, &mut errors, |alias| {
            if let Some(alias) = alias {
                if alias.is_empty() {
                    return Some("must not be empty");
                }
                let pattern = Regex::new(r"^[a-z]+(-[a-z]+)*$").unwrap();
                if !pattern.is_match(alias) {
                    return Some("can only consist of lowercase letters and hyphens");
                }
            }
            None
        });

        validate_field!(self.author, &mut errors, |author| {
            if author.is_empty() {
                return Some("must not be empty");
            }

            let pattern = Regex::new(r"^[a-zA-Z0-9]+(-[a-zA-Z0-9]+)*$").unwrap();
            if !pattern.is_match(author) {
                return Some("can only consist of alphanumeric characters and hyphens");
            }

            None
        });

        validate_field!(self.version, &mut errors, |version| {
            if version.is_empty() {
                return Some("must not be empty");
            }

            let pattern = Regex::new(r"^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-([0-9A-Za-z-]+(?:\.[0-9A-Za-z-]+)*))?(?:\+([0-9A-Za-z-]+(?:\.[0-9A-Za-z-]+)*))?$").unwrap();
            if !pattern.is_match(version) {
                return Some("must be a valid semantic version");
            }

            None
        });

        // Validate all commands
        if let Err(command_errors) = self.commands.validate() {
            for error in command_errors {
                match error {
                    Error::FieldValidationError(mut field_error) => {
                        field_error.path.insert(0, "commands".to_string());
                        errors.push(Error::FieldValidationError(field_error));
                    }
                    Error::ParseError(mut field_error) => {
                        field_error.path.insert(0, "commands".to_string());
                        errors.push(Error::ParseError(field_error));
                    }
                    _ => (),
                }
            }
        }

        // Validate all inputs
        for (index, input) in self.inputs.iter().enumerate() {
            if let Err(input_errors) = input.validate() {
                for error in input_errors {
                    match error {
                        Error::FieldValidationError(mut field_error) => {
                            field_error.path.insert(0, "inputs".to_string());
                            field_error.path.insert(1, index.to_string());
                            errors.push(Error::FieldValidationError(field_error));
                        }
                        Error::ParseError(mut field_error) => {
                            field_error.path.insert(0, "inputs".to_string());
                            field_error.path.insert(1, index.to_string());
                            errors.push(Error::ParseError(field_error));
                        }
                        _ => (),
                    }
                }
            }
        }

        // Validate all configurations
        for (index, configuration) in self.configuration.iter().enumerate() {
            if let Err(configuration_errors) = configuration.validate() {
                for error in configuration_errors {
                    match error {
                        Error::FieldValidationError(mut field_error) => {
                            field_error.path.insert(0, "configurations".to_string());
                            field_error.path.insert(1, index.to_string());
                            errors.push(Error::FieldValidationError(field_error));
                        }
                        Error::ParseError(mut field_error) => {
                            field_error.path.insert(0, "configurations".to_string());
                            field_error.path.insert(1, index.to_string());
                            errors.push(Error::ParseError(field_error));
                        }
                        _ => (),
                    }
                }
            }
        }

        let pattern = Regex::new(r"^[a-z]+(-[a-z]+)*$").unwrap();
        // Validate all outputs
        for (index, output) in self.outputs.iter().enumerate() {
            if !pattern.is_match(output) {
                errors.push(Error::FieldValidationError(crate::error::FieldError {
                    path: vec!["outputs".to_string(), index.to_string()],
                    message: "can only consist of lowercase letters and hyphens".to_string(),
                }));
            }
        }

        // Make sure there are no duplicate input services
        let mut input_services = Vec::new();
        for (index, input) in self.inputs.iter().enumerate() {
            if input_services.contains(&input.service) {
                let msg = format!(
                    "input services must be unique, but {} exists more than once",
                    input.service
                );
                errors.push(Error::FieldValidationError(crate::error::FieldError {
                    path: vec!["inputs".to_string(), index.to_string()],
                    message: msg,
                }));
            } else {
                input_services.push(input.service.clone());
            }
        }

        // Make sure there are no duplicate output services
        let mut output_services = Vec::new();
        for (index, output) in self.outputs.iter().enumerate() {
            if output_services.contains(&output) {
                let msg = format!(
                    "output services must be unique, but {} exists more than once",
                    output
                );
                errors.push(Error::FieldValidationError(crate::error::FieldError {
                    path: vec!["outputs".to_string(), index.to_string()],
                    message: msg,
                }));
            } else {
                output_services.push(output);
            }
        }

        if errors.is_empty() {
            Ok(ValidatedService(self.clone()))
        } else {
            Err(errors)
        }
    }
}

impl Validate<bool> for gen::Commands {
    fn validate(&self) -> Result<bool> {
        let mut errors = Vec::new();

        validate_field!(self.run, &mut errors, |run| {
            if run.is_empty() {
                Some("must not be empty")
            } else {
                None
            }
        });

        if errors.is_empty() {
            Ok(true)
        } else {
            Err(errors)
        }
    }
}

impl Validate<bool> for gen::Input {
    fn validate(&self) -> Result<bool> {
        let mut errors = Vec::new();

        validate_field!(self.service, &mut errors, |service| {
            if service.is_empty() {
                return Some("must not be empty");
            }

            let pattern = Regex::new(r"^[a-z]+(-[a-z]+)*$").unwrap();
            if !pattern.is_match(service) {
                return Some("can only consist of lowercase letters and hyphens");
            }

            None
        });

        let mut input_streams = Vec::new(); // Make sure there are no duplicate input streams
        let pattern = Regex::new(r"^[a-z]+(-[a-z]+)*$").unwrap();
        for (index, stream) in self.streams.iter().enumerate() {
            if !pattern.is_match(stream) {
                errors.push(Error::FieldValidationError(crate::error::FieldError {
                    path: vec!["streams".to_string(), index.to_string()],
                    message: "can only consist of lowercase letters and hyphens".to_string(),
                }));
            }

            if input_streams.contains(&stream) {
                let msg = format!(
                    "input streams must be unique, but {} exists more than once",
                    stream
                );
                errors.push(Error::FieldValidationError(crate::error::FieldError {
                    path: vec!["streams".to_string(), index.to_string()],
                    message: msg,
                }));
            } else {
                input_streams.push(stream);
            }
        }

        if errors.is_empty() {
            Ok(true)
        } else {
            Err(errors)
        }
    }
}

impl Validate<bool> for gen::Configuration {
    fn validate(&self) -> Result<bool> {
        let mut errors = Vec::new();

        validate_field!(self.name, &mut errors, |name| {
            if name.is_empty() {
                return Some("must not be empty");
            }

            let pattern = Regex::new(r"^[a-z]+(-[a-z]+)*$").unwrap();
            if !pattern.is_match(name) {
                return Some("can only consist of lowercase letters and hyphens");
            }

            None
        });

        validate_field!(self.value, &mut errors, |value| {
            // Checks depend on the type of the value
            match value {
                gen::Value::String(s) => {
                    if s.is_empty() {
                        return Some("must not be empty");
                    }
                    // Does this correspond with self.configuration_type?
                    if let Some(gen::Type::Number) = &self.configuration_type {
                        return Some("is parsed as a string, but was specified as a float, make sure to use float syntax");
                    }
                }
                gen::Value::Double(f) => {
                    if f.is_nan() {
                        return Some("must be a valid float");
                    }
                    // Does this correspond with self.configuration_type?
                    if let Some(gen::Type::String) = self.configuration_type {
                        return Some("is parsed as a float, but was specified as a string, make sure to encapsulate your string in quotes");
                    }
                }
            }
            None
        });

        if errors.is_empty() {
            Ok(true)
        } else {
            Err(errors)
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const TEST_FILES_LOCATION: &str = "./src/testfiles/service-yaml";

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
            let service: gen::Service = serde_yaml::from_str(&file_content).unwrap();

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
            let service: gen::Service = serde_yaml::from_str(&file_content).unwrap();

            assert!(
                service.validate().is_err(),
                "Validation should have failed for file: {}",
                file_name
            );
        }
    }
}

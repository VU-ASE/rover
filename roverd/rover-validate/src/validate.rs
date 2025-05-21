use crate::error::{Error, Result};

pub fn validate_helper<T, F>(
    field_value: &T,
    field_name: &str,
    errors: &mut Vec<Error>,
    validator: F,
) where
    F: FnOnce(&T) -> Option<&'static str>,
{
    if let Some(msg) = validator(field_value) {
        errors.push(Error::FieldValidationError(crate::error::FieldError {
            path: vec![field_name.to_string()],
            message: msg.to_string(),
        }));
    }
}

// Macro definition to validate a field by its name and automatically add the name to error messages
#[macro_export]
macro_rules! validate_field {
    ($root:ident . $($field:ident).+ , $errors:expr, $validator:expr) => {{
        use $crate::validate::validate_helper;
        let field_name = stringify!($($field).+);
        validate_helper(&$root.$($field).+, field_name, $errors, $validator);
    }};
}

pub trait Validate<T> {
    fn validate(&self) -> Result<T>;
}

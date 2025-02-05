mod gen;
mod validate;

/**
 * This module exposes validation and parsing logic for the roverd configuration file (roverd.yaml)
 */
pub use crate::validate::Validate;
pub use gen::*;
pub use validate::*;

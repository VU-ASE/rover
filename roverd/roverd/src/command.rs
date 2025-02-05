use crate::error;

/// The service defines run and build commands which are encapsulated into this
/// convenience struct.
#[derive(Debug)]
pub struct ParsedCommand {
    pub program: String,
    pub arguments: Vec<String>,
}

/// Commands are read from the yaml file directly and need to be parsed into
/// the program name and vector of arguments.
impl TryFrom<&String> for ParsedCommand {
    type Error = error::Error;

    fn try_from(value: &String) -> Result<Self, error::Error> {
        let command = value.clone();
        let separated: Vec<&str> = command.split_whitespace().collect();
        let (program, arguments) = separated
            .split_first()
            .ok_or_else(|| error::Error::ParsingRunCommand)?;
        let program = program.to_string();
        let arguments: Vec<String> = arguments.iter().map(|arg| arg.to_string()).collect();

        Ok(Self { program, arguments })
    }
}

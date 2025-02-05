use tracing::Level;

/// Initializes logging to stdout, by default Level::Trace shows all logs,
/// otherwise change to Level::WARN for just warn! & error!
pub fn init() {
    tracing_subscriber::fmt()
        .with_max_level(Level::INFO)
        .without_time()
        .with_line_number(false)
        .with_file(false)
        .init();
}

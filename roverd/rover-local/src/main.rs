use rover_bootspec::BootSpecs;
use rover_types::error::Error;
use rover_validate::{
    pipeline::interface::Pipeline,
    service::{Service, ValidatedService},
    validate::Validate,
};

fn main() -> Result<(), Error> {
    // Todo: get these from the command line
    let service_paths = [
        "./example-pipelines/imaging.yaml",
        "./example-pipelines/controller.yaml",
        "./example-pipelines/actuator.yaml",
    ];

    let mut enabled_services: Vec<ValidatedService> = vec![];

    for service_path in &service_paths {
        let service_file = std::fs::read_to_string(service_path).map_err(|_| {
            Error::ServiceNotFound(format!("could not find or read {}", service_path))
        })?;
        let service: Service = serde_yaml::from_str(&service_file)?;
        let validated = service.validate()?;
        enabled_services.push(validated);
    }

    let runnable_pipeline = Pipeline::new(enabled_services).validate()?;

    let bootspecs_map = BootSpecs::new(runnable_pipeline);

    println!("Todo: ");
    dbg!(bootspecs_map);

    Ok(())
}

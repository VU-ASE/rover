use std::collections::HashMap;

use rover_constants::*;

use rover_validate::pipeline::interface::RunnablePipeline;
use serde::{Deserialize, Serialize};

use rover_types::service::FqBuf;

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct Stream {
    pub name: String,
    pub address: String,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Input {
    pub service: String,
    pub streams: Vec<Stream>,
}

#[derive(Debug, Serialize, Deserialize)]
pub enum BootSpecDataType {
    String(String),
    Number(f64),
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct BootSpecTuning {
    pub enabled: bool,
    pub address: String,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct BootSpec {
    pub name: String,
    pub version: String,
    pub inputs: Vec<Input>,
    pub outputs: Vec<Stream>,
    pub configuration: Vec<rover_validate::service::Configuration>,
    pub tuning: BootSpecTuning,
}

#[repr(transparent)]
#[derive(Debug)]
pub struct BootSpecs(pub HashMap<FqBuf, BootSpec>);

impl BootSpecs {
    pub fn new(pipeline: RunnablePipeline) -> Self {
        let mut services = pipeline.0.services;

        // Transceiver outputs to START_PORT
        let mut tuning = BootSpecTuning {
            enabled: false,
            address: format!("{}:{}", DATA_ADDRESS, START_PORT),
        };

        let transeiver_service = (0..services.len()).find_map(|i| {
            if services[i].0.get_original_name() == "transceiver" {
                tuning.enabled = true;
                Some(services.swap_remove(i))
            } else {
                None
            }
        });

        let mut transceiver_inputs = vec![];

        let mut start_port = START_PORT + 1;

        let mut result = HashMap::new();

        // Create a mapping for all outputs, such that we can lookup a (service, stream)
        // and get the assigned address.
        let mut mappings: HashMap<(String, String), String> = HashMap::new();

        for validated in &services {
            let s = &validated.0;
            // For each service assign an address to all of its outputs and
            // save the resulting address in the mapping.
            for out_stream in &s.outputs {
                let address = format!("{}:{}", DATA_ADDRESS, start_port);
                let service_name = s.get_pipeline_name();
                mappings.insert((service_name, out_stream.clone()), address.clone());
                start_port += 1;
            }
        }

        // Now that we know the mappings we can iterate over all services again and set
        // each output and input field by looking up the address from the previous step
        for validated in &services {
            let s = &validated.0;
            let service_name = s.get_pipeline_name();

            let fq = FqBuf::from(validated);

            let mut outputs = vec![];
            for out_stream in s.outputs.iter() {
                let stream_name = out_stream.clone();

                if let Some(address) = mappings.get(&(service_name.clone(), stream_name.clone())) {
                    // For outputs, the address should be in the form of tcp://*:port instead of tcp://localhost:port
                    // (required for zmq bind). So we replace localhost with *.
                    let bind_address = address.clone().replace("localhost", "*");

                    outputs.push(Stream {
                        name: stream_name.clone(),
                        address: bind_address,
                    });
                }
            }

            // If we have a transceiver, it gets all outputs as its inputs
            if transeiver_service.is_some() {
                transceiver_inputs.push(Input {
                    service: service_name.clone(),
                    streams: outputs
                        .clone()
                        .iter()
                        .map(|s| {
                            Stream {
                                name: s.name.clone(),
                                // Conversely, for inputs, the address should be in the form of tcp://localhost:port instead of tcp://*:port
                                // (required for zmq connect). So we replace * with localhost.
                                address: s.address.clone().replace("*", "localhost"),
                            }
                        })
                        .collect(),
                });
            }

            let mut inputs = vec![];
            for input_stream in s.inputs.iter() {
                let service_name = &input_stream.service;

                let mut streams = vec![];

                for stream_name in input_stream.streams.iter() {
                    if let Some(address) =
                        mappings.get(&(service_name.clone(), stream_name.clone()))
                    {
                        streams.push(Stream {
                            name: stream_name.clone(),
                            address: address.clone(),
                        });
                    }
                }

                inputs.push(Input {
                    service: service_name.clone(),
                    streams,
                });
            }

            let b = BootSpec {
                name: s.get_pipeline_name(),
                version: s.version.clone(),
                inputs,
                outputs,
                configuration: s.configuration.clone(),
                tuning: tuning.clone(),
            };

            result.insert(fq, b);
        }

        // Add the battery stream as input to transceiver
        if transeiver_service.is_some() {
            transceiver_inputs.push(Input {
                service: "battery".to_string(),
                streams: vec![Stream {
                    name: BATTERY_STREAM_NAME.to_string(),
                    address: format!("{}:{}", DATA_ADDRESS, BATTERY_PORT),
                }],
            })
        }

        if let Some(s) = transeiver_service {
            let fq = FqBuf::from(s.clone());

            let service_name = s.0.get_pipeline_name();

            let b = BootSpec {
                name: service_name.clone(),
                version: s.0.version.clone(),
                inputs: transceiver_inputs,
                outputs: vec![Stream {
                    name: service_name,
                    // For outputs, the address should be in the form of tcp://*:port instead of tcp://localhost:port
                    address: tuning.address.clone().replace("localhost", "*"),
                }],
                configuration: s.0.configuration.clone(),
                // This might seem weird, but the transceiver itself does not listen to tuning from another service
                tuning: BootSpecTuning {
                    enabled: false,
                    address: "disabled".to_string(),
                },
            };

            result.insert(fq, b);
        }

        Self(result)
    }
}

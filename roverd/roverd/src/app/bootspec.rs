use std::collections::HashMap;

use crate::constants::*;

use rovervalidate::service::ValidatedService;
use serde::{Deserialize, Serialize};

use super::service::FqBuf;

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
    pub configuration: Vec<rovervalidate::service::Configuration>,
    pub tuning: BootSpecTuning,
}

#[repr(transparent)]
pub struct BootSpecs(pub HashMap<FqBuf, BootSpec>);

impl BootSpecs {
    pub fn new(mut services: Vec<ValidatedService>) -> Self {
        // Transceiver outputs to START_PORT
        let mut tuning = BootSpecTuning {
            enabled: false,
            address: format!("{}:{}", DATA_ADDRESS, START_PORT),
        };

        let transeiver_service = (0..services.len()).find_map(|i| {
            if services[i].0.name == "transceiver" {
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
                let stream_name = out_stream.clone();
                mappings.insert((s.name.clone(), stream_name.clone()), address.clone());
                start_port += 1;
            }
        }

        // Now that we know the mappings we can iterate over all service again
        // and set each output and input field
        for validated in &services {
            let s = &validated.0;
            let service_name = &s.name;
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

                // If we have a transceiver, it gets all inputs
                if transeiver_service.is_some() {
                    transceiver_inputs.push(Input {
                        service: service_name.clone(),
                        streams: streams.clone(),
                    });
                }

                inputs.push(Input {
                    service: service_name.clone(),
                    streams,
                });
            }

            let b = BootSpec {
                name: s.name.clone(),
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
            let b = BootSpec {
                name: s.0.name.clone(),
                version: s.0.version.clone(),
                inputs: transceiver_inputs,
                outputs: vec![Stream {
                    name: s.0.name.clone(),
                    address: tuning.address.clone(),
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

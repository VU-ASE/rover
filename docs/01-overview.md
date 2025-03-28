# Overview

The [rover repository](https://github.com/VU-ASE/rover) houses the software framework that makes the Rover runnable and controllable. Most importantly, it exposes:

- `roverd`: the daemon that runs on the Rover and listens to commands on an exposed HTTP REST API
- `roverctl`: the CLI that runs on your local device and interfaces with the `roverd` API to control your Rover

Additional tools are available for debugging and tuning your Rover.

![Schematic of roverd and roverctl](https://github.com/user-attachments/assets/8b9e1c8b-192e-48ba-9dae-d300caf61290)

## Versioning and compatibility

Both `roverd` and `roverctl` contain partly auto-generated code based on several [specs](https://ase.vu.nl/docs/framework/glossary/spec), expressed as OpenAPI specs and JSON schemas (you can find them [here](https://github.com/VU-ASE/rover/tree/main/spec)). Because all tools are incorporated into one repository, all software is recompiled on every release. This means that different software under the same release is **guaranteed to work**. I.e. `roverctl` version 1.8.1 is one-to-one compatible with `roverd` version 1.8.1. 

Other versions _may_ work but backwards compatibility is not a priority currently. `roverctl` will report about version mismatches.
# Overview

This piece of software is the always-running http endpoint that comes with all rovers. It runs in the background and does not need to be explicitly enable or disable by users, since it is **always** enabled. The main functionality of this program is summarized as follows:

* The [roverctl](https://ase.vu.nl/docs/category/roverctl) utility interacts with the API exposed by roverd
* Roverd is in charge of starting and stopping any user-defined programs
* Roverd has numerous safety measures in place like making sure the Debix shuts down when the battery is low


Since this roverd runs in the background, it is referred to as a daemon (hence the "d" in roverd) and it works with the following two concepts: **services** and a **pipeline**. Services can be though of as any program that might run on the rover and a pipeline is a collection of those services that share information during runtime. All the services in a pipeline get started and stopped together and if one service crashes all other services in the pipeline are terminated. The definition of a pipeline is simply a list of *enabled* services. In the case of roverd, it only saves pipelines that are valid, meaning that all dependencies of all services are met. Since it will reject invalid pipelines, it is easier to reason about the state since we know that at any given moment the stored pipeline (in `/etc/roverd/rover.yaml`) is always a valid one.

The following shows the three states of a pipeline: `Empty`, `Startable` and `Started`. From the `Empty` state one can set a pipeline. If that pipeline is invalid, it will be rejected an we remain in `Empty`. On the other hand, if it is valid, then we transition to the `Startable` state from where we can start the rover. From this state any changes made to the pipeline will be checked again so if a new pipeline is invalid, it will be sent back to the `Empty` state.

![StateMachine](https://github.com/user-attachments/assets/36534655-1904-40ce-b170-e1b6fb5e0cc7)

After starting the rover from the `Startable` state, the pipeline moves to the `Started` state. From there, if any process from a service exits, all other processes will be terminated and we are back in the `Startable` state. The stop command will terminate all processes and bring us back to the `Startable` state.

The file system (with the `/etc/roverd/rover.yaml`) holds the source of truth in this case, so no runtime state is stored in memory. All actions performed by roverd check the filesystem first in case any changes have been made on disk.

## API
Roverd is an always running process on the rover (daemon) which exposes endpoints that allow programs like `roverctl` to interact with the rover. This repo also defines the API specification which clients need to implement in order to use the provided functionality ([apispec.yaml](https://github.com/VU-ASE/roverd/blob/main/roverd/spec/apispec.yaml)). In short, roverd lets you view system status, upload services and start/stop a pipeline.

Roverd is an always running process on the rover (daemon) which exposes endpoints that allow programs like `roverctl` or `web-monitor` to interact with the rover. This repo also defines the API specification which clients need to implement in order to use the provided functionality ([apispec.yaml](https://github.com/VU-ASE/roverd/blob/main/roverd/spec/apispec.yaml)). In short, roverd lets you view system status, upload services and start/stop a collection of services (a pipeline).



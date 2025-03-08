
![rover-repo-overview](https://github.com/user-attachments/assets/12c267ba-9896-40fa-876f-07dc92a0736f)

<hr/>

**The `rover` repository is the heart and soul of everything ASE. The various software components bundled here allow for pipeline execution, service development and remote control of the Rovers we offer. All software is always bundled and released in parallel, so you can always rest assured that versions are perfectly compatible.**

## Software

- `roverd`: the central service execution daemon, running on the Rover. This daemon exposes a REST API for service scheduling and execution, as defined in [the OpenAPI spec](/spec)
- `roverctl`: convenient CLI to communicate with the `roverd` REST API that integrates a proxy server for debugging between the Rover and the browser using webRTC
- `roverctl-web`: a web interface to get finegrained control over rover execution, an extension of `roverctl`

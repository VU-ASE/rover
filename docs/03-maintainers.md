# For Maintainers

The following page describes how to get started developing with roverd and holds all information future maintainers should know. If you are a student this page is likely completely irrelevant.

## Development

All dependencies are bundled in the devcontainer, as well as a debix user and filesystem setup identical to that of a rover. Run `make dev` for development. If changes are made to [`apispec.yaml`](https://github.com/VU-ASE/roverd/blob/main/roverd/spec/apispec.yaml), then the openapi definitions must be generated again with `cd roverd ; make build`.

> Important: due to bugs in the openapi generator, some tuple structs have private members, which needs to be updated manually after re-generating the openapi definitions. After running `make build` the following need to be edited by hand after which everything should compile.

[`roverd/openapi/src/models.rs`](https://github.com/VU-ASE/roverd/blob/main/roverd/openapi/src/models.rs)
```
                                 +++
pub struct DuplicateServiceError(pub String);

                                                                             +++
pub struct ServicesAuthorServiceVersionGet200ResponseConfigurationInnerValue(pub Box<serde_json::value::RawValue> );
```

For interacting with the API, the Swagger extension (already installed through devcontainer) is extremely helpful. It lets you test authorized API requests based on the specification.

## CI/CD

The Github actions runner can be run on any architecture since cross-compilation is supported. It builds the roverd binary for the Debix Model A by first building the same docker image used for the devcontainer and then running the `make build-arm` target inside the container.

## Directories
* `/roverd` - source code for roverd
* `/roverd/spec` - openapi and bootspec specifications
* `/roverd/example-pipelines` - dummy services that can be used for testing
* `/roverd/openapi` - generated rust code from openapi
* `/rovervalidate` - library that performs validation of service and configuration files
* `/scripts` - useful for testing


## Future Improvements
This repo has an unnecessarily **large** amount of code due to type conversions between types generated from openapi and validation types in rovervalidate. Furthermore, boilerplate code could be largely reduced by not generating Rust code from openapi, but by generating a openapi defintion form Rust types. Furthermore, there is currently a mix of `anyhow` Errors as well as typed errors which might not be the best design, since the API should always report as much information to the logs as possible.

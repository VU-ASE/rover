# \PipelineAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LogsAuthorNameVersionGet**](PipelineAPI.md#LogsAuthorNameVersionGet) | **Get** /logs/{author}/{name}/{version} | Retrieve logs for any service. Logs from running or previously run services can be viewed and will be kept until rover reboot.
[**PipelineGet**](PipelineAPI.md#PipelineGet) | **Get** /pipeline | Retrieve pipeline status and process execution information
[**PipelinePost**](PipelineAPI.md#PipelinePost) | **Post** /pipeline | Set the services that are enabled in this pipeline, by specifying the fully qualified services
[**PipelineStartPost**](PipelineAPI.md#PipelineStartPost) | **Post** /pipeline/start | Start the pipeline
[**PipelineStopPost**](PipelineAPI.md#PipelineStopPost) | **Post** /pipeline/stop | Stop the pipeline



## LogsAuthorNameVersionGet

> []string LogsAuthorNameVersionGet(ctx, author, name, version).Lines(lines).Execute()

Retrieve logs for any service. Logs from running or previously run services can be viewed and will be kept until rover reboot.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	author := "vu-ase" // string | The author of the service.
	name := "imaging" // string | The name of the service.
	version := "1.0.0" // string | The version of the service.
	lines := int32(100) // int32 | The number of log lines to retrieve (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelineAPI.LogsAuthorNameVersionGet(context.Background(), author, name, version).Lines(lines).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelineAPI.LogsAuthorNameVersionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LogsAuthorNameVersionGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `PipelineAPI.LogsAuthorNameVersionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**author** | **string** | The author of the service. | 
**name** | **string** | The name of the service. | 
**version** | **string** | The version of the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogsAuthorNameVersionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lines** | **int32** | The number of log lines to retrieve | [default to 50]

### Return type

**[]string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PipelineGet

> PipelineGet200Response PipelineGet(ctx).Execute()

Retrieve pipeline status and process execution information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelineAPI.PipelineGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelineAPI.PipelineGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PipelineGet`: PipelineGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PipelineAPI.PipelineGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPipelineGetRequest struct via the builder pattern


### Return type

[**PipelineGet200Response**](PipelineGet200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PipelinePost

> PipelinePost(ctx).PipelinePostRequestInner(pipelinePostRequestInner).Execute()

Set the services that are enabled in this pipeline, by specifying the fully qualified services

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pipelinePostRequestInner := []openapiclient.PipelinePostRequestInner{*openapiclient.NewPipelinePostRequestInner("imaging", "1.0.0", "vu-ase")} // []PipelinePostRequestInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PipelineAPI.PipelinePost(context.Background()).PipelinePostRequestInner(pipelinePostRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelineAPI.PipelinePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPipelinePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pipelinePostRequestInner** | [**[]PipelinePostRequestInner**](PipelinePostRequestInner.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PipelineStartPost

> PipelineStartPost(ctx).Execute()

Start the pipeline

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PipelineAPI.PipelineStartPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelineAPI.PipelineStartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPipelineStartPostRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PipelineStopPost

> PipelineStopPost(ctx).Execute()

Stop the pipeline

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PipelineAPI.PipelineStopPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelineAPI.PipelineStopPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPipelineStopPostRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


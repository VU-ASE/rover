# \SourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SourcesDelete**](SourcesAPI.md#SourcesDelete) | **Delete** /sources | Delete a source
[**SourcesGet**](SourcesAPI.md#SourcesGet) | **Get** /sources | Retrieve all sources
[**SourcesPost**](SourcesAPI.md#SourcesPost) | **Post** /sources | Downloads and installs a new source, overwriting the prior version (if any) and adding it to the &#39;downloaded&#39; section in rover.yaml (checks for duplicate source names)



## SourcesDelete

> SourcesDelete(ctx).SourcesPostRequest(sourcesPostRequest).Execute()

Delete a source

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/VU-ASE/roverctl"
)

func main() {
	sourcesPostRequest := *openapiclient.NewSourcesPostRequest("imaging", "github.com/VU-ASE/imaging", "1.0.0") // SourcesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SourcesAPI.SourcesDelete(context.Background()).SourcesPostRequest(sourcesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.SourcesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSourcesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourcesPostRequest** | [**SourcesPostRequest**](SourcesPostRequest.md) |  | 

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


## SourcesGet

> []SourcesGet200ResponseInner SourcesGet(ctx).Execute()

Retrieve all sources

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/VU-ASE/roverctl"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourcesAPI.SourcesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.SourcesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourcesGet`: []SourcesGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SourcesAPI.SourcesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSourcesGetRequest struct via the builder pattern


### Return type

[**[]SourcesGet200ResponseInner**](SourcesGet200ResponseInner.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourcesPost

> SourcesPost(ctx).SourcesPostRequest(sourcesPostRequest).Execute()

Downloads and installs a new source, overwriting the prior version (if any) and adding it to the 'downloaded' section in rover.yaml (checks for duplicate source names)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/VU-ASE/roverctl"
)

func main() {
	sourcesPostRequest := *openapiclient.NewSourcesPostRequest("imaging", "github.com/VU-ASE/imaging", "1.0.0") // SourcesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SourcesAPI.SourcesPost(context.Background()).SourcesPostRequest(sourcesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourcesAPI.SourcesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSourcesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourcesPostRequest** | [**SourcesPostRequest**](SourcesPostRequest.md) |  | 

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


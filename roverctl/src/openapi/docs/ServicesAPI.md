# \ServicesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPost**](ServicesAPI.md#FetchPost) | **Post** /fetch | Fetches the zip file from the given URL and installs the service onto the filesystem
[**FqnsGet**](ServicesAPI.md#FqnsGet) | **Get** /fqns | Retrieve a list of all fully qualified services
[**ServicesAuthorGet**](ServicesAPI.md#ServicesAuthorGet) | **Get** /services/{author} | Retrieve the list of parsable services for a specific author
[**ServicesAuthorServiceGet**](ServicesAPI.md#ServicesAuthorServiceGet) | **Get** /services/{author}/{service} | Retrieve the list of parsable service versions for a specific author and service
[**ServicesAuthorServiceVersionDelete**](ServicesAPI.md#ServicesAuthorServiceVersionDelete) | **Delete** /services/{author}/{service}/{version} | Delete a specific version of a service
[**ServicesAuthorServiceVersionGet**](ServicesAPI.md#ServicesAuthorServiceVersionGet) | **Get** /services/{author}/{service}/{version} | Retrieve the status of a specific version of a service
[**ServicesAuthorServiceVersionPost**](ServicesAPI.md#ServicesAuthorServiceVersionPost) | **Post** /services/{author}/{service}/{version} | Build a fully qualified service version
[**ServicesGet**](ServicesAPI.md#ServicesGet) | **Get** /services | Retrieve the list of all authors that have parsable services. With these authors you can query further for services
[**UploadPost**](ServicesAPI.md#UploadPost) | **Post** /upload | Upload a new service or new version to the rover by uploading a ZIP file



## FetchPost

> FetchPost200Response FetchPost(ctx).FetchPostRequest(fetchPostRequest).Execute()

Fetches the zip file from the given URL and installs the service onto the filesystem

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
	fetchPostRequest := *openapiclient.NewFetchPostRequest("https://downloads.ase.vu.nl/api/imaging/v1.0.0") // FetchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.FetchPost(context.Background()).FetchPostRequest(fetchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.FetchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchPost`: FetchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.FetchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetchPostRequest** | [**FetchPostRequest**](FetchPostRequest.md) |  | 

### Return type

[**FetchPost200Response**](FetchPost200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FqnsGet

> []FqnsGet200ResponseInner FqnsGet(ctx).Execute()

Retrieve a list of all fully qualified services

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
	resp, r, err := apiClient.ServicesAPI.FqnsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.FqnsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FqnsGet`: []FqnsGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.FqnsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFqnsGetRequest struct via the builder pattern


### Return type

[**[]FqnsGet200ResponseInner**](FqnsGet200ResponseInner.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesAuthorGet

> []string ServicesAuthorGet(ctx, author).Execute()

Retrieve the list of parsable services for a specific author

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
	author := "vu-ase" // string | The author name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ServicesAuthorGet(context.Background(), author).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesAuthorGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesAuthorGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesAuthorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**author** | **string** | The author name | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicesAuthorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ServicesAuthorServiceGet

> []string ServicesAuthorServiceGet(ctx, author, service).Execute()

Retrieve the list of parsable service versions for a specific author and service

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
	author := "vu-ase" // string | The author name
	service := "imaging" // string | The service name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ServicesAuthorServiceGet(context.Background(), author, service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesAuthorServiceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesAuthorServiceGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesAuthorServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**author** | **string** | The author name | 
**service** | **string** | The service name | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicesAuthorServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## ServicesAuthorServiceVersionDelete

> ServicesAuthorServiceVersionDelete200Response ServicesAuthorServiceVersionDelete(ctx, author, service, version).Execute()

Delete a specific version of a service

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
	author := "vu-ase" // string | The author name
	service := "imaging" // string | The service name
	version := "1.0.0" // string | The version of the service

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ServicesAuthorServiceVersionDelete(context.Background(), author, service, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesAuthorServiceVersionDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesAuthorServiceVersionDelete`: ServicesAuthorServiceVersionDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesAuthorServiceVersionDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**author** | **string** | The author name | 
**service** | **string** | The service name | 
**version** | **string** | The version of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicesAuthorServiceVersionDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServicesAuthorServiceVersionDelete200Response**](ServicesAuthorServiceVersionDelete200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesAuthorServiceVersionGet

> ServicesAuthorServiceVersionGet200Response ServicesAuthorServiceVersionGet(ctx, author, service, version).Execute()

Retrieve the status of a specific version of a service

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
	author := "vu-ase" // string | The author name
	service := "imaging" // string | The service name
	version := "1.0.0" // string | The version of the service

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ServicesAuthorServiceVersionGet(context.Background(), author, service, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesAuthorServiceVersionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesAuthorServiceVersionGet`: ServicesAuthorServiceVersionGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesAuthorServiceVersionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**author** | **string** | The author name | 
**service** | **string** | The service name | 
**version** | **string** | The version of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicesAuthorServiceVersionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServicesAuthorServiceVersionGet200Response**](ServicesAuthorServiceVersionGet200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesAuthorServiceVersionPost

> ServicesAuthorServiceVersionPost(ctx, author, service, version).Execute()

Build a fully qualified service version

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
	author := "vu-ase" // string | The author name
	service := "imaging" // string | The service name
	version := "1.0.0" // string | The version of the service

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServicesAPI.ServicesAuthorServiceVersionPost(context.Background(), author, service, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesAuthorServiceVersionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**author** | **string** | The author name | 
**service** | **string** | The service name | 
**version** | **string** | The version of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicesAuthorServiceVersionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## ServicesGet

> []string ServicesGet(ctx).Execute()

Retrieve the list of all authors that have parsable services. With these authors you can query further for services

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
	resp, r, err := apiClient.ServicesAPI.ServicesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ServicesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicesGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ServicesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesGetRequest struct via the builder pattern


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


## UploadPost

> FetchPost200Response UploadPost(ctx).Content(content).Execute()

Upload a new service or new version to the rover by uploading a ZIP file

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
	content := os.NewFile(1234, "some_file") // *os.File | The content of the ZIP file to upload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.UploadPost(context.Background()).Content(content).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.UploadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadPost`: FetchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.UploadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **content** | ***os.File** | The content of the ZIP file to upload | 

### Return type

[**FetchPost200Response**](FetchPost200Response.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


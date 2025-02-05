# ServicesPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the service | 
**Version** | **string** | The version of the service | 
**Author** | **string** | The author of the service | 
**InvalidatedPipeline** | **bool** | Whether the pipeline was invalidated by this service upload | 

## Methods

### NewServicesPost200Response

`func NewServicesPost200Response(name string, version string, author string, invalidatedPipeline bool, ) *ServicesPost200Response`

NewServicesPost200Response instantiates a new ServicesPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesPost200ResponseWithDefaults

`func NewServicesPost200ResponseWithDefaults() *ServicesPost200Response`

NewServicesPost200ResponseWithDefaults instantiates a new ServicesPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServicesPost200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicesPost200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicesPost200Response) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *ServicesPost200Response) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServicesPost200Response) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServicesPost200Response) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetAuthor

`func (o *ServicesPost200Response) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ServicesPost200Response) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ServicesPost200Response) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetInvalidatedPipeline

`func (o *ServicesPost200Response) GetInvalidatedPipeline() bool`

GetInvalidatedPipeline returns the InvalidatedPipeline field if non-nil, zero value otherwise.

### GetInvalidatedPipelineOk

`func (o *ServicesPost200Response) GetInvalidatedPipelineOk() (*bool, bool)`

GetInvalidatedPipelineOk returns a tuple with the InvalidatedPipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidatedPipeline

`func (o *ServicesPost200Response) SetInvalidatedPipeline(v bool)`

SetInvalidatedPipeline sets InvalidatedPipeline field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SourcesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the source | 
**Url** | **string** | The URL of the source (without scheme) | 
**Version** | **string** | The version of the source | 

## Methods

### NewSourcesPostRequest

`func NewSourcesPostRequest(name string, url string, version string, ) *SourcesPostRequest`

NewSourcesPostRequest instantiates a new SourcesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcesPostRequestWithDefaults

`func NewSourcesPostRequestWithDefaults() *SourcesPostRequest`

NewSourcesPostRequestWithDefaults instantiates a new SourcesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SourcesPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourcesPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourcesPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *SourcesPostRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SourcesPostRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SourcesPostRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVersion

`func (o *SourcesPostRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SourcesPostRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SourcesPostRequest) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



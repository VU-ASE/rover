# SourcesGet200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the source | 
**Url** | **string** | The URL of the source (without scheme) | 
**Version** | **string** |  | 
**Sha** | Pointer to **string** | The SHA256 hash of the source download, computed over the ZIP file downloaded | [optional] 

## Methods

### NewSourcesGet200ResponseInner

`func NewSourcesGet200ResponseInner(name string, url string, version string, ) *SourcesGet200ResponseInner`

NewSourcesGet200ResponseInner instantiates a new SourcesGet200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcesGet200ResponseInnerWithDefaults

`func NewSourcesGet200ResponseInnerWithDefaults() *SourcesGet200ResponseInner`

NewSourcesGet200ResponseInnerWithDefaults instantiates a new SourcesGet200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SourcesGet200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourcesGet200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourcesGet200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *SourcesGet200ResponseInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SourcesGet200ResponseInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SourcesGet200ResponseInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVersion

`func (o *SourcesGet200ResponseInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SourcesGet200ResponseInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SourcesGet200ResponseInner) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetSha

`func (o *SourcesGet200ResponseInner) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *SourcesGet200ResponseInner) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *SourcesGet200ResponseInner) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *SourcesGet200ResponseInner) HasSha() bool`

HasSha returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



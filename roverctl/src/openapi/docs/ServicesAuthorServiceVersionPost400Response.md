# ServicesAuthorServiceVersionPost400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | The error message | 
**BuildLog** | **[]string** | The build log (one log line per item) | 

## Methods

### NewServicesAuthorServiceVersionPost400Response

`func NewServicesAuthorServiceVersionPost400Response(message string, buildLog []string, ) *ServicesAuthorServiceVersionPost400Response`

NewServicesAuthorServiceVersionPost400Response instantiates a new ServicesAuthorServiceVersionPost400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesAuthorServiceVersionPost400ResponseWithDefaults

`func NewServicesAuthorServiceVersionPost400ResponseWithDefaults() *ServicesAuthorServiceVersionPost400Response`

NewServicesAuthorServiceVersionPost400ResponseWithDefaults instantiates a new ServicesAuthorServiceVersionPost400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ServicesAuthorServiceVersionPost400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServicesAuthorServiceVersionPost400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServicesAuthorServiceVersionPost400Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetBuildLog

`func (o *ServicesAuthorServiceVersionPost400Response) GetBuildLog() []string`

GetBuildLog returns the BuildLog field if non-nil, zero value otherwise.

### GetBuildLogOk

`func (o *ServicesAuthorServiceVersionPost400Response) GetBuildLogOk() (*[]string, bool)`

GetBuildLogOk returns a tuple with the BuildLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildLog

`func (o *ServicesAuthorServiceVersionPost400Response) SetBuildLog(v []string)`

SetBuildLog sets BuildLog field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



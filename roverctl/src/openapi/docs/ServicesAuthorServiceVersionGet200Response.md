# ServicesAuthorServiceVersionGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuiltAt** | Pointer to **int64** | The time this version was last built as milliseconds since epoch, not set if the service was never built | [optional] 
**Inputs** | [**[]ServicesAuthorServiceVersionGet200ResponseInputsInner**](ServicesAuthorServiceVersionGet200ResponseInputsInner.md) | The dependencies/inputs of this service version | 
**Outputs** | **[]string** | The output streams of this service version | 
**Configuration** | [**[]ServicesAuthorServiceVersionGet200ResponseConfigurationInner**](ServicesAuthorServiceVersionGet200ResponseConfigurationInner.md) | All configuration values of this service version and their tunability | 

## Methods

### NewServicesAuthorServiceVersionGet200Response

`func NewServicesAuthorServiceVersionGet200Response(inputs []ServicesAuthorServiceVersionGet200ResponseInputsInner, outputs []string, configuration []ServicesAuthorServiceVersionGet200ResponseConfigurationInner, ) *ServicesAuthorServiceVersionGet200Response`

NewServicesAuthorServiceVersionGet200Response instantiates a new ServicesAuthorServiceVersionGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesAuthorServiceVersionGet200ResponseWithDefaults

`func NewServicesAuthorServiceVersionGet200ResponseWithDefaults() *ServicesAuthorServiceVersionGet200Response`

NewServicesAuthorServiceVersionGet200ResponseWithDefaults instantiates a new ServicesAuthorServiceVersionGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuiltAt

`func (o *ServicesAuthorServiceVersionGet200Response) GetBuiltAt() int64`

GetBuiltAt returns the BuiltAt field if non-nil, zero value otherwise.

### GetBuiltAtOk

`func (o *ServicesAuthorServiceVersionGet200Response) GetBuiltAtOk() (*int64, bool)`

GetBuiltAtOk returns a tuple with the BuiltAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltAt

`func (o *ServicesAuthorServiceVersionGet200Response) SetBuiltAt(v int64)`

SetBuiltAt sets BuiltAt field to given value.

### HasBuiltAt

`func (o *ServicesAuthorServiceVersionGet200Response) HasBuiltAt() bool`

HasBuiltAt returns a boolean if a field has been set.

### GetInputs

`func (o *ServicesAuthorServiceVersionGet200Response) GetInputs() []ServicesAuthorServiceVersionGet200ResponseInputsInner`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ServicesAuthorServiceVersionGet200Response) GetInputsOk() (*[]ServicesAuthorServiceVersionGet200ResponseInputsInner, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ServicesAuthorServiceVersionGet200Response) SetInputs(v []ServicesAuthorServiceVersionGet200ResponseInputsInner)`

SetInputs sets Inputs field to given value.


### GetOutputs

`func (o *ServicesAuthorServiceVersionGet200Response) GetOutputs() []string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *ServicesAuthorServiceVersionGet200Response) GetOutputsOk() (*[]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *ServicesAuthorServiceVersionGet200Response) SetOutputs(v []string)`

SetOutputs sets Outputs field to given value.


### GetConfiguration

`func (o *ServicesAuthorServiceVersionGet200Response) GetConfiguration() []ServicesAuthorServiceVersionGet200ResponseConfigurationInner`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ServicesAuthorServiceVersionGet200Response) GetConfigurationOk() (*[]ServicesAuthorServiceVersionGet200ResponseConfigurationInner, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ServicesAuthorServiceVersionGet200Response) SetConfiguration(v []ServicesAuthorServiceVersionGet200ResponseConfigurationInner)`

SetConfiguration sets Configuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



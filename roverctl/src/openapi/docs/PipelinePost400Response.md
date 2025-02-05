# PipelinePost400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Additional information | [optional] 
**ValidationErrors** | [**PipelinePost400ResponseValidationErrors**](PipelinePost400ResponseValidationErrors.md) |  | 

## Methods

### NewPipelinePost400Response

`func NewPipelinePost400Response(validationErrors PipelinePost400ResponseValidationErrors, ) *PipelinePost400Response`

NewPipelinePost400Response instantiates a new PipelinePost400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelinePost400ResponseWithDefaults

`func NewPipelinePost400ResponseWithDefaults() *PipelinePost400Response`

NewPipelinePost400ResponseWithDefaults instantiates a new PipelinePost400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PipelinePost400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PipelinePost400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PipelinePost400Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PipelinePost400Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetValidationErrors

`func (o *PipelinePost400Response) GetValidationErrors() PipelinePost400ResponseValidationErrors`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *PipelinePost400Response) GetValidationErrorsOk() (*PipelinePost400ResponseValidationErrors, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *PipelinePost400Response) SetValidationErrors(v PipelinePost400ResponseValidationErrors)`

SetValidationErrors sets ValidationErrors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



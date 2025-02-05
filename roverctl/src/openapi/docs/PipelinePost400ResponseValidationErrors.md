# PipelinePost400ResponseValidationErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnmetStreams** | Pointer to [**[]UnmetStreamError**](UnmetStreamError.md) |  | [optional] 
**UnmetServices** | Pointer to [**[]UnmetServiceError**](UnmetServiceError.md) |  | [optional] 
**DuplicateService** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPipelinePost400ResponseValidationErrors

`func NewPipelinePost400ResponseValidationErrors() *PipelinePost400ResponseValidationErrors`

NewPipelinePost400ResponseValidationErrors instantiates a new PipelinePost400ResponseValidationErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelinePost400ResponseValidationErrorsWithDefaults

`func NewPipelinePost400ResponseValidationErrorsWithDefaults() *PipelinePost400ResponseValidationErrors`

NewPipelinePost400ResponseValidationErrorsWithDefaults instantiates a new PipelinePost400ResponseValidationErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnmetStreams

`func (o *PipelinePost400ResponseValidationErrors) GetUnmetStreams() []UnmetStreamError`

GetUnmetStreams returns the UnmetStreams field if non-nil, zero value otherwise.

### GetUnmetStreamsOk

`func (o *PipelinePost400ResponseValidationErrors) GetUnmetStreamsOk() (*[]UnmetStreamError, bool)`

GetUnmetStreamsOk returns a tuple with the UnmetStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmetStreams

`func (o *PipelinePost400ResponseValidationErrors) SetUnmetStreams(v []UnmetStreamError)`

SetUnmetStreams sets UnmetStreams field to given value.

### HasUnmetStreams

`func (o *PipelinePost400ResponseValidationErrors) HasUnmetStreams() bool`

HasUnmetStreams returns a boolean if a field has been set.

### GetUnmetServices

`func (o *PipelinePost400ResponseValidationErrors) GetUnmetServices() []UnmetServiceError`

GetUnmetServices returns the UnmetServices field if non-nil, zero value otherwise.

### GetUnmetServicesOk

`func (o *PipelinePost400ResponseValidationErrors) GetUnmetServicesOk() (*[]UnmetServiceError, bool)`

GetUnmetServicesOk returns a tuple with the UnmetServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmetServices

`func (o *PipelinePost400ResponseValidationErrors) SetUnmetServices(v []UnmetServiceError)`

SetUnmetServices sets UnmetServices field to given value.

### HasUnmetServices

`func (o *PipelinePost400ResponseValidationErrors) HasUnmetServices() bool`

HasUnmetServices returns a boolean if a field has been set.

### GetDuplicateService

`func (o *PipelinePost400ResponseValidationErrors) GetDuplicateService() []string`

GetDuplicateService returns the DuplicateService field if non-nil, zero value otherwise.

### GetDuplicateServiceOk

`func (o *PipelinePost400ResponseValidationErrors) GetDuplicateServiceOk() (*[]string, bool)`

GetDuplicateServiceOk returns a tuple with the DuplicateService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateService

`func (o *PipelinePost400ResponseValidationErrors) SetDuplicateService(v []string)`

SetDuplicateService sets DuplicateService field to given value.

### HasDuplicateService

`func (o *PipelinePost400ResponseValidationErrors) HasDuplicateService() bool`

HasDuplicateService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



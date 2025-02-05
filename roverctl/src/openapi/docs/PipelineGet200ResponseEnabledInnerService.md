# PipelineGet200ResponseEnabledInnerService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the service | 
**Version** | **string** | The version of the service | 
**Author** | **string** | The author of the service | 
**Faults** | **int32** | The number of faults that have occurred (causing the pipeline to restart) since pipeline.last_start | 
**Exit** | **int32** | The most recent exit code returned by the process. | 

## Methods

### NewPipelineGet200ResponseEnabledInnerService

`func NewPipelineGet200ResponseEnabledInnerService(name string, version string, author string, faults int32, exit int32, ) *PipelineGet200ResponseEnabledInnerService`

NewPipelineGet200ResponseEnabledInnerService instantiates a new PipelineGet200ResponseEnabledInnerService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineGet200ResponseEnabledInnerServiceWithDefaults

`func NewPipelineGet200ResponseEnabledInnerServiceWithDefaults() *PipelineGet200ResponseEnabledInnerService`

NewPipelineGet200ResponseEnabledInnerServiceWithDefaults instantiates a new PipelineGet200ResponseEnabledInnerService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PipelineGet200ResponseEnabledInnerService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PipelineGet200ResponseEnabledInnerService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PipelineGet200ResponseEnabledInnerService) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *PipelineGet200ResponseEnabledInnerService) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PipelineGet200ResponseEnabledInnerService) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PipelineGet200ResponseEnabledInnerService) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetAuthor

`func (o *PipelineGet200ResponseEnabledInnerService) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *PipelineGet200ResponseEnabledInnerService) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *PipelineGet200ResponseEnabledInnerService) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetFaults

`func (o *PipelineGet200ResponseEnabledInnerService) GetFaults() int32`

GetFaults returns the Faults field if non-nil, zero value otherwise.

### GetFaultsOk

`func (o *PipelineGet200ResponseEnabledInnerService) GetFaultsOk() (*int32, bool)`

GetFaultsOk returns a tuple with the Faults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaults

`func (o *PipelineGet200ResponseEnabledInnerService) SetFaults(v int32)`

SetFaults sets Faults field to given value.


### GetExit

`func (o *PipelineGet200ResponseEnabledInnerService) GetExit() int32`

GetExit returns the Exit field if non-nil, zero value otherwise.

### GetExitOk

`func (o *PipelineGet200ResponseEnabledInnerService) GetExitOk() (*int32, bool)`

GetExitOk returns a tuple with the Exit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExit

`func (o *PipelineGet200ResponseEnabledInnerService) SetExit(v int32)`

SetExit sets Exit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



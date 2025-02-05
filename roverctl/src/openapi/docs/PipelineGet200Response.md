# PipelineGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**PipelineStatus**](PipelineStatus.md) |  | 
**LastStart** | Pointer to **int64** | Milliseconds since epoch when the pipeline was manually started | [optional] 
**LastStop** | Pointer to **int64** | Milliseconds since epoch when the pipeline was manually stopped | [optional] 
**LastRestart** | Pointer to **int64** | Milliseconds since epoch when the pipeline was automatically restarted (on process faults) | [optional] 
**Enabled** | [**[]PipelineGet200ResponseEnabledInner**](PipelineGet200ResponseEnabledInner.md) | The list of fully qualified services that are enabled in this pipeline. If the pipeline was started, this includes a process for each service | 

## Methods

### NewPipelineGet200Response

`func NewPipelineGet200Response(status PipelineStatus, enabled []PipelineGet200ResponseEnabledInner, ) *PipelineGet200Response`

NewPipelineGet200Response instantiates a new PipelineGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineGet200ResponseWithDefaults

`func NewPipelineGet200ResponseWithDefaults() *PipelineGet200Response`

NewPipelineGet200ResponseWithDefaults instantiates a new PipelineGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PipelineGet200Response) GetStatus() PipelineStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PipelineGet200Response) GetStatusOk() (*PipelineStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PipelineGet200Response) SetStatus(v PipelineStatus)`

SetStatus sets Status field to given value.


### GetLastStart

`func (o *PipelineGet200Response) GetLastStart() int64`

GetLastStart returns the LastStart field if non-nil, zero value otherwise.

### GetLastStartOk

`func (o *PipelineGet200Response) GetLastStartOk() (*int64, bool)`

GetLastStartOk returns a tuple with the LastStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStart

`func (o *PipelineGet200Response) SetLastStart(v int64)`

SetLastStart sets LastStart field to given value.

### HasLastStart

`func (o *PipelineGet200Response) HasLastStart() bool`

HasLastStart returns a boolean if a field has been set.

### GetLastStop

`func (o *PipelineGet200Response) GetLastStop() int64`

GetLastStop returns the LastStop field if non-nil, zero value otherwise.

### GetLastStopOk

`func (o *PipelineGet200Response) GetLastStopOk() (*int64, bool)`

GetLastStopOk returns a tuple with the LastStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStop

`func (o *PipelineGet200Response) SetLastStop(v int64)`

SetLastStop sets LastStop field to given value.

### HasLastStop

`func (o *PipelineGet200Response) HasLastStop() bool`

HasLastStop returns a boolean if a field has been set.

### GetLastRestart

`func (o *PipelineGet200Response) GetLastRestart() int64`

GetLastRestart returns the LastRestart field if non-nil, zero value otherwise.

### GetLastRestartOk

`func (o *PipelineGet200Response) GetLastRestartOk() (*int64, bool)`

GetLastRestartOk returns a tuple with the LastRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRestart

`func (o *PipelineGet200Response) SetLastRestart(v int64)`

SetLastRestart sets LastRestart field to given value.

### HasLastRestart

`func (o *PipelineGet200Response) HasLastRestart() bool`

HasLastRestart returns a boolean if a field has been set.

### GetEnabled

`func (o *PipelineGet200Response) GetEnabled() []PipelineGet200ResponseEnabledInner`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PipelineGet200Response) GetEnabledOk() (*[]PipelineGet200ResponseEnabledInner, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PipelineGet200Response) SetEnabled(v []PipelineGet200ResponseEnabledInner)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



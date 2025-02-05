# PipelineGet200ResponseEnabledInnerProcess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pid** | **int32** | The process ID. Depending on the status, this PID might not exist anymore | 
**Status** | [**ProcessStatus**](ProcessStatus.md) |  | 
**Uptime** | **int64** | The number of milliseconds the process has been running | 
**Memory** | **int32** | The amount of memory used by the process in megabytes | 
**Cpu** | **int32** | The percentage of CPU used by the process | 

## Methods

### NewPipelineGet200ResponseEnabledInnerProcess

`func NewPipelineGet200ResponseEnabledInnerProcess(pid int32, status ProcessStatus, uptime int64, memory int32, cpu int32, ) *PipelineGet200ResponseEnabledInnerProcess`

NewPipelineGet200ResponseEnabledInnerProcess instantiates a new PipelineGet200ResponseEnabledInnerProcess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineGet200ResponseEnabledInnerProcessWithDefaults

`func NewPipelineGet200ResponseEnabledInnerProcessWithDefaults() *PipelineGet200ResponseEnabledInnerProcess`

NewPipelineGet200ResponseEnabledInnerProcessWithDefaults instantiates a new PipelineGet200ResponseEnabledInnerProcess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPid

`func (o *PipelineGet200ResponseEnabledInnerProcess) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *PipelineGet200ResponseEnabledInnerProcess) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *PipelineGet200ResponseEnabledInnerProcess) SetPid(v int32)`

SetPid sets Pid field to given value.


### GetStatus

`func (o *PipelineGet200ResponseEnabledInnerProcess) GetStatus() ProcessStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PipelineGet200ResponseEnabledInnerProcess) GetStatusOk() (*ProcessStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PipelineGet200ResponseEnabledInnerProcess) SetStatus(v ProcessStatus)`

SetStatus sets Status field to given value.


### GetUptime

`func (o *PipelineGet200ResponseEnabledInnerProcess) GetUptime() int64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *PipelineGet200ResponseEnabledInnerProcess) GetUptimeOk() (*int64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *PipelineGet200ResponseEnabledInnerProcess) SetUptime(v int64)`

SetUptime sets Uptime field to given value.


### GetMemory

`func (o *PipelineGet200ResponseEnabledInnerProcess) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *PipelineGet200ResponseEnabledInnerProcess) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *PipelineGet200ResponseEnabledInnerProcess) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### GetCpu

`func (o *PipelineGet200ResponseEnabledInnerProcess) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *PipelineGet200ResponseEnabledInnerProcess) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *PipelineGet200ResponseEnabledInnerProcess) SetCpu(v int32)`

SetCpu sets Cpu field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



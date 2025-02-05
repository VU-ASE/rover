# StatusGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**DaemonStatus**](DaemonStatus.md) |  | 
**ErrorMessage** | Pointer to **string** | Error message of the daemon status | [optional] 
**Version** | **string** | The version of the roverd daemon | 
**Uptime** | **int64** | The number of milliseconds the roverd daemon process has been running | 
**Os** | **string** | The operating system of the rover | 
**Systime** | **int64** | The system time of the rover as milliseconds since epoch | 
**RoverId** | Pointer to **int32** | The unique identifier of the rover | [optional] 
**RoverName** | Pointer to **string** | The unique name of the rover | [optional] 
**Memory** | [**StatusGet200ResponseMemory**](StatusGet200ResponseMemory.md) |  | 
**Cpu** | [**[]StatusGet200ResponseCpuInner**](StatusGet200ResponseCpuInner.md) | The CPU usage of the roverd process | 

## Methods

### NewStatusGet200Response

`func NewStatusGet200Response(status DaemonStatus, version string, uptime int64, os string, systime int64, memory StatusGet200ResponseMemory, cpu []StatusGet200ResponseCpuInner, ) *StatusGet200Response`

NewStatusGet200Response instantiates a new StatusGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusGet200ResponseWithDefaults

`func NewStatusGet200ResponseWithDefaults() *StatusGet200Response`

NewStatusGet200ResponseWithDefaults instantiates a new StatusGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *StatusGet200Response) GetStatus() DaemonStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StatusGet200Response) GetStatusOk() (*DaemonStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StatusGet200Response) SetStatus(v DaemonStatus)`

SetStatus sets Status field to given value.


### GetErrorMessage

`func (o *StatusGet200Response) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *StatusGet200Response) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *StatusGet200Response) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *StatusGet200Response) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetVersion

`func (o *StatusGet200Response) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StatusGet200Response) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StatusGet200Response) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetUptime

`func (o *StatusGet200Response) GetUptime() int64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *StatusGet200Response) GetUptimeOk() (*int64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *StatusGet200Response) SetUptime(v int64)`

SetUptime sets Uptime field to given value.


### GetOs

`func (o *StatusGet200Response) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *StatusGet200Response) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *StatusGet200Response) SetOs(v string)`

SetOs sets Os field to given value.


### GetSystime

`func (o *StatusGet200Response) GetSystime() int64`

GetSystime returns the Systime field if non-nil, zero value otherwise.

### GetSystimeOk

`func (o *StatusGet200Response) GetSystimeOk() (*int64, bool)`

GetSystimeOk returns a tuple with the Systime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystime

`func (o *StatusGet200Response) SetSystime(v int64)`

SetSystime sets Systime field to given value.


### GetRoverId

`func (o *StatusGet200Response) GetRoverId() int32`

GetRoverId returns the RoverId field if non-nil, zero value otherwise.

### GetRoverIdOk

`func (o *StatusGet200Response) GetRoverIdOk() (*int32, bool)`

GetRoverIdOk returns a tuple with the RoverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoverId

`func (o *StatusGet200Response) SetRoverId(v int32)`

SetRoverId sets RoverId field to given value.

### HasRoverId

`func (o *StatusGet200Response) HasRoverId() bool`

HasRoverId returns a boolean if a field has been set.

### GetRoverName

`func (o *StatusGet200Response) GetRoverName() string`

GetRoverName returns the RoverName field if non-nil, zero value otherwise.

### GetRoverNameOk

`func (o *StatusGet200Response) GetRoverNameOk() (*string, bool)`

GetRoverNameOk returns a tuple with the RoverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoverName

`func (o *StatusGet200Response) SetRoverName(v string)`

SetRoverName sets RoverName field to given value.

### HasRoverName

`func (o *StatusGet200Response) HasRoverName() bool`

HasRoverName returns a boolean if a field has been set.

### GetMemory

`func (o *StatusGet200Response) GetMemory() StatusGet200ResponseMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *StatusGet200Response) GetMemoryOk() (*StatusGet200ResponseMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *StatusGet200Response) SetMemory(v StatusGet200ResponseMemory)`

SetMemory sets Memory field to given value.


### GetCpu

`func (o *StatusGet200Response) GetCpu() []StatusGet200ResponseCpuInner`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *StatusGet200Response) GetCpuOk() (*[]StatusGet200ResponseCpuInner, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *StatusGet200Response) SetCpu(v []StatusGet200ResponseCpuInner)`

SetCpu sets Cpu field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



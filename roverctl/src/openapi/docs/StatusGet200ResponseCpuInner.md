# StatusGet200ResponseCpuInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Core** | **int32** | The core number | 
**Total** | **int32** | The total amount of CPU available on the core | 
**Used** | **int32** | The amount of CPU used on the core | 

## Methods

### NewStatusGet200ResponseCpuInner

`func NewStatusGet200ResponseCpuInner(core int32, total int32, used int32, ) *StatusGet200ResponseCpuInner`

NewStatusGet200ResponseCpuInner instantiates a new StatusGet200ResponseCpuInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusGet200ResponseCpuInnerWithDefaults

`func NewStatusGet200ResponseCpuInnerWithDefaults() *StatusGet200ResponseCpuInner`

NewStatusGet200ResponseCpuInnerWithDefaults instantiates a new StatusGet200ResponseCpuInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCore

`func (o *StatusGet200ResponseCpuInner) GetCore() int32`

GetCore returns the Core field if non-nil, zero value otherwise.

### GetCoreOk

`func (o *StatusGet200ResponseCpuInner) GetCoreOk() (*int32, bool)`

GetCoreOk returns a tuple with the Core field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCore

`func (o *StatusGet200ResponseCpuInner) SetCore(v int32)`

SetCore sets Core field to given value.


### GetTotal

`func (o *StatusGet200ResponseCpuInner) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *StatusGet200ResponseCpuInner) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *StatusGet200ResponseCpuInner) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetUsed

`func (o *StatusGet200ResponseCpuInner) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *StatusGet200ResponseCpuInner) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *StatusGet200ResponseCpuInner) SetUsed(v int32)`

SetUsed sets Used field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StatusGet200ResponseMemory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | The total amount of memory available on the rover in megabytes | 
**Used** | **int32** | The amount of memory used on the rover in megabytes | 

## Methods

### NewStatusGet200ResponseMemory

`func NewStatusGet200ResponseMemory(total int32, used int32, ) *StatusGet200ResponseMemory`

NewStatusGet200ResponseMemory instantiates a new StatusGet200ResponseMemory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusGet200ResponseMemoryWithDefaults

`func NewStatusGet200ResponseMemoryWithDefaults() *StatusGet200ResponseMemory`

NewStatusGet200ResponseMemoryWithDefaults instantiates a new StatusGet200ResponseMemory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *StatusGet200ResponseMemory) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *StatusGet200ResponseMemory) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *StatusGet200ResponseMemory) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetUsed

`func (o *StatusGet200ResponseMemory) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *StatusGet200ResponseMemory) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *StatusGet200ResponseMemory) SetUsed(v int32)`

SetUsed sets Used field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



/*
roverd REST API

API exposed from each rover to allow process, service, source and file management

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the StatusGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusGet200Response{}

// StatusGet200Response struct for StatusGet200Response
type StatusGet200Response struct {
	Status DaemonStatus `json:"status"`
	// Error message of the daemon status
	ErrorMessage *string `json:"error_message,omitempty"`
	// The version of the roverd daemon
	Version string `json:"version"`
	// The number of milliseconds the roverd daemon process has been running
	Uptime int64 `json:"uptime"`
	// The operating system of the rover
	Os string `json:"os"`
	// The system time of the rover as milliseconds since epoch
	Systime int64 `json:"systime"`
	// The unique identifier of the rover
	RoverId *int32 `json:"rover_id,omitempty"`
	// The unique name of the rover
	RoverName *string `json:"rover_name,omitempty"`
	Memory StatusGet200ResponseMemory `json:"memory"`
	// The CPU usage of the roverd process
	Cpu []StatusGet200ResponseCpuInner `json:"cpu"`
}

type _StatusGet200Response StatusGet200Response

// NewStatusGet200Response instantiates a new StatusGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusGet200Response(status DaemonStatus, version string, uptime int64, os string, systime int64, memory StatusGet200ResponseMemory, cpu []StatusGet200ResponseCpuInner) *StatusGet200Response {
	this := StatusGet200Response{}
	this.Status = status
	this.Version = version
	this.Uptime = uptime
	this.Os = os
	this.Systime = systime
	this.Memory = memory
	this.Cpu = cpu
	return &this
}

// NewStatusGet200ResponseWithDefaults instantiates a new StatusGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusGet200ResponseWithDefaults() *StatusGet200Response {
	this := StatusGet200Response{}
	return &this
}

// GetStatus returns the Status field value
func (o *StatusGet200Response) GetStatus() DaemonStatus {
	if o == nil {
		var ret DaemonStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *StatusGet200Response) GetStatusOk() (*DaemonStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *StatusGet200Response) SetStatus(v DaemonStatus) {
	o.Status = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *StatusGet200Response) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusGet200Response) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *StatusGet200Response) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *StatusGet200Response) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetVersion returns the Version field value
func (o *StatusGet200Response) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *StatusGet200Response) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *StatusGet200Response) SetVersion(v string) {
	o.Version = v
}

// GetUptime returns the Uptime field value
func (o *StatusGet200Response) GetUptime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value
// and a boolean to check if the value has been set.
func (o *StatusGet200Response) GetUptimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uptime, true
}

// SetUptime sets field value
func (o *StatusGet200Response) SetUptime(v int64) {
	o.Uptime = v
}

// GetOs returns the Os field value
func (o *StatusGet200Response) GetOs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Os
}

// GetOsOk returns a tuple with the Os field value
// and a boolean to check if the value has been set.
func (o *StatusGet200Response) GetOsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Os, true
}

// SetOs sets field value
func (o *StatusGet200Response) SetOs(v string) {
	o.Os = v
}

// GetSystime returns the Systime field value
func (o *StatusGet200Response) GetSystime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Systime
}

// GetSystimeOk returns a tuple with the Systime field value
// and a boolean to check if the value has been set.
func (o *StatusGet200Response) GetSystimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Systime, true
}

// SetSystime sets field value
func (o *StatusGet200Response) SetSystime(v int64) {
	o.Systime = v
}

// GetRoverId returns the RoverId field value if set, zero value otherwise.
func (o *StatusGet200Response) GetRoverId() int32 {
	if o == nil || IsNil(o.RoverId) {
		var ret int32
		return ret
	}
	return *o.RoverId
}

// GetRoverIdOk returns a tuple with the RoverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusGet200Response) GetRoverIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RoverId) {
		return nil, false
	}
	return o.RoverId, true
}

// HasRoverId returns a boolean if a field has been set.
func (o *StatusGet200Response) HasRoverId() bool {
	if o != nil && !IsNil(o.RoverId) {
		return true
	}

	return false
}

// SetRoverId gets a reference to the given int32 and assigns it to the RoverId field.
func (o *StatusGet200Response) SetRoverId(v int32) {
	o.RoverId = &v
}

// GetRoverName returns the RoverName field value if set, zero value otherwise.
func (o *StatusGet200Response) GetRoverName() string {
	if o == nil || IsNil(o.RoverName) {
		var ret string
		return ret
	}
	return *o.RoverName
}

// GetRoverNameOk returns a tuple with the RoverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusGet200Response) GetRoverNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoverName) {
		return nil, false
	}
	return o.RoverName, true
}

// HasRoverName returns a boolean if a field has been set.
func (o *StatusGet200Response) HasRoverName() bool {
	if o != nil && !IsNil(o.RoverName) {
		return true
	}

	return false
}

// SetRoverName gets a reference to the given string and assigns it to the RoverName field.
func (o *StatusGet200Response) SetRoverName(v string) {
	o.RoverName = &v
}

// GetMemory returns the Memory field value
func (o *StatusGet200Response) GetMemory() StatusGet200ResponseMemory {
	if o == nil {
		var ret StatusGet200ResponseMemory
		return ret
	}

	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *StatusGet200Response) GetMemoryOk() (*StatusGet200ResponseMemory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memory, true
}

// SetMemory sets field value
func (o *StatusGet200Response) SetMemory(v StatusGet200ResponseMemory) {
	o.Memory = v
}

// GetCpu returns the Cpu field value
func (o *StatusGet200Response) GetCpu() []StatusGet200ResponseCpuInner {
	if o == nil {
		var ret []StatusGet200ResponseCpuInner
		return ret
	}

	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *StatusGet200Response) GetCpuOk() ([]StatusGet200ResponseCpuInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpu, true
}

// SetCpu sets field value
func (o *StatusGet200Response) SetCpu(v []StatusGet200ResponseCpuInner) {
	o.Cpu = v
}

func (o StatusGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if !IsNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	toSerialize["version"] = o.Version
	toSerialize["uptime"] = o.Uptime
	toSerialize["os"] = o.Os
	toSerialize["systime"] = o.Systime
	if !IsNil(o.RoverId) {
		toSerialize["rover_id"] = o.RoverId
	}
	if !IsNil(o.RoverName) {
		toSerialize["rover_name"] = o.RoverName
	}
	toSerialize["memory"] = o.Memory
	toSerialize["cpu"] = o.Cpu
	return toSerialize, nil
}

func (o *StatusGet200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"version",
		"uptime",
		"os",
		"systime",
		"memory",
		"cpu",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varStatusGet200Response := _StatusGet200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStatusGet200Response)

	if err != nil {
		return err
	}

	*o = StatusGet200Response(varStatusGet200Response)

	return err
}

type NullableStatusGet200Response struct {
	value *StatusGet200Response
	isSet bool
}

func (v NullableStatusGet200Response) Get() *StatusGet200Response {
	return v.value
}

func (v *NullableStatusGet200Response) Set(val *StatusGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusGet200Response(val *StatusGet200Response) *NullableStatusGet200Response {
	return &NullableStatusGet200Response{value: val, isSet: true}
}

func (v NullableStatusGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



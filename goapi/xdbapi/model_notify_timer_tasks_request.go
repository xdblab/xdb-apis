/*
XDB APIs

This APIs between xdb service and SDKs

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xdbapi

import (
	"encoding/json"
)

// checks if the NotifyTimerTasksRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifyTimerTasksRequest{}

// NotifyTimerTasksRequest struct for NotifyTimerTasksRequest
type NotifyTimerTasksRequest struct {
	ShardId int32 `json:"shardId"`
	// the fire timestamp of all timer tasks to pull
	FireTimestamps []int64 `json:"fireTimestamps"`
	// optional field for distributed database without global secondary index, to pull for specific task rather than a page
	Namespace *string `json:"namespace,omitempty"`
	// optional field for distributed database without global secondary index, to pull for specific task rather than a page
	ProcessId *string `json:"processId,omitempty"`
	// optional field for distributed database without global secondary index, to pull for specific task rather than a page
	ProcessExecutionId *string `json:"processExecutionId,omitempty"`
}

// NewNotifyTimerTasksRequest instantiates a new NotifyTimerTasksRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyTimerTasksRequest(shardId int32, fireTimestamps []int64) *NotifyTimerTasksRequest {
	this := NotifyTimerTasksRequest{}
	this.ShardId = shardId
	this.FireTimestamps = fireTimestamps
	return &this
}

// NewNotifyTimerTasksRequestWithDefaults instantiates a new NotifyTimerTasksRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyTimerTasksRequestWithDefaults() *NotifyTimerTasksRequest {
	this := NotifyTimerTasksRequest{}
	return &this
}

// GetShardId returns the ShardId field value
func (o *NotifyTimerTasksRequest) GetShardId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ShardId
}

// GetShardIdOk returns a tuple with the ShardId field value
// and a boolean to check if the value has been set.
func (o *NotifyTimerTasksRequest) GetShardIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShardId, true
}

// SetShardId sets field value
func (o *NotifyTimerTasksRequest) SetShardId(v int32) {
	o.ShardId = v
}

// GetFireTimestamps returns the FireTimestamps field value
func (o *NotifyTimerTasksRequest) GetFireTimestamps() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.FireTimestamps
}

// GetFireTimestampsOk returns a tuple with the FireTimestamps field value
// and a boolean to check if the value has been set.
func (o *NotifyTimerTasksRequest) GetFireTimestampsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FireTimestamps, true
}

// SetFireTimestamps sets field value
func (o *NotifyTimerTasksRequest) SetFireTimestamps(v []int64) {
	o.FireTimestamps = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *NotifyTimerTasksRequest) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyTimerTasksRequest) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *NotifyTimerTasksRequest) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *NotifyTimerTasksRequest) SetNamespace(v string) {
	o.Namespace = &v
}

// GetProcessId returns the ProcessId field value if set, zero value otherwise.
func (o *NotifyTimerTasksRequest) GetProcessId() string {
	if o == nil || IsNil(o.ProcessId) {
		var ret string
		return ret
	}
	return *o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyTimerTasksRequest) GetProcessIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessId) {
		return nil, false
	}
	return o.ProcessId, true
}

// HasProcessId returns a boolean if a field has been set.
func (o *NotifyTimerTasksRequest) HasProcessId() bool {
	if o != nil && !IsNil(o.ProcessId) {
		return true
	}

	return false
}

// SetProcessId gets a reference to the given string and assigns it to the ProcessId field.
func (o *NotifyTimerTasksRequest) SetProcessId(v string) {
	o.ProcessId = &v
}

// GetProcessExecutionId returns the ProcessExecutionId field value if set, zero value otherwise.
func (o *NotifyTimerTasksRequest) GetProcessExecutionId() string {
	if o == nil || IsNil(o.ProcessExecutionId) {
		var ret string
		return ret
	}
	return *o.ProcessExecutionId
}

// GetProcessExecutionIdOk returns a tuple with the ProcessExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyTimerTasksRequest) GetProcessExecutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessExecutionId) {
		return nil, false
	}
	return o.ProcessExecutionId, true
}

// HasProcessExecutionId returns a boolean if a field has been set.
func (o *NotifyTimerTasksRequest) HasProcessExecutionId() bool {
	if o != nil && !IsNil(o.ProcessExecutionId) {
		return true
	}

	return false
}

// SetProcessExecutionId gets a reference to the given string and assigns it to the ProcessExecutionId field.
func (o *NotifyTimerTasksRequest) SetProcessExecutionId(v string) {
	o.ProcessExecutionId = &v
}

func (o NotifyTimerTasksRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifyTimerTasksRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shardId"] = o.ShardId
	toSerialize["fireTimestamps"] = o.FireTimestamps
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.ProcessId) {
		toSerialize["processId"] = o.ProcessId
	}
	if !IsNil(o.ProcessExecutionId) {
		toSerialize["processExecutionId"] = o.ProcessExecutionId
	}
	return toSerialize, nil
}

type NullableNotifyTimerTasksRequest struct {
	value *NotifyTimerTasksRequest
	isSet bool
}

func (v NullableNotifyTimerTasksRequest) Get() *NotifyTimerTasksRequest {
	return v.value
}

func (v *NullableNotifyTimerTasksRequest) Set(val *NotifyTimerTasksRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyTimerTasksRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyTimerTasksRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyTimerTasksRequest(val *NotifyTimerTasksRequest) *NullableNotifyTimerTasksRequest {
	return &NullableNotifyTimerTasksRequest{value: val, isSet: true}
}

func (v NullableNotifyTimerTasksRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyTimerTasksRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

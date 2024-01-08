/*
xCherry APIs

This APIs between xCherry service and SDKs

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xcapi

import (
	"encoding/json"
)

// checks if the ProcessExecutionListInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessExecutionListInfo{}

// ProcessExecutionListInfo struct for ProcessExecutionListInfo
type ProcessExecutionListInfo struct {
	Namespace          *string        `json:"namespace,omitempty"`
	ProcessId          *string        `json:"processId,omitempty"`
	ProcessExecutionId *string        `json:"processExecutionId,omitempty"`
	ProcessType        *string        `json:"processType,omitempty"`
	StartTimestamp     *int32         `json:"startTimestamp,omitempty"`
	CloseTimestamp     *int32         `json:"closeTimestamp,omitempty"`
	Status             *ProcessStatus `json:"status,omitempty"`
}

// NewProcessExecutionListInfo instantiates a new ProcessExecutionListInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessExecutionListInfo() *ProcessExecutionListInfo {
	this := ProcessExecutionListInfo{}
	return &this
}

// NewProcessExecutionListInfoWithDefaults instantiates a new ProcessExecutionListInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessExecutionListInfoWithDefaults() *ProcessExecutionListInfo {
	this := ProcessExecutionListInfo{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ProcessExecutionListInfo) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessExecutionListInfo) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ProcessExecutionListInfo) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ProcessExecutionListInfo) SetNamespace(v string) {
	o.Namespace = &v
}

// GetProcessId returns the ProcessId field value if set, zero value otherwise.
func (o *ProcessExecutionListInfo) GetProcessId() string {
	if o == nil || IsNil(o.ProcessId) {
		var ret string
		return ret
	}
	return *o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessExecutionListInfo) GetProcessIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessId) {
		return nil, false
	}
	return o.ProcessId, true
}

// HasProcessId returns a boolean if a field has been set.
func (o *ProcessExecutionListInfo) HasProcessId() bool {
	if o != nil && !IsNil(o.ProcessId) {
		return true
	}

	return false
}

// SetProcessId gets a reference to the given string and assigns it to the ProcessId field.
func (o *ProcessExecutionListInfo) SetProcessId(v string) {
	o.ProcessId = &v
}

// GetProcessExecutionId returns the ProcessExecutionId field value if set, zero value otherwise.
func (o *ProcessExecutionListInfo) GetProcessExecutionId() string {
	if o == nil || IsNil(o.ProcessExecutionId) {
		var ret string
		return ret
	}
	return *o.ProcessExecutionId
}

// GetProcessExecutionIdOk returns a tuple with the ProcessExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessExecutionListInfo) GetProcessExecutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessExecutionId) {
		return nil, false
	}
	return o.ProcessExecutionId, true
}

// HasProcessExecutionId returns a boolean if a field has been set.
func (o *ProcessExecutionListInfo) HasProcessExecutionId() bool {
	if o != nil && !IsNil(o.ProcessExecutionId) {
		return true
	}

	return false
}

// SetProcessExecutionId gets a reference to the given string and assigns it to the ProcessExecutionId field.
func (o *ProcessExecutionListInfo) SetProcessExecutionId(v string) {
	o.ProcessExecutionId = &v
}

// GetProcessType returns the ProcessType field value if set, zero value otherwise.
func (o *ProcessExecutionListInfo) GetProcessType() string {
	if o == nil || IsNil(o.ProcessType) {
		var ret string
		return ret
	}
	return *o.ProcessType
}

// GetProcessTypeOk returns a tuple with the ProcessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessExecutionListInfo) GetProcessTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessType) {
		return nil, false
	}
	return o.ProcessType, true
}

// HasProcessType returns a boolean if a field has been set.
func (o *ProcessExecutionListInfo) HasProcessType() bool {
	if o != nil && !IsNil(o.ProcessType) {
		return true
	}

	return false
}

// SetProcessType gets a reference to the given string and assigns it to the ProcessType field.
func (o *ProcessExecutionListInfo) SetProcessType(v string) {
	o.ProcessType = &v
}

// GetStartTimestamp returns the StartTimestamp field value if set, zero value otherwise.
func (o *ProcessExecutionListInfo) GetStartTimestamp() int32 {
	if o == nil || IsNil(o.StartTimestamp) {
		var ret int32
		return ret
	}
	return *o.StartTimestamp
}

// GetStartTimestampOk returns a tuple with the StartTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessExecutionListInfo) GetStartTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.StartTimestamp) {
		return nil, false
	}
	return o.StartTimestamp, true
}

// HasStartTimestamp returns a boolean if a field has been set.
func (o *ProcessExecutionListInfo) HasStartTimestamp() bool {
	if o != nil && !IsNil(o.StartTimestamp) {
		return true
	}

	return false
}

// SetStartTimestamp gets a reference to the given int32 and assigns it to the StartTimestamp field.
func (o *ProcessExecutionListInfo) SetStartTimestamp(v int32) {
	o.StartTimestamp = &v
}

// GetCloseTimestamp returns the CloseTimestamp field value if set, zero value otherwise.
func (o *ProcessExecutionListInfo) GetCloseTimestamp() int32 {
	if o == nil || IsNil(o.CloseTimestamp) {
		var ret int32
		return ret
	}
	return *o.CloseTimestamp
}

// GetCloseTimestampOk returns a tuple with the CloseTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessExecutionListInfo) GetCloseTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.CloseTimestamp) {
		return nil, false
	}
	return o.CloseTimestamp, true
}

// HasCloseTimestamp returns a boolean if a field has been set.
func (o *ProcessExecutionListInfo) HasCloseTimestamp() bool {
	if o != nil && !IsNil(o.CloseTimestamp) {
		return true
	}

	return false
}

// SetCloseTimestamp gets a reference to the given int32 and assigns it to the CloseTimestamp field.
func (o *ProcessExecutionListInfo) SetCloseTimestamp(v int32) {
	o.CloseTimestamp = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProcessExecutionListInfo) GetStatus() ProcessStatus {
	if o == nil || IsNil(o.Status) {
		var ret ProcessStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessExecutionListInfo) GetStatusOk() (*ProcessStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProcessExecutionListInfo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ProcessStatus and assigns it to the Status field.
func (o *ProcessExecutionListInfo) SetStatus(v ProcessStatus) {
	o.Status = &v
}

func (o ProcessExecutionListInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessExecutionListInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.ProcessId) {
		toSerialize["processId"] = o.ProcessId
	}
	if !IsNil(o.ProcessExecutionId) {
		toSerialize["processExecutionId"] = o.ProcessExecutionId
	}
	if !IsNil(o.ProcessType) {
		toSerialize["processType"] = o.ProcessType
	}
	if !IsNil(o.StartTimestamp) {
		toSerialize["startTimestamp"] = o.StartTimestamp
	}
	if !IsNil(o.CloseTimestamp) {
		toSerialize["closeTimestamp"] = o.CloseTimestamp
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableProcessExecutionListInfo struct {
	value *ProcessExecutionListInfo
	isSet bool
}

func (v NullableProcessExecutionListInfo) Get() *ProcessExecutionListInfo {
	return v.value
}

func (v *NullableProcessExecutionListInfo) Set(val *ProcessExecutionListInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessExecutionListInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessExecutionListInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessExecutionListInfo(val *ProcessExecutionListInfo) *NullableProcessExecutionListInfo {
	return &NullableProcessExecutionListInfo{value: val, isSet: true}
}

func (v NullableProcessExecutionListInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessExecutionListInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

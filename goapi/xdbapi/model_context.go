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

// checks if the Context type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Context{}

// Context struct for Context
type Context struct {
	ProcessId               string `json:"processId"`
	ProcessExecutionId      string `json:"processExecutionId"`
	ProcessStartedTimestamp int64  `json:"processStartedTimestamp"`
	// stateExecutionId is for async state API only
	StateExecutionId *string `json:"stateExecutionId,omitempty"`
	// for async state API only(during backoff retry)
	FirstAttemptTimestamp *int64 `json:"firstAttemptTimestamp,omitempty"`
	// for async state API only(during backoff retry)
	Attempt *int32 `json:"attempt,omitempty"`
}

// NewContext instantiates a new Context object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContext(processId string, processExecutionId string, processStartedTimestamp int64) *Context {
	this := Context{}
	this.ProcessId = processId
	this.ProcessExecutionId = processExecutionId
	this.ProcessStartedTimestamp = processStartedTimestamp
	return &this
}

// NewContextWithDefaults instantiates a new Context object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextWithDefaults() *Context {
	this := Context{}
	return &this
}

// GetProcessId returns the ProcessId field value
func (o *Context) GetProcessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value
// and a boolean to check if the value has been set.
func (o *Context) GetProcessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessId, true
}

// SetProcessId sets field value
func (o *Context) SetProcessId(v string) {
	o.ProcessId = v
}

// GetProcessExecutionId returns the ProcessExecutionId field value
func (o *Context) GetProcessExecutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessExecutionId
}

// GetProcessExecutionIdOk returns a tuple with the ProcessExecutionId field value
// and a boolean to check if the value has been set.
func (o *Context) GetProcessExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessExecutionId, true
}

// SetProcessExecutionId sets field value
func (o *Context) SetProcessExecutionId(v string) {
	o.ProcessExecutionId = v
}

// GetProcessStartedTimestamp returns the ProcessStartedTimestamp field value
func (o *Context) GetProcessStartedTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ProcessStartedTimestamp
}

// GetProcessStartedTimestampOk returns a tuple with the ProcessStartedTimestamp field value
// and a boolean to check if the value has been set.
func (o *Context) GetProcessStartedTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessStartedTimestamp, true
}

// SetProcessStartedTimestamp sets field value
func (o *Context) SetProcessStartedTimestamp(v int64) {
	o.ProcessStartedTimestamp = v
}

// GetStateExecutionId returns the StateExecutionId field value if set, zero value otherwise.
func (o *Context) GetStateExecutionId() string {
	if o == nil || IsNil(o.StateExecutionId) {
		var ret string
		return ret
	}
	return *o.StateExecutionId
}

// GetStateExecutionIdOk returns a tuple with the StateExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetStateExecutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.StateExecutionId) {
		return nil, false
	}
	return o.StateExecutionId, true
}

// HasStateExecutionId returns a boolean if a field has been set.
func (o *Context) HasStateExecutionId() bool {
	if o != nil && !IsNil(o.StateExecutionId) {
		return true
	}

	return false
}

// SetStateExecutionId gets a reference to the given string and assigns it to the StateExecutionId field.
func (o *Context) SetStateExecutionId(v string) {
	o.StateExecutionId = &v
}

// GetFirstAttemptTimestamp returns the FirstAttemptTimestamp field value if set, zero value otherwise.
func (o *Context) GetFirstAttemptTimestamp() int64 {
	if o == nil || IsNil(o.FirstAttemptTimestamp) {
		var ret int64
		return ret
	}
	return *o.FirstAttemptTimestamp
}

// GetFirstAttemptTimestampOk returns a tuple with the FirstAttemptTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetFirstAttemptTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.FirstAttemptTimestamp) {
		return nil, false
	}
	return o.FirstAttemptTimestamp, true
}

// HasFirstAttemptTimestamp returns a boolean if a field has been set.
func (o *Context) HasFirstAttemptTimestamp() bool {
	if o != nil && !IsNil(o.FirstAttemptTimestamp) {
		return true
	}

	return false
}

// SetFirstAttemptTimestamp gets a reference to the given int64 and assigns it to the FirstAttemptTimestamp field.
func (o *Context) SetFirstAttemptTimestamp(v int64) {
	o.FirstAttemptTimestamp = &v
}

// GetAttempt returns the Attempt field value if set, zero value otherwise.
func (o *Context) GetAttempt() int32 {
	if o == nil || IsNil(o.Attempt) {
		var ret int32
		return ret
	}
	return *o.Attempt
}

// GetAttemptOk returns a tuple with the Attempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetAttemptOk() (*int32, bool) {
	if o == nil || IsNil(o.Attempt) {
		return nil, false
	}
	return o.Attempt, true
}

// HasAttempt returns a boolean if a field has been set.
func (o *Context) HasAttempt() bool {
	if o != nil && !IsNil(o.Attempt) {
		return true
	}

	return false
}

// SetAttempt gets a reference to the given int32 and assigns it to the Attempt field.
func (o *Context) SetAttempt(v int32) {
	o.Attempt = &v
}

func (o Context) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Context) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["processId"] = o.ProcessId
	toSerialize["processExecutionId"] = o.ProcessExecutionId
	toSerialize["processStartedTimestamp"] = o.ProcessStartedTimestamp
	if !IsNil(o.StateExecutionId) {
		toSerialize["stateExecutionId"] = o.StateExecutionId
	}
	if !IsNil(o.FirstAttemptTimestamp) {
		toSerialize["firstAttemptTimestamp"] = o.FirstAttemptTimestamp
	}
	if !IsNil(o.Attempt) {
		toSerialize["attempt"] = o.Attempt
	}
	return toSerialize, nil
}

type NullableContext struct {
	value *Context
	isSet bool
}

func (v NullableContext) Get() *Context {
	return v.value
}

func (v *NullableContext) Set(val *Context) {
	v.value = val
	v.isSet = true
}

func (v NullableContext) IsSet() bool {
	return v.isSet
}

func (v *NullableContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContext(val *Context) *NullableContext {
	return &NullableContext{value: val, isSet: true}
}

func (v NullableContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

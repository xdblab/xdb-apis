/*
xCherry APIs

This APIs between xCherry service and SDKs

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xcapi

import (
	"encoding/json"
	"fmt"
)

// checks if the ProcessExecutionWaitForCompletionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessExecutionWaitForCompletionRequest{}

// ProcessExecutionWaitForCompletionRequest the request for waiting for a process completion
type ProcessExecutionWaitForCompletionRequest struct {
	Namespace string `json:"namespace"`
	ProcessId string `json:"processId"`
	// the timeout for the waiting operation
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`
}

type _ProcessExecutionWaitForCompletionRequest ProcessExecutionWaitForCompletionRequest

// NewProcessExecutionWaitForCompletionRequest instantiates a new ProcessExecutionWaitForCompletionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessExecutionWaitForCompletionRequest(namespace string, processId string) *ProcessExecutionWaitForCompletionRequest {
	this := ProcessExecutionWaitForCompletionRequest{}
	this.Namespace = namespace
	this.ProcessId = processId
	return &this
}

// NewProcessExecutionWaitForCompletionRequestWithDefaults instantiates a new ProcessExecutionWaitForCompletionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessExecutionWaitForCompletionRequestWithDefaults() *ProcessExecutionWaitForCompletionRequest {
	this := ProcessExecutionWaitForCompletionRequest{}
	return &this
}

// GetNamespace returns the Namespace field value
func (o *ProcessExecutionWaitForCompletionRequest) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *ProcessExecutionWaitForCompletionRequest) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *ProcessExecutionWaitForCompletionRequest) SetNamespace(v string) {
	o.Namespace = v
}

// GetProcessId returns the ProcessId field value
func (o *ProcessExecutionWaitForCompletionRequest) GetProcessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value
// and a boolean to check if the value has been set.
func (o *ProcessExecutionWaitForCompletionRequest) GetProcessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessId, true
}

// SetProcessId sets field value
func (o *ProcessExecutionWaitForCompletionRequest) SetProcessId(v string) {
	o.ProcessId = v
}

// GetTimeoutSeconds returns the TimeoutSeconds field value if set, zero value otherwise.
func (o *ProcessExecutionWaitForCompletionRequest) GetTimeoutSeconds() int32 {
	if o == nil || IsNil(o.TimeoutSeconds) {
		var ret int32
		return ret
	}
	return *o.TimeoutSeconds
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessExecutionWaitForCompletionRequest) GetTimeoutSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeoutSeconds) {
		return nil, false
	}
	return o.TimeoutSeconds, true
}

// HasTimeoutSeconds returns a boolean if a field has been set.
func (o *ProcessExecutionWaitForCompletionRequest) HasTimeoutSeconds() bool {
	if o != nil && !IsNil(o.TimeoutSeconds) {
		return true
	}

	return false
}

// SetTimeoutSeconds gets a reference to the given int32 and assigns it to the TimeoutSeconds field.
func (o *ProcessExecutionWaitForCompletionRequest) SetTimeoutSeconds(v int32) {
	o.TimeoutSeconds = &v
}

func (o ProcessExecutionWaitForCompletionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessExecutionWaitForCompletionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["namespace"] = o.Namespace
	toSerialize["processId"] = o.ProcessId
	if !IsNil(o.TimeoutSeconds) {
		toSerialize["timeoutSeconds"] = o.TimeoutSeconds
	}
	return toSerialize, nil
}

func (o *ProcessExecutionWaitForCompletionRequest) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"namespace",
		"processId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProcessExecutionWaitForCompletionRequest := _ProcessExecutionWaitForCompletionRequest{}

	err = json.Unmarshal(bytes, &varProcessExecutionWaitForCompletionRequest)

	if err != nil {
		return err
	}

	*o = ProcessExecutionWaitForCompletionRequest(varProcessExecutionWaitForCompletionRequest)

	return err
}

type NullableProcessExecutionWaitForCompletionRequest struct {
	value *ProcessExecutionWaitForCompletionRequest
	isSet bool
}

func (v NullableProcessExecutionWaitForCompletionRequest) Get() *ProcessExecutionWaitForCompletionRequest {
	return v.value
}

func (v *NullableProcessExecutionWaitForCompletionRequest) Set(val *ProcessExecutionWaitForCompletionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessExecutionWaitForCompletionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessExecutionWaitForCompletionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessExecutionWaitForCompletionRequest(val *ProcessExecutionWaitForCompletionRequest) *NullableProcessExecutionWaitForCompletionRequest {
	return &NullableProcessExecutionWaitForCompletionRequest{value: val, isSet: true}
}

func (v NullableProcessExecutionWaitForCompletionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessExecutionWaitForCompletionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

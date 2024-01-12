/*
xCherry APIs

This APIs between xCherry service and SDKs

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xcapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ProcessExecutionDescribeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessExecutionDescribeRequest{}

// ProcessExecutionDescribeRequest struct for ProcessExecutionDescribeRequest
type ProcessExecutionDescribeRequest struct {
	Namespace string `json:"namespace"`
	ProcessId string `json:"processId"`
}

type _ProcessExecutionDescribeRequest ProcessExecutionDescribeRequest

// NewProcessExecutionDescribeRequest instantiates a new ProcessExecutionDescribeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessExecutionDescribeRequest(namespace string, processId string) *ProcessExecutionDescribeRequest {
	this := ProcessExecutionDescribeRequest{}
	this.Namespace = namespace
	this.ProcessId = processId
	return &this
}

// NewProcessExecutionDescribeRequestWithDefaults instantiates a new ProcessExecutionDescribeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessExecutionDescribeRequestWithDefaults() *ProcessExecutionDescribeRequest {
	this := ProcessExecutionDescribeRequest{}
	return &this
}

// GetNamespace returns the Namespace field value
func (o *ProcessExecutionDescribeRequest) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *ProcessExecutionDescribeRequest) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *ProcessExecutionDescribeRequest) SetNamespace(v string) {
	o.Namespace = v
}

// GetProcessId returns the ProcessId field value
func (o *ProcessExecutionDescribeRequest) GetProcessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value
// and a boolean to check if the value has been set.
func (o *ProcessExecutionDescribeRequest) GetProcessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessId, true
}

// SetProcessId sets field value
func (o *ProcessExecutionDescribeRequest) SetProcessId(v string) {
	o.ProcessId = v
}

func (o ProcessExecutionDescribeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessExecutionDescribeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["namespace"] = o.Namespace
	toSerialize["processId"] = o.ProcessId
	return toSerialize, nil
}

func (o *ProcessExecutionDescribeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"namespace",
		"processId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProcessExecutionDescribeRequest := _ProcessExecutionDescribeRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProcessExecutionDescribeRequest)

	if err != nil {
		return err
	}

	*o = ProcessExecutionDescribeRequest(varProcessExecutionDescribeRequest)

	return err
}

type NullableProcessExecutionDescribeRequest struct {
	value *ProcessExecutionDescribeRequest
	isSet bool
}

func (v NullableProcessExecutionDescribeRequest) Get() *ProcessExecutionDescribeRequest {
	return v.value
}

func (v *NullableProcessExecutionDescribeRequest) Set(val *ProcessExecutionDescribeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessExecutionDescribeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessExecutionDescribeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessExecutionDescribeRequest(val *ProcessExecutionDescribeRequest) *NullableProcessExecutionDescribeRequest {
	return &NullableProcessExecutionDescribeRequest{value: val, isSet: true}
}

func (v NullableProcessExecutionDescribeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessExecutionDescribeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

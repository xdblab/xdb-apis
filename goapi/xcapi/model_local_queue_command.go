/*
xCherry APIs

This APIs between xCherry service and SDKs

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xcapi

import (
	"encoding/json"
)

// checks if the LocalQueueCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalQueueCommand{}

// LocalQueueCommand struct for LocalQueueCommand
type LocalQueueCommand struct {
	QueueName string `json:"queueName"`
	// the number of identical messages to await
	Count *int32 `json:"count,omitempty"`
}

// NewLocalQueueCommand instantiates a new LocalQueueCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalQueueCommand(queueName string) *LocalQueueCommand {
	this := LocalQueueCommand{}
	this.QueueName = queueName
	var count int32 = 1
	this.Count = &count
	return &this
}

// NewLocalQueueCommandWithDefaults instantiates a new LocalQueueCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalQueueCommandWithDefaults() *LocalQueueCommand {
	this := LocalQueueCommand{}
	var count int32 = 1
	this.Count = &count
	return &this
}

// GetQueueName returns the QueueName field value
func (o *LocalQueueCommand) GetQueueName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueueName
}

// GetQueueNameOk returns a tuple with the QueueName field value
// and a boolean to check if the value has been set.
func (o *LocalQueueCommand) GetQueueNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueueName, true
}

// SetQueueName sets field value
func (o *LocalQueueCommand) SetQueueName(v string) {
	o.QueueName = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *LocalQueueCommand) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalQueueCommand) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *LocalQueueCommand) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *LocalQueueCommand) SetCount(v int32) {
	o.Count = &v
}

func (o LocalQueueCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalQueueCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queueName"] = o.QueueName
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableLocalQueueCommand struct {
	value *LocalQueueCommand
	isSet bool
}

func (v NullableLocalQueueCommand) Get() *LocalQueueCommand {
	return v.value
}

func (v *NullableLocalQueueCommand) Set(val *LocalQueueCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalQueueCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalQueueCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalQueueCommand(val *LocalQueueCommand) *NullableLocalQueueCommand {
	return &NullableLocalQueueCommand{value: val, isSet: true}
}

func (v NullableLocalQueueCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalQueueCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

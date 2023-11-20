/*
xCherry APIs

This APIs between xCherry service and SDKs

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xcapi

import (
	"encoding/json"
	"fmt"
)

// checks if the TimerCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimerCommand{}

// TimerCommand struct for TimerCommand
type TimerCommand struct {
	DelayInSeconds int64 `json:"delayInSeconds"`
}

type _TimerCommand TimerCommand

// NewTimerCommand instantiates a new TimerCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimerCommand(delayInSeconds int64) *TimerCommand {
	this := TimerCommand{}
	this.DelayInSeconds = delayInSeconds
	return &this
}

// NewTimerCommandWithDefaults instantiates a new TimerCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimerCommandWithDefaults() *TimerCommand {
	this := TimerCommand{}
	return &this
}

// GetDelayInSeconds returns the DelayInSeconds field value
func (o *TimerCommand) GetDelayInSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DelayInSeconds
}

// GetDelayInSecondsOk returns a tuple with the DelayInSeconds field value
// and a boolean to check if the value has been set.
func (o *TimerCommand) GetDelayInSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DelayInSeconds, true
}

// SetDelayInSeconds sets field value
func (o *TimerCommand) SetDelayInSeconds(v int64) {
	o.DelayInSeconds = v
}

func (o TimerCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimerCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["delayInSeconds"] = o.DelayInSeconds
	return toSerialize, nil
}

func (o *TimerCommand) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"delayInSeconds",
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

	varTimerCommand := _TimerCommand{}

	err = json.Unmarshal(bytes, &varTimerCommand)

	if err != nil {
		return err
	}

	*o = TimerCommand(varTimerCommand)

	return err
}

type NullableTimerCommand struct {
	value *TimerCommand
	isSet bool
}

func (v NullableTimerCommand) Get() *TimerCommand {
	return v.value
}

func (v *NullableTimerCommand) Set(val *TimerCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableTimerCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableTimerCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimerCommand(val *TimerCommand) *NullableTimerCommand {
	return &NullableTimerCommand{value: val, isSet: true}
}

func (v NullableTimerCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimerCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

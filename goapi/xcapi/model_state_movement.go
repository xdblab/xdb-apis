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

// checks if the StateMovement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StateMovement{}

// StateMovement struct for StateMovement
type StateMovement struct {
	StateId     string            `json:"stateId"`
	StateInput  *EncodedObject    `json:"stateInput,omitempty"`
	StateConfig *AsyncStateConfig `json:"stateConfig,omitempty"`
}

type _StateMovement StateMovement

// NewStateMovement instantiates a new StateMovement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateMovement(stateId string) *StateMovement {
	this := StateMovement{}
	this.StateId = stateId
	return &this
}

// NewStateMovementWithDefaults instantiates a new StateMovement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateMovementWithDefaults() *StateMovement {
	this := StateMovement{}
	return &this
}

// GetStateId returns the StateId field value
func (o *StateMovement) GetStateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateId
}

// GetStateIdOk returns a tuple with the StateId field value
// and a boolean to check if the value has been set.
func (o *StateMovement) GetStateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateId, true
}

// SetStateId sets field value
func (o *StateMovement) SetStateId(v string) {
	o.StateId = v
}

// GetStateInput returns the StateInput field value if set, zero value otherwise.
func (o *StateMovement) GetStateInput() EncodedObject {
	if o == nil || IsNil(o.StateInput) {
		var ret EncodedObject
		return ret
	}
	return *o.StateInput
}

// GetStateInputOk returns a tuple with the StateInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateMovement) GetStateInputOk() (*EncodedObject, bool) {
	if o == nil || IsNil(o.StateInput) {
		return nil, false
	}
	return o.StateInput, true
}

// HasStateInput returns a boolean if a field has been set.
func (o *StateMovement) HasStateInput() bool {
	if o != nil && !IsNil(o.StateInput) {
		return true
	}

	return false
}

// SetStateInput gets a reference to the given EncodedObject and assigns it to the StateInput field.
func (o *StateMovement) SetStateInput(v EncodedObject) {
	o.StateInput = &v
}

// GetStateConfig returns the StateConfig field value if set, zero value otherwise.
func (o *StateMovement) GetStateConfig() AsyncStateConfig {
	if o == nil || IsNil(o.StateConfig) {
		var ret AsyncStateConfig
		return ret
	}
	return *o.StateConfig
}

// GetStateConfigOk returns a tuple with the StateConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateMovement) GetStateConfigOk() (*AsyncStateConfig, bool) {
	if o == nil || IsNil(o.StateConfig) {
		return nil, false
	}
	return o.StateConfig, true
}

// HasStateConfig returns a boolean if a field has been set.
func (o *StateMovement) HasStateConfig() bool {
	if o != nil && !IsNil(o.StateConfig) {
		return true
	}

	return false
}

// SetStateConfig gets a reference to the given AsyncStateConfig and assigns it to the StateConfig field.
func (o *StateMovement) SetStateConfig(v AsyncStateConfig) {
	o.StateConfig = &v
}

func (o StateMovement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StateMovement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stateId"] = o.StateId
	if !IsNil(o.StateInput) {
		toSerialize["stateInput"] = o.StateInput
	}
	if !IsNil(o.StateConfig) {
		toSerialize["stateConfig"] = o.StateConfig
	}
	return toSerialize, nil
}

func (o *StateMovement) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"stateId",
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

	varStateMovement := _StateMovement{}

	err = json.Unmarshal(bytes, &varStateMovement)

	if err != nil {
		return err
	}

	*o = StateMovement(varStateMovement)

	return err
}

type NullableStateMovement struct {
	value *StateMovement
	isSet bool
}

func (v NullableStateMovement) Get() *StateMovement {
	return v.value
}

func (v *NullableStateMovement) Set(val *StateMovement) {
	v.value = val
	v.isSet = true
}

func (v NullableStateMovement) IsSet() bool {
	return v.isSet
}

func (v *NullableStateMovement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateMovement(val *StateMovement) *NullableStateMovement {
	return &NullableStateMovement{value: val, isSet: true}
}

func (v NullableStateMovement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateMovement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

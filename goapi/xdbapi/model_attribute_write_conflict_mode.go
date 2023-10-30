/*
XDB APIs

This APIs between xdb service and SDKs

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xdbapi

import (
	"encoding/json"
	"fmt"
)

// AttributeWriteConflictMode the model 'AttributeWriteConflictMode'
type AttributeWriteConflictMode string

// List of AttributeWriteConflictMode
const (
	RETURN_ERROR_ON_CONFLICT AttributeWriteConflictMode = "ReturnErrorOnConflict"
	IGNORE_CONFLICT          AttributeWriteConflictMode = "IgnoreConflict"
	OVERRIDE_ON_CONFLICT     AttributeWriteConflictMode = "OverrideOnConflict"
)

// All allowed values of AttributeWriteConflictMode enum
var AllowedAttributeWriteConflictModeEnumValues = []AttributeWriteConflictMode{
	"ReturnErrorOnConflict",
	"IgnoreConflict",
	"OverrideOnConflict",
}

func (v *AttributeWriteConflictMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AttributeWriteConflictMode(value)
	for _, existing := range AllowedAttributeWriteConflictModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AttributeWriteConflictMode", value)
}

// NewAttributeWriteConflictModeFromValue returns a pointer to a valid AttributeWriteConflictMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAttributeWriteConflictModeFromValue(v string) (*AttributeWriteConflictMode, error) {
	ev := AttributeWriteConflictMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AttributeWriteConflictMode: valid values are %v", v, AllowedAttributeWriteConflictModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AttributeWriteConflictMode) IsValid() bool {
	for _, existing := range AllowedAttributeWriteConflictModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AttributeWriteConflictMode value
func (v AttributeWriteConflictMode) Ptr() *AttributeWriteConflictMode {
	return &v
}

type NullableAttributeWriteConflictMode struct {
	value *AttributeWriteConflictMode
	isSet bool
}

func (v NullableAttributeWriteConflictMode) Get() *AttributeWriteConflictMode {
	return v.value
}

func (v *NullableAttributeWriteConflictMode) Set(val *AttributeWriteConflictMode) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeWriteConflictMode) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeWriteConflictMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeWriteConflictMode(val *AttributeWriteConflictMode) *NullableAttributeWriteConflictMode {
	return &NullableAttributeWriteConflictMode{value: val, isSet: true}
}

func (v NullableAttributeWriteConflictMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeWriteConflictMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

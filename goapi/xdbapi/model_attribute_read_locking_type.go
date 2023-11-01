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

// AttributeReadLockingType the model 'AttributeReadLockingType'
type AttributeReadLockingType string

// List of AttributeReadLockingType
const (
	NO_LOCKING     AttributeReadLockingType = "NO_LOCKING"
	SHARE_LOCK     AttributeReadLockingType = "SHARE_LOCK"
	EXCLUSIVE_LOCK AttributeReadLockingType = "EXCLUSIVE_LOCK"
)

// All allowed values of AttributeReadLockingType enum
var AllowedAttributeReadLockingTypeEnumValues = []AttributeReadLockingType{
	"NO_LOCKING",
	"SHARE_LOCK",
	"EXCLUSIVE_LOCK",
}

func (v *AttributeReadLockingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AttributeReadLockingType(value)
	for _, existing := range AllowedAttributeReadLockingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AttributeReadLockingType", value)
}

// NewAttributeReadLockingTypeFromValue returns a pointer to a valid AttributeReadLockingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAttributeReadLockingTypeFromValue(v string) (*AttributeReadLockingType, error) {
	ev := AttributeReadLockingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AttributeReadLockingType: valid values are %v", v, AllowedAttributeReadLockingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AttributeReadLockingType) IsValid() bool {
	for _, existing := range AllowedAttributeReadLockingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AttributeReadLockingType value
func (v AttributeReadLockingType) Ptr() *AttributeReadLockingType {
	return &v
}

type NullableAttributeReadLockingType struct {
	value *AttributeReadLockingType
	isSet bool
}

func (v NullableAttributeReadLockingType) Get() *AttributeReadLockingType {
	return v.value
}

func (v *NullableAttributeReadLockingType) Set(val *AttributeReadLockingType) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeReadLockingType) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeReadLockingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeReadLockingType(val *AttributeReadLockingType) *NullableAttributeReadLockingType {
	return &NullableAttributeReadLockingType{value: val, isSet: true}
}

func (v NullableAttributeReadLockingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeReadLockingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
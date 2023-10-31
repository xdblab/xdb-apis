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

// LocalQueueStatus the model 'LocalQueueStatus'
type LocalQueueStatus string

// List of LocalQueueStatus
const (
	WAITING_LOCAL_QUEUE   LocalQueueStatus = "WAITING_LOCAL_QUEUE"
	COMPLETED_LOCAL_QUEUE LocalQueueStatus = "COMPLETED_LOCAL_QUEUE"
	SKIPPED_LOCAL_QUEUE   LocalQueueStatus = "SKIPPED_LOCAL_QUEUE"
)

// All allowed values of LocalQueueStatus enum
var AllowedLocalQueueStatusEnumValues = []LocalQueueStatus{
	"WAITING_LOCAL_QUEUE",
	"COMPLETED_LOCAL_QUEUE",
	"SKIPPED_LOCAL_QUEUE",
}

func (v *LocalQueueStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LocalQueueStatus(value)
	for _, existing := range AllowedLocalQueueStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LocalQueueStatus", value)
}

// NewLocalQueueStatusFromValue returns a pointer to a valid LocalQueueStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocalQueueStatusFromValue(v string) (*LocalQueueStatus, error) {
	ev := LocalQueueStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocalQueueStatus: valid values are %v", v, AllowedLocalQueueStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocalQueueStatus) IsValid() bool {
	for _, existing := range AllowedLocalQueueStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LocalQueueStatus value
func (v LocalQueueStatus) Ptr() *LocalQueueStatus {
	return &v
}

type NullableLocalQueueStatus struct {
	value *LocalQueueStatus
	isSet bool
}

func (v NullableLocalQueueStatus) Get() *LocalQueueStatus {
	return v.value
}

func (v *NullableLocalQueueStatus) Set(val *LocalQueueStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalQueueStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalQueueStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalQueueStatus(val *LocalQueueStatus) *NullableLocalQueueStatus {
	return &NullableLocalQueueStatus{value: val, isSet: true}
}

func (v NullableLocalQueueStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalQueueStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

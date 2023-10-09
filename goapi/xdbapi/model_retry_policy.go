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

// checks if the RetryPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetryPolicy{}

// RetryPolicy struct for RetryPolicy
type RetryPolicy struct {
	// the initial interval for the first retry, default to 1 second
	InitialIntervalSeconds *int32 `json:"initialIntervalSeconds,omitempty"`
	// the backoff coefficient for the next retry, default to 2
	BackoffCoefficient *float32 `json:"backoffCoefficient,omitempty"`
	// the maximum interval for the next retry, default to 100x of initial interval
	MaximumIntervalSeconds *int32 `json:"maximumIntervalSeconds,omitempty"`
	// the maximum number of attempts, default to 0, means unlimited
	MaximumAttempts *int32 `json:"maximumAttempts,omitempty"`
	// the maximum duration of all attempts, default to 0, means unlimited
	MaximumAttemptsDurationSeconds *int32 `json:"maximumAttemptsDurationSeconds,omitempty"`
}

// NewRetryPolicy instantiates a new RetryPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetryPolicy() *RetryPolicy {
	this := RetryPolicy{}
	return &this
}

// NewRetryPolicyWithDefaults instantiates a new RetryPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetryPolicyWithDefaults() *RetryPolicy {
	this := RetryPolicy{}
	return &this
}

// GetInitialIntervalSeconds returns the InitialIntervalSeconds field value if set, zero value otherwise.
func (o *RetryPolicy) GetInitialIntervalSeconds() int32 {
	if o == nil || IsNil(o.InitialIntervalSeconds) {
		var ret int32
		return ret
	}
	return *o.InitialIntervalSeconds
}

// GetInitialIntervalSecondsOk returns a tuple with the InitialIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetryPolicy) GetInitialIntervalSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.InitialIntervalSeconds) {
		return nil, false
	}
	return o.InitialIntervalSeconds, true
}

// HasInitialIntervalSeconds returns a boolean if a field has been set.
func (o *RetryPolicy) HasInitialIntervalSeconds() bool {
	if o != nil && !IsNil(o.InitialIntervalSeconds) {
		return true
	}

	return false
}

// SetInitialIntervalSeconds gets a reference to the given int32 and assigns it to the InitialIntervalSeconds field.
func (o *RetryPolicy) SetInitialIntervalSeconds(v int32) {
	o.InitialIntervalSeconds = &v
}

// GetBackoffCoefficient returns the BackoffCoefficient field value if set, zero value otherwise.
func (o *RetryPolicy) GetBackoffCoefficient() float32 {
	if o == nil || IsNil(o.BackoffCoefficient) {
		var ret float32
		return ret
	}
	return *o.BackoffCoefficient
}

// GetBackoffCoefficientOk returns a tuple with the BackoffCoefficient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetryPolicy) GetBackoffCoefficientOk() (*float32, bool) {
	if o == nil || IsNil(o.BackoffCoefficient) {
		return nil, false
	}
	return o.BackoffCoefficient, true
}

// HasBackoffCoefficient returns a boolean if a field has been set.
func (o *RetryPolicy) HasBackoffCoefficient() bool {
	if o != nil && !IsNil(o.BackoffCoefficient) {
		return true
	}

	return false
}

// SetBackoffCoefficient gets a reference to the given float32 and assigns it to the BackoffCoefficient field.
func (o *RetryPolicy) SetBackoffCoefficient(v float32) {
	o.BackoffCoefficient = &v
}

// GetMaximumIntervalSeconds returns the MaximumIntervalSeconds field value if set, zero value otherwise.
func (o *RetryPolicy) GetMaximumIntervalSeconds() int32 {
	if o == nil || IsNil(o.MaximumIntervalSeconds) {
		var ret int32
		return ret
	}
	return *o.MaximumIntervalSeconds
}

// GetMaximumIntervalSecondsOk returns a tuple with the MaximumIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetryPolicy) GetMaximumIntervalSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumIntervalSeconds) {
		return nil, false
	}
	return o.MaximumIntervalSeconds, true
}

// HasMaximumIntervalSeconds returns a boolean if a field has been set.
func (o *RetryPolicy) HasMaximumIntervalSeconds() bool {
	if o != nil && !IsNil(o.MaximumIntervalSeconds) {
		return true
	}

	return false
}

// SetMaximumIntervalSeconds gets a reference to the given int32 and assigns it to the MaximumIntervalSeconds field.
func (o *RetryPolicy) SetMaximumIntervalSeconds(v int32) {
	o.MaximumIntervalSeconds = &v
}

// GetMaximumAttempts returns the MaximumAttempts field value if set, zero value otherwise.
func (o *RetryPolicy) GetMaximumAttempts() int32 {
	if o == nil || IsNil(o.MaximumAttempts) {
		var ret int32
		return ret
	}
	return *o.MaximumAttempts
}

// GetMaximumAttemptsOk returns a tuple with the MaximumAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetryPolicy) GetMaximumAttemptsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumAttempts) {
		return nil, false
	}
	return o.MaximumAttempts, true
}

// HasMaximumAttempts returns a boolean if a field has been set.
func (o *RetryPolicy) HasMaximumAttempts() bool {
	if o != nil && !IsNil(o.MaximumAttempts) {
		return true
	}

	return false
}

// SetMaximumAttempts gets a reference to the given int32 and assigns it to the MaximumAttempts field.
func (o *RetryPolicy) SetMaximumAttempts(v int32) {
	o.MaximumAttempts = &v
}

// GetMaximumAttemptsDurationSeconds returns the MaximumAttemptsDurationSeconds field value if set, zero value otherwise.
func (o *RetryPolicy) GetMaximumAttemptsDurationSeconds() int32 {
	if o == nil || IsNil(o.MaximumAttemptsDurationSeconds) {
		var ret int32
		return ret
	}
	return *o.MaximumAttemptsDurationSeconds
}

// GetMaximumAttemptsDurationSecondsOk returns a tuple with the MaximumAttemptsDurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetryPolicy) GetMaximumAttemptsDurationSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumAttemptsDurationSeconds) {
		return nil, false
	}
	return o.MaximumAttemptsDurationSeconds, true
}

// HasMaximumAttemptsDurationSeconds returns a boolean if a field has been set.
func (o *RetryPolicy) HasMaximumAttemptsDurationSeconds() bool {
	if o != nil && !IsNil(o.MaximumAttemptsDurationSeconds) {
		return true
	}

	return false
}

// SetMaximumAttemptsDurationSeconds gets a reference to the given int32 and assigns it to the MaximumAttemptsDurationSeconds field.
func (o *RetryPolicy) SetMaximumAttemptsDurationSeconds(v int32) {
	o.MaximumAttemptsDurationSeconds = &v
}

func (o RetryPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetryPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InitialIntervalSeconds) {
		toSerialize["initialIntervalSeconds"] = o.InitialIntervalSeconds
	}
	if !IsNil(o.BackoffCoefficient) {
		toSerialize["backoffCoefficient"] = o.BackoffCoefficient
	}
	if !IsNil(o.MaximumIntervalSeconds) {
		toSerialize["maximumIntervalSeconds"] = o.MaximumIntervalSeconds
	}
	if !IsNil(o.MaximumAttempts) {
		toSerialize["maximumAttempts"] = o.MaximumAttempts
	}
	if !IsNil(o.MaximumAttemptsDurationSeconds) {
		toSerialize["maximumAttemptsDurationSeconds"] = o.MaximumAttemptsDurationSeconds
	}
	return toSerialize, nil
}

type NullableRetryPolicy struct {
	value *RetryPolicy
	isSet bool
}

func (v NullableRetryPolicy) Get() *RetryPolicy {
	return v.value
}

func (v *NullableRetryPolicy) Set(val *RetryPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableRetryPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableRetryPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetryPolicy(val *RetryPolicy) *NullableRetryPolicy {
	return &NullableRetryPolicy{value: val, isSet: true}
}

func (v NullableRetryPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetryPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

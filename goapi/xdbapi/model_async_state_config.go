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

// checks if the AsyncStateConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsyncStateConfig{}

// AsyncStateConfig struct for AsyncStateConfig
type AsyncStateConfig struct {
	SkipWaitUntil *bool `json:"skipWaitUntil,omitempty"`
	// the timeout for the single attempt of AsyncState.waitUntil API
	WaitUntilApiTimeoutSeconds *int32 `json:"waitUntilApiTimeoutSeconds,omitempty"`
	// the timeout for the single attempt of AsyncState.execute API
	ExecuteApiTimeoutSeconds    *int32                       `json:"executeApiTimeoutSeconds,omitempty"`
	WaitUntilApiRetryPolicy     *RetryPolicy                 `json:"waitUntilApiRetryPolicy,omitempty"`
	ExecuteApiRetryPolicy       *RetryPolicy                 `json:"executeApiRetryPolicy,omitempty"`
	StateFailureRecoveryOptions *StateFailureRecoveryOptions `json:"stateFailureRecoveryOptions,omitempty"`
	LoadGlobalAttributesRequest *LoadGlobalAttributesRequest `json:"loadGlobalAttributesRequest,omitempty"`
}

// NewAsyncStateConfig instantiates a new AsyncStateConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsyncStateConfig() *AsyncStateConfig {
	this := AsyncStateConfig{}
	return &this
}

// NewAsyncStateConfigWithDefaults instantiates a new AsyncStateConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsyncStateConfigWithDefaults() *AsyncStateConfig {
	this := AsyncStateConfig{}
	return &this
}

// GetSkipWaitUntil returns the SkipWaitUntil field value if set, zero value otherwise.
func (o *AsyncStateConfig) GetSkipWaitUntil() bool {
	if o == nil || IsNil(o.SkipWaitUntil) {
		var ret bool
		return ret
	}
	return *o.SkipWaitUntil
}

// GetSkipWaitUntilOk returns a tuple with the SkipWaitUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncStateConfig) GetSkipWaitUntilOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipWaitUntil) {
		return nil, false
	}
	return o.SkipWaitUntil, true
}

// HasSkipWaitUntil returns a boolean if a field has been set.
func (o *AsyncStateConfig) HasSkipWaitUntil() bool {
	if o != nil && !IsNil(o.SkipWaitUntil) {
		return true
	}

	return false
}

// SetSkipWaitUntil gets a reference to the given bool and assigns it to the SkipWaitUntil field.
func (o *AsyncStateConfig) SetSkipWaitUntil(v bool) {
	o.SkipWaitUntil = &v
}

// GetWaitUntilApiTimeoutSeconds returns the WaitUntilApiTimeoutSeconds field value if set, zero value otherwise.
func (o *AsyncStateConfig) GetWaitUntilApiTimeoutSeconds() int32 {
	if o == nil || IsNil(o.WaitUntilApiTimeoutSeconds) {
		var ret int32
		return ret
	}
	return *o.WaitUntilApiTimeoutSeconds
}

// GetWaitUntilApiTimeoutSecondsOk returns a tuple with the WaitUntilApiTimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncStateConfig) GetWaitUntilApiTimeoutSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitUntilApiTimeoutSeconds) {
		return nil, false
	}
	return o.WaitUntilApiTimeoutSeconds, true
}

// HasWaitUntilApiTimeoutSeconds returns a boolean if a field has been set.
func (o *AsyncStateConfig) HasWaitUntilApiTimeoutSeconds() bool {
	if o != nil && !IsNil(o.WaitUntilApiTimeoutSeconds) {
		return true
	}

	return false
}

// SetWaitUntilApiTimeoutSeconds gets a reference to the given int32 and assigns it to the WaitUntilApiTimeoutSeconds field.
func (o *AsyncStateConfig) SetWaitUntilApiTimeoutSeconds(v int32) {
	o.WaitUntilApiTimeoutSeconds = &v
}

// GetExecuteApiTimeoutSeconds returns the ExecuteApiTimeoutSeconds field value if set, zero value otherwise.
func (o *AsyncStateConfig) GetExecuteApiTimeoutSeconds() int32 {
	if o == nil || IsNil(o.ExecuteApiTimeoutSeconds) {
		var ret int32
		return ret
	}
	return *o.ExecuteApiTimeoutSeconds
}

// GetExecuteApiTimeoutSecondsOk returns a tuple with the ExecuteApiTimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncStateConfig) GetExecuteApiTimeoutSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.ExecuteApiTimeoutSeconds) {
		return nil, false
	}
	return o.ExecuteApiTimeoutSeconds, true
}

// HasExecuteApiTimeoutSeconds returns a boolean if a field has been set.
func (o *AsyncStateConfig) HasExecuteApiTimeoutSeconds() bool {
	if o != nil && !IsNil(o.ExecuteApiTimeoutSeconds) {
		return true
	}

	return false
}

// SetExecuteApiTimeoutSeconds gets a reference to the given int32 and assigns it to the ExecuteApiTimeoutSeconds field.
func (o *AsyncStateConfig) SetExecuteApiTimeoutSeconds(v int32) {
	o.ExecuteApiTimeoutSeconds = &v
}

// GetWaitUntilApiRetryPolicy returns the WaitUntilApiRetryPolicy field value if set, zero value otherwise.
func (o *AsyncStateConfig) GetWaitUntilApiRetryPolicy() RetryPolicy {
	if o == nil || IsNil(o.WaitUntilApiRetryPolicy) {
		var ret RetryPolicy
		return ret
	}
	return *o.WaitUntilApiRetryPolicy
}

// GetWaitUntilApiRetryPolicyOk returns a tuple with the WaitUntilApiRetryPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncStateConfig) GetWaitUntilApiRetryPolicyOk() (*RetryPolicy, bool) {
	if o == nil || IsNil(o.WaitUntilApiRetryPolicy) {
		return nil, false
	}
	return o.WaitUntilApiRetryPolicy, true
}

// HasWaitUntilApiRetryPolicy returns a boolean if a field has been set.
func (o *AsyncStateConfig) HasWaitUntilApiRetryPolicy() bool {
	if o != nil && !IsNil(o.WaitUntilApiRetryPolicy) {
		return true
	}

	return false
}

// SetWaitUntilApiRetryPolicy gets a reference to the given RetryPolicy and assigns it to the WaitUntilApiRetryPolicy field.
func (o *AsyncStateConfig) SetWaitUntilApiRetryPolicy(v RetryPolicy) {
	o.WaitUntilApiRetryPolicy = &v
}

// GetExecuteApiRetryPolicy returns the ExecuteApiRetryPolicy field value if set, zero value otherwise.
func (o *AsyncStateConfig) GetExecuteApiRetryPolicy() RetryPolicy {
	if o == nil || IsNil(o.ExecuteApiRetryPolicy) {
		var ret RetryPolicy
		return ret
	}
	return *o.ExecuteApiRetryPolicy
}

// GetExecuteApiRetryPolicyOk returns a tuple with the ExecuteApiRetryPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncStateConfig) GetExecuteApiRetryPolicyOk() (*RetryPolicy, bool) {
	if o == nil || IsNil(o.ExecuteApiRetryPolicy) {
		return nil, false
	}
	return o.ExecuteApiRetryPolicy, true
}

// HasExecuteApiRetryPolicy returns a boolean if a field has been set.
func (o *AsyncStateConfig) HasExecuteApiRetryPolicy() bool {
	if o != nil && !IsNil(o.ExecuteApiRetryPolicy) {
		return true
	}

	return false
}

// SetExecuteApiRetryPolicy gets a reference to the given RetryPolicy and assigns it to the ExecuteApiRetryPolicy field.
func (o *AsyncStateConfig) SetExecuteApiRetryPolicy(v RetryPolicy) {
	o.ExecuteApiRetryPolicy = &v
}

// GetStateFailureRecoveryOptions returns the StateFailureRecoveryOptions field value if set, zero value otherwise.
func (o *AsyncStateConfig) GetStateFailureRecoveryOptions() StateFailureRecoveryOptions {
	if o == nil || IsNil(o.StateFailureRecoveryOptions) {
		var ret StateFailureRecoveryOptions
		return ret
	}
	return *o.StateFailureRecoveryOptions
}

// GetStateFailureRecoveryOptionsOk returns a tuple with the StateFailureRecoveryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncStateConfig) GetStateFailureRecoveryOptionsOk() (*StateFailureRecoveryOptions, bool) {
	if o == nil || IsNil(o.StateFailureRecoveryOptions) {
		return nil, false
	}
	return o.StateFailureRecoveryOptions, true
}

// HasStateFailureRecoveryOptions returns a boolean if a field has been set.
func (o *AsyncStateConfig) HasStateFailureRecoveryOptions() bool {
	if o != nil && !IsNil(o.StateFailureRecoveryOptions) {
		return true
	}

	return false
}

// SetStateFailureRecoveryOptions gets a reference to the given StateFailureRecoveryOptions and assigns it to the StateFailureRecoveryOptions field.
func (o *AsyncStateConfig) SetStateFailureRecoveryOptions(v StateFailureRecoveryOptions) {
	o.StateFailureRecoveryOptions = &v
}

// GetLoadGlobalAttributesRequest returns the LoadGlobalAttributesRequest field value if set, zero value otherwise.
func (o *AsyncStateConfig) GetLoadGlobalAttributesRequest() LoadGlobalAttributesRequest {
	if o == nil || IsNil(o.LoadGlobalAttributesRequest) {
		var ret LoadGlobalAttributesRequest
		return ret
	}
	return *o.LoadGlobalAttributesRequest
}

// GetLoadGlobalAttributesRequestOk returns a tuple with the LoadGlobalAttributesRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncStateConfig) GetLoadGlobalAttributesRequestOk() (*LoadGlobalAttributesRequest, bool) {
	if o == nil || IsNil(o.LoadGlobalAttributesRequest) {
		return nil, false
	}
	return o.LoadGlobalAttributesRequest, true
}

// HasLoadGlobalAttributesRequest returns a boolean if a field has been set.
func (o *AsyncStateConfig) HasLoadGlobalAttributesRequest() bool {
	if o != nil && !IsNil(o.LoadGlobalAttributesRequest) {
		return true
	}

	return false
}

// SetLoadGlobalAttributesRequest gets a reference to the given LoadGlobalAttributesRequest and assigns it to the LoadGlobalAttributesRequest field.
func (o *AsyncStateConfig) SetLoadGlobalAttributesRequest(v LoadGlobalAttributesRequest) {
	o.LoadGlobalAttributesRequest = &v
}

func (o AsyncStateConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AsyncStateConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SkipWaitUntil) {
		toSerialize["skipWaitUntil"] = o.SkipWaitUntil
	}
	if !IsNil(o.WaitUntilApiTimeoutSeconds) {
		toSerialize["waitUntilApiTimeoutSeconds"] = o.WaitUntilApiTimeoutSeconds
	}
	if !IsNil(o.ExecuteApiTimeoutSeconds) {
		toSerialize["executeApiTimeoutSeconds"] = o.ExecuteApiTimeoutSeconds
	}
	if !IsNil(o.WaitUntilApiRetryPolicy) {
		toSerialize["waitUntilApiRetryPolicy"] = o.WaitUntilApiRetryPolicy
	}
	if !IsNil(o.ExecuteApiRetryPolicy) {
		toSerialize["executeApiRetryPolicy"] = o.ExecuteApiRetryPolicy
	}
	if !IsNil(o.StateFailureRecoveryOptions) {
		toSerialize["stateFailureRecoveryOptions"] = o.StateFailureRecoveryOptions
	}
	if !IsNil(o.LoadGlobalAttributesRequest) {
		toSerialize["loadGlobalAttributesRequest"] = o.LoadGlobalAttributesRequest
	}
	return toSerialize, nil
}

type NullableAsyncStateConfig struct {
	value *AsyncStateConfig
	isSet bool
}

func (v NullableAsyncStateConfig) Get() *AsyncStateConfig {
	return v.value
}

func (v *NullableAsyncStateConfig) Set(val *AsyncStateConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAsyncStateConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAsyncStateConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsyncStateConfig(val *AsyncStateConfig) *NullableAsyncStateConfig {
	return &NullableAsyncStateConfig{value: val, isSet: true}
}

func (v NullableAsyncStateConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsyncStateConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

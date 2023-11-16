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

// checks if the AsyncStateExecuteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsyncStateExecuteResponse{}

// AsyncStateExecuteResponse the output of the execute API
type AsyncStateExecuteResponse struct {
	StateDecision           StateDecision                   `json:"stateDecision"`
	PublishToLocalQueue     []LocalQueueMessage             `json:"publishToLocalQueue,omitempty"`
	WriteToGlobalAttributes []GlobalAttributeTableRowUpdate `json:"writeToGlobalAttributes,omitempty"`
	WriteToLocalAttributes  []KeyValue                      `json:"writeToLocalAttributes,omitempty"`
}

type _AsyncStateExecuteResponse AsyncStateExecuteResponse

// NewAsyncStateExecuteResponse instantiates a new AsyncStateExecuteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsyncStateExecuteResponse(stateDecision StateDecision) *AsyncStateExecuteResponse {
	this := AsyncStateExecuteResponse{}
	this.StateDecision = stateDecision
	return &this
}

// NewAsyncStateExecuteResponseWithDefaults instantiates a new AsyncStateExecuteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsyncStateExecuteResponseWithDefaults() *AsyncStateExecuteResponse {
	this := AsyncStateExecuteResponse{}
	return &this
}

// GetStateDecision returns the StateDecision field value
func (o *AsyncStateExecuteResponse) GetStateDecision() StateDecision {
	if o == nil {
		var ret StateDecision
		return ret
	}

	return o.StateDecision
}

// GetStateDecisionOk returns a tuple with the StateDecision field value
// and a boolean to check if the value has been set.
func (o *AsyncStateExecuteResponse) GetStateDecisionOk() (*StateDecision, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateDecision, true
}

// SetStateDecision sets field value
func (o *AsyncStateExecuteResponse) SetStateDecision(v StateDecision) {
	o.StateDecision = v
}

// GetPublishToLocalQueue returns the PublishToLocalQueue field value if set, zero value otherwise.
func (o *AsyncStateExecuteResponse) GetPublishToLocalQueue() []LocalQueueMessage {
	if o == nil || IsNil(o.PublishToLocalQueue) {
		var ret []LocalQueueMessage
		return ret
	}
	return o.PublishToLocalQueue
}

// GetPublishToLocalQueueOk returns a tuple with the PublishToLocalQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncStateExecuteResponse) GetPublishToLocalQueueOk() ([]LocalQueueMessage, bool) {
	if o == nil || IsNil(o.PublishToLocalQueue) {
		return nil, false
	}
	return o.PublishToLocalQueue, true
}

// HasPublishToLocalQueue returns a boolean if a field has been set.
func (o *AsyncStateExecuteResponse) HasPublishToLocalQueue() bool {
	if o != nil && !IsNil(o.PublishToLocalQueue) {
		return true
	}

	return false
}

// SetPublishToLocalQueue gets a reference to the given []LocalQueueMessage and assigns it to the PublishToLocalQueue field.
func (o *AsyncStateExecuteResponse) SetPublishToLocalQueue(v []LocalQueueMessage) {
	o.PublishToLocalQueue = v
}

// GetWriteToGlobalAttributes returns the WriteToGlobalAttributes field value if set, zero value otherwise.
func (o *AsyncStateExecuteResponse) GetWriteToGlobalAttributes() []GlobalAttributeTableRowUpdate {
	if o == nil || IsNil(o.WriteToGlobalAttributes) {
		var ret []GlobalAttributeTableRowUpdate
		return ret
	}
	return o.WriteToGlobalAttributes
}

// GetWriteToGlobalAttributesOk returns a tuple with the WriteToGlobalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncStateExecuteResponse) GetWriteToGlobalAttributesOk() ([]GlobalAttributeTableRowUpdate, bool) {
	if o == nil || IsNil(o.WriteToGlobalAttributes) {
		return nil, false
	}
	return o.WriteToGlobalAttributes, true
}

// HasWriteToGlobalAttributes returns a boolean if a field has been set.
func (o *AsyncStateExecuteResponse) HasWriteToGlobalAttributes() bool {
	if o != nil && !IsNil(o.WriteToGlobalAttributes) {
		return true
	}

	return false
}

// SetWriteToGlobalAttributes gets a reference to the given []GlobalAttributeTableRowUpdate and assigns it to the WriteToGlobalAttributes field.
func (o *AsyncStateExecuteResponse) SetWriteToGlobalAttributes(v []GlobalAttributeTableRowUpdate) {
	o.WriteToGlobalAttributes = v
}

// GetWriteToLocalAttributes returns the WriteToLocalAttributes field value if set, zero value otherwise.
func (o *AsyncStateExecuteResponse) GetWriteToLocalAttributes() []KeyValue {
	if o == nil || IsNil(o.WriteToLocalAttributes) {
		var ret []KeyValue
		return ret
	}
	return o.WriteToLocalAttributes
}

// GetWriteToLocalAttributesOk returns a tuple with the WriteToLocalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncStateExecuteResponse) GetWriteToLocalAttributesOk() ([]KeyValue, bool) {
	if o == nil || IsNil(o.WriteToLocalAttributes) {
		return nil, false
	}
	return o.WriteToLocalAttributes, true
}

// HasWriteToLocalAttributes returns a boolean if a field has been set.
func (o *AsyncStateExecuteResponse) HasWriteToLocalAttributes() bool {
	if o != nil && !IsNil(o.WriteToLocalAttributes) {
		return true
	}

	return false
}

// SetWriteToLocalAttributes gets a reference to the given []KeyValue and assigns it to the WriteToLocalAttributes field.
func (o *AsyncStateExecuteResponse) SetWriteToLocalAttributes(v []KeyValue) {
	o.WriteToLocalAttributes = v
}

func (o AsyncStateExecuteResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AsyncStateExecuteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stateDecision"] = o.StateDecision
	if !IsNil(o.PublishToLocalQueue) {
		toSerialize["publishToLocalQueue"] = o.PublishToLocalQueue
	}
	if !IsNil(o.WriteToGlobalAttributes) {
		toSerialize["writeToGlobalAttributes"] = o.WriteToGlobalAttributes
	}
	if !IsNil(o.WriteToLocalAttributes) {
		toSerialize["writeToLocalAttributes"] = o.WriteToLocalAttributes
	}
	return toSerialize, nil
}

func (o *AsyncStateExecuteResponse) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"stateDecision",
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

	varAsyncStateExecuteResponse := _AsyncStateExecuteResponse{}

	err = json.Unmarshal(bytes, &varAsyncStateExecuteResponse)

	if err != nil {
		return err
	}

	*o = AsyncStateExecuteResponse(varAsyncStateExecuteResponse)

	return err
}

type NullableAsyncStateExecuteResponse struct {
	value *AsyncStateExecuteResponse
	isSet bool
}

func (v NullableAsyncStateExecuteResponse) Get() *AsyncStateExecuteResponse {
	return v.value
}

func (v *NullableAsyncStateExecuteResponse) Set(val *AsyncStateExecuteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAsyncStateExecuteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAsyncStateExecuteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsyncStateExecuteResponse(val *AsyncStateExecuteResponse) *NullableAsyncStateExecuteResponse {
	return &NullableAsyncStateExecuteResponse{value: val, isSet: true}
}

func (v NullableAsyncStateExecuteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsyncStateExecuteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
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

// checks if the AsyncStateWaitUntilResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsyncStateWaitUntilResponse{}

// AsyncStateWaitUntilResponse the output of the waitUntil API
type AsyncStateWaitUntilResponse struct {
	CommandRequest      CommandRequest      `json:"commandRequest"`
	PublishToLocalQueue []LocalQueueMessage `json:"publishToLocalQueue,omitempty"`
}

// NewAsyncStateWaitUntilResponse instantiates a new AsyncStateWaitUntilResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsyncStateWaitUntilResponse(commandRequest CommandRequest) *AsyncStateWaitUntilResponse {
	this := AsyncStateWaitUntilResponse{}
	this.CommandRequest = commandRequest
	return &this
}

// NewAsyncStateWaitUntilResponseWithDefaults instantiates a new AsyncStateWaitUntilResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsyncStateWaitUntilResponseWithDefaults() *AsyncStateWaitUntilResponse {
	this := AsyncStateWaitUntilResponse{}
	return &this
}

// GetCommandRequest returns the CommandRequest field value
func (o *AsyncStateWaitUntilResponse) GetCommandRequest() CommandRequest {
	if o == nil {
		var ret CommandRequest
		return ret
	}

	return o.CommandRequest
}

// GetCommandRequestOk returns a tuple with the CommandRequest field value
// and a boolean to check if the value has been set.
func (o *AsyncStateWaitUntilResponse) GetCommandRequestOk() (*CommandRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommandRequest, true
}

// SetCommandRequest sets field value
func (o *AsyncStateWaitUntilResponse) SetCommandRequest(v CommandRequest) {
	o.CommandRequest = v
}

// GetPublishToLocalQueue returns the PublishToLocalQueue field value if set, zero value otherwise.
func (o *AsyncStateWaitUntilResponse) GetPublishToLocalQueue() []LocalQueueMessage {
	if o == nil || IsNil(o.PublishToLocalQueue) {
		var ret []LocalQueueMessage
		return ret
	}
	return o.PublishToLocalQueue
}

// GetPublishToLocalQueueOk returns a tuple with the PublishToLocalQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncStateWaitUntilResponse) GetPublishToLocalQueueOk() ([]LocalQueueMessage, bool) {
	if o == nil || IsNil(o.PublishToLocalQueue) {
		return nil, false
	}
	return o.PublishToLocalQueue, true
}

// HasPublishToLocalQueue returns a boolean if a field has been set.
func (o *AsyncStateWaitUntilResponse) HasPublishToLocalQueue() bool {
	if o != nil && !IsNil(o.PublishToLocalQueue) {
		return true
	}

	return false
}

// SetPublishToLocalQueue gets a reference to the given []LocalQueueMessage and assigns it to the PublishToLocalQueue field.
func (o *AsyncStateWaitUntilResponse) SetPublishToLocalQueue(v []LocalQueueMessage) {
	o.PublishToLocalQueue = v
}

func (o AsyncStateWaitUntilResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AsyncStateWaitUntilResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["commandRequest"] = o.CommandRequest
	if !IsNil(o.PublishToLocalQueue) {
		toSerialize["publishToLocalQueue"] = o.PublishToLocalQueue
	}
	return toSerialize, nil
}

type NullableAsyncStateWaitUntilResponse struct {
	value *AsyncStateWaitUntilResponse
	isSet bool
}

func (v NullableAsyncStateWaitUntilResponse) Get() *AsyncStateWaitUntilResponse {
	return v.value
}

func (v *NullableAsyncStateWaitUntilResponse) Set(val *AsyncStateWaitUntilResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAsyncStateWaitUntilResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAsyncStateWaitUntilResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsyncStateWaitUntilResponse(val *AsyncStateWaitUntilResponse) *NullableAsyncStateWaitUntilResponse {
	return &NullableAsyncStateWaitUntilResponse{value: val, isSet: true}
}

func (v NullableAsyncStateWaitUntilResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsyncStateWaitUntilResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

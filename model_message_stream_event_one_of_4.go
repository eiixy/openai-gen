/*
OpenAI API

The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openai

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MessageStreamEventOneOf4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageStreamEventOneOf4{}

// MessageStreamEventOneOf4 Occurs when a [message](/docs/api-reference/messages/object) ends before it is completed.
type MessageStreamEventOneOf4 struct {
	Event string `json:"event"`
	Data MessageObject `json:"data"`
}

type _MessageStreamEventOneOf4 MessageStreamEventOneOf4

// NewMessageStreamEventOneOf4 instantiates a new MessageStreamEventOneOf4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageStreamEventOneOf4(event string, data MessageObject) *MessageStreamEventOneOf4 {
	this := MessageStreamEventOneOf4{}
	this.Event = event
	this.Data = data
	return &this
}

// NewMessageStreamEventOneOf4WithDefaults instantiates a new MessageStreamEventOneOf4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageStreamEventOneOf4WithDefaults() *MessageStreamEventOneOf4 {
	this := MessageStreamEventOneOf4{}
	return &this
}

// GetEvent returns the Event field value
func (o *MessageStreamEventOneOf4) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *MessageStreamEventOneOf4) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *MessageStreamEventOneOf4) SetEvent(v string) {
	o.Event = v
}

// GetData returns the Data field value
func (o *MessageStreamEventOneOf4) GetData() MessageObject {
	if o == nil {
		var ret MessageObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MessageStreamEventOneOf4) GetDataOk() (*MessageObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MessageStreamEventOneOf4) SetData(v MessageObject) {
	o.Data = v
}

func (o MessageStreamEventOneOf4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageStreamEventOneOf4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *MessageStreamEventOneOf4) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event",
		"data",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMessageStreamEventOneOf4 := _MessageStreamEventOneOf4{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageStreamEventOneOf4)

	if err != nil {
		return err
	}

	*o = MessageStreamEventOneOf4(varMessageStreamEventOneOf4)

	return err
}

type NullableMessageStreamEventOneOf4 struct {
	value *MessageStreamEventOneOf4
	isSet bool
}

func (v NullableMessageStreamEventOneOf4) Get() *MessageStreamEventOneOf4 {
	return v.value
}

func (v *NullableMessageStreamEventOneOf4) Set(val *MessageStreamEventOneOf4) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageStreamEventOneOf4) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageStreamEventOneOf4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageStreamEventOneOf4(val *MessageStreamEventOneOf4) *NullableMessageStreamEventOneOf4 {
	return &NullableMessageStreamEventOneOf4{value: val, isSet: true}
}

func (v NullableMessageStreamEventOneOf4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageStreamEventOneOf4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



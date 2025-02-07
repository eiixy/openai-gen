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

// checks if the MessageDeltaContentImageUrlObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageDeltaContentImageUrlObject{}

// MessageDeltaContentImageUrlObject References an image URL in the content of a message.
type MessageDeltaContentImageUrlObject struct {
	// The index of the content part in the message.
	Index int32 `json:"index"`
	// Always `image_url`.
	Type string `json:"type"`
	ImageUrl *MessageDeltaContentImageUrlObjectImageUrl `json:"image_url,omitempty"`
}

type _MessageDeltaContentImageUrlObject MessageDeltaContentImageUrlObject

// NewMessageDeltaContentImageUrlObject instantiates a new MessageDeltaContentImageUrlObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageDeltaContentImageUrlObject(index int32, type_ string) *MessageDeltaContentImageUrlObject {
	this := MessageDeltaContentImageUrlObject{}
	this.Index = index
	this.Type = type_
	return &this
}

// NewMessageDeltaContentImageUrlObjectWithDefaults instantiates a new MessageDeltaContentImageUrlObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageDeltaContentImageUrlObjectWithDefaults() *MessageDeltaContentImageUrlObject {
	this := MessageDeltaContentImageUrlObject{}
	return &this
}

// GetIndex returns the Index field value
func (o *MessageDeltaContentImageUrlObject) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *MessageDeltaContentImageUrlObject) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *MessageDeltaContentImageUrlObject) SetIndex(v int32) {
	o.Index = v
}

// GetType returns the Type field value
func (o *MessageDeltaContentImageUrlObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MessageDeltaContentImageUrlObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MessageDeltaContentImageUrlObject) SetType(v string) {
	o.Type = v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *MessageDeltaContentImageUrlObject) GetImageUrl() MessageDeltaContentImageUrlObjectImageUrl {
	if o == nil || IsNil(o.ImageUrl) {
		var ret MessageDeltaContentImageUrlObjectImageUrl
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDeltaContentImageUrlObject) GetImageUrlOk() (*MessageDeltaContentImageUrlObjectImageUrl, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *MessageDeltaContentImageUrlObject) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given MessageDeltaContentImageUrlObjectImageUrl and assigns it to the ImageUrl field.
func (o *MessageDeltaContentImageUrlObject) SetImageUrl(v MessageDeltaContentImageUrlObjectImageUrl) {
	o.ImageUrl = &v
}

func (o MessageDeltaContentImageUrlObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageDeltaContentImageUrlObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["type"] = o.Type
	if !IsNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	return toSerialize, nil
}

func (o *MessageDeltaContentImageUrlObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
		"type",
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

	varMessageDeltaContentImageUrlObject := _MessageDeltaContentImageUrlObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageDeltaContentImageUrlObject)

	if err != nil {
		return err
	}

	*o = MessageDeltaContentImageUrlObject(varMessageDeltaContentImageUrlObject)

	return err
}

type NullableMessageDeltaContentImageUrlObject struct {
	value *MessageDeltaContentImageUrlObject
	isSet bool
}

func (v NullableMessageDeltaContentImageUrlObject) Get() *MessageDeltaContentImageUrlObject {
	return v.value
}

func (v *NullableMessageDeltaContentImageUrlObject) Set(val *MessageDeltaContentImageUrlObject) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageDeltaContentImageUrlObject) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageDeltaContentImageUrlObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageDeltaContentImageUrlObject(val *MessageDeltaContentImageUrlObject) *NullableMessageDeltaContentImageUrlObject {
	return &NullableMessageDeltaContentImageUrlObject{value: val, isSet: true}
}

func (v NullableMessageDeltaContentImageUrlObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageDeltaContentImageUrlObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



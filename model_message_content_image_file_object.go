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

// checks if the MessageContentImageFileObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageContentImageFileObject{}

// MessageContentImageFileObject References an image [File](/docs/api-reference/files) in the content of a message.
type MessageContentImageFileObject struct {
	// Always `image_file`.
	Type string `json:"type"`
	ImageFile MessageContentImageFileObjectImageFile `json:"image_file"`
}

type _MessageContentImageFileObject MessageContentImageFileObject

// NewMessageContentImageFileObject instantiates a new MessageContentImageFileObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageContentImageFileObject(type_ string, imageFile MessageContentImageFileObjectImageFile) *MessageContentImageFileObject {
	this := MessageContentImageFileObject{}
	this.Type = type_
	this.ImageFile = imageFile
	return &this
}

// NewMessageContentImageFileObjectWithDefaults instantiates a new MessageContentImageFileObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageContentImageFileObjectWithDefaults() *MessageContentImageFileObject {
	this := MessageContentImageFileObject{}
	return &this
}

// GetType returns the Type field value
func (o *MessageContentImageFileObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MessageContentImageFileObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MessageContentImageFileObject) SetType(v string) {
	o.Type = v
}

// GetImageFile returns the ImageFile field value
func (o *MessageContentImageFileObject) GetImageFile() MessageContentImageFileObjectImageFile {
	if o == nil {
		var ret MessageContentImageFileObjectImageFile
		return ret
	}

	return o.ImageFile
}

// GetImageFileOk returns a tuple with the ImageFile field value
// and a boolean to check if the value has been set.
func (o *MessageContentImageFileObject) GetImageFileOk() (*MessageContentImageFileObjectImageFile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageFile, true
}

// SetImageFile sets field value
func (o *MessageContentImageFileObject) SetImageFile(v MessageContentImageFileObjectImageFile) {
	o.ImageFile = v
}

func (o MessageContentImageFileObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageContentImageFileObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["image_file"] = o.ImageFile
	return toSerialize, nil
}

func (o *MessageContentImageFileObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"image_file",
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

	varMessageContentImageFileObject := _MessageContentImageFileObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageContentImageFileObject)

	if err != nil {
		return err
	}

	*o = MessageContentImageFileObject(varMessageContentImageFileObject)

	return err
}

type NullableMessageContentImageFileObject struct {
	value *MessageContentImageFileObject
	isSet bool
}

func (v NullableMessageContentImageFileObject) Get() *MessageContentImageFileObject {
	return v.value
}

func (v *NullableMessageContentImageFileObject) Set(val *MessageContentImageFileObject) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageContentImageFileObject) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageContentImageFileObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageContentImageFileObject(val *MessageContentImageFileObject) *NullableMessageContentImageFileObject {
	return &NullableMessageContentImageFileObject{value: val, isSet: true}
}

func (v NullableMessageContentImageFileObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageContentImageFileObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



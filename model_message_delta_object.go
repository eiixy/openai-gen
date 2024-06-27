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

// checks if the MessageDeltaObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageDeltaObject{}

// MessageDeltaObject Represents a message delta i.e. any changed fields on a message during streaming. 
type MessageDeltaObject struct {
	// The identifier of the message, which can be referenced in API endpoints.
	Id string `json:"id"`
	// The object type, which is always `thread.message.delta`.
	Object string `json:"object"`
	Delta MessageDeltaObjectDelta `json:"delta"`
}

type _MessageDeltaObject MessageDeltaObject

// NewMessageDeltaObject instantiates a new MessageDeltaObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageDeltaObject(id string, object string, delta MessageDeltaObjectDelta) *MessageDeltaObject {
	this := MessageDeltaObject{}
	this.Id = id
	this.Object = object
	this.Delta = delta
	return &this
}

// NewMessageDeltaObjectWithDefaults instantiates a new MessageDeltaObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageDeltaObjectWithDefaults() *MessageDeltaObject {
	this := MessageDeltaObject{}
	return &this
}

// GetId returns the Id field value
func (o *MessageDeltaObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MessageDeltaObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MessageDeltaObject) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *MessageDeltaObject) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *MessageDeltaObject) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *MessageDeltaObject) SetObject(v string) {
	o.Object = v
}

// GetDelta returns the Delta field value
func (o *MessageDeltaObject) GetDelta() MessageDeltaObjectDelta {
	if o == nil {
		var ret MessageDeltaObjectDelta
		return ret
	}

	return o.Delta
}

// GetDeltaOk returns a tuple with the Delta field value
// and a boolean to check if the value has been set.
func (o *MessageDeltaObject) GetDeltaOk() (*MessageDeltaObjectDelta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delta, true
}

// SetDelta sets field value
func (o *MessageDeltaObject) SetDelta(v MessageDeltaObjectDelta) {
	o.Delta = v
}

func (o MessageDeltaObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageDeltaObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["delta"] = o.Delta
	return toSerialize, nil
}

func (o *MessageDeltaObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"delta",
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

	varMessageDeltaObject := _MessageDeltaObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageDeltaObject)

	if err != nil {
		return err
	}

	*o = MessageDeltaObject(varMessageDeltaObject)

	return err
}

type NullableMessageDeltaObject struct {
	value *MessageDeltaObject
	isSet bool
}

func (v NullableMessageDeltaObject) Get() *MessageDeltaObject {
	return v.value
}

func (v *NullableMessageDeltaObject) Set(val *MessageDeltaObject) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageDeltaObject) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageDeltaObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageDeltaObject(val *MessageDeltaObject) *NullableMessageDeltaObject {
	return &NullableMessageDeltaObject{value: val, isSet: true}
}

func (v NullableMessageDeltaObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageDeltaObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



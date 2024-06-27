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

// checks if the CreateTranslationResponseJson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTranslationResponseJson{}

// CreateTranslationResponseJson struct for CreateTranslationResponseJson
type CreateTranslationResponseJson struct {
	Text string `json:"text"`
}

type _CreateTranslationResponseJson CreateTranslationResponseJson

// NewCreateTranslationResponseJson instantiates a new CreateTranslationResponseJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTranslationResponseJson(text string) *CreateTranslationResponseJson {
	this := CreateTranslationResponseJson{}
	this.Text = text
	return &this
}

// NewCreateTranslationResponseJsonWithDefaults instantiates a new CreateTranslationResponseJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTranslationResponseJsonWithDefaults() *CreateTranslationResponseJson {
	this := CreateTranslationResponseJson{}
	return &this
}

// GetText returns the Text field value
func (o *CreateTranslationResponseJson) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *CreateTranslationResponseJson) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *CreateTranslationResponseJson) SetText(v string) {
	o.Text = v
}

func (o CreateTranslationResponseJson) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTranslationResponseJson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	return toSerialize, nil
}

func (o *CreateTranslationResponseJson) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"text",
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

	varCreateTranslationResponseJson := _CreateTranslationResponseJson{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTranslationResponseJson)

	if err != nil {
		return err
	}

	*o = CreateTranslationResponseJson(varCreateTranslationResponseJson)

	return err
}

type NullableCreateTranslationResponseJson struct {
	value *CreateTranslationResponseJson
	isSet bool
}

func (v NullableCreateTranslationResponseJson) Get() *CreateTranslationResponseJson {
	return v.value
}

func (v *NullableCreateTranslationResponseJson) Set(val *CreateTranslationResponseJson) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTranslationResponseJson) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTranslationResponseJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTranslationResponseJson(val *CreateTranslationResponseJson) *NullableCreateTranslationResponseJson {
	return &NullableCreateTranslationResponseJson{value: val, isSet: true}
}

func (v NullableCreateTranslationResponseJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTranslationResponseJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



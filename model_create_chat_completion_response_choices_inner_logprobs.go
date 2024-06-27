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

// checks if the CreateChatCompletionResponseChoicesInnerLogprobs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateChatCompletionResponseChoicesInnerLogprobs{}

// CreateChatCompletionResponseChoicesInnerLogprobs Log probability information for the choice.
type CreateChatCompletionResponseChoicesInnerLogprobs struct {
	// A list of message content tokens with log probability information.
	Content []ChatCompletionTokenLogprob `json:"content"`
}

type _CreateChatCompletionResponseChoicesInnerLogprobs CreateChatCompletionResponseChoicesInnerLogprobs

// NewCreateChatCompletionResponseChoicesInnerLogprobs instantiates a new CreateChatCompletionResponseChoicesInnerLogprobs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateChatCompletionResponseChoicesInnerLogprobs(content []ChatCompletionTokenLogprob) *CreateChatCompletionResponseChoicesInnerLogprobs {
	this := CreateChatCompletionResponseChoicesInnerLogprobs{}
	this.Content = content
	return &this
}

// NewCreateChatCompletionResponseChoicesInnerLogprobsWithDefaults instantiates a new CreateChatCompletionResponseChoicesInnerLogprobs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateChatCompletionResponseChoicesInnerLogprobsWithDefaults() *CreateChatCompletionResponseChoicesInnerLogprobs {
	this := CreateChatCompletionResponseChoicesInnerLogprobs{}
	return &this
}

// GetContent returns the Content field value
// If the value is explicit nil, the zero value for []ChatCompletionTokenLogprob will be returned
func (o *CreateChatCompletionResponseChoicesInnerLogprobs) GetContent() []ChatCompletionTokenLogprob {
	if o == nil {
		var ret []ChatCompletionTokenLogprob
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateChatCompletionResponseChoicesInnerLogprobs) GetContentOk() ([]ChatCompletionTokenLogprob, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// SetContent sets field value
func (o *CreateChatCompletionResponseChoicesInnerLogprobs) SetContent(v []ChatCompletionTokenLogprob) {
	o.Content = v
}

func (o CreateChatCompletionResponseChoicesInnerLogprobs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateChatCompletionResponseChoicesInnerLogprobs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

func (o *CreateChatCompletionResponseChoicesInnerLogprobs) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content",
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

	varCreateChatCompletionResponseChoicesInnerLogprobs := _CreateChatCompletionResponseChoicesInnerLogprobs{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateChatCompletionResponseChoicesInnerLogprobs)

	if err != nil {
		return err
	}

	*o = CreateChatCompletionResponseChoicesInnerLogprobs(varCreateChatCompletionResponseChoicesInnerLogprobs)

	return err
}

type NullableCreateChatCompletionResponseChoicesInnerLogprobs struct {
	value *CreateChatCompletionResponseChoicesInnerLogprobs
	isSet bool
}

func (v NullableCreateChatCompletionResponseChoicesInnerLogprobs) Get() *CreateChatCompletionResponseChoicesInnerLogprobs {
	return v.value
}

func (v *NullableCreateChatCompletionResponseChoicesInnerLogprobs) Set(val *CreateChatCompletionResponseChoicesInnerLogprobs) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateChatCompletionResponseChoicesInnerLogprobs) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateChatCompletionResponseChoicesInnerLogprobs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateChatCompletionResponseChoicesInnerLogprobs(val *CreateChatCompletionResponseChoicesInnerLogprobs) *NullableCreateChatCompletionResponseChoicesInnerLogprobs {
	return &NullableCreateChatCompletionResponseChoicesInnerLogprobs{value: val, isSet: true}
}

func (v NullableCreateChatCompletionResponseChoicesInnerLogprobs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateChatCompletionResponseChoicesInnerLogprobs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



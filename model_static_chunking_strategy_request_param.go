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

// checks if the StaticChunkingStrategyRequestParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticChunkingStrategyRequestParam{}

// StaticChunkingStrategyRequestParam struct for StaticChunkingStrategyRequestParam
type StaticChunkingStrategyRequestParam struct {
	// Always `static`.
	Type string `json:"type"`
	Static StaticChunkingStrategy `json:"static"`
}

type _StaticChunkingStrategyRequestParam StaticChunkingStrategyRequestParam

// NewStaticChunkingStrategyRequestParam instantiates a new StaticChunkingStrategyRequestParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticChunkingStrategyRequestParam(type_ string, static StaticChunkingStrategy) *StaticChunkingStrategyRequestParam {
	this := StaticChunkingStrategyRequestParam{}
	this.Type = type_
	this.Static = static
	return &this
}

// NewStaticChunkingStrategyRequestParamWithDefaults instantiates a new StaticChunkingStrategyRequestParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticChunkingStrategyRequestParamWithDefaults() *StaticChunkingStrategyRequestParam {
	this := StaticChunkingStrategyRequestParam{}
	return &this
}

// GetType returns the Type field value
func (o *StaticChunkingStrategyRequestParam) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StaticChunkingStrategyRequestParam) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StaticChunkingStrategyRequestParam) SetType(v string) {
	o.Type = v
}

// GetStatic returns the Static field value
func (o *StaticChunkingStrategyRequestParam) GetStatic() StaticChunkingStrategy {
	if o == nil {
		var ret StaticChunkingStrategy
		return ret
	}

	return o.Static
}

// GetStaticOk returns a tuple with the Static field value
// and a boolean to check if the value has been set.
func (o *StaticChunkingStrategyRequestParam) GetStaticOk() (*StaticChunkingStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Static, true
}

// SetStatic sets field value
func (o *StaticChunkingStrategyRequestParam) SetStatic(v StaticChunkingStrategy) {
	o.Static = v
}

func (o StaticChunkingStrategyRequestParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticChunkingStrategyRequestParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["static"] = o.Static
	return toSerialize, nil
}

func (o *StaticChunkingStrategyRequestParam) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"static",
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

	varStaticChunkingStrategyRequestParam := _StaticChunkingStrategyRequestParam{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStaticChunkingStrategyRequestParam)

	if err != nil {
		return err
	}

	*o = StaticChunkingStrategyRequestParam(varStaticChunkingStrategyRequestParam)

	return err
}

type NullableStaticChunkingStrategyRequestParam struct {
	value *StaticChunkingStrategyRequestParam
	isSet bool
}

func (v NullableStaticChunkingStrategyRequestParam) Get() *StaticChunkingStrategyRequestParam {
	return v.value
}

func (v *NullableStaticChunkingStrategyRequestParam) Set(val *StaticChunkingStrategyRequestParam) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticChunkingStrategyRequestParam) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticChunkingStrategyRequestParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticChunkingStrategyRequestParam(val *StaticChunkingStrategyRequestParam) *NullableStaticChunkingStrategyRequestParam {
	return &NullableStaticChunkingStrategyRequestParam{value: val, isSet: true}
}

func (v NullableStaticChunkingStrategyRequestParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticChunkingStrategyRequestParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



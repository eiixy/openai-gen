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

// checks if the RunCompletionUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunCompletionUsage{}

// RunCompletionUsage Usage statistics related to the run. This value will be `null` if the run is not in a terminal state (i.e. `in_progress`, `queued`, etc.).
type RunCompletionUsage struct {
	// Number of completion tokens used over the course of the run.
	CompletionTokens int32 `json:"completion_tokens"`
	// Number of prompt tokens used over the course of the run.
	PromptTokens int32 `json:"prompt_tokens"`
	// Total number of tokens used (prompt + completion).
	TotalTokens int32 `json:"total_tokens"`
}

type _RunCompletionUsage RunCompletionUsage

// NewRunCompletionUsage instantiates a new RunCompletionUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunCompletionUsage(completionTokens int32, promptTokens int32, totalTokens int32) *RunCompletionUsage {
	this := RunCompletionUsage{}
	this.CompletionTokens = completionTokens
	this.PromptTokens = promptTokens
	this.TotalTokens = totalTokens
	return &this
}

// NewRunCompletionUsageWithDefaults instantiates a new RunCompletionUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunCompletionUsageWithDefaults() *RunCompletionUsage {
	this := RunCompletionUsage{}
	return &this
}

// GetCompletionTokens returns the CompletionTokens field value
func (o *RunCompletionUsage) GetCompletionTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CompletionTokens
}

// GetCompletionTokensOk returns a tuple with the CompletionTokens field value
// and a boolean to check if the value has been set.
func (o *RunCompletionUsage) GetCompletionTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletionTokens, true
}

// SetCompletionTokens sets field value
func (o *RunCompletionUsage) SetCompletionTokens(v int32) {
	o.CompletionTokens = v
}

// GetPromptTokens returns the PromptTokens field value
func (o *RunCompletionUsage) GetPromptTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PromptTokens
}

// GetPromptTokensOk returns a tuple with the PromptTokens field value
// and a boolean to check if the value has been set.
func (o *RunCompletionUsage) GetPromptTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromptTokens, true
}

// SetPromptTokens sets field value
func (o *RunCompletionUsage) SetPromptTokens(v int32) {
	o.PromptTokens = v
}

// GetTotalTokens returns the TotalTokens field value
func (o *RunCompletionUsage) GetTotalTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalTokens
}

// GetTotalTokensOk returns a tuple with the TotalTokens field value
// and a boolean to check if the value has been set.
func (o *RunCompletionUsage) GetTotalTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTokens, true
}

// SetTotalTokens sets field value
func (o *RunCompletionUsage) SetTotalTokens(v int32) {
	o.TotalTokens = v
}

func (o RunCompletionUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunCompletionUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["completion_tokens"] = o.CompletionTokens
	toSerialize["prompt_tokens"] = o.PromptTokens
	toSerialize["total_tokens"] = o.TotalTokens
	return toSerialize, nil
}

func (o *RunCompletionUsage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"completion_tokens",
		"prompt_tokens",
		"total_tokens",
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

	varRunCompletionUsage := _RunCompletionUsage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunCompletionUsage)

	if err != nil {
		return err
	}

	*o = RunCompletionUsage(varRunCompletionUsage)

	return err
}

type NullableRunCompletionUsage struct {
	value *RunCompletionUsage
	isSet bool
}

func (v NullableRunCompletionUsage) Get() *RunCompletionUsage {
	return v.value
}

func (v *NullableRunCompletionUsage) Set(val *RunCompletionUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableRunCompletionUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableRunCompletionUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunCompletionUsage(val *RunCompletionUsage) *NullableRunCompletionUsage {
	return &NullableRunCompletionUsage{value: val, isSet: true}
}

func (v NullableRunCompletionUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunCompletionUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



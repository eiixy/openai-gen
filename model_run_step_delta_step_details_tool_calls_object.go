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

// checks if the RunStepDeltaStepDetailsToolCallsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStepDeltaStepDetailsToolCallsObject{}

// RunStepDeltaStepDetailsToolCallsObject Details of the tool call.
type RunStepDeltaStepDetailsToolCallsObject struct {
	// Always `tool_calls`.
	Type string `json:"type"`
	// An array of tool calls the run step was involved in. These can be associated with one of three types of tools: `code_interpreter`, `file_search`, or `function`. 
	ToolCalls []RunStepDeltaStepDetailsToolCallsObjectToolCallsInner `json:"tool_calls,omitempty"`
}

type _RunStepDeltaStepDetailsToolCallsObject RunStepDeltaStepDetailsToolCallsObject

// NewRunStepDeltaStepDetailsToolCallsObject instantiates a new RunStepDeltaStepDetailsToolCallsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStepDeltaStepDetailsToolCallsObject(type_ string) *RunStepDeltaStepDetailsToolCallsObject {
	this := RunStepDeltaStepDetailsToolCallsObject{}
	this.Type = type_
	return &this
}

// NewRunStepDeltaStepDetailsToolCallsObjectWithDefaults instantiates a new RunStepDeltaStepDetailsToolCallsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStepDeltaStepDetailsToolCallsObjectWithDefaults() *RunStepDeltaStepDetailsToolCallsObject {
	this := RunStepDeltaStepDetailsToolCallsObject{}
	return &this
}

// GetType returns the Type field value
func (o *RunStepDeltaStepDetailsToolCallsObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RunStepDeltaStepDetailsToolCallsObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RunStepDeltaStepDetailsToolCallsObject) SetType(v string) {
	o.Type = v
}

// GetToolCalls returns the ToolCalls field value if set, zero value otherwise.
func (o *RunStepDeltaStepDetailsToolCallsObject) GetToolCalls() []RunStepDeltaStepDetailsToolCallsObjectToolCallsInner {
	if o == nil || IsNil(o.ToolCalls) {
		var ret []RunStepDeltaStepDetailsToolCallsObjectToolCallsInner
		return ret
	}
	return o.ToolCalls
}

// GetToolCallsOk returns a tuple with the ToolCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStepDeltaStepDetailsToolCallsObject) GetToolCallsOk() ([]RunStepDeltaStepDetailsToolCallsObjectToolCallsInner, bool) {
	if o == nil || IsNil(o.ToolCalls) {
		return nil, false
	}
	return o.ToolCalls, true
}

// HasToolCalls returns a boolean if a field has been set.
func (o *RunStepDeltaStepDetailsToolCallsObject) HasToolCalls() bool {
	if o != nil && !IsNil(o.ToolCalls) {
		return true
	}

	return false
}

// SetToolCalls gets a reference to the given []RunStepDeltaStepDetailsToolCallsObjectToolCallsInner and assigns it to the ToolCalls field.
func (o *RunStepDeltaStepDetailsToolCallsObject) SetToolCalls(v []RunStepDeltaStepDetailsToolCallsObjectToolCallsInner) {
	o.ToolCalls = v
}

func (o RunStepDeltaStepDetailsToolCallsObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStepDeltaStepDetailsToolCallsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.ToolCalls) {
		toSerialize["tool_calls"] = o.ToolCalls
	}
	return toSerialize, nil
}

func (o *RunStepDeltaStepDetailsToolCallsObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varRunStepDeltaStepDetailsToolCallsObject := _RunStepDeltaStepDetailsToolCallsObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunStepDeltaStepDetailsToolCallsObject)

	if err != nil {
		return err
	}

	*o = RunStepDeltaStepDetailsToolCallsObject(varRunStepDeltaStepDetailsToolCallsObject)

	return err
}

type NullableRunStepDeltaStepDetailsToolCallsObject struct {
	value *RunStepDeltaStepDetailsToolCallsObject
	isSet bool
}

func (v NullableRunStepDeltaStepDetailsToolCallsObject) Get() *RunStepDeltaStepDetailsToolCallsObject {
	return v.value
}

func (v *NullableRunStepDeltaStepDetailsToolCallsObject) Set(val *RunStepDeltaStepDetailsToolCallsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStepDeltaStepDetailsToolCallsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStepDeltaStepDetailsToolCallsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStepDeltaStepDetailsToolCallsObject(val *RunStepDeltaStepDetailsToolCallsObject) *NullableRunStepDeltaStepDetailsToolCallsObject {
	return &NullableRunStepDeltaStepDetailsToolCallsObject{value: val, isSet: true}
}

func (v NullableRunStepDeltaStepDetailsToolCallsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStepDeltaStepDetailsToolCallsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



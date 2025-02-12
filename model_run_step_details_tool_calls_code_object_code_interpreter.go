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

// checks if the RunStepDetailsToolCallsCodeObjectCodeInterpreter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStepDetailsToolCallsCodeObjectCodeInterpreter{}

// RunStepDetailsToolCallsCodeObjectCodeInterpreter The Code Interpreter tool call definition.
type RunStepDetailsToolCallsCodeObjectCodeInterpreter struct {
	// The input to the Code Interpreter tool call.
	Input string `json:"input"`
	// The outputs from the Code Interpreter tool call. Code Interpreter can output one or more items, including text (`logs`) or images (`image`). Each of these are represented by a different object type.
	Outputs []RunStepDetailsToolCallsCodeObjectCodeInterpreterOutputsInner `json:"outputs"`
}

type _RunStepDetailsToolCallsCodeObjectCodeInterpreter RunStepDetailsToolCallsCodeObjectCodeInterpreter

// NewRunStepDetailsToolCallsCodeObjectCodeInterpreter instantiates a new RunStepDetailsToolCallsCodeObjectCodeInterpreter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStepDetailsToolCallsCodeObjectCodeInterpreter(input string, outputs []RunStepDetailsToolCallsCodeObjectCodeInterpreterOutputsInner) *RunStepDetailsToolCallsCodeObjectCodeInterpreter {
	this := RunStepDetailsToolCallsCodeObjectCodeInterpreter{}
	this.Input = input
	this.Outputs = outputs
	return &this
}

// NewRunStepDetailsToolCallsCodeObjectCodeInterpreterWithDefaults instantiates a new RunStepDetailsToolCallsCodeObjectCodeInterpreter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStepDetailsToolCallsCodeObjectCodeInterpreterWithDefaults() *RunStepDetailsToolCallsCodeObjectCodeInterpreter {
	this := RunStepDetailsToolCallsCodeObjectCodeInterpreter{}
	return &this
}

// GetInput returns the Input field value
func (o *RunStepDetailsToolCallsCodeObjectCodeInterpreter) GetInput() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *RunStepDetailsToolCallsCodeObjectCodeInterpreter) GetInputOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *RunStepDetailsToolCallsCodeObjectCodeInterpreter) SetInput(v string) {
	o.Input = v
}

// GetOutputs returns the Outputs field value
func (o *RunStepDetailsToolCallsCodeObjectCodeInterpreter) GetOutputs() []RunStepDetailsToolCallsCodeObjectCodeInterpreterOutputsInner {
	if o == nil {
		var ret []RunStepDetailsToolCallsCodeObjectCodeInterpreterOutputsInner
		return ret
	}

	return o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value
// and a boolean to check if the value has been set.
func (o *RunStepDetailsToolCallsCodeObjectCodeInterpreter) GetOutputsOk() ([]RunStepDetailsToolCallsCodeObjectCodeInterpreterOutputsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Outputs, true
}

// SetOutputs sets field value
func (o *RunStepDetailsToolCallsCodeObjectCodeInterpreter) SetOutputs(v []RunStepDetailsToolCallsCodeObjectCodeInterpreterOutputsInner) {
	o.Outputs = v
}

func (o RunStepDetailsToolCallsCodeObjectCodeInterpreter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStepDetailsToolCallsCodeObjectCodeInterpreter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["input"] = o.Input
	toSerialize["outputs"] = o.Outputs
	return toSerialize, nil
}

func (o *RunStepDetailsToolCallsCodeObjectCodeInterpreter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"input",
		"outputs",
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

	varRunStepDetailsToolCallsCodeObjectCodeInterpreter := _RunStepDetailsToolCallsCodeObjectCodeInterpreter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunStepDetailsToolCallsCodeObjectCodeInterpreter)

	if err != nil {
		return err
	}

	*o = RunStepDetailsToolCallsCodeObjectCodeInterpreter(varRunStepDetailsToolCallsCodeObjectCodeInterpreter)

	return err
}

type NullableRunStepDetailsToolCallsCodeObjectCodeInterpreter struct {
	value *RunStepDetailsToolCallsCodeObjectCodeInterpreter
	isSet bool
}

func (v NullableRunStepDetailsToolCallsCodeObjectCodeInterpreter) Get() *RunStepDetailsToolCallsCodeObjectCodeInterpreter {
	return v.value
}

func (v *NullableRunStepDetailsToolCallsCodeObjectCodeInterpreter) Set(val *RunStepDetailsToolCallsCodeObjectCodeInterpreter) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStepDetailsToolCallsCodeObjectCodeInterpreter) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStepDetailsToolCallsCodeObjectCodeInterpreter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStepDetailsToolCallsCodeObjectCodeInterpreter(val *RunStepDetailsToolCallsCodeObjectCodeInterpreter) *NullableRunStepDetailsToolCallsCodeObjectCodeInterpreter {
	return &NullableRunStepDetailsToolCallsCodeObjectCodeInterpreter{value: val, isSet: true}
}

func (v NullableRunStepDetailsToolCallsCodeObjectCodeInterpreter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStepDetailsToolCallsCodeObjectCodeInterpreter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



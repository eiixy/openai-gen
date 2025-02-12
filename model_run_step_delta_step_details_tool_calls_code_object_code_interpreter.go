/*
OpenAI API

The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openai

import (
	"encoding/json"
)

// checks if the RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter{}

// RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter The Code Interpreter tool call definition.
type RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter struct {
	// The input to the Code Interpreter tool call.
	Input *string `json:"input,omitempty"`
	// The outputs from the Code Interpreter tool call. Code Interpreter can output one or more items, including text (`logs`) or images (`image`). Each of these are represented by a different object type.
	Outputs []RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreterOutputsInner `json:"outputs,omitempty"`
}

// NewRunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter instantiates a new RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter() *RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter {
	this := RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter{}
	return &this
}

// NewRunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreterWithDefaults instantiates a new RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreterWithDefaults() *RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter {
	this := RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter{}
	return &this
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter) GetInput() string {
	if o == nil || IsNil(o.Input) {
		var ret string
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter) GetInputOk() (*string, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given string and assigns it to the Input field.
func (o *RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter) SetInput(v string) {
	o.Input = &v
}

// GetOutputs returns the Outputs field value if set, zero value otherwise.
func (o *RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter) GetOutputs() []RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreterOutputsInner {
	if o == nil || IsNil(o.Outputs) {
		var ret []RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreterOutputsInner
		return ret
	}
	return o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter) GetOutputsOk() ([]RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreterOutputsInner, bool) {
	if o == nil || IsNil(o.Outputs) {
		return nil, false
	}
	return o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter) HasOutputs() bool {
	if o != nil && !IsNil(o.Outputs) {
		return true
	}

	return false
}

// SetOutputs gets a reference to the given []RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreterOutputsInner and assigns it to the Outputs field.
func (o *RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter) SetOutputs(v []RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreterOutputsInner) {
	o.Outputs = v
}

func (o RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	if !IsNil(o.Outputs) {
		toSerialize["outputs"] = o.Outputs
	}
	return toSerialize, nil
}

type NullableRunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter struct {
	value *RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter
	isSet bool
}

func (v NullableRunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter) Get() *RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter {
	return v.value
}

func (v *NullableRunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter) Set(val *RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter(val *RunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter) *NullableRunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter {
	return &NullableRunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter{value: val, isSet: true}
}

func (v NullableRunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStepDeltaStepDetailsToolCallsCodeObjectCodeInterpreter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



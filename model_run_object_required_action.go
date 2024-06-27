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

// checks if the RunObjectRequiredAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunObjectRequiredAction{}

// RunObjectRequiredAction Details on the action required to continue the run. Will be `null` if no action is required.
type RunObjectRequiredAction struct {
	// For now, this is always `submit_tool_outputs`.
	Type string `json:"type"`
	SubmitToolOutputs RunObjectRequiredActionSubmitToolOutputs `json:"submit_tool_outputs"`
}

type _RunObjectRequiredAction RunObjectRequiredAction

// NewRunObjectRequiredAction instantiates a new RunObjectRequiredAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunObjectRequiredAction(type_ string, submitToolOutputs RunObjectRequiredActionSubmitToolOutputs) *RunObjectRequiredAction {
	this := RunObjectRequiredAction{}
	this.Type = type_
	this.SubmitToolOutputs = submitToolOutputs
	return &this
}

// NewRunObjectRequiredActionWithDefaults instantiates a new RunObjectRequiredAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunObjectRequiredActionWithDefaults() *RunObjectRequiredAction {
	this := RunObjectRequiredAction{}
	return &this
}

// GetType returns the Type field value
func (o *RunObjectRequiredAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RunObjectRequiredAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RunObjectRequiredAction) SetType(v string) {
	o.Type = v
}

// GetSubmitToolOutputs returns the SubmitToolOutputs field value
func (o *RunObjectRequiredAction) GetSubmitToolOutputs() RunObjectRequiredActionSubmitToolOutputs {
	if o == nil {
		var ret RunObjectRequiredActionSubmitToolOutputs
		return ret
	}

	return o.SubmitToolOutputs
}

// GetSubmitToolOutputsOk returns a tuple with the SubmitToolOutputs field value
// and a boolean to check if the value has been set.
func (o *RunObjectRequiredAction) GetSubmitToolOutputsOk() (*RunObjectRequiredActionSubmitToolOutputs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubmitToolOutputs, true
}

// SetSubmitToolOutputs sets field value
func (o *RunObjectRequiredAction) SetSubmitToolOutputs(v RunObjectRequiredActionSubmitToolOutputs) {
	o.SubmitToolOutputs = v
}

func (o RunObjectRequiredAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunObjectRequiredAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["submit_tool_outputs"] = o.SubmitToolOutputs
	return toSerialize, nil
}

func (o *RunObjectRequiredAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"submit_tool_outputs",
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

	varRunObjectRequiredAction := _RunObjectRequiredAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunObjectRequiredAction)

	if err != nil {
		return err
	}

	*o = RunObjectRequiredAction(varRunObjectRequiredAction)

	return err
}

type NullableRunObjectRequiredAction struct {
	value *RunObjectRequiredAction
	isSet bool
}

func (v NullableRunObjectRequiredAction) Get() *RunObjectRequiredAction {
	return v.value
}

func (v *NullableRunObjectRequiredAction) Set(val *RunObjectRequiredAction) {
	v.value = val
	v.isSet = true
}

func (v NullableRunObjectRequiredAction) IsSet() bool {
	return v.isSet
}

func (v *NullableRunObjectRequiredAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunObjectRequiredAction(val *RunObjectRequiredAction) *NullableRunObjectRequiredAction {
	return &NullableRunObjectRequiredAction{value: val, isSet: true}
}

func (v NullableRunObjectRequiredAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunObjectRequiredAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



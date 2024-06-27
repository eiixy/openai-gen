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

// checks if the RunStepDetailsToolCallsCodeObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStepDetailsToolCallsCodeObject{}

// RunStepDetailsToolCallsCodeObject Details of the Code Interpreter tool call the run step was involved in.
type RunStepDetailsToolCallsCodeObject struct {
	// The ID of the tool call.
	Id string `json:"id"`
	// The type of tool call. This is always going to be `code_interpreter` for this type of tool call.
	Type string `json:"type"`
	CodeInterpreter RunStepDetailsToolCallsCodeObjectCodeInterpreter `json:"code_interpreter"`
}

type _RunStepDetailsToolCallsCodeObject RunStepDetailsToolCallsCodeObject

// NewRunStepDetailsToolCallsCodeObject instantiates a new RunStepDetailsToolCallsCodeObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStepDetailsToolCallsCodeObject(id string, type_ string, codeInterpreter RunStepDetailsToolCallsCodeObjectCodeInterpreter) *RunStepDetailsToolCallsCodeObject {
	this := RunStepDetailsToolCallsCodeObject{}
	this.Id = id
	this.Type = type_
	this.CodeInterpreter = codeInterpreter
	return &this
}

// NewRunStepDetailsToolCallsCodeObjectWithDefaults instantiates a new RunStepDetailsToolCallsCodeObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStepDetailsToolCallsCodeObjectWithDefaults() *RunStepDetailsToolCallsCodeObject {
	this := RunStepDetailsToolCallsCodeObject{}
	return &this
}

// GetId returns the Id field value
func (o *RunStepDetailsToolCallsCodeObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RunStepDetailsToolCallsCodeObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RunStepDetailsToolCallsCodeObject) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *RunStepDetailsToolCallsCodeObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RunStepDetailsToolCallsCodeObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RunStepDetailsToolCallsCodeObject) SetType(v string) {
	o.Type = v
}

// GetCodeInterpreter returns the CodeInterpreter field value
func (o *RunStepDetailsToolCallsCodeObject) GetCodeInterpreter() RunStepDetailsToolCallsCodeObjectCodeInterpreter {
	if o == nil {
		var ret RunStepDetailsToolCallsCodeObjectCodeInterpreter
		return ret
	}

	return o.CodeInterpreter
}

// GetCodeInterpreterOk returns a tuple with the CodeInterpreter field value
// and a boolean to check if the value has been set.
func (o *RunStepDetailsToolCallsCodeObject) GetCodeInterpreterOk() (*RunStepDetailsToolCallsCodeObjectCodeInterpreter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeInterpreter, true
}

// SetCodeInterpreter sets field value
func (o *RunStepDetailsToolCallsCodeObject) SetCodeInterpreter(v RunStepDetailsToolCallsCodeObjectCodeInterpreter) {
	o.CodeInterpreter = v
}

func (o RunStepDetailsToolCallsCodeObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStepDetailsToolCallsCodeObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["code_interpreter"] = o.CodeInterpreter
	return toSerialize, nil
}

func (o *RunStepDetailsToolCallsCodeObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"code_interpreter",
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

	varRunStepDetailsToolCallsCodeObject := _RunStepDetailsToolCallsCodeObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunStepDetailsToolCallsCodeObject)

	if err != nil {
		return err
	}

	*o = RunStepDetailsToolCallsCodeObject(varRunStepDetailsToolCallsCodeObject)

	return err
}

type NullableRunStepDetailsToolCallsCodeObject struct {
	value *RunStepDetailsToolCallsCodeObject
	isSet bool
}

func (v NullableRunStepDetailsToolCallsCodeObject) Get() *RunStepDetailsToolCallsCodeObject {
	return v.value
}

func (v *NullableRunStepDetailsToolCallsCodeObject) Set(val *RunStepDetailsToolCallsCodeObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStepDetailsToolCallsCodeObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStepDetailsToolCallsCodeObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStepDetailsToolCallsCodeObject(val *RunStepDetailsToolCallsCodeObject) *NullableRunStepDetailsToolCallsCodeObject {
	return &NullableRunStepDetailsToolCallsCodeObject{value: val, isSet: true}
}

func (v NullableRunStepDetailsToolCallsCodeObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStepDetailsToolCallsCodeObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject{}

// RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject Text output from the Code Interpreter tool call as part of a run step.
type RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject struct {
	// The index of the output in the outputs array.
	Index int32 `json:"index"`
	// Always `logs`.
	Type string `json:"type"`
	// The text output from the Code Interpreter tool call.
	Logs *string `json:"logs,omitempty"`
}

type _RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject

// NewRunStepDeltaStepDetailsToolCallsCodeOutputLogsObject instantiates a new RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStepDeltaStepDetailsToolCallsCodeOutputLogsObject(index int32, type_ string) *RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject {
	this := RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject{}
	this.Index = index
	this.Type = type_
	return &this
}

// NewRunStepDeltaStepDetailsToolCallsCodeOutputLogsObjectWithDefaults instantiates a new RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStepDeltaStepDetailsToolCallsCodeOutputLogsObjectWithDefaults() *RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject {
	this := RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject{}
	return &this
}

// GetIndex returns the Index field value
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) SetIndex(v int32) {
	o.Index = v
}

// GetType returns the Type field value
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) SetType(v string) {
	o.Type = v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) GetLogs() string {
	if o == nil || IsNil(o.Logs) {
		var ret string
		return ret
	}
	return *o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) GetLogsOk() (*string, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given string and assigns it to the Logs field.
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) SetLogs(v string) {
	o.Logs = &v
}

func (o RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["type"] = o.Type
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	return toSerialize, nil
}

func (o *RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
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

	varRunStepDeltaStepDetailsToolCallsCodeOutputLogsObject := _RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunStepDeltaStepDetailsToolCallsCodeOutputLogsObject)

	if err != nil {
		return err
	}

	*o = RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject(varRunStepDeltaStepDetailsToolCallsCodeOutputLogsObject)

	return err
}

type NullableRunStepDeltaStepDetailsToolCallsCodeOutputLogsObject struct {
	value *RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject
	isSet bool
}

func (v NullableRunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) Get() *RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject {
	return v.value
}

func (v *NullableRunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) Set(val *RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStepDeltaStepDetailsToolCallsCodeOutputLogsObject(val *RunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) *NullableRunStepDeltaStepDetailsToolCallsCodeOutputLogsObject {
	return &NullableRunStepDeltaStepDetailsToolCallsCodeOutputLogsObject{value: val, isSet: true}
}

func (v NullableRunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStepDeltaStepDetailsToolCallsCodeOutputLogsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



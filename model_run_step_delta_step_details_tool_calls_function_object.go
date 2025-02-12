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

// checks if the RunStepDeltaStepDetailsToolCallsFunctionObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStepDeltaStepDetailsToolCallsFunctionObject{}

// RunStepDeltaStepDetailsToolCallsFunctionObject struct for RunStepDeltaStepDetailsToolCallsFunctionObject
type RunStepDeltaStepDetailsToolCallsFunctionObject struct {
	// The index of the tool call in the tool calls array.
	Index int32 `json:"index"`
	// The ID of the tool call object.
	Id *string `json:"id,omitempty"`
	// The type of tool call. This is always going to be `function` for this type of tool call.
	Type string `json:"type"`
	Function *RunStepDeltaStepDetailsToolCallsFunctionObjectFunction `json:"function,omitempty"`
}

type _RunStepDeltaStepDetailsToolCallsFunctionObject RunStepDeltaStepDetailsToolCallsFunctionObject

// NewRunStepDeltaStepDetailsToolCallsFunctionObject instantiates a new RunStepDeltaStepDetailsToolCallsFunctionObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStepDeltaStepDetailsToolCallsFunctionObject(index int32, type_ string) *RunStepDeltaStepDetailsToolCallsFunctionObject {
	this := RunStepDeltaStepDetailsToolCallsFunctionObject{}
	this.Index = index
	this.Type = type_
	return &this
}

// NewRunStepDeltaStepDetailsToolCallsFunctionObjectWithDefaults instantiates a new RunStepDeltaStepDetailsToolCallsFunctionObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStepDeltaStepDetailsToolCallsFunctionObjectWithDefaults() *RunStepDeltaStepDetailsToolCallsFunctionObject {
	this := RunStepDeltaStepDetailsToolCallsFunctionObject{}
	return &this
}

// GetIndex returns the Index field value
func (o *RunStepDeltaStepDetailsToolCallsFunctionObject) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *RunStepDeltaStepDetailsToolCallsFunctionObject) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *RunStepDeltaStepDetailsToolCallsFunctionObject) SetIndex(v int32) {
	o.Index = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RunStepDeltaStepDetailsToolCallsFunctionObject) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStepDeltaStepDetailsToolCallsFunctionObject) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RunStepDeltaStepDetailsToolCallsFunctionObject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RunStepDeltaStepDetailsToolCallsFunctionObject) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value
func (o *RunStepDeltaStepDetailsToolCallsFunctionObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RunStepDeltaStepDetailsToolCallsFunctionObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RunStepDeltaStepDetailsToolCallsFunctionObject) SetType(v string) {
	o.Type = v
}

// GetFunction returns the Function field value if set, zero value otherwise.
func (o *RunStepDeltaStepDetailsToolCallsFunctionObject) GetFunction() RunStepDeltaStepDetailsToolCallsFunctionObjectFunction {
	if o == nil || IsNil(o.Function) {
		var ret RunStepDeltaStepDetailsToolCallsFunctionObjectFunction
		return ret
	}
	return *o.Function
}

// GetFunctionOk returns a tuple with the Function field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStepDeltaStepDetailsToolCallsFunctionObject) GetFunctionOk() (*RunStepDeltaStepDetailsToolCallsFunctionObjectFunction, bool) {
	if o == nil || IsNil(o.Function) {
		return nil, false
	}
	return o.Function, true
}

// HasFunction returns a boolean if a field has been set.
func (o *RunStepDeltaStepDetailsToolCallsFunctionObject) HasFunction() bool {
	if o != nil && !IsNil(o.Function) {
		return true
	}

	return false
}

// SetFunction gets a reference to the given RunStepDeltaStepDetailsToolCallsFunctionObjectFunction and assigns it to the Function field.
func (o *RunStepDeltaStepDetailsToolCallsFunctionObject) SetFunction(v RunStepDeltaStepDetailsToolCallsFunctionObjectFunction) {
	o.Function = &v
}

func (o RunStepDeltaStepDetailsToolCallsFunctionObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStepDeltaStepDetailsToolCallsFunctionObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Function) {
		toSerialize["function"] = o.Function
	}
	return toSerialize, nil
}

func (o *RunStepDeltaStepDetailsToolCallsFunctionObject) UnmarshalJSON(data []byte) (err error) {
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

	varRunStepDeltaStepDetailsToolCallsFunctionObject := _RunStepDeltaStepDetailsToolCallsFunctionObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunStepDeltaStepDetailsToolCallsFunctionObject)

	if err != nil {
		return err
	}

	*o = RunStepDeltaStepDetailsToolCallsFunctionObject(varRunStepDeltaStepDetailsToolCallsFunctionObject)

	return err
}

type NullableRunStepDeltaStepDetailsToolCallsFunctionObject struct {
	value *RunStepDeltaStepDetailsToolCallsFunctionObject
	isSet bool
}

func (v NullableRunStepDeltaStepDetailsToolCallsFunctionObject) Get() *RunStepDeltaStepDetailsToolCallsFunctionObject {
	return v.value
}

func (v *NullableRunStepDeltaStepDetailsToolCallsFunctionObject) Set(val *RunStepDeltaStepDetailsToolCallsFunctionObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStepDeltaStepDetailsToolCallsFunctionObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStepDeltaStepDetailsToolCallsFunctionObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStepDeltaStepDetailsToolCallsFunctionObject(val *RunStepDeltaStepDetailsToolCallsFunctionObject) *NullableRunStepDeltaStepDetailsToolCallsFunctionObject {
	return &NullableRunStepDeltaStepDetailsToolCallsFunctionObject{value: val, isSet: true}
}

func (v NullableRunStepDeltaStepDetailsToolCallsFunctionObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStepDeltaStepDetailsToolCallsFunctionObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the RunToolCallObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunToolCallObject{}

// RunToolCallObject Tool call objects
type RunToolCallObject struct {
	// The ID of the tool call. This ID must be referenced when you submit the tool outputs in using the [Submit tool outputs to run](/docs/api-reference/runs/submitToolOutputs) endpoint.
	Id string `json:"id"`
	// The type of tool call the output is required for. For now, this is always `function`.
	Type string `json:"type"`
	Function RunToolCallObjectFunction `json:"function"`
}

type _RunToolCallObject RunToolCallObject

// NewRunToolCallObject instantiates a new RunToolCallObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunToolCallObject(id string, type_ string, function RunToolCallObjectFunction) *RunToolCallObject {
	this := RunToolCallObject{}
	this.Id = id
	this.Type = type_
	this.Function = function
	return &this
}

// NewRunToolCallObjectWithDefaults instantiates a new RunToolCallObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunToolCallObjectWithDefaults() *RunToolCallObject {
	this := RunToolCallObject{}
	return &this
}

// GetId returns the Id field value
func (o *RunToolCallObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RunToolCallObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RunToolCallObject) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *RunToolCallObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RunToolCallObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RunToolCallObject) SetType(v string) {
	o.Type = v
}

// GetFunction returns the Function field value
func (o *RunToolCallObject) GetFunction() RunToolCallObjectFunction {
	if o == nil {
		var ret RunToolCallObjectFunction
		return ret
	}

	return o.Function
}

// GetFunctionOk returns a tuple with the Function field value
// and a boolean to check if the value has been set.
func (o *RunToolCallObject) GetFunctionOk() (*RunToolCallObjectFunction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Function, true
}

// SetFunction sets field value
func (o *RunToolCallObject) SetFunction(v RunToolCallObjectFunction) {
	o.Function = v
}

func (o RunToolCallObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunToolCallObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["function"] = o.Function
	return toSerialize, nil
}

func (o *RunToolCallObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"function",
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

	varRunToolCallObject := _RunToolCallObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunToolCallObject)

	if err != nil {
		return err
	}

	*o = RunToolCallObject(varRunToolCallObject)

	return err
}

type NullableRunToolCallObject struct {
	value *RunToolCallObject
	isSet bool
}

func (v NullableRunToolCallObject) Get() *RunToolCallObject {
	return v.value
}

func (v *NullableRunToolCallObject) Set(val *RunToolCallObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRunToolCallObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRunToolCallObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunToolCallObject(val *RunToolCallObject) *NullableRunToolCallObject {
	return &NullableRunToolCallObject{value: val, isSet: true}
}

func (v NullableRunToolCallObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunToolCallObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



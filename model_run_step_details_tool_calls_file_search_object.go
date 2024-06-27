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

// checks if the RunStepDetailsToolCallsFileSearchObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStepDetailsToolCallsFileSearchObject{}

// RunStepDetailsToolCallsFileSearchObject struct for RunStepDetailsToolCallsFileSearchObject
type RunStepDetailsToolCallsFileSearchObject struct {
	// The ID of the tool call object.
	Id string `json:"id"`
	// The type of tool call. This is always going to be `file_search` for this type of tool call.
	Type string `json:"type"`
	// For now, this is always going to be an empty object.
	FileSearch map[string]interface{} `json:"file_search"`
}

type _RunStepDetailsToolCallsFileSearchObject RunStepDetailsToolCallsFileSearchObject

// NewRunStepDetailsToolCallsFileSearchObject instantiates a new RunStepDetailsToolCallsFileSearchObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStepDetailsToolCallsFileSearchObject(id string, type_ string, fileSearch map[string]interface{}) *RunStepDetailsToolCallsFileSearchObject {
	this := RunStepDetailsToolCallsFileSearchObject{}
	this.Id = id
	this.Type = type_
	this.FileSearch = fileSearch
	return &this
}

// NewRunStepDetailsToolCallsFileSearchObjectWithDefaults instantiates a new RunStepDetailsToolCallsFileSearchObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStepDetailsToolCallsFileSearchObjectWithDefaults() *RunStepDetailsToolCallsFileSearchObject {
	this := RunStepDetailsToolCallsFileSearchObject{}
	return &this
}

// GetId returns the Id field value
func (o *RunStepDetailsToolCallsFileSearchObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RunStepDetailsToolCallsFileSearchObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RunStepDetailsToolCallsFileSearchObject) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *RunStepDetailsToolCallsFileSearchObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RunStepDetailsToolCallsFileSearchObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RunStepDetailsToolCallsFileSearchObject) SetType(v string) {
	o.Type = v
}

// GetFileSearch returns the FileSearch field value
func (o *RunStepDetailsToolCallsFileSearchObject) GetFileSearch() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.FileSearch
}

// GetFileSearchOk returns a tuple with the FileSearch field value
// and a boolean to check if the value has been set.
func (o *RunStepDetailsToolCallsFileSearchObject) GetFileSearchOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.FileSearch, true
}

// SetFileSearch sets field value
func (o *RunStepDetailsToolCallsFileSearchObject) SetFileSearch(v map[string]interface{}) {
	o.FileSearch = v
}

func (o RunStepDetailsToolCallsFileSearchObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStepDetailsToolCallsFileSearchObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["file_search"] = o.FileSearch
	return toSerialize, nil
}

func (o *RunStepDetailsToolCallsFileSearchObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"file_search",
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

	varRunStepDetailsToolCallsFileSearchObject := _RunStepDetailsToolCallsFileSearchObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunStepDetailsToolCallsFileSearchObject)

	if err != nil {
		return err
	}

	*o = RunStepDetailsToolCallsFileSearchObject(varRunStepDetailsToolCallsFileSearchObject)

	return err
}

type NullableRunStepDetailsToolCallsFileSearchObject struct {
	value *RunStepDetailsToolCallsFileSearchObject
	isSet bool
}

func (v NullableRunStepDetailsToolCallsFileSearchObject) Get() *RunStepDetailsToolCallsFileSearchObject {
	return v.value
}

func (v *NullableRunStepDetailsToolCallsFileSearchObject) Set(val *RunStepDetailsToolCallsFileSearchObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStepDetailsToolCallsFileSearchObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStepDetailsToolCallsFileSearchObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStepDetailsToolCallsFileSearchObject(val *RunStepDetailsToolCallsFileSearchObject) *NullableRunStepDetailsToolCallsFileSearchObject {
	return &NullableRunStepDetailsToolCallsFileSearchObject{value: val, isSet: true}
}

func (v NullableRunStepDetailsToolCallsFileSearchObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStepDetailsToolCallsFileSearchObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the ModifyAssistantRequestToolResourcesCodeInterpreter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyAssistantRequestToolResourcesCodeInterpreter{}

// ModifyAssistantRequestToolResourcesCodeInterpreter struct for ModifyAssistantRequestToolResourcesCodeInterpreter
type ModifyAssistantRequestToolResourcesCodeInterpreter struct {
	// Overrides the list of [file](/docs/api-reference/files) IDs made available to the `code_interpreter` tool. There can be a maximum of 20 files associated with the tool. 
	FileIds []string `json:"file_ids,omitempty"`
}

// NewModifyAssistantRequestToolResourcesCodeInterpreter instantiates a new ModifyAssistantRequestToolResourcesCodeInterpreter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyAssistantRequestToolResourcesCodeInterpreter() *ModifyAssistantRequestToolResourcesCodeInterpreter {
	this := ModifyAssistantRequestToolResourcesCodeInterpreter{}
	return &this
}

// NewModifyAssistantRequestToolResourcesCodeInterpreterWithDefaults instantiates a new ModifyAssistantRequestToolResourcesCodeInterpreter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyAssistantRequestToolResourcesCodeInterpreterWithDefaults() *ModifyAssistantRequestToolResourcesCodeInterpreter {
	this := ModifyAssistantRequestToolResourcesCodeInterpreter{}
	return &this
}

// GetFileIds returns the FileIds field value if set, zero value otherwise.
func (o *ModifyAssistantRequestToolResourcesCodeInterpreter) GetFileIds() []string {
	if o == nil || IsNil(o.FileIds) {
		var ret []string
		return ret
	}
	return o.FileIds
}

// GetFileIdsOk returns a tuple with the FileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyAssistantRequestToolResourcesCodeInterpreter) GetFileIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.FileIds) {
		return nil, false
	}
	return o.FileIds, true
}

// HasFileIds returns a boolean if a field has been set.
func (o *ModifyAssistantRequestToolResourcesCodeInterpreter) HasFileIds() bool {
	if o != nil && !IsNil(o.FileIds) {
		return true
	}

	return false
}

// SetFileIds gets a reference to the given []string and assigns it to the FileIds field.
func (o *ModifyAssistantRequestToolResourcesCodeInterpreter) SetFileIds(v []string) {
	o.FileIds = v
}

func (o ModifyAssistantRequestToolResourcesCodeInterpreter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyAssistantRequestToolResourcesCodeInterpreter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileIds) {
		toSerialize["file_ids"] = o.FileIds
	}
	return toSerialize, nil
}

type NullableModifyAssistantRequestToolResourcesCodeInterpreter struct {
	value *ModifyAssistantRequestToolResourcesCodeInterpreter
	isSet bool
}

func (v NullableModifyAssistantRequestToolResourcesCodeInterpreter) Get() *ModifyAssistantRequestToolResourcesCodeInterpreter {
	return v.value
}

func (v *NullableModifyAssistantRequestToolResourcesCodeInterpreter) Set(val *ModifyAssistantRequestToolResourcesCodeInterpreter) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyAssistantRequestToolResourcesCodeInterpreter) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyAssistantRequestToolResourcesCodeInterpreter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyAssistantRequestToolResourcesCodeInterpreter(val *ModifyAssistantRequestToolResourcesCodeInterpreter) *NullableModifyAssistantRequestToolResourcesCodeInterpreter {
	return &NullableModifyAssistantRequestToolResourcesCodeInterpreter{value: val, isSet: true}
}

func (v NullableModifyAssistantRequestToolResourcesCodeInterpreter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyAssistantRequestToolResourcesCodeInterpreter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



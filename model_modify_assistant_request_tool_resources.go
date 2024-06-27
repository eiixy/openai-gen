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

// checks if the ModifyAssistantRequestToolResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyAssistantRequestToolResources{}

// ModifyAssistantRequestToolResources A set of resources that are used by the assistant's tools. The resources are specific to the type of tool. For example, the `code_interpreter` tool requires a list of file IDs, while the `file_search` tool requires a list of vector store IDs. 
type ModifyAssistantRequestToolResources struct {
	CodeInterpreter *ModifyAssistantRequestToolResourcesCodeInterpreter `json:"code_interpreter,omitempty"`
	FileSearch *ModifyAssistantRequestToolResourcesFileSearch `json:"file_search,omitempty"`
}

// NewModifyAssistantRequestToolResources instantiates a new ModifyAssistantRequestToolResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyAssistantRequestToolResources() *ModifyAssistantRequestToolResources {
	this := ModifyAssistantRequestToolResources{}
	return &this
}

// NewModifyAssistantRequestToolResourcesWithDefaults instantiates a new ModifyAssistantRequestToolResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyAssistantRequestToolResourcesWithDefaults() *ModifyAssistantRequestToolResources {
	this := ModifyAssistantRequestToolResources{}
	return &this
}

// GetCodeInterpreter returns the CodeInterpreter field value if set, zero value otherwise.
func (o *ModifyAssistantRequestToolResources) GetCodeInterpreter() ModifyAssistantRequestToolResourcesCodeInterpreter {
	if o == nil || IsNil(o.CodeInterpreter) {
		var ret ModifyAssistantRequestToolResourcesCodeInterpreter
		return ret
	}
	return *o.CodeInterpreter
}

// GetCodeInterpreterOk returns a tuple with the CodeInterpreter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyAssistantRequestToolResources) GetCodeInterpreterOk() (*ModifyAssistantRequestToolResourcesCodeInterpreter, bool) {
	if o == nil || IsNil(o.CodeInterpreter) {
		return nil, false
	}
	return o.CodeInterpreter, true
}

// HasCodeInterpreter returns a boolean if a field has been set.
func (o *ModifyAssistantRequestToolResources) HasCodeInterpreter() bool {
	if o != nil && !IsNil(o.CodeInterpreter) {
		return true
	}

	return false
}

// SetCodeInterpreter gets a reference to the given ModifyAssistantRequestToolResourcesCodeInterpreter and assigns it to the CodeInterpreter field.
func (o *ModifyAssistantRequestToolResources) SetCodeInterpreter(v ModifyAssistantRequestToolResourcesCodeInterpreter) {
	o.CodeInterpreter = &v
}

// GetFileSearch returns the FileSearch field value if set, zero value otherwise.
func (o *ModifyAssistantRequestToolResources) GetFileSearch() ModifyAssistantRequestToolResourcesFileSearch {
	if o == nil || IsNil(o.FileSearch) {
		var ret ModifyAssistantRequestToolResourcesFileSearch
		return ret
	}
	return *o.FileSearch
}

// GetFileSearchOk returns a tuple with the FileSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyAssistantRequestToolResources) GetFileSearchOk() (*ModifyAssistantRequestToolResourcesFileSearch, bool) {
	if o == nil || IsNil(o.FileSearch) {
		return nil, false
	}
	return o.FileSearch, true
}

// HasFileSearch returns a boolean if a field has been set.
func (o *ModifyAssistantRequestToolResources) HasFileSearch() bool {
	if o != nil && !IsNil(o.FileSearch) {
		return true
	}

	return false
}

// SetFileSearch gets a reference to the given ModifyAssistantRequestToolResourcesFileSearch and assigns it to the FileSearch field.
func (o *ModifyAssistantRequestToolResources) SetFileSearch(v ModifyAssistantRequestToolResourcesFileSearch) {
	o.FileSearch = &v
}

func (o ModifyAssistantRequestToolResources) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyAssistantRequestToolResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CodeInterpreter) {
		toSerialize["code_interpreter"] = o.CodeInterpreter
	}
	if !IsNil(o.FileSearch) {
		toSerialize["file_search"] = o.FileSearch
	}
	return toSerialize, nil
}

type NullableModifyAssistantRequestToolResources struct {
	value *ModifyAssistantRequestToolResources
	isSet bool
}

func (v NullableModifyAssistantRequestToolResources) Get() *ModifyAssistantRequestToolResources {
	return v.value
}

func (v *NullableModifyAssistantRequestToolResources) Set(val *ModifyAssistantRequestToolResources) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyAssistantRequestToolResources) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyAssistantRequestToolResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyAssistantRequestToolResources(val *ModifyAssistantRequestToolResources) *NullableModifyAssistantRequestToolResources {
	return &NullableModifyAssistantRequestToolResources{value: val, isSet: true}
}

func (v NullableModifyAssistantRequestToolResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyAssistantRequestToolResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the AssistantObjectToolResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssistantObjectToolResources{}

// AssistantObjectToolResources A set of resources that are used by the assistant's tools. The resources are specific to the type of tool. For example, the `code_interpreter` tool requires a list of file IDs, while the `file_search` tool requires a list of vector store IDs. 
type AssistantObjectToolResources struct {
	CodeInterpreter *AssistantObjectToolResourcesCodeInterpreter `json:"code_interpreter,omitempty"`
	FileSearch *AssistantObjectToolResourcesFileSearch `json:"file_search,omitempty"`
}

// NewAssistantObjectToolResources instantiates a new AssistantObjectToolResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssistantObjectToolResources() *AssistantObjectToolResources {
	this := AssistantObjectToolResources{}
	return &this
}

// NewAssistantObjectToolResourcesWithDefaults instantiates a new AssistantObjectToolResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssistantObjectToolResourcesWithDefaults() *AssistantObjectToolResources {
	this := AssistantObjectToolResources{}
	return &this
}

// GetCodeInterpreter returns the CodeInterpreter field value if set, zero value otherwise.
func (o *AssistantObjectToolResources) GetCodeInterpreter() AssistantObjectToolResourcesCodeInterpreter {
	if o == nil || IsNil(o.CodeInterpreter) {
		var ret AssistantObjectToolResourcesCodeInterpreter
		return ret
	}
	return *o.CodeInterpreter
}

// GetCodeInterpreterOk returns a tuple with the CodeInterpreter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssistantObjectToolResources) GetCodeInterpreterOk() (*AssistantObjectToolResourcesCodeInterpreter, bool) {
	if o == nil || IsNil(o.CodeInterpreter) {
		return nil, false
	}
	return o.CodeInterpreter, true
}

// HasCodeInterpreter returns a boolean if a field has been set.
func (o *AssistantObjectToolResources) HasCodeInterpreter() bool {
	if o != nil && !IsNil(o.CodeInterpreter) {
		return true
	}

	return false
}

// SetCodeInterpreter gets a reference to the given AssistantObjectToolResourcesCodeInterpreter and assigns it to the CodeInterpreter field.
func (o *AssistantObjectToolResources) SetCodeInterpreter(v AssistantObjectToolResourcesCodeInterpreter) {
	o.CodeInterpreter = &v
}

// GetFileSearch returns the FileSearch field value if set, zero value otherwise.
func (o *AssistantObjectToolResources) GetFileSearch() AssistantObjectToolResourcesFileSearch {
	if o == nil || IsNil(o.FileSearch) {
		var ret AssistantObjectToolResourcesFileSearch
		return ret
	}
	return *o.FileSearch
}

// GetFileSearchOk returns a tuple with the FileSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssistantObjectToolResources) GetFileSearchOk() (*AssistantObjectToolResourcesFileSearch, bool) {
	if o == nil || IsNil(o.FileSearch) {
		return nil, false
	}
	return o.FileSearch, true
}

// HasFileSearch returns a boolean if a field has been set.
func (o *AssistantObjectToolResources) HasFileSearch() bool {
	if o != nil && !IsNil(o.FileSearch) {
		return true
	}

	return false
}

// SetFileSearch gets a reference to the given AssistantObjectToolResourcesFileSearch and assigns it to the FileSearch field.
func (o *AssistantObjectToolResources) SetFileSearch(v AssistantObjectToolResourcesFileSearch) {
	o.FileSearch = &v
}

func (o AssistantObjectToolResources) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssistantObjectToolResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CodeInterpreter) {
		toSerialize["code_interpreter"] = o.CodeInterpreter
	}
	if !IsNil(o.FileSearch) {
		toSerialize["file_search"] = o.FileSearch
	}
	return toSerialize, nil
}

type NullableAssistantObjectToolResources struct {
	value *AssistantObjectToolResources
	isSet bool
}

func (v NullableAssistantObjectToolResources) Get() *AssistantObjectToolResources {
	return v.value
}

func (v *NullableAssistantObjectToolResources) Set(val *AssistantObjectToolResources) {
	v.value = val
	v.isSet = true
}

func (v NullableAssistantObjectToolResources) IsSet() bool {
	return v.isSet
}

func (v *NullableAssistantObjectToolResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssistantObjectToolResources(val *AssistantObjectToolResources) *NullableAssistantObjectToolResources {
	return &NullableAssistantObjectToolResources{value: val, isSet: true}
}

func (v NullableAssistantObjectToolResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssistantObjectToolResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



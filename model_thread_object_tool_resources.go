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

// checks if the ThreadObjectToolResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreadObjectToolResources{}

// ThreadObjectToolResources A set of resources that are made available to the assistant's tools in this thread. The resources are specific to the type of tool. For example, the `code_interpreter` tool requires a list of file IDs, while the `file_search` tool requires a list of vector store IDs. 
type ThreadObjectToolResources struct {
	CodeInterpreter *CreateAssistantRequestToolResourcesCodeInterpreter `json:"code_interpreter,omitempty"`
	FileSearch *ThreadObjectToolResourcesFileSearch `json:"file_search,omitempty"`
}

// NewThreadObjectToolResources instantiates a new ThreadObjectToolResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreadObjectToolResources() *ThreadObjectToolResources {
	this := ThreadObjectToolResources{}
	return &this
}

// NewThreadObjectToolResourcesWithDefaults instantiates a new ThreadObjectToolResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreadObjectToolResourcesWithDefaults() *ThreadObjectToolResources {
	this := ThreadObjectToolResources{}
	return &this
}

// GetCodeInterpreter returns the CodeInterpreter field value if set, zero value otherwise.
func (o *ThreadObjectToolResources) GetCodeInterpreter() CreateAssistantRequestToolResourcesCodeInterpreter {
	if o == nil || IsNil(o.CodeInterpreter) {
		var ret CreateAssistantRequestToolResourcesCodeInterpreter
		return ret
	}
	return *o.CodeInterpreter
}

// GetCodeInterpreterOk returns a tuple with the CodeInterpreter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadObjectToolResources) GetCodeInterpreterOk() (*CreateAssistantRequestToolResourcesCodeInterpreter, bool) {
	if o == nil || IsNil(o.CodeInterpreter) {
		return nil, false
	}
	return o.CodeInterpreter, true
}

// HasCodeInterpreter returns a boolean if a field has been set.
func (o *ThreadObjectToolResources) HasCodeInterpreter() bool {
	if o != nil && !IsNil(o.CodeInterpreter) {
		return true
	}

	return false
}

// SetCodeInterpreter gets a reference to the given CreateAssistantRequestToolResourcesCodeInterpreter and assigns it to the CodeInterpreter field.
func (o *ThreadObjectToolResources) SetCodeInterpreter(v CreateAssistantRequestToolResourcesCodeInterpreter) {
	o.CodeInterpreter = &v
}

// GetFileSearch returns the FileSearch field value if set, zero value otherwise.
func (o *ThreadObjectToolResources) GetFileSearch() ThreadObjectToolResourcesFileSearch {
	if o == nil || IsNil(o.FileSearch) {
		var ret ThreadObjectToolResourcesFileSearch
		return ret
	}
	return *o.FileSearch
}

// GetFileSearchOk returns a tuple with the FileSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadObjectToolResources) GetFileSearchOk() (*ThreadObjectToolResourcesFileSearch, bool) {
	if o == nil || IsNil(o.FileSearch) {
		return nil, false
	}
	return o.FileSearch, true
}

// HasFileSearch returns a boolean if a field has been set.
func (o *ThreadObjectToolResources) HasFileSearch() bool {
	if o != nil && !IsNil(o.FileSearch) {
		return true
	}

	return false
}

// SetFileSearch gets a reference to the given ThreadObjectToolResourcesFileSearch and assigns it to the FileSearch field.
func (o *ThreadObjectToolResources) SetFileSearch(v ThreadObjectToolResourcesFileSearch) {
	o.FileSearch = &v
}

func (o ThreadObjectToolResources) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreadObjectToolResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CodeInterpreter) {
		toSerialize["code_interpreter"] = o.CodeInterpreter
	}
	if !IsNil(o.FileSearch) {
		toSerialize["file_search"] = o.FileSearch
	}
	return toSerialize, nil
}

type NullableThreadObjectToolResources struct {
	value *ThreadObjectToolResources
	isSet bool
}

func (v NullableThreadObjectToolResources) Get() *ThreadObjectToolResources {
	return v.value
}

func (v *NullableThreadObjectToolResources) Set(val *ThreadObjectToolResources) {
	v.value = val
	v.isSet = true
}

func (v NullableThreadObjectToolResources) IsSet() bool {
	return v.isSet
}

func (v *NullableThreadObjectToolResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreadObjectToolResources(val *ThreadObjectToolResources) *NullableThreadObjectToolResources {
	return &NullableThreadObjectToolResources{value: val, isSet: true}
}

func (v NullableThreadObjectToolResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreadObjectToolResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



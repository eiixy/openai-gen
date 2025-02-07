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

// checks if the ModifyRunRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyRunRequest{}

// ModifyRunRequest struct for ModifyRunRequest
type ModifyRunRequest struct {
	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long. 
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewModifyRunRequest instantiates a new ModifyRunRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyRunRequest() *ModifyRunRequest {
	this := ModifyRunRequest{}
	return &this
}

// NewModifyRunRequestWithDefaults instantiates a new ModifyRunRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyRunRequestWithDefaults() *ModifyRunRequest {
	this := ModifyRunRequest{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModifyRunRequest) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModifyRunRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ModifyRunRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ModifyRunRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ModifyRunRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyRunRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableModifyRunRequest struct {
	value *ModifyRunRequest
	isSet bool
}

func (v NullableModifyRunRequest) Get() *ModifyRunRequest {
	return v.value
}

func (v *NullableModifyRunRequest) Set(val *ModifyRunRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyRunRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyRunRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyRunRequest(val *ModifyRunRequest) *NullableModifyRunRequest {
	return &NullableModifyRunRequest{value: val, isSet: true}
}

func (v NullableModifyRunRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyRunRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



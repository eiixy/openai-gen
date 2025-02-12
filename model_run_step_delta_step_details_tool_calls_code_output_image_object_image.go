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

// checks if the RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage{}

// RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage struct for RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage
type RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage struct {
	// The [file](/docs/api-reference/files) ID of the image.
	FileId *string `json:"file_id,omitempty"`
}

// NewRunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage instantiates a new RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage() *RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage {
	this := RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage{}
	return &this
}

// NewRunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImageWithDefaults instantiates a new RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImageWithDefaults() *RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage {
	this := RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage{}
	return &this
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage) GetFileId() string {
	if o == nil || IsNil(o.FileId) {
		var ret string
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage) GetFileIdOk() (*string, bool) {
	if o == nil || IsNil(o.FileId) {
		return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage) HasFileId() bool {
	if o != nil && !IsNil(o.FileId) {
		return true
	}

	return false
}

// SetFileId gets a reference to the given string and assigns it to the FileId field.
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage) SetFileId(v string) {
	o.FileId = &v
}

func (o RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileId) {
		toSerialize["file_id"] = o.FileId
	}
	return toSerialize, nil
}

type NullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage struct {
	value *RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage
	isSet bool
}

func (v NullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage) Get() *RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage {
	return v.value
}

func (v *NullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage) Set(val *RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage(val *RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage) *NullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage {
	return &NullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage{value: val, isSet: true}
}

func (v NullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



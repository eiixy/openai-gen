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

// checks if the RunStepDeltaStepDetailsToolCallsCodeOutputImageObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStepDeltaStepDetailsToolCallsCodeOutputImageObject{}

// RunStepDeltaStepDetailsToolCallsCodeOutputImageObject struct for RunStepDeltaStepDetailsToolCallsCodeOutputImageObject
type RunStepDeltaStepDetailsToolCallsCodeOutputImageObject struct {
	// The index of the output in the outputs array.
	Index int32 `json:"index"`
	// Always `image`.
	Type string `json:"type"`
	Image *RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage `json:"image,omitempty"`
}

type _RunStepDeltaStepDetailsToolCallsCodeOutputImageObject RunStepDeltaStepDetailsToolCallsCodeOutputImageObject

// NewRunStepDeltaStepDetailsToolCallsCodeOutputImageObject instantiates a new RunStepDeltaStepDetailsToolCallsCodeOutputImageObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStepDeltaStepDetailsToolCallsCodeOutputImageObject(index int32, type_ string) *RunStepDeltaStepDetailsToolCallsCodeOutputImageObject {
	this := RunStepDeltaStepDetailsToolCallsCodeOutputImageObject{}
	this.Index = index
	this.Type = type_
	return &this
}

// NewRunStepDeltaStepDetailsToolCallsCodeOutputImageObjectWithDefaults instantiates a new RunStepDeltaStepDetailsToolCallsCodeOutputImageObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStepDeltaStepDetailsToolCallsCodeOutputImageObjectWithDefaults() *RunStepDeltaStepDetailsToolCallsCodeOutputImageObject {
	this := RunStepDeltaStepDetailsToolCallsCodeOutputImageObject{}
	return &this
}

// GetIndex returns the Index field value
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputImageObject) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputImageObject) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputImageObject) SetIndex(v int32) {
	o.Index = v
}

// GetType returns the Type field value
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputImageObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputImageObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputImageObject) SetType(v string) {
	o.Type = v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputImageObject) GetImage() RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage {
	if o == nil || IsNil(o.Image) {
		var ret RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputImageObject) GetImageOk() (*RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputImageObject) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage and assigns it to the Image field.
func (o *RunStepDeltaStepDetailsToolCallsCodeOutputImageObject) SetImage(v RunStepDeltaStepDetailsToolCallsCodeOutputImageObjectImage) {
	o.Image = &v
}

func (o RunStepDeltaStepDetailsToolCallsCodeOutputImageObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStepDeltaStepDetailsToolCallsCodeOutputImageObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["type"] = o.Type
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	return toSerialize, nil
}

func (o *RunStepDeltaStepDetailsToolCallsCodeOutputImageObject) UnmarshalJSON(data []byte) (err error) {
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

	varRunStepDeltaStepDetailsToolCallsCodeOutputImageObject := _RunStepDeltaStepDetailsToolCallsCodeOutputImageObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunStepDeltaStepDetailsToolCallsCodeOutputImageObject)

	if err != nil {
		return err
	}

	*o = RunStepDeltaStepDetailsToolCallsCodeOutputImageObject(varRunStepDeltaStepDetailsToolCallsCodeOutputImageObject)

	return err
}

type NullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObject struct {
	value *RunStepDeltaStepDetailsToolCallsCodeOutputImageObject
	isSet bool
}

func (v NullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObject) Get() *RunStepDeltaStepDetailsToolCallsCodeOutputImageObject {
	return v.value
}

func (v *NullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObject) Set(val *RunStepDeltaStepDetailsToolCallsCodeOutputImageObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObject(val *RunStepDeltaStepDetailsToolCallsCodeOutputImageObject) *NullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObject {
	return &NullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObject{value: val, isSet: true}
}

func (v NullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStepDeltaStepDetailsToolCallsCodeOutputImageObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



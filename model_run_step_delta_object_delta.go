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

// checks if the RunStepDeltaObjectDelta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStepDeltaObjectDelta{}

// RunStepDeltaObjectDelta The delta containing the fields that have changed on the run step.
type RunStepDeltaObjectDelta struct {
	StepDetails *RunStepDeltaObjectDeltaStepDetails `json:"step_details,omitempty"`
}

// NewRunStepDeltaObjectDelta instantiates a new RunStepDeltaObjectDelta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStepDeltaObjectDelta() *RunStepDeltaObjectDelta {
	this := RunStepDeltaObjectDelta{}
	return &this
}

// NewRunStepDeltaObjectDeltaWithDefaults instantiates a new RunStepDeltaObjectDelta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStepDeltaObjectDeltaWithDefaults() *RunStepDeltaObjectDelta {
	this := RunStepDeltaObjectDelta{}
	return &this
}

// GetStepDetails returns the StepDetails field value if set, zero value otherwise.
func (o *RunStepDeltaObjectDelta) GetStepDetails() RunStepDeltaObjectDeltaStepDetails {
	if o == nil || IsNil(o.StepDetails) {
		var ret RunStepDeltaObjectDeltaStepDetails
		return ret
	}
	return *o.StepDetails
}

// GetStepDetailsOk returns a tuple with the StepDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStepDeltaObjectDelta) GetStepDetailsOk() (*RunStepDeltaObjectDeltaStepDetails, bool) {
	if o == nil || IsNil(o.StepDetails) {
		return nil, false
	}
	return o.StepDetails, true
}

// HasStepDetails returns a boolean if a field has been set.
func (o *RunStepDeltaObjectDelta) HasStepDetails() bool {
	if o != nil && !IsNil(o.StepDetails) {
		return true
	}

	return false
}

// SetStepDetails gets a reference to the given RunStepDeltaObjectDeltaStepDetails and assigns it to the StepDetails field.
func (o *RunStepDeltaObjectDelta) SetStepDetails(v RunStepDeltaObjectDeltaStepDetails) {
	o.StepDetails = &v
}

func (o RunStepDeltaObjectDelta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStepDeltaObjectDelta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StepDetails) {
		toSerialize["step_details"] = o.StepDetails
	}
	return toSerialize, nil
}

type NullableRunStepDeltaObjectDelta struct {
	value *RunStepDeltaObjectDelta
	isSet bool
}

func (v NullableRunStepDeltaObjectDelta) Get() *RunStepDeltaObjectDelta {
	return v.value
}

func (v *NullableRunStepDeltaObjectDelta) Set(val *RunStepDeltaObjectDelta) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStepDeltaObjectDelta) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStepDeltaObjectDelta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStepDeltaObjectDelta(val *RunStepDeltaObjectDelta) *NullableRunStepDeltaObjectDelta {
	return &NullableRunStepDeltaObjectDelta{value: val, isSet: true}
}

func (v NullableRunStepDeltaObjectDelta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStepDeltaObjectDelta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



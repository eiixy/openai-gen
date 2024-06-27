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

// checks if the FineTuningJobHyperparameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FineTuningJobHyperparameters{}

// FineTuningJobHyperparameters The hyperparameters used for the fine-tuning job. See the [fine-tuning guide](/docs/guides/fine-tuning) for more details.
type FineTuningJobHyperparameters struct {
	NEpochs FineTuningJobHyperparametersNEpochs `json:"n_epochs"`
}

type _FineTuningJobHyperparameters FineTuningJobHyperparameters

// NewFineTuningJobHyperparameters instantiates a new FineTuningJobHyperparameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFineTuningJobHyperparameters(nEpochs FineTuningJobHyperparametersNEpochs) *FineTuningJobHyperparameters {
	this := FineTuningJobHyperparameters{}
	this.NEpochs = nEpochs
	return &this
}

// NewFineTuningJobHyperparametersWithDefaults instantiates a new FineTuningJobHyperparameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFineTuningJobHyperparametersWithDefaults() *FineTuningJobHyperparameters {
	this := FineTuningJobHyperparameters{}
	var nEpochs FineTuningJobHyperparametersNEpochs = auto
	this.NEpochs = nEpochs
	return &this
}

// GetNEpochs returns the NEpochs field value
func (o *FineTuningJobHyperparameters) GetNEpochs() FineTuningJobHyperparametersNEpochs {
	if o == nil {
		var ret FineTuningJobHyperparametersNEpochs
		return ret
	}

	return o.NEpochs
}

// GetNEpochsOk returns a tuple with the NEpochs field value
// and a boolean to check if the value has been set.
func (o *FineTuningJobHyperparameters) GetNEpochsOk() (*FineTuningJobHyperparametersNEpochs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NEpochs, true
}

// SetNEpochs sets field value
func (o *FineTuningJobHyperparameters) SetNEpochs(v FineTuningJobHyperparametersNEpochs) {
	o.NEpochs = v
}

func (o FineTuningJobHyperparameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FineTuningJobHyperparameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["n_epochs"] = o.NEpochs
	return toSerialize, nil
}

func (o *FineTuningJobHyperparameters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"n_epochs",
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

	varFineTuningJobHyperparameters := _FineTuningJobHyperparameters{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFineTuningJobHyperparameters)

	if err != nil {
		return err
	}

	*o = FineTuningJobHyperparameters(varFineTuningJobHyperparameters)

	return err
}

type NullableFineTuningJobHyperparameters struct {
	value *FineTuningJobHyperparameters
	isSet bool
}

func (v NullableFineTuningJobHyperparameters) Get() *FineTuningJobHyperparameters {
	return v.value
}

func (v *NullableFineTuningJobHyperparameters) Set(val *FineTuningJobHyperparameters) {
	v.value = val
	v.isSet = true
}

func (v NullableFineTuningJobHyperparameters) IsSet() bool {
	return v.isSet
}

func (v *NullableFineTuningJobHyperparameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFineTuningJobHyperparameters(val *FineTuningJobHyperparameters) *NullableFineTuningJobHyperparameters {
	return &NullableFineTuningJobHyperparameters{value: val, isSet: true}
}

func (v NullableFineTuningJobHyperparameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFineTuningJobHyperparameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



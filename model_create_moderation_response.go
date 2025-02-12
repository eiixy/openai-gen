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

// checks if the CreateModerationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateModerationResponse{}

// CreateModerationResponse Represents if a given text input is potentially harmful.
type CreateModerationResponse struct {
	// The unique identifier for the moderation request.
	Id string `json:"id"`
	// The model used to generate the moderation results.
	Model string `json:"model"`
	// A list of moderation objects.
	Results []CreateModerationResponseResultsInner `json:"results"`
}

type _CreateModerationResponse CreateModerationResponse

// NewCreateModerationResponse instantiates a new CreateModerationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateModerationResponse(id string, model string, results []CreateModerationResponseResultsInner) *CreateModerationResponse {
	this := CreateModerationResponse{}
	this.Id = id
	this.Model = model
	this.Results = results
	return &this
}

// NewCreateModerationResponseWithDefaults instantiates a new CreateModerationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateModerationResponseWithDefaults() *CreateModerationResponse {
	this := CreateModerationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CreateModerationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateModerationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateModerationResponse) SetId(v string) {
	o.Id = v
}

// GetModel returns the Model field value
func (o *CreateModerationResponse) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *CreateModerationResponse) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *CreateModerationResponse) SetModel(v string) {
	o.Model = v
}

// GetResults returns the Results field value
func (o *CreateModerationResponse) GetResults() []CreateModerationResponseResultsInner {
	if o == nil {
		var ret []CreateModerationResponseResultsInner
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CreateModerationResponse) GetResultsOk() ([]CreateModerationResponseResultsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CreateModerationResponse) SetResults(v []CreateModerationResponseResultsInner) {
	o.Results = v
}

func (o CreateModerationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateModerationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["model"] = o.Model
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *CreateModerationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"model",
		"results",
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

	varCreateModerationResponse := _CreateModerationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateModerationResponse)

	if err != nil {
		return err
	}

	*o = CreateModerationResponse(varCreateModerationResponse)

	return err
}

type NullableCreateModerationResponse struct {
	value *CreateModerationResponse
	isSet bool
}

func (v NullableCreateModerationResponse) Get() *CreateModerationResponse {
	return v.value
}

func (v *NullableCreateModerationResponse) Set(val *CreateModerationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateModerationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateModerationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateModerationResponse(val *CreateModerationResponse) *NullableCreateModerationResponse {
	return &NullableCreateModerationResponse{value: val, isSet: true}
}

func (v NullableCreateModerationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateModerationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



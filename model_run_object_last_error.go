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

// checks if the RunObjectLastError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunObjectLastError{}

// RunObjectLastError The last error associated with this run. Will be `null` if there are no errors.
type RunObjectLastError struct {
	// One of `server_error`, `rate_limit_exceeded`, or `invalid_prompt`.
	Code string `json:"code"`
	// A human-readable description of the error.
	Message string `json:"message"`
}

type _RunObjectLastError RunObjectLastError

// NewRunObjectLastError instantiates a new RunObjectLastError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunObjectLastError(code string, message string) *RunObjectLastError {
	this := RunObjectLastError{}
	this.Code = code
	this.Message = message
	return &this
}

// NewRunObjectLastErrorWithDefaults instantiates a new RunObjectLastError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunObjectLastErrorWithDefaults() *RunObjectLastError {
	this := RunObjectLastError{}
	return &this
}

// GetCode returns the Code field value
func (o *RunObjectLastError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *RunObjectLastError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *RunObjectLastError) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *RunObjectLastError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *RunObjectLastError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *RunObjectLastError) SetMessage(v string) {
	o.Message = v
}

func (o RunObjectLastError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunObjectLastError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *RunObjectLastError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
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

	varRunObjectLastError := _RunObjectLastError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunObjectLastError)

	if err != nil {
		return err
	}

	*o = RunObjectLastError(varRunObjectLastError)

	return err
}

type NullableRunObjectLastError struct {
	value *RunObjectLastError
	isSet bool
}

func (v NullableRunObjectLastError) Get() *RunObjectLastError {
	return v.value
}

func (v *NullableRunObjectLastError) Set(val *RunObjectLastError) {
	v.value = val
	v.isSet = true
}

func (v NullableRunObjectLastError) IsSet() bool {
	return v.isSet
}

func (v *NullableRunObjectLastError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunObjectLastError(val *RunObjectLastError) *NullableRunObjectLastError {
	return &NullableRunObjectLastError{value: val, isSet: true}
}

func (v NullableRunObjectLastError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunObjectLastError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



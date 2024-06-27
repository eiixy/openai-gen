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

// checks if the BatchRequestInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchRequestInput{}

// BatchRequestInput The per-line object of the batch input file
type BatchRequestInput struct {
	// A developer-provided per-request id that will be used to match outputs to inputs. Must be unique for each request in a batch.
	CustomId *string `json:"custom_id,omitempty"`
	// The HTTP method to be used for the request. Currently only `POST` is supported.
	Method *string `json:"method,omitempty"`
	// The OpenAI API relative URL to be used for the request. Currently `/v1/chat/completions`, `/v1/embeddings`, and `/v1/completions` are supported.
	Url *string `json:"url,omitempty"`
}

// NewBatchRequestInput instantiates a new BatchRequestInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchRequestInput() *BatchRequestInput {
	this := BatchRequestInput{}
	return &this
}

// NewBatchRequestInputWithDefaults instantiates a new BatchRequestInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchRequestInputWithDefaults() *BatchRequestInput {
	this := BatchRequestInput{}
	return &this
}

// GetCustomId returns the CustomId field value if set, zero value otherwise.
func (o *BatchRequestInput) GetCustomId() string {
	if o == nil || IsNil(o.CustomId) {
		var ret string
		return ret
	}
	return *o.CustomId
}

// GetCustomIdOk returns a tuple with the CustomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRequestInput) GetCustomIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomId) {
		return nil, false
	}
	return o.CustomId, true
}

// HasCustomId returns a boolean if a field has been set.
func (o *BatchRequestInput) HasCustomId() bool {
	if o != nil && !IsNil(o.CustomId) {
		return true
	}

	return false
}

// SetCustomId gets a reference to the given string and assigns it to the CustomId field.
func (o *BatchRequestInput) SetCustomId(v string) {
	o.CustomId = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *BatchRequestInput) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRequestInput) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *BatchRequestInput) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *BatchRequestInput) SetMethod(v string) {
	o.Method = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *BatchRequestInput) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRequestInput) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *BatchRequestInput) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *BatchRequestInput) SetUrl(v string) {
	o.Url = &v
}

func (o BatchRequestInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchRequestInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomId) {
		toSerialize["custom_id"] = o.CustomId
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableBatchRequestInput struct {
	value *BatchRequestInput
	isSet bool
}

func (v NullableBatchRequestInput) Get() *BatchRequestInput {
	return v.value
}

func (v *NullableBatchRequestInput) Set(val *BatchRequestInput) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchRequestInput) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchRequestInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchRequestInput(val *BatchRequestInput) *NullableBatchRequestInput {
	return &NullableBatchRequestInput{value: val, isSet: true}
}

func (v NullableBatchRequestInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchRequestInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the CreateChatCompletionFunctionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateChatCompletionFunctionResponse{}

// CreateChatCompletionFunctionResponse Represents a chat completion response returned by model, based on the provided input.
type CreateChatCompletionFunctionResponse struct {
	// A unique identifier for the chat completion.
	Id string `json:"id"`
	// A list of chat completion choices. Can be more than one if `n` is greater than 1.
	Choices []CreateChatCompletionFunctionResponseChoicesInner `json:"choices"`
	// The Unix timestamp (in seconds) of when the chat completion was created.
	Created int32 `json:"created"`
	// The model used for the chat completion.
	Model string `json:"model"`
	// This fingerprint represents the backend configuration that the model runs with.  Can be used in conjunction with the `seed` request parameter to understand when backend changes have been made that might impact determinism. 
	SystemFingerprint *string `json:"system_fingerprint,omitempty"`
	// The object type, which is always `chat.completion`.
	Object string `json:"object"`
	Usage *CompletionUsage `json:"usage,omitempty"`
}

type _CreateChatCompletionFunctionResponse CreateChatCompletionFunctionResponse

// NewCreateChatCompletionFunctionResponse instantiates a new CreateChatCompletionFunctionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateChatCompletionFunctionResponse(id string, choices []CreateChatCompletionFunctionResponseChoicesInner, created int32, model string, object string) *CreateChatCompletionFunctionResponse {
	this := CreateChatCompletionFunctionResponse{}
	this.Id = id
	this.Choices = choices
	this.Created = created
	this.Model = model
	this.Object = object
	return &this
}

// NewCreateChatCompletionFunctionResponseWithDefaults instantiates a new CreateChatCompletionFunctionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateChatCompletionFunctionResponseWithDefaults() *CreateChatCompletionFunctionResponse {
	this := CreateChatCompletionFunctionResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CreateChatCompletionFunctionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionFunctionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateChatCompletionFunctionResponse) SetId(v string) {
	o.Id = v
}

// GetChoices returns the Choices field value
func (o *CreateChatCompletionFunctionResponse) GetChoices() []CreateChatCompletionFunctionResponseChoicesInner {
	if o == nil {
		var ret []CreateChatCompletionFunctionResponseChoicesInner
		return ret
	}

	return o.Choices
}

// GetChoicesOk returns a tuple with the Choices field value
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionFunctionResponse) GetChoicesOk() ([]CreateChatCompletionFunctionResponseChoicesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Choices, true
}

// SetChoices sets field value
func (o *CreateChatCompletionFunctionResponse) SetChoices(v []CreateChatCompletionFunctionResponseChoicesInner) {
	o.Choices = v
}

// GetCreated returns the Created field value
func (o *CreateChatCompletionFunctionResponse) GetCreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionFunctionResponse) GetCreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *CreateChatCompletionFunctionResponse) SetCreated(v int32) {
	o.Created = v
}

// GetModel returns the Model field value
func (o *CreateChatCompletionFunctionResponse) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionFunctionResponse) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *CreateChatCompletionFunctionResponse) SetModel(v string) {
	o.Model = v
}

// GetSystemFingerprint returns the SystemFingerprint field value if set, zero value otherwise.
func (o *CreateChatCompletionFunctionResponse) GetSystemFingerprint() string {
	if o == nil || IsNil(o.SystemFingerprint) {
		var ret string
		return ret
	}
	return *o.SystemFingerprint
}

// GetSystemFingerprintOk returns a tuple with the SystemFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionFunctionResponse) GetSystemFingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.SystemFingerprint) {
		return nil, false
	}
	return o.SystemFingerprint, true
}

// HasSystemFingerprint returns a boolean if a field has been set.
func (o *CreateChatCompletionFunctionResponse) HasSystemFingerprint() bool {
	if o != nil && !IsNil(o.SystemFingerprint) {
		return true
	}

	return false
}

// SetSystemFingerprint gets a reference to the given string and assigns it to the SystemFingerprint field.
func (o *CreateChatCompletionFunctionResponse) SetSystemFingerprint(v string) {
	o.SystemFingerprint = &v
}

// GetObject returns the Object field value
func (o *CreateChatCompletionFunctionResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionFunctionResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CreateChatCompletionFunctionResponse) SetObject(v string) {
	o.Object = v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *CreateChatCompletionFunctionResponse) GetUsage() CompletionUsage {
	if o == nil || IsNil(o.Usage) {
		var ret CompletionUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionFunctionResponse) GetUsageOk() (*CompletionUsage, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *CreateChatCompletionFunctionResponse) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given CompletionUsage and assigns it to the Usage field.
func (o *CreateChatCompletionFunctionResponse) SetUsage(v CompletionUsage) {
	o.Usage = &v
}

func (o CreateChatCompletionFunctionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateChatCompletionFunctionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["choices"] = o.Choices
	toSerialize["created"] = o.Created
	toSerialize["model"] = o.Model
	if !IsNil(o.SystemFingerprint) {
		toSerialize["system_fingerprint"] = o.SystemFingerprint
	}
	toSerialize["object"] = o.Object
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

func (o *CreateChatCompletionFunctionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"choices",
		"created",
		"model",
		"object",
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

	varCreateChatCompletionFunctionResponse := _CreateChatCompletionFunctionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateChatCompletionFunctionResponse)

	if err != nil {
		return err
	}

	*o = CreateChatCompletionFunctionResponse(varCreateChatCompletionFunctionResponse)

	return err
}

type NullableCreateChatCompletionFunctionResponse struct {
	value *CreateChatCompletionFunctionResponse
	isSet bool
}

func (v NullableCreateChatCompletionFunctionResponse) Get() *CreateChatCompletionFunctionResponse {
	return v.value
}

func (v *NullableCreateChatCompletionFunctionResponse) Set(val *CreateChatCompletionFunctionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateChatCompletionFunctionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateChatCompletionFunctionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateChatCompletionFunctionResponse(val *CreateChatCompletionFunctionResponse) *NullableCreateChatCompletionFunctionResponse {
	return &NullableCreateChatCompletionFunctionResponse{value: val, isSet: true}
}

func (v NullableCreateChatCompletionFunctionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateChatCompletionFunctionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



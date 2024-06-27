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

// checks if the CreateCompletionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCompletionResponse{}

// CreateCompletionResponse Represents a completion response from the API. Note: both the streamed and non-streamed response objects share the same shape (unlike the chat endpoint). 
type CreateCompletionResponse struct {
	// A unique identifier for the completion.
	Id string `json:"id"`
	// The list of completion choices the model generated for the input prompt.
	Choices []CreateCompletionResponseChoicesInner `json:"choices"`
	// The Unix timestamp (in seconds) of when the completion was created.
	Created int32 `json:"created"`
	// The model used for completion.
	Model string `json:"model"`
	// This fingerprint represents the backend configuration that the model runs with.  Can be used in conjunction with the `seed` request parameter to understand when backend changes have been made that might impact determinism. 
	SystemFingerprint *string `json:"system_fingerprint,omitempty"`
	// The object type, which is always \"text_completion\"
	Object string `json:"object"`
	Usage *CompletionUsage `json:"usage,omitempty"`
}

type _CreateCompletionResponse CreateCompletionResponse

// NewCreateCompletionResponse instantiates a new CreateCompletionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCompletionResponse(id string, choices []CreateCompletionResponseChoicesInner, created int32, model string, object string) *CreateCompletionResponse {
	this := CreateCompletionResponse{}
	this.Id = id
	this.Choices = choices
	this.Created = created
	this.Model = model
	this.Object = object
	return &this
}

// NewCreateCompletionResponseWithDefaults instantiates a new CreateCompletionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCompletionResponseWithDefaults() *CreateCompletionResponse {
	this := CreateCompletionResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CreateCompletionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateCompletionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateCompletionResponse) SetId(v string) {
	o.Id = v
}

// GetChoices returns the Choices field value
func (o *CreateCompletionResponse) GetChoices() []CreateCompletionResponseChoicesInner {
	if o == nil {
		var ret []CreateCompletionResponseChoicesInner
		return ret
	}

	return o.Choices
}

// GetChoicesOk returns a tuple with the Choices field value
// and a boolean to check if the value has been set.
func (o *CreateCompletionResponse) GetChoicesOk() ([]CreateCompletionResponseChoicesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Choices, true
}

// SetChoices sets field value
func (o *CreateCompletionResponse) SetChoices(v []CreateCompletionResponseChoicesInner) {
	o.Choices = v
}

// GetCreated returns the Created field value
func (o *CreateCompletionResponse) GetCreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *CreateCompletionResponse) GetCreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *CreateCompletionResponse) SetCreated(v int32) {
	o.Created = v
}

// GetModel returns the Model field value
func (o *CreateCompletionResponse) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *CreateCompletionResponse) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *CreateCompletionResponse) SetModel(v string) {
	o.Model = v
}

// GetSystemFingerprint returns the SystemFingerprint field value if set, zero value otherwise.
func (o *CreateCompletionResponse) GetSystemFingerprint() string {
	if o == nil || IsNil(o.SystemFingerprint) {
		var ret string
		return ret
	}
	return *o.SystemFingerprint
}

// GetSystemFingerprintOk returns a tuple with the SystemFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCompletionResponse) GetSystemFingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.SystemFingerprint) {
		return nil, false
	}
	return o.SystemFingerprint, true
}

// HasSystemFingerprint returns a boolean if a field has been set.
func (o *CreateCompletionResponse) HasSystemFingerprint() bool {
	if o != nil && !IsNil(o.SystemFingerprint) {
		return true
	}

	return false
}

// SetSystemFingerprint gets a reference to the given string and assigns it to the SystemFingerprint field.
func (o *CreateCompletionResponse) SetSystemFingerprint(v string) {
	o.SystemFingerprint = &v
}

// GetObject returns the Object field value
func (o *CreateCompletionResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CreateCompletionResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CreateCompletionResponse) SetObject(v string) {
	o.Object = v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *CreateCompletionResponse) GetUsage() CompletionUsage {
	if o == nil || IsNil(o.Usage) {
		var ret CompletionUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCompletionResponse) GetUsageOk() (*CompletionUsage, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *CreateCompletionResponse) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given CompletionUsage and assigns it to the Usage field.
func (o *CreateCompletionResponse) SetUsage(v CompletionUsage) {
	o.Usage = &v
}

func (o CreateCompletionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCompletionResponse) ToMap() (map[string]interface{}, error) {
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

func (o *CreateCompletionResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCreateCompletionResponse := _CreateCompletionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateCompletionResponse)

	if err != nil {
		return err
	}

	*o = CreateCompletionResponse(varCreateCompletionResponse)

	return err
}

type NullableCreateCompletionResponse struct {
	value *CreateCompletionResponse
	isSet bool
}

func (v NullableCreateCompletionResponse) Get() *CreateCompletionResponse {
	return v.value
}

func (v *NullableCreateCompletionResponse) Set(val *CreateCompletionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCompletionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCompletionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCompletionResponse(val *CreateCompletionResponse) *NullableCreateCompletionResponse {
	return &NullableCreateCompletionResponse{value: val, isSet: true}
}

func (v NullableCreateCompletionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCompletionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



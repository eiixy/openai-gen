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

// checks if the CreateEmbeddingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEmbeddingResponse{}

// CreateEmbeddingResponse struct for CreateEmbeddingResponse
type CreateEmbeddingResponse struct {
	// The list of embeddings generated by the model.
	Data []Embedding `json:"data"`
	// The name of the model used to generate the embedding.
	Model string `json:"model"`
	// The object type, which is always \"list\".
	Object string `json:"object"`
	Usage CreateEmbeddingResponseUsage `json:"usage"`
}

type _CreateEmbeddingResponse CreateEmbeddingResponse

// NewCreateEmbeddingResponse instantiates a new CreateEmbeddingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEmbeddingResponse(data []Embedding, model string, object string, usage CreateEmbeddingResponseUsage) *CreateEmbeddingResponse {
	this := CreateEmbeddingResponse{}
	this.Data = data
	this.Model = model
	this.Object = object
	this.Usage = usage
	return &this
}

// NewCreateEmbeddingResponseWithDefaults instantiates a new CreateEmbeddingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEmbeddingResponseWithDefaults() *CreateEmbeddingResponse {
	this := CreateEmbeddingResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CreateEmbeddingResponse) GetData() []Embedding {
	if o == nil {
		var ret []Embedding
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateEmbeddingResponse) GetDataOk() ([]Embedding, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CreateEmbeddingResponse) SetData(v []Embedding) {
	o.Data = v
}

// GetModel returns the Model field value
func (o *CreateEmbeddingResponse) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *CreateEmbeddingResponse) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *CreateEmbeddingResponse) SetModel(v string) {
	o.Model = v
}

// GetObject returns the Object field value
func (o *CreateEmbeddingResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CreateEmbeddingResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CreateEmbeddingResponse) SetObject(v string) {
	o.Object = v
}

// GetUsage returns the Usage field value
func (o *CreateEmbeddingResponse) GetUsage() CreateEmbeddingResponseUsage {
	if o == nil {
		var ret CreateEmbeddingResponseUsage
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *CreateEmbeddingResponse) GetUsageOk() (*CreateEmbeddingResponseUsage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value
func (o *CreateEmbeddingResponse) SetUsage(v CreateEmbeddingResponseUsage) {
	o.Usage = v
}

func (o CreateEmbeddingResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEmbeddingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["model"] = o.Model
	toSerialize["object"] = o.Object
	toSerialize["usage"] = o.Usage
	return toSerialize, nil
}

func (o *CreateEmbeddingResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"model",
		"object",
		"usage",
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

	varCreateEmbeddingResponse := _CreateEmbeddingResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateEmbeddingResponse)

	if err != nil {
		return err
	}

	*o = CreateEmbeddingResponse(varCreateEmbeddingResponse)

	return err
}

type NullableCreateEmbeddingResponse struct {
	value *CreateEmbeddingResponse
	isSet bool
}

func (v NullableCreateEmbeddingResponse) Get() *CreateEmbeddingResponse {
	return v.value
}

func (v *NullableCreateEmbeddingResponse) Set(val *CreateEmbeddingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEmbeddingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEmbeddingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEmbeddingResponse(val *CreateEmbeddingResponse) *NullableCreateEmbeddingResponse {
	return &NullableCreateEmbeddingResponse{value: val, isSet: true}
}

func (v NullableCreateEmbeddingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEmbeddingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



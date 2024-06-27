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

// checks if the Embedding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Embedding{}

// Embedding Represents an embedding vector returned by embedding endpoint. 
type Embedding struct {
	// The index of the embedding in the list of embeddings.
	Index int32 `json:"index"`
	// The embedding vector, which is a list of floats. The length of vector depends on the model as listed in the [embedding guide](/docs/guides/embeddings). 
	Embedding []float32 `json:"embedding"`
	// The object type, which is always \"embedding\".
	Object string `json:"object"`
}

type _Embedding Embedding

// NewEmbedding instantiates a new Embedding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbedding(index int32, embedding []float32, object string) *Embedding {
	this := Embedding{}
	this.Index = index
	this.Embedding = embedding
	this.Object = object
	return &this
}

// NewEmbeddingWithDefaults instantiates a new Embedding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddingWithDefaults() *Embedding {
	this := Embedding{}
	return &this
}

// GetIndex returns the Index field value
func (o *Embedding) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *Embedding) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *Embedding) SetIndex(v int32) {
	o.Index = v
}

// GetEmbedding returns the Embedding field value
func (o *Embedding) GetEmbedding() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.Embedding
}

// GetEmbeddingOk returns a tuple with the Embedding field value
// and a boolean to check if the value has been set.
func (o *Embedding) GetEmbeddingOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Embedding, true
}

// SetEmbedding sets field value
func (o *Embedding) SetEmbedding(v []float32) {
	o.Embedding = v
}

// GetObject returns the Object field value
func (o *Embedding) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Embedding) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Embedding) SetObject(v string) {
	o.Object = v
}

func (o Embedding) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Embedding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["embedding"] = o.Embedding
	toSerialize["object"] = o.Object
	return toSerialize, nil
}

func (o *Embedding) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
		"embedding",
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

	varEmbedding := _Embedding{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmbedding)

	if err != nil {
		return err
	}

	*o = Embedding(varEmbedding)

	return err
}

type NullableEmbedding struct {
	value *Embedding
	isSet bool
}

func (v NullableEmbedding) Get() *Embedding {
	return v.value
}

func (v *NullableEmbedding) Set(val *Embedding) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbedding) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbedding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbedding(val *Embedding) *NullableEmbedding {
	return &NullableEmbedding{value: val, isSet: true}
}

func (v NullableEmbedding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbedding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



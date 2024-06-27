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

// checks if the VectorStoreExpirationAfter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VectorStoreExpirationAfter{}

// VectorStoreExpirationAfter The expiration policy for a vector store.
type VectorStoreExpirationAfter struct {
	// Anchor timestamp after which the expiration policy applies. Supported anchors: `last_active_at`.
	Anchor string `json:"anchor"`
	// The number of days after the anchor time that the vector store will expire.
	Days int32 `json:"days"`
}

type _VectorStoreExpirationAfter VectorStoreExpirationAfter

// NewVectorStoreExpirationAfter instantiates a new VectorStoreExpirationAfter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVectorStoreExpirationAfter(anchor string, days int32) *VectorStoreExpirationAfter {
	this := VectorStoreExpirationAfter{}
	this.Anchor = anchor
	this.Days = days
	return &this
}

// NewVectorStoreExpirationAfterWithDefaults instantiates a new VectorStoreExpirationAfter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVectorStoreExpirationAfterWithDefaults() *VectorStoreExpirationAfter {
	this := VectorStoreExpirationAfter{}
	return &this
}

// GetAnchor returns the Anchor field value
func (o *VectorStoreExpirationAfter) GetAnchor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Anchor
}

// GetAnchorOk returns a tuple with the Anchor field value
// and a boolean to check if the value has been set.
func (o *VectorStoreExpirationAfter) GetAnchorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Anchor, true
}

// SetAnchor sets field value
func (o *VectorStoreExpirationAfter) SetAnchor(v string) {
	o.Anchor = v
}

// GetDays returns the Days field value
func (o *VectorStoreExpirationAfter) GetDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Days
}

// GetDaysOk returns a tuple with the Days field value
// and a boolean to check if the value has been set.
func (o *VectorStoreExpirationAfter) GetDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Days, true
}

// SetDays sets field value
func (o *VectorStoreExpirationAfter) SetDays(v int32) {
	o.Days = v
}

func (o VectorStoreExpirationAfter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VectorStoreExpirationAfter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["anchor"] = o.Anchor
	toSerialize["days"] = o.Days
	return toSerialize, nil
}

func (o *VectorStoreExpirationAfter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"anchor",
		"days",
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

	varVectorStoreExpirationAfter := _VectorStoreExpirationAfter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVectorStoreExpirationAfter)

	if err != nil {
		return err
	}

	*o = VectorStoreExpirationAfter(varVectorStoreExpirationAfter)

	return err
}

type NullableVectorStoreExpirationAfter struct {
	value *VectorStoreExpirationAfter
	isSet bool
}

func (v NullableVectorStoreExpirationAfter) Get() *VectorStoreExpirationAfter {
	return v.value
}

func (v *NullableVectorStoreExpirationAfter) Set(val *VectorStoreExpirationAfter) {
	v.value = val
	v.isSet = true
}

func (v NullableVectorStoreExpirationAfter) IsSet() bool {
	return v.isSet
}

func (v *NullableVectorStoreExpirationAfter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVectorStoreExpirationAfter(val *VectorStoreExpirationAfter) *NullableVectorStoreExpirationAfter {
	return &NullableVectorStoreExpirationAfter{value: val, isSet: true}
}

func (v NullableVectorStoreExpirationAfter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVectorStoreExpirationAfter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



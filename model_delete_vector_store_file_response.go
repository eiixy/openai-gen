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

// checks if the DeleteVectorStoreFileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteVectorStoreFileResponse{}

// DeleteVectorStoreFileResponse struct for DeleteVectorStoreFileResponse
type DeleteVectorStoreFileResponse struct {
	Id string `json:"id"`
	Deleted bool `json:"deleted"`
	Object string `json:"object"`
}

type _DeleteVectorStoreFileResponse DeleteVectorStoreFileResponse

// NewDeleteVectorStoreFileResponse instantiates a new DeleteVectorStoreFileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteVectorStoreFileResponse(id string, deleted bool, object string) *DeleteVectorStoreFileResponse {
	this := DeleteVectorStoreFileResponse{}
	this.Id = id
	this.Deleted = deleted
	this.Object = object
	return &this
}

// NewDeleteVectorStoreFileResponseWithDefaults instantiates a new DeleteVectorStoreFileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteVectorStoreFileResponseWithDefaults() *DeleteVectorStoreFileResponse {
	this := DeleteVectorStoreFileResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DeleteVectorStoreFileResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeleteVectorStoreFileResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeleteVectorStoreFileResponse) SetId(v string) {
	o.Id = v
}

// GetDeleted returns the Deleted field value
func (o *DeleteVectorStoreFileResponse) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *DeleteVectorStoreFileResponse) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *DeleteVectorStoreFileResponse) SetDeleted(v bool) {
	o.Deleted = v
}

// GetObject returns the Object field value
func (o *DeleteVectorStoreFileResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *DeleteVectorStoreFileResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *DeleteVectorStoreFileResponse) SetObject(v string) {
	o.Object = v
}

func (o DeleteVectorStoreFileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteVectorStoreFileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["deleted"] = o.Deleted
	toSerialize["object"] = o.Object
	return toSerialize, nil
}

func (o *DeleteVectorStoreFileResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"deleted",
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

	varDeleteVectorStoreFileResponse := _DeleteVectorStoreFileResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteVectorStoreFileResponse)

	if err != nil {
		return err
	}

	*o = DeleteVectorStoreFileResponse(varDeleteVectorStoreFileResponse)

	return err
}

type NullableDeleteVectorStoreFileResponse struct {
	value *DeleteVectorStoreFileResponse
	isSet bool
}

func (v NullableDeleteVectorStoreFileResponse) Get() *DeleteVectorStoreFileResponse {
	return v.value
}

func (v *NullableDeleteVectorStoreFileResponse) Set(val *DeleteVectorStoreFileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteVectorStoreFileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteVectorStoreFileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteVectorStoreFileResponse(val *DeleteVectorStoreFileResponse) *NullableDeleteVectorStoreFileResponse {
	return &NullableDeleteVectorStoreFileResponse{value: val, isSet: true}
}

func (v NullableDeleteVectorStoreFileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteVectorStoreFileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



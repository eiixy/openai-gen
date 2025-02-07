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

// checks if the BatchRequestOutputResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchRequestOutputResponse{}

// BatchRequestOutputResponse struct for BatchRequestOutputResponse
type BatchRequestOutputResponse struct {
	// The HTTP status code of the response
	StatusCode *int32 `json:"status_code,omitempty"`
	// An unique identifier for the OpenAI API request. Please include this request ID when contacting support.
	RequestId *string `json:"request_id,omitempty"`
	// The JSON body of the response
	Body map[string]interface{} `json:"body,omitempty"`
}

// NewBatchRequestOutputResponse instantiates a new BatchRequestOutputResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchRequestOutputResponse() *BatchRequestOutputResponse {
	this := BatchRequestOutputResponse{}
	return &this
}

// NewBatchRequestOutputResponseWithDefaults instantiates a new BatchRequestOutputResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchRequestOutputResponseWithDefaults() *BatchRequestOutputResponse {
	this := BatchRequestOutputResponse{}
	return &this
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *BatchRequestOutputResponse) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRequestOutputResponse) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *BatchRequestOutputResponse) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *BatchRequestOutputResponse) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *BatchRequestOutputResponse) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRequestOutputResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *BatchRequestOutputResponse) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *BatchRequestOutputResponse) SetRequestId(v string) {
	o.RequestId = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BatchRequestOutputResponse) GetBody() map[string]interface{} {
	if o == nil || IsNil(o.Body) {
		var ret map[string]interface{}
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRequestOutputResponse) GetBodyOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Body) {
		return map[string]interface{}{}, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BatchRequestOutputResponse) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given map[string]interface{} and assigns it to the Body field.
func (o *BatchRequestOutputResponse) SetBody(v map[string]interface{}) {
	o.Body = v
}

func (o BatchRequestOutputResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchRequestOutputResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	return toSerialize, nil
}

type NullableBatchRequestOutputResponse struct {
	value *BatchRequestOutputResponse
	isSet bool
}

func (v NullableBatchRequestOutputResponse) Get() *BatchRequestOutputResponse {
	return v.value
}

func (v *NullableBatchRequestOutputResponse) Set(val *BatchRequestOutputResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchRequestOutputResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchRequestOutputResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchRequestOutputResponse(val *BatchRequestOutputResponse) *NullableBatchRequestOutputResponse {
	return &NullableBatchRequestOutputResponse{value: val, isSet: true}
}

func (v NullableBatchRequestOutputResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchRequestOutputResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the CreateBatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBatchRequest{}

// CreateBatchRequest struct for CreateBatchRequest
type CreateBatchRequest struct {
	// The ID of an uploaded file that contains requests for the new batch.  See [upload file](/docs/api-reference/files/create) for how to upload a file.  Your input file must be formatted as a [JSONL file](/docs/api-reference/batch/request-input), and must be uploaded with the purpose `batch`. The file can contain up to 50,000 requests, and can be up to 100 MB in size. 
	InputFileId string `json:"input_file_id"`
	// The endpoint to be used for all requests in the batch. Currently `/v1/chat/completions`, `/v1/embeddings`, and `/v1/completions` are supported. Note that `/v1/embeddings` batches are also restricted to a maximum of 50,000 embedding inputs across all requests in the batch.
	Endpoint string `json:"endpoint"`
	// The time frame within which the batch should be processed. Currently only `24h` is supported.
	CompletionWindow string `json:"completion_window"`
	// Optional custom metadata for the batch.
	Metadata map[string]string `json:"metadata,omitempty"`
}

type _CreateBatchRequest CreateBatchRequest

// NewCreateBatchRequest instantiates a new CreateBatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBatchRequest(inputFileId string, endpoint string, completionWindow string) *CreateBatchRequest {
	this := CreateBatchRequest{}
	this.InputFileId = inputFileId
	this.Endpoint = endpoint
	this.CompletionWindow = completionWindow
	return &this
}

// NewCreateBatchRequestWithDefaults instantiates a new CreateBatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBatchRequestWithDefaults() *CreateBatchRequest {
	this := CreateBatchRequest{}
	return &this
}

// GetInputFileId returns the InputFileId field value
func (o *CreateBatchRequest) GetInputFileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InputFileId
}

// GetInputFileIdOk returns a tuple with the InputFileId field value
// and a boolean to check if the value has been set.
func (o *CreateBatchRequest) GetInputFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputFileId, true
}

// SetInputFileId sets field value
func (o *CreateBatchRequest) SetInputFileId(v string) {
	o.InputFileId = v
}

// GetEndpoint returns the Endpoint field value
func (o *CreateBatchRequest) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *CreateBatchRequest) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *CreateBatchRequest) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetCompletionWindow returns the CompletionWindow field value
func (o *CreateBatchRequest) GetCompletionWindow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompletionWindow
}

// GetCompletionWindowOk returns a tuple with the CompletionWindow field value
// and a boolean to check if the value has been set.
func (o *CreateBatchRequest) GetCompletionWindowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletionWindow, true
}

// SetCompletionWindow sets field value
func (o *CreateBatchRequest) SetCompletionWindow(v string) {
	o.CompletionWindow = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateBatchRequest) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateBatchRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateBatchRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CreateBatchRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}

func (o CreateBatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["input_file_id"] = o.InputFileId
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["completion_window"] = o.CompletionWindow
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *CreateBatchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"input_file_id",
		"endpoint",
		"completion_window",
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

	varCreateBatchRequest := _CreateBatchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateBatchRequest)

	if err != nil {
		return err
	}

	*o = CreateBatchRequest(varCreateBatchRequest)

	return err
}

type NullableCreateBatchRequest struct {
	value *CreateBatchRequest
	isSet bool
}

func (v NullableCreateBatchRequest) Get() *CreateBatchRequest {
	return v.value
}

func (v *NullableCreateBatchRequest) Set(val *CreateBatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBatchRequest(val *CreateBatchRequest) *NullableCreateBatchRequest {
	return &NullableCreateBatchRequest{value: val, isSet: true}
}

func (v NullableCreateBatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



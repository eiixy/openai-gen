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

// checks if the CreateFineTuningJobRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFineTuningJobRequest{}

// CreateFineTuningJobRequest struct for CreateFineTuningJobRequest
type CreateFineTuningJobRequest struct {
	Model CreateFineTuningJobRequestModel `json:"model"`
	// The ID of an uploaded file that contains training data.  See [upload file](/docs/api-reference/files/create) for how to upload a file.  Your dataset must be formatted as a JSONL file. Additionally, you must upload your file with the purpose `fine-tune`.  The contents of the file should differ depending on if the model uses the [chat](/docs/api-reference/fine-tuning/chat-input) or [completions](/docs/api-reference/fine-tuning/completions-input) format.  See the [fine-tuning guide](/docs/guides/fine-tuning) for more details. 
	TrainingFile string `json:"training_file"`
	Hyperparameters *CreateFineTuningJobRequestHyperparameters `json:"hyperparameters,omitempty"`
	// A string of up to 18 characters that will be added to your fine-tuned model name.  For example, a `suffix` of \"custom-model-name\" would produce a model name like `ft:gpt-3.5-turbo:openai:custom-model-name:7p4lURel`. 
	Suffix NullableString `json:"suffix,omitempty"`
	// The ID of an uploaded file that contains validation data.  If you provide this file, the data is used to generate validation metrics periodically during fine-tuning. These metrics can be viewed in the fine-tuning results file. The same data should not be present in both train and validation files.  Your dataset must be formatted as a JSONL file. You must upload your file with the purpose `fine-tune`.  See the [fine-tuning guide](/docs/guides/fine-tuning) for more details. 
	ValidationFile NullableString `json:"validation_file,omitempty"`
	// A list of integrations to enable for your fine-tuning job.
	Integrations []CreateFineTuningJobRequestIntegrationsInner `json:"integrations,omitempty"`
	// The seed controls the reproducibility of the job. Passing in the same seed and job parameters should produce the same results, but may differ in rare cases. If a seed is not specified, one will be generated for you. 
	Seed NullableInt32 `json:"seed,omitempty"`
}

type _CreateFineTuningJobRequest CreateFineTuningJobRequest

// NewCreateFineTuningJobRequest instantiates a new CreateFineTuningJobRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFineTuningJobRequest(model CreateFineTuningJobRequestModel, trainingFile string) *CreateFineTuningJobRequest {
	this := CreateFineTuningJobRequest{}
	this.Model = model
	this.TrainingFile = trainingFile
	return &this
}

// NewCreateFineTuningJobRequestWithDefaults instantiates a new CreateFineTuningJobRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFineTuningJobRequestWithDefaults() *CreateFineTuningJobRequest {
	this := CreateFineTuningJobRequest{}
	return &this
}

// GetModel returns the Model field value
func (o *CreateFineTuningJobRequest) GetModel() CreateFineTuningJobRequestModel {
	if o == nil {
		var ret CreateFineTuningJobRequestModel
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *CreateFineTuningJobRequest) GetModelOk() (*CreateFineTuningJobRequestModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *CreateFineTuningJobRequest) SetModel(v CreateFineTuningJobRequestModel) {
	o.Model = v
}

// GetTrainingFile returns the TrainingFile field value
func (o *CreateFineTuningJobRequest) GetTrainingFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrainingFile
}

// GetTrainingFileOk returns a tuple with the TrainingFile field value
// and a boolean to check if the value has been set.
func (o *CreateFineTuningJobRequest) GetTrainingFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrainingFile, true
}

// SetTrainingFile sets field value
func (o *CreateFineTuningJobRequest) SetTrainingFile(v string) {
	o.TrainingFile = v
}

// GetHyperparameters returns the Hyperparameters field value if set, zero value otherwise.
func (o *CreateFineTuningJobRequest) GetHyperparameters() CreateFineTuningJobRequestHyperparameters {
	if o == nil || IsNil(o.Hyperparameters) {
		var ret CreateFineTuningJobRequestHyperparameters
		return ret
	}
	return *o.Hyperparameters
}

// GetHyperparametersOk returns a tuple with the Hyperparameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFineTuningJobRequest) GetHyperparametersOk() (*CreateFineTuningJobRequestHyperparameters, bool) {
	if o == nil || IsNil(o.Hyperparameters) {
		return nil, false
	}
	return o.Hyperparameters, true
}

// HasHyperparameters returns a boolean if a field has been set.
func (o *CreateFineTuningJobRequest) HasHyperparameters() bool {
	if o != nil && !IsNil(o.Hyperparameters) {
		return true
	}

	return false
}

// SetHyperparameters gets a reference to the given CreateFineTuningJobRequestHyperparameters and assigns it to the Hyperparameters field.
func (o *CreateFineTuningJobRequest) SetHyperparameters(v CreateFineTuningJobRequestHyperparameters) {
	o.Hyperparameters = &v
}

// GetSuffix returns the Suffix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFineTuningJobRequest) GetSuffix() string {
	if o == nil || IsNil(o.Suffix.Get()) {
		var ret string
		return ret
	}
	return *o.Suffix.Get()
}

// GetSuffixOk returns a tuple with the Suffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFineTuningJobRequest) GetSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Suffix.Get(), o.Suffix.IsSet()
}

// HasSuffix returns a boolean if a field has been set.
func (o *CreateFineTuningJobRequest) HasSuffix() bool {
	if o != nil && o.Suffix.IsSet() {
		return true
	}

	return false
}

// SetSuffix gets a reference to the given NullableString and assigns it to the Suffix field.
func (o *CreateFineTuningJobRequest) SetSuffix(v string) {
	o.Suffix.Set(&v)
}
// SetSuffixNil sets the value for Suffix to be an explicit nil
func (o *CreateFineTuningJobRequest) SetSuffixNil() {
	o.Suffix.Set(nil)
}

// UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
func (o *CreateFineTuningJobRequest) UnsetSuffix() {
	o.Suffix.Unset()
}

// GetValidationFile returns the ValidationFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFineTuningJobRequest) GetValidationFile() string {
	if o == nil || IsNil(o.ValidationFile.Get()) {
		var ret string
		return ret
	}
	return *o.ValidationFile.Get()
}

// GetValidationFileOk returns a tuple with the ValidationFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFineTuningJobRequest) GetValidationFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidationFile.Get(), o.ValidationFile.IsSet()
}

// HasValidationFile returns a boolean if a field has been set.
func (o *CreateFineTuningJobRequest) HasValidationFile() bool {
	if o != nil && o.ValidationFile.IsSet() {
		return true
	}

	return false
}

// SetValidationFile gets a reference to the given NullableString and assigns it to the ValidationFile field.
func (o *CreateFineTuningJobRequest) SetValidationFile(v string) {
	o.ValidationFile.Set(&v)
}
// SetValidationFileNil sets the value for ValidationFile to be an explicit nil
func (o *CreateFineTuningJobRequest) SetValidationFileNil() {
	o.ValidationFile.Set(nil)
}

// UnsetValidationFile ensures that no value is present for ValidationFile, not even an explicit nil
func (o *CreateFineTuningJobRequest) UnsetValidationFile() {
	o.ValidationFile.Unset()
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFineTuningJobRequest) GetIntegrations() []CreateFineTuningJobRequestIntegrationsInner {
	if o == nil {
		var ret []CreateFineTuningJobRequestIntegrationsInner
		return ret
	}
	return o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFineTuningJobRequest) GetIntegrationsOk() ([]CreateFineTuningJobRequestIntegrationsInner, bool) {
	if o == nil || IsNil(o.Integrations) {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *CreateFineTuningJobRequest) HasIntegrations() bool {
	if o != nil && !IsNil(o.Integrations) {
		return true
	}

	return false
}

// SetIntegrations gets a reference to the given []CreateFineTuningJobRequestIntegrationsInner and assigns it to the Integrations field.
func (o *CreateFineTuningJobRequest) SetIntegrations(v []CreateFineTuningJobRequestIntegrationsInner) {
	o.Integrations = v
}

// GetSeed returns the Seed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFineTuningJobRequest) GetSeed() int32 {
	if o == nil || IsNil(o.Seed.Get()) {
		var ret int32
		return ret
	}
	return *o.Seed.Get()
}

// GetSeedOk returns a tuple with the Seed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFineTuningJobRequest) GetSeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Seed.Get(), o.Seed.IsSet()
}

// HasSeed returns a boolean if a field has been set.
func (o *CreateFineTuningJobRequest) HasSeed() bool {
	if o != nil && o.Seed.IsSet() {
		return true
	}

	return false
}

// SetSeed gets a reference to the given NullableInt32 and assigns it to the Seed field.
func (o *CreateFineTuningJobRequest) SetSeed(v int32) {
	o.Seed.Set(&v)
}
// SetSeedNil sets the value for Seed to be an explicit nil
func (o *CreateFineTuningJobRequest) SetSeedNil() {
	o.Seed.Set(nil)
}

// UnsetSeed ensures that no value is present for Seed, not even an explicit nil
func (o *CreateFineTuningJobRequest) UnsetSeed() {
	o.Seed.Unset()
}

func (o CreateFineTuningJobRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFineTuningJobRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["model"] = o.Model
	toSerialize["training_file"] = o.TrainingFile
	if !IsNil(o.Hyperparameters) {
		toSerialize["hyperparameters"] = o.Hyperparameters
	}
	if o.Suffix.IsSet() {
		toSerialize["suffix"] = o.Suffix.Get()
	}
	if o.ValidationFile.IsSet() {
		toSerialize["validation_file"] = o.ValidationFile.Get()
	}
	if o.Integrations != nil {
		toSerialize["integrations"] = o.Integrations
	}
	if o.Seed.IsSet() {
		toSerialize["seed"] = o.Seed.Get()
	}
	return toSerialize, nil
}

func (o *CreateFineTuningJobRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"model",
		"training_file",
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

	varCreateFineTuningJobRequest := _CreateFineTuningJobRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateFineTuningJobRequest)

	if err != nil {
		return err
	}

	*o = CreateFineTuningJobRequest(varCreateFineTuningJobRequest)

	return err
}

type NullableCreateFineTuningJobRequest struct {
	value *CreateFineTuningJobRequest
	isSet bool
}

func (v NullableCreateFineTuningJobRequest) Get() *CreateFineTuningJobRequest {
	return v.value
}

func (v *NullableCreateFineTuningJobRequest) Set(val *CreateFineTuningJobRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFineTuningJobRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFineTuningJobRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFineTuningJobRequest(val *CreateFineTuningJobRequest) *NullableCreateFineTuningJobRequest {
	return &NullableCreateFineTuningJobRequest{value: val, isSet: true}
}

func (v NullableCreateFineTuningJobRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFineTuningJobRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



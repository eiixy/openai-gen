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

// checks if the CreateFineTuningJobRequestIntegrationsInnerWandb type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFineTuningJobRequestIntegrationsInnerWandb{}

// CreateFineTuningJobRequestIntegrationsInnerWandb The settings for your integration with Weights and Biases. This payload specifies the project that metrics will be sent to. Optionally, you can set an explicit display name for your run, add tags to your run, and set a default entity (team, username, etc) to be associated with your run. 
type CreateFineTuningJobRequestIntegrationsInnerWandb struct {
	// The name of the project that the new run will be created under. 
	Project string `json:"project"`
	// A display name to set for the run. If not set, we will use the Job ID as the name. 
	Name NullableString `json:"name,omitempty"`
	// The entity to use for the run. This allows you to set the team or username of the WandB user that you would like associated with the run. If not set, the default entity for the registered WandB API key is used. 
	Entity NullableString `json:"entity,omitempty"`
	// A list of tags to be attached to the newly created run. These tags are passed through directly to WandB. Some default tags are generated by OpenAI: \"openai/finetune\", \"openai/{base-model}\", \"openai/{ftjob-abcdef}\". 
	Tags []string `json:"tags,omitempty"`
}

type _CreateFineTuningJobRequestIntegrationsInnerWandb CreateFineTuningJobRequestIntegrationsInnerWandb

// NewCreateFineTuningJobRequestIntegrationsInnerWandb instantiates a new CreateFineTuningJobRequestIntegrationsInnerWandb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFineTuningJobRequestIntegrationsInnerWandb(project string) *CreateFineTuningJobRequestIntegrationsInnerWandb {
	this := CreateFineTuningJobRequestIntegrationsInnerWandb{}
	this.Project = project
	return &this
}

// NewCreateFineTuningJobRequestIntegrationsInnerWandbWithDefaults instantiates a new CreateFineTuningJobRequestIntegrationsInnerWandb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFineTuningJobRequestIntegrationsInnerWandbWithDefaults() *CreateFineTuningJobRequestIntegrationsInnerWandb {
	this := CreateFineTuningJobRequestIntegrationsInnerWandb{}
	return &this
}

// GetProject returns the Project field value
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) SetProject(v string) {
	o.Project = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) UnsetName() {
	o.Name.Unset()
}

// GetEntity returns the Entity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) GetEntity() string {
	if o == nil || IsNil(o.Entity.Get()) {
		var ret string
		return ret
	}
	return *o.Entity.Get()
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) GetEntityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entity.Get(), o.Entity.IsSet()
}

// HasEntity returns a boolean if a field has been set.
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) HasEntity() bool {
	if o != nil && o.Entity.IsSet() {
		return true
	}

	return false
}

// SetEntity gets a reference to the given NullableString and assigns it to the Entity field.
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) SetEntity(v string) {
	o.Entity.Set(&v)
}
// SetEntityNil sets the value for Entity to be an explicit nil
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) SetEntityNil() {
	o.Entity.Set(nil)
}

// UnsetEntity ensures that no value is present for Entity, not even an explicit nil
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) UnsetEntity() {
	o.Entity.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) SetTags(v []string) {
	o.Tags = v
}

func (o CreateFineTuningJobRequestIntegrationsInnerWandb) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFineTuningJobRequestIntegrationsInnerWandb) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["project"] = o.Project
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Entity.IsSet() {
		toSerialize["entity"] = o.Entity.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

func (o *CreateFineTuningJobRequestIntegrationsInnerWandb) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"project",
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

	varCreateFineTuningJobRequestIntegrationsInnerWandb := _CreateFineTuningJobRequestIntegrationsInnerWandb{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateFineTuningJobRequestIntegrationsInnerWandb)

	if err != nil {
		return err
	}

	*o = CreateFineTuningJobRequestIntegrationsInnerWandb(varCreateFineTuningJobRequestIntegrationsInnerWandb)

	return err
}

type NullableCreateFineTuningJobRequestIntegrationsInnerWandb struct {
	value *CreateFineTuningJobRequestIntegrationsInnerWandb
	isSet bool
}

func (v NullableCreateFineTuningJobRequestIntegrationsInnerWandb) Get() *CreateFineTuningJobRequestIntegrationsInnerWandb {
	return v.value
}

func (v *NullableCreateFineTuningJobRequestIntegrationsInnerWandb) Set(val *CreateFineTuningJobRequestIntegrationsInnerWandb) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFineTuningJobRequestIntegrationsInnerWandb) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFineTuningJobRequestIntegrationsInnerWandb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFineTuningJobRequestIntegrationsInnerWandb(val *CreateFineTuningJobRequestIntegrationsInnerWandb) *NullableCreateFineTuningJobRequestIntegrationsInnerWandb {
	return &NullableCreateFineTuningJobRequestIntegrationsInnerWandb{value: val, isSet: true}
}

func (v NullableCreateFineTuningJobRequestIntegrationsInnerWandb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFineTuningJobRequestIntegrationsInnerWandb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



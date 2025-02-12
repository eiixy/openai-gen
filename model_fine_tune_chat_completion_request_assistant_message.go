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

// checks if the FineTuneChatCompletionRequestAssistantMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FineTuneChatCompletionRequestAssistantMessage{}

// FineTuneChatCompletionRequestAssistantMessage struct for FineTuneChatCompletionRequestAssistantMessage
type FineTuneChatCompletionRequestAssistantMessage struct {
	// The contents of the assistant message. Required unless `tool_calls` or `function_call` is specified. 
	Content NullableString `json:"content,omitempty"`
	// The role of the messages author, in this case `assistant`.
	Role string `json:"role"`
	// An optional name for the participant. Provides the model information to differentiate between participants of the same role.
	Name *string `json:"name,omitempty"`
	// The tool calls generated by the model, such as function calls.
	ToolCalls []ChatCompletionMessageToolCall `json:"tool_calls,omitempty"`
	// Deprecated
	FunctionCall NullableChatCompletionRequestAssistantMessageFunctionCall `json:"function_call,omitempty"`
	// Controls whether the assistant message is trained against (0 or 1)
	Weight *int32 `json:"weight,omitempty"`
}

type _FineTuneChatCompletionRequestAssistantMessage FineTuneChatCompletionRequestAssistantMessage

// NewFineTuneChatCompletionRequestAssistantMessage instantiates a new FineTuneChatCompletionRequestAssistantMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFineTuneChatCompletionRequestAssistantMessage(role string) *FineTuneChatCompletionRequestAssistantMessage {
	this := FineTuneChatCompletionRequestAssistantMessage{}
	this.Role = role
	return &this
}

// NewFineTuneChatCompletionRequestAssistantMessageWithDefaults instantiates a new FineTuneChatCompletionRequestAssistantMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFineTuneChatCompletionRequestAssistantMessageWithDefaults() *FineTuneChatCompletionRequestAssistantMessage {
	this := FineTuneChatCompletionRequestAssistantMessage{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FineTuneChatCompletionRequestAssistantMessage) GetContent() string {
	if o == nil || IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FineTuneChatCompletionRequestAssistantMessage) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *FineTuneChatCompletionRequestAssistantMessage) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *FineTuneChatCompletionRequestAssistantMessage) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *FineTuneChatCompletionRequestAssistantMessage) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *FineTuneChatCompletionRequestAssistantMessage) UnsetContent() {
	o.Content.Unset()
}

// GetRole returns the Role field value
func (o *FineTuneChatCompletionRequestAssistantMessage) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *FineTuneChatCompletionRequestAssistantMessage) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *FineTuneChatCompletionRequestAssistantMessage) SetRole(v string) {
	o.Role = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FineTuneChatCompletionRequestAssistantMessage) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FineTuneChatCompletionRequestAssistantMessage) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FineTuneChatCompletionRequestAssistantMessage) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FineTuneChatCompletionRequestAssistantMessage) SetName(v string) {
	o.Name = &v
}

// GetToolCalls returns the ToolCalls field value if set, zero value otherwise.
func (o *FineTuneChatCompletionRequestAssistantMessage) GetToolCalls() []ChatCompletionMessageToolCall {
	if o == nil || IsNil(o.ToolCalls) {
		var ret []ChatCompletionMessageToolCall
		return ret
	}
	return o.ToolCalls
}

// GetToolCallsOk returns a tuple with the ToolCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FineTuneChatCompletionRequestAssistantMessage) GetToolCallsOk() ([]ChatCompletionMessageToolCall, bool) {
	if o == nil || IsNil(o.ToolCalls) {
		return nil, false
	}
	return o.ToolCalls, true
}

// HasToolCalls returns a boolean if a field has been set.
func (o *FineTuneChatCompletionRequestAssistantMessage) HasToolCalls() bool {
	if o != nil && !IsNil(o.ToolCalls) {
		return true
	}

	return false
}

// SetToolCalls gets a reference to the given []ChatCompletionMessageToolCall and assigns it to the ToolCalls field.
func (o *FineTuneChatCompletionRequestAssistantMessage) SetToolCalls(v []ChatCompletionMessageToolCall) {
	o.ToolCalls = v
}

// GetFunctionCall returns the FunctionCall field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *FineTuneChatCompletionRequestAssistantMessage) GetFunctionCall() ChatCompletionRequestAssistantMessageFunctionCall {
	if o == nil || IsNil(o.FunctionCall.Get()) {
		var ret ChatCompletionRequestAssistantMessageFunctionCall
		return ret
	}
	return *o.FunctionCall.Get()
}

// GetFunctionCallOk returns a tuple with the FunctionCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *FineTuneChatCompletionRequestAssistantMessage) GetFunctionCallOk() (*ChatCompletionRequestAssistantMessageFunctionCall, bool) {
	if o == nil {
		return nil, false
	}
	return o.FunctionCall.Get(), o.FunctionCall.IsSet()
}

// HasFunctionCall returns a boolean if a field has been set.
func (o *FineTuneChatCompletionRequestAssistantMessage) HasFunctionCall() bool {
	if o != nil && o.FunctionCall.IsSet() {
		return true
	}

	return false
}

// SetFunctionCall gets a reference to the given NullableChatCompletionRequestAssistantMessageFunctionCall and assigns it to the FunctionCall field.
// Deprecated
func (o *FineTuneChatCompletionRequestAssistantMessage) SetFunctionCall(v ChatCompletionRequestAssistantMessageFunctionCall) {
	o.FunctionCall.Set(&v)
}
// SetFunctionCallNil sets the value for FunctionCall to be an explicit nil
func (o *FineTuneChatCompletionRequestAssistantMessage) SetFunctionCallNil() {
	o.FunctionCall.Set(nil)
}

// UnsetFunctionCall ensures that no value is present for FunctionCall, not even an explicit nil
func (o *FineTuneChatCompletionRequestAssistantMessage) UnsetFunctionCall() {
	o.FunctionCall.Unset()
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *FineTuneChatCompletionRequestAssistantMessage) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FineTuneChatCompletionRequestAssistantMessage) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *FineTuneChatCompletionRequestAssistantMessage) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *FineTuneChatCompletionRequestAssistantMessage) SetWeight(v int32) {
	o.Weight = &v
}

func (o FineTuneChatCompletionRequestAssistantMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FineTuneChatCompletionRequestAssistantMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	toSerialize["role"] = o.Role
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ToolCalls) {
		toSerialize["tool_calls"] = o.ToolCalls
	}
	if o.FunctionCall.IsSet() {
		toSerialize["function_call"] = o.FunctionCall.Get()
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

func (o *FineTuneChatCompletionRequestAssistantMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"role",
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

	varFineTuneChatCompletionRequestAssistantMessage := _FineTuneChatCompletionRequestAssistantMessage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFineTuneChatCompletionRequestAssistantMessage)

	if err != nil {
		return err
	}

	*o = FineTuneChatCompletionRequestAssistantMessage(varFineTuneChatCompletionRequestAssistantMessage)

	return err
}

type NullableFineTuneChatCompletionRequestAssistantMessage struct {
	value *FineTuneChatCompletionRequestAssistantMessage
	isSet bool
}

func (v NullableFineTuneChatCompletionRequestAssistantMessage) Get() *FineTuneChatCompletionRequestAssistantMessage {
	return v.value
}

func (v *NullableFineTuneChatCompletionRequestAssistantMessage) Set(val *FineTuneChatCompletionRequestAssistantMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableFineTuneChatCompletionRequestAssistantMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableFineTuneChatCompletionRequestAssistantMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFineTuneChatCompletionRequestAssistantMessage(val *FineTuneChatCompletionRequestAssistantMessage) *NullableFineTuneChatCompletionRequestAssistantMessage {
	return &NullableFineTuneChatCompletionRequestAssistantMessage{value: val, isSet: true}
}

func (v NullableFineTuneChatCompletionRequestAssistantMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFineTuneChatCompletionRequestAssistantMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



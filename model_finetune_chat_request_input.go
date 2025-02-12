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

// checks if the FinetuneChatRequestInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinetuneChatRequestInput{}

// FinetuneChatRequestInput The per-line training example of a fine-tuning input file for chat models
type FinetuneChatRequestInput struct {
	Messages []FinetuneChatRequestInputMessagesInner `json:"messages,omitempty"`
	// A list of tools the model may generate JSON inputs for.
	Tools []ChatCompletionTool `json:"tools,omitempty"`
	// Whether to enable [parallel function calling](/docs/guides/function-calling/parallel-function-calling) during tool use.
	ParallelToolCalls *bool `json:"parallel_tool_calls,omitempty"`
	// A list of functions the model may generate JSON inputs for.
	// Deprecated
	Functions []ChatCompletionFunctions `json:"functions,omitempty"`
}

// NewFinetuneChatRequestInput instantiates a new FinetuneChatRequestInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinetuneChatRequestInput() *FinetuneChatRequestInput {
	this := FinetuneChatRequestInput{}
	var parallelToolCalls bool = true
	this.ParallelToolCalls = &parallelToolCalls
	return &this
}

// NewFinetuneChatRequestInputWithDefaults instantiates a new FinetuneChatRequestInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinetuneChatRequestInputWithDefaults() *FinetuneChatRequestInput {
	this := FinetuneChatRequestInput{}
	var parallelToolCalls bool = true
	this.ParallelToolCalls = &parallelToolCalls
	return &this
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *FinetuneChatRequestInput) GetMessages() []FinetuneChatRequestInputMessagesInner {
	if o == nil || IsNil(o.Messages) {
		var ret []FinetuneChatRequestInputMessagesInner
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinetuneChatRequestInput) GetMessagesOk() ([]FinetuneChatRequestInputMessagesInner, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *FinetuneChatRequestInput) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []FinetuneChatRequestInputMessagesInner and assigns it to the Messages field.
func (o *FinetuneChatRequestInput) SetMessages(v []FinetuneChatRequestInputMessagesInner) {
	o.Messages = v
}

// GetTools returns the Tools field value if set, zero value otherwise.
func (o *FinetuneChatRequestInput) GetTools() []ChatCompletionTool {
	if o == nil || IsNil(o.Tools) {
		var ret []ChatCompletionTool
		return ret
	}
	return o.Tools
}

// GetToolsOk returns a tuple with the Tools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinetuneChatRequestInput) GetToolsOk() ([]ChatCompletionTool, bool) {
	if o == nil || IsNil(o.Tools) {
		return nil, false
	}
	return o.Tools, true
}

// HasTools returns a boolean if a field has been set.
func (o *FinetuneChatRequestInput) HasTools() bool {
	if o != nil && !IsNil(o.Tools) {
		return true
	}

	return false
}

// SetTools gets a reference to the given []ChatCompletionTool and assigns it to the Tools field.
func (o *FinetuneChatRequestInput) SetTools(v []ChatCompletionTool) {
	o.Tools = v
}

// GetParallelToolCalls returns the ParallelToolCalls field value if set, zero value otherwise.
func (o *FinetuneChatRequestInput) GetParallelToolCalls() bool {
	if o == nil || IsNil(o.ParallelToolCalls) {
		var ret bool
		return ret
	}
	return *o.ParallelToolCalls
}

// GetParallelToolCallsOk returns a tuple with the ParallelToolCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinetuneChatRequestInput) GetParallelToolCallsOk() (*bool, bool) {
	if o == nil || IsNil(o.ParallelToolCalls) {
		return nil, false
	}
	return o.ParallelToolCalls, true
}

// HasParallelToolCalls returns a boolean if a field has been set.
func (o *FinetuneChatRequestInput) HasParallelToolCalls() bool {
	if o != nil && !IsNil(o.ParallelToolCalls) {
		return true
	}

	return false
}

// SetParallelToolCalls gets a reference to the given bool and assigns it to the ParallelToolCalls field.
func (o *FinetuneChatRequestInput) SetParallelToolCalls(v bool) {
	o.ParallelToolCalls = &v
}

// GetFunctions returns the Functions field value if set, zero value otherwise.
// Deprecated
func (o *FinetuneChatRequestInput) GetFunctions() []ChatCompletionFunctions {
	if o == nil || IsNil(o.Functions) {
		var ret []ChatCompletionFunctions
		return ret
	}
	return o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *FinetuneChatRequestInput) GetFunctionsOk() ([]ChatCompletionFunctions, bool) {
	if o == nil || IsNil(o.Functions) {
		return nil, false
	}
	return o.Functions, true
}

// HasFunctions returns a boolean if a field has been set.
func (o *FinetuneChatRequestInput) HasFunctions() bool {
	if o != nil && !IsNil(o.Functions) {
		return true
	}

	return false
}

// SetFunctions gets a reference to the given []ChatCompletionFunctions and assigns it to the Functions field.
// Deprecated
func (o *FinetuneChatRequestInput) SetFunctions(v []ChatCompletionFunctions) {
	o.Functions = v
}

func (o FinetuneChatRequestInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinetuneChatRequestInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !IsNil(o.Tools) {
		toSerialize["tools"] = o.Tools
	}
	if !IsNil(o.ParallelToolCalls) {
		toSerialize["parallel_tool_calls"] = o.ParallelToolCalls
	}
	if !IsNil(o.Functions) {
		toSerialize["functions"] = o.Functions
	}
	return toSerialize, nil
}

type NullableFinetuneChatRequestInput struct {
	value *FinetuneChatRequestInput
	isSet bool
}

func (v NullableFinetuneChatRequestInput) Get() *FinetuneChatRequestInput {
	return v.value
}

func (v *NullableFinetuneChatRequestInput) Set(val *FinetuneChatRequestInput) {
	v.value = val
	v.isSet = true
}

func (v NullableFinetuneChatRequestInput) IsSet() bool {
	return v.isSet
}

func (v *NullableFinetuneChatRequestInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinetuneChatRequestInput(val *FinetuneChatRequestInput) *NullableFinetuneChatRequestInput {
	return &NullableFinetuneChatRequestInput{value: val, isSet: true}
}

func (v NullableFinetuneChatRequestInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinetuneChatRequestInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



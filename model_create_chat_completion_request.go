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

// checks if the CreateChatCompletionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateChatCompletionRequest{}

// CreateChatCompletionRequest struct for CreateChatCompletionRequest
type CreateChatCompletionRequest struct {
	// A list of messages comprising the conversation so far. [Example Python code](https://cookbook.openai.com/examples/how_to_format_inputs_to_chatgpt_models).
	Messages []ChatCompletionRequestMessage `json:"messages"`
	Model CreateChatCompletionRequestModel `json:"model"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.  [See more information about frequency and presence penalties.](/docs/guides/text-generation/parameter-details) 
	FrequencyPenalty NullableFloat32 `json:"frequency_penalty,omitempty"`
	// Modify the likelihood of specified tokens appearing in the completion.  Accepts a JSON object that maps tokens (specified by their token ID in the tokenizer) to an associated bias value from -100 to 100. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token. 
	LogitBias map[string]int32 `json:"logit_bias,omitempty"`
	// Whether to return log probabilities of the output tokens or not. If true, returns the log probabilities of each output token returned in the `content` of `message`.
	Logprobs NullableBool `json:"logprobs,omitempty"`
	// An integer between 0 and 20 specifying the number of most likely tokens to return at each token position, each with an associated log probability. `logprobs` must be set to `true` if this parameter is used.
	TopLogprobs NullableInt32 `json:"top_logprobs,omitempty"`
	// The maximum number of [tokens](/tokenizer) that can be generated in the chat completion.  The total length of input tokens and generated tokens is limited by the model's context length. [Example Python code](https://cookbook.openai.com/examples/how_to_count_tokens_with_tiktoken) for counting tokens. 
	MaxTokens NullableInt32 `json:"max_tokens,omitempty"`
	// How many chat completion choices to generate for each input message. Note that you will be charged based on the number of generated tokens across all of the choices. Keep `n` as `1` to minimize costs.
	N NullableInt32 `json:"n,omitempty"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.  [See more information about frequency and presence penalties.](/docs/guides/text-generation/parameter-details) 
	PresencePenalty NullableFloat32 `json:"presence_penalty,omitempty"`
	ResponseFormat *CreateChatCompletionRequestResponseFormat `json:"response_format,omitempty"`
	// This feature is in Beta. If specified, our system will make a best effort to sample deterministically, such that repeated requests with the same `seed` and parameters should return the same result. Determinism is not guaranteed, and you should refer to the `system_fingerprint` response parameter to monitor changes in the backend. 
	Seed NullableInt32 `json:"seed,omitempty"`
	// Specifies the latency tier to use for processing the request. This parameter is relevant for customers subscribed to the scale tier service:   - If set to 'auto', the system will utilize scale tier credits until they are exhausted.   - If set to 'default', the request will be processed using the default service tier with a lower uptime SLA and no latency guarentee.    When this parameter is set, the response body will include the `service_tier` utilized. 
	ServiceTier NullableString `json:"service_tier,omitempty"`
	Stop *CreateChatCompletionRequestStop `json:"stop,omitempty"`
	// If set, partial message deltas will be sent, like in ChatGPT. Tokens will be sent as data-only [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format) as they become available, with the stream terminated by a `data: [DONE]` message. [Example Python code](https://cookbook.openai.com/examples/how_to_stream_completions). 
	Stream NullableBool `json:"stream,omitempty"`
	StreamOptions NullableChatCompletionStreamOptions `json:"stream_options,omitempty"`
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.  We generally recommend altering this or `top_p` but not both. 
	Temperature NullableFloat32 `json:"temperature,omitempty"`
	// An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.  We generally recommend altering this or `temperature` but not both. 
	TopP NullableFloat32 `json:"top_p,omitempty"`
	// A list of tools the model may call. Currently, only functions are supported as a tool. Use this to provide a list of functions the model may generate JSON inputs for. A max of 128 functions are supported. 
	Tools []ChatCompletionTool `json:"tools,omitempty"`
	ToolChoice *ChatCompletionToolChoiceOption `json:"tool_choice,omitempty"`
	// Whether to enable [parallel function calling](/docs/guides/function-calling/parallel-function-calling) during tool use.
	ParallelToolCalls *bool `json:"parallel_tool_calls,omitempty"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids). 
	User *string `json:"user,omitempty"`
	// Deprecated
	FunctionCall *CreateChatCompletionRequestFunctionCall `json:"function_call,omitempty"`
	// Deprecated in favor of `tools`.  A list of functions the model may generate JSON inputs for. 
	// Deprecated
	Functions []ChatCompletionFunctions `json:"functions,omitempty"`
}

type _CreateChatCompletionRequest CreateChatCompletionRequest

// NewCreateChatCompletionRequest instantiates a new CreateChatCompletionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateChatCompletionRequest(messages []ChatCompletionRequestMessage, model CreateChatCompletionRequestModel) *CreateChatCompletionRequest {
	this := CreateChatCompletionRequest{}
	this.Messages = messages
	this.Model = model
	var frequencyPenalty float32 = 0
	this.FrequencyPenalty = *NewNullableFloat32(&frequencyPenalty)
	var logprobs bool = false
	this.Logprobs = *NewNullableBool(&logprobs)
	var n int32 = 1
	this.N = *NewNullableInt32(&n)
	var presencePenalty float32 = 0
	this.PresencePenalty = *NewNullableFloat32(&presencePenalty)
	var stop CreateChatCompletionRequestStop = null
	this.Stop = &stop
	var stream bool = false
	this.Stream = *NewNullableBool(&stream)
	var temperature float32 = 1
	this.Temperature = *NewNullableFloat32(&temperature)
	var topP float32 = 1
	this.TopP = *NewNullableFloat32(&topP)
	var parallelToolCalls bool = true
	this.ParallelToolCalls = &parallelToolCalls
	return &this
}

// NewCreateChatCompletionRequestWithDefaults instantiates a new CreateChatCompletionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateChatCompletionRequestWithDefaults() *CreateChatCompletionRequest {
	this := CreateChatCompletionRequest{}
	var frequencyPenalty float32 = 0
	this.FrequencyPenalty = *NewNullableFloat32(&frequencyPenalty)
	var logprobs bool = false
	this.Logprobs = *NewNullableBool(&logprobs)
	var n int32 = 1
	this.N = *NewNullableInt32(&n)
	var presencePenalty float32 = 0
	this.PresencePenalty = *NewNullableFloat32(&presencePenalty)
	var stop CreateChatCompletionRequestStop = null
	this.Stop = &stop
	var stream bool = false
	this.Stream = *NewNullableBool(&stream)
	var temperature float32 = 1
	this.Temperature = *NewNullableFloat32(&temperature)
	var topP float32 = 1
	this.TopP = *NewNullableFloat32(&topP)
	var parallelToolCalls bool = true
	this.ParallelToolCalls = &parallelToolCalls
	return &this
}

// GetMessages returns the Messages field value
func (o *CreateChatCompletionRequest) GetMessages() []ChatCompletionRequestMessage {
	if o == nil {
		var ret []ChatCompletionRequestMessage
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionRequest) GetMessagesOk() ([]ChatCompletionRequestMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *CreateChatCompletionRequest) SetMessages(v []ChatCompletionRequestMessage) {
	o.Messages = v
}

// GetModel returns the Model field value
func (o *CreateChatCompletionRequest) GetModel() CreateChatCompletionRequestModel {
	if o == nil {
		var ret CreateChatCompletionRequestModel
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionRequest) GetModelOk() (*CreateChatCompletionRequestModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *CreateChatCompletionRequest) SetModel(v CreateChatCompletionRequestModel) {
	o.Model = v
}

// GetFrequencyPenalty returns the FrequencyPenalty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateChatCompletionRequest) GetFrequencyPenalty() float32 {
	if o == nil || IsNil(o.FrequencyPenalty.Get()) {
		var ret float32
		return ret
	}
	return *o.FrequencyPenalty.Get()
}

// GetFrequencyPenaltyOk returns a tuple with the FrequencyPenalty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateChatCompletionRequest) GetFrequencyPenaltyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FrequencyPenalty.Get(), o.FrequencyPenalty.IsSet()
}

// HasFrequencyPenalty returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasFrequencyPenalty() bool {
	if o != nil && o.FrequencyPenalty.IsSet() {
		return true
	}

	return false
}

// SetFrequencyPenalty gets a reference to the given NullableFloat32 and assigns it to the FrequencyPenalty field.
func (o *CreateChatCompletionRequest) SetFrequencyPenalty(v float32) {
	o.FrequencyPenalty.Set(&v)
}
// SetFrequencyPenaltyNil sets the value for FrequencyPenalty to be an explicit nil
func (o *CreateChatCompletionRequest) SetFrequencyPenaltyNil() {
	o.FrequencyPenalty.Set(nil)
}

// UnsetFrequencyPenalty ensures that no value is present for FrequencyPenalty, not even an explicit nil
func (o *CreateChatCompletionRequest) UnsetFrequencyPenalty() {
	o.FrequencyPenalty.Unset()
}

// GetLogitBias returns the LogitBias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateChatCompletionRequest) GetLogitBias() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}
	return o.LogitBias
}

// GetLogitBiasOk returns a tuple with the LogitBias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateChatCompletionRequest) GetLogitBiasOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.LogitBias) {
		return nil, false
	}
	return &o.LogitBias, true
}

// HasLogitBias returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasLogitBias() bool {
	if o != nil && !IsNil(o.LogitBias) {
		return true
	}

	return false
}

// SetLogitBias gets a reference to the given map[string]int32 and assigns it to the LogitBias field.
func (o *CreateChatCompletionRequest) SetLogitBias(v map[string]int32) {
	o.LogitBias = v
}

// GetLogprobs returns the Logprobs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateChatCompletionRequest) GetLogprobs() bool {
	if o == nil || IsNil(o.Logprobs.Get()) {
		var ret bool
		return ret
	}
	return *o.Logprobs.Get()
}

// GetLogprobsOk returns a tuple with the Logprobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateChatCompletionRequest) GetLogprobsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logprobs.Get(), o.Logprobs.IsSet()
}

// HasLogprobs returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasLogprobs() bool {
	if o != nil && o.Logprobs.IsSet() {
		return true
	}

	return false
}

// SetLogprobs gets a reference to the given NullableBool and assigns it to the Logprobs field.
func (o *CreateChatCompletionRequest) SetLogprobs(v bool) {
	o.Logprobs.Set(&v)
}
// SetLogprobsNil sets the value for Logprobs to be an explicit nil
func (o *CreateChatCompletionRequest) SetLogprobsNil() {
	o.Logprobs.Set(nil)
}

// UnsetLogprobs ensures that no value is present for Logprobs, not even an explicit nil
func (o *CreateChatCompletionRequest) UnsetLogprobs() {
	o.Logprobs.Unset()
}

// GetTopLogprobs returns the TopLogprobs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateChatCompletionRequest) GetTopLogprobs() int32 {
	if o == nil || IsNil(o.TopLogprobs.Get()) {
		var ret int32
		return ret
	}
	return *o.TopLogprobs.Get()
}

// GetTopLogprobsOk returns a tuple with the TopLogprobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateChatCompletionRequest) GetTopLogprobsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopLogprobs.Get(), o.TopLogprobs.IsSet()
}

// HasTopLogprobs returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasTopLogprobs() bool {
	if o != nil && o.TopLogprobs.IsSet() {
		return true
	}

	return false
}

// SetTopLogprobs gets a reference to the given NullableInt32 and assigns it to the TopLogprobs field.
func (o *CreateChatCompletionRequest) SetTopLogprobs(v int32) {
	o.TopLogprobs.Set(&v)
}
// SetTopLogprobsNil sets the value for TopLogprobs to be an explicit nil
func (o *CreateChatCompletionRequest) SetTopLogprobsNil() {
	o.TopLogprobs.Set(nil)
}

// UnsetTopLogprobs ensures that no value is present for TopLogprobs, not even an explicit nil
func (o *CreateChatCompletionRequest) UnsetTopLogprobs() {
	o.TopLogprobs.Unset()
}

// GetMaxTokens returns the MaxTokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateChatCompletionRequest) GetMaxTokens() int32 {
	if o == nil || IsNil(o.MaxTokens.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxTokens.Get()
}

// GetMaxTokensOk returns a tuple with the MaxTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateChatCompletionRequest) GetMaxTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxTokens.Get(), o.MaxTokens.IsSet()
}

// HasMaxTokens returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasMaxTokens() bool {
	if o != nil && o.MaxTokens.IsSet() {
		return true
	}

	return false
}

// SetMaxTokens gets a reference to the given NullableInt32 and assigns it to the MaxTokens field.
func (o *CreateChatCompletionRequest) SetMaxTokens(v int32) {
	o.MaxTokens.Set(&v)
}
// SetMaxTokensNil sets the value for MaxTokens to be an explicit nil
func (o *CreateChatCompletionRequest) SetMaxTokensNil() {
	o.MaxTokens.Set(nil)
}

// UnsetMaxTokens ensures that no value is present for MaxTokens, not even an explicit nil
func (o *CreateChatCompletionRequest) UnsetMaxTokens() {
	o.MaxTokens.Unset()
}

// GetN returns the N field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateChatCompletionRequest) GetN() int32 {
	if o == nil || IsNil(o.N.Get()) {
		var ret int32
		return ret
	}
	return *o.N.Get()
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateChatCompletionRequest) GetNOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.N.Get(), o.N.IsSet()
}

// HasN returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasN() bool {
	if o != nil && o.N.IsSet() {
		return true
	}

	return false
}

// SetN gets a reference to the given NullableInt32 and assigns it to the N field.
func (o *CreateChatCompletionRequest) SetN(v int32) {
	o.N.Set(&v)
}
// SetNNil sets the value for N to be an explicit nil
func (o *CreateChatCompletionRequest) SetNNil() {
	o.N.Set(nil)
}

// UnsetN ensures that no value is present for N, not even an explicit nil
func (o *CreateChatCompletionRequest) UnsetN() {
	o.N.Unset()
}

// GetPresencePenalty returns the PresencePenalty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateChatCompletionRequest) GetPresencePenalty() float32 {
	if o == nil || IsNil(o.PresencePenalty.Get()) {
		var ret float32
		return ret
	}
	return *o.PresencePenalty.Get()
}

// GetPresencePenaltyOk returns a tuple with the PresencePenalty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateChatCompletionRequest) GetPresencePenaltyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PresencePenalty.Get(), o.PresencePenalty.IsSet()
}

// HasPresencePenalty returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasPresencePenalty() bool {
	if o != nil && o.PresencePenalty.IsSet() {
		return true
	}

	return false
}

// SetPresencePenalty gets a reference to the given NullableFloat32 and assigns it to the PresencePenalty field.
func (o *CreateChatCompletionRequest) SetPresencePenalty(v float32) {
	o.PresencePenalty.Set(&v)
}
// SetPresencePenaltyNil sets the value for PresencePenalty to be an explicit nil
func (o *CreateChatCompletionRequest) SetPresencePenaltyNil() {
	o.PresencePenalty.Set(nil)
}

// UnsetPresencePenalty ensures that no value is present for PresencePenalty, not even an explicit nil
func (o *CreateChatCompletionRequest) UnsetPresencePenalty() {
	o.PresencePenalty.Unset()
}

// GetResponseFormat returns the ResponseFormat field value if set, zero value otherwise.
func (o *CreateChatCompletionRequest) GetResponseFormat() CreateChatCompletionRequestResponseFormat {
	if o == nil || IsNil(o.ResponseFormat) {
		var ret CreateChatCompletionRequestResponseFormat
		return ret
	}
	return *o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionRequest) GetResponseFormatOk() (*CreateChatCompletionRequestResponseFormat, bool) {
	if o == nil || IsNil(o.ResponseFormat) {
		return nil, false
	}
	return o.ResponseFormat, true
}

// HasResponseFormat returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasResponseFormat() bool {
	if o != nil && !IsNil(o.ResponseFormat) {
		return true
	}

	return false
}

// SetResponseFormat gets a reference to the given CreateChatCompletionRequestResponseFormat and assigns it to the ResponseFormat field.
func (o *CreateChatCompletionRequest) SetResponseFormat(v CreateChatCompletionRequestResponseFormat) {
	o.ResponseFormat = &v
}

// GetSeed returns the Seed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateChatCompletionRequest) GetSeed() int32 {
	if o == nil || IsNil(o.Seed.Get()) {
		var ret int32
		return ret
	}
	return *o.Seed.Get()
}

// GetSeedOk returns a tuple with the Seed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateChatCompletionRequest) GetSeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Seed.Get(), o.Seed.IsSet()
}

// HasSeed returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasSeed() bool {
	if o != nil && o.Seed.IsSet() {
		return true
	}

	return false
}

// SetSeed gets a reference to the given NullableInt32 and assigns it to the Seed field.
func (o *CreateChatCompletionRequest) SetSeed(v int32) {
	o.Seed.Set(&v)
}
// SetSeedNil sets the value for Seed to be an explicit nil
func (o *CreateChatCompletionRequest) SetSeedNil() {
	o.Seed.Set(nil)
}

// UnsetSeed ensures that no value is present for Seed, not even an explicit nil
func (o *CreateChatCompletionRequest) UnsetSeed() {
	o.Seed.Unset()
}

// GetServiceTier returns the ServiceTier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateChatCompletionRequest) GetServiceTier() string {
	if o == nil || IsNil(o.ServiceTier.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceTier.Get()
}

// GetServiceTierOk returns a tuple with the ServiceTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateChatCompletionRequest) GetServiceTierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceTier.Get(), o.ServiceTier.IsSet()
}

// HasServiceTier returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasServiceTier() bool {
	if o != nil && o.ServiceTier.IsSet() {
		return true
	}

	return false
}

// SetServiceTier gets a reference to the given NullableString and assigns it to the ServiceTier field.
func (o *CreateChatCompletionRequest) SetServiceTier(v string) {
	o.ServiceTier.Set(&v)
}
// SetServiceTierNil sets the value for ServiceTier to be an explicit nil
func (o *CreateChatCompletionRequest) SetServiceTierNil() {
	o.ServiceTier.Set(nil)
}

// UnsetServiceTier ensures that no value is present for ServiceTier, not even an explicit nil
func (o *CreateChatCompletionRequest) UnsetServiceTier() {
	o.ServiceTier.Unset()
}

// GetStop returns the Stop field value if set, zero value otherwise.
func (o *CreateChatCompletionRequest) GetStop() CreateChatCompletionRequestStop {
	if o == nil || IsNil(o.Stop) {
		var ret CreateChatCompletionRequestStop
		return ret
	}
	return *o.Stop
}

// GetStopOk returns a tuple with the Stop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionRequest) GetStopOk() (*CreateChatCompletionRequestStop, bool) {
	if o == nil || IsNil(o.Stop) {
		return nil, false
	}
	return o.Stop, true
}

// HasStop returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasStop() bool {
	if o != nil && !IsNil(o.Stop) {
		return true
	}

	return false
}

// SetStop gets a reference to the given CreateChatCompletionRequestStop and assigns it to the Stop field.
func (o *CreateChatCompletionRequest) SetStop(v CreateChatCompletionRequestStop) {
	o.Stop = &v
}

// GetStream returns the Stream field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateChatCompletionRequest) GetStream() bool {
	if o == nil || IsNil(o.Stream.Get()) {
		var ret bool
		return ret
	}
	return *o.Stream.Get()
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateChatCompletionRequest) GetStreamOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stream.Get(), o.Stream.IsSet()
}

// HasStream returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasStream() bool {
	if o != nil && o.Stream.IsSet() {
		return true
	}

	return false
}

// SetStream gets a reference to the given NullableBool and assigns it to the Stream field.
func (o *CreateChatCompletionRequest) SetStream(v bool) {
	o.Stream.Set(&v)
}
// SetStreamNil sets the value for Stream to be an explicit nil
func (o *CreateChatCompletionRequest) SetStreamNil() {
	o.Stream.Set(nil)
}

// UnsetStream ensures that no value is present for Stream, not even an explicit nil
func (o *CreateChatCompletionRequest) UnsetStream() {
	o.Stream.Unset()
}

// GetStreamOptions returns the StreamOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateChatCompletionRequest) GetStreamOptions() ChatCompletionStreamOptions {
	if o == nil || IsNil(o.StreamOptions.Get()) {
		var ret ChatCompletionStreamOptions
		return ret
	}
	return *o.StreamOptions.Get()
}

// GetStreamOptionsOk returns a tuple with the StreamOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateChatCompletionRequest) GetStreamOptionsOk() (*ChatCompletionStreamOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreamOptions.Get(), o.StreamOptions.IsSet()
}

// HasStreamOptions returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasStreamOptions() bool {
	if o != nil && o.StreamOptions.IsSet() {
		return true
	}

	return false
}

// SetStreamOptions gets a reference to the given NullableChatCompletionStreamOptions and assigns it to the StreamOptions field.
func (o *CreateChatCompletionRequest) SetStreamOptions(v ChatCompletionStreamOptions) {
	o.StreamOptions.Set(&v)
}
// SetStreamOptionsNil sets the value for StreamOptions to be an explicit nil
func (o *CreateChatCompletionRequest) SetStreamOptionsNil() {
	o.StreamOptions.Set(nil)
}

// UnsetStreamOptions ensures that no value is present for StreamOptions, not even an explicit nil
func (o *CreateChatCompletionRequest) UnsetStreamOptions() {
	o.StreamOptions.Unset()
}

// GetTemperature returns the Temperature field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateChatCompletionRequest) GetTemperature() float32 {
	if o == nil || IsNil(o.Temperature.Get()) {
		var ret float32
		return ret
	}
	return *o.Temperature.Get()
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateChatCompletionRequest) GetTemperatureOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Temperature.Get(), o.Temperature.IsSet()
}

// HasTemperature returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasTemperature() bool {
	if o != nil && o.Temperature.IsSet() {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given NullableFloat32 and assigns it to the Temperature field.
func (o *CreateChatCompletionRequest) SetTemperature(v float32) {
	o.Temperature.Set(&v)
}
// SetTemperatureNil sets the value for Temperature to be an explicit nil
func (o *CreateChatCompletionRequest) SetTemperatureNil() {
	o.Temperature.Set(nil)
}

// UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
func (o *CreateChatCompletionRequest) UnsetTemperature() {
	o.Temperature.Unset()
}

// GetTopP returns the TopP field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateChatCompletionRequest) GetTopP() float32 {
	if o == nil || IsNil(o.TopP.Get()) {
		var ret float32
		return ret
	}
	return *o.TopP.Get()
}

// GetTopPOk returns a tuple with the TopP field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateChatCompletionRequest) GetTopPOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopP.Get(), o.TopP.IsSet()
}

// HasTopP returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasTopP() bool {
	if o != nil && o.TopP.IsSet() {
		return true
	}

	return false
}

// SetTopP gets a reference to the given NullableFloat32 and assigns it to the TopP field.
func (o *CreateChatCompletionRequest) SetTopP(v float32) {
	o.TopP.Set(&v)
}
// SetTopPNil sets the value for TopP to be an explicit nil
func (o *CreateChatCompletionRequest) SetTopPNil() {
	o.TopP.Set(nil)
}

// UnsetTopP ensures that no value is present for TopP, not even an explicit nil
func (o *CreateChatCompletionRequest) UnsetTopP() {
	o.TopP.Unset()
}

// GetTools returns the Tools field value if set, zero value otherwise.
func (o *CreateChatCompletionRequest) GetTools() []ChatCompletionTool {
	if o == nil || IsNil(o.Tools) {
		var ret []ChatCompletionTool
		return ret
	}
	return o.Tools
}

// GetToolsOk returns a tuple with the Tools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionRequest) GetToolsOk() ([]ChatCompletionTool, bool) {
	if o == nil || IsNil(o.Tools) {
		return nil, false
	}
	return o.Tools, true
}

// HasTools returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasTools() bool {
	if o != nil && !IsNil(o.Tools) {
		return true
	}

	return false
}

// SetTools gets a reference to the given []ChatCompletionTool and assigns it to the Tools field.
func (o *CreateChatCompletionRequest) SetTools(v []ChatCompletionTool) {
	o.Tools = v
}

// GetToolChoice returns the ToolChoice field value if set, zero value otherwise.
func (o *CreateChatCompletionRequest) GetToolChoice() ChatCompletionToolChoiceOption {
	if o == nil || IsNil(o.ToolChoice) {
		var ret ChatCompletionToolChoiceOption
		return ret
	}
	return *o.ToolChoice
}

// GetToolChoiceOk returns a tuple with the ToolChoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionRequest) GetToolChoiceOk() (*ChatCompletionToolChoiceOption, bool) {
	if o == nil || IsNil(o.ToolChoice) {
		return nil, false
	}
	return o.ToolChoice, true
}

// HasToolChoice returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasToolChoice() bool {
	if o != nil && !IsNil(o.ToolChoice) {
		return true
	}

	return false
}

// SetToolChoice gets a reference to the given ChatCompletionToolChoiceOption and assigns it to the ToolChoice field.
func (o *CreateChatCompletionRequest) SetToolChoice(v ChatCompletionToolChoiceOption) {
	o.ToolChoice = &v
}

// GetParallelToolCalls returns the ParallelToolCalls field value if set, zero value otherwise.
func (o *CreateChatCompletionRequest) GetParallelToolCalls() bool {
	if o == nil || IsNil(o.ParallelToolCalls) {
		var ret bool
		return ret
	}
	return *o.ParallelToolCalls
}

// GetParallelToolCallsOk returns a tuple with the ParallelToolCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionRequest) GetParallelToolCallsOk() (*bool, bool) {
	if o == nil || IsNil(o.ParallelToolCalls) {
		return nil, false
	}
	return o.ParallelToolCalls, true
}

// HasParallelToolCalls returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasParallelToolCalls() bool {
	if o != nil && !IsNil(o.ParallelToolCalls) {
		return true
	}

	return false
}

// SetParallelToolCalls gets a reference to the given bool and assigns it to the ParallelToolCalls field.
func (o *CreateChatCompletionRequest) SetParallelToolCalls(v bool) {
	o.ParallelToolCalls = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CreateChatCompletionRequest) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChatCompletionRequest) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *CreateChatCompletionRequest) SetUser(v string) {
	o.User = &v
}

// GetFunctionCall returns the FunctionCall field value if set, zero value otherwise.
// Deprecated
func (o *CreateChatCompletionRequest) GetFunctionCall() CreateChatCompletionRequestFunctionCall {
	if o == nil || IsNil(o.FunctionCall) {
		var ret CreateChatCompletionRequestFunctionCall
		return ret
	}
	return *o.FunctionCall
}

// GetFunctionCallOk returns a tuple with the FunctionCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateChatCompletionRequest) GetFunctionCallOk() (*CreateChatCompletionRequestFunctionCall, bool) {
	if o == nil || IsNil(o.FunctionCall) {
		return nil, false
	}
	return o.FunctionCall, true
}

// HasFunctionCall returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasFunctionCall() bool {
	if o != nil && !IsNil(o.FunctionCall) {
		return true
	}

	return false
}

// SetFunctionCall gets a reference to the given CreateChatCompletionRequestFunctionCall and assigns it to the FunctionCall field.
// Deprecated
func (o *CreateChatCompletionRequest) SetFunctionCall(v CreateChatCompletionRequestFunctionCall) {
	o.FunctionCall = &v
}

// GetFunctions returns the Functions field value if set, zero value otherwise.
// Deprecated
func (o *CreateChatCompletionRequest) GetFunctions() []ChatCompletionFunctions {
	if o == nil || IsNil(o.Functions) {
		var ret []ChatCompletionFunctions
		return ret
	}
	return o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateChatCompletionRequest) GetFunctionsOk() ([]ChatCompletionFunctions, bool) {
	if o == nil || IsNil(o.Functions) {
		return nil, false
	}
	return o.Functions, true
}

// HasFunctions returns a boolean if a field has been set.
func (o *CreateChatCompletionRequest) HasFunctions() bool {
	if o != nil && !IsNil(o.Functions) {
		return true
	}

	return false
}

// SetFunctions gets a reference to the given []ChatCompletionFunctions and assigns it to the Functions field.
// Deprecated
func (o *CreateChatCompletionRequest) SetFunctions(v []ChatCompletionFunctions) {
	o.Functions = v
}

func (o CreateChatCompletionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateChatCompletionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["messages"] = o.Messages
	toSerialize["model"] = o.Model
	if o.FrequencyPenalty.IsSet() {
		toSerialize["frequency_penalty"] = o.FrequencyPenalty.Get()
	}
	if o.LogitBias != nil {
		toSerialize["logit_bias"] = o.LogitBias
	}
	if o.Logprobs.IsSet() {
		toSerialize["logprobs"] = o.Logprobs.Get()
	}
	if o.TopLogprobs.IsSet() {
		toSerialize["top_logprobs"] = o.TopLogprobs.Get()
	}
	if o.MaxTokens.IsSet() {
		toSerialize["max_tokens"] = o.MaxTokens.Get()
	}
	if o.N.IsSet() {
		toSerialize["n"] = o.N.Get()
	}
	if o.PresencePenalty.IsSet() {
		toSerialize["presence_penalty"] = o.PresencePenalty.Get()
	}
	if !IsNil(o.ResponseFormat) {
		toSerialize["response_format"] = o.ResponseFormat
	}
	if o.Seed.IsSet() {
		toSerialize["seed"] = o.Seed.Get()
	}
	if o.ServiceTier.IsSet() {
		toSerialize["service_tier"] = o.ServiceTier.Get()
	}
	if !IsNil(o.Stop) {
		toSerialize["stop"] = o.Stop
	}
	if o.Stream.IsSet() {
		toSerialize["stream"] = o.Stream.Get()
	}
	if o.StreamOptions.IsSet() {
		toSerialize["stream_options"] = o.StreamOptions.Get()
	}
	if o.Temperature.IsSet() {
		toSerialize["temperature"] = o.Temperature.Get()
	}
	if o.TopP.IsSet() {
		toSerialize["top_p"] = o.TopP.Get()
	}
	if !IsNil(o.Tools) {
		toSerialize["tools"] = o.Tools
	}
	if !IsNil(o.ToolChoice) {
		toSerialize["tool_choice"] = o.ToolChoice
	}
	if !IsNil(o.ParallelToolCalls) {
		toSerialize["parallel_tool_calls"] = o.ParallelToolCalls
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.FunctionCall) {
		toSerialize["function_call"] = o.FunctionCall
	}
	if !IsNil(o.Functions) {
		toSerialize["functions"] = o.Functions
	}
	return toSerialize, nil
}

func (o *CreateChatCompletionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"messages",
		"model",
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

	varCreateChatCompletionRequest := _CreateChatCompletionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateChatCompletionRequest)

	if err != nil {
		return err
	}

	*o = CreateChatCompletionRequest(varCreateChatCompletionRequest)

	return err
}

type NullableCreateChatCompletionRequest struct {
	value *CreateChatCompletionRequest
	isSet bool
}

func (v NullableCreateChatCompletionRequest) Get() *CreateChatCompletionRequest {
	return v.value
}

func (v *NullableCreateChatCompletionRequest) Set(val *CreateChatCompletionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateChatCompletionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateChatCompletionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateChatCompletionRequest(val *CreateChatCompletionRequest) *NullableCreateChatCompletionRequest {
	return &NullableCreateChatCompletionRequest{value: val, isSet: true}
}

func (v NullableCreateChatCompletionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateChatCompletionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



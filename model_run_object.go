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

// checks if the RunObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunObject{}

// RunObject Represents an execution run on a [thread](/docs/api-reference/threads).
type RunObject struct {
	// The identifier, which can be referenced in API endpoints.
	Id string `json:"id"`
	// The object type, which is always `thread.run`.
	Object string `json:"object"`
	// The Unix timestamp (in seconds) for when the run was created.
	CreatedAt int32 `json:"created_at"`
	// The ID of the [thread](/docs/api-reference/threads) that was executed on as a part of this run.
	ThreadId string `json:"thread_id"`
	// The ID of the [assistant](/docs/api-reference/assistants) used for execution of this run.
	AssistantId string `json:"assistant_id"`
	// The status of the run, which can be either `queued`, `in_progress`, `requires_action`, `cancelling`, `cancelled`, `failed`, `completed`, `incomplete`, or `expired`.
	Status string `json:"status"`
	RequiredAction NullableRunObjectRequiredAction `json:"required_action"`
	LastError NullableRunObjectLastError `json:"last_error"`
	// The Unix timestamp (in seconds) for when the run will expire.
	ExpiresAt NullableInt32 `json:"expires_at"`
	// The Unix timestamp (in seconds) for when the run was started.
	StartedAt NullableInt32 `json:"started_at"`
	// The Unix timestamp (in seconds) for when the run was cancelled.
	CancelledAt NullableInt32 `json:"cancelled_at"`
	// The Unix timestamp (in seconds) for when the run failed.
	FailedAt NullableInt32 `json:"failed_at"`
	// The Unix timestamp (in seconds) for when the run was completed.
	CompletedAt NullableInt32 `json:"completed_at"`
	IncompleteDetails NullableRunObjectIncompleteDetails `json:"incomplete_details"`
	// The model that the [assistant](/docs/api-reference/assistants) used for this run.
	Model string `json:"model"`
	// The instructions that the [assistant](/docs/api-reference/assistants) used for this run.
	Instructions string `json:"instructions"`
	// The list of tools that the [assistant](/docs/api-reference/assistants) used for this run.
	Tools []AssistantObjectToolsInner `json:"tools"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long. 
	Metadata map[string]interface{} `json:"metadata"`
	Usage NullableRunCompletionUsage `json:"usage"`
	// The sampling temperature used for this run. If not set, defaults to 1.
	Temperature NullableFloat32 `json:"temperature,omitempty"`
	// The nucleus sampling value used for this run. If not set, defaults to 1.
	TopP NullableFloat32 `json:"top_p,omitempty"`
	// The maximum number of prompt tokens specified to have been used over the course of the run. 
	MaxPromptTokens NullableInt32 `json:"max_prompt_tokens"`
	// The maximum number of completion tokens specified to have been used over the course of the run. 
	MaxCompletionTokens NullableInt32 `json:"max_completion_tokens"`
	TruncationStrategy TruncationObject `json:"truncation_strategy"`
	ToolChoice AssistantsApiToolChoiceOption `json:"tool_choice"`
	// Whether to enable [parallel function calling](/docs/guides/function-calling/parallel-function-calling) during tool use.
	ParallelToolCalls bool `json:"parallel_tool_calls"`
	ResponseFormat AssistantsApiResponseFormatOption `json:"response_format"`
}

type _RunObject RunObject

// NewRunObject instantiates a new RunObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunObject(id string, object string, createdAt int32, threadId string, assistantId string, status string, requiredAction NullableRunObjectRequiredAction, lastError NullableRunObjectLastError, expiresAt NullableInt32, startedAt NullableInt32, cancelledAt NullableInt32, failedAt NullableInt32, completedAt NullableInt32, incompleteDetails NullableRunObjectIncompleteDetails, model string, instructions string, tools []AssistantObjectToolsInner, metadata map[string]interface{}, usage NullableRunCompletionUsage, maxPromptTokens NullableInt32, maxCompletionTokens NullableInt32, truncationStrategy TruncationObject, toolChoice AssistantsApiToolChoiceOption, parallelToolCalls bool, responseFormat AssistantsApiResponseFormatOption) *RunObject {
	this := RunObject{}
	this.Id = id
	this.Object = object
	this.CreatedAt = createdAt
	this.ThreadId = threadId
	this.AssistantId = assistantId
	this.Status = status
	this.RequiredAction = requiredAction
	this.LastError = lastError
	this.ExpiresAt = expiresAt
	this.StartedAt = startedAt
	this.CancelledAt = cancelledAt
	this.FailedAt = failedAt
	this.CompletedAt = completedAt
	this.IncompleteDetails = incompleteDetails
	this.Model = model
	this.Instructions = instructions
	this.Tools = tools
	this.Metadata = metadata
	this.Usage = usage
	this.MaxPromptTokens = maxPromptTokens
	this.MaxCompletionTokens = maxCompletionTokens
	this.TruncationStrategy = truncationStrategy
	this.ToolChoice = toolChoice
	this.ParallelToolCalls = parallelToolCalls
	this.ResponseFormat = responseFormat
	return &this
}

// NewRunObjectWithDefaults instantiates a new RunObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunObjectWithDefaults() *RunObject {
	this := RunObject{}
	var parallelToolCalls bool = true
	this.ParallelToolCalls = parallelToolCalls
	return &this
}

// GetId returns the Id field value
func (o *RunObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RunObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RunObject) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *RunObject) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *RunObject) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *RunObject) SetObject(v string) {
	o.Object = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RunObject) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RunObject) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RunObject) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetThreadId returns the ThreadId field value
func (o *RunObject) GetThreadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value
// and a boolean to check if the value has been set.
func (o *RunObject) GetThreadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadId, true
}

// SetThreadId sets field value
func (o *RunObject) SetThreadId(v string) {
	o.ThreadId = v
}

// GetAssistantId returns the AssistantId field value
func (o *RunObject) GetAssistantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssistantId
}

// GetAssistantIdOk returns a tuple with the AssistantId field value
// and a boolean to check if the value has been set.
func (o *RunObject) GetAssistantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssistantId, true
}

// SetAssistantId sets field value
func (o *RunObject) SetAssistantId(v string) {
	o.AssistantId = v
}

// GetStatus returns the Status field value
func (o *RunObject) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RunObject) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RunObject) SetStatus(v string) {
	o.Status = v
}

// GetRequiredAction returns the RequiredAction field value
// If the value is explicit nil, the zero value for RunObjectRequiredAction will be returned
func (o *RunObject) GetRequiredAction() RunObjectRequiredAction {
	if o == nil || o.RequiredAction.Get() == nil {
		var ret RunObjectRequiredAction
		return ret
	}

	return *o.RequiredAction.Get()
}

// GetRequiredActionOk returns a tuple with the RequiredAction field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunObject) GetRequiredActionOk() (*RunObjectRequiredAction, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiredAction.Get(), o.RequiredAction.IsSet()
}

// SetRequiredAction sets field value
func (o *RunObject) SetRequiredAction(v RunObjectRequiredAction) {
	o.RequiredAction.Set(&v)
}

// GetLastError returns the LastError field value
// If the value is explicit nil, the zero value for RunObjectLastError will be returned
func (o *RunObject) GetLastError() RunObjectLastError {
	if o == nil || o.LastError.Get() == nil {
		var ret RunObjectLastError
		return ret
	}

	return *o.LastError.Get()
}

// GetLastErrorOk returns a tuple with the LastError field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunObject) GetLastErrorOk() (*RunObjectLastError, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastError.Get(), o.LastError.IsSet()
}

// SetLastError sets field value
func (o *RunObject) SetLastError(v RunObjectLastError) {
	o.LastError.Set(&v)
}

// GetExpiresAt returns the ExpiresAt field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *RunObject) GetExpiresAt() int32 {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunObject) GetExpiresAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// SetExpiresAt sets field value
func (o *RunObject) SetExpiresAt(v int32) {
	o.ExpiresAt.Set(&v)
}

// GetStartedAt returns the StartedAt field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *RunObject) GetStartedAt() int32 {
	if o == nil || o.StartedAt.Get() == nil {
		var ret int32
		return ret
	}

	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunObject) GetStartedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// SetStartedAt sets field value
func (o *RunObject) SetStartedAt(v int32) {
	o.StartedAt.Set(&v)
}

// GetCancelledAt returns the CancelledAt field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *RunObject) GetCancelledAt() int32 {
	if o == nil || o.CancelledAt.Get() == nil {
		var ret int32
		return ret
	}

	return *o.CancelledAt.Get()
}

// GetCancelledAtOk returns a tuple with the CancelledAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunObject) GetCancelledAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CancelledAt.Get(), o.CancelledAt.IsSet()
}

// SetCancelledAt sets field value
func (o *RunObject) SetCancelledAt(v int32) {
	o.CancelledAt.Set(&v)
}

// GetFailedAt returns the FailedAt field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *RunObject) GetFailedAt() int32 {
	if o == nil || o.FailedAt.Get() == nil {
		var ret int32
		return ret
	}

	return *o.FailedAt.Get()
}

// GetFailedAtOk returns a tuple with the FailedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunObject) GetFailedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailedAt.Get(), o.FailedAt.IsSet()
}

// SetFailedAt sets field value
func (o *RunObject) SetFailedAt(v int32) {
	o.FailedAt.Set(&v)
}

// GetCompletedAt returns the CompletedAt field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *RunObject) GetCompletedAt() int32 {
	if o == nil || o.CompletedAt.Get() == nil {
		var ret int32
		return ret
	}

	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunObject) GetCompletedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// SetCompletedAt sets field value
func (o *RunObject) SetCompletedAt(v int32) {
	o.CompletedAt.Set(&v)
}

// GetIncompleteDetails returns the IncompleteDetails field value
// If the value is explicit nil, the zero value for RunObjectIncompleteDetails will be returned
func (o *RunObject) GetIncompleteDetails() RunObjectIncompleteDetails {
	if o == nil || o.IncompleteDetails.Get() == nil {
		var ret RunObjectIncompleteDetails
		return ret
	}

	return *o.IncompleteDetails.Get()
}

// GetIncompleteDetailsOk returns a tuple with the IncompleteDetails field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunObject) GetIncompleteDetailsOk() (*RunObjectIncompleteDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncompleteDetails.Get(), o.IncompleteDetails.IsSet()
}

// SetIncompleteDetails sets field value
func (o *RunObject) SetIncompleteDetails(v RunObjectIncompleteDetails) {
	o.IncompleteDetails.Set(&v)
}

// GetModel returns the Model field value
func (o *RunObject) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *RunObject) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *RunObject) SetModel(v string) {
	o.Model = v
}

// GetInstructions returns the Instructions field value
func (o *RunObject) GetInstructions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value
// and a boolean to check if the value has been set.
func (o *RunObject) GetInstructionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instructions, true
}

// SetInstructions sets field value
func (o *RunObject) SetInstructions(v string) {
	o.Instructions = v
}

// GetTools returns the Tools field value
func (o *RunObject) GetTools() []AssistantObjectToolsInner {
	if o == nil {
		var ret []AssistantObjectToolsInner
		return ret
	}

	return o.Tools
}

// GetToolsOk returns a tuple with the Tools field value
// and a boolean to check if the value has been set.
func (o *RunObject) GetToolsOk() ([]AssistantObjectToolsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tools, true
}

// SetTools sets field value
func (o *RunObject) SetTools(v []AssistantObjectToolsInner) {
	o.Tools = v
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RunObject) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunObject) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *RunObject) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetUsage returns the Usage field value
// If the value is explicit nil, the zero value for RunCompletionUsage will be returned
func (o *RunObject) GetUsage() RunCompletionUsage {
	if o == nil || o.Usage.Get() == nil {
		var ret RunCompletionUsage
		return ret
	}

	return *o.Usage.Get()
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunObject) GetUsageOk() (*RunCompletionUsage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Usage.Get(), o.Usage.IsSet()
}

// SetUsage sets field value
func (o *RunObject) SetUsage(v RunCompletionUsage) {
	o.Usage.Set(&v)
}

// GetTemperature returns the Temperature field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunObject) GetTemperature() float32 {
	if o == nil || IsNil(o.Temperature.Get()) {
		var ret float32
		return ret
	}
	return *o.Temperature.Get()
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunObject) GetTemperatureOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Temperature.Get(), o.Temperature.IsSet()
}

// HasTemperature returns a boolean if a field has been set.
func (o *RunObject) HasTemperature() bool {
	if o != nil && o.Temperature.IsSet() {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given NullableFloat32 and assigns it to the Temperature field.
func (o *RunObject) SetTemperature(v float32) {
	o.Temperature.Set(&v)
}
// SetTemperatureNil sets the value for Temperature to be an explicit nil
func (o *RunObject) SetTemperatureNil() {
	o.Temperature.Set(nil)
}

// UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
func (o *RunObject) UnsetTemperature() {
	o.Temperature.Unset()
}

// GetTopP returns the TopP field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunObject) GetTopP() float32 {
	if o == nil || IsNil(o.TopP.Get()) {
		var ret float32
		return ret
	}
	return *o.TopP.Get()
}

// GetTopPOk returns a tuple with the TopP field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunObject) GetTopPOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopP.Get(), o.TopP.IsSet()
}

// HasTopP returns a boolean if a field has been set.
func (o *RunObject) HasTopP() bool {
	if o != nil && o.TopP.IsSet() {
		return true
	}

	return false
}

// SetTopP gets a reference to the given NullableFloat32 and assigns it to the TopP field.
func (o *RunObject) SetTopP(v float32) {
	o.TopP.Set(&v)
}
// SetTopPNil sets the value for TopP to be an explicit nil
func (o *RunObject) SetTopPNil() {
	o.TopP.Set(nil)
}

// UnsetTopP ensures that no value is present for TopP, not even an explicit nil
func (o *RunObject) UnsetTopP() {
	o.TopP.Unset()
}

// GetMaxPromptTokens returns the MaxPromptTokens field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *RunObject) GetMaxPromptTokens() int32 {
	if o == nil || o.MaxPromptTokens.Get() == nil {
		var ret int32
		return ret
	}

	return *o.MaxPromptTokens.Get()
}

// GetMaxPromptTokensOk returns a tuple with the MaxPromptTokens field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunObject) GetMaxPromptTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxPromptTokens.Get(), o.MaxPromptTokens.IsSet()
}

// SetMaxPromptTokens sets field value
func (o *RunObject) SetMaxPromptTokens(v int32) {
	o.MaxPromptTokens.Set(&v)
}

// GetMaxCompletionTokens returns the MaxCompletionTokens field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *RunObject) GetMaxCompletionTokens() int32 {
	if o == nil || o.MaxCompletionTokens.Get() == nil {
		var ret int32
		return ret
	}

	return *o.MaxCompletionTokens.Get()
}

// GetMaxCompletionTokensOk returns a tuple with the MaxCompletionTokens field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunObject) GetMaxCompletionTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxCompletionTokens.Get(), o.MaxCompletionTokens.IsSet()
}

// SetMaxCompletionTokens sets field value
func (o *RunObject) SetMaxCompletionTokens(v int32) {
	o.MaxCompletionTokens.Set(&v)
}

// GetTruncationStrategy returns the TruncationStrategy field value
func (o *RunObject) GetTruncationStrategy() TruncationObject {
	if o == nil {
		var ret TruncationObject
		return ret
	}

	return o.TruncationStrategy
}

// GetTruncationStrategyOk returns a tuple with the TruncationStrategy field value
// and a boolean to check if the value has been set.
func (o *RunObject) GetTruncationStrategyOk() (*TruncationObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TruncationStrategy, true
}

// SetTruncationStrategy sets field value
func (o *RunObject) SetTruncationStrategy(v TruncationObject) {
	o.TruncationStrategy = v
}

// GetToolChoice returns the ToolChoice field value
func (o *RunObject) GetToolChoice() AssistantsApiToolChoiceOption {
	if o == nil {
		var ret AssistantsApiToolChoiceOption
		return ret
	}

	return o.ToolChoice
}

// GetToolChoiceOk returns a tuple with the ToolChoice field value
// and a boolean to check if the value has been set.
func (o *RunObject) GetToolChoiceOk() (*AssistantsApiToolChoiceOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolChoice, true
}

// SetToolChoice sets field value
func (o *RunObject) SetToolChoice(v AssistantsApiToolChoiceOption) {
	o.ToolChoice = v
}

// GetParallelToolCalls returns the ParallelToolCalls field value
func (o *RunObject) GetParallelToolCalls() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ParallelToolCalls
}

// GetParallelToolCallsOk returns a tuple with the ParallelToolCalls field value
// and a boolean to check if the value has been set.
func (o *RunObject) GetParallelToolCallsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParallelToolCalls, true
}

// SetParallelToolCalls sets field value
func (o *RunObject) SetParallelToolCalls(v bool) {
	o.ParallelToolCalls = v
}

// GetResponseFormat returns the ResponseFormat field value
func (o *RunObject) GetResponseFormat() AssistantsApiResponseFormatOption {
	if o == nil {
		var ret AssistantsApiResponseFormatOption
		return ret
	}

	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *RunObject) GetResponseFormatOk() (*AssistantsApiResponseFormatOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value
func (o *RunObject) SetResponseFormat(v AssistantsApiResponseFormatOption) {
	o.ResponseFormat = v
}

func (o RunObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["thread_id"] = o.ThreadId
	toSerialize["assistant_id"] = o.AssistantId
	toSerialize["status"] = o.Status
	toSerialize["required_action"] = o.RequiredAction.Get()
	toSerialize["last_error"] = o.LastError.Get()
	toSerialize["expires_at"] = o.ExpiresAt.Get()
	toSerialize["started_at"] = o.StartedAt.Get()
	toSerialize["cancelled_at"] = o.CancelledAt.Get()
	toSerialize["failed_at"] = o.FailedAt.Get()
	toSerialize["completed_at"] = o.CompletedAt.Get()
	toSerialize["incomplete_details"] = o.IncompleteDetails.Get()
	toSerialize["model"] = o.Model
	toSerialize["instructions"] = o.Instructions
	toSerialize["tools"] = o.Tools
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["usage"] = o.Usage.Get()
	if o.Temperature.IsSet() {
		toSerialize["temperature"] = o.Temperature.Get()
	}
	if o.TopP.IsSet() {
		toSerialize["top_p"] = o.TopP.Get()
	}
	toSerialize["max_prompt_tokens"] = o.MaxPromptTokens.Get()
	toSerialize["max_completion_tokens"] = o.MaxCompletionTokens.Get()
	toSerialize["truncation_strategy"] = o.TruncationStrategy
	toSerialize["tool_choice"] = o.ToolChoice
	toSerialize["parallel_tool_calls"] = o.ParallelToolCalls
	toSerialize["response_format"] = o.ResponseFormat
	return toSerialize, nil
}

func (o *RunObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"created_at",
		"thread_id",
		"assistant_id",
		"status",
		"required_action",
		"last_error",
		"expires_at",
		"started_at",
		"cancelled_at",
		"failed_at",
		"completed_at",
		"incomplete_details",
		"model",
		"instructions",
		"tools",
		"metadata",
		"usage",
		"max_prompt_tokens",
		"max_completion_tokens",
		"truncation_strategy",
		"tool_choice",
		"parallel_tool_calls",
		"response_format",
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

	varRunObject := _RunObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunObject)

	if err != nil {
		return err
	}

	*o = RunObject(varRunObject)

	return err
}

type NullableRunObject struct {
	value *RunObject
	isSet bool
}

func (v NullableRunObject) Get() *RunObject {
	return v.value
}

func (v *NullableRunObject) Set(val *RunObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRunObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRunObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunObject(val *RunObject) *NullableRunObject {
	return &NullableRunObject{value: val, isSet: true}
}

func (v NullableRunObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



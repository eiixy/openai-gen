# ChatCompletionResponseMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **NullableString** | The contents of the message. | 
**ToolCalls** | Pointer to [**[]ChatCompletionMessageToolCall**](ChatCompletionMessageToolCall.md) | The tool calls generated by the model, such as function calls. | [optional] 
**Role** | **string** | The role of the author of this message. | 
**FunctionCall** | Pointer to [**ChatCompletionResponseMessageFunctionCall**](ChatCompletionResponseMessageFunctionCall.md) |  | [optional] 

## Methods

### NewChatCompletionResponseMessage

`func NewChatCompletionResponseMessage(content NullableString, role string, ) *ChatCompletionResponseMessage`

NewChatCompletionResponseMessage instantiates a new ChatCompletionResponseMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionResponseMessageWithDefaults

`func NewChatCompletionResponseMessageWithDefaults() *ChatCompletionResponseMessage`

NewChatCompletionResponseMessageWithDefaults instantiates a new ChatCompletionResponseMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ChatCompletionResponseMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ChatCompletionResponseMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ChatCompletionResponseMessage) SetContent(v string)`

SetContent sets Content field to given value.


### SetContentNil

`func (o *ChatCompletionResponseMessage) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *ChatCompletionResponseMessage) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetToolCalls

`func (o *ChatCompletionResponseMessage) GetToolCalls() []ChatCompletionMessageToolCall`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *ChatCompletionResponseMessage) GetToolCallsOk() (*[]ChatCompletionMessageToolCall, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *ChatCompletionResponseMessage) SetToolCalls(v []ChatCompletionMessageToolCall)`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *ChatCompletionResponseMessage) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.

### GetRole

`func (o *ChatCompletionResponseMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ChatCompletionResponseMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ChatCompletionResponseMessage) SetRole(v string)`

SetRole sets Role field to given value.


### GetFunctionCall

`func (o *ChatCompletionResponseMessage) GetFunctionCall() ChatCompletionResponseMessageFunctionCall`

GetFunctionCall returns the FunctionCall field if non-nil, zero value otherwise.

### GetFunctionCallOk

`func (o *ChatCompletionResponseMessage) GetFunctionCallOk() (*ChatCompletionResponseMessageFunctionCall, bool)`

GetFunctionCallOk returns a tuple with the FunctionCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionCall

`func (o *ChatCompletionResponseMessage) SetFunctionCall(v ChatCompletionResponseMessageFunctionCall)`

SetFunctionCall sets FunctionCall field to given value.

### HasFunctionCall

`func (o *ChatCompletionResponseMessage) HasFunctionCall() bool`

HasFunctionCall returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



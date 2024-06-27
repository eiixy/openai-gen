/*
OpenAI API

The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openai

import (
	"encoding/json"
	"fmt"
)

// ChatCompletionRequestMessage - struct for ChatCompletionRequestMessage
type ChatCompletionRequestMessage struct {
	ChatCompletionRequestAssistantMessage *ChatCompletionRequestAssistantMessage
	ChatCompletionRequestFunctionMessage *ChatCompletionRequestFunctionMessage
	ChatCompletionRequestSystemMessage *ChatCompletionRequestSystemMessage
	ChatCompletionRequestToolMessage *ChatCompletionRequestToolMessage
	ChatCompletionRequestUserMessage *ChatCompletionRequestUserMessage
}

// ChatCompletionRequestAssistantMessageAsChatCompletionRequestMessage is a convenience function that returns ChatCompletionRequestAssistantMessage wrapped in ChatCompletionRequestMessage
func ChatCompletionRequestAssistantMessageAsChatCompletionRequestMessage(v *ChatCompletionRequestAssistantMessage) ChatCompletionRequestMessage {
	return ChatCompletionRequestMessage{
		ChatCompletionRequestAssistantMessage: v,
	}
}

// ChatCompletionRequestFunctionMessageAsChatCompletionRequestMessage is a convenience function that returns ChatCompletionRequestFunctionMessage wrapped in ChatCompletionRequestMessage
func ChatCompletionRequestFunctionMessageAsChatCompletionRequestMessage(v *ChatCompletionRequestFunctionMessage) ChatCompletionRequestMessage {
	return ChatCompletionRequestMessage{
		ChatCompletionRequestFunctionMessage: v,
	}
}

// ChatCompletionRequestSystemMessageAsChatCompletionRequestMessage is a convenience function that returns ChatCompletionRequestSystemMessage wrapped in ChatCompletionRequestMessage
func ChatCompletionRequestSystemMessageAsChatCompletionRequestMessage(v *ChatCompletionRequestSystemMessage) ChatCompletionRequestMessage {
	return ChatCompletionRequestMessage{
		ChatCompletionRequestSystemMessage: v,
	}
}

// ChatCompletionRequestToolMessageAsChatCompletionRequestMessage is a convenience function that returns ChatCompletionRequestToolMessage wrapped in ChatCompletionRequestMessage
func ChatCompletionRequestToolMessageAsChatCompletionRequestMessage(v *ChatCompletionRequestToolMessage) ChatCompletionRequestMessage {
	return ChatCompletionRequestMessage{
		ChatCompletionRequestToolMessage: v,
	}
}

// ChatCompletionRequestUserMessageAsChatCompletionRequestMessage is a convenience function that returns ChatCompletionRequestUserMessage wrapped in ChatCompletionRequestMessage
func ChatCompletionRequestUserMessageAsChatCompletionRequestMessage(v *ChatCompletionRequestUserMessage) ChatCompletionRequestMessage {
	return ChatCompletionRequestMessage{
		ChatCompletionRequestUserMessage: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ChatCompletionRequestMessage) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ChatCompletionRequestAssistantMessage
	err = newStrictDecoder(data).Decode(&dst.ChatCompletionRequestAssistantMessage)
	if err == nil {
		jsonChatCompletionRequestAssistantMessage, _ := json.Marshal(dst.ChatCompletionRequestAssistantMessage)
		if string(jsonChatCompletionRequestAssistantMessage) == "{}" { // empty struct
			dst.ChatCompletionRequestAssistantMessage = nil
		} else {
			match++
		}
	} else {
		dst.ChatCompletionRequestAssistantMessage = nil
	}

	// try to unmarshal data into ChatCompletionRequestFunctionMessage
	err = newStrictDecoder(data).Decode(&dst.ChatCompletionRequestFunctionMessage)
	if err == nil {
		jsonChatCompletionRequestFunctionMessage, _ := json.Marshal(dst.ChatCompletionRequestFunctionMessage)
		if string(jsonChatCompletionRequestFunctionMessage) == "{}" { // empty struct
			dst.ChatCompletionRequestFunctionMessage = nil
		} else {
			match++
		}
	} else {
		dst.ChatCompletionRequestFunctionMessage = nil
	}

	// try to unmarshal data into ChatCompletionRequestSystemMessage
	err = newStrictDecoder(data).Decode(&dst.ChatCompletionRequestSystemMessage)
	if err == nil {
		jsonChatCompletionRequestSystemMessage, _ := json.Marshal(dst.ChatCompletionRequestSystemMessage)
		if string(jsonChatCompletionRequestSystemMessage) == "{}" { // empty struct
			dst.ChatCompletionRequestSystemMessage = nil
		} else {
			match++
		}
	} else {
		dst.ChatCompletionRequestSystemMessage = nil
	}

	// try to unmarshal data into ChatCompletionRequestToolMessage
	err = newStrictDecoder(data).Decode(&dst.ChatCompletionRequestToolMessage)
	if err == nil {
		jsonChatCompletionRequestToolMessage, _ := json.Marshal(dst.ChatCompletionRequestToolMessage)
		if string(jsonChatCompletionRequestToolMessage) == "{}" { // empty struct
			dst.ChatCompletionRequestToolMessage = nil
		} else {
			match++
		}
	} else {
		dst.ChatCompletionRequestToolMessage = nil
	}

	// try to unmarshal data into ChatCompletionRequestUserMessage
	err = newStrictDecoder(data).Decode(&dst.ChatCompletionRequestUserMessage)
	if err == nil {
		jsonChatCompletionRequestUserMessage, _ := json.Marshal(dst.ChatCompletionRequestUserMessage)
		if string(jsonChatCompletionRequestUserMessage) == "{}" { // empty struct
			dst.ChatCompletionRequestUserMessage = nil
		} else {
			match++
		}
	} else {
		dst.ChatCompletionRequestUserMessage = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ChatCompletionRequestAssistantMessage = nil
		dst.ChatCompletionRequestFunctionMessage = nil
		dst.ChatCompletionRequestSystemMessage = nil
		dst.ChatCompletionRequestToolMessage = nil
		dst.ChatCompletionRequestUserMessage = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ChatCompletionRequestMessage)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ChatCompletionRequestMessage)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ChatCompletionRequestMessage) MarshalJSON() ([]byte, error) {
	if src.ChatCompletionRequestAssistantMessage != nil {
		return json.Marshal(&src.ChatCompletionRequestAssistantMessage)
	}

	if src.ChatCompletionRequestFunctionMessage != nil {
		return json.Marshal(&src.ChatCompletionRequestFunctionMessage)
	}

	if src.ChatCompletionRequestSystemMessage != nil {
		return json.Marshal(&src.ChatCompletionRequestSystemMessage)
	}

	if src.ChatCompletionRequestToolMessage != nil {
		return json.Marshal(&src.ChatCompletionRequestToolMessage)
	}

	if src.ChatCompletionRequestUserMessage != nil {
		return json.Marshal(&src.ChatCompletionRequestUserMessage)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ChatCompletionRequestMessage) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ChatCompletionRequestAssistantMessage != nil {
		return obj.ChatCompletionRequestAssistantMessage
	}

	if obj.ChatCompletionRequestFunctionMessage != nil {
		return obj.ChatCompletionRequestFunctionMessage
	}

	if obj.ChatCompletionRequestSystemMessage != nil {
		return obj.ChatCompletionRequestSystemMessage
	}

	if obj.ChatCompletionRequestToolMessage != nil {
		return obj.ChatCompletionRequestToolMessage
	}

	if obj.ChatCompletionRequestUserMessage != nil {
		return obj.ChatCompletionRequestUserMessage
	}

	// all schemas are nil
	return nil
}

type NullableChatCompletionRequestMessage struct {
	value *ChatCompletionRequestMessage
	isSet bool
}

func (v NullableChatCompletionRequestMessage) Get() *ChatCompletionRequestMessage {
	return v.value
}

func (v *NullableChatCompletionRequestMessage) Set(val *ChatCompletionRequestMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableChatCompletionRequestMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableChatCompletionRequestMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatCompletionRequestMessage(val *ChatCompletionRequestMessage) *NullableChatCompletionRequestMessage {
	return &NullableChatCompletionRequestMessage{value: val, isSet: true}
}

func (v NullableChatCompletionRequestMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatCompletionRequestMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// VectorStoreFileObjectChunkingStrategy - The strategy used to chunk the file.
type VectorStoreFileObjectChunkingStrategy struct {
	OtherChunkingStrategyResponseParam *OtherChunkingStrategyResponseParam
	StaticChunkingStrategyResponseParam *StaticChunkingStrategyResponseParam
}

// OtherChunkingStrategyResponseParamAsVectorStoreFileObjectChunkingStrategy is a convenience function that returns OtherChunkingStrategyResponseParam wrapped in VectorStoreFileObjectChunkingStrategy
func OtherChunkingStrategyResponseParamAsVectorStoreFileObjectChunkingStrategy(v *OtherChunkingStrategyResponseParam) VectorStoreFileObjectChunkingStrategy {
	return VectorStoreFileObjectChunkingStrategy{
		OtherChunkingStrategyResponseParam: v,
	}
}

// StaticChunkingStrategyResponseParamAsVectorStoreFileObjectChunkingStrategy is a convenience function that returns StaticChunkingStrategyResponseParam wrapped in VectorStoreFileObjectChunkingStrategy
func StaticChunkingStrategyResponseParamAsVectorStoreFileObjectChunkingStrategy(v *StaticChunkingStrategyResponseParam) VectorStoreFileObjectChunkingStrategy {
	return VectorStoreFileObjectChunkingStrategy{
		StaticChunkingStrategyResponseParam: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *VectorStoreFileObjectChunkingStrategy) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OtherChunkingStrategyResponseParam
	err = newStrictDecoder(data).Decode(&dst.OtherChunkingStrategyResponseParam)
	if err == nil {
		jsonOtherChunkingStrategyResponseParam, _ := json.Marshal(dst.OtherChunkingStrategyResponseParam)
		if string(jsonOtherChunkingStrategyResponseParam) == "{}" { // empty struct
			dst.OtherChunkingStrategyResponseParam = nil
		} else {
			match++
		}
	} else {
		dst.OtherChunkingStrategyResponseParam = nil
	}

	// try to unmarshal data into StaticChunkingStrategyResponseParam
	err = newStrictDecoder(data).Decode(&dst.StaticChunkingStrategyResponseParam)
	if err == nil {
		jsonStaticChunkingStrategyResponseParam, _ := json.Marshal(dst.StaticChunkingStrategyResponseParam)
		if string(jsonStaticChunkingStrategyResponseParam) == "{}" { // empty struct
			dst.StaticChunkingStrategyResponseParam = nil
		} else {
			match++
		}
	} else {
		dst.StaticChunkingStrategyResponseParam = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.OtherChunkingStrategyResponseParam = nil
		dst.StaticChunkingStrategyResponseParam = nil

		return fmt.Errorf("data matches more than one schema in oneOf(VectorStoreFileObjectChunkingStrategy)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(VectorStoreFileObjectChunkingStrategy)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VectorStoreFileObjectChunkingStrategy) MarshalJSON() ([]byte, error) {
	if src.OtherChunkingStrategyResponseParam != nil {
		return json.Marshal(&src.OtherChunkingStrategyResponseParam)
	}

	if src.StaticChunkingStrategyResponseParam != nil {
		return json.Marshal(&src.StaticChunkingStrategyResponseParam)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VectorStoreFileObjectChunkingStrategy) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.OtherChunkingStrategyResponseParam != nil {
		return obj.OtherChunkingStrategyResponseParam
	}

	if obj.StaticChunkingStrategyResponseParam != nil {
		return obj.StaticChunkingStrategyResponseParam
	}

	// all schemas are nil
	return nil
}

type NullableVectorStoreFileObjectChunkingStrategy struct {
	value *VectorStoreFileObjectChunkingStrategy
	isSet bool
}

func (v NullableVectorStoreFileObjectChunkingStrategy) Get() *VectorStoreFileObjectChunkingStrategy {
	return v.value
}

func (v *NullableVectorStoreFileObjectChunkingStrategy) Set(val *VectorStoreFileObjectChunkingStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableVectorStoreFileObjectChunkingStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableVectorStoreFileObjectChunkingStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVectorStoreFileObjectChunkingStrategy(val *VectorStoreFileObjectChunkingStrategy) *NullableVectorStoreFileObjectChunkingStrategy {
	return &NullableVectorStoreFileObjectChunkingStrategy{value: val, isSet: true}
}

func (v NullableVectorStoreFileObjectChunkingStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVectorStoreFileObjectChunkingStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



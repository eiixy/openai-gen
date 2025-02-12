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

// checks if the DoneEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DoneEvent{}

// DoneEvent Occurs when a stream ends.
type DoneEvent struct {
	Event string `json:"event"`
	Data string `json:"data"`
}

type _DoneEvent DoneEvent

// NewDoneEvent instantiates a new DoneEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDoneEvent(event string, data string) *DoneEvent {
	this := DoneEvent{}
	this.Event = event
	this.Data = data
	return &this
}

// NewDoneEventWithDefaults instantiates a new DoneEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDoneEventWithDefaults() *DoneEvent {
	this := DoneEvent{}
	return &this
}

// GetEvent returns the Event field value
func (o *DoneEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *DoneEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *DoneEvent) SetEvent(v string) {
	o.Event = v
}

// GetData returns the Data field value
func (o *DoneEvent) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DoneEvent) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DoneEvent) SetData(v string) {
	o.Data = v
}

func (o DoneEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DoneEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *DoneEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event",
		"data",
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

	varDoneEvent := _DoneEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDoneEvent)

	if err != nil {
		return err
	}

	*o = DoneEvent(varDoneEvent)

	return err
}

type NullableDoneEvent struct {
	value *DoneEvent
	isSet bool
}

func (v NullableDoneEvent) Get() *DoneEvent {
	return v.value
}

func (v *NullableDoneEvent) Set(val *DoneEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableDoneEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableDoneEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDoneEvent(val *DoneEvent) *NullableDoneEvent {
	return &NullableDoneEvent{value: val, isSet: true}
}

func (v NullableDoneEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDoneEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



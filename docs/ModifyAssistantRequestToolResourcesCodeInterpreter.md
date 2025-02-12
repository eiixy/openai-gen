# ModifyAssistantRequestToolResourcesCodeInterpreter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileIds** | Pointer to **[]string** | Overrides the list of [file](/docs/api-reference/files) IDs made available to the &#x60;code_interpreter&#x60; tool. There can be a maximum of 20 files associated with the tool.  | [optional] [default to []]

## Methods

### NewModifyAssistantRequestToolResourcesCodeInterpreter

`func NewModifyAssistantRequestToolResourcesCodeInterpreter() *ModifyAssistantRequestToolResourcesCodeInterpreter`

NewModifyAssistantRequestToolResourcesCodeInterpreter instantiates a new ModifyAssistantRequestToolResourcesCodeInterpreter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyAssistantRequestToolResourcesCodeInterpreterWithDefaults

`func NewModifyAssistantRequestToolResourcesCodeInterpreterWithDefaults() *ModifyAssistantRequestToolResourcesCodeInterpreter`

NewModifyAssistantRequestToolResourcesCodeInterpreterWithDefaults instantiates a new ModifyAssistantRequestToolResourcesCodeInterpreter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileIds

`func (o *ModifyAssistantRequestToolResourcesCodeInterpreter) GetFileIds() []string`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *ModifyAssistantRequestToolResourcesCodeInterpreter) GetFileIdsOk() (*[]string, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *ModifyAssistantRequestToolResourcesCodeInterpreter) SetFileIds(v []string)`

SetFileIds sets FileIds field to given value.

### HasFileIds

`func (o *ModifyAssistantRequestToolResourcesCodeInterpreter) HasFileIds() bool`

HasFileIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



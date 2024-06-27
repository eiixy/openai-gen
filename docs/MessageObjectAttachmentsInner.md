# MessageObjectAttachmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | Pointer to **string** | The ID of the file to attach to the message. | [optional] 
**Tools** | Pointer to [**[]MessageObjectAttachmentsInnerToolsInner**](MessageObjectAttachmentsInnerToolsInner.md) | The tools to add this file to. | [optional] 

## Methods

### NewMessageObjectAttachmentsInner

`func NewMessageObjectAttachmentsInner() *MessageObjectAttachmentsInner`

NewMessageObjectAttachmentsInner instantiates a new MessageObjectAttachmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageObjectAttachmentsInnerWithDefaults

`func NewMessageObjectAttachmentsInnerWithDefaults() *MessageObjectAttachmentsInner`

NewMessageObjectAttachmentsInnerWithDefaults instantiates a new MessageObjectAttachmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *MessageObjectAttachmentsInner) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *MessageObjectAttachmentsInner) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *MessageObjectAttachmentsInner) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *MessageObjectAttachmentsInner) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetTools

`func (o *MessageObjectAttachmentsInner) GetTools() []MessageObjectAttachmentsInnerToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *MessageObjectAttachmentsInner) GetToolsOk() (*[]MessageObjectAttachmentsInnerToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *MessageObjectAttachmentsInner) SetTools(v []MessageObjectAttachmentsInnerToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *MessageObjectAttachmentsInner) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



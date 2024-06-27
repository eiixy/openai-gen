# AssistantToolsFileSearchFileSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxNumResults** | Pointer to **int32** | The maximum number of results the file search tool should output. The default is 20 for gpt-4* models and 5 for gpt-3.5-turbo. This number should be between 1 and 50 inclusive.  Note that the file search tool may output fewer than &#x60;max_num_results&#x60; results. See the [file search tool documentation](/docs/assistants/tools/file-search/number-of-chunks-returned) for more information.  | [optional] 

## Methods

### NewAssistantToolsFileSearchFileSearch

`func NewAssistantToolsFileSearchFileSearch() *AssistantToolsFileSearchFileSearch`

NewAssistantToolsFileSearchFileSearch instantiates a new AssistantToolsFileSearchFileSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantToolsFileSearchFileSearchWithDefaults

`func NewAssistantToolsFileSearchFileSearchWithDefaults() *AssistantToolsFileSearchFileSearch`

NewAssistantToolsFileSearchFileSearchWithDefaults instantiates a new AssistantToolsFileSearchFileSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxNumResults

`func (o *AssistantToolsFileSearchFileSearch) GetMaxNumResults() int32`

GetMaxNumResults returns the MaxNumResults field if non-nil, zero value otherwise.

### GetMaxNumResultsOk

`func (o *AssistantToolsFileSearchFileSearch) GetMaxNumResultsOk() (*int32, bool)`

GetMaxNumResultsOk returns a tuple with the MaxNumResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumResults

`func (o *AssistantToolsFileSearchFileSearch) SetMaxNumResults(v int32)`

SetMaxNumResults sets MaxNumResults field to given value.

### HasMaxNumResults

`func (o *AssistantToolsFileSearchFileSearch) HasMaxNumResults() bool`

HasMaxNumResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



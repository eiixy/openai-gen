# StaticChunkingStrategyRequestParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Always &#x60;static&#x60;. | 
**Static** | [**StaticChunkingStrategy**](StaticChunkingStrategy.md) |  | 

## Methods

### NewStaticChunkingStrategyRequestParam

`func NewStaticChunkingStrategyRequestParam(type_ string, static StaticChunkingStrategy, ) *StaticChunkingStrategyRequestParam`

NewStaticChunkingStrategyRequestParam instantiates a new StaticChunkingStrategyRequestParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticChunkingStrategyRequestParamWithDefaults

`func NewStaticChunkingStrategyRequestParamWithDefaults() *StaticChunkingStrategyRequestParam`

NewStaticChunkingStrategyRequestParamWithDefaults instantiates a new StaticChunkingStrategyRequestParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StaticChunkingStrategyRequestParam) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StaticChunkingStrategyRequestParam) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StaticChunkingStrategyRequestParam) SetType(v string)`

SetType sets Type field to given value.


### GetStatic

`func (o *StaticChunkingStrategyRequestParam) GetStatic() StaticChunkingStrategy`

GetStatic returns the Static field if non-nil, zero value otherwise.

### GetStaticOk

`func (o *StaticChunkingStrategyRequestParam) GetStaticOk() (*StaticChunkingStrategy, bool)`

GetStaticOk returns a tuple with the Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic

`func (o *StaticChunkingStrategyRequestParam) SetStatic(v StaticChunkingStrategy)`

SetStatic sets Static field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SetGroupTopicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The group ID | 
**Topic** | Pointer to **string** | The group topic/description. Leave empty to remove the topic. | [optional] 

## Methods

### NewSetGroupTopicRequest

`func NewSetGroupTopicRequest(groupId string, ) *SetGroupTopicRequest`

NewSetGroupTopicRequest instantiates a new SetGroupTopicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetGroupTopicRequestWithDefaults

`func NewSetGroupTopicRequestWithDefaults() *SetGroupTopicRequest`

NewSetGroupTopicRequestWithDefaults instantiates a new SetGroupTopicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *SetGroupTopicRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SetGroupTopicRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SetGroupTopicRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetTopic

`func (o *SetGroupTopicRequest) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *SetGroupTopicRequest) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *SetGroupTopicRequest) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *SetGroupTopicRequest) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



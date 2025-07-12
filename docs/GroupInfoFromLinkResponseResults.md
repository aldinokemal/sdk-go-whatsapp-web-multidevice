# GroupInfoFromLinkResponseResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | The group ID | [optional] 
**Name** | Pointer to **string** | The group name | [optional] 
**Topic** | Pointer to **string** | The group topic/description | [optional] 
**CreatedAt** | Pointer to **time.Time** | When the group was created | [optional] 
**ParticipantCount** | Pointer to **int32** | Number of participants in the group | [optional] 
**IsLocked** | Pointer to **bool** | Whether the group is locked (only admins can modify group info) | [optional] 
**IsAnnounce** | Pointer to **bool** | Whether the group is in announce mode (only admins can send messages) | [optional] 
**IsEphemeral** | Pointer to **bool** | Whether the group has disappearing messages enabled | [optional] 
**Description** | Pointer to **string** | Additional description of the group | [optional] 

## Methods

### NewGroupInfoFromLinkResponseResults

`func NewGroupInfoFromLinkResponseResults() *GroupInfoFromLinkResponseResults`

NewGroupInfoFromLinkResponseResults instantiates a new GroupInfoFromLinkResponseResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupInfoFromLinkResponseResultsWithDefaults

`func NewGroupInfoFromLinkResponseResultsWithDefaults() *GroupInfoFromLinkResponseResults`

NewGroupInfoFromLinkResponseResultsWithDefaults instantiates a new GroupInfoFromLinkResponseResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GroupInfoFromLinkResponseResults) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupInfoFromLinkResponseResults) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupInfoFromLinkResponseResults) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupInfoFromLinkResponseResults) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *GroupInfoFromLinkResponseResults) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupInfoFromLinkResponseResults) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupInfoFromLinkResponseResults) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupInfoFromLinkResponseResults) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTopic

`func (o *GroupInfoFromLinkResponseResults) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *GroupInfoFromLinkResponseResults) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *GroupInfoFromLinkResponseResults) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *GroupInfoFromLinkResponseResults) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GroupInfoFromLinkResponseResults) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupInfoFromLinkResponseResults) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupInfoFromLinkResponseResults) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupInfoFromLinkResponseResults) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetParticipantCount

`func (o *GroupInfoFromLinkResponseResults) GetParticipantCount() int32`

GetParticipantCount returns the ParticipantCount field if non-nil, zero value otherwise.

### GetParticipantCountOk

`func (o *GroupInfoFromLinkResponseResults) GetParticipantCountOk() (*int32, bool)`

GetParticipantCountOk returns a tuple with the ParticipantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantCount

`func (o *GroupInfoFromLinkResponseResults) SetParticipantCount(v int32)`

SetParticipantCount sets ParticipantCount field to given value.

### HasParticipantCount

`func (o *GroupInfoFromLinkResponseResults) HasParticipantCount() bool`

HasParticipantCount returns a boolean if a field has been set.

### GetIsLocked

`func (o *GroupInfoFromLinkResponseResults) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *GroupInfoFromLinkResponseResults) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *GroupInfoFromLinkResponseResults) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *GroupInfoFromLinkResponseResults) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsAnnounce

`func (o *GroupInfoFromLinkResponseResults) GetIsAnnounce() bool`

GetIsAnnounce returns the IsAnnounce field if non-nil, zero value otherwise.

### GetIsAnnounceOk

`func (o *GroupInfoFromLinkResponseResults) GetIsAnnounceOk() (*bool, bool)`

GetIsAnnounceOk returns a tuple with the IsAnnounce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnnounce

`func (o *GroupInfoFromLinkResponseResults) SetIsAnnounce(v bool)`

SetIsAnnounce sets IsAnnounce field to given value.

### HasIsAnnounce

`func (o *GroupInfoFromLinkResponseResults) HasIsAnnounce() bool`

HasIsAnnounce returns a boolean if a field has been set.

### GetIsEphemeral

`func (o *GroupInfoFromLinkResponseResults) GetIsEphemeral() bool`

GetIsEphemeral returns the IsEphemeral field if non-nil, zero value otherwise.

### GetIsEphemeralOk

`func (o *GroupInfoFromLinkResponseResults) GetIsEphemeralOk() (*bool, bool)`

GetIsEphemeralOk returns a tuple with the IsEphemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEphemeral

`func (o *GroupInfoFromLinkResponseResults) SetIsEphemeral(v bool)`

SetIsEphemeral sets IsEphemeral field to given value.

### HasIsEphemeral

`func (o *GroupInfoFromLinkResponseResults) HasIsEphemeral() bool`

HasIsEphemeral returns a boolean if a field has been set.

### GetDescription

`func (o *GroupInfoFromLinkResponseResults) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupInfoFromLinkResponseResults) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupInfoFromLinkResponseResults) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupInfoFromLinkResponseResults) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GroupParticipantsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Participants** | Pointer to [**[]GroupParticipantItem**](GroupParticipantItem.md) |  | [optional] 

## Methods

### NewGroupParticipantsResult

`func NewGroupParticipantsResult() *GroupParticipantsResult`

NewGroupParticipantsResult instantiates a new GroupParticipantsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupParticipantsResultWithDefaults

`func NewGroupParticipantsResultWithDefaults() *GroupParticipantsResult`

NewGroupParticipantsResultWithDefaults instantiates a new GroupParticipantsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GroupParticipantsResult) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupParticipantsResult) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupParticipantsResult) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupParticipantsResult) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *GroupParticipantsResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupParticipantsResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupParticipantsResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupParticipantsResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParticipants

`func (o *GroupParticipantsResult) GetParticipants() []GroupParticipantItem`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *GroupParticipantsResult) GetParticipantsOk() (*[]GroupParticipantItem, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *GroupParticipantsResult) SetParticipants(v []GroupParticipantItem)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *GroupParticipantsResult) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



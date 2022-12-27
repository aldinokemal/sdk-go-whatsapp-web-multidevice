# UserGroupResponseResultsDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JID** | Pointer to **string** |  | [optional] 
**OwnerJID** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NameSetAt** | Pointer to **string** |  | [optional] 
**NameSetBy** | Pointer to **string** |  | [optional] 
**GroupCreated** | Pointer to **string** |  | [optional] 
**ParticipantVersionID** | Pointer to **string** |  | [optional] 
**Participants** | Pointer to [**[]UserGroupResponseResultsDataInnerParticipantsInner**](UserGroupResponseResultsDataInnerParticipantsInner.md) |  | [optional] 

## Methods

### NewUserGroupResponseResultsDataInner

`func NewUserGroupResponseResultsDataInner() *UserGroupResponseResultsDataInner`

NewUserGroupResponseResultsDataInner instantiates a new UserGroupResponseResultsDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupResponseResultsDataInnerWithDefaults

`func NewUserGroupResponseResultsDataInnerWithDefaults() *UserGroupResponseResultsDataInner`

NewUserGroupResponseResultsDataInnerWithDefaults instantiates a new UserGroupResponseResultsDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJID

`func (o *UserGroupResponseResultsDataInner) GetJID() string`

GetJID returns the JID field if non-nil, zero value otherwise.

### GetJIDOk

`func (o *UserGroupResponseResultsDataInner) GetJIDOk() (*string, bool)`

GetJIDOk returns a tuple with the JID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJID

`func (o *UserGroupResponseResultsDataInner) SetJID(v string)`

SetJID sets JID field to given value.

### HasJID

`func (o *UserGroupResponseResultsDataInner) HasJID() bool`

HasJID returns a boolean if a field has been set.

### GetOwnerJID

`func (o *UserGroupResponseResultsDataInner) GetOwnerJID() string`

GetOwnerJID returns the OwnerJID field if non-nil, zero value otherwise.

### GetOwnerJIDOk

`func (o *UserGroupResponseResultsDataInner) GetOwnerJIDOk() (*string, bool)`

GetOwnerJIDOk returns a tuple with the OwnerJID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerJID

`func (o *UserGroupResponseResultsDataInner) SetOwnerJID(v string)`

SetOwnerJID sets OwnerJID field to given value.

### HasOwnerJID

`func (o *UserGroupResponseResultsDataInner) HasOwnerJID() bool`

HasOwnerJID returns a boolean if a field has been set.

### GetName

`func (o *UserGroupResponseResultsDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserGroupResponseResultsDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserGroupResponseResultsDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserGroupResponseResultsDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameSetAt

`func (o *UserGroupResponseResultsDataInner) GetNameSetAt() string`

GetNameSetAt returns the NameSetAt field if non-nil, zero value otherwise.

### GetNameSetAtOk

`func (o *UserGroupResponseResultsDataInner) GetNameSetAtOk() (*string, bool)`

GetNameSetAtOk returns a tuple with the NameSetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSetAt

`func (o *UserGroupResponseResultsDataInner) SetNameSetAt(v string)`

SetNameSetAt sets NameSetAt field to given value.

### HasNameSetAt

`func (o *UserGroupResponseResultsDataInner) HasNameSetAt() bool`

HasNameSetAt returns a boolean if a field has been set.

### GetNameSetBy

`func (o *UserGroupResponseResultsDataInner) GetNameSetBy() string`

GetNameSetBy returns the NameSetBy field if non-nil, zero value otherwise.

### GetNameSetByOk

`func (o *UserGroupResponseResultsDataInner) GetNameSetByOk() (*string, bool)`

GetNameSetByOk returns a tuple with the NameSetBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSetBy

`func (o *UserGroupResponseResultsDataInner) SetNameSetBy(v string)`

SetNameSetBy sets NameSetBy field to given value.

### HasNameSetBy

`func (o *UserGroupResponseResultsDataInner) HasNameSetBy() bool`

HasNameSetBy returns a boolean if a field has been set.

### GetGroupCreated

`func (o *UserGroupResponseResultsDataInner) GetGroupCreated() string`

GetGroupCreated returns the GroupCreated field if non-nil, zero value otherwise.

### GetGroupCreatedOk

`func (o *UserGroupResponseResultsDataInner) GetGroupCreatedOk() (*string, bool)`

GetGroupCreatedOk returns a tuple with the GroupCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCreated

`func (o *UserGroupResponseResultsDataInner) SetGroupCreated(v string)`

SetGroupCreated sets GroupCreated field to given value.

### HasGroupCreated

`func (o *UserGroupResponseResultsDataInner) HasGroupCreated() bool`

HasGroupCreated returns a boolean if a field has been set.

### GetParticipantVersionID

`func (o *UserGroupResponseResultsDataInner) GetParticipantVersionID() string`

GetParticipantVersionID returns the ParticipantVersionID field if non-nil, zero value otherwise.

### GetParticipantVersionIDOk

`func (o *UserGroupResponseResultsDataInner) GetParticipantVersionIDOk() (*string, bool)`

GetParticipantVersionIDOk returns a tuple with the ParticipantVersionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantVersionID

`func (o *UserGroupResponseResultsDataInner) SetParticipantVersionID(v string)`

SetParticipantVersionID sets ParticipantVersionID field to given value.

### HasParticipantVersionID

`func (o *UserGroupResponseResultsDataInner) HasParticipantVersionID() bool`

HasParticipantVersionID returns a boolean if a field has been set.

### GetParticipants

`func (o *UserGroupResponseResultsDataInner) GetParticipants() []UserGroupResponseResultsDataInnerParticipantsInner`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *UserGroupResponseResultsDataInner) GetParticipantsOk() (*[]UserGroupResponseResultsDataInnerParticipantsInner, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *UserGroupResponseResultsDataInner) SetParticipants(v []UserGroupResponseResultsDataInnerParticipantsInner)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *UserGroupResponseResultsDataInner) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



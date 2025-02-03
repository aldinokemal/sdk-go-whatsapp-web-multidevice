# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JID** | Pointer to **string** |  | [optional] 
**OwnerJID** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NameSetAt** | Pointer to **time.Time** |  | [optional] 
**NameSetBy** | Pointer to **string** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**TopicID** | Pointer to **string** |  | [optional] 
**TopicSetAt** | Pointer to **time.Time** |  | [optional] 
**TopicSetBy** | Pointer to **string** |  | [optional] 
**TopicDeleted** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**IsAnnounce** | Pointer to **bool** |  | [optional] 
**AnnounceVersionID** | Pointer to **string** |  | [optional] 
**IsEphemeral** | Pointer to **bool** |  | [optional] 
**DisappearingTimer** | Pointer to **int32** |  | [optional] 
**IsIncognito** | Pointer to **bool** |  | [optional] 
**IsParent** | Pointer to **bool** |  | [optional] 
**DefaultMembershipApprovalMode** | Pointer to **string** |  | [optional] 
**LinkedParentJID** | Pointer to **string** |  | [optional] 
**IsDefaultSubGroup** | Pointer to **bool** |  | [optional] 
**IsJoinApprovalRequired** | Pointer to **bool** |  | [optional] 
**GroupCreated** | Pointer to **time.Time** |  | [optional] 
**ParticipantVersionID** | Pointer to **string** |  | [optional] 
**Participants** | Pointer to [**[]Participant**](Participant.md) |  | [optional] 
**MemberAddMode** | Pointer to **string** |  | [optional] 

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJID

`func (o *Group) GetJID() string`

GetJID returns the JID field if non-nil, zero value otherwise.

### GetJIDOk

`func (o *Group) GetJIDOk() (*string, bool)`

GetJIDOk returns a tuple with the JID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJID

`func (o *Group) SetJID(v string)`

SetJID sets JID field to given value.

### HasJID

`func (o *Group) HasJID() bool`

HasJID returns a boolean if a field has been set.

### GetOwnerJID

`func (o *Group) GetOwnerJID() string`

GetOwnerJID returns the OwnerJID field if non-nil, zero value otherwise.

### GetOwnerJIDOk

`func (o *Group) GetOwnerJIDOk() (*string, bool)`

GetOwnerJIDOk returns a tuple with the OwnerJID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerJID

`func (o *Group) SetOwnerJID(v string)`

SetOwnerJID sets OwnerJID field to given value.

### HasOwnerJID

`func (o *Group) HasOwnerJID() bool`

HasOwnerJID returns a boolean if a field has been set.

### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Group) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameSetAt

`func (o *Group) GetNameSetAt() time.Time`

GetNameSetAt returns the NameSetAt field if non-nil, zero value otherwise.

### GetNameSetAtOk

`func (o *Group) GetNameSetAtOk() (*time.Time, bool)`

GetNameSetAtOk returns a tuple with the NameSetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSetAt

`func (o *Group) SetNameSetAt(v time.Time)`

SetNameSetAt sets NameSetAt field to given value.

### HasNameSetAt

`func (o *Group) HasNameSetAt() bool`

HasNameSetAt returns a boolean if a field has been set.

### GetNameSetBy

`func (o *Group) GetNameSetBy() string`

GetNameSetBy returns the NameSetBy field if non-nil, zero value otherwise.

### GetNameSetByOk

`func (o *Group) GetNameSetByOk() (*string, bool)`

GetNameSetByOk returns a tuple with the NameSetBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSetBy

`func (o *Group) SetNameSetBy(v string)`

SetNameSetBy sets NameSetBy field to given value.

### HasNameSetBy

`func (o *Group) HasNameSetBy() bool`

HasNameSetBy returns a boolean if a field has been set.

### GetTopic

`func (o *Group) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *Group) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *Group) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *Group) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetTopicID

`func (o *Group) GetTopicID() string`

GetTopicID returns the TopicID field if non-nil, zero value otherwise.

### GetTopicIDOk

`func (o *Group) GetTopicIDOk() (*string, bool)`

GetTopicIDOk returns a tuple with the TopicID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicID

`func (o *Group) SetTopicID(v string)`

SetTopicID sets TopicID field to given value.

### HasTopicID

`func (o *Group) HasTopicID() bool`

HasTopicID returns a boolean if a field has been set.

### GetTopicSetAt

`func (o *Group) GetTopicSetAt() time.Time`

GetTopicSetAt returns the TopicSetAt field if non-nil, zero value otherwise.

### GetTopicSetAtOk

`func (o *Group) GetTopicSetAtOk() (*time.Time, bool)`

GetTopicSetAtOk returns a tuple with the TopicSetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicSetAt

`func (o *Group) SetTopicSetAt(v time.Time)`

SetTopicSetAt sets TopicSetAt field to given value.

### HasTopicSetAt

`func (o *Group) HasTopicSetAt() bool`

HasTopicSetAt returns a boolean if a field has been set.

### GetTopicSetBy

`func (o *Group) GetTopicSetBy() string`

GetTopicSetBy returns the TopicSetBy field if non-nil, zero value otherwise.

### GetTopicSetByOk

`func (o *Group) GetTopicSetByOk() (*string, bool)`

GetTopicSetByOk returns a tuple with the TopicSetBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicSetBy

`func (o *Group) SetTopicSetBy(v string)`

SetTopicSetBy sets TopicSetBy field to given value.

### HasTopicSetBy

`func (o *Group) HasTopicSetBy() bool`

HasTopicSetBy returns a boolean if a field has been set.

### GetTopicDeleted

`func (o *Group) GetTopicDeleted() bool`

GetTopicDeleted returns the TopicDeleted field if non-nil, zero value otherwise.

### GetTopicDeletedOk

`func (o *Group) GetTopicDeletedOk() (*bool, bool)`

GetTopicDeletedOk returns a tuple with the TopicDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicDeleted

`func (o *Group) SetTopicDeleted(v bool)`

SetTopicDeleted sets TopicDeleted field to given value.

### HasTopicDeleted

`func (o *Group) HasTopicDeleted() bool`

HasTopicDeleted returns a boolean if a field has been set.

### GetIsLocked

`func (o *Group) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *Group) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *Group) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *Group) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsAnnounce

`func (o *Group) GetIsAnnounce() bool`

GetIsAnnounce returns the IsAnnounce field if non-nil, zero value otherwise.

### GetIsAnnounceOk

`func (o *Group) GetIsAnnounceOk() (*bool, bool)`

GetIsAnnounceOk returns a tuple with the IsAnnounce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnnounce

`func (o *Group) SetIsAnnounce(v bool)`

SetIsAnnounce sets IsAnnounce field to given value.

### HasIsAnnounce

`func (o *Group) HasIsAnnounce() bool`

HasIsAnnounce returns a boolean if a field has been set.

### GetAnnounceVersionID

`func (o *Group) GetAnnounceVersionID() string`

GetAnnounceVersionID returns the AnnounceVersionID field if non-nil, zero value otherwise.

### GetAnnounceVersionIDOk

`func (o *Group) GetAnnounceVersionIDOk() (*string, bool)`

GetAnnounceVersionIDOk returns a tuple with the AnnounceVersionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnounceVersionID

`func (o *Group) SetAnnounceVersionID(v string)`

SetAnnounceVersionID sets AnnounceVersionID field to given value.

### HasAnnounceVersionID

`func (o *Group) HasAnnounceVersionID() bool`

HasAnnounceVersionID returns a boolean if a field has been set.

### GetIsEphemeral

`func (o *Group) GetIsEphemeral() bool`

GetIsEphemeral returns the IsEphemeral field if non-nil, zero value otherwise.

### GetIsEphemeralOk

`func (o *Group) GetIsEphemeralOk() (*bool, bool)`

GetIsEphemeralOk returns a tuple with the IsEphemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEphemeral

`func (o *Group) SetIsEphemeral(v bool)`

SetIsEphemeral sets IsEphemeral field to given value.

### HasIsEphemeral

`func (o *Group) HasIsEphemeral() bool`

HasIsEphemeral returns a boolean if a field has been set.

### GetDisappearingTimer

`func (o *Group) GetDisappearingTimer() int32`

GetDisappearingTimer returns the DisappearingTimer field if non-nil, zero value otherwise.

### GetDisappearingTimerOk

`func (o *Group) GetDisappearingTimerOk() (*int32, bool)`

GetDisappearingTimerOk returns a tuple with the DisappearingTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisappearingTimer

`func (o *Group) SetDisappearingTimer(v int32)`

SetDisappearingTimer sets DisappearingTimer field to given value.

### HasDisappearingTimer

`func (o *Group) HasDisappearingTimer() bool`

HasDisappearingTimer returns a boolean if a field has been set.

### GetIsIncognito

`func (o *Group) GetIsIncognito() bool`

GetIsIncognito returns the IsIncognito field if non-nil, zero value otherwise.

### GetIsIncognitoOk

`func (o *Group) GetIsIncognitoOk() (*bool, bool)`

GetIsIncognitoOk returns a tuple with the IsIncognito field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncognito

`func (o *Group) SetIsIncognito(v bool)`

SetIsIncognito sets IsIncognito field to given value.

### HasIsIncognito

`func (o *Group) HasIsIncognito() bool`

HasIsIncognito returns a boolean if a field has been set.

### GetIsParent

`func (o *Group) GetIsParent() bool`

GetIsParent returns the IsParent field if non-nil, zero value otherwise.

### GetIsParentOk

`func (o *Group) GetIsParentOk() (*bool, bool)`

GetIsParentOk returns a tuple with the IsParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsParent

`func (o *Group) SetIsParent(v bool)`

SetIsParent sets IsParent field to given value.

### HasIsParent

`func (o *Group) HasIsParent() bool`

HasIsParent returns a boolean if a field has been set.

### GetDefaultMembershipApprovalMode

`func (o *Group) GetDefaultMembershipApprovalMode() string`

GetDefaultMembershipApprovalMode returns the DefaultMembershipApprovalMode field if non-nil, zero value otherwise.

### GetDefaultMembershipApprovalModeOk

`func (o *Group) GetDefaultMembershipApprovalModeOk() (*string, bool)`

GetDefaultMembershipApprovalModeOk returns a tuple with the DefaultMembershipApprovalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMembershipApprovalMode

`func (o *Group) SetDefaultMembershipApprovalMode(v string)`

SetDefaultMembershipApprovalMode sets DefaultMembershipApprovalMode field to given value.

### HasDefaultMembershipApprovalMode

`func (o *Group) HasDefaultMembershipApprovalMode() bool`

HasDefaultMembershipApprovalMode returns a boolean if a field has been set.

### GetLinkedParentJID

`func (o *Group) GetLinkedParentJID() string`

GetLinkedParentJID returns the LinkedParentJID field if non-nil, zero value otherwise.

### GetLinkedParentJIDOk

`func (o *Group) GetLinkedParentJIDOk() (*string, bool)`

GetLinkedParentJIDOk returns a tuple with the LinkedParentJID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedParentJID

`func (o *Group) SetLinkedParentJID(v string)`

SetLinkedParentJID sets LinkedParentJID field to given value.

### HasLinkedParentJID

`func (o *Group) HasLinkedParentJID() bool`

HasLinkedParentJID returns a boolean if a field has been set.

### GetIsDefaultSubGroup

`func (o *Group) GetIsDefaultSubGroup() bool`

GetIsDefaultSubGroup returns the IsDefaultSubGroup field if non-nil, zero value otherwise.

### GetIsDefaultSubGroupOk

`func (o *Group) GetIsDefaultSubGroupOk() (*bool, bool)`

GetIsDefaultSubGroupOk returns a tuple with the IsDefaultSubGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultSubGroup

`func (o *Group) SetIsDefaultSubGroup(v bool)`

SetIsDefaultSubGroup sets IsDefaultSubGroup field to given value.

### HasIsDefaultSubGroup

`func (o *Group) HasIsDefaultSubGroup() bool`

HasIsDefaultSubGroup returns a boolean if a field has been set.

### GetIsJoinApprovalRequired

`func (o *Group) GetIsJoinApprovalRequired() bool`

GetIsJoinApprovalRequired returns the IsJoinApprovalRequired field if non-nil, zero value otherwise.

### GetIsJoinApprovalRequiredOk

`func (o *Group) GetIsJoinApprovalRequiredOk() (*bool, bool)`

GetIsJoinApprovalRequiredOk returns a tuple with the IsJoinApprovalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsJoinApprovalRequired

`func (o *Group) SetIsJoinApprovalRequired(v bool)`

SetIsJoinApprovalRequired sets IsJoinApprovalRequired field to given value.

### HasIsJoinApprovalRequired

`func (o *Group) HasIsJoinApprovalRequired() bool`

HasIsJoinApprovalRequired returns a boolean if a field has been set.

### GetGroupCreated

`func (o *Group) GetGroupCreated() time.Time`

GetGroupCreated returns the GroupCreated field if non-nil, zero value otherwise.

### GetGroupCreatedOk

`func (o *Group) GetGroupCreatedOk() (*time.Time, bool)`

GetGroupCreatedOk returns a tuple with the GroupCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCreated

`func (o *Group) SetGroupCreated(v time.Time)`

SetGroupCreated sets GroupCreated field to given value.

### HasGroupCreated

`func (o *Group) HasGroupCreated() bool`

HasGroupCreated returns a boolean if a field has been set.

### GetParticipantVersionID

`func (o *Group) GetParticipantVersionID() string`

GetParticipantVersionID returns the ParticipantVersionID field if non-nil, zero value otherwise.

### GetParticipantVersionIDOk

`func (o *Group) GetParticipantVersionIDOk() (*string, bool)`

GetParticipantVersionIDOk returns a tuple with the ParticipantVersionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantVersionID

`func (o *Group) SetParticipantVersionID(v string)`

SetParticipantVersionID sets ParticipantVersionID field to given value.

### HasParticipantVersionID

`func (o *Group) HasParticipantVersionID() bool`

HasParticipantVersionID returns a boolean if a field has been set.

### GetParticipants

`func (o *Group) GetParticipants() []Participant`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *Group) GetParticipantsOk() (*[]Participant, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *Group) SetParticipants(v []Participant)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *Group) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetMemberAddMode

`func (o *Group) GetMemberAddMode() string`

GetMemberAddMode returns the MemberAddMode field if non-nil, zero value otherwise.

### GetMemberAddModeOk

`func (o *Group) GetMemberAddModeOk() (*string, bool)`

GetMemberAddModeOk returns a tuple with the MemberAddMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberAddMode

`func (o *Group) SetMemberAddMode(v string)`

SetMemberAddMode sets MemberAddMode field to given value.

### HasMemberAddMode

`func (o *Group) HasMemberAddMode() bool`

HasMemberAddMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



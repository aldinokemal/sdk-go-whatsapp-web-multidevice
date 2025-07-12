# ApproveGroupParticipantRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The group ID | 
**Participants** | **[]string** | Array of participant WhatsApp IDs to approve | 

## Methods

### NewApproveGroupParticipantRequestRequest

`func NewApproveGroupParticipantRequestRequest(groupId string, participants []string, ) *ApproveGroupParticipantRequestRequest`

NewApproveGroupParticipantRequestRequest instantiates a new ApproveGroupParticipantRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproveGroupParticipantRequestRequestWithDefaults

`func NewApproveGroupParticipantRequestRequestWithDefaults() *ApproveGroupParticipantRequestRequest`

NewApproveGroupParticipantRequestRequestWithDefaults instantiates a new ApproveGroupParticipantRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ApproveGroupParticipantRequestRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApproveGroupParticipantRequestRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApproveGroupParticipantRequestRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetParticipants

`func (o *ApproveGroupParticipantRequestRequest) GetParticipants() []string`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *ApproveGroupParticipantRequestRequest) GetParticipantsOk() (*[]string, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *ApproveGroupParticipantRequestRequest) SetParticipants(v []string)`

SetParticipants sets Participants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



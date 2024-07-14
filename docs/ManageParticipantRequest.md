# ManageParticipantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** |  | [optional] 
**Participants** | Pointer to **[]string** |  | [optional] 

## Methods

### NewManageParticipantRequest

`func NewManageParticipantRequest() *ManageParticipantRequest`

NewManageParticipantRequest instantiates a new ManageParticipantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageParticipantRequestWithDefaults

`func NewManageParticipantRequestWithDefaults() *ManageParticipantRequest`

NewManageParticipantRequestWithDefaults instantiates a new ManageParticipantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ManageParticipantRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ManageParticipantRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ManageParticipantRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ManageParticipantRequest) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetParticipants

`func (o *ManageParticipantRequest) GetParticipants() []string`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *ManageParticipantRequest) GetParticipantsOk() (*[]string, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *ManageParticipantRequest) SetParticipants(v []string)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *ManageParticipantRequest) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



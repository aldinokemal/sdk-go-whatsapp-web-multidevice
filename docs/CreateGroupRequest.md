# CreateGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Participants** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateGroupRequest

`func NewCreateGroupRequest() *CreateGroupRequest`

NewCreateGroupRequest instantiates a new CreateGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupRequestWithDefaults

`func NewCreateGroupRequestWithDefaults() *CreateGroupRequest`

NewCreateGroupRequestWithDefaults instantiates a new CreateGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateGroupRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateGroupRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateGroupRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateGroupRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetParticipants

`func (o *CreateGroupRequest) GetParticipants() []string`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *CreateGroupRequest) GetParticipantsOk() (*[]string, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *CreateGroupRequest) SetParticipants(v []string)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *CreateGroupRequest) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



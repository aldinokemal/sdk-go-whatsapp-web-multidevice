# SetGroupLockedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The group ID | 
**Locked** | **bool** | Whether to lock the group (true) or unlock it (false) | 

## Methods

### NewSetGroupLockedRequest

`func NewSetGroupLockedRequest(groupId string, locked bool, ) *SetGroupLockedRequest`

NewSetGroupLockedRequest instantiates a new SetGroupLockedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetGroupLockedRequestWithDefaults

`func NewSetGroupLockedRequestWithDefaults() *SetGroupLockedRequest`

NewSetGroupLockedRequestWithDefaults instantiates a new SetGroupLockedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *SetGroupLockedRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SetGroupLockedRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SetGroupLockedRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetLocked

`func (o *SetGroupLockedRequest) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *SetGroupLockedRequest) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *SetGroupLockedRequest) SetLocked(v bool)`

SetLocked sets Locked field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



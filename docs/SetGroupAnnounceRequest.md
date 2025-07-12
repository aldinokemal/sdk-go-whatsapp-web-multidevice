# SetGroupAnnounceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The group ID | 
**Announce** | **bool** | Whether to enable announce mode (true) or disable it (false) | 

## Methods

### NewSetGroupAnnounceRequest

`func NewSetGroupAnnounceRequest(groupId string, announce bool, ) *SetGroupAnnounceRequest`

NewSetGroupAnnounceRequest instantiates a new SetGroupAnnounceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetGroupAnnounceRequestWithDefaults

`func NewSetGroupAnnounceRequestWithDefaults() *SetGroupAnnounceRequest`

NewSetGroupAnnounceRequestWithDefaults instantiates a new SetGroupAnnounceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *SetGroupAnnounceRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SetGroupAnnounceRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SetGroupAnnounceRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetAnnounce

`func (o *SetGroupAnnounceRequest) GetAnnounce() bool`

GetAnnounce returns the Announce field if non-nil, zero value otherwise.

### GetAnnounceOk

`func (o *SetGroupAnnounceRequest) GetAnnounceOk() (*bool, bool)`

GetAnnounceOk returns a tuple with the Announce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnounce

`func (o *SetGroupAnnounceRequest) SetAnnounce(v bool)`

SetAnnounce sets Announce field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



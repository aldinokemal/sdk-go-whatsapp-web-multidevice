# SetGroupNameRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The group ID | 
**Name** | **string** | The new group name (max 25 characters) | 

## Methods

### NewSetGroupNameRequest

`func NewSetGroupNameRequest(groupId string, name string, ) *SetGroupNameRequest`

NewSetGroupNameRequest instantiates a new SetGroupNameRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetGroupNameRequestWithDefaults

`func NewSetGroupNameRequestWithDefaults() *SetGroupNameRequest`

NewSetGroupNameRequestWithDefaults instantiates a new SetGroupNameRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *SetGroupNameRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SetGroupNameRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SetGroupNameRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetName

`func (o *SetGroupNameRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SetGroupNameRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SetGroupNameRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ChatwootSyncHistoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | Device ID to sync (uses CHATWOOT_DEVICE_ID if not specified) | [optional] 
**DaysLimit** | Pointer to **int32** | Number of days of history to sync | [optional] [default to 3]
**IncludeMedia** | Pointer to **bool** | Include media attachments in sync | [optional] [default to true]
**IncludeGroups** | Pointer to **bool** | Include group chats in sync | [optional] [default to true]

## Methods

### NewChatwootSyncHistoryRequest

`func NewChatwootSyncHistoryRequest() *ChatwootSyncHistoryRequest`

NewChatwootSyncHistoryRequest instantiates a new ChatwootSyncHistoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatwootSyncHistoryRequestWithDefaults

`func NewChatwootSyncHistoryRequestWithDefaults() *ChatwootSyncHistoryRequest`

NewChatwootSyncHistoryRequestWithDefaults instantiates a new ChatwootSyncHistoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *ChatwootSyncHistoryRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ChatwootSyncHistoryRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ChatwootSyncHistoryRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ChatwootSyncHistoryRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDaysLimit

`func (o *ChatwootSyncHistoryRequest) GetDaysLimit() int32`

GetDaysLimit returns the DaysLimit field if non-nil, zero value otherwise.

### GetDaysLimitOk

`func (o *ChatwootSyncHistoryRequest) GetDaysLimitOk() (*int32, bool)`

GetDaysLimitOk returns a tuple with the DaysLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysLimit

`func (o *ChatwootSyncHistoryRequest) SetDaysLimit(v int32)`

SetDaysLimit sets DaysLimit field to given value.

### HasDaysLimit

`func (o *ChatwootSyncHistoryRequest) HasDaysLimit() bool`

HasDaysLimit returns a boolean if a field has been set.

### GetIncludeMedia

`func (o *ChatwootSyncHistoryRequest) GetIncludeMedia() bool`

GetIncludeMedia returns the IncludeMedia field if non-nil, zero value otherwise.

### GetIncludeMediaOk

`func (o *ChatwootSyncHistoryRequest) GetIncludeMediaOk() (*bool, bool)`

GetIncludeMediaOk returns a tuple with the IncludeMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMedia

`func (o *ChatwootSyncHistoryRequest) SetIncludeMedia(v bool)`

SetIncludeMedia sets IncludeMedia field to given value.

### HasIncludeMedia

`func (o *ChatwootSyncHistoryRequest) HasIncludeMedia() bool`

HasIncludeMedia returns a boolean if a field has been set.

### GetIncludeGroups

`func (o *ChatwootSyncHistoryRequest) GetIncludeGroups() bool`

GetIncludeGroups returns the IncludeGroups field if non-nil, zero value otherwise.

### GetIncludeGroupsOk

`func (o *ChatwootSyncHistoryRequest) GetIncludeGroupsOk() (*bool, bool)`

GetIncludeGroupsOk returns a tuple with the IncludeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeGroups

`func (o *ChatwootSyncHistoryRequest) SetIncludeGroups(v bool)`

SetIncludeGroups sets IncludeGroups field to given value.

### HasIncludeGroups

`func (o *ChatwootSyncHistoryRequest) HasIncludeGroups() bool`

HasIncludeGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



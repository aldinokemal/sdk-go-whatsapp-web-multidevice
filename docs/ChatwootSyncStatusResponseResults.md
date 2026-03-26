# ChatwootSyncStatusResponseResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TotalChats** | Pointer to **int32** |  | [optional] 
**SyncedChats** | Pointer to **int32** |  | [optional] 
**TotalMessages** | Pointer to **int32** |  | [optional] 
**SyncedMessages** | Pointer to **int32** |  | [optional] 
**FailedMessages** | Pointer to **int32** |  | [optional] 
**CurrentChat** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewChatwootSyncStatusResponseResults

`func NewChatwootSyncStatusResponseResults() *ChatwootSyncStatusResponseResults`

NewChatwootSyncStatusResponseResults instantiates a new ChatwootSyncStatusResponseResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatwootSyncStatusResponseResultsWithDefaults

`func NewChatwootSyncStatusResponseResultsWithDefaults() *ChatwootSyncStatusResponseResults`

NewChatwootSyncStatusResponseResultsWithDefaults instantiates a new ChatwootSyncStatusResponseResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *ChatwootSyncStatusResponseResults) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ChatwootSyncStatusResponseResults) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ChatwootSyncStatusResponseResults) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ChatwootSyncStatusResponseResults) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetStatus

`func (o *ChatwootSyncStatusResponseResults) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChatwootSyncStatusResponseResults) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChatwootSyncStatusResponseResults) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChatwootSyncStatusResponseResults) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalChats

`func (o *ChatwootSyncStatusResponseResults) GetTotalChats() int32`

GetTotalChats returns the TotalChats field if non-nil, zero value otherwise.

### GetTotalChatsOk

`func (o *ChatwootSyncStatusResponseResults) GetTotalChatsOk() (*int32, bool)`

GetTotalChatsOk returns a tuple with the TotalChats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalChats

`func (o *ChatwootSyncStatusResponseResults) SetTotalChats(v int32)`

SetTotalChats sets TotalChats field to given value.

### HasTotalChats

`func (o *ChatwootSyncStatusResponseResults) HasTotalChats() bool`

HasTotalChats returns a boolean if a field has been set.

### GetSyncedChats

`func (o *ChatwootSyncStatusResponseResults) GetSyncedChats() int32`

GetSyncedChats returns the SyncedChats field if non-nil, zero value otherwise.

### GetSyncedChatsOk

`func (o *ChatwootSyncStatusResponseResults) GetSyncedChatsOk() (*int32, bool)`

GetSyncedChatsOk returns a tuple with the SyncedChats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedChats

`func (o *ChatwootSyncStatusResponseResults) SetSyncedChats(v int32)`

SetSyncedChats sets SyncedChats field to given value.

### HasSyncedChats

`func (o *ChatwootSyncStatusResponseResults) HasSyncedChats() bool`

HasSyncedChats returns a boolean if a field has been set.

### GetTotalMessages

`func (o *ChatwootSyncStatusResponseResults) GetTotalMessages() int32`

GetTotalMessages returns the TotalMessages field if non-nil, zero value otherwise.

### GetTotalMessagesOk

`func (o *ChatwootSyncStatusResponseResults) GetTotalMessagesOk() (*int32, bool)`

GetTotalMessagesOk returns a tuple with the TotalMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMessages

`func (o *ChatwootSyncStatusResponseResults) SetTotalMessages(v int32)`

SetTotalMessages sets TotalMessages field to given value.

### HasTotalMessages

`func (o *ChatwootSyncStatusResponseResults) HasTotalMessages() bool`

HasTotalMessages returns a boolean if a field has been set.

### GetSyncedMessages

`func (o *ChatwootSyncStatusResponseResults) GetSyncedMessages() int32`

GetSyncedMessages returns the SyncedMessages field if non-nil, zero value otherwise.

### GetSyncedMessagesOk

`func (o *ChatwootSyncStatusResponseResults) GetSyncedMessagesOk() (*int32, bool)`

GetSyncedMessagesOk returns a tuple with the SyncedMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedMessages

`func (o *ChatwootSyncStatusResponseResults) SetSyncedMessages(v int32)`

SetSyncedMessages sets SyncedMessages field to given value.

### HasSyncedMessages

`func (o *ChatwootSyncStatusResponseResults) HasSyncedMessages() bool`

HasSyncedMessages returns a boolean if a field has been set.

### GetFailedMessages

`func (o *ChatwootSyncStatusResponseResults) GetFailedMessages() int32`

GetFailedMessages returns the FailedMessages field if non-nil, zero value otherwise.

### GetFailedMessagesOk

`func (o *ChatwootSyncStatusResponseResults) GetFailedMessagesOk() (*int32, bool)`

GetFailedMessagesOk returns a tuple with the FailedMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMessages

`func (o *ChatwootSyncStatusResponseResults) SetFailedMessages(v int32)`

SetFailedMessages sets FailedMessages field to given value.

### HasFailedMessages

`func (o *ChatwootSyncStatusResponseResults) HasFailedMessages() bool`

HasFailedMessages returns a boolean if a field has been set.

### GetCurrentChat

`func (o *ChatwootSyncStatusResponseResults) GetCurrentChat() string`

GetCurrentChat returns the CurrentChat field if non-nil, zero value otherwise.

### GetCurrentChatOk

`func (o *ChatwootSyncStatusResponseResults) GetCurrentChatOk() (*string, bool)`

GetCurrentChatOk returns a tuple with the CurrentChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChat

`func (o *ChatwootSyncStatusResponseResults) SetCurrentChat(v string)`

SetCurrentChat sets CurrentChat field to given value.

### HasCurrentChat

`func (o *ChatwootSyncStatusResponseResults) HasCurrentChat() bool`

HasCurrentChat returns a boolean if a field has been set.

### GetStartedAt

`func (o *ChatwootSyncStatusResponseResults) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ChatwootSyncStatusResponseResults) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ChatwootSyncStatusResponseResults) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ChatwootSyncStatusResponseResults) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *ChatwootSyncStatusResponseResults) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ChatwootSyncStatusResponseResults) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ChatwootSyncStatusResponseResults) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ChatwootSyncStatusResponseResults) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetError

`func (o *ChatwootSyncStatusResponseResults) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ChatwootSyncStatusResponseResults) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ChatwootSyncStatusResponseResults) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ChatwootSyncStatusResponseResults) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



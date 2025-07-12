# Chat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jid** | Pointer to **string** | Chat JID identifier | [optional] 
**Name** | Pointer to **string** | Chat display name | [optional] 
**LastMessageTime** | Pointer to **time.Time** | Timestamp of the last message | [optional] 
**EphemeralExpiration** | Pointer to **int32** | Ephemeral message expiration time in seconds (0 &#x3D; disabled) | [optional] 
**CreatedAt** | Pointer to **time.Time** | Chat creation timestamp | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Chat last update timestamp | [optional] 

## Methods

### NewChat

`func NewChat() *Chat`

NewChat instantiates a new Chat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatWithDefaults

`func NewChatWithDefaults() *Chat`

NewChatWithDefaults instantiates a new Chat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJid

`func (o *Chat) GetJid() string`

GetJid returns the Jid field if non-nil, zero value otherwise.

### GetJidOk

`func (o *Chat) GetJidOk() (*string, bool)`

GetJidOk returns a tuple with the Jid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJid

`func (o *Chat) SetJid(v string)`

SetJid sets Jid field to given value.

### HasJid

`func (o *Chat) HasJid() bool`

HasJid returns a boolean if a field has been set.

### GetName

`func (o *Chat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Chat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Chat) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Chat) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLastMessageTime

`func (o *Chat) GetLastMessageTime() time.Time`

GetLastMessageTime returns the LastMessageTime field if non-nil, zero value otherwise.

### GetLastMessageTimeOk

`func (o *Chat) GetLastMessageTimeOk() (*time.Time, bool)`

GetLastMessageTimeOk returns a tuple with the LastMessageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageTime

`func (o *Chat) SetLastMessageTime(v time.Time)`

SetLastMessageTime sets LastMessageTime field to given value.

### HasLastMessageTime

`func (o *Chat) HasLastMessageTime() bool`

HasLastMessageTime returns a boolean if a field has been set.

### GetEphemeralExpiration

`func (o *Chat) GetEphemeralExpiration() int32`

GetEphemeralExpiration returns the EphemeralExpiration field if non-nil, zero value otherwise.

### GetEphemeralExpirationOk

`func (o *Chat) GetEphemeralExpirationOk() (*int32, bool)`

GetEphemeralExpirationOk returns a tuple with the EphemeralExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralExpiration

`func (o *Chat) SetEphemeralExpiration(v int32)`

SetEphemeralExpiration sets EphemeralExpiration field to given value.

### HasEphemeralExpiration

`func (o *Chat) HasEphemeralExpiration() bool`

HasEphemeralExpiration returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Chat) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Chat) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Chat) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Chat) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Chat) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Chat) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Chat) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Chat) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



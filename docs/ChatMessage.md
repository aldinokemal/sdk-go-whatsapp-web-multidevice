# ChatMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Message ID | [optional] 
**ChatJid** | Pointer to **string** | Chat JID this message belongs to | [optional] 
**SenderJid** | Pointer to **string** | Sender JID | [optional] 
**Content** | Pointer to **string** | Message text content | [optional] 
**Timestamp** | Pointer to **time.Time** | Message timestamp | [optional] 
**IsFromMe** | Pointer to **bool** | Whether this message was sent by the current user | [optional] 
**MediaType** | Pointer to **NullableString** | Type of media (image, video, audio, document, etc.) | [optional] 
**Filename** | Pointer to **NullableString** | Original filename for media messages | [optional] 
**Url** | Pointer to **NullableString** | Media file URL | [optional] 
**FileLength** | Pointer to **NullableInt32** | File size in bytes for media messages | [optional] 
**CreatedAt** | Pointer to **time.Time** | Record creation timestamp | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Record last update timestamp | [optional] 

## Methods

### NewChatMessage

`func NewChatMessage() *ChatMessage`

NewChatMessage instantiates a new ChatMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMessageWithDefaults

`func NewChatMessageWithDefaults() *ChatMessage`

NewChatMessageWithDefaults instantiates a new ChatMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChatMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChatMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChatJid

`func (o *ChatMessage) GetChatJid() string`

GetChatJid returns the ChatJid field if non-nil, zero value otherwise.

### GetChatJidOk

`func (o *ChatMessage) GetChatJidOk() (*string, bool)`

GetChatJidOk returns a tuple with the ChatJid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatJid

`func (o *ChatMessage) SetChatJid(v string)`

SetChatJid sets ChatJid field to given value.

### HasChatJid

`func (o *ChatMessage) HasChatJid() bool`

HasChatJid returns a boolean if a field has been set.

### GetSenderJid

`func (o *ChatMessage) GetSenderJid() string`

GetSenderJid returns the SenderJid field if non-nil, zero value otherwise.

### GetSenderJidOk

`func (o *ChatMessage) GetSenderJidOk() (*string, bool)`

GetSenderJidOk returns a tuple with the SenderJid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderJid

`func (o *ChatMessage) SetSenderJid(v string)`

SetSenderJid sets SenderJid field to given value.

### HasSenderJid

`func (o *ChatMessage) HasSenderJid() bool`

HasSenderJid returns a boolean if a field has been set.

### GetContent

`func (o *ChatMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ChatMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ChatMessage) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ChatMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetTimestamp

`func (o *ChatMessage) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ChatMessage) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ChatMessage) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ChatMessage) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetIsFromMe

`func (o *ChatMessage) GetIsFromMe() bool`

GetIsFromMe returns the IsFromMe field if non-nil, zero value otherwise.

### GetIsFromMeOk

`func (o *ChatMessage) GetIsFromMeOk() (*bool, bool)`

GetIsFromMeOk returns a tuple with the IsFromMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFromMe

`func (o *ChatMessage) SetIsFromMe(v bool)`

SetIsFromMe sets IsFromMe field to given value.

### HasIsFromMe

`func (o *ChatMessage) HasIsFromMe() bool`

HasIsFromMe returns a boolean if a field has been set.

### GetMediaType

`func (o *ChatMessage) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *ChatMessage) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *ChatMessage) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *ChatMessage) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### SetMediaTypeNil

`func (o *ChatMessage) SetMediaTypeNil(b bool)`

 SetMediaTypeNil sets the value for MediaType to be an explicit nil

### UnsetMediaType
`func (o *ChatMessage) UnsetMediaType()`

UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil
### GetFilename

`func (o *ChatMessage) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ChatMessage) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ChatMessage) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ChatMessage) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### SetFilenameNil

`func (o *ChatMessage) SetFilenameNil(b bool)`

 SetFilenameNil sets the value for Filename to be an explicit nil

### UnsetFilename
`func (o *ChatMessage) UnsetFilename()`

UnsetFilename ensures that no value is present for Filename, not even an explicit nil
### GetUrl

`func (o *ChatMessage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChatMessage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChatMessage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ChatMessage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ChatMessage) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ChatMessage) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetFileLength

`func (o *ChatMessage) GetFileLength() int32`

GetFileLength returns the FileLength field if non-nil, zero value otherwise.

### GetFileLengthOk

`func (o *ChatMessage) GetFileLengthOk() (*int32, bool)`

GetFileLengthOk returns a tuple with the FileLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLength

`func (o *ChatMessage) SetFileLength(v int32)`

SetFileLength sets FileLength field to given value.

### HasFileLength

`func (o *ChatMessage) HasFileLength() bool`

HasFileLength returns a boolean if a field has been set.

### SetFileLengthNil

`func (o *ChatMessage) SetFileLengthNil(b bool)`

 SetFileLengthNil sets the value for FileLength to be an explicit nil

### UnsetFileLength
`func (o *ChatMessage) UnsetFileLength()`

UnsetFileLength ensures that no value is present for FileLength, not even an explicit nil
### GetCreatedAt

`func (o *ChatMessage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChatMessage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChatMessage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChatMessage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ChatMessage) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChatMessage) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChatMessage) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ChatMessage) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



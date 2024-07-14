# ReactMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** | Phone number with country code | [optional] 
**Emoji** | Pointer to **string** | Emoji to react | [optional] 

## Methods

### NewReactMessageRequest

`func NewReactMessageRequest() *ReactMessageRequest`

NewReactMessageRequest instantiates a new ReactMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactMessageRequestWithDefaults

`func NewReactMessageRequestWithDefaults() *ReactMessageRequest`

NewReactMessageRequestWithDefaults instantiates a new ReactMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *ReactMessageRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ReactMessageRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ReactMessageRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ReactMessageRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmoji

`func (o *ReactMessageRequest) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *ReactMessageRequest) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *ReactMessageRequest) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *ReactMessageRequest) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



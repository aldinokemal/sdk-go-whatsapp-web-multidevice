# SendMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** | Phone number with country code | [optional] 
**Message** | Pointer to **string** | Message to send | [optional] 
**ReplyMessageId** | Pointer to **string** | Message ID that you want reply | [optional] 
**IsForwarded** | Pointer to **bool** | Whether this is a forwarded message | [optional] 
**Duration** | Pointer to **int32** | Disappearing message duration in seconds (optional) | [optional] 

## Methods

### NewSendMessageRequest

`func NewSendMessageRequest() *SendMessageRequest`

NewSendMessageRequest instantiates a new SendMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessageRequestWithDefaults

`func NewSendMessageRequestWithDefaults() *SendMessageRequest`

NewSendMessageRequestWithDefaults instantiates a new SendMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendMessageRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendMessageRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendMessageRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendMessageRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetMessage

`func (o *SendMessageRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SendMessageRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SendMessageRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SendMessageRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReplyMessageId

`func (o *SendMessageRequest) GetReplyMessageId() string`

GetReplyMessageId returns the ReplyMessageId field if non-nil, zero value otherwise.

### GetReplyMessageIdOk

`func (o *SendMessageRequest) GetReplyMessageIdOk() (*string, bool)`

GetReplyMessageIdOk returns a tuple with the ReplyMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMessageId

`func (o *SendMessageRequest) SetReplyMessageId(v string)`

SetReplyMessageId sets ReplyMessageId field to given value.

### HasReplyMessageId

`func (o *SendMessageRequest) HasReplyMessageId() bool`

HasReplyMessageId returns a boolean if a field has been set.

### GetIsForwarded

`func (o *SendMessageRequest) GetIsForwarded() bool`

GetIsForwarded returns the IsForwarded field if non-nil, zero value otherwise.

### GetIsForwardedOk

`func (o *SendMessageRequest) GetIsForwardedOk() (*bool, bool)`

GetIsForwardedOk returns a tuple with the IsForwarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForwarded

`func (o *SendMessageRequest) SetIsForwarded(v bool)`

SetIsForwarded sets IsForwarded field to given value.

### HasIsForwarded

`func (o *SendMessageRequest) HasIsForwarded() bool`

HasIsForwarded returns a boolean if a field has been set.

### GetDuration

`func (o *SendMessageRequest) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SendMessageRequest) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SendMessageRequest) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SendMessageRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



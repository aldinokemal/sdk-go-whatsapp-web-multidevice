# SendChatPresenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | **string** | Phone number with country code | 
**Action** | **string** | Action to perform - \&quot;start\&quot; to begin typing indicator, \&quot;stop\&quot; to end typing indicator | 

## Methods

### NewSendChatPresenceRequest

`func NewSendChatPresenceRequest(phone string, action string, ) *SendChatPresenceRequest`

NewSendChatPresenceRequest instantiates a new SendChatPresenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendChatPresenceRequestWithDefaults

`func NewSendChatPresenceRequestWithDefaults() *SendChatPresenceRequest`

NewSendChatPresenceRequestWithDefaults instantiates a new SendChatPresenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendChatPresenceRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendChatPresenceRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendChatPresenceRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetAction

`func (o *SendChatPresenceRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SendChatPresenceRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SendChatPresenceRequest) SetAction(v string)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



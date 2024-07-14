# SendLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** | Phone number with country code | [optional] 
**Link** | Pointer to **string** | Link to send | [optional] 
**Caption** | Pointer to **string** | Caption to send | [optional] 

## Methods

### NewSendLinkRequest

`func NewSendLinkRequest() *SendLinkRequest`

NewSendLinkRequest instantiates a new SendLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendLinkRequestWithDefaults

`func NewSendLinkRequestWithDefaults() *SendLinkRequest`

NewSendLinkRequestWithDefaults instantiates a new SendLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendLinkRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendLinkRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendLinkRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendLinkRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetLink

`func (o *SendLinkRequest) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *SendLinkRequest) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *SendLinkRequest) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *SendLinkRequest) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetCaption

`func (o *SendLinkRequest) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *SendLinkRequest) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *SendLinkRequest) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *SendLinkRequest) HasCaption() bool`

HasCaption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



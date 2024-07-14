# RevokeMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** | Phone number with country code | [optional] 

## Methods

### NewRevokeMessageRequest

`func NewRevokeMessageRequest() *RevokeMessageRequest`

NewRevokeMessageRequest instantiates a new RevokeMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevokeMessageRequestWithDefaults

`func NewRevokeMessageRequestWithDefaults() *RevokeMessageRequest`

NewRevokeMessageRequestWithDefaults instantiates a new RevokeMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *RevokeMessageRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *RevokeMessageRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *RevokeMessageRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *RevokeMessageRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SendContactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** | Phone number with country code | [optional] 
**ContactName** | Pointer to **string** | Contact name | [optional] 
**ContactPhone** | Pointer to **string** | Contact phone number | [optional] 

## Methods

### NewSendContactRequest

`func NewSendContactRequest() *SendContactRequest`

NewSendContactRequest instantiates a new SendContactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendContactRequestWithDefaults

`func NewSendContactRequestWithDefaults() *SendContactRequest`

NewSendContactRequestWithDefaults instantiates a new SendContactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendContactRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendContactRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendContactRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendContactRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetContactName

`func (o *SendContactRequest) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *SendContactRequest) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *SendContactRequest) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *SendContactRequest) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetContactPhone

`func (o *SendContactRequest) GetContactPhone() string`

GetContactPhone returns the ContactPhone field if non-nil, zero value otherwise.

### GetContactPhoneOk

`func (o *SendContactRequest) GetContactPhoneOk() (*string, bool)`

GetContactPhoneOk returns a tuple with the ContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhone

`func (o *SendContactRequest) SetContactPhone(v string)`

SetContactPhone sets ContactPhone field to given value.

### HasContactPhone

`func (o *SendContactRequest) HasContactPhone() bool`

HasContactPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SendPresenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The presence type to send | 
**IsForwarded** | Pointer to **bool** | Whether this is a forwarded message | [optional] 

## Methods

### NewSendPresenceRequest

`func NewSendPresenceRequest(type_ string, ) *SendPresenceRequest`

NewSendPresenceRequest instantiates a new SendPresenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendPresenceRequestWithDefaults

`func NewSendPresenceRequestWithDefaults() *SendPresenceRequest`

NewSendPresenceRequestWithDefaults instantiates a new SendPresenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SendPresenceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SendPresenceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SendPresenceRequest) SetType(v string)`

SetType sets Type field to given value.


### GetIsForwarded

`func (o *SendPresenceRequest) GetIsForwarded() bool`

GetIsForwarded returns the IsForwarded field if non-nil, zero value otherwise.

### GetIsForwardedOk

`func (o *SendPresenceRequest) GetIsForwardedOk() (*bool, bool)`

GetIsForwardedOk returns a tuple with the IsForwarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForwarded

`func (o *SendPresenceRequest) SetIsForwarded(v bool)`

SetIsForwarded sets IsForwarded field to given value.

### HasIsForwarded

`func (o *SendPresenceRequest) HasIsForwarded() bool`

HasIsForwarded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SendPresenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Presence** | **string** | The presence status to send | 

## Methods

### NewSendPresenceRequest

`func NewSendPresenceRequest(presence string, ) *SendPresenceRequest`

NewSendPresenceRequest instantiates a new SendPresenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendPresenceRequestWithDefaults

`func NewSendPresenceRequestWithDefaults() *SendPresenceRequest`

NewSendPresenceRequestWithDefaults instantiates a new SendPresenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPresence

`func (o *SendPresenceRequest) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *SendPresenceRequest) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *SendPresenceRequest) SetPresence(v string)`

SetPresence sets Presence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



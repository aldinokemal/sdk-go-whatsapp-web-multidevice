# SetDisappearingTimerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimerSeconds** | **int32** | Timer in seconds: 0 (off), 86400 (24h), 604800 (7d), 7776000 (90d) | 

## Methods

### NewSetDisappearingTimerRequest

`func NewSetDisappearingTimerRequest(timerSeconds int32, ) *SetDisappearingTimerRequest`

NewSetDisappearingTimerRequest instantiates a new SetDisappearingTimerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetDisappearingTimerRequestWithDefaults

`func NewSetDisappearingTimerRequestWithDefaults() *SetDisappearingTimerRequest`

NewSetDisappearingTimerRequestWithDefaults instantiates a new SetDisappearingTimerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimerSeconds

`func (o *SetDisappearingTimerRequest) GetTimerSeconds() int32`

GetTimerSeconds returns the TimerSeconds field if non-nil, zero value otherwise.

### GetTimerSecondsOk

`func (o *SetDisappearingTimerRequest) GetTimerSecondsOk() (*int32, bool)`

GetTimerSecondsOk returns a tuple with the TimerSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerSeconds

`func (o *SetDisappearingTimerRequest) SetTimerSeconds(v int32)`

SetTimerSeconds sets TimerSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SendLocationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** | Phone number with country code | [optional] 
**Latitude** | Pointer to **string** | Latitude coordinate | [optional] 
**Longitude** | Pointer to **string** | Longitude coordinate | [optional] 

## Methods

### NewSendLocationRequest

`func NewSendLocationRequest() *SendLocationRequest`

NewSendLocationRequest instantiates a new SendLocationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendLocationRequestWithDefaults

`func NewSendLocationRequestWithDefaults() *SendLocationRequest`

NewSendLocationRequestWithDefaults instantiates a new SendLocationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendLocationRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendLocationRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendLocationRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendLocationRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetLatitude

`func (o *SendLocationRequest) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *SendLocationRequest) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *SendLocationRequest) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *SendLocationRequest) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *SendLocationRequest) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *SendLocationRequest) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *SendLocationRequest) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *SendLocationRequest) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



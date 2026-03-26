# AppStatus200ResponseResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsConnected** | Pointer to **bool** | Whether the device is connected to WhatsApp servers | [optional] 
**IsLoggedIn** | Pointer to **bool** | Whether the device is logged in to WhatsApp | [optional] 
**DeviceId** | Pointer to **string** | The device ID | [optional] 

## Methods

### NewAppStatus200ResponseResults

`func NewAppStatus200ResponseResults() *AppStatus200ResponseResults`

NewAppStatus200ResponseResults instantiates a new AppStatus200ResponseResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStatus200ResponseResultsWithDefaults

`func NewAppStatus200ResponseResultsWithDefaults() *AppStatus200ResponseResults`

NewAppStatus200ResponseResultsWithDefaults instantiates a new AppStatus200ResponseResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsConnected

`func (o *AppStatus200ResponseResults) GetIsConnected() bool`

GetIsConnected returns the IsConnected field if non-nil, zero value otherwise.

### GetIsConnectedOk

`func (o *AppStatus200ResponseResults) GetIsConnectedOk() (*bool, bool)`

GetIsConnectedOk returns a tuple with the IsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConnected

`func (o *AppStatus200ResponseResults) SetIsConnected(v bool)`

SetIsConnected sets IsConnected field to given value.

### HasIsConnected

`func (o *AppStatus200ResponseResults) HasIsConnected() bool`

HasIsConnected returns a boolean if a field has been set.

### GetIsLoggedIn

`func (o *AppStatus200ResponseResults) GetIsLoggedIn() bool`

GetIsLoggedIn returns the IsLoggedIn field if non-nil, zero value otherwise.

### GetIsLoggedInOk

`func (o *AppStatus200ResponseResults) GetIsLoggedInOk() (*bool, bool)`

GetIsLoggedInOk returns a tuple with the IsLoggedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLoggedIn

`func (o *AppStatus200ResponseResults) SetIsLoggedIn(v bool)`

SetIsLoggedIn sets IsLoggedIn field to given value.

### HasIsLoggedIn

`func (o *AppStatus200ResponseResults) HasIsLoggedIn() bool`

HasIsLoggedIn returns a boolean if a field has been set.

### GetDeviceId

`func (o *AppStatus200ResponseResults) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *AppStatus200ResponseResults) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *AppStatus200ResponseResults) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *AppStatus200ResponseResults) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



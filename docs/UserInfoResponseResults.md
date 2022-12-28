# UserInfoResponseResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerifiedName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**PictureId** | Pointer to **string** |  | [optional] 
**Devices** | Pointer to [**[]UserInfoResponseResultsDevicesInner**](UserInfoResponseResultsDevicesInner.md) |  | [optional] 

## Methods

### NewUserInfoResponseResults

`func NewUserInfoResponseResults() *UserInfoResponseResults`

NewUserInfoResponseResults instantiates a new UserInfoResponseResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInfoResponseResultsWithDefaults

`func NewUserInfoResponseResultsWithDefaults() *UserInfoResponseResults`

NewUserInfoResponseResultsWithDefaults instantiates a new UserInfoResponseResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerifiedName

`func (o *UserInfoResponseResults) GetVerifiedName() string`

GetVerifiedName returns the VerifiedName field if non-nil, zero value otherwise.

### GetVerifiedNameOk

`func (o *UserInfoResponseResults) GetVerifiedNameOk() (*string, bool)`

GetVerifiedNameOk returns a tuple with the VerifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedName

`func (o *UserInfoResponseResults) SetVerifiedName(v string)`

SetVerifiedName sets VerifiedName field to given value.

### HasVerifiedName

`func (o *UserInfoResponseResults) HasVerifiedName() bool`

HasVerifiedName returns a boolean if a field has been set.

### GetStatus

`func (o *UserInfoResponseResults) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserInfoResponseResults) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserInfoResponseResults) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserInfoResponseResults) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPictureId

`func (o *UserInfoResponseResults) GetPictureId() string`

GetPictureId returns the PictureId field if non-nil, zero value otherwise.

### GetPictureIdOk

`func (o *UserInfoResponseResults) GetPictureIdOk() (*string, bool)`

GetPictureIdOk returns a tuple with the PictureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureId

`func (o *UserInfoResponseResults) SetPictureId(v string)`

SetPictureId sets PictureId field to given value.

### HasPictureId

`func (o *UserInfoResponseResults) HasPictureId() bool`

HasPictureId returns a boolean if a field has been set.

### GetDevices

`func (o *UserInfoResponseResults) GetDevices() []UserInfoResponseResultsDevicesInner`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *UserInfoResponseResults) GetDevicesOk() (*[]UserInfoResponseResultsDevicesInner, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *UserInfoResponseResults) SetDevices(v []UserInfoResponseResultsDevicesInner)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *UserInfoResponseResults) HasDevices() bool`

HasDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



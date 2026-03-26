# DeviceInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Results** | Pointer to [**DeviceInfo**](DeviceInfo.md) |  | [optional] 

## Methods

### NewDeviceInfoResponse

`func NewDeviceInfoResponse() *DeviceInfoResponse`

NewDeviceInfoResponse instantiates a new DeviceInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInfoResponseWithDefaults

`func NewDeviceInfoResponseWithDefaults() *DeviceInfoResponse`

NewDeviceInfoResponseWithDefaults instantiates a new DeviceInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DeviceInfoResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DeviceInfoResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DeviceInfoResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *DeviceInfoResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *DeviceInfoResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeviceInfoResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeviceInfoResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeviceInfoResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceInfoResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceInfoResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceInfoResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceInfoResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResults

`func (o *DeviceInfoResponse) GetResults() DeviceInfo`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DeviceInfoResponse) GetResultsOk() (*DeviceInfo, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DeviceInfoResponse) SetResults(v DeviceInfo)`

SetResults sets Results field to given value.

### HasResults

`func (o *DeviceInfoResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



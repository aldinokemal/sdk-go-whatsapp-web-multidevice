# DeviceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Results** | Pointer to [**[]DeviceInfo**](DeviceInfo.md) |  | [optional] 

## Methods

### NewDeviceListResponse

`func NewDeviceListResponse() *DeviceListResponse`

NewDeviceListResponse instantiates a new DeviceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceListResponseWithDefaults

`func NewDeviceListResponseWithDefaults() *DeviceListResponse`

NewDeviceListResponseWithDefaults instantiates a new DeviceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DeviceListResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DeviceListResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DeviceListResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *DeviceListResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *DeviceListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeviceListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeviceListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeviceListResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceListResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceListResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceListResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResults

`func (o *DeviceListResponse) GetResults() []DeviceInfo`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DeviceListResponse) GetResultsOk() (*[]DeviceInfo, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DeviceListResponse) SetResults(v []DeviceInfo)`

SetResults sets Results field to given value.

### HasResults

`func (o *DeviceListResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



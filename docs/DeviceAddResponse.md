# DeviceAddResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Results** | Pointer to [**DeviceInfo**](DeviceInfo.md) |  | [optional] 

## Methods

### NewDeviceAddResponse

`func NewDeviceAddResponse() *DeviceAddResponse`

NewDeviceAddResponse instantiates a new DeviceAddResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAddResponseWithDefaults

`func NewDeviceAddResponseWithDefaults() *DeviceAddResponse`

NewDeviceAddResponseWithDefaults instantiates a new DeviceAddResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DeviceAddResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DeviceAddResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DeviceAddResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *DeviceAddResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *DeviceAddResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeviceAddResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeviceAddResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeviceAddResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceAddResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceAddResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceAddResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceAddResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResults

`func (o *DeviceAddResponse) GetResults() DeviceInfo`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DeviceAddResponse) GetResultsOk() (*DeviceInfo, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DeviceAddResponse) SetResults(v DeviceInfo)`

SetResults sets Results field to given value.

### HasResults

`func (o *DeviceAddResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



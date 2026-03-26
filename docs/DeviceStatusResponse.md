# DeviceStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Results** | Pointer to [**DeviceStatusResponseResults**](DeviceStatusResponseResults.md) |  | [optional] 

## Methods

### NewDeviceStatusResponse

`func NewDeviceStatusResponse() *DeviceStatusResponse`

NewDeviceStatusResponse instantiates a new DeviceStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceStatusResponseWithDefaults

`func NewDeviceStatusResponseWithDefaults() *DeviceStatusResponse`

NewDeviceStatusResponseWithDefaults instantiates a new DeviceStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DeviceStatusResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DeviceStatusResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DeviceStatusResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *DeviceStatusResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *DeviceStatusResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeviceStatusResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeviceStatusResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeviceStatusResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceStatusResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceStatusResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceStatusResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResults

`func (o *DeviceStatusResponse) GetResults() DeviceStatusResponseResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DeviceStatusResponse) GetResultsOk() (*DeviceStatusResponseResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DeviceStatusResponse) SetResults(v DeviceStatusResponseResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *DeviceStatusResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



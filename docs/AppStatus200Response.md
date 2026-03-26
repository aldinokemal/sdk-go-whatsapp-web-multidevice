# AppStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**AppStatus200ResponseResults**](AppStatus200ResponseResults.md) |  | [optional] 

## Methods

### NewAppStatus200Response

`func NewAppStatus200Response() *AppStatus200Response`

NewAppStatus200Response instantiates a new AppStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStatus200ResponseWithDefaults

`func NewAppStatus200ResponseWithDefaults() *AppStatus200Response`

NewAppStatus200ResponseWithDefaults instantiates a new AppStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AppStatus200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppStatus200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppStatus200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppStatus200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *AppStatus200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AppStatus200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AppStatus200Response) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AppStatus200Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *AppStatus200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AppStatus200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AppStatus200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AppStatus200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *AppStatus200Response) GetResults() AppStatus200ResponseResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AppStatus200Response) GetResultsOk() (*AppStatus200ResponseResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AppStatus200Response) SetResults(v AppStatus200ResponseResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *AppStatus200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



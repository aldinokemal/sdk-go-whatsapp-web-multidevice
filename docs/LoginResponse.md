# LoginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**LoginResponseResults**](LoginResponseResults.md) |  | [optional] 

## Methods

### NewLoginResponse

`func NewLoginResponse() *LoginResponse`

NewLoginResponse instantiates a new LoginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginResponseWithDefaults

`func NewLoginResponseWithDefaults() *LoginResponse`

NewLoginResponseWithDefaults instantiates a new LoginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *LoginResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LoginResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *LoginResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *LoginResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *LoginResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LoginResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LoginResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LoginResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *LoginResponse) GetResults() LoginResponseResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *LoginResponse) GetResultsOk() (*LoginResponseResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *LoginResponse) SetResults(v LoginResponseResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *LoginResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



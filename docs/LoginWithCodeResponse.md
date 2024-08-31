# LoginWithCodeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**LoginWithCodeResponseResults**](LoginWithCodeResponseResults.md) |  | [optional] 

## Methods

### NewLoginWithCodeResponse

`func NewLoginWithCodeResponse() *LoginWithCodeResponse`

NewLoginWithCodeResponse instantiates a new LoginWithCodeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginWithCodeResponseWithDefaults

`func NewLoginWithCodeResponseWithDefaults() *LoginWithCodeResponse`

NewLoginWithCodeResponseWithDefaults instantiates a new LoginWithCodeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *LoginWithCodeResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LoginWithCodeResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *LoginWithCodeResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *LoginWithCodeResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *LoginWithCodeResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LoginWithCodeResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LoginWithCodeResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LoginWithCodeResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *LoginWithCodeResponse) GetResults() LoginWithCodeResponseResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *LoginWithCodeResponse) GetResultsOk() (*LoginWithCodeResponseResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *LoginWithCodeResponse) SetResults(v LoginWithCodeResponseResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *LoginWithCodeResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



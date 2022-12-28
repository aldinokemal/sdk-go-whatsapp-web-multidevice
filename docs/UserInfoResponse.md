# UserInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**UserInfoResponseResults**](UserInfoResponseResults.md) |  | [optional] 

## Methods

### NewUserInfoResponse

`func NewUserInfoResponse() *UserInfoResponse`

NewUserInfoResponse instantiates a new UserInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInfoResponseWithDefaults

`func NewUserInfoResponseWithDefaults() *UserInfoResponse`

NewUserInfoResponseWithDefaults instantiates a new UserInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UserInfoResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UserInfoResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UserInfoResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UserInfoResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *UserInfoResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserInfoResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserInfoResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UserInfoResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *UserInfoResponse) GetResults() UserInfoResponseResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *UserInfoResponse) GetResultsOk() (*UserInfoResponseResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *UserInfoResponse) SetResults(v UserInfoResponseResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *UserInfoResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



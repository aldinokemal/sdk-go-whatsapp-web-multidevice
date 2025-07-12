# UserCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**UserCheckResponseResults**](UserCheckResponseResults.md) |  | [optional] 

## Methods

### NewUserCheckResponse

`func NewUserCheckResponse() *UserCheckResponse`

NewUserCheckResponse instantiates a new UserCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCheckResponseWithDefaults

`func NewUserCheckResponseWithDefaults() *UserCheckResponse`

NewUserCheckResponseWithDefaults instantiates a new UserCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UserCheckResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UserCheckResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UserCheckResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UserCheckResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *UserCheckResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserCheckResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserCheckResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UserCheckResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *UserCheckResponse) GetResults() UserCheckResponseResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *UserCheckResponse) GetResultsOk() (*UserCheckResponseResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *UserCheckResponse) SetResults(v UserCheckResponseResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *UserCheckResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



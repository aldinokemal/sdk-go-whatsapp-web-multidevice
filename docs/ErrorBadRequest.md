# ErrorBadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | HTTP Status Code | [optional] 
**Message** | Pointer to **string** | Detail error message | [optional] 
**Results** | Pointer to **map[string]interface{}** | additional data | [optional] 

## Methods

### NewErrorBadRequest

`func NewErrorBadRequest() *ErrorBadRequest`

NewErrorBadRequest instantiates a new ErrorBadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorBadRequestWithDefaults

`func NewErrorBadRequestWithDefaults() *ErrorBadRequest`

NewErrorBadRequestWithDefaults instantiates a new ErrorBadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorBadRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorBadRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorBadRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorBadRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorBadRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorBadRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorBadRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorBadRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *ErrorBadRequest) GetResults() map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ErrorBadRequest) GetResultsOk() (*map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ErrorBadRequest) SetResults(v map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *ErrorBadRequest) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



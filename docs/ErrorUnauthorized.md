# ErrorUnauthorized

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | HTTP Status Code | [optional] 
**Message** | Pointer to **string** | Detail error message | [optional] 
**Results** | Pointer to **map[string]interface{}** | additional data | [optional] 

## Methods

### NewErrorUnauthorized

`func NewErrorUnauthorized() *ErrorUnauthorized`

NewErrorUnauthorized instantiates a new ErrorUnauthorized object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorUnauthorizedWithDefaults

`func NewErrorUnauthorizedWithDefaults() *ErrorUnauthorized`

NewErrorUnauthorizedWithDefaults instantiates a new ErrorUnauthorized object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorUnauthorized) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorUnauthorized) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorUnauthorized) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorUnauthorized) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorUnauthorized) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorUnauthorized) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorUnauthorized) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorUnauthorized) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *ErrorUnauthorized) GetResults() map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ErrorUnauthorized) GetResultsOk() (*map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ErrorUnauthorized) SetResults(v map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *ErrorUnauthorized) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



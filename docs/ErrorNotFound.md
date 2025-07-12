# ErrorNotFound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | HTTP Status Code | [optional] 
**Message** | Pointer to **string** | Detail error message | [optional] 
**Results** | Pointer to **map[string]interface{}** | additional data | [optional] 

## Methods

### NewErrorNotFound

`func NewErrorNotFound() *ErrorNotFound`

NewErrorNotFound instantiates a new ErrorNotFound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorNotFoundWithDefaults

`func NewErrorNotFoundWithDefaults() *ErrorNotFound`

NewErrorNotFoundWithDefaults instantiates a new ErrorNotFound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorNotFound) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorNotFound) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorNotFound) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorNotFound) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorNotFound) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorNotFound) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorNotFound) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorNotFound) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *ErrorNotFound) GetResults() map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ErrorNotFound) GetResultsOk() (*map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ErrorNotFound) SetResults(v map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *ErrorNotFound) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



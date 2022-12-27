# GenericResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Results** | Pointer to **string** |  | [optional] 

## Methods

### NewGenericResponse

`func NewGenericResponse() *GenericResponse`

NewGenericResponse instantiates a new GenericResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericResponseWithDefaults

`func NewGenericResponseWithDefaults() *GenericResponse`

NewGenericResponseWithDefaults instantiates a new GenericResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GenericResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GenericResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GenericResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GenericResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *GenericResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GenericResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GenericResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GenericResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *GenericResponse) GetResults() string`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GenericResponse) GetResultsOk() (*string, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GenericResponse) SetResults(v string)`

SetResults sets Results field to given value.

### HasResults

`func (o *GenericResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



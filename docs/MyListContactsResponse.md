# MyListContactsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**MyListContactsResponseResults**](MyListContactsResponseResults.md) |  | [optional] 

## Methods

### NewMyListContactsResponse

`func NewMyListContactsResponse() *MyListContactsResponse`

NewMyListContactsResponse instantiates a new MyListContactsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyListContactsResponseWithDefaults

`func NewMyListContactsResponseWithDefaults() *MyListContactsResponse`

NewMyListContactsResponseWithDefaults instantiates a new MyListContactsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *MyListContactsResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MyListContactsResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MyListContactsResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *MyListContactsResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *MyListContactsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MyListContactsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MyListContactsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MyListContactsResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *MyListContactsResponse) GetResults() MyListContactsResponseResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *MyListContactsResponse) GetResultsOk() (*MyListContactsResponseResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *MyListContactsResponse) SetResults(v MyListContactsResponseResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *MyListContactsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



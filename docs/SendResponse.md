# SendResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**SendResponseResults**](SendResponseResults.md) |  | [optional] 

## Methods

### NewSendResponse

`func NewSendResponse() *SendResponse`

NewSendResponse instantiates a new SendResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendResponseWithDefaults

`func NewSendResponseWithDefaults() *SendResponse`

NewSendResponseWithDefaults instantiates a new SendResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SendResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SendResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SendResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SendResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *SendResponse) GetResults() SendResponseResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SendResponse) GetResultsOk() (*SendResponseResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SendResponse) SetResults(v SendResponseResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *SendResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ChatMessagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**ChatMessagesResponseResults**](ChatMessagesResponseResults.md) |  | [optional] 

## Methods

### NewChatMessagesResponse

`func NewChatMessagesResponse() *ChatMessagesResponse`

NewChatMessagesResponse instantiates a new ChatMessagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMessagesResponseWithDefaults

`func NewChatMessagesResponseWithDefaults() *ChatMessagesResponse`

NewChatMessagesResponseWithDefaults instantiates a new ChatMessagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ChatMessagesResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ChatMessagesResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ChatMessagesResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ChatMessagesResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ChatMessagesResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ChatMessagesResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ChatMessagesResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ChatMessagesResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *ChatMessagesResponse) GetResults() ChatMessagesResponseResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ChatMessagesResponse) GetResultsOk() (*ChatMessagesResponseResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ChatMessagesResponse) SetResults(v ChatMessagesResponseResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *ChatMessagesResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



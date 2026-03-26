# ChatwootSyncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**ChatwootSyncResponseResults**](ChatwootSyncResponseResults.md) |  | [optional] 

## Methods

### NewChatwootSyncResponse

`func NewChatwootSyncResponse() *ChatwootSyncResponse`

NewChatwootSyncResponse instantiates a new ChatwootSyncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatwootSyncResponseWithDefaults

`func NewChatwootSyncResponseWithDefaults() *ChatwootSyncResponse`

NewChatwootSyncResponseWithDefaults instantiates a new ChatwootSyncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ChatwootSyncResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChatwootSyncResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChatwootSyncResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChatwootSyncResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *ChatwootSyncResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ChatwootSyncResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ChatwootSyncResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ChatwootSyncResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ChatwootSyncResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ChatwootSyncResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ChatwootSyncResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ChatwootSyncResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *ChatwootSyncResponse) GetResults() ChatwootSyncResponseResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ChatwootSyncResponse) GetResultsOk() (*ChatwootSyncResponseResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ChatwootSyncResponse) SetResults(v ChatwootSyncResponseResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *ChatwootSyncResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



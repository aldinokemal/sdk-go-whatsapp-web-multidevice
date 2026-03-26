# ArchiveChatResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**ArchiveChatResponseResults**](ArchiveChatResponseResults.md) |  | [optional] 

## Methods

### NewArchiveChatResponse

`func NewArchiveChatResponse() *ArchiveChatResponse`

NewArchiveChatResponse instantiates a new ArchiveChatResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveChatResponseWithDefaults

`func NewArchiveChatResponseWithDefaults() *ArchiveChatResponse`

NewArchiveChatResponseWithDefaults instantiates a new ArchiveChatResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ArchiveChatResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArchiveChatResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArchiveChatResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArchiveChatResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *ArchiveChatResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ArchiveChatResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ArchiveChatResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ArchiveChatResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ArchiveChatResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ArchiveChatResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ArchiveChatResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ArchiveChatResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *ArchiveChatResponse) GetResults() ArchiveChatResponseResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ArchiveChatResponse) GetResultsOk() (*ArchiveChatResponseResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ArchiveChatResponse) SetResults(v ArchiveChatResponseResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *ArchiveChatResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



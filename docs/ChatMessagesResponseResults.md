# ChatMessagesResponseResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ChatMessage**](ChatMessage.md) |  | [optional] 
**Pagination** | Pointer to [**ChatMessagesResponseResultsPagination**](ChatMessagesResponseResultsPagination.md) |  | [optional] 
**ChatInfo** | Pointer to [**Chat**](Chat.md) |  | [optional] 

## Methods

### NewChatMessagesResponseResults

`func NewChatMessagesResponseResults() *ChatMessagesResponseResults`

NewChatMessagesResponseResults instantiates a new ChatMessagesResponseResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMessagesResponseResultsWithDefaults

`func NewChatMessagesResponseResultsWithDefaults() *ChatMessagesResponseResults`

NewChatMessagesResponseResultsWithDefaults instantiates a new ChatMessagesResponseResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ChatMessagesResponseResults) GetData() []ChatMessage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ChatMessagesResponseResults) GetDataOk() (*[]ChatMessage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ChatMessagesResponseResults) SetData(v []ChatMessage)`

SetData sets Data field to given value.

### HasData

`func (o *ChatMessagesResponseResults) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ChatMessagesResponseResults) GetPagination() ChatMessagesResponseResultsPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ChatMessagesResponseResults) GetPaginationOk() (*ChatMessagesResponseResultsPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ChatMessagesResponseResults) SetPagination(v ChatMessagesResponseResultsPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ChatMessagesResponseResults) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetChatInfo

`func (o *ChatMessagesResponseResults) GetChatInfo() Chat`

GetChatInfo returns the ChatInfo field if non-nil, zero value otherwise.

### GetChatInfoOk

`func (o *ChatMessagesResponseResults) GetChatInfoOk() (*Chat, bool)`

GetChatInfoOk returns a tuple with the ChatInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatInfo

`func (o *ChatMessagesResponseResults) SetChatInfo(v Chat)`

SetChatInfo sets ChatInfo field to given value.

### HasChatInfo

`func (o *ChatMessagesResponseResults) HasChatInfo() bool`

HasChatInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



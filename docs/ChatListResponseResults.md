# ChatListResponseResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Chat**](Chat.md) |  | [optional] 
**Pagination** | Pointer to [**ChatListResponseResultsPagination**](ChatListResponseResultsPagination.md) |  | [optional] 

## Methods

### NewChatListResponseResults

`func NewChatListResponseResults() *ChatListResponseResults`

NewChatListResponseResults instantiates a new ChatListResponseResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatListResponseResultsWithDefaults

`func NewChatListResponseResultsWithDefaults() *ChatListResponseResults`

NewChatListResponseResultsWithDefaults instantiates a new ChatListResponseResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ChatListResponseResults) GetData() []Chat`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ChatListResponseResults) GetDataOk() (*[]Chat, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ChatListResponseResults) SetData(v []Chat)`

SetData sets Data field to given value.

### HasData

`func (o *ChatListResponseResults) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ChatListResponseResults) GetPagination() ChatListResponseResultsPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ChatListResponseResults) GetPaginationOk() (*ChatListResponseResultsPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ChatListResponseResults) SetPagination(v ChatListResponseResultsPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ChatListResponseResults) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



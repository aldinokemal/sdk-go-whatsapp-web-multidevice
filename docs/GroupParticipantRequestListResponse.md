# GroupParticipantRequestListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**GroupParticipantRequestListResponseResults**](GroupParticipantRequestListResponseResults.md) |  | [optional] 

## Methods

### NewGroupParticipantRequestListResponse

`func NewGroupParticipantRequestListResponse() *GroupParticipantRequestListResponse`

NewGroupParticipantRequestListResponse instantiates a new GroupParticipantRequestListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupParticipantRequestListResponseWithDefaults

`func NewGroupParticipantRequestListResponseWithDefaults() *GroupParticipantRequestListResponse`

NewGroupParticipantRequestListResponseWithDefaults instantiates a new GroupParticipantRequestListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GroupParticipantRequestListResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GroupParticipantRequestListResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GroupParticipantRequestListResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GroupParticipantRequestListResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *GroupParticipantRequestListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GroupParticipantRequestListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GroupParticipantRequestListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GroupParticipantRequestListResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *GroupParticipantRequestListResponse) GetResults() GroupParticipantRequestListResponseResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GroupParticipantRequestListResponse) GetResultsOk() (*GroupParticipantRequestListResponseResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GroupParticipantRequestListResponse) SetResults(v GroupParticipantRequestListResponseResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *GroupParticipantRequestListResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



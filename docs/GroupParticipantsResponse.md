# GroupParticipantsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**GroupParticipantsResult**](GroupParticipantsResult.md) |  | [optional] 

## Methods

### NewGroupParticipantsResponse

`func NewGroupParticipantsResponse() *GroupParticipantsResponse`

NewGroupParticipantsResponse instantiates a new GroupParticipantsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupParticipantsResponseWithDefaults

`func NewGroupParticipantsResponseWithDefaults() *GroupParticipantsResponse`

NewGroupParticipantsResponseWithDefaults instantiates a new GroupParticipantsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GroupParticipantsResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GroupParticipantsResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GroupParticipantsResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GroupParticipantsResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *GroupParticipantsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GroupParticipantsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GroupParticipantsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GroupParticipantsResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *GroupParticipantsResponse) GetResults() GroupParticipantsResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GroupParticipantsResponse) GetResultsOk() (*GroupParticipantsResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GroupParticipantsResponse) SetResults(v GroupParticipantsResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *GroupParticipantsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



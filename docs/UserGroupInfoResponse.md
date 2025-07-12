# UserGroupInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Results** | Pointer to **map[string]interface{}** | Group information object (structure may vary) | [optional] 

## Methods

### NewUserGroupInfoResponse

`func NewUserGroupInfoResponse() *UserGroupInfoResponse`

NewUserGroupInfoResponse instantiates a new UserGroupInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupInfoResponseWithDefaults

`func NewUserGroupInfoResponseWithDefaults() *UserGroupInfoResponse`

NewUserGroupInfoResponseWithDefaults instantiates a new UserGroupInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UserGroupInfoResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserGroupInfoResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserGroupInfoResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserGroupInfoResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *UserGroupInfoResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UserGroupInfoResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UserGroupInfoResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UserGroupInfoResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *UserGroupInfoResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserGroupInfoResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserGroupInfoResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UserGroupInfoResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *UserGroupInfoResponse) GetResults() map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *UserGroupInfoResponse) GetResultsOk() (*map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *UserGroupInfoResponse) SetResults(v map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *UserGroupInfoResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



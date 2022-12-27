# UserAvatarResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**UserAvatarResponseResults**](UserAvatarResponseResults.md) |  | [optional] 

## Methods

### NewUserAvatarResponse

`func NewUserAvatarResponse() *UserAvatarResponse`

NewUserAvatarResponse instantiates a new UserAvatarResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAvatarResponseWithDefaults

`func NewUserAvatarResponseWithDefaults() *UserAvatarResponse`

NewUserAvatarResponseWithDefaults instantiates a new UserAvatarResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UserAvatarResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserAvatarResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserAvatarResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UserAvatarResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *UserAvatarResponse) GetResults() UserAvatarResponseResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *UserAvatarResponse) GetResultsOk() (*UserAvatarResponseResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *UserAvatarResponse) SetResults(v UserAvatarResponseResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *UserAvatarResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



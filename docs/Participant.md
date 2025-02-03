# Participant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JID** | Pointer to **string** |  | [optional] 
**LID** | Pointer to **string** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**IsSuperAdmin** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **int32** |  | [optional] 
**AddRequest** | Pointer to **string** |  | [optional] 

## Methods

### NewParticipant

`func NewParticipant() *Participant`

NewParticipant instantiates a new Participant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantWithDefaults

`func NewParticipantWithDefaults() *Participant`

NewParticipantWithDefaults instantiates a new Participant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJID

`func (o *Participant) GetJID() string`

GetJID returns the JID field if non-nil, zero value otherwise.

### GetJIDOk

`func (o *Participant) GetJIDOk() (*string, bool)`

GetJIDOk returns a tuple with the JID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJID

`func (o *Participant) SetJID(v string)`

SetJID sets JID field to given value.

### HasJID

`func (o *Participant) HasJID() bool`

HasJID returns a boolean if a field has been set.

### GetLID

`func (o *Participant) GetLID() string`

GetLID returns the LID field if non-nil, zero value otherwise.

### GetLIDOk

`func (o *Participant) GetLIDOk() (*string, bool)`

GetLIDOk returns a tuple with the LID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLID

`func (o *Participant) SetLID(v string)`

SetLID sets LID field to given value.

### HasLID

`func (o *Participant) HasLID() bool`

HasLID returns a boolean if a field has been set.

### GetIsAdmin

`func (o *Participant) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *Participant) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *Participant) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *Participant) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetIsSuperAdmin

`func (o *Participant) GetIsSuperAdmin() bool`

GetIsSuperAdmin returns the IsSuperAdmin field if non-nil, zero value otherwise.

### GetIsSuperAdminOk

`func (o *Participant) GetIsSuperAdminOk() (*bool, bool)`

GetIsSuperAdminOk returns a tuple with the IsSuperAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperAdmin

`func (o *Participant) SetIsSuperAdmin(v bool)`

SetIsSuperAdmin sets IsSuperAdmin field to given value.

### HasIsSuperAdmin

`func (o *Participant) HasIsSuperAdmin() bool`

HasIsSuperAdmin returns a boolean if a field has been set.

### GetDisplayName

`func (o *Participant) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Participant) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Participant) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Participant) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetError

`func (o *Participant) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Participant) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Participant) SetError(v int32)`

SetError sets Error field to given value.

### HasError

`func (o *Participant) HasError() bool`

HasError returns a boolean if a field has been set.

### GetAddRequest

`func (o *Participant) GetAddRequest() string`

GetAddRequest returns the AddRequest field if non-nil, zero value otherwise.

### GetAddRequestOk

`func (o *Participant) GetAddRequestOk() (*string, bool)`

GetAddRequestOk returns a tuple with the AddRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddRequest

`func (o *Participant) SetAddRequest(v string)`

SetAddRequest sets AddRequest field to given value.

### HasAddRequest

`func (o *Participant) HasAddRequest() bool`

HasAddRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



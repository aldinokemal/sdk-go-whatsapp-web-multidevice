# BusinessProfileResponseResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jid** | Pointer to **string** | Business account JID | [optional] 
**Email** | Pointer to **string** | Business email address | [optional] 
**Address** | Pointer to **string** | Business physical address | [optional] 
**Categories** | Pointer to [**[]BusinessProfileResponseResultsCategoriesInner**](BusinessProfileResponseResultsCategoriesInner.md) | Business categories | [optional] 
**ProfileOptions** | Pointer to **map[string]interface{}** | Additional profile options | [optional] 
**BusinessHoursTimezone** | Pointer to **string** | Business hours timezone | [optional] 
**BusinessHours** | Pointer to [**[]BusinessProfileResponseResultsBusinessHoursInner**](BusinessProfileResponseResultsBusinessHoursInner.md) | Business operating hours | [optional] 

## Methods

### NewBusinessProfileResponseResults

`func NewBusinessProfileResponseResults() *BusinessProfileResponseResults`

NewBusinessProfileResponseResults instantiates a new BusinessProfileResponseResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessProfileResponseResultsWithDefaults

`func NewBusinessProfileResponseResultsWithDefaults() *BusinessProfileResponseResults`

NewBusinessProfileResponseResultsWithDefaults instantiates a new BusinessProfileResponseResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJid

`func (o *BusinessProfileResponseResults) GetJid() string`

GetJid returns the Jid field if non-nil, zero value otherwise.

### GetJidOk

`func (o *BusinessProfileResponseResults) GetJidOk() (*string, bool)`

GetJidOk returns a tuple with the Jid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJid

`func (o *BusinessProfileResponseResults) SetJid(v string)`

SetJid sets Jid field to given value.

### HasJid

`func (o *BusinessProfileResponseResults) HasJid() bool`

HasJid returns a boolean if a field has been set.

### GetEmail

`func (o *BusinessProfileResponseResults) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BusinessProfileResponseResults) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BusinessProfileResponseResults) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BusinessProfileResponseResults) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAddress

`func (o *BusinessProfileResponseResults) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BusinessProfileResponseResults) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BusinessProfileResponseResults) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BusinessProfileResponseResults) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCategories

`func (o *BusinessProfileResponseResults) GetCategories() []BusinessProfileResponseResultsCategoriesInner`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *BusinessProfileResponseResults) GetCategoriesOk() (*[]BusinessProfileResponseResultsCategoriesInner, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *BusinessProfileResponseResults) SetCategories(v []BusinessProfileResponseResultsCategoriesInner)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *BusinessProfileResponseResults) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetProfileOptions

`func (o *BusinessProfileResponseResults) GetProfileOptions() map[string]interface{}`

GetProfileOptions returns the ProfileOptions field if non-nil, zero value otherwise.

### GetProfileOptionsOk

`func (o *BusinessProfileResponseResults) GetProfileOptionsOk() (*map[string]interface{}, bool)`

GetProfileOptionsOk returns a tuple with the ProfileOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileOptions

`func (o *BusinessProfileResponseResults) SetProfileOptions(v map[string]interface{})`

SetProfileOptions sets ProfileOptions field to given value.

### HasProfileOptions

`func (o *BusinessProfileResponseResults) HasProfileOptions() bool`

HasProfileOptions returns a boolean if a field has been set.

### GetBusinessHoursTimezone

`func (o *BusinessProfileResponseResults) GetBusinessHoursTimezone() string`

GetBusinessHoursTimezone returns the BusinessHoursTimezone field if non-nil, zero value otherwise.

### GetBusinessHoursTimezoneOk

`func (o *BusinessProfileResponseResults) GetBusinessHoursTimezoneOk() (*string, bool)`

GetBusinessHoursTimezoneOk returns a tuple with the BusinessHoursTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessHoursTimezone

`func (o *BusinessProfileResponseResults) SetBusinessHoursTimezone(v string)`

SetBusinessHoursTimezone sets BusinessHoursTimezone field to given value.

### HasBusinessHoursTimezone

`func (o *BusinessProfileResponseResults) HasBusinessHoursTimezone() bool`

HasBusinessHoursTimezone returns a boolean if a field has been set.

### GetBusinessHours

`func (o *BusinessProfileResponseResults) GetBusinessHours() []BusinessProfileResponseResultsBusinessHoursInner`

GetBusinessHours returns the BusinessHours field if non-nil, zero value otherwise.

### GetBusinessHoursOk

`func (o *BusinessProfileResponseResults) GetBusinessHoursOk() (*[]BusinessProfileResponseResultsBusinessHoursInner, bool)`

GetBusinessHoursOk returns a tuple with the BusinessHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessHours

`func (o *BusinessProfileResponseResults) SetBusinessHours(v []BusinessProfileResponseResultsBusinessHoursInner)`

SetBusinessHours sets BusinessHours field to given value.

### HasBusinessHours

`func (o *BusinessProfileResponseResults) HasBusinessHours() bool`

HasBusinessHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



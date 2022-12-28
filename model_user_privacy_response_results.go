/*
WhatsApp API MultiDevice

This API is used for sending whatsapp via API

API version: 3.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserPrivacyResponseResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPrivacyResponseResults{}

// UserPrivacyResponseResults struct for UserPrivacyResponseResults
type UserPrivacyResponseResults struct {
	GroupAdd *string `json:"group_add,omitempty"`
	LastSeen *string `json:"last_seen,omitempty"`
	Status *string `json:"status,omitempty"`
	Profile *string `json:"profile,omitempty"`
	ReadReceipts *string `json:"read_receipts,omitempty"`
}

// NewUserPrivacyResponseResults instantiates a new UserPrivacyResponseResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPrivacyResponseResults() *UserPrivacyResponseResults {
	this := UserPrivacyResponseResults{}
	return &this
}

// NewUserPrivacyResponseResultsWithDefaults instantiates a new UserPrivacyResponseResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPrivacyResponseResultsWithDefaults() *UserPrivacyResponseResults {
	this := UserPrivacyResponseResults{}
	return &this
}

// GetGroupAdd returns the GroupAdd field value if set, zero value otherwise.
func (o *UserPrivacyResponseResults) GetGroupAdd() string {
	if o == nil || isNil(o.GroupAdd) {
		var ret string
		return ret
	}
	return *o.GroupAdd
}

// GetGroupAddOk returns a tuple with the GroupAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrivacyResponseResults) GetGroupAddOk() (*string, bool) {
	if o == nil || isNil(o.GroupAdd) {
		return nil, false
	}
	return o.GroupAdd, true
}

// HasGroupAdd returns a boolean if a field has been set.
func (o *UserPrivacyResponseResults) HasGroupAdd() bool {
	if o != nil && !isNil(o.GroupAdd) {
		return true
	}

	return false
}

// SetGroupAdd gets a reference to the given string and assigns it to the GroupAdd field.
func (o *UserPrivacyResponseResults) SetGroupAdd(v string) {
	o.GroupAdd = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *UserPrivacyResponseResults) GetLastSeen() string {
	if o == nil || isNil(o.LastSeen) {
		var ret string
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrivacyResponseResults) GetLastSeenOk() (*string, bool) {
	if o == nil || isNil(o.LastSeen) {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *UserPrivacyResponseResults) HasLastSeen() bool {
	if o != nil && !isNil(o.LastSeen) {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given string and assigns it to the LastSeen field.
func (o *UserPrivacyResponseResults) SetLastSeen(v string) {
	o.LastSeen = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserPrivacyResponseResults) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrivacyResponseResults) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserPrivacyResponseResults) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UserPrivacyResponseResults) SetStatus(v string) {
	o.Status = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserPrivacyResponseResults) GetProfile() string {
	if o == nil || isNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrivacyResponseResults) GetProfileOk() (*string, bool) {
	if o == nil || isNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserPrivacyResponseResults) HasProfile() bool {
	if o != nil && !isNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *UserPrivacyResponseResults) SetProfile(v string) {
	o.Profile = &v
}

// GetReadReceipts returns the ReadReceipts field value if set, zero value otherwise.
func (o *UserPrivacyResponseResults) GetReadReceipts() string {
	if o == nil || isNil(o.ReadReceipts) {
		var ret string
		return ret
	}
	return *o.ReadReceipts
}

// GetReadReceiptsOk returns a tuple with the ReadReceipts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrivacyResponseResults) GetReadReceiptsOk() (*string, bool) {
	if o == nil || isNil(o.ReadReceipts) {
		return nil, false
	}
	return o.ReadReceipts, true
}

// HasReadReceipts returns a boolean if a field has been set.
func (o *UserPrivacyResponseResults) HasReadReceipts() bool {
	if o != nil && !isNil(o.ReadReceipts) {
		return true
	}

	return false
}

// SetReadReceipts gets a reference to the given string and assigns it to the ReadReceipts field.
func (o *UserPrivacyResponseResults) SetReadReceipts(v string) {
	o.ReadReceipts = &v
}

func (o UserPrivacyResponseResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPrivacyResponseResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GroupAdd) {
		toSerialize["group_add"] = o.GroupAdd
	}
	if !isNil(o.LastSeen) {
		toSerialize["last_seen"] = o.LastSeen
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !isNil(o.ReadReceipts) {
		toSerialize["read_receipts"] = o.ReadReceipts
	}
	return toSerialize, nil
}

type NullableUserPrivacyResponseResults struct {
	value *UserPrivacyResponseResults
	isSet bool
}

func (v NullableUserPrivacyResponseResults) Get() *UserPrivacyResponseResults {
	return v.value
}

func (v *NullableUserPrivacyResponseResults) Set(val *UserPrivacyResponseResults) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPrivacyResponseResults) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPrivacyResponseResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPrivacyResponseResults(val *UserPrivacyResponseResults) *NullableUserPrivacyResponseResults {
	return &NullableUserPrivacyResponseResults{value: val, isSet: true}
}

func (v NullableUserPrivacyResponseResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPrivacyResponseResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



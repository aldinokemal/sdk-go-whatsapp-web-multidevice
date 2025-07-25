/*
WhatsApp API MultiDevice

This API is used for sending whatsapp via API

API version: 6.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SdkWhatsappWebMultiDevice

import (
	"encoding/json"
)

// checks if the UserGroupResponseResultsDataInnerParticipantsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGroupResponseResultsDataInnerParticipantsInner{}

// UserGroupResponseResultsDataInnerParticipantsInner struct for UserGroupResponseResultsDataInnerParticipantsInner
type UserGroupResponseResultsDataInnerParticipantsInner struct {
	JID *string `json:"JID,omitempty"`
	IsAdmin *bool `json:"IsAdmin,omitempty"`
	IsSuperAdmin *bool `json:"IsSuperAdmin,omitempty"`
	Error *float32 `json:"Error,omitempty"`
}

// NewUserGroupResponseResultsDataInnerParticipantsInner instantiates a new UserGroupResponseResultsDataInnerParticipantsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupResponseResultsDataInnerParticipantsInner() *UserGroupResponseResultsDataInnerParticipantsInner {
	this := UserGroupResponseResultsDataInnerParticipantsInner{}
	return &this
}

// NewUserGroupResponseResultsDataInnerParticipantsInnerWithDefaults instantiates a new UserGroupResponseResultsDataInnerParticipantsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupResponseResultsDataInnerParticipantsInnerWithDefaults() *UserGroupResponseResultsDataInnerParticipantsInner {
	this := UserGroupResponseResultsDataInnerParticipantsInner{}
	return &this
}

// GetJID returns the JID field value if set, zero value otherwise.
func (o *UserGroupResponseResultsDataInnerParticipantsInner) GetJID() string {
	if o == nil || IsNil(o.JID) {
		var ret string
		return ret
	}
	return *o.JID
}

// GetJIDOk returns a tuple with the JID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupResponseResultsDataInnerParticipantsInner) GetJIDOk() (*string, bool) {
	if o == nil || IsNil(o.JID) {
		return nil, false
	}
	return o.JID, true
}

// HasJID returns a boolean if a field has been set.
func (o *UserGroupResponseResultsDataInnerParticipantsInner) HasJID() bool {
	if o != nil && !IsNil(o.JID) {
		return true
	}

	return false
}

// SetJID gets a reference to the given string and assigns it to the JID field.
func (o *UserGroupResponseResultsDataInnerParticipantsInner) SetJID(v string) {
	o.JID = &v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *UserGroupResponseResultsDataInnerParticipantsInner) GetIsAdmin() bool {
	if o == nil || IsNil(o.IsAdmin) {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupResponseResultsDataInnerParticipantsInner) GetIsAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAdmin) {
		return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *UserGroupResponseResultsDataInnerParticipantsInner) HasIsAdmin() bool {
	if o != nil && !IsNil(o.IsAdmin) {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *UserGroupResponseResultsDataInnerParticipantsInner) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

// GetIsSuperAdmin returns the IsSuperAdmin field value if set, zero value otherwise.
func (o *UserGroupResponseResultsDataInnerParticipantsInner) GetIsSuperAdmin() bool {
	if o == nil || IsNil(o.IsSuperAdmin) {
		var ret bool
		return ret
	}
	return *o.IsSuperAdmin
}

// GetIsSuperAdminOk returns a tuple with the IsSuperAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupResponseResultsDataInnerParticipantsInner) GetIsSuperAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSuperAdmin) {
		return nil, false
	}
	return o.IsSuperAdmin, true
}

// HasIsSuperAdmin returns a boolean if a field has been set.
func (o *UserGroupResponseResultsDataInnerParticipantsInner) HasIsSuperAdmin() bool {
	if o != nil && !IsNil(o.IsSuperAdmin) {
		return true
	}

	return false
}

// SetIsSuperAdmin gets a reference to the given bool and assigns it to the IsSuperAdmin field.
func (o *UserGroupResponseResultsDataInnerParticipantsInner) SetIsSuperAdmin(v bool) {
	o.IsSuperAdmin = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *UserGroupResponseResultsDataInnerParticipantsInner) GetError() float32 {
	if o == nil || IsNil(o.Error) {
		var ret float32
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupResponseResultsDataInnerParticipantsInner) GetErrorOk() (*float32, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *UserGroupResponseResultsDataInnerParticipantsInner) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given float32 and assigns it to the Error field.
func (o *UserGroupResponseResultsDataInnerParticipantsInner) SetError(v float32) {
	o.Error = &v
}

func (o UserGroupResponseResultsDataInnerParticipantsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGroupResponseResultsDataInnerParticipantsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JID) {
		toSerialize["JID"] = o.JID
	}
	if !IsNil(o.IsAdmin) {
		toSerialize["IsAdmin"] = o.IsAdmin
	}
	if !IsNil(o.IsSuperAdmin) {
		toSerialize["IsSuperAdmin"] = o.IsSuperAdmin
	}
	if !IsNil(o.Error) {
		toSerialize["Error"] = o.Error
	}
	return toSerialize, nil
}

type NullableUserGroupResponseResultsDataInnerParticipantsInner struct {
	value *UserGroupResponseResultsDataInnerParticipantsInner
	isSet bool
}

func (v NullableUserGroupResponseResultsDataInnerParticipantsInner) Get() *UserGroupResponseResultsDataInnerParticipantsInner {
	return v.value
}

func (v *NullableUserGroupResponseResultsDataInnerParticipantsInner) Set(val *UserGroupResponseResultsDataInnerParticipantsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupResponseResultsDataInnerParticipantsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupResponseResultsDataInnerParticipantsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupResponseResultsDataInnerParticipantsInner(val *UserGroupResponseResultsDataInnerParticipantsInner) *NullableUserGroupResponseResultsDataInnerParticipantsInner {
	return &NullableUserGroupResponseResultsDataInnerParticipantsInner{value: val, isSet: true}
}

func (v NullableUserGroupResponseResultsDataInnerParticipantsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupResponseResultsDataInnerParticipantsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



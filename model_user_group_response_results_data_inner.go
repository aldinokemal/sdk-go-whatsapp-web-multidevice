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

// checks if the UserGroupResponseResultsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGroupResponseResultsDataInner{}

// UserGroupResponseResultsDataInner struct for UserGroupResponseResultsDataInner
type UserGroupResponseResultsDataInner struct {
	JID *string `json:"JID,omitempty"`
	OwnerJID *string `json:"OwnerJID,omitempty"`
	Name *string `json:"Name,omitempty"`
	NameSetAt *string `json:"NameSetAt,omitempty"`
	NameSetBy *string `json:"NameSetBy,omitempty"`
	GroupCreated *string `json:"GroupCreated,omitempty"`
	ParticipantVersionID *string `json:"ParticipantVersionID,omitempty"`
	Participants []UserGroupResponseResultsDataInnerParticipantsInner `json:"Participants,omitempty"`
}

// NewUserGroupResponseResultsDataInner instantiates a new UserGroupResponseResultsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupResponseResultsDataInner() *UserGroupResponseResultsDataInner {
	this := UserGroupResponseResultsDataInner{}
	return &this
}

// NewUserGroupResponseResultsDataInnerWithDefaults instantiates a new UserGroupResponseResultsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupResponseResultsDataInnerWithDefaults() *UserGroupResponseResultsDataInner {
	this := UserGroupResponseResultsDataInner{}
	return &this
}

// GetJID returns the JID field value if set, zero value otherwise.
func (o *UserGroupResponseResultsDataInner) GetJID() string {
	if o == nil || isNil(o.JID) {
		var ret string
		return ret
	}
	return *o.JID
}

// GetJIDOk returns a tuple with the JID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupResponseResultsDataInner) GetJIDOk() (*string, bool) {
	if o == nil || isNil(o.JID) {
		return nil, false
	}
	return o.JID, true
}

// HasJID returns a boolean if a field has been set.
func (o *UserGroupResponseResultsDataInner) HasJID() bool {
	if o != nil && !isNil(o.JID) {
		return true
	}

	return false
}

// SetJID gets a reference to the given string and assigns it to the JID field.
func (o *UserGroupResponseResultsDataInner) SetJID(v string) {
	o.JID = &v
}

// GetOwnerJID returns the OwnerJID field value if set, zero value otherwise.
func (o *UserGroupResponseResultsDataInner) GetOwnerJID() string {
	if o == nil || isNil(o.OwnerJID) {
		var ret string
		return ret
	}
	return *o.OwnerJID
}

// GetOwnerJIDOk returns a tuple with the OwnerJID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupResponseResultsDataInner) GetOwnerJIDOk() (*string, bool) {
	if o == nil || isNil(o.OwnerJID) {
		return nil, false
	}
	return o.OwnerJID, true
}

// HasOwnerJID returns a boolean if a field has been set.
func (o *UserGroupResponseResultsDataInner) HasOwnerJID() bool {
	if o != nil && !isNil(o.OwnerJID) {
		return true
	}

	return false
}

// SetOwnerJID gets a reference to the given string and assigns it to the OwnerJID field.
func (o *UserGroupResponseResultsDataInner) SetOwnerJID(v string) {
	o.OwnerJID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserGroupResponseResultsDataInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupResponseResultsDataInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserGroupResponseResultsDataInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserGroupResponseResultsDataInner) SetName(v string) {
	o.Name = &v
}

// GetNameSetAt returns the NameSetAt field value if set, zero value otherwise.
func (o *UserGroupResponseResultsDataInner) GetNameSetAt() string {
	if o == nil || isNil(o.NameSetAt) {
		var ret string
		return ret
	}
	return *o.NameSetAt
}

// GetNameSetAtOk returns a tuple with the NameSetAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupResponseResultsDataInner) GetNameSetAtOk() (*string, bool) {
	if o == nil || isNil(o.NameSetAt) {
		return nil, false
	}
	return o.NameSetAt, true
}

// HasNameSetAt returns a boolean if a field has been set.
func (o *UserGroupResponseResultsDataInner) HasNameSetAt() bool {
	if o != nil && !isNil(o.NameSetAt) {
		return true
	}

	return false
}

// SetNameSetAt gets a reference to the given string and assigns it to the NameSetAt field.
func (o *UserGroupResponseResultsDataInner) SetNameSetAt(v string) {
	o.NameSetAt = &v
}

// GetNameSetBy returns the NameSetBy field value if set, zero value otherwise.
func (o *UserGroupResponseResultsDataInner) GetNameSetBy() string {
	if o == nil || isNil(o.NameSetBy) {
		var ret string
		return ret
	}
	return *o.NameSetBy
}

// GetNameSetByOk returns a tuple with the NameSetBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupResponseResultsDataInner) GetNameSetByOk() (*string, bool) {
	if o == nil || isNil(o.NameSetBy) {
		return nil, false
	}
	return o.NameSetBy, true
}

// HasNameSetBy returns a boolean if a field has been set.
func (o *UserGroupResponseResultsDataInner) HasNameSetBy() bool {
	if o != nil && !isNil(o.NameSetBy) {
		return true
	}

	return false
}

// SetNameSetBy gets a reference to the given string and assigns it to the NameSetBy field.
func (o *UserGroupResponseResultsDataInner) SetNameSetBy(v string) {
	o.NameSetBy = &v
}

// GetGroupCreated returns the GroupCreated field value if set, zero value otherwise.
func (o *UserGroupResponseResultsDataInner) GetGroupCreated() string {
	if o == nil || isNil(o.GroupCreated) {
		var ret string
		return ret
	}
	return *o.GroupCreated
}

// GetGroupCreatedOk returns a tuple with the GroupCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupResponseResultsDataInner) GetGroupCreatedOk() (*string, bool) {
	if o == nil || isNil(o.GroupCreated) {
		return nil, false
	}
	return o.GroupCreated, true
}

// HasGroupCreated returns a boolean if a field has been set.
func (o *UserGroupResponseResultsDataInner) HasGroupCreated() bool {
	if o != nil && !isNil(o.GroupCreated) {
		return true
	}

	return false
}

// SetGroupCreated gets a reference to the given string and assigns it to the GroupCreated field.
func (o *UserGroupResponseResultsDataInner) SetGroupCreated(v string) {
	o.GroupCreated = &v
}

// GetParticipantVersionID returns the ParticipantVersionID field value if set, zero value otherwise.
func (o *UserGroupResponseResultsDataInner) GetParticipantVersionID() string {
	if o == nil || isNil(o.ParticipantVersionID) {
		var ret string
		return ret
	}
	return *o.ParticipantVersionID
}

// GetParticipantVersionIDOk returns a tuple with the ParticipantVersionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupResponseResultsDataInner) GetParticipantVersionIDOk() (*string, bool) {
	if o == nil || isNil(o.ParticipantVersionID) {
		return nil, false
	}
	return o.ParticipantVersionID, true
}

// HasParticipantVersionID returns a boolean if a field has been set.
func (o *UserGroupResponseResultsDataInner) HasParticipantVersionID() bool {
	if o != nil && !isNil(o.ParticipantVersionID) {
		return true
	}

	return false
}

// SetParticipantVersionID gets a reference to the given string and assigns it to the ParticipantVersionID field.
func (o *UserGroupResponseResultsDataInner) SetParticipantVersionID(v string) {
	o.ParticipantVersionID = &v
}

// GetParticipants returns the Participants field value if set, zero value otherwise.
func (o *UserGroupResponseResultsDataInner) GetParticipants() []UserGroupResponseResultsDataInnerParticipantsInner {
	if o == nil || isNil(o.Participants) {
		var ret []UserGroupResponseResultsDataInnerParticipantsInner
		return ret
	}
	return o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupResponseResultsDataInner) GetParticipantsOk() ([]UserGroupResponseResultsDataInnerParticipantsInner, bool) {
	if o == nil || isNil(o.Participants) {
		return nil, false
	}
	return o.Participants, true
}

// HasParticipants returns a boolean if a field has been set.
func (o *UserGroupResponseResultsDataInner) HasParticipants() bool {
	if o != nil && !isNil(o.Participants) {
		return true
	}

	return false
}

// SetParticipants gets a reference to the given []UserGroupResponseResultsDataInnerParticipantsInner and assigns it to the Participants field.
func (o *UserGroupResponseResultsDataInner) SetParticipants(v []UserGroupResponseResultsDataInnerParticipantsInner) {
	o.Participants = v
}

func (o UserGroupResponseResultsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGroupResponseResultsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JID) {
		toSerialize["JID"] = o.JID
	}
	if !isNil(o.OwnerJID) {
		toSerialize["OwnerJID"] = o.OwnerJID
	}
	if !isNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !isNil(o.NameSetAt) {
		toSerialize["NameSetAt"] = o.NameSetAt
	}
	if !isNil(o.NameSetBy) {
		toSerialize["NameSetBy"] = o.NameSetBy
	}
	if !isNil(o.GroupCreated) {
		toSerialize["GroupCreated"] = o.GroupCreated
	}
	if !isNil(o.ParticipantVersionID) {
		toSerialize["ParticipantVersionID"] = o.ParticipantVersionID
	}
	if !isNil(o.Participants) {
		toSerialize["Participants"] = o.Participants
	}
	return toSerialize, nil
}

type NullableUserGroupResponseResultsDataInner struct {
	value *UserGroupResponseResultsDataInner
	isSet bool
}

func (v NullableUserGroupResponseResultsDataInner) Get() *UserGroupResponseResultsDataInner {
	return v.value
}

func (v *NullableUserGroupResponseResultsDataInner) Set(val *UserGroupResponseResultsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupResponseResultsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupResponseResultsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupResponseResultsDataInner(val *UserGroupResponseResultsDataInner) *NullableUserGroupResponseResultsDataInner {
	return &NullableUserGroupResponseResultsDataInner{value: val, isSet: true}
}

func (v NullableUserGroupResponseResultsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupResponseResultsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



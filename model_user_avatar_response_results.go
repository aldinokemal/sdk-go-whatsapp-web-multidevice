/*
WhatsApp API MultiDevice

This API is used for sending whatsapp via API

API version: 3.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk-go-whatsapp-web-multidevice

import (
	"encoding/json"
)

// checks if the UserAvatarResponseResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAvatarResponseResults{}

// UserAvatarResponseResults struct for UserAvatarResponseResults
type UserAvatarResponseResults struct {
	Url *string `json:"url,omitempty"`
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewUserAvatarResponseResults instantiates a new UserAvatarResponseResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAvatarResponseResults() *UserAvatarResponseResults {
	this := UserAvatarResponseResults{}
	return &this
}

// NewUserAvatarResponseResultsWithDefaults instantiates a new UserAvatarResponseResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAvatarResponseResultsWithDefaults() *UserAvatarResponseResults {
	this := UserAvatarResponseResults{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UserAvatarResponseResults) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAvatarResponseResults) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UserAvatarResponseResults) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UserAvatarResponseResults) SetUrl(v string) {
	o.Url = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserAvatarResponseResults) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAvatarResponseResults) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserAvatarResponseResults) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserAvatarResponseResults) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserAvatarResponseResults) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAvatarResponseResults) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserAvatarResponseResults) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserAvatarResponseResults) SetType(v string) {
	o.Type = &v
}

func (o UserAvatarResponseResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAvatarResponseResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUserAvatarResponseResults struct {
	value *UserAvatarResponseResults
	isSet bool
}

func (v NullableUserAvatarResponseResults) Get() *UserAvatarResponseResults {
	return v.value
}

func (v *NullableUserAvatarResponseResults) Set(val *UserAvatarResponseResults) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAvatarResponseResults) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAvatarResponseResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAvatarResponseResults(val *UserAvatarResponseResults) *NullableUserAvatarResponseResults {
	return &NullableUserAvatarResponseResults{value: val, isSet: true}
}

func (v NullableUserAvatarResponseResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAvatarResponseResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



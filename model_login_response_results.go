/*
WhatsApp API MultiDevice

This API is used for sending whatsapp via API

API version: 3.10.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk_go_whatsapp_web_multidevice

import (
	"encoding/json"
)

// checks if the LoginResponseResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginResponseResults{}

// LoginResponseResults struct for LoginResponseResults
type LoginResponseResults struct {
	QrDuration *int32 `json:"qr_duration,omitempty"`
	QrLink *string `json:"qr_link,omitempty"`
}

// NewLoginResponseResults instantiates a new LoginResponseResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginResponseResults() *LoginResponseResults {
	this := LoginResponseResults{}
	return &this
}

// NewLoginResponseResultsWithDefaults instantiates a new LoginResponseResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginResponseResultsWithDefaults() *LoginResponseResults {
	this := LoginResponseResults{}
	return &this
}

// GetQrDuration returns the QrDuration field value if set, zero value otherwise.
func (o *LoginResponseResults) GetQrDuration() int32 {
	if o == nil || IsNil(o.QrDuration) {
		var ret int32
		return ret
	}
	return *o.QrDuration
}

// GetQrDurationOk returns a tuple with the QrDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponseResults) GetQrDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.QrDuration) {
		return nil, false
	}
	return o.QrDuration, true
}

// HasQrDuration returns a boolean if a field has been set.
func (o *LoginResponseResults) HasQrDuration() bool {
	if o != nil && !IsNil(o.QrDuration) {
		return true
	}

	return false
}

// SetQrDuration gets a reference to the given int32 and assigns it to the QrDuration field.
func (o *LoginResponseResults) SetQrDuration(v int32) {
	o.QrDuration = &v
}

// GetQrLink returns the QrLink field value if set, zero value otherwise.
func (o *LoginResponseResults) GetQrLink() string {
	if o == nil || IsNil(o.QrLink) {
		var ret string
		return ret
	}
	return *o.QrLink
}

// GetQrLinkOk returns a tuple with the QrLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponseResults) GetQrLinkOk() (*string, bool) {
	if o == nil || IsNil(o.QrLink) {
		return nil, false
	}
	return o.QrLink, true
}

// HasQrLink returns a boolean if a field has been set.
func (o *LoginResponseResults) HasQrLink() bool {
	if o != nil && !IsNil(o.QrLink) {
		return true
	}

	return false
}

// SetQrLink gets a reference to the given string and assigns it to the QrLink field.
func (o *LoginResponseResults) SetQrLink(v string) {
	o.QrLink = &v
}

func (o LoginResponseResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginResponseResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QrDuration) {
		toSerialize["qr_duration"] = o.QrDuration
	}
	if !IsNil(o.QrLink) {
		toSerialize["qr_link"] = o.QrLink
	}
	return toSerialize, nil
}

type NullableLoginResponseResults struct {
	value *LoginResponseResults
	isSet bool
}

func (v NullableLoginResponseResults) Get() *LoginResponseResults {
	return v.value
}

func (v *NullableLoginResponseResults) Set(val *LoginResponseResults) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginResponseResults) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginResponseResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginResponseResults(val *LoginResponseResults) *NullableLoginResponseResults {
	return &NullableLoginResponseResults{value: val, isSet: true}
}

func (v NullableLoginResponseResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginResponseResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



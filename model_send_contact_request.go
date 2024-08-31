/*
WhatsApp API MultiDevice

This API is used for sending whatsapp via API

API version: 4.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk_go_whatsapp_web_multidevice

import (
	"encoding/json"
)

// checks if the SendContactRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendContactRequest{}

// SendContactRequest struct for SendContactRequest
type SendContactRequest struct {
	// Phone number with country code
	Phone *string `json:"phone,omitempty"`
	// Contact name
	ContactName *string `json:"contact_name,omitempty"`
	// Contact phone number
	ContactPhone *string `json:"contact_phone,omitempty"`
}

// NewSendContactRequest instantiates a new SendContactRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendContactRequest() *SendContactRequest {
	this := SendContactRequest{}
	return &this
}

// NewSendContactRequestWithDefaults instantiates a new SendContactRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendContactRequestWithDefaults() *SendContactRequest {
	this := SendContactRequest{}
	return &this
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *SendContactRequest) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendContactRequest) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *SendContactRequest) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *SendContactRequest) SetPhone(v string) {
	o.Phone = &v
}

// GetContactName returns the ContactName field value if set, zero value otherwise.
func (o *SendContactRequest) GetContactName() string {
	if o == nil || IsNil(o.ContactName) {
		var ret string
		return ret
	}
	return *o.ContactName
}

// GetContactNameOk returns a tuple with the ContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendContactRequest) GetContactNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContactName) {
		return nil, false
	}
	return o.ContactName, true
}

// HasContactName returns a boolean if a field has been set.
func (o *SendContactRequest) HasContactName() bool {
	if o != nil && !IsNil(o.ContactName) {
		return true
	}

	return false
}

// SetContactName gets a reference to the given string and assigns it to the ContactName field.
func (o *SendContactRequest) SetContactName(v string) {
	o.ContactName = &v
}

// GetContactPhone returns the ContactPhone field value if set, zero value otherwise.
func (o *SendContactRequest) GetContactPhone() string {
	if o == nil || IsNil(o.ContactPhone) {
		var ret string
		return ret
	}
	return *o.ContactPhone
}

// GetContactPhoneOk returns a tuple with the ContactPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendContactRequest) GetContactPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.ContactPhone) {
		return nil, false
	}
	return o.ContactPhone, true
}

// HasContactPhone returns a boolean if a field has been set.
func (o *SendContactRequest) HasContactPhone() bool {
	if o != nil && !IsNil(o.ContactPhone) {
		return true
	}

	return false
}

// SetContactPhone gets a reference to the given string and assigns it to the ContactPhone field.
func (o *SendContactRequest) SetContactPhone(v string) {
	o.ContactPhone = &v
}

func (o SendContactRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendContactRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.ContactName) {
		toSerialize["contact_name"] = o.ContactName
	}
	if !IsNil(o.ContactPhone) {
		toSerialize["contact_phone"] = o.ContactPhone
	}
	return toSerialize, nil
}

type NullableSendContactRequest struct {
	value *SendContactRequest
	isSet bool
}

func (v NullableSendContactRequest) Get() *SendContactRequest {
	return v.value
}

func (v *NullableSendContactRequest) Set(val *SendContactRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendContactRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendContactRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendContactRequest(val *SendContactRequest) *NullableSendContactRequest {
	return &NullableSendContactRequest{value: val, isSet: true}
}

func (v NullableSendContactRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendContactRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



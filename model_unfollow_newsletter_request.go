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

// checks if the UnfollowNewsletterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnfollowNewsletterRequest{}

// UnfollowNewsletterRequest struct for UnfollowNewsletterRequest
type UnfollowNewsletterRequest struct {
	NewsletterId *string `json:"newsletter_id,omitempty"`
}

// NewUnfollowNewsletterRequest instantiates a new UnfollowNewsletterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnfollowNewsletterRequest() *UnfollowNewsletterRequest {
	this := UnfollowNewsletterRequest{}
	return &this
}

// NewUnfollowNewsletterRequestWithDefaults instantiates a new UnfollowNewsletterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnfollowNewsletterRequestWithDefaults() *UnfollowNewsletterRequest {
	this := UnfollowNewsletterRequest{}
	return &this
}

// GetNewsletterId returns the NewsletterId field value if set, zero value otherwise.
func (o *UnfollowNewsletterRequest) GetNewsletterId() string {
	if o == nil || IsNil(o.NewsletterId) {
		var ret string
		return ret
	}
	return *o.NewsletterId
}

// GetNewsletterIdOk returns a tuple with the NewsletterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnfollowNewsletterRequest) GetNewsletterIdOk() (*string, bool) {
	if o == nil || IsNil(o.NewsletterId) {
		return nil, false
	}
	return o.NewsletterId, true
}

// HasNewsletterId returns a boolean if a field has been set.
func (o *UnfollowNewsletterRequest) HasNewsletterId() bool {
	if o != nil && !IsNil(o.NewsletterId) {
		return true
	}

	return false
}

// SetNewsletterId gets a reference to the given string and assigns it to the NewsletterId field.
func (o *UnfollowNewsletterRequest) SetNewsletterId(v string) {
	o.NewsletterId = &v
}

func (o UnfollowNewsletterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnfollowNewsletterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewsletterId) {
		toSerialize["newsletter_id"] = o.NewsletterId
	}
	return toSerialize, nil
}

type NullableUnfollowNewsletterRequest struct {
	value *UnfollowNewsletterRequest
	isSet bool
}

func (v NullableUnfollowNewsletterRequest) Get() *UnfollowNewsletterRequest {
	return v.value
}

func (v *NullableUnfollowNewsletterRequest) Set(val *UnfollowNewsletterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnfollowNewsletterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnfollowNewsletterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnfollowNewsletterRequest(val *UnfollowNewsletterRequest) *NullableUnfollowNewsletterRequest {
	return &NullableUnfollowNewsletterRequest{value: val, isSet: true}
}

func (v NullableUnfollowNewsletterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnfollowNewsletterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



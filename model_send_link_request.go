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

// checks if the SendLinkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendLinkRequest{}

// SendLinkRequest struct for SendLinkRequest
type SendLinkRequest struct {
	// Phone number with country code
	Phone *string `json:"phone,omitempty"`
	// Link to send
	Link *string `json:"link,omitempty"`
	// Caption to send
	Caption *string `json:"caption,omitempty"`
}

// NewSendLinkRequest instantiates a new SendLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendLinkRequest() *SendLinkRequest {
	this := SendLinkRequest{}
	return &this
}

// NewSendLinkRequestWithDefaults instantiates a new SendLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendLinkRequestWithDefaults() *SendLinkRequest {
	this := SendLinkRequest{}
	return &this
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *SendLinkRequest) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendLinkRequest) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *SendLinkRequest) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *SendLinkRequest) SetPhone(v string) {
	o.Phone = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *SendLinkRequest) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendLinkRequest) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *SendLinkRequest) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *SendLinkRequest) SetLink(v string) {
	o.Link = &v
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *SendLinkRequest) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendLinkRequest) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *SendLinkRequest) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *SendLinkRequest) SetCaption(v string) {
	o.Caption = &v
}

func (o SendLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendLinkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Caption) {
		toSerialize["caption"] = o.Caption
	}
	return toSerialize, nil
}

type NullableSendLinkRequest struct {
	value *SendLinkRequest
	isSet bool
}

func (v NullableSendLinkRequest) Get() *SendLinkRequest {
	return v.value
}

func (v *NullableSendLinkRequest) Set(val *SendLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendLinkRequest(val *SendLinkRequest) *NullableSendLinkRequest {
	return &NullableSendLinkRequest{value: val, isSet: true}
}

func (v NullableSendLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



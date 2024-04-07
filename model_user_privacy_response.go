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

// checks if the UserPrivacyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPrivacyResponse{}

// UserPrivacyResponse struct for UserPrivacyResponse
type UserPrivacyResponse struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Results *UserPrivacyResponseResults `json:"results,omitempty"`
}

// NewUserPrivacyResponse instantiates a new UserPrivacyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPrivacyResponse() *UserPrivacyResponse {
	this := UserPrivacyResponse{}
	return &this
}

// NewUserPrivacyResponseWithDefaults instantiates a new UserPrivacyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPrivacyResponseWithDefaults() *UserPrivacyResponse {
	this := UserPrivacyResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UserPrivacyResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrivacyResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UserPrivacyResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UserPrivacyResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UserPrivacyResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrivacyResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UserPrivacyResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UserPrivacyResponse) SetMessage(v string) {
	o.Message = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *UserPrivacyResponse) GetResults() UserPrivacyResponseResults {
	if o == nil || IsNil(o.Results) {
		var ret UserPrivacyResponseResults
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrivacyResponse) GetResultsOk() (*UserPrivacyResponseResults, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *UserPrivacyResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given UserPrivacyResponseResults and assigns it to the Results field.
func (o *UserPrivacyResponse) SetResults(v UserPrivacyResponseResults) {
	o.Results = &v
}

func (o UserPrivacyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPrivacyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableUserPrivacyResponse struct {
	value *UserPrivacyResponse
	isSet bool
}

func (v NullableUserPrivacyResponse) Get() *UserPrivacyResponse {
	return v.value
}

func (v *NullableUserPrivacyResponse) Set(val *UserPrivacyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPrivacyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPrivacyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPrivacyResponse(val *UserPrivacyResponse) *NullableUserPrivacyResponse {
	return &NullableUserPrivacyResponse{value: val, isSet: true}
}

func (v NullableUserPrivacyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPrivacyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



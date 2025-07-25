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

// checks if the UserGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGroupResponse{}

// UserGroupResponse struct for UserGroupResponse
type UserGroupResponse struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Results *UserGroupResponseResults `json:"results,omitempty"`
}

// NewUserGroupResponse instantiates a new UserGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupResponse() *UserGroupResponse {
	this := UserGroupResponse{}
	return &this
}

// NewUserGroupResponseWithDefaults instantiates a new UserGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupResponseWithDefaults() *UserGroupResponse {
	this := UserGroupResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UserGroupResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UserGroupResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UserGroupResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UserGroupResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UserGroupResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UserGroupResponse) SetMessage(v string) {
	o.Message = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *UserGroupResponse) GetResults() UserGroupResponseResults {
	if o == nil || IsNil(o.Results) {
		var ret UserGroupResponseResults
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupResponse) GetResultsOk() (*UserGroupResponseResults, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *UserGroupResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given UserGroupResponseResults and assigns it to the Results field.
func (o *UserGroupResponse) SetResults(v UserGroupResponseResults) {
	o.Results = &v
}

func (o UserGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGroupResponse) ToMap() (map[string]interface{}, error) {
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

type NullableUserGroupResponse struct {
	value *UserGroupResponse
	isSet bool
}

func (v NullableUserGroupResponse) Get() *UserGroupResponse {
	return v.value
}

func (v *NullableUserGroupResponse) Set(val *UserGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupResponse(val *UserGroupResponse) *NullableUserGroupResponse {
	return &NullableUserGroupResponse{value: val, isSet: true}
}

func (v NullableUserGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



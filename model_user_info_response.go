/*
WhatsApp API MultiDevice

This API is used for sending whatsapp via API

API version: 6.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SdkWhatsappWebMultiDevice

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInfoResponse{}

// UserInfoResponse struct for UserInfoResponse
type UserInfoResponse struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Results UserInfoResponseResults `json:"results"`
}

type _UserInfoResponse UserInfoResponse

// NewUserInfoResponse instantiates a new UserInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInfoResponse(code string, message string, results UserInfoResponseResults) *UserInfoResponse {
	this := UserInfoResponse{}
	this.Code = code
	this.Message = message
	this.Results = results
	return &this
}

// NewUserInfoResponseWithDefaults instantiates a new UserInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInfoResponseWithDefaults() *UserInfoResponse {
	this := UserInfoResponse{}
	return &this
}

// GetCode returns the Code field value
func (o *UserInfoResponse) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *UserInfoResponse) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *UserInfoResponse) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *UserInfoResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *UserInfoResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *UserInfoResponse) SetMessage(v string) {
	o.Message = v
}

// GetResults returns the Results field value
func (o *UserInfoResponse) GetResults() UserInfoResponseResults {
	if o == nil {
		var ret UserInfoResponseResults
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *UserInfoResponse) GetResultsOk() (*UserInfoResponseResults, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *UserInfoResponse) SetResults(v UserInfoResponseResults) {
	o.Results = v
}

func (o UserInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *UserInfoResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
		"results",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUserInfoResponse := _UserInfoResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserInfoResponse)

	if err != nil {
		return err
	}

	*o = UserInfoResponse(varUserInfoResponse)

	return err
}

type NullableUserInfoResponse struct {
	value *UserInfoResponse
	isSet bool
}

func (v NullableUserInfoResponse) Get() *UserInfoResponse {
	return v.value
}

func (v *NullableUserInfoResponse) Set(val *UserInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInfoResponse(val *UserInfoResponse) *NullableUserInfoResponse {
	return &NullableUserInfoResponse{value: val, isSet: true}
}

func (v NullableUserInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the BusinessProfileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BusinessProfileResponse{}

// BusinessProfileResponse struct for BusinessProfileResponse
type BusinessProfileResponse struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Results *BusinessProfileResponseResults `json:"results,omitempty"`
}

// NewBusinessProfileResponse instantiates a new BusinessProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessProfileResponse() *BusinessProfileResponse {
	this := BusinessProfileResponse{}
	return &this
}

// NewBusinessProfileResponseWithDefaults instantiates a new BusinessProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessProfileResponseWithDefaults() *BusinessProfileResponse {
	this := BusinessProfileResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *BusinessProfileResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessProfileResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *BusinessProfileResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *BusinessProfileResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BusinessProfileResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessProfileResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BusinessProfileResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BusinessProfileResponse) SetMessage(v string) {
	o.Message = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *BusinessProfileResponse) GetResults() BusinessProfileResponseResults {
	if o == nil || IsNil(o.Results) {
		var ret BusinessProfileResponseResults
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessProfileResponse) GetResultsOk() (*BusinessProfileResponseResults, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *BusinessProfileResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given BusinessProfileResponseResults and assigns it to the Results field.
func (o *BusinessProfileResponse) SetResults(v BusinessProfileResponseResults) {
	o.Results = &v
}

func (o BusinessProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessProfileResponse) ToMap() (map[string]interface{}, error) {
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

type NullableBusinessProfileResponse struct {
	value *BusinessProfileResponse
	isSet bool
}

func (v NullableBusinessProfileResponse) Get() *BusinessProfileResponse {
	return v.value
}

func (v *NullableBusinessProfileResponse) Set(val *BusinessProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessProfileResponse(val *BusinessProfileResponse) *NullableBusinessProfileResponse {
	return &NullableBusinessProfileResponse{value: val, isSet: true}
}

func (v NullableBusinessProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



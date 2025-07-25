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

// checks if the ErrorUnauthorized type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorUnauthorized{}

// ErrorUnauthorized struct for ErrorUnauthorized
type ErrorUnauthorized struct {
	// HTTP Status Code
	Code *string `json:"code,omitempty"`
	// Detail error message
	Message *string `json:"message,omitempty"`
	// additional data
	Results map[string]interface{} `json:"results,omitempty"`
}

// NewErrorUnauthorized instantiates a new ErrorUnauthorized object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorUnauthorized() *ErrorUnauthorized {
	this := ErrorUnauthorized{}
	return &this
}

// NewErrorUnauthorizedWithDefaults instantiates a new ErrorUnauthorized object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorUnauthorizedWithDefaults() *ErrorUnauthorized {
	this := ErrorUnauthorized{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ErrorUnauthorized) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorUnauthorized) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ErrorUnauthorized) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ErrorUnauthorized) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorUnauthorized) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorUnauthorized) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorUnauthorized) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorUnauthorized) SetMessage(v string) {
	o.Message = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ErrorUnauthorized) GetResults() map[string]interface{} {
	if o == nil || IsNil(o.Results) {
		var ret map[string]interface{}
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorUnauthorized) GetResultsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Results) {
		return map[string]interface{}{}, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ErrorUnauthorized) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given map[string]interface{} and assigns it to the Results field.
func (o *ErrorUnauthorized) SetResults(v map[string]interface{}) {
	o.Results = v
}

func (o ErrorUnauthorized) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorUnauthorized) ToMap() (map[string]interface{}, error) {
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

type NullableErrorUnauthorized struct {
	value *ErrorUnauthorized
	isSet bool
}

func (v NullableErrorUnauthorized) Get() *ErrorUnauthorized {
	return v.value
}

func (v *NullableErrorUnauthorized) Set(val *ErrorUnauthorized) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorUnauthorized) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorUnauthorized) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorUnauthorized(val *ErrorUnauthorized) *NullableErrorUnauthorized {
	return &NullableErrorUnauthorized{value: val, isSet: true}
}

func (v NullableErrorUnauthorized) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorUnauthorized) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



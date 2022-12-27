/*
WhatsApp API MultiDevice

This API is used for sending whatsapp via API

API version: 3.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SendResponse struct for SendResponse
type SendResponse struct {
	Message *string `json:"message,omitempty"`
	Results *SendResponseResults `json:"results,omitempty"`
}

// NewSendResponse instantiates a new SendResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendResponse() *SendResponse {
	this := SendResponse{}
	return &this
}

// NewSendResponseWithDefaults instantiates a new SendResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendResponseWithDefaults() *SendResponse {
	this := SendResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SendResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SendResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SendResponse) SetMessage(v string) {
	o.Message = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *SendResponse) GetResults() SendResponseResults {
	if o == nil || isNil(o.Results) {
		var ret SendResponseResults
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendResponse) GetResultsOk() (*SendResponseResults, bool) {
	if o == nil || isNil(o.Results) {
    return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *SendResponse) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given SendResponseResults and assigns it to the Results field.
func (o *SendResponse) SetResults(v SendResponseResults) {
	o.Results = &v
}

func (o SendResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableSendResponse struct {
	value *SendResponse
	isSet bool
}

func (v NullableSendResponse) Get() *SendResponse {
	return v.value
}

func (v *NullableSendResponse) Set(val *SendResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSendResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSendResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendResponse(val *SendResponse) *NullableSendResponse {
	return &NullableSendResponse{value: val, isSet: true}
}

func (v NullableSendResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



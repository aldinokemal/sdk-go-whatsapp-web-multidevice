/*
WhatsApp API MultiDevice

This API is used for sending whatsapp via API

API version: 3.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DeviceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceResponse{}

// DeviceResponse struct for DeviceResponse
type DeviceResponse struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Results []DeviceResponseResultsInner `json:"results,omitempty"`
}

// NewDeviceResponse instantiates a new DeviceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceResponse() *DeviceResponse {
	this := DeviceResponse{}
	return &this
}

// NewDeviceResponseWithDefaults instantiates a new DeviceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceResponseWithDefaults() *DeviceResponse {
	this := DeviceResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *DeviceResponse) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceResponse) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *DeviceResponse) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *DeviceResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DeviceResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DeviceResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DeviceResponse) SetMessage(v string) {
	o.Message = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *DeviceResponse) GetResults() []DeviceResponseResultsInner {
	if o == nil || isNil(o.Results) {
		var ret []DeviceResponseResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceResponse) GetResultsOk() ([]DeviceResponseResultsInner, bool) {
	if o == nil || isNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *DeviceResponse) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DeviceResponseResultsInner and assigns it to the Results field.
func (o *DeviceResponse) SetResults(v []DeviceResponseResultsInner) {
	o.Results = v
}

func (o DeviceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableDeviceResponse struct {
	value *DeviceResponse
	isSet bool
}

func (v NullableDeviceResponse) Get() *DeviceResponse {
	return v.value
}

func (v *NullableDeviceResponse) Set(val *DeviceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceResponse(val *DeviceResponse) *NullableDeviceResponse {
	return &NullableDeviceResponse{value: val, isSet: true}
}

func (v NullableDeviceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



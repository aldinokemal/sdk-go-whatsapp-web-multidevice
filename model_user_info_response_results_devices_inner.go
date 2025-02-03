/*
WhatsApp API MultiDevice

This API is used for sending whatsapp via API

API version: 5.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SdkWhatsappWebMultiDevice

import (
	"encoding/json"
)

// checks if the UserInfoResponseResultsDevicesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInfoResponseResultsDevicesInner{}

// UserInfoResponseResultsDevicesInner struct for UserInfoResponseResultsDevicesInner
type UserInfoResponseResultsDevicesInner struct {
	User *string `json:"User,omitempty"`
	Agent *int32 `json:"Agent,omitempty"`
	Device *string `json:"Device,omitempty"`
	Server *string `json:"Server,omitempty"`
	AD *bool `json:"AD,omitempty"`
}

// NewUserInfoResponseResultsDevicesInner instantiates a new UserInfoResponseResultsDevicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInfoResponseResultsDevicesInner() *UserInfoResponseResultsDevicesInner {
	this := UserInfoResponseResultsDevicesInner{}
	return &this
}

// NewUserInfoResponseResultsDevicesInnerWithDefaults instantiates a new UserInfoResponseResultsDevicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInfoResponseResultsDevicesInnerWithDefaults() *UserInfoResponseResultsDevicesInner {
	this := UserInfoResponseResultsDevicesInner{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UserInfoResponseResultsDevicesInner) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfoResponseResultsDevicesInner) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UserInfoResponseResultsDevicesInner) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *UserInfoResponseResultsDevicesInner) SetUser(v string) {
	o.User = &v
}

// GetAgent returns the Agent field value if set, zero value otherwise.
func (o *UserInfoResponseResultsDevicesInner) GetAgent() int32 {
	if o == nil || IsNil(o.Agent) {
		var ret int32
		return ret
	}
	return *o.Agent
}

// GetAgentOk returns a tuple with the Agent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfoResponseResultsDevicesInner) GetAgentOk() (*int32, bool) {
	if o == nil || IsNil(o.Agent) {
		return nil, false
	}
	return o.Agent, true
}

// HasAgent returns a boolean if a field has been set.
func (o *UserInfoResponseResultsDevicesInner) HasAgent() bool {
	if o != nil && !IsNil(o.Agent) {
		return true
	}

	return false
}

// SetAgent gets a reference to the given int32 and assigns it to the Agent field.
func (o *UserInfoResponseResultsDevicesInner) SetAgent(v int32) {
	o.Agent = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *UserInfoResponseResultsDevicesInner) GetDevice() string {
	if o == nil || IsNil(o.Device) {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfoResponseResultsDevicesInner) GetDeviceOk() (*string, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *UserInfoResponseResultsDevicesInner) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *UserInfoResponseResultsDevicesInner) SetDevice(v string) {
	o.Device = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *UserInfoResponseResultsDevicesInner) GetServer() string {
	if o == nil || IsNil(o.Server) {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfoResponseResultsDevicesInner) GetServerOk() (*string, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *UserInfoResponseResultsDevicesInner) HasServer() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *UserInfoResponseResultsDevicesInner) SetServer(v string) {
	o.Server = &v
}

// GetAD returns the AD field value if set, zero value otherwise.
func (o *UserInfoResponseResultsDevicesInner) GetAD() bool {
	if o == nil || IsNil(o.AD) {
		var ret bool
		return ret
	}
	return *o.AD
}

// GetADOk returns a tuple with the AD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfoResponseResultsDevicesInner) GetADOk() (*bool, bool) {
	if o == nil || IsNil(o.AD) {
		return nil, false
	}
	return o.AD, true
}

// HasAD returns a boolean if a field has been set.
func (o *UserInfoResponseResultsDevicesInner) HasAD() bool {
	if o != nil && !IsNil(o.AD) {
		return true
	}

	return false
}

// SetAD gets a reference to the given bool and assigns it to the AD field.
func (o *UserInfoResponseResultsDevicesInner) SetAD(v bool) {
	o.AD = &v
}

func (o UserInfoResponseResultsDevicesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInfoResponseResultsDevicesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["User"] = o.User
	}
	if !IsNil(o.Agent) {
		toSerialize["Agent"] = o.Agent
	}
	if !IsNil(o.Device) {
		toSerialize["Device"] = o.Device
	}
	if !IsNil(o.Server) {
		toSerialize["Server"] = o.Server
	}
	if !IsNil(o.AD) {
		toSerialize["AD"] = o.AD
	}
	return toSerialize, nil
}

type NullableUserInfoResponseResultsDevicesInner struct {
	value *UserInfoResponseResultsDevicesInner
	isSet bool
}

func (v NullableUserInfoResponseResultsDevicesInner) Get() *UserInfoResponseResultsDevicesInner {
	return v.value
}

func (v *NullableUserInfoResponseResultsDevicesInner) Set(val *UserInfoResponseResultsDevicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInfoResponseResultsDevicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInfoResponseResultsDevicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInfoResponseResultsDevicesInner(val *UserInfoResponseResultsDevicesInner) *NullableUserInfoResponseResultsDevicesInner {
	return &NullableUserInfoResponseResultsDevicesInner{value: val, isSet: true}
}

func (v NullableUserInfoResponseResultsDevicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInfoResponseResultsDevicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



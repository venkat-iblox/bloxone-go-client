/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.   

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns_config

import (
	"encoding/json"
)

// checks if the ConfigCreateAuthNSGResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigCreateAuthNSGResponse{}

// ConfigCreateAuthNSGResponse The AuthNSG object create response format.
type ConfigCreateAuthNSGResponse struct {
	Result *ConfigAuthNSG `json:"result,omitempty"`
}

// NewConfigCreateAuthNSGResponse instantiates a new ConfigCreateAuthNSGResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigCreateAuthNSGResponse() *ConfigCreateAuthNSGResponse {
	this := ConfigCreateAuthNSGResponse{}
	return &this
}

// NewConfigCreateAuthNSGResponseWithDefaults instantiates a new ConfigCreateAuthNSGResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigCreateAuthNSGResponseWithDefaults() *ConfigCreateAuthNSGResponse {
	this := ConfigCreateAuthNSGResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ConfigCreateAuthNSGResponse) GetResult() ConfigAuthNSG {
	if o == nil || IsNil(o.Result) {
		var ret ConfigAuthNSG
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigCreateAuthNSGResponse) GetResultOk() (*ConfigAuthNSG, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ConfigCreateAuthNSGResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ConfigAuthNSG and assigns it to the Result field.
func (o *ConfigCreateAuthNSGResponse) SetResult(v ConfigAuthNSG) {
	o.Result = &v
}

func (o ConfigCreateAuthNSGResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigCreateAuthNSGResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableConfigCreateAuthNSGResponse struct {
	value *ConfigCreateAuthNSGResponse
	isSet bool
}

func (v NullableConfigCreateAuthNSGResponse) Get() *ConfigCreateAuthNSGResponse {
	return v.value
}

func (v *NullableConfigCreateAuthNSGResponse) Set(val *ConfigCreateAuthNSGResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigCreateAuthNSGResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigCreateAuthNSGResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigCreateAuthNSGResponse(val *ConfigCreateAuthNSGResponse) *NullableConfigCreateAuthNSGResponse {
	return &NullableConfigCreateAuthNSGResponse{value: val, isSet: true}
}

func (v NullableConfigCreateAuthNSGResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigCreateAuthNSGResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



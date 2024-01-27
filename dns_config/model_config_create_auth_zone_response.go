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

// checks if the ConfigCreateAuthZoneResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigCreateAuthZoneResponse{}

// ConfigCreateAuthZoneResponse The Authoritative Zone object create response format.
type ConfigCreateAuthZoneResponse struct {
	Result *ConfigAuthZone `json:"result,omitempty"`
}

// NewConfigCreateAuthZoneResponse instantiates a new ConfigCreateAuthZoneResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigCreateAuthZoneResponse() *ConfigCreateAuthZoneResponse {
	this := ConfigCreateAuthZoneResponse{}
	return &this
}

// NewConfigCreateAuthZoneResponseWithDefaults instantiates a new ConfigCreateAuthZoneResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigCreateAuthZoneResponseWithDefaults() *ConfigCreateAuthZoneResponse {
	this := ConfigCreateAuthZoneResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ConfigCreateAuthZoneResponse) GetResult() ConfigAuthZone {
	if o == nil || IsNil(o.Result) {
		var ret ConfigAuthZone
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigCreateAuthZoneResponse) GetResultOk() (*ConfigAuthZone, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ConfigCreateAuthZoneResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ConfigAuthZone and assigns it to the Result field.
func (o *ConfigCreateAuthZoneResponse) SetResult(v ConfigAuthZone) {
	o.Result = &v
}

func (o ConfigCreateAuthZoneResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigCreateAuthZoneResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableConfigCreateAuthZoneResponse struct {
	value *ConfigCreateAuthZoneResponse
	isSet bool
}

func (v NullableConfigCreateAuthZoneResponse) Get() *ConfigCreateAuthZoneResponse {
	return v.value
}

func (v *NullableConfigCreateAuthZoneResponse) Set(val *ConfigCreateAuthZoneResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigCreateAuthZoneResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigCreateAuthZoneResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigCreateAuthZoneResponse(val *ConfigCreateAuthZoneResponse) *NullableConfigCreateAuthZoneResponse {
	return &NullableConfigCreateAuthZoneResponse{value: val, isSet: true}
}

func (v NullableConfigCreateAuthZoneResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigCreateAuthZoneResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



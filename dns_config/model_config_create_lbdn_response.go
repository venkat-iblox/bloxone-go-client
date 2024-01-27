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

// checks if the ConfigCreateLBDNResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigCreateLBDNResponse{}

// ConfigCreateLBDNResponse The __LBDN__ object create response format.
type ConfigCreateLBDNResponse struct {
	Result *ConfigLBDN `json:"result,omitempty"`
}

// NewConfigCreateLBDNResponse instantiates a new ConfigCreateLBDNResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigCreateLBDNResponse() *ConfigCreateLBDNResponse {
	this := ConfigCreateLBDNResponse{}
	return &this
}

// NewConfigCreateLBDNResponseWithDefaults instantiates a new ConfigCreateLBDNResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigCreateLBDNResponseWithDefaults() *ConfigCreateLBDNResponse {
	this := ConfigCreateLBDNResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ConfigCreateLBDNResponse) GetResult() ConfigLBDN {
	if o == nil || IsNil(o.Result) {
		var ret ConfigLBDN
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigCreateLBDNResponse) GetResultOk() (*ConfigLBDN, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ConfigCreateLBDNResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ConfigLBDN and assigns it to the Result field.
func (o *ConfigCreateLBDNResponse) SetResult(v ConfigLBDN) {
	o.Result = &v
}

func (o ConfigCreateLBDNResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigCreateLBDNResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableConfigCreateLBDNResponse struct {
	value *ConfigCreateLBDNResponse
	isSet bool
}

func (v NullableConfigCreateLBDNResponse) Get() *ConfigCreateLBDNResponse {
	return v.value
}

func (v *NullableConfigCreateLBDNResponse) Set(val *ConfigCreateLBDNResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigCreateLBDNResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigCreateLBDNResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigCreateLBDNResponse(val *ConfigCreateLBDNResponse) *NullableConfigCreateLBDNResponse {
	return &NullableConfigCreateLBDNResponse{value: val, isSet: true}
}

func (v NullableConfigCreateLBDNResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigCreateLBDNResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



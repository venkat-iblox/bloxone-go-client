/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the IpamsvcIgnoreItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcIgnoreItem{}

// IpamsvcIgnoreItem An Ignore Item object (_dhcp/ignore_item_) represents an item in a DHCP ignore list.
type IpamsvcIgnoreItem struct {
	// Type of ignore matching: client to ignore by client identifier (client hex or client text) or hardware to ignore by hardware identifier (MAC address). It can have one of the following values:  * _client_hex_,  * _client_text_,  * _hardware_.
	Type string `json:"type"`
	// Value to match.
	Value string `json:"value"`
}

// NewIpamsvcIgnoreItem instantiates a new IpamsvcIgnoreItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcIgnoreItem(type_ string, value string) *IpamsvcIgnoreItem {
	this := IpamsvcIgnoreItem{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewIpamsvcIgnoreItemWithDefaults instantiates a new IpamsvcIgnoreItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcIgnoreItemWithDefaults() *IpamsvcIgnoreItem {
	this := IpamsvcIgnoreItem{}
	return &this
}

// GetType returns the Type field value
func (o *IpamsvcIgnoreItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IpamsvcIgnoreItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IpamsvcIgnoreItem) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *IpamsvcIgnoreItem) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *IpamsvcIgnoreItem) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *IpamsvcIgnoreItem) SetValue(v string) {
	o.Value = v
}

func (o IpamsvcIgnoreItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcIgnoreItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableIpamsvcIgnoreItem struct {
	value *IpamsvcIgnoreItem
	isSet bool
}

func (v NullableIpamsvcIgnoreItem) Get() *IpamsvcIgnoreItem {
	return v.value
}

func (v *NullableIpamsvcIgnoreItem) Set(val *IpamsvcIgnoreItem) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcIgnoreItem) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcIgnoreItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcIgnoreItem(val *IpamsvcIgnoreItem) *NullableIpamsvcIgnoreItem {
	return &NullableIpamsvcIgnoreItem{value: val, isSet: true}
}

func (v NullableIpamsvcIgnoreItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcIgnoreItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the IpamsvcCreateOptionGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcCreateOptionGroupResponse{}

// IpamsvcCreateOptionGroupResponse The response format to create the __OptionGroup__ object.
type IpamsvcCreateOptionGroupResponse struct {
	Result *IpamsvcOptionGroup `json:"result,omitempty"`
}

// NewIpamsvcCreateOptionGroupResponse instantiates a new IpamsvcCreateOptionGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcCreateOptionGroupResponse() *IpamsvcCreateOptionGroupResponse {
	this := IpamsvcCreateOptionGroupResponse{}
	return &this
}

// NewIpamsvcCreateOptionGroupResponseWithDefaults instantiates a new IpamsvcCreateOptionGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcCreateOptionGroupResponseWithDefaults() *IpamsvcCreateOptionGroupResponse {
	this := IpamsvcCreateOptionGroupResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *IpamsvcCreateOptionGroupResponse) GetResult() IpamsvcOptionGroup {
	if o == nil || IsNil(o.Result) {
		var ret IpamsvcOptionGroup
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcCreateOptionGroupResponse) GetResultOk() (*IpamsvcOptionGroup, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *IpamsvcCreateOptionGroupResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given IpamsvcOptionGroup and assigns it to the Result field.
func (o *IpamsvcCreateOptionGroupResponse) SetResult(v IpamsvcOptionGroup) {
	o.Result = &v
}

func (o IpamsvcCreateOptionGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcCreateOptionGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableIpamsvcCreateOptionGroupResponse struct {
	value *IpamsvcCreateOptionGroupResponse
	isSet bool
}

func (v NullableIpamsvcCreateOptionGroupResponse) Get() *IpamsvcCreateOptionGroupResponse {
	return v.value
}

func (v *NullableIpamsvcCreateOptionGroupResponse) Set(val *IpamsvcCreateOptionGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcCreateOptionGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcCreateOptionGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcCreateOptionGroupResponse(val *IpamsvcCreateOptionGroupResponse) *NullableIpamsvcCreateOptionGroupResponse {
	return &NullableIpamsvcCreateOptionGroupResponse{value: val, isSet: true}
}

func (v NullableIpamsvcCreateOptionGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcCreateOptionGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



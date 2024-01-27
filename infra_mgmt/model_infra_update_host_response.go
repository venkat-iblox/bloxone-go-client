/*
Infrastructure Management API

The **Infrastructure Management API** provides a RESTful interface to manage Infrastructure Hosts and Services objects.  The following is a list of the different Services and their string types (the string types are to be used with the APIs for the `service_type` field):  | Service name | Service type |   | ------ | ------ |   | Access Authentication | authn |   | Anycast | anycast |   | Data Connector | cdc |   | DHCP | dhcp |   | DNS | dns |   | DNS Forwarding Proxy | dfp |   | NIOS Grid Connector | orpheus |   | MS AD Sync | msad |   | NTP | ntp |   | BGP | bgp |   | RIP | rip |   | OSPF | ospf |    ---   ### Hosts API  The Hosts API is used to manage the Infrastructure Host resources. These include various operations related to hosts such as viewing, creating, updating, replacing, disconnecting, and deleting Hosts. Management of Hosts is done from the Cloud Services Portal (CSP) by navigating to the Manage -> Infrastructure -> Hosts tab.  ---   ### Services API  The Services API is used to manage the Infrastructure Service resources (a.k.a. BloxOne applications). These include various operations related to hosts such as viewing, creating, updating, starting/stopping, configuring, and deleting Services. Management of Services is done from the Cloud Services Portal (CSP) by navigating to the Manage -> Infrastructure -> Services tab.  ---   ### Detail APIs  The Detail APIs are read-only APIs used to list all the Infrastructure resources (Hosts and Services). Each resource record returned also contains information about its other associated resources and the status information for itself and the associated resource(s) (i.e., Host/Service status).  ---   

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infra_mgmt

import (
	"encoding/json"
)

// checks if the InfraUpdateHostResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfraUpdateHostResponse{}

// InfraUpdateHostResponse struct for InfraUpdateHostResponse
type InfraUpdateHostResponse struct {
	Result *InfraHost `json:"result,omitempty"`
}

// NewInfraUpdateHostResponse instantiates a new InfraUpdateHostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfraUpdateHostResponse() *InfraUpdateHostResponse {
	this := InfraUpdateHostResponse{}
	return &this
}

// NewInfraUpdateHostResponseWithDefaults instantiates a new InfraUpdateHostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfraUpdateHostResponseWithDefaults() *InfraUpdateHostResponse {
	this := InfraUpdateHostResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *InfraUpdateHostResponse) GetResult() InfraHost {
	if o == nil || IsNil(o.Result) {
		var ret InfraHost
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraUpdateHostResponse) GetResultOk() (*InfraHost, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *InfraUpdateHostResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given InfraHost and assigns it to the Result field.
func (o *InfraUpdateHostResponse) SetResult(v InfraHost) {
	o.Result = &v
}

func (o InfraUpdateHostResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfraUpdateHostResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableInfraUpdateHostResponse struct {
	value *InfraUpdateHostResponse
	isSet bool
}

func (v NullableInfraUpdateHostResponse) Get() *InfraUpdateHostResponse {
	return v.value
}

func (v *NullableInfraUpdateHostResponse) Set(val *InfraUpdateHostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInfraUpdateHostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInfraUpdateHostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfraUpdateHostResponse(val *InfraUpdateHostResponse) *NullableInfraUpdateHostResponse {
	return &NullableInfraUpdateHostResponse{value: val, isSet: true}
}

func (v NullableInfraUpdateHostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfraUpdateHostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the InfraListDetailServicesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfraListDetailServicesResponse{}

// InfraListDetailServicesResponse struct for InfraListDetailServicesResponse
type InfraListDetailServicesResponse struct {
	Page *ApiPageInfo `json:"page,omitempty"`
	Results []InfraDetailService `json:"results,omitempty"`
}

// NewInfraListDetailServicesResponse instantiates a new InfraListDetailServicesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfraListDetailServicesResponse() *InfraListDetailServicesResponse {
	this := InfraListDetailServicesResponse{}
	return &this
}

// NewInfraListDetailServicesResponseWithDefaults instantiates a new InfraListDetailServicesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfraListDetailServicesResponseWithDefaults() *InfraListDetailServicesResponse {
	this := InfraListDetailServicesResponse{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *InfraListDetailServicesResponse) GetPage() ApiPageInfo {
	if o == nil || IsNil(o.Page) {
		var ret ApiPageInfo
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraListDetailServicesResponse) GetPageOk() (*ApiPageInfo, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *InfraListDetailServicesResponse) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given ApiPageInfo and assigns it to the Page field.
func (o *InfraListDetailServicesResponse) SetPage(v ApiPageInfo) {
	o.Page = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *InfraListDetailServicesResponse) GetResults() []InfraDetailService {
	if o == nil || IsNil(o.Results) {
		var ret []InfraDetailService
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraListDetailServicesResponse) GetResultsOk() ([]InfraDetailService, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *InfraListDetailServicesResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []InfraDetailService and assigns it to the Results field.
func (o *InfraListDetailServicesResponse) SetResults(v []InfraDetailService) {
	o.Results = v
}

func (o InfraListDetailServicesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfraListDetailServicesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableInfraListDetailServicesResponse struct {
	value *InfraListDetailServicesResponse
	isSet bool
}

func (v NullableInfraListDetailServicesResponse) Get() *InfraListDetailServicesResponse {
	return v.value
}

func (v *NullableInfraListDetailServicesResponse) Set(val *InfraListDetailServicesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInfraListDetailServicesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInfraListDetailServicesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfraListDetailServicesResponse(val *InfraListDetailServicesResponse) *NullableInfraListDetailServicesResponse {
	return &NullableInfraListDetailServicesResponse{value: val, isSet: true}
}

func (v NullableInfraListDetailServicesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfraListDetailServicesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



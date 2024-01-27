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

// checks if the Inheritance2AssignedHost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Inheritance2AssignedHost{}

// Inheritance2AssignedHost _ddi/assigned_host_ represents a BloxOne DDI host assigned to an object.
type Inheritance2AssignedHost struct {
	// The human-readable display name for the host referred to by _ophid_.
	DisplayName *string `json:"display_name,omitempty"`
	// The resource identifier.
	Host *string `json:"host,omitempty"`
	// The on-prem host ID.
	Ophid *string `json:"ophid,omitempty"`
}

// NewInheritance2AssignedHost instantiates a new Inheritance2AssignedHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInheritance2AssignedHost() *Inheritance2AssignedHost {
	this := Inheritance2AssignedHost{}
	return &this
}

// NewInheritance2AssignedHostWithDefaults instantiates a new Inheritance2AssignedHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInheritance2AssignedHostWithDefaults() *Inheritance2AssignedHost {
	this := Inheritance2AssignedHost{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Inheritance2AssignedHost) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inheritance2AssignedHost) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Inheritance2AssignedHost) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Inheritance2AssignedHost) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *Inheritance2AssignedHost) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inheritance2AssignedHost) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *Inheritance2AssignedHost) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *Inheritance2AssignedHost) SetHost(v string) {
	o.Host = &v
}

// GetOphid returns the Ophid field value if set, zero value otherwise.
func (o *Inheritance2AssignedHost) GetOphid() string {
	if o == nil || IsNil(o.Ophid) {
		var ret string
		return ret
	}
	return *o.Ophid
}

// GetOphidOk returns a tuple with the Ophid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inheritance2AssignedHost) GetOphidOk() (*string, bool) {
	if o == nil || IsNil(o.Ophid) {
		return nil, false
	}
	return o.Ophid, true
}

// HasOphid returns a boolean if a field has been set.
func (o *Inheritance2AssignedHost) HasOphid() bool {
	if o != nil && !IsNil(o.Ophid) {
		return true
	}

	return false
}

// SetOphid gets a reference to the given string and assigns it to the Ophid field.
func (o *Inheritance2AssignedHost) SetOphid(v string) {
	o.Ophid = &v
}

func (o Inheritance2AssignedHost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Inheritance2AssignedHost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Ophid) {
		toSerialize["ophid"] = o.Ophid
	}
	return toSerialize, nil
}

type NullableInheritance2AssignedHost struct {
	value *Inheritance2AssignedHost
	isSet bool
}

func (v NullableInheritance2AssignedHost) Get() *Inheritance2AssignedHost {
	return v.value
}

func (v *NullableInheritance2AssignedHost) Set(val *Inheritance2AssignedHost) {
	v.value = val
	v.isSet = true
}

func (v NullableInheritance2AssignedHost) IsSet() bool {
	return v.isSet
}

func (v *NullableInheritance2AssignedHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInheritance2AssignedHost(val *Inheritance2AssignedHost) *NullableInheritance2AssignedHost {
	return &NullableInheritance2AssignedHost{value: val, isSet: true}
}

func (v NullableInheritance2AssignedHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInheritance2AssignedHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



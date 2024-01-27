/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
	"time"
)

// checks if the IpamsvcHAGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcHAGroup{}

// IpamsvcHAGroup An __HAGroup__ object (_dhcp/ha_group_) represents on-prem hosts that can serve the same leases for HA.
type IpamsvcHAGroup struct {
	// The resource identifier.
	AnycastConfigId *string `json:"anycast_config_id,omitempty"`
	// The description for the HA group. May contain 0 to 1024 characters. Can include UTF-8.
	Comment *string `json:"comment,omitempty"`
	// Time when the object has been created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The list of hosts.
	Hosts []IpamsvcHAGroupHost `json:"hosts"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The resource identifier.
	IpSpace *string `json:"ip_space,omitempty"`
	// The mode of the HA group.  Valid values are: * _active-active_: Both on-prem hosts remain active. * _active-passive_: One on-prem host remains active and one remains passive. When the active on-prem host is down, the passive on-prem host takes over. * _advanced-active-passive_: One on-prem host may be part of multiple HA groups. When the active on-prem host is down, the passive on-prem host takes over.
	Mode *string `json:"mode,omitempty"`
	// The name of the HA group. Must contain 1 to 256 characters. Can include UTF-8.
	Name string `json:"name"`
	// Status of the HA group. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request.
	Status *string `json:"status,omitempty"`
	// The tags for the HA group.
	Tags map[string]interface{} `json:"tags,omitempty"`
	// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewIpamsvcHAGroup instantiates a new IpamsvcHAGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcHAGroup(hosts []IpamsvcHAGroupHost, name string) *IpamsvcHAGroup {
	this := IpamsvcHAGroup{}
	this.Hosts = hosts
	this.Name = name
	return &this
}

// NewIpamsvcHAGroupWithDefaults instantiates a new IpamsvcHAGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcHAGroupWithDefaults() *IpamsvcHAGroup {
	this := IpamsvcHAGroup{}
	return &this
}

// GetAnycastConfigId returns the AnycastConfigId field value if set, zero value otherwise.
func (o *IpamsvcHAGroup) GetAnycastConfigId() string {
	if o == nil || IsNil(o.AnycastConfigId) {
		var ret string
		return ret
	}
	return *o.AnycastConfigId
}

// GetAnycastConfigIdOk returns a tuple with the AnycastConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHAGroup) GetAnycastConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.AnycastConfigId) {
		return nil, false
	}
	return o.AnycastConfigId, true
}

// HasAnycastConfigId returns a boolean if a field has been set.
func (o *IpamsvcHAGroup) HasAnycastConfigId() bool {
	if o != nil && !IsNil(o.AnycastConfigId) {
		return true
	}

	return false
}

// SetAnycastConfigId gets a reference to the given string and assigns it to the AnycastConfigId field.
func (o *IpamsvcHAGroup) SetAnycastConfigId(v string) {
	o.AnycastConfigId = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *IpamsvcHAGroup) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHAGroup) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *IpamsvcHAGroup) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *IpamsvcHAGroup) SetComment(v string) {
	o.Comment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IpamsvcHAGroup) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHAGroup) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IpamsvcHAGroup) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *IpamsvcHAGroup) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHosts returns the Hosts field value
func (o *IpamsvcHAGroup) GetHosts() []IpamsvcHAGroupHost {
	if o == nil {
		var ret []IpamsvcHAGroupHost
		return ret
	}

	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value
// and a boolean to check if the value has been set.
func (o *IpamsvcHAGroup) GetHostsOk() ([]IpamsvcHAGroupHost, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hosts, true
}

// SetHosts sets field value
func (o *IpamsvcHAGroup) SetHosts(v []IpamsvcHAGroupHost) {
	o.Hosts = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IpamsvcHAGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHAGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IpamsvcHAGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IpamsvcHAGroup) SetId(v string) {
	o.Id = &v
}

// GetIpSpace returns the IpSpace field value if set, zero value otherwise.
func (o *IpamsvcHAGroup) GetIpSpace() string {
	if o == nil || IsNil(o.IpSpace) {
		var ret string
		return ret
	}
	return *o.IpSpace
}

// GetIpSpaceOk returns a tuple with the IpSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHAGroup) GetIpSpaceOk() (*string, bool) {
	if o == nil || IsNil(o.IpSpace) {
		return nil, false
	}
	return o.IpSpace, true
}

// HasIpSpace returns a boolean if a field has been set.
func (o *IpamsvcHAGroup) HasIpSpace() bool {
	if o != nil && !IsNil(o.IpSpace) {
		return true
	}

	return false
}

// SetIpSpace gets a reference to the given string and assigns it to the IpSpace field.
func (o *IpamsvcHAGroup) SetIpSpace(v string) {
	o.IpSpace = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *IpamsvcHAGroup) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHAGroup) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *IpamsvcHAGroup) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *IpamsvcHAGroup) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value
func (o *IpamsvcHAGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IpamsvcHAGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IpamsvcHAGroup) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IpamsvcHAGroup) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHAGroup) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IpamsvcHAGroup) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IpamsvcHAGroup) SetStatus(v string) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IpamsvcHAGroup) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHAGroup) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IpamsvcHAGroup) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *IpamsvcHAGroup) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IpamsvcHAGroup) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcHAGroup) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IpamsvcHAGroup) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *IpamsvcHAGroup) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o IpamsvcHAGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcHAGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnycastConfigId) {
		toSerialize["anycast_config_id"] = o.AnycastConfigId
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	toSerialize["hosts"] = o.Hosts
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IpSpace) {
		toSerialize["ip_space"] = o.IpSpace
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableIpamsvcHAGroup struct {
	value *IpamsvcHAGroup
	isSet bool
}

func (v NullableIpamsvcHAGroup) Get() *IpamsvcHAGroup {
	return v.value
}

func (v *NullableIpamsvcHAGroup) Set(val *IpamsvcHAGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcHAGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcHAGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcHAGroup(val *IpamsvcHAGroup) *NullableIpamsvcHAGroup {
	return &NullableIpamsvcHAGroup{value: val, isSet: true}
}

func (v NullableIpamsvcHAGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcHAGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



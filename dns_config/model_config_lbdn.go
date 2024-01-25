/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns_config

import (
	"encoding/json"
	"fmt"
)

// checks if the ConfigLBDN type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigLBDN{}

// ConfigLBDN A LBDN (_dtc/lbdn_) represents a load-balanced domain name
type ConfigLBDN struct {
	// Optional. Comment for __LBDN__.
	Comment *string `json:"comment,omitempty"`
	// Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.
	Disabled  *bool            `json:"disabled,omitempty"`
	DtcPolicy *ConfigDTCPolicy `json:"dtc_policy,omitempty"`
	// The resource identifier.
	Id                 *string               `json:"id,omitempty"`
	InheritanceSources *ConfigTTLInheritance `json:"inheritance_sources,omitempty"`
	// Name of __LBDN__.
	Name string `json:"name"`
	// Optional. Precedence.
	Precedence *int64 `json:"precedence,omitempty"`
	// Optional. The tags for __LBDN__ in JSON format.
	Tags map[string]interface{} `json:"tags,omitempty"`
	// Optional. Time to live value (in seconds) to be used for records in DTC response. Unsigned integer, min: 0, max 2147483647 (31-bits per RFC-2181).
	Ttl *int64 `json:"ttl,omitempty"`
	// The resource identifier.
	View string `json:"view"`
}

type _ConfigLBDN ConfigLBDN

// NewConfigLBDN instantiates a new ConfigLBDN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigLBDN(name string, view string) *ConfigLBDN {
	this := ConfigLBDN{}
	this.Name = name
	this.View = view
	return &this
}

// NewConfigLBDNWithDefaults instantiates a new ConfigLBDN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigLBDNWithDefaults() *ConfigLBDN {
	this := ConfigLBDN{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ConfigLBDN) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigLBDN) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ConfigLBDN) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ConfigLBDN) SetComment(v string) {
	o.Comment = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *ConfigLBDN) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigLBDN) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *ConfigLBDN) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *ConfigLBDN) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetDtcPolicy returns the DtcPolicy field value if set, zero value otherwise.
func (o *ConfigLBDN) GetDtcPolicy() ConfigDTCPolicy {
	if o == nil || IsNil(o.DtcPolicy) {
		var ret ConfigDTCPolicy
		return ret
	}
	return *o.DtcPolicy
}

// GetDtcPolicyOk returns a tuple with the DtcPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigLBDN) GetDtcPolicyOk() (*ConfigDTCPolicy, bool) {
	if o == nil || IsNil(o.DtcPolicy) {
		return nil, false
	}
	return o.DtcPolicy, true
}

// HasDtcPolicy returns a boolean if a field has been set.
func (o *ConfigLBDN) HasDtcPolicy() bool {
	if o != nil && !IsNil(o.DtcPolicy) {
		return true
	}

	return false
}

// SetDtcPolicy gets a reference to the given ConfigDTCPolicy and assigns it to the DtcPolicy field.
func (o *ConfigLBDN) SetDtcPolicy(v ConfigDTCPolicy) {
	o.DtcPolicy = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConfigLBDN) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigLBDN) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConfigLBDN) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConfigLBDN) SetId(v string) {
	o.Id = &v
}

// GetInheritanceSources returns the InheritanceSources field value if set, zero value otherwise.
func (o *ConfigLBDN) GetInheritanceSources() ConfigTTLInheritance {
	if o == nil || IsNil(o.InheritanceSources) {
		var ret ConfigTTLInheritance
		return ret
	}
	return *o.InheritanceSources
}

// GetInheritanceSourcesOk returns a tuple with the InheritanceSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigLBDN) GetInheritanceSourcesOk() (*ConfigTTLInheritance, bool) {
	if o == nil || IsNil(o.InheritanceSources) {
		return nil, false
	}
	return o.InheritanceSources, true
}

// HasInheritanceSources returns a boolean if a field has been set.
func (o *ConfigLBDN) HasInheritanceSources() bool {
	if o != nil && !IsNil(o.InheritanceSources) {
		return true
	}

	return false
}

// SetInheritanceSources gets a reference to the given ConfigTTLInheritance and assigns it to the InheritanceSources field.
func (o *ConfigLBDN) SetInheritanceSources(v ConfigTTLInheritance) {
	o.InheritanceSources = &v
}

// GetName returns the Name field value
func (o *ConfigLBDN) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigLBDN) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConfigLBDN) SetName(v string) {
	o.Name = v
}

// GetPrecedence returns the Precedence field value if set, zero value otherwise.
func (o *ConfigLBDN) GetPrecedence() int64 {
	if o == nil || IsNil(o.Precedence) {
		var ret int64
		return ret
	}
	return *o.Precedence
}

// GetPrecedenceOk returns a tuple with the Precedence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigLBDN) GetPrecedenceOk() (*int64, bool) {
	if o == nil || IsNil(o.Precedence) {
		return nil, false
	}
	return o.Precedence, true
}

// HasPrecedence returns a boolean if a field has been set.
func (o *ConfigLBDN) HasPrecedence() bool {
	if o != nil && !IsNil(o.Precedence) {
		return true
	}

	return false
}

// SetPrecedence gets a reference to the given int64 and assigns it to the Precedence field.
func (o *ConfigLBDN) SetPrecedence(v int64) {
	o.Precedence = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConfigLBDN) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigLBDN) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConfigLBDN) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *ConfigLBDN) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *ConfigLBDN) GetTtl() int64 {
	if o == nil || IsNil(o.Ttl) {
		var ret int64
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigLBDN) GetTtlOk() (*int64, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *ConfigLBDN) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int64 and assigns it to the Ttl field.
func (o *ConfigLBDN) SetTtl(v int64) {
	o.Ttl = &v
}

// GetView returns the View field value
func (o *ConfigLBDN) GetView() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.View
}

// GetViewOk returns a tuple with the View field value
// and a boolean to check if the value has been set.
func (o *ConfigLBDN) GetViewOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.View, true
}

// SetView sets field value
func (o *ConfigLBDN) SetView(v string) {
	o.View = v
}

func (o ConfigLBDN) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigLBDN) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.DtcPolicy) {
		toSerialize["dtc_policy"] = o.DtcPolicy
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InheritanceSources) {
		toSerialize["inheritance_sources"] = o.InheritanceSources
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Precedence) {
		toSerialize["precedence"] = o.Precedence
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	toSerialize["view"] = o.View
	return toSerialize, nil
}

func (o *ConfigLBDN) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"view",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varConfigLBDN := _ConfigLBDN{}

	err = json.Unmarshal(bytes, &varConfigLBDN)

	if err != nil {
		return err
	}

	*o = ConfigLBDN(varConfigLBDN)

	return err
}

type NullableConfigLBDN struct {
	value *ConfigLBDN
	isSet bool
}

func (v NullableConfigLBDN) Get() *ConfigLBDN {
	return v.value
}

func (v *NullableConfigLBDN) Set(val *ConfigLBDN) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigLBDN) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigLBDN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigLBDN(val *ConfigLBDN) *NullableConfigLBDN {
	return &NullableConfigLBDN{value: val, isSet: true}
}

func (v NullableConfigLBDN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigLBDN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

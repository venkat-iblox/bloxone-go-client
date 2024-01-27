/*
Host Activation Service

Host activation service provides a RESTful interface to manage cert and join token object. Join tokens are essentially a password that allows on-prem hosts to auto-associate themselves to a customer's account and receive a signed cert.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infra_provision

import (
	"encoding/json"
)

// checks if the TypesJSONValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TypesJSONValue{}

// TypesJSONValue struct for TypesJSONValue
type TypesJSONValue struct {
	Value *string `json:"value,omitempty"`
}

// NewTypesJSONValue instantiates a new TypesJSONValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesJSONValue() *TypesJSONValue {
	this := TypesJSONValue{}
	return &this
}

// NewTypesJSONValueWithDefaults instantiates a new TypesJSONValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesJSONValueWithDefaults() *TypesJSONValue {
	this := TypesJSONValue{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TypesJSONValue) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesJSONValue) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TypesJSONValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TypesJSONValue) SetValue(v string) {
	o.Value = &v
}

func (o TypesJSONValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TypesJSONValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTypesJSONValue struct {
	value *TypesJSONValue
	isSet bool
}

func (v NullableTypesJSONValue) Get() *TypesJSONValue {
	return v.value
}

func (v *NullableTypesJSONValue) Set(val *TypesJSONValue) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesJSONValue) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesJSONValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesJSONValue(val *TypesJSONValue) *NullableTypesJSONValue {
	return &NullableTypesJSONValue{value: val, isSet: true}
}

func (v NullableTypesJSONValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesJSONValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



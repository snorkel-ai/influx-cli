/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TemplateSummaryCheckAllOf struct for TemplateSummaryCheckAllOf
type TemplateSummaryCheckAllOf struct {
	Id          uint64  `json:"id" yaml:"id"`
	Name        string  `json:"name" yaml:"name"`
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
}

// NewTemplateSummaryCheckAllOf instantiates a new TemplateSummaryCheckAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryCheckAllOf(id uint64, name string) *TemplateSummaryCheckAllOf {
	this := TemplateSummaryCheckAllOf{}
	this.Id = id
	this.Name = name
	return &this
}

// NewTemplateSummaryCheckAllOfWithDefaults instantiates a new TemplateSummaryCheckAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryCheckAllOfWithDefaults() *TemplateSummaryCheckAllOf {
	this := TemplateSummaryCheckAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *TemplateSummaryCheckAllOf) GetId() uint64 {
	if o == nil {
		var ret uint64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryCheckAllOf) GetIdOk() (*uint64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TemplateSummaryCheckAllOf) SetId(v uint64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TemplateSummaryCheckAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryCheckAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateSummaryCheckAllOf) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateSummaryCheckAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryCheckAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateSummaryCheckAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplateSummaryCheckAllOf) SetDescription(v string) {
	o.Description = &v
}

func (o TemplateSummaryCheckAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryCheckAllOf struct {
	value *TemplateSummaryCheckAllOf
	isSet bool
}

func (v NullableTemplateSummaryCheckAllOf) Get() *TemplateSummaryCheckAllOf {
	return v.value
}

func (v *NullableTemplateSummaryCheckAllOf) Set(val *TemplateSummaryCheckAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryCheckAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryCheckAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryCheckAllOf(val *TemplateSummaryCheckAllOf) *NullableTemplateSummaryCheckAllOf {
	return &NullableTemplateSummaryCheckAllOf{value: val, isSet: true}
}

func (v NullableTemplateSummaryCheckAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryCheckAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
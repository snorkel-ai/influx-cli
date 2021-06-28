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

// TemplateSummaryCommonAllOf struct for TemplateSummaryCommonAllOf
type TemplateSummaryCommonAllOf struct {
	LabelAssociations []TemplateSummaryLabel `json:"labelAssociations" yaml:"labelAssociations"`
}

// NewTemplateSummaryCommonAllOf instantiates a new TemplateSummaryCommonAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryCommonAllOf(labelAssociations []TemplateSummaryLabel) *TemplateSummaryCommonAllOf {
	this := TemplateSummaryCommonAllOf{}
	this.LabelAssociations = labelAssociations
	return &this
}

// NewTemplateSummaryCommonAllOfWithDefaults instantiates a new TemplateSummaryCommonAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryCommonAllOfWithDefaults() *TemplateSummaryCommonAllOf {
	this := TemplateSummaryCommonAllOf{}
	return &this
}

// GetLabelAssociations returns the LabelAssociations field value
func (o *TemplateSummaryCommonAllOf) GetLabelAssociations() []TemplateSummaryLabel {
	if o == nil {
		var ret []TemplateSummaryLabel
		return ret
	}

	return o.LabelAssociations
}

// GetLabelAssociationsOk returns a tuple with the LabelAssociations field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryCommonAllOf) GetLabelAssociationsOk() (*[]TemplateSummaryLabel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelAssociations, true
}

// SetLabelAssociations sets field value
func (o *TemplateSummaryCommonAllOf) SetLabelAssociations(v []TemplateSummaryLabel) {
	o.LabelAssociations = v
}

func (o TemplateSummaryCommonAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["labelAssociations"] = o.LabelAssociations
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryCommonAllOf struct {
	value *TemplateSummaryCommonAllOf
	isSet bool
}

func (v NullableTemplateSummaryCommonAllOf) Get() *TemplateSummaryCommonAllOf {
	return v.value
}

func (v *NullableTemplateSummaryCommonAllOf) Set(val *TemplateSummaryCommonAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryCommonAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryCommonAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryCommonAllOf(val *TemplateSummaryCommonAllOf) *NullableTemplateSummaryCommonAllOf {
	return &NullableTemplateSummaryCommonAllOf{value: val, isSet: true}
}

func (v NullableTemplateSummaryCommonAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryCommonAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

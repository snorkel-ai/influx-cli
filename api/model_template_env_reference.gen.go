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

// TemplateEnvReference struct for TemplateEnvReference
type TemplateEnvReference struct {
	// Field the environment reference corresponds too
	ResourceField string `json:"resourceField" yaml:"resourceField"`
	// Key identified as environment reference and is the key identified in the template
	EnvRefKey string `json:"envRefKey" yaml:"envRefKey"`
	// Value provided to fulfill reference
	Value interface{} `json:"value,omitempty" yaml:"value,omitempty"`
	// Default value that will be provided for the reference when no value is provided
	DefaultValue interface{} `json:"defaultValue,omitempty" yaml:"defaultValue,omitempty"`
}

// NewTemplateEnvReference instantiates a new TemplateEnvReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateEnvReference(resourceField string, envRefKey string) *TemplateEnvReference {
	this := TemplateEnvReference{}
	this.ResourceField = resourceField
	this.EnvRefKey = envRefKey
	return &this
}

// NewTemplateEnvReferenceWithDefaults instantiates a new TemplateEnvReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateEnvReferenceWithDefaults() *TemplateEnvReference {
	this := TemplateEnvReference{}
	return &this
}

// GetResourceField returns the ResourceField field value
func (o *TemplateEnvReference) GetResourceField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceField
}

// GetResourceFieldOk returns a tuple with the ResourceField field value
// and a boolean to check if the value has been set.
func (o *TemplateEnvReference) GetResourceFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceField, true
}

// SetResourceField sets field value
func (o *TemplateEnvReference) SetResourceField(v string) {
	o.ResourceField = v
}

// GetEnvRefKey returns the EnvRefKey field value
func (o *TemplateEnvReference) GetEnvRefKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvRefKey
}

// GetEnvRefKeyOk returns a tuple with the EnvRefKey field value
// and a boolean to check if the value has been set.
func (o *TemplateEnvReference) GetEnvRefKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvRefKey, true
}

// SetEnvRefKey sets field value
func (o *TemplateEnvReference) SetEnvRefKey(v string) {
	o.EnvRefKey = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvReference) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvReference) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TemplateEnvReference) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *TemplateEnvReference) SetValue(v interface{}) {
	o.Value = v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvReference) GetDefaultValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvReference) GetDefaultValueOk() (*interface{}, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return &o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *TemplateEnvReference) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given interface{} and assigns it to the DefaultValue field.
func (o *TemplateEnvReference) SetDefaultValue(v interface{}) {
	o.DefaultValue = v
}

func (o TemplateEnvReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resourceField"] = o.ResourceField
	}
	if true {
		toSerialize["envRefKey"] = o.EnvRefKey
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateEnvReference struct {
	value *TemplateEnvReference
	isSet bool
}

func (v NullableTemplateEnvReference) Get() *TemplateEnvReference {
	return v.value
}

func (v *NullableTemplateEnvReference) Set(val *TemplateEnvReference) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateEnvReference) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateEnvReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateEnvReference(val *TemplateEnvReference) *NullableTemplateEnvReference {
	return &NullableTemplateEnvReference{value: val, isSet: true}
}

func (v NullableTemplateEnvReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateEnvReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

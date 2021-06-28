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

// TemplateSummaryNotificationEndpointAllOf struct for TemplateSummaryNotificationEndpointAllOf
type TemplateSummaryNotificationEndpointAllOf struct {
	Id          string  `json:"id" yaml:"id"`
	Name        string  `json:"name" yaml:"name"`
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	Status      string  `json:"status" yaml:"status"`
}

// NewTemplateSummaryNotificationEndpointAllOf instantiates a new TemplateSummaryNotificationEndpointAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryNotificationEndpointAllOf(id string, name string, status string) *TemplateSummaryNotificationEndpointAllOf {
	this := TemplateSummaryNotificationEndpointAllOf{}
	this.Id = id
	this.Name = name
	this.Status = status
	return &this
}

// NewTemplateSummaryNotificationEndpointAllOfWithDefaults instantiates a new TemplateSummaryNotificationEndpointAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryNotificationEndpointAllOfWithDefaults() *TemplateSummaryNotificationEndpointAllOf {
	this := TemplateSummaryNotificationEndpointAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *TemplateSummaryNotificationEndpointAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpointAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TemplateSummaryNotificationEndpointAllOf) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TemplateSummaryNotificationEndpointAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpointAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateSummaryNotificationEndpointAllOf) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateSummaryNotificationEndpointAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpointAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateSummaryNotificationEndpointAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplateSummaryNotificationEndpointAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value
func (o *TemplateSummaryNotificationEndpointAllOf) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryNotificationEndpointAllOf) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TemplateSummaryNotificationEndpointAllOf) SetStatus(v string) {
	o.Status = v
}

func (o TemplateSummaryNotificationEndpointAllOf) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryNotificationEndpointAllOf struct {
	value *TemplateSummaryNotificationEndpointAllOf
	isSet bool
}

func (v NullableTemplateSummaryNotificationEndpointAllOf) Get() *TemplateSummaryNotificationEndpointAllOf {
	return v.value
}

func (v *NullableTemplateSummaryNotificationEndpointAllOf) Set(val *TemplateSummaryNotificationEndpointAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryNotificationEndpointAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryNotificationEndpointAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryNotificationEndpointAllOf(val *TemplateSummaryNotificationEndpointAllOf) *NullableTemplateSummaryNotificationEndpointAllOf {
	return &NullableTemplateSummaryNotificationEndpointAllOf{value: val, isSet: true}
}

func (v NullableTemplateSummaryNotificationEndpointAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryNotificationEndpointAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

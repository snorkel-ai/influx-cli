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

// Logs struct for Logs
type Logs struct {
	Events *[]LogEvent `json:"events,omitempty" yaml:"events,omitempty"`
}

// NewLogs instantiates a new Logs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogs() *Logs {
	this := Logs{}
	return &this
}

// NewLogsWithDefaults instantiates a new Logs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsWithDefaults() *Logs {
	this := Logs{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *Logs) GetEvents() []LogEvent {
	if o == nil || o.Events == nil {
		var ret []LogEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Logs) GetEventsOk() (*[]LogEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *Logs) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []LogEvent and assigns it to the Events field.
func (o *Logs) SetEvents(v []LogEvent) {
	o.Events = &v
}

func (o Logs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableLogs struct {
	value *Logs
	isSet bool
}

func (v NullableLogs) Get() *Logs {
	return v.value
}

func (v *NullableLogs) Set(val *Logs) {
	v.value = val
	v.isSet = true
}

func (v NullableLogs) IsSet() bool {
	return v.isSet
}

func (v *NullableLogs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogs(val *Logs) *NullableLogs {
	return &NullableLogs{value: val, isSet: true}
}

func (v NullableLogs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

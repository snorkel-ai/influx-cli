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

// RetentionRule struct for RetentionRule
type RetentionRule struct {
	Type string `json:"type" yaml:"type"`
	// Duration in seconds for how long data will be kept in the database. 0 means infinite.
	EverySeconds int64 `json:"everySeconds" yaml:"everySeconds"`
	// Shard duration measured in seconds.
	ShardGroupDurationSeconds *int64 `json:"shardGroupDurationSeconds,omitempty" yaml:"shardGroupDurationSeconds,omitempty"`
}

// NewRetentionRule instantiates a new RetentionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetentionRule(type_ string, everySeconds int64) *RetentionRule {
	this := RetentionRule{}
	this.Type = type_
	this.EverySeconds = everySeconds
	return &this
}

// NewRetentionRuleWithDefaults instantiates a new RetentionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetentionRuleWithDefaults() *RetentionRule {
	this := RetentionRule{}
	var type_ string = "expire"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *RetentionRule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RetentionRule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RetentionRule) SetType(v string) {
	o.Type = v
}

// GetEverySeconds returns the EverySeconds field value
func (o *RetentionRule) GetEverySeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EverySeconds
}

// GetEverySecondsOk returns a tuple with the EverySeconds field value
// and a boolean to check if the value has been set.
func (o *RetentionRule) GetEverySecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EverySeconds, true
}

// SetEverySeconds sets field value
func (o *RetentionRule) SetEverySeconds(v int64) {
	o.EverySeconds = v
}

// GetShardGroupDurationSeconds returns the ShardGroupDurationSeconds field value if set, zero value otherwise.
func (o *RetentionRule) GetShardGroupDurationSeconds() int64 {
	if o == nil || o.ShardGroupDurationSeconds == nil {
		var ret int64
		return ret
	}
	return *o.ShardGroupDurationSeconds
}

// GetShardGroupDurationSecondsOk returns a tuple with the ShardGroupDurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionRule) GetShardGroupDurationSecondsOk() (*int64, bool) {
	if o == nil || o.ShardGroupDurationSeconds == nil {
		return nil, false
	}
	return o.ShardGroupDurationSeconds, true
}

// HasShardGroupDurationSeconds returns a boolean if a field has been set.
func (o *RetentionRule) HasShardGroupDurationSeconds() bool {
	if o != nil && o.ShardGroupDurationSeconds != nil {
		return true
	}

	return false
}

// SetShardGroupDurationSeconds gets a reference to the given int64 and assigns it to the ShardGroupDurationSeconds field.
func (o *RetentionRule) SetShardGroupDurationSeconds(v int64) {
	o.ShardGroupDurationSeconds = &v
}

func (o RetentionRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["everySeconds"] = o.EverySeconds
	}
	if o.ShardGroupDurationSeconds != nil {
		toSerialize["shardGroupDurationSeconds"] = o.ShardGroupDurationSeconds
	}
	return json.Marshal(toSerialize)
}

type NullableRetentionRule struct {
	value *RetentionRule
	isSet bool
}

func (v NullableRetentionRule) Get() *RetentionRule {
	return v.value
}

func (v *NullableRetentionRule) Set(val *RetentionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableRetentionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableRetentionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetentionRule(val *RetentionRule) *NullableRetentionRule {
	return &NullableRetentionRule{value: val, isSet: true}
}

func (v NullableRetentionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetentionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

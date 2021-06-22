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

// RestoredBucketMappings struct for RestoredBucketMappings
type RestoredBucketMappings struct {
	// New ID of the restored bucket
	Id            string               `json:"id"`
	Name          string               `json:"name"`
	ShardMappings []BucketShardMapping `json:"shardMappings"`
}

// NewRestoredBucketMappings instantiates a new RestoredBucketMappings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoredBucketMappings(id string, name string, shardMappings []BucketShardMapping) *RestoredBucketMappings {
	this := RestoredBucketMappings{}
	this.Id = id
	this.Name = name
	this.ShardMappings = shardMappings
	return &this
}

// NewRestoredBucketMappingsWithDefaults instantiates a new RestoredBucketMappings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoredBucketMappingsWithDefaults() *RestoredBucketMappings {
	this := RestoredBucketMappings{}
	return &this
}

// GetId returns the Id field value
func (o *RestoredBucketMappings) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RestoredBucketMappings) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RestoredBucketMappings) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *RestoredBucketMappings) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RestoredBucketMappings) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RestoredBucketMappings) SetName(v string) {
	o.Name = v
}

// GetShardMappings returns the ShardMappings field value
func (o *RestoredBucketMappings) GetShardMappings() []BucketShardMapping {
	if o == nil {
		var ret []BucketShardMapping
		return ret
	}

	return o.ShardMappings
}

// GetShardMappingsOk returns a tuple with the ShardMappings field value
// and a boolean to check if the value has been set.
func (o *RestoredBucketMappings) GetShardMappingsOk() (*[]BucketShardMapping, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShardMappings, true
}

// SetShardMappings sets field value
func (o *RestoredBucketMappings) SetShardMappings(v []BucketShardMapping) {
	o.ShardMappings = v
}

func (o RestoredBucketMappings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["shardMappings"] = o.ShardMappings
	}
	return json.Marshal(toSerialize)
}

type NullableRestoredBucketMappings struct {
	value *RestoredBucketMappings
	isSet bool
}

func (v NullableRestoredBucketMappings) Get() *RestoredBucketMappings {
	return v.value
}

func (v *NullableRestoredBucketMappings) Set(val *RestoredBucketMappings) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoredBucketMappings) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoredBucketMappings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoredBucketMappings(val *RestoredBucketMappings) *NullableRestoredBucketMappings {
	return &NullableRestoredBucketMappings{value: val, isSet: true}
}

func (v NullableRestoredBucketMappings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoredBucketMappings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
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
	"time"
)

// Authorization struct for Authorization
type Authorization struct {
	// If inactive the token is inactive and requests using the token will be rejected.
	Status *string `json:"status,omitempty" yaml:"status,omitempty"`
	// A description of the token.
	Description *string    `json:"description,omitempty" yaml:"description,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty" yaml:"updatedAt,omitempty"`
	// ID of org that authorization is scoped to.
	OrgID string `json:"orgID" yaml:"orgID"`
	// List of permissions for an auth.  An auth must have at least one Permission.
	Permissions []Permission `json:"permissions" yaml:"permissions"`
	Id          *string      `json:"id,omitempty" yaml:"id,omitempty"`
	// Passed via the Authorization Header and Token Authentication type.
	Token *string `json:"token,omitempty" yaml:"token,omitempty"`
	// ID of user that created and owns the token.
	UserID *string `json:"userID,omitempty" yaml:"userID,omitempty"`
	// Name of user that created and owns the token.
	User *string `json:"user,omitempty" yaml:"user,omitempty"`
	// Name of the org token is scoped to.
	Org   *string                  `json:"org,omitempty" yaml:"org,omitempty"`
	Links *AuthorizationAllOfLinks `json:"links,omitempty" yaml:"links,omitempty"`
}

// NewAuthorization instantiates a new Authorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorization(orgID string, permissions []Permission) *Authorization {
	this := Authorization{}
	var status string = "active"
	this.Status = &status
	this.OrgID = orgID
	this.Permissions = permissions
	return &this
}

// NewAuthorizationWithDefaults instantiates a new Authorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationWithDefaults() *Authorization {
	this := Authorization{}
	var status string = "active"
	this.Status = &status
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Authorization) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authorization) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Authorization) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Authorization) SetStatus(v string) {
	o.Status = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Authorization) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authorization) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Authorization) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Authorization) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Authorization) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authorization) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Authorization) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Authorization) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Authorization) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authorization) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Authorization) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Authorization) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetOrgID returns the OrgID field value
func (o *Authorization) GetOrgID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgID
}

// GetOrgIDOk returns a tuple with the OrgID field value
// and a boolean to check if the value has been set.
func (o *Authorization) GetOrgIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgID, true
}

// SetOrgID sets field value
func (o *Authorization) SetOrgID(v string) {
	o.OrgID = v
}

// GetPermissions returns the Permissions field value
func (o *Authorization) GetPermissions() []Permission {
	if o == nil {
		var ret []Permission
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *Authorization) GetPermissionsOk() (*[]Permission, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value
func (o *Authorization) SetPermissions(v []Permission) {
	o.Permissions = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Authorization) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authorization) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Authorization) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Authorization) SetId(v string) {
	o.Id = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *Authorization) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authorization) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *Authorization) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *Authorization) SetToken(v string) {
	o.Token = &v
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *Authorization) GetUserID() string {
	if o == nil || o.UserID == nil {
		var ret string
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authorization) GetUserIDOk() (*string, bool) {
	if o == nil || o.UserID == nil {
		return nil, false
	}
	return o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *Authorization) HasUserID() bool {
	if o != nil && o.UserID != nil {
		return true
	}

	return false
}

// SetUserID gets a reference to the given string and assigns it to the UserID field.
func (o *Authorization) SetUserID(v string) {
	o.UserID = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Authorization) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authorization) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Authorization) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *Authorization) SetUser(v string) {
	o.User = &v
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *Authorization) GetOrg() string {
	if o == nil || o.Org == nil {
		var ret string
		return ret
	}
	return *o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authorization) GetOrgOk() (*string, bool) {
	if o == nil || o.Org == nil {
		return nil, false
	}
	return o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *Authorization) HasOrg() bool {
	if o != nil && o.Org != nil {
		return true
	}

	return false
}

// SetOrg gets a reference to the given string and assigns it to the Org field.
func (o *Authorization) SetOrg(v string) {
	o.Org = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Authorization) GetLinks() AuthorizationAllOfLinks {
	if o == nil || o.Links == nil {
		var ret AuthorizationAllOfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authorization) GetLinksOk() (*AuthorizationAllOfLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Authorization) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AuthorizationAllOfLinks and assigns it to the Links field.
func (o *Authorization) SetLinks(v AuthorizationAllOfLinks) {
	o.Links = &v
}

func (o Authorization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if true {
		toSerialize["orgID"] = o.OrgID
	}
	if true {
		toSerialize["permissions"] = o.Permissions
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UserID != nil {
		toSerialize["userID"] = o.UserID
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Org != nil {
		toSerialize["org"] = o.Org
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorization struct {
	value *Authorization
	isSet bool
}

func (v NullableAuthorization) Get() *Authorization {
	return v.value
}

func (v *NullableAuthorization) Set(val *Authorization) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorization(val *Authorization) *NullableAuthorization {
	return &NullableAuthorization{value: val, isSet: true}
}

func (v NullableAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

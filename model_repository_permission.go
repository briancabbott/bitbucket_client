/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bitbucket_client

import (
	"encoding/json"
)

// RepositoryPermission A user's permission for a given repository.
type RepositoryPermission struct {
	Type string `json:"type"`
	Permission *string `json:"permission,omitempty"`
	User *User `json:"user,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
}

// NewRepositoryPermission instantiates a new RepositoryPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryPermission(type_ string) *RepositoryPermission {
	this := RepositoryPermission{}
	this.Type = type_
	return &this
}

// NewRepositoryPermissionWithDefaults instantiates a new RepositoryPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryPermissionWithDefaults() *RepositoryPermission {
	this := RepositoryPermission{}
	return &this
}

// GetType returns the Type field value
func (o *RepositoryPermission) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RepositoryPermission) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RepositoryPermission) SetType(v string) {
	o.Type = v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *RepositoryPermission) GetPermission() string {
	if o == nil || o.Permission == nil {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryPermission) GetPermissionOk() (*string, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *RepositoryPermission) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *RepositoryPermission) SetPermission(v string) {
	o.Permission = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RepositoryPermission) GetUser() User {
	if o == nil || o.User == nil {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryPermission) GetUserOk() (*User, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RepositoryPermission) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *RepositoryPermission) SetUser(v User) {
	o.User = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *RepositoryPermission) GetRepository() Repository {
	if o == nil || o.Repository == nil {
		var ret Repository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryPermission) GetRepositoryOk() (*Repository, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *RepositoryPermission) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given Repository and assigns it to the Repository field.
func (o *RepositoryPermission) SetRepository(v Repository) {
	o.Repository = &v
}

func (o RepositoryPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Permission != nil {
		toSerialize["permission"] = o.Permission
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryPermission struct {
	value *RepositoryPermission
	isSet bool
}

func (v NullableRepositoryPermission) Get() *RepositoryPermission {
	return v.value
}

func (v *NullableRepositoryPermission) Set(val *RepositoryPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryPermission(val *RepositoryPermission) *NullableRepositoryPermission {
	return &NullableRepositoryPermission{value: val, isSet: true}
}

func (v NullableRepositoryPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// DeploymentsDdevDeploymentEnvironment struct for DeploymentsDdevDeploymentEnvironment
type DeploymentsDdevDeploymentEnvironment struct {
	Object
}

// NewDeploymentsDdevDeploymentEnvironment instantiates a new DeploymentsDdevDeploymentEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentsDdevDeploymentEnvironment(type_ string) *DeploymentsDdevDeploymentEnvironment {
	this := DeploymentsDdevDeploymentEnvironment{}
	this.Type = type_
	return &this
}

// NewDeploymentsDdevDeploymentEnvironmentWithDefaults instantiates a new DeploymentsDdevDeploymentEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentsDdevDeploymentEnvironmentWithDefaults() *DeploymentsDdevDeploymentEnvironment {
	this := DeploymentsDdevDeploymentEnvironment{}
	return &this
}

func (o DeploymentsDdevDeploymentEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedObject, errObject := json.Marshal(o.Object)
	if errObject != nil {
		return []byte{}, errObject
	}
	errObject = json.Unmarshal([]byte(serializedObject), &toSerialize)
	if errObject != nil {
		return []byte{}, errObject
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentsDdevDeploymentEnvironment struct {
	value *DeploymentsDdevDeploymentEnvironment
	isSet bool
}

func (v NullableDeploymentsDdevDeploymentEnvironment) Get() *DeploymentsDdevDeploymentEnvironment {
	return v.value
}

func (v *NullableDeploymentsDdevDeploymentEnvironment) Set(val *DeploymentsDdevDeploymentEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentsDdevDeploymentEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentsDdevDeploymentEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentsDdevDeploymentEnvironment(val *DeploymentsDdevDeploymentEnvironment) *NullableDeploymentsDdevDeploymentEnvironment {
	return &NullableDeploymentsDdevDeploymentEnvironment{value: val, isSet: true}
}

func (v NullableDeploymentsDdevDeploymentEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentsDdevDeploymentEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



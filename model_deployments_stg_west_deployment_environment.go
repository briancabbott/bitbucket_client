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

// DeploymentsStgWestDeploymentEnvironment struct for DeploymentsStgWestDeploymentEnvironment
type DeploymentsStgWestDeploymentEnvironment struct {
	Object
}

// NewDeploymentsStgWestDeploymentEnvironment instantiates a new DeploymentsStgWestDeploymentEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentsStgWestDeploymentEnvironment(type_ string) *DeploymentsStgWestDeploymentEnvironment {
	this := DeploymentsStgWestDeploymentEnvironment{}
	this.Type = type_
	return &this
}

// NewDeploymentsStgWestDeploymentEnvironmentWithDefaults instantiates a new DeploymentsStgWestDeploymentEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentsStgWestDeploymentEnvironmentWithDefaults() *DeploymentsStgWestDeploymentEnvironment {
	this := DeploymentsStgWestDeploymentEnvironment{}
	return &this
}

func (o DeploymentsStgWestDeploymentEnvironment) MarshalJSON() ([]byte, error) {
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

type NullableDeploymentsStgWestDeploymentEnvironment struct {
	value *DeploymentsStgWestDeploymentEnvironment
	isSet bool
}

func (v NullableDeploymentsStgWestDeploymentEnvironment) Get() *DeploymentsStgWestDeploymentEnvironment {
	return v.value
}

func (v *NullableDeploymentsStgWestDeploymentEnvironment) Set(val *DeploymentsStgWestDeploymentEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentsStgWestDeploymentEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentsStgWestDeploymentEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentsStgWestDeploymentEnvironment(val *DeploymentsStgWestDeploymentEnvironment) *NullableDeploymentsStgWestDeploymentEnvironment {
	return &NullableDeploymentsStgWestDeploymentEnvironment{value: val, isSet: true}
}

func (v NullableDeploymentsStgWestDeploymentEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentsStgWestDeploymentEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



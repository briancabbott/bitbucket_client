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

// PipelineStateCompletedResult struct for PipelineStateCompletedResult
type PipelineStateCompletedResult struct {
	Object
}

// NewPipelineStateCompletedResult instantiates a new PipelineStateCompletedResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStateCompletedResult(type_ string) *PipelineStateCompletedResult {
	this := PipelineStateCompletedResult{}
	this.Type = type_
	return &this
}

// NewPipelineStateCompletedResultWithDefaults instantiates a new PipelineStateCompletedResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStateCompletedResultWithDefaults() *PipelineStateCompletedResult {
	this := PipelineStateCompletedResult{}
	return &this
}

func (o PipelineStateCompletedResult) MarshalJSON() ([]byte, error) {
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

type NullablePipelineStateCompletedResult struct {
	value *PipelineStateCompletedResult
	isSet bool
}

func (v NullablePipelineStateCompletedResult) Get() *PipelineStateCompletedResult {
	return v.value
}

func (v *NullablePipelineStateCompletedResult) Set(val *PipelineStateCompletedResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStateCompletedResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStateCompletedResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStateCompletedResult(val *PipelineStateCompletedResult) *NullablePipelineStateCompletedResult {
	return &NullablePipelineStateCompletedResult{value: val, isSet: true}
}

func (v NullablePipelineStateCompletedResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStateCompletedResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



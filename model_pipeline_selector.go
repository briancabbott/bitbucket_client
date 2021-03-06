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

// PipelineSelector struct for PipelineSelector
type PipelineSelector struct {
	Object
	// The type of selector.
	Type string `json:"type"`
	// The name of the matching pipeline definition.
	Pattern *string `json:"pattern,omitempty"`
}

// NewPipelineSelector instantiates a new PipelineSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineSelector(type_ string) *PipelineSelector {
	this := PipelineSelector{}
	this.Type = type_
	return &this
}

// NewPipelineSelectorWithDefaults instantiates a new PipelineSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineSelectorWithDefaults() *PipelineSelector {
	this := PipelineSelector{}
	return &this
}

// GetType returns the Type field value
func (o *PipelineSelector) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PipelineSelector) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PipelineSelector) SetType(v string) {
	o.Type = v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *PipelineSelector) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineSelector) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *PipelineSelector) HasPattern() bool {
	if o != nil && o.Pattern != nil {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *PipelineSelector) SetPattern(v string) {
	o.Pattern = &v
}

func (o PipelineSelector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedObject, errObject := json.Marshal(o.Object)
	if errObject != nil {
		return []byte{}, errObject
	}
	errObject = json.Unmarshal([]byte(serializedObject), &toSerialize)
	if errObject != nil {
		return []byte{}, errObject
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineSelector struct {
	value *PipelineSelector
	isSet bool
}

func (v NullablePipelineSelector) Get() *PipelineSelector {
	return v.value
}

func (v *NullablePipelineSelector) Set(val *PipelineSelector) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineSelector) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineSelector(val *PipelineSelector) *NullablePipelineSelector {
	return &NullablePipelineSelector{value: val, isSet: true}
}

func (v NullablePipelineSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



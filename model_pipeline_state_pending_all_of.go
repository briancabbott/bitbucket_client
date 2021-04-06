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

// PipelineStatePendingAllOf A Bitbucket Pipelines PENDING pipeline state.
type PipelineStatePendingAllOf struct {
	// The name of pipeline state (PENDING).
	Name *string `json:"name,omitempty"`
}

// NewPipelineStatePendingAllOf instantiates a new PipelineStatePendingAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStatePendingAllOf() *PipelineStatePendingAllOf {
	this := PipelineStatePendingAllOf{}
	return &this
}

// NewPipelineStatePendingAllOfWithDefaults instantiates a new PipelineStatePendingAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStatePendingAllOfWithDefaults() *PipelineStatePendingAllOf {
	this := PipelineStatePendingAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PipelineStatePendingAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStatePendingAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PipelineStatePendingAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PipelineStatePendingAllOf) SetName(v string) {
	o.Name = &v
}

func (o PipelineStatePendingAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineStatePendingAllOf struct {
	value *PipelineStatePendingAllOf
	isSet bool
}

func (v NullablePipelineStatePendingAllOf) Get() *PipelineStatePendingAllOf {
	return v.value
}

func (v *NullablePipelineStatePendingAllOf) Set(val *PipelineStatePendingAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStatePendingAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStatePendingAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStatePendingAllOf(val *PipelineStatePendingAllOf) *NullablePipelineStatePendingAllOf {
	return &NullablePipelineStatePendingAllOf{value: val, isSet: true}
}

func (v NullablePipelineStatePendingAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStatePendingAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// PipelineStepErrorAllOf An error causing a step failure.
type PipelineStepErrorAllOf struct {
	// The error key.
	Key *string `json:"key,omitempty"`
	// The error message.
	Message *string `json:"message,omitempty"`
}

// NewPipelineStepErrorAllOf instantiates a new PipelineStepErrorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStepErrorAllOf() *PipelineStepErrorAllOf {
	this := PipelineStepErrorAllOf{}
	return &this
}

// NewPipelineStepErrorAllOfWithDefaults instantiates a new PipelineStepErrorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStepErrorAllOfWithDefaults() *PipelineStepErrorAllOf {
	this := PipelineStepErrorAllOf{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *PipelineStepErrorAllOf) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStepErrorAllOf) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *PipelineStepErrorAllOf) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *PipelineStepErrorAllOf) SetKey(v string) {
	o.Key = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PipelineStepErrorAllOf) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStepErrorAllOf) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PipelineStepErrorAllOf) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PipelineStepErrorAllOf) SetMessage(v string) {
	o.Message = &v
}

func (o PipelineStepErrorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineStepErrorAllOf struct {
	value *PipelineStepErrorAllOf
	isSet bool
}

func (v NullablePipelineStepErrorAllOf) Get() *PipelineStepErrorAllOf {
	return v.value
}

func (v *NullablePipelineStepErrorAllOf) Set(val *PipelineStepErrorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStepErrorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStepErrorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStepErrorAllOf(val *PipelineStepErrorAllOf) *NullablePipelineStepErrorAllOf {
	return &NullablePipelineStepErrorAllOf{value: val, isSet: true}
}

func (v NullablePipelineStepErrorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStepErrorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



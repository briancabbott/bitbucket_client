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

// PullrequestComment struct for PullrequestComment
type PullrequestComment struct {
	Comment
	Type string `json:"type"`
	Pullrequest *Pullrequest `json:"pullrequest,omitempty"`
}

// NewPullrequestComment instantiates a new PullrequestComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullrequestComment(type_ string) *PullrequestComment {
	this := PullrequestComment{}
	this.Type = type_
	return &this
}

// NewPullrequestCommentWithDefaults instantiates a new PullrequestComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullrequestCommentWithDefaults() *PullrequestComment {
	this := PullrequestComment{}
	return &this
}

// GetType returns the Type field value
func (o *PullrequestComment) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PullrequestComment) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PullrequestComment) SetType(v string) {
	o.Type = v
}

// GetPullrequest returns the Pullrequest field value if set, zero value otherwise.
func (o *PullrequestComment) GetPullrequest() Pullrequest {
	if o == nil || o.Pullrequest == nil {
		var ret Pullrequest
		return ret
	}
	return *o.Pullrequest
}

// GetPullrequestOk returns a tuple with the Pullrequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullrequestComment) GetPullrequestOk() (*Pullrequest, bool) {
	if o == nil || o.Pullrequest == nil {
		return nil, false
	}
	return o.Pullrequest, true
}

// HasPullrequest returns a boolean if a field has been set.
func (o *PullrequestComment) HasPullrequest() bool {
	if o != nil && o.Pullrequest != nil {
		return true
	}

	return false
}

// SetPullrequest gets a reference to the given Pullrequest and assigns it to the Pullrequest field.
func (o *PullrequestComment) SetPullrequest(v Pullrequest) {
	o.Pullrequest = &v
}

func (o PullrequestComment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedComment, errComment := json.Marshal(o.Comment)
	if errComment != nil {
		return []byte{}, errComment
	}
	errComment = json.Unmarshal([]byte(serializedComment), &toSerialize)
	if errComment != nil {
		return []byte{}, errComment
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Pullrequest != nil {
		toSerialize["pullrequest"] = o.Pullrequest
	}
	return json.Marshal(toSerialize)
}

type NullablePullrequestComment struct {
	value *PullrequestComment
	isSet bool
}

func (v NullablePullrequestComment) Get() *PullrequestComment {
	return v.value
}

func (v *NullablePullrequestComment) Set(val *PullrequestComment) {
	v.value = val
	v.isSet = true
}

func (v NullablePullrequestComment) IsSet() bool {
	return v.isSet
}

func (v *NullablePullrequestComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullrequestComment(val *PullrequestComment) *NullablePullrequestComment {
	return &NullablePullrequestComment{value: val, isSet: true}
}

func (v NullablePullrequestComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullrequestComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



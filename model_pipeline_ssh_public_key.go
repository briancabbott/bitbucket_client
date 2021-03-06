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

// PipelineSshPublicKey struct for PipelineSshPublicKey
type PipelineSshPublicKey struct {
	Object
	// The type of the public key.
	KeyType *string `json:"key_type,omitempty"`
	// The base64 encoded public key.
	Key *string `json:"key,omitempty"`
	// The MD5 fingerprint of the public key.
	Md5Fingerprint *string `json:"md5_fingerprint,omitempty"`
	// The SHA-256 fingerprint of the public key.
	Sha256Fingerprint *string `json:"sha256_fingerprint,omitempty"`
}

// NewPipelineSshPublicKey instantiates a new PipelineSshPublicKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineSshPublicKey(type_ string) *PipelineSshPublicKey {
	this := PipelineSshPublicKey{}
	this.Type = type_
	return &this
}

// NewPipelineSshPublicKeyWithDefaults instantiates a new PipelineSshPublicKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineSshPublicKeyWithDefaults() *PipelineSshPublicKey {
	this := PipelineSshPublicKey{}
	return &this
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *PipelineSshPublicKey) GetKeyType() string {
	if o == nil || o.KeyType == nil {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineSshPublicKey) GetKeyTypeOk() (*string, bool) {
	if o == nil || o.KeyType == nil {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *PipelineSshPublicKey) HasKeyType() bool {
	if o != nil && o.KeyType != nil {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *PipelineSshPublicKey) SetKeyType(v string) {
	o.KeyType = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *PipelineSshPublicKey) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineSshPublicKey) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *PipelineSshPublicKey) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *PipelineSshPublicKey) SetKey(v string) {
	o.Key = &v
}

// GetMd5Fingerprint returns the Md5Fingerprint field value if set, zero value otherwise.
func (o *PipelineSshPublicKey) GetMd5Fingerprint() string {
	if o == nil || o.Md5Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Md5Fingerprint
}

// GetMd5FingerprintOk returns a tuple with the Md5Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineSshPublicKey) GetMd5FingerprintOk() (*string, bool) {
	if o == nil || o.Md5Fingerprint == nil {
		return nil, false
	}
	return o.Md5Fingerprint, true
}

// HasMd5Fingerprint returns a boolean if a field has been set.
func (o *PipelineSshPublicKey) HasMd5Fingerprint() bool {
	if o != nil && o.Md5Fingerprint != nil {
		return true
	}

	return false
}

// SetMd5Fingerprint gets a reference to the given string and assigns it to the Md5Fingerprint field.
func (o *PipelineSshPublicKey) SetMd5Fingerprint(v string) {
	o.Md5Fingerprint = &v
}

// GetSha256Fingerprint returns the Sha256Fingerprint field value if set, zero value otherwise.
func (o *PipelineSshPublicKey) GetSha256Fingerprint() string {
	if o == nil || o.Sha256Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Sha256Fingerprint
}

// GetSha256FingerprintOk returns a tuple with the Sha256Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineSshPublicKey) GetSha256FingerprintOk() (*string, bool) {
	if o == nil || o.Sha256Fingerprint == nil {
		return nil, false
	}
	return o.Sha256Fingerprint, true
}

// HasSha256Fingerprint returns a boolean if a field has been set.
func (o *PipelineSshPublicKey) HasSha256Fingerprint() bool {
	if o != nil && o.Sha256Fingerprint != nil {
		return true
	}

	return false
}

// SetSha256Fingerprint gets a reference to the given string and assigns it to the Sha256Fingerprint field.
func (o *PipelineSshPublicKey) SetSha256Fingerprint(v string) {
	o.Sha256Fingerprint = &v
}

func (o PipelineSshPublicKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedObject, errObject := json.Marshal(o.Object)
	if errObject != nil {
		return []byte{}, errObject
	}
	errObject = json.Unmarshal([]byte(serializedObject), &toSerialize)
	if errObject != nil {
		return []byte{}, errObject
	}
	if o.KeyType != nil {
		toSerialize["key_type"] = o.KeyType
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Md5Fingerprint != nil {
		toSerialize["md5_fingerprint"] = o.Md5Fingerprint
	}
	if o.Sha256Fingerprint != nil {
		toSerialize["sha256_fingerprint"] = o.Sha256Fingerprint
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineSshPublicKey struct {
	value *PipelineSshPublicKey
	isSet bool
}

func (v NullablePipelineSshPublicKey) Get() *PipelineSshPublicKey {
	return v.value
}

func (v *NullablePipelineSshPublicKey) Set(val *PipelineSshPublicKey) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineSshPublicKey) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineSshPublicKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineSshPublicKey(val *PipelineSshPublicKey) *NullablePipelineSshPublicKey {
	return &NullablePipelineSshPublicKey{value: val, isSet: true}
}

func (v NullablePipelineSshPublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineSshPublicKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



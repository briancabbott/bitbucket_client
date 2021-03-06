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
	"time"
)

// PipelineCache struct for PipelineCache
type PipelineCache struct {
	Object
	// The UUID identifying the pipeline cache.
	Uuid *string `json:"uuid,omitempty"`
	// The UUID of the pipeline that created the cache.
	PipelineUuid *string `json:"pipeline_uuid,omitempty"`
	// The uuid of the step that created the cache.
	StepUuid *string `json:"step_uuid,omitempty"`
	// The name of the cache.
	Name *string `json:"name,omitempty"`
	// The path where the cache contents were retrieved from.
	Path *string `json:"path,omitempty"`
	// The size of the file containing the archive of the cache.
	FileSizeBytes *int32 `json:"file_size_bytes,omitempty"`
	// The timestamp when the cache was created.
	CreatedOn *time.Time `json:"created_on,omitempty"`
}

// NewPipelineCache instantiates a new PipelineCache object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineCache(type_ string) *PipelineCache {
	this := PipelineCache{}
	this.Type = type_
	return &this
}

// NewPipelineCacheWithDefaults instantiates a new PipelineCache object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineCacheWithDefaults() *PipelineCache {
	this := PipelineCache{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *PipelineCache) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineCache) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *PipelineCache) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *PipelineCache) SetUuid(v string) {
	o.Uuid = &v
}

// GetPipelineUuid returns the PipelineUuid field value if set, zero value otherwise.
func (o *PipelineCache) GetPipelineUuid() string {
	if o == nil || o.PipelineUuid == nil {
		var ret string
		return ret
	}
	return *o.PipelineUuid
}

// GetPipelineUuidOk returns a tuple with the PipelineUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineCache) GetPipelineUuidOk() (*string, bool) {
	if o == nil || o.PipelineUuid == nil {
		return nil, false
	}
	return o.PipelineUuid, true
}

// HasPipelineUuid returns a boolean if a field has been set.
func (o *PipelineCache) HasPipelineUuid() bool {
	if o != nil && o.PipelineUuid != nil {
		return true
	}

	return false
}

// SetPipelineUuid gets a reference to the given string and assigns it to the PipelineUuid field.
func (o *PipelineCache) SetPipelineUuid(v string) {
	o.PipelineUuid = &v
}

// GetStepUuid returns the StepUuid field value if set, zero value otherwise.
func (o *PipelineCache) GetStepUuid() string {
	if o == nil || o.StepUuid == nil {
		var ret string
		return ret
	}
	return *o.StepUuid
}

// GetStepUuidOk returns a tuple with the StepUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineCache) GetStepUuidOk() (*string, bool) {
	if o == nil || o.StepUuid == nil {
		return nil, false
	}
	return o.StepUuid, true
}

// HasStepUuid returns a boolean if a field has been set.
func (o *PipelineCache) HasStepUuid() bool {
	if o != nil && o.StepUuid != nil {
		return true
	}

	return false
}

// SetStepUuid gets a reference to the given string and assigns it to the StepUuid field.
func (o *PipelineCache) SetStepUuid(v string) {
	o.StepUuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PipelineCache) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineCache) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PipelineCache) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PipelineCache) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PipelineCache) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineCache) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PipelineCache) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *PipelineCache) SetPath(v string) {
	o.Path = &v
}

// GetFileSizeBytes returns the FileSizeBytes field value if set, zero value otherwise.
func (o *PipelineCache) GetFileSizeBytes() int32 {
	if o == nil || o.FileSizeBytes == nil {
		var ret int32
		return ret
	}
	return *o.FileSizeBytes
}

// GetFileSizeBytesOk returns a tuple with the FileSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineCache) GetFileSizeBytesOk() (*int32, bool) {
	if o == nil || o.FileSizeBytes == nil {
		return nil, false
	}
	return o.FileSizeBytes, true
}

// HasFileSizeBytes returns a boolean if a field has been set.
func (o *PipelineCache) HasFileSizeBytes() bool {
	if o != nil && o.FileSizeBytes != nil {
		return true
	}

	return false
}

// SetFileSizeBytes gets a reference to the given int32 and assigns it to the FileSizeBytes field.
func (o *PipelineCache) SetFileSizeBytes(v int32) {
	o.FileSizeBytes = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *PipelineCache) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineCache) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *PipelineCache) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *PipelineCache) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

func (o PipelineCache) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedObject, errObject := json.Marshal(o.Object)
	if errObject != nil {
		return []byte{}, errObject
	}
	errObject = json.Unmarshal([]byte(serializedObject), &toSerialize)
	if errObject != nil {
		return []byte{}, errObject
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.PipelineUuid != nil {
		toSerialize["pipeline_uuid"] = o.PipelineUuid
	}
	if o.StepUuid != nil {
		toSerialize["step_uuid"] = o.StepUuid
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.FileSizeBytes != nil {
		toSerialize["file_size_bytes"] = o.FileSizeBytes
	}
	if o.CreatedOn != nil {
		toSerialize["created_on"] = o.CreatedOn
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineCache struct {
	value *PipelineCache
	isSet bool
}

func (v NullablePipelineCache) Get() *PipelineCache {
	return v.value
}

func (v *NullablePipelineCache) Set(val *PipelineCache) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineCache) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineCache) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineCache(val *PipelineCache) *NullablePipelineCache {
	return &NullablePipelineCache{value: val, isSet: true}
}

func (v NullablePipelineCache) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineCache) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



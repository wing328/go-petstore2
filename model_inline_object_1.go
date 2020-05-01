/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"os"
)

// InlineObject1 struct for InlineObject1
type InlineObject1 struct {
	// Additional data to pass to server
	AdditionalMetadata *string `json:"additionalMetadata,omitempty"`
	// file to upload
	File **os.File `json:"file,omitempty"`
}

// NewInlineObject1 instantiates a new InlineObject1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1() *InlineObject1 {
	this := InlineObject1{}
	return &this
}

// NewInlineObject1WithDefaults instantiates a new InlineObject1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1WithDefaults() *InlineObject1 {
	this := InlineObject1{}
	return &this
}

// GetAdditionalMetadata returns the AdditionalMetadata field value if set, zero value otherwise.
func (o *InlineObject1) GetAdditionalMetadata() string {
	if o == nil || o.AdditionalMetadata == nil {
		var ret string
		return ret
	}
	return *o.AdditionalMetadata
}

// GetAdditionalMetadataOk returns a tuple with the AdditionalMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetAdditionalMetadataOk() (*string, bool) {
	if o == nil || o.AdditionalMetadata == nil {
		return nil, false
	}
	return o.AdditionalMetadata, true
}

// HasAdditionalMetadata returns a boolean if a field has been set.
func (o *InlineObject1) HasAdditionalMetadata() bool {
	if o != nil && o.AdditionalMetadata != nil {
		return true
	}

	return false
}

// SetAdditionalMetadata gets a reference to the given string and assigns it to the AdditionalMetadata field.
func (o *InlineObject1) SetAdditionalMetadata(v string) {
	o.AdditionalMetadata = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *InlineObject1) GetFile() *os.File {
	if o == nil || o.File == nil {
		var ret *os.File
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetFileOk() (**os.File, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *InlineObject1) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given *os.File and assigns it to the File field.
func (o *InlineObject1) SetFile(v *os.File) {
	o.File = &v
}

func (o InlineObject1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalMetadata != nil {
		toSerialize["additionalMetadata"] = o.AdditionalMetadata
	}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1 struct {
	value *InlineObject1
	isSet bool
}

func (v NullableInlineObject1) Get() *InlineObject1 {
	return v.value
}

func (v *NullableInlineObject1) Set(val *InlineObject1) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1(val *InlineObject1) *NullableInlineObject1 {
	return &NullableInlineObject1{value: val, isSet: true}
}

func (v NullableInlineObject1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
FirstQuadrant API

The FirstQuadrant API is used to interact with FirstQuadrant programmatically. We also have SDKs available (coming soon).

API version: 0.0.0
Contact: inquiry@firstquadrant.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Success type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Success{}

// Success struct for Success
type Success struct {
	Success interface{} `json:"success"`
}

// NewSuccess instantiates a new Success object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccess(success interface{}) *Success {
	this := Success{}
	this.Success = success
	return &this
}

// NewSuccessWithDefaults instantiates a new Success object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessWithDefaults() *Success {
	this := Success{}
	return &this
}

// GetSuccess returns the Success field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Success) GetSuccess() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Success) GetSuccessOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *Success) SetSuccess(v interface{}) {
	o.Success = v
}

func (o Success) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Success) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableSuccess struct {
	value *Success
	isSet bool
}

func (v NullableSuccess) Get() *Success {
	return v.value
}

func (v *NullableSuccess) Set(val *Success) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccess(val *Success) *NullableSuccess {
	return &NullableSuccess{value: val, isSet: true}
}

func (v NullableSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



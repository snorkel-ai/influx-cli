/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LineProtocolError struct for LineProtocolError
type LineProtocolError struct {
	Code LineProtocolErrorCode `json:"code"`
	// Message is a human-readable message.
	Message string `json:"message"`
	// Op describes the logical code operation during error. Useful for debugging.
	Op string `json:"op"`
	// Err is a stack of errors that occurred during processing of the request. Useful for debugging.
	Err string `json:"err"`
	// First line within sent body containing malformed data
	Line *int32 `json:"line,omitempty"`
}

// NewLineProtocolError instantiates a new LineProtocolError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineProtocolError(code LineProtocolErrorCode, message string, op string, err string) *LineProtocolError {
	this := LineProtocolError{}
	this.Code = code
	this.Message = message
	this.Op = op
	this.Err = err
	return &this
}

// NewLineProtocolErrorWithDefaults instantiates a new LineProtocolError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineProtocolErrorWithDefaults() *LineProtocolError {
	this := LineProtocolError{}
	return &this
}

// GetCode returns the Code field value
func (o *LineProtocolError) GetCode() LineProtocolErrorCode {
	if o == nil {
		var ret LineProtocolErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *LineProtocolError) GetCodeOk() (*LineProtocolErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *LineProtocolError) SetCode(v LineProtocolErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *LineProtocolError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *LineProtocolError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *LineProtocolError) SetMessage(v string) {
	o.Message = v
}

// GetOp returns the Op field value
func (o *LineProtocolError) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *LineProtocolError) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *LineProtocolError) SetOp(v string) {
	o.Op = v
}

// GetErr returns the Err field value
func (o *LineProtocolError) GetErr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Err
}

// GetErrOk returns a tuple with the Err field value
// and a boolean to check if the value has been set.
func (o *LineProtocolError) GetErrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Err, true
}

// SetErr sets field value
func (o *LineProtocolError) SetErr(v string) {
	o.Err = v
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *LineProtocolError) GetLine() int32 {
	if o == nil || o.Line == nil {
		var ret int32
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineProtocolError) GetLineOk() (*int32, bool) {
	if o == nil || o.Line == nil {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *LineProtocolError) HasLine() bool {
	if o != nil && o.Line != nil {
		return true
	}

	return false
}

// SetLine gets a reference to the given int32 and assigns it to the Line field.
func (o *LineProtocolError) SetLine(v int32) {
	o.Line = &v
}

func (o LineProtocolError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["op"] = o.Op
	}
	if true {
		toSerialize["err"] = o.Err
	}
	if o.Line != nil {
		toSerialize["line"] = o.Line
	}
	return json.Marshal(toSerialize)
}

type NullableLineProtocolError struct {
	value *LineProtocolError
	isSet bool
}

func (v NullableLineProtocolError) Get() *LineProtocolError {
	return v.value
}

func (v *NullableLineProtocolError) Set(val *LineProtocolError) {
	v.value = val
	v.isSet = true
}

func (v NullableLineProtocolError) IsSet() bool {
	return v.isSet
}

func (v *NullableLineProtocolError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineProtocolError(val *LineProtocolError) *NullableLineProtocolError {
	return &NullableLineProtocolError{value: val, isSet: true}
}

func (v NullableLineProtocolError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineProtocolError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
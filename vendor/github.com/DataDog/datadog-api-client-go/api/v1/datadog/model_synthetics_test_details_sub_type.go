/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// SyntheticsTestDetailsSubType The sub-type of the Synthetic API test, either `http` or `ssl`.
type SyntheticsTestDetailsSubType string

// List of SyntheticsTestDetailsSubType
const (
	SYNTHETICSTESTDETAILSSUBTYPE_HTTP SyntheticsTestDetailsSubType = "http"
	SYNTHETICSTESTDETAILSSUBTYPE_SSL  SyntheticsTestDetailsSubType = "ssl"
)

func (v *SyntheticsTestDetailsSubType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SyntheticsTestDetailsSubType(value)
	for _, existing := range []SyntheticsTestDetailsSubType{"http", "ssl"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SyntheticsTestDetailsSubType", *v)
}

// Ptr returns reference to SyntheticsTestDetailsSubType value
func (v SyntheticsTestDetailsSubType) Ptr() *SyntheticsTestDetailsSubType {
	return &v
}

type NullableSyntheticsTestDetailsSubType struct {
	value *SyntheticsTestDetailsSubType
	isSet bool
}

func (v NullableSyntheticsTestDetailsSubType) Get() *SyntheticsTestDetailsSubType {
	return v.value
}

func (v *NullableSyntheticsTestDetailsSubType) Set(val *SyntheticsTestDetailsSubType) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsTestDetailsSubType) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsTestDetailsSubType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsTestDetailsSubType(val *SyntheticsTestDetailsSubType) *NullableSyntheticsTestDetailsSubType {
	return &NullableSyntheticsTestDetailsSubType{value: val, isSet: true}
}

func (v NullableSyntheticsTestDetailsSubType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsTestDetailsSubType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
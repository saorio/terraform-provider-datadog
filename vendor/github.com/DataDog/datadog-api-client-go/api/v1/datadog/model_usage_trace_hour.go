/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"time"
)

// UsageTraceHour The hours of trace usage.
type UsageTraceHour struct {
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// Contains the number of Analyzed Spans indexed.
	IndexedEventsCount *int64 `json:"indexed_events_count,omitempty"`
}

// NewUsageTraceHour instantiates a new UsageTraceHour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageTraceHour() *UsageTraceHour {
	this := UsageTraceHour{}
	return &this
}

// NewUsageTraceHourWithDefaults instantiates a new UsageTraceHour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageTraceHourWithDefaults() *UsageTraceHour {
	this := UsageTraceHour{}
	return &this
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageTraceHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageTraceHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageTraceHour) HasHour() bool {
	if o != nil && o.Hour != nil {
		return true
	}

	return false
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageTraceHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetIndexedEventsCount returns the IndexedEventsCount field value if set, zero value otherwise.
func (o *UsageTraceHour) GetIndexedEventsCount() int64 {
	if o == nil || o.IndexedEventsCount == nil {
		var ret int64
		return ret
	}
	return *o.IndexedEventsCount
}

// GetIndexedEventsCountOk returns a tuple with the IndexedEventsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageTraceHour) GetIndexedEventsCountOk() (*int64, bool) {
	if o == nil || o.IndexedEventsCount == nil {
		return nil, false
	}
	return o.IndexedEventsCount, true
}

// HasIndexedEventsCount returns a boolean if a field has been set.
func (o *UsageTraceHour) HasIndexedEventsCount() bool {
	if o != nil && o.IndexedEventsCount != nil {
		return true
	}

	return false
}

// SetIndexedEventsCount gets a reference to the given int64 and assigns it to the IndexedEventsCount field.
func (o *UsageTraceHour) SetIndexedEventsCount(v int64) {
	o.IndexedEventsCount = &v
}

func (o UsageTraceHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hour != nil {
		toSerialize["hour"] = o.Hour
	}
	if o.IndexedEventsCount != nil {
		toSerialize["indexed_events_count"] = o.IndexedEventsCount
	}
	return json.Marshal(toSerialize)
}

type NullableUsageTraceHour struct {
	value *UsageTraceHour
	isSet bool
}

func (v NullableUsageTraceHour) Get() *UsageTraceHour {
	return v.value
}

func (v *NullableUsageTraceHour) Set(val *UsageTraceHour) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageTraceHour) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageTraceHour) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageTraceHour(val *UsageTraceHour) *NullableUsageTraceHour {
	return &NullableUsageTraceHour{value: val, isSet: true}
}

func (v NullableUsageTraceHour) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageTraceHour) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

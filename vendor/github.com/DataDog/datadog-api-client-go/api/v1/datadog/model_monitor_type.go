/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// MonitorType The type of the monitor.
type MonitorType string

// List of MonitorType
const (
	MONITORTYPE_COMPOSITE             MonitorType = "composite"
	MONITORTYPE_EVENT_ALERT           MonitorType = "event alert"
	MONITORTYPE_LOG_ALERT             MonitorType = "log alert"
	MONITORTYPE_METRIC_ALERT          MonitorType = "metric alert"
	MONITORTYPE_PROCESS_ALERT         MonitorType = "process alert"
	MONITORTYPE_QUERY_ALERT           MonitorType = "query alert"
	MONITORTYPE_RUM_ALERT             MonitorType = "rum alert"
	MONITORTYPE_SERVICE_CHECK         MonitorType = "service check"
	MONITORTYPE_SYNTHETICS_ALERT      MonitorType = "synthetics alert"
	MONITORTYPE_TRACE_ANALYTICS_ALERT MonitorType = "trace-analytics alert"
)

// Ptr returns reference to MonitorType value
func (v MonitorType) Ptr() *MonitorType {
	return &v
}

type NullableMonitorType struct {
	value *MonitorType
	isSet bool
}

func (v NullableMonitorType) Get() *MonitorType {
	return v.value
}

func (v *NullableMonitorType) Set(val *MonitorType) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorType) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorType(val *MonitorType) *NullableMonitorType {
	return &NullableMonitorType{value: val, isSet: true}
}

func (v NullableMonitorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

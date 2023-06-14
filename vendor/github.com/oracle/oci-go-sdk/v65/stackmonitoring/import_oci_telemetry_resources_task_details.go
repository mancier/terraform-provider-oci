// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ImportOciTelemetryResourcesTaskDetails Request details for importing resources from Telemetry like resources from OCI Native Services and prometheus.
type ImportOciTelemetryResourcesTaskDetails struct {

	// Name space to be used for OCI Native service resources discovery.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The resource group to use while fetching metrics from telemetry.
	// If not specified, resource group will be skipped in the list metrics request.
	ResourceGroup *string `mandatory:"false" json:"resourceGroup"`

	// List of comma separated properties to be used while getting resource data from telemetry.
	// Resources with matching dimensions are only fetched from telemetry.
	IdentifyingProperties []string `mandatory:"false" json:"identifyingProperties"`

	// Source from where the metrics pushed to telemetry.
	// Possible values:
	//   * OCI_TELEMETRY_NATIVE      - The metrics are pushed to telemetry from OCI Native Services.
	//   * OCI_TELEMETRY_PROMETHEUS  - The metrics are pushed to telemetry from Prometheus.
	Source ImportOciTelemetryResourcesTaskDetailsSourceEnum `mandatory:"true" json:"source"`
}

func (m ImportOciTelemetryResourcesTaskDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImportOciTelemetryResourcesTaskDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingImportOciTelemetryResourcesTaskDetailsSourceEnum(string(m.Source)); !ok && m.Source != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Source: %s. Supported values are: %s.", m.Source, strings.Join(GetImportOciTelemetryResourcesTaskDetailsSourceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ImportOciTelemetryResourcesTaskDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeImportOciTelemetryResourcesTaskDetails ImportOciTelemetryResourcesTaskDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeImportOciTelemetryResourcesTaskDetails
	}{
		"IMPORT_OCI_TELEMETRY_RESOURCES",
		(MarshalTypeImportOciTelemetryResourcesTaskDetails)(m),
	}

	return json.Marshal(&s)
}

// ImportOciTelemetryResourcesTaskDetailsSourceEnum Enum with underlying type: string
type ImportOciTelemetryResourcesTaskDetailsSourceEnum string

// Set of constants representing the allowable values for ImportOciTelemetryResourcesTaskDetailsSourceEnum
const (
	ImportOciTelemetryResourcesTaskDetailsSourceNative     ImportOciTelemetryResourcesTaskDetailsSourceEnum = "OCI_TELEMETRY_NATIVE"
	ImportOciTelemetryResourcesTaskDetailsSourcePrometheus ImportOciTelemetryResourcesTaskDetailsSourceEnum = "OCI_TELEMETRY_PROMETHEUS"
)

var mappingImportOciTelemetryResourcesTaskDetailsSourceEnum = map[string]ImportOciTelemetryResourcesTaskDetailsSourceEnum{
	"OCI_TELEMETRY_NATIVE":     ImportOciTelemetryResourcesTaskDetailsSourceNative,
	"OCI_TELEMETRY_PROMETHEUS": ImportOciTelemetryResourcesTaskDetailsSourcePrometheus,
}

var mappingImportOciTelemetryResourcesTaskDetailsSourceEnumLowerCase = map[string]ImportOciTelemetryResourcesTaskDetailsSourceEnum{
	"oci_telemetry_native":     ImportOciTelemetryResourcesTaskDetailsSourceNative,
	"oci_telemetry_prometheus": ImportOciTelemetryResourcesTaskDetailsSourcePrometheus,
}

// GetImportOciTelemetryResourcesTaskDetailsSourceEnumValues Enumerates the set of values for ImportOciTelemetryResourcesTaskDetailsSourceEnum
func GetImportOciTelemetryResourcesTaskDetailsSourceEnumValues() []ImportOciTelemetryResourcesTaskDetailsSourceEnum {
	values := make([]ImportOciTelemetryResourcesTaskDetailsSourceEnum, 0)
	for _, v := range mappingImportOciTelemetryResourcesTaskDetailsSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetImportOciTelemetryResourcesTaskDetailsSourceEnumStringValues Enumerates the set of values in String for ImportOciTelemetryResourcesTaskDetailsSourceEnum
func GetImportOciTelemetryResourcesTaskDetailsSourceEnumStringValues() []string {
	return []string{
		"OCI_TELEMETRY_NATIVE",
		"OCI_TELEMETRY_PROMETHEUS",
	}
}

// GetMappingImportOciTelemetryResourcesTaskDetailsSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImportOciTelemetryResourcesTaskDetailsSourceEnum(val string) (ImportOciTelemetryResourcesTaskDetailsSourceEnum, bool) {
	enum, ok := mappingImportOciTelemetryResourcesTaskDetailsSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

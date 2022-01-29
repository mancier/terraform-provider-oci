// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v56/common"
	"strings"
)

// AmdMilanBmGpuPlatformConfig The platform configuration of a bare metal instance that uses a GPU shape on the AMD Milan platform.
type AmdMilanBmGpuPlatformConfig struct {

	// Whether Secure Boot is enabled on the instance.
	IsSecureBootEnabled *bool `mandatory:"false" json:"isSecureBootEnabled"`

	// Whether the Trusted Platform Module (TPM) is enabled on the instance.
	IsTrustedPlatformModuleEnabled *bool `mandatory:"false" json:"isTrustedPlatformModuleEnabled"`

	// Whether the Measured Boot feature is enabled on the instance.
	IsMeasuredBootEnabled *bool `mandatory:"false" json:"isMeasuredBootEnabled"`

	// Whether the instance is a confidential instance. If this value is `true`, the instance is a confidential instance. The default value is `false`.
	IsMemoryEncryptionEnabled *bool `mandatory:"false" json:"isMemoryEncryptionEnabled"`

	// Whether symmetric multi-threading is enabled on the instance.
	IsSymmetricMultiThreadingEnabled *bool `mandatory:"false" json:"isSymmetricMultiThreadingEnabled"`

	// Whether the Access Control Service is enabled on the instance. When enabled,
	// the platform can enforce PCIe device isolation, required for VFIO device passthrough.
	IsAccessControlServiceEnabled *bool `mandatory:"false" json:"isAccessControlServiceEnabled"`

	// Whether virtualization instructions are available.
	AreVirtualInstructionsEnabled *bool `mandatory:"false" json:"areVirtualInstructionsEnabled"`

	// Whether the input-output memory management unit is enabled.
	IsInputOutputMemoryManagementUnitEnabled *bool `mandatory:"false" json:"isInputOutputMemoryManagementUnitEnabled"`

	// The number of NUMA nodes per socket (NPS).
	NumaNodesPerSocket AmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnum `mandatory:"false" json:"numaNodesPerSocket,omitempty"`
}

//GetIsSecureBootEnabled returns IsSecureBootEnabled
func (m AmdMilanBmGpuPlatformConfig) GetIsSecureBootEnabled() *bool {
	return m.IsSecureBootEnabled
}

//GetIsTrustedPlatformModuleEnabled returns IsTrustedPlatformModuleEnabled
func (m AmdMilanBmGpuPlatformConfig) GetIsTrustedPlatformModuleEnabled() *bool {
	return m.IsTrustedPlatformModuleEnabled
}

//GetIsMeasuredBootEnabled returns IsMeasuredBootEnabled
func (m AmdMilanBmGpuPlatformConfig) GetIsMeasuredBootEnabled() *bool {
	return m.IsMeasuredBootEnabled
}

//GetIsMemoryEncryptionEnabled returns IsMemoryEncryptionEnabled
func (m AmdMilanBmGpuPlatformConfig) GetIsMemoryEncryptionEnabled() *bool {
	return m.IsMemoryEncryptionEnabled
}

func (m AmdMilanBmGpuPlatformConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AmdMilanBmGpuPlatformConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingAmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnum[string(m.NumaNodesPerSocket)]; !ok && m.NumaNodesPerSocket != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NumaNodesPerSocket: %s. Supported values are: %s.", m.NumaNodesPerSocket, strings.Join(GetAmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AmdMilanBmGpuPlatformConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAmdMilanBmGpuPlatformConfig AmdMilanBmGpuPlatformConfig
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAmdMilanBmGpuPlatformConfig
	}{
		"AMD_MILAN_BM_GPU",
		(MarshalTypeAmdMilanBmGpuPlatformConfig)(m),
	}

	return json.Marshal(&s)
}

// AmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnum Enum with underlying type: string
type AmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnum string

// Set of constants representing the allowable values for AmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnum
const (
	AmdMilanBmGpuPlatformConfigNumaNodesPerSocketNps0 AmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnum = "NPS0"
	AmdMilanBmGpuPlatformConfigNumaNodesPerSocketNps1 AmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnum = "NPS1"
	AmdMilanBmGpuPlatformConfigNumaNodesPerSocketNps2 AmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnum = "NPS2"
	AmdMilanBmGpuPlatformConfigNumaNodesPerSocketNps4 AmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnum = "NPS4"
)

var mappingAmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnum = map[string]AmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnum{
	"NPS0": AmdMilanBmGpuPlatformConfigNumaNodesPerSocketNps0,
	"NPS1": AmdMilanBmGpuPlatformConfigNumaNodesPerSocketNps1,
	"NPS2": AmdMilanBmGpuPlatformConfigNumaNodesPerSocketNps2,
	"NPS4": AmdMilanBmGpuPlatformConfigNumaNodesPerSocketNps4,
}

// GetAmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnumValues Enumerates the set of values for AmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnum
func GetAmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnumValues() []AmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnum {
	values := make([]AmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnum, 0)
	for _, v := range mappingAmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnum {
		values = append(values, v)
	}
	return values
}

// GetAmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnumStringValues Enumerates the set of values in String for AmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnum
func GetAmdMilanBmGpuPlatformConfigNumaNodesPerSocketEnumStringValues() []string {
	return []string{
		"NPS0",
		"NPS1",
		"NPS2",
		"NPS4",
	}
}

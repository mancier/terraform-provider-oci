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

// IntelIcelakeBmPlatformConfig The platform configuration of a bare metal instance that uses the Intel Icelake platform.
type IntelIcelakeBmPlatformConfig struct {

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

	// Whether the input-output memory management unit is enabled.
	IsInputOutputMemoryManagementUnitEnabled *bool `mandatory:"false" json:"isInputOutputMemoryManagementUnitEnabled"`

	// The percentage of cores enabled.
	PercentageOfCoresEnabled *int `mandatory:"false" json:"percentageOfCoresEnabled"`

	// The number of NUMA nodes per socket (NPS).
	NumaNodesPerSocket IntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum `mandatory:"false" json:"numaNodesPerSocket,omitempty"`
}

//GetIsSecureBootEnabled returns IsSecureBootEnabled
func (m IntelIcelakeBmPlatformConfig) GetIsSecureBootEnabled() *bool {
	return m.IsSecureBootEnabled
}

//GetIsTrustedPlatformModuleEnabled returns IsTrustedPlatformModuleEnabled
func (m IntelIcelakeBmPlatformConfig) GetIsTrustedPlatformModuleEnabled() *bool {
	return m.IsTrustedPlatformModuleEnabled
}

//GetIsMeasuredBootEnabled returns IsMeasuredBootEnabled
func (m IntelIcelakeBmPlatformConfig) GetIsMeasuredBootEnabled() *bool {
	return m.IsMeasuredBootEnabled
}

//GetIsMemoryEncryptionEnabled returns IsMemoryEncryptionEnabled
func (m IntelIcelakeBmPlatformConfig) GetIsMemoryEncryptionEnabled() *bool {
	return m.IsMemoryEncryptionEnabled
}

func (m IntelIcelakeBmPlatformConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IntelIcelakeBmPlatformConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingIntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum[string(m.NumaNodesPerSocket)]; !ok && m.NumaNodesPerSocket != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NumaNodesPerSocket: %s. Supported values are: %s.", m.NumaNodesPerSocket, strings.Join(GetIntelIcelakeBmPlatformConfigNumaNodesPerSocketEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m IntelIcelakeBmPlatformConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeIntelIcelakeBmPlatformConfig IntelIcelakeBmPlatformConfig
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeIntelIcelakeBmPlatformConfig
	}{
		"INTEL_ICELAKE_BM",
		(MarshalTypeIntelIcelakeBmPlatformConfig)(m),
	}

	return json.Marshal(&s)
}

// IntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum Enum with underlying type: string
type IntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum string

// Set of constants representing the allowable values for IntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum
const (
	IntelIcelakeBmPlatformConfigNumaNodesPerSocketNps1 IntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum = "NPS1"
	IntelIcelakeBmPlatformConfigNumaNodesPerSocketNps2 IntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum = "NPS2"
)

var mappingIntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum = map[string]IntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum{
	"NPS1": IntelIcelakeBmPlatformConfigNumaNodesPerSocketNps1,
	"NPS2": IntelIcelakeBmPlatformConfigNumaNodesPerSocketNps2,
}

// GetIntelIcelakeBmPlatformConfigNumaNodesPerSocketEnumValues Enumerates the set of values for IntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum
func GetIntelIcelakeBmPlatformConfigNumaNodesPerSocketEnumValues() []IntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum {
	values := make([]IntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum, 0)
	for _, v := range mappingIntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum {
		values = append(values, v)
	}
	return values
}

// GetIntelIcelakeBmPlatformConfigNumaNodesPerSocketEnumStringValues Enumerates the set of values in String for IntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum
func GetIntelIcelakeBmPlatformConfigNumaNodesPerSocketEnumStringValues() []string {
	return []string{
		"NPS1",
		"NPS2",
	}
}

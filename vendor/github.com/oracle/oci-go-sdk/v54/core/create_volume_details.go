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
	"github.com/oracle/oci-go-sdk/v54/common"
)

// CreateVolumeDetails The details of the volume to create. For CreateVolume operation, this field is required in the request,
// see CreateVolume.
type CreateVolumeDetails struct {

	// The OCID of the compartment that contains the volume.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The availability domain of the volume. Omissible for cloning a volume. The new volume will be created in the availability domain of the source volume.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// If provided, specifies the ID of the volume backup policy to assign to the newly
	// created volume. If omitted, no policy will be assigned.
	BackupPolicyId *string `mandatory:"false" json:"backupPolicyId"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The OCID of the Key Management key to assign as the master encryption key
	// for the volume.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The number of volume performance units (VPUs) that will be applied to this volume per GB,
	// representing the Block Volume service's elastic performance options.
	// See Block Volume Elastic Performance (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeelasticperformance.htm) for more information.
	// Allowed values:
	//   * `0`: Represents Lower Cost option.
	//   * `10`: Represents Balanced option.
	//   * `20`: Represents Higher Performance option.
	VpusPerGB *int64 `mandatory:"false" json:"vpusPerGB"`

	// The size of the volume in GBs.
	SizeInGBs *int64 `mandatory:"false" json:"sizeInGBs"`

	// The size of the volume in MBs. The value must be a multiple of 1024.
	// This field is deprecated. Use sizeInGBs instead.
	SizeInMBs *int64 `mandatory:"false" json:"sizeInMBs"`

	SourceDetails VolumeSourceDetails `mandatory:"false" json:"sourceDetails"`

	// The OCID of the volume backup from which the data should be restored on the newly created volume.
	// This field is deprecated. Use the sourceDetails field instead to specify the
	// backup for the volume.
	VolumeBackupId *string `mandatory:"false" json:"volumeBackupId"`

	// Specifies whether the auto-tune performance is enabled for this volume.
	IsAutoTuneEnabled *bool `mandatory:"false" json:"isAutoTuneEnabled"`

	// The list of block volume replicas to be enabled for this volume
	// in the specified destination availability domains.
	BlockVolumeReplicas []BlockVolumeReplicaDetails `mandatory:"false" json:"blockVolumeReplicas"`
}

func (m CreateVolumeDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateVolumeDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		AvailabilityDomain  *string                           `json:"availabilityDomain"`
		BackupPolicyId      *string                           `json:"backupPolicyId"`
		DefinedTags         map[string]map[string]interface{} `json:"definedTags"`
		DisplayName         *string                           `json:"displayName"`
		FreeformTags        map[string]string                 `json:"freeformTags"`
		KmsKeyId            *string                           `json:"kmsKeyId"`
		VpusPerGB           *int64                            `json:"vpusPerGB"`
		SizeInGBs           *int64                            `json:"sizeInGBs"`
		SizeInMBs           *int64                            `json:"sizeInMBs"`
		SourceDetails       volumesourcedetails               `json:"sourceDetails"`
		VolumeBackupId      *string                           `json:"volumeBackupId"`
		IsAutoTuneEnabled   *bool                             `json:"isAutoTuneEnabled"`
		BlockVolumeReplicas []BlockVolumeReplicaDetails       `json:"blockVolumeReplicas"`
		CompartmentId       *string                           `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.AvailabilityDomain = model.AvailabilityDomain

	m.BackupPolicyId = model.BackupPolicyId

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.KmsKeyId = model.KmsKeyId

	m.VpusPerGB = model.VpusPerGB

	m.SizeInGBs = model.SizeInGBs

	m.SizeInMBs = model.SizeInMBs

	nn, e = model.SourceDetails.UnmarshalPolymorphicJSON(model.SourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SourceDetails = nn.(VolumeSourceDetails)
	} else {
		m.SourceDetails = nil
	}

	m.VolumeBackupId = model.VolumeBackupId

	m.IsAutoTuneEnabled = model.IsAutoTuneEnabled

	m.BlockVolumeReplicas = make([]BlockVolumeReplicaDetails, len(model.BlockVolumeReplicas))
	for i, n := range model.BlockVolumeReplicas {
		m.BlockVolumeReplicas[i] = n
	}

	m.CompartmentId = model.CompartmentId

	return
}

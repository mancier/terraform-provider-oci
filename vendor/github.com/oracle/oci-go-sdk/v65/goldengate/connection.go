// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Connection Represents the metadata description of a connection used by deployments in the same compartment.
type Connection interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the connection being
	// referenced.
	GetId() *string

	// An object's Display Name.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	GetCompartmentId() *string

	// Possible lifecycle states for connection.
	GetLifecycleState() ConnectionLifecycleStateEnum

	// The time the resource was created. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	GetTimeCreated() *common.SDKTime

	// The time the resource was last updated. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	GetTimeUpdated() *common.SDKTime

	// Metadata about this specific object.
	GetDescription() *string

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// The system tags associated with this resource, if any. The system tags are set by Oracle
	// Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more
	// information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	GetSystemTags() map[string]map[string]interface{}

	// Describes the object's current state in detail. For example, it can be used to provide
	// actionable information for a resource in a Failed state.
	GetLifecycleDetails() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the customer vault being
	// referenced.
	// If provided, this will reference a vault which the customer will be required to ensure
	// the policies are established to permit the GoldenGate Service to manage secrets contained
	// within this vault.
	GetVaultId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the customer "Master" key being
	// referenced.
	// If provided, this will reference a key which the customer will be required to ensure
	// the policies are established to permit the GoldenGate Service to utilize this key to
	// manage secrets.
	GetKeyId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet being referenced.
	GetSubnetId() *string

	// List of ingress IP addresses, from where the GoldenGate deployment connects to this connection's privateIp.
	GetIngressIps() []IngressIpDetails

	// An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
	GetNsgIds() []string
}

type connection struct {
	JsonData         []byte
	Id               *string                           `mandatory:"true" json:"id"`
	DisplayName      *string                           `mandatory:"true" json:"displayName"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	LifecycleState   ConnectionLifecycleStateEnum      `mandatory:"true" json:"lifecycleState"`
	TimeCreated      *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated      *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	Description      *string                           `mandatory:"false" json:"description"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	LifecycleDetails *string                           `mandatory:"false" json:"lifecycleDetails"`
	VaultId          *string                           `mandatory:"false" json:"vaultId"`
	KeyId            *string                           `mandatory:"false" json:"keyId"`
	SubnetId         *string                           `mandatory:"false" json:"subnetId"`
	IngressIps       []IngressIpDetails                `mandatory:"false" json:"ingressIps"`
	NsgIds           []string                          `mandatory:"false" json:"nsgIds"`
	ConnectionType   string                            `json:"connectionType"`
}

// UnmarshalJSON unmarshals json
func (m *connection) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconnection connection
	s := struct {
		Model Unmarshalerconnection
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.VaultId = s.Model.VaultId
	m.KeyId = s.Model.KeyId
	m.SubnetId = s.Model.SubnetId
	m.IngressIps = s.Model.IngressIps
	m.NsgIds = s.Model.NsgIds
	m.ConnectionType = s.Model.ConnectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *connection) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectionType {
	case "KAFKA":
		mm := KafkaConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "POSTGRESQL":
		mm := PostgresqlConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_OBJECT_STORAGE":
		mm := OciObjectStorageConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KAFKA_SCHEMA_REGISTRY":
		mm := KafkaSchemaRegistryConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GOLDENGATE":
		mm := GoldenGateConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL":
		mm := MysqlConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE":
		mm := OracleConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AZURE_DATA_LAKE_STORAGE":
		mm := AzureDataLakeStorageConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AZURE_SYNAPSE_ANALYTICS":
		mm := AzureSynapseConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Connection: %s.", m.ConnectionType)
		return *m, nil
	}
}

//GetId returns Id
func (m connection) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m connection) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m connection) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetLifecycleState returns LifecycleState
func (m connection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

//GetTimeCreated returns TimeCreated
func (m connection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m connection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetDescription returns Description
func (m connection) GetDescription() *string {
	return m.Description
}

//GetFreeformTags returns FreeformTags
func (m connection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m connection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m connection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

//GetLifecycleDetails returns LifecycleDetails
func (m connection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetVaultId returns VaultId
func (m connection) GetVaultId() *string {
	return m.VaultId
}

//GetKeyId returns KeyId
func (m connection) GetKeyId() *string {
	return m.KeyId
}

//GetSubnetId returns SubnetId
func (m connection) GetSubnetId() *string {
	return m.SubnetId
}

//GetIngressIps returns IngressIps
func (m connection) GetIngressIps() []IngressIpDetails {
	return m.IngressIps
}

//GetNsgIds returns NsgIds
func (m connection) GetNsgIds() []string {
	return m.NsgIds
}

func (m connection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m connection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConnectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConnectionLifecycleStateEnum Enum with underlying type: string
type ConnectionLifecycleStateEnum string

// Set of constants representing the allowable values for ConnectionLifecycleStateEnum
const (
	ConnectionLifecycleStateCreating ConnectionLifecycleStateEnum = "CREATING"
	ConnectionLifecycleStateUpdating ConnectionLifecycleStateEnum = "UPDATING"
	ConnectionLifecycleStateActive   ConnectionLifecycleStateEnum = "ACTIVE"
	ConnectionLifecycleStateDeleting ConnectionLifecycleStateEnum = "DELETING"
	ConnectionLifecycleStateDeleted  ConnectionLifecycleStateEnum = "DELETED"
	ConnectionLifecycleStateFailed   ConnectionLifecycleStateEnum = "FAILED"
)

var mappingConnectionLifecycleStateEnum = map[string]ConnectionLifecycleStateEnum{
	"CREATING": ConnectionLifecycleStateCreating,
	"UPDATING": ConnectionLifecycleStateUpdating,
	"ACTIVE":   ConnectionLifecycleStateActive,
	"DELETING": ConnectionLifecycleStateDeleting,
	"DELETED":  ConnectionLifecycleStateDeleted,
	"FAILED":   ConnectionLifecycleStateFailed,
}

var mappingConnectionLifecycleStateEnumLowerCase = map[string]ConnectionLifecycleStateEnum{
	"creating": ConnectionLifecycleStateCreating,
	"updating": ConnectionLifecycleStateUpdating,
	"active":   ConnectionLifecycleStateActive,
	"deleting": ConnectionLifecycleStateDeleting,
	"deleted":  ConnectionLifecycleStateDeleted,
	"failed":   ConnectionLifecycleStateFailed,
}

// GetConnectionLifecycleStateEnumValues Enumerates the set of values for ConnectionLifecycleStateEnum
func GetConnectionLifecycleStateEnumValues() []ConnectionLifecycleStateEnum {
	values := make([]ConnectionLifecycleStateEnum, 0)
	for _, v := range mappingConnectionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectionLifecycleStateEnumStringValues Enumerates the set of values in String for ConnectionLifecycleStateEnum
func GetConnectionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingConnectionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectionLifecycleStateEnum(val string) (ConnectionLifecycleStateEnum, bool) {
	enum, ok := mappingConnectionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

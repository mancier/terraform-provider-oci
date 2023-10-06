// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PipelineCustomScriptStepDetails The type of step where user provides the step artifact to be executed on an execution engine managed by the pipelines service.
type PipelineCustomScriptStepDetails struct {

	// The name of the step. It must be unique within the pipeline. This is used to create the pipeline DAG.
	StepName *string `mandatory:"true" json:"stepName"`

	// A short description of the step.
	Description *string `mandatory:"false" json:"description"`

	// The list of step names this current step depends on for execution.
	DependsOn []string `mandatory:"false" json:"dependsOn"`

	StepConfigurationDetails *PipelineStepConfigurationDetails `mandatory:"false" json:"stepConfigurationDetails"`

	StepInfrastructureConfigurationDetails *PipelineInfrastructureConfigurationDetails `mandatory:"false" json:"stepInfrastructureConfigurationDetails"`

	// A flag to indicate whether the artifact has been uploaded for this step or not.
	IsArtifactUploaded *bool `mandatory:"false" json:"isArtifactUploaded"`
}

//GetStepName returns StepName
func (m PipelineCustomScriptStepDetails) GetStepName() *string {
	return m.StepName
}

//GetDescription returns Description
func (m PipelineCustomScriptStepDetails) GetDescription() *string {
	return m.Description
}

//GetDependsOn returns DependsOn
func (m PipelineCustomScriptStepDetails) GetDependsOn() []string {
	return m.DependsOn
}

//GetStepConfigurationDetails returns StepConfigurationDetails
func (m PipelineCustomScriptStepDetails) GetStepConfigurationDetails() *PipelineStepConfigurationDetails {
	return m.StepConfigurationDetails
}

func (m PipelineCustomScriptStepDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PipelineCustomScriptStepDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PipelineCustomScriptStepDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePipelineCustomScriptStepDetails PipelineCustomScriptStepDetails
	s := struct {
		DiscriminatorParam string `json:"stepType"`
		MarshalTypePipelineCustomScriptStepDetails
	}{
		"CUSTOM_SCRIPT",
		(MarshalTypePipelineCustomScriptStepDetails)(m),
	}

	return json.Marshal(&s)
}

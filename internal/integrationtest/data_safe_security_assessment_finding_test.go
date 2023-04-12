// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafesecurityAssessmentFindingDataSourceRepresentation = map[string]interface{}{
		"security_assessment_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_security_assessment.test_security_assessment_findings.id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"finding_key":               acctest.Representation{RepType: acctest.Optional, Create: `findingKey`},
		"severity":                  acctest.Representation{RepType: acctest.Optional, Create: `HIGH`},
	}
	DataSafeSecurityAssessmentFindingReferencesDataSourceRepresentation = map[string]interface{}{}

	DataSafeSecurityAssessmentFindingResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment_findings", acctest.Required, acctest.Create, securityAssessmentRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityAssessmentFindingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityAssessmentFindingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	datasourceName := "data.oci_data_safe_security_assessment_findings.test_security_assessment_findings"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_assessment_findings", "test_security_assessment_findings", acctest.Required, acctest.Create, DataSafesecurityAssessmentFindingDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSecurityAssessmentFindingResourceConfig + targetIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "security_assessment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.0.key"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.0.remarks"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.0.severity"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.0.summary"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.0.target_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.0.title"),
			),
		},
	})
}

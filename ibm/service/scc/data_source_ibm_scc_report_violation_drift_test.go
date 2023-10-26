// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccReportViolationDriftDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccReportViolationDriftDataSourceConfigBasic(acc.SccInstanceID, acc.SccReportID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_violation_drift.scc_report_violation_drift_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_violation_drift.scc_report_violation_drift_instance", "report_id"),
				),
			},
		},
	})
}

func testAccCheckIbmSccReportViolationDriftDataSourceConfigBasic(instanceID, reportID string) string {
	return fmt.Sprintf(`
		data "ibm_scc_report_violation_drift" "scc_report_violation_drift_instance" {
			instance_id = "%s"
			report_id = "%s"
		}
	`, instanceID, reportID)
}

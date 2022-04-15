// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISLBSDatasource_basic(t *testing.T) {
	name := fmt.Sprintf("tflb-name-%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	var lb string

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{

			{
				Config: testDSCheckIBMISLBSConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_lb", lb),
					resource.TestCheckResourceAttr(
						"data.ibm_is_lb.ds_lb", "name", name),
				),
			},
			{
				Config: testDSCheckIBMISLBSDatasourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_lbs.test_lbs", "load_balancers.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lbs.test_lbs", "load_balancers.0.route_mode"),
				),
			},
		},
	})

}
func TestAccIBMISLBSDatasource_ReservedIp(t *testing.T) {
	name := fmt.Sprintf("tflb-name-%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	var lb string

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{

			{
				Config: testDSCheckIBMISLBSConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_lb", lb),
					resource.TestCheckResourceAttr(
						"data.ibm_is_lb.ds_lb", "name", name),
				),
			},
			{
				Config: testDSCheckIBMISLBSDatasourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_lbs.test_lbs", "load_balancers.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lbs.test_lbs", "load_balancers.0.route_mode"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lbs.test_lbs", "load_balancers.0.private_ip.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lbs.test_lbs", "load_balancers.0.private_ip.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lbs.test_lbs", "load_balancers.0.private_ip.0.address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lbs.test_lbs", "load_balancers.0.private_ip.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lbs.test_lbs", "load_balancers.0.private_ip.0.reserved_ip"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lbs.test_lbs", "load_balancers.0.private_ip.0.resource_type"),
				),
			},
		},
	})

}

func testDSCheckIBMISLBSConfig(vpcname, subnetname, zone, cidr, name string) string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  resource "ibm_is_lb" "testacc_lb" {
		name    = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
	  }
	  data "ibm_is_lb" "ds_lb" {
		name = ibm_is_lb.testacc_lb.name
	  }`, vpcname, subnetname, zone, cidr, name)
}
func testDSCheckIBMISLBSDatasourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
      data "ibm_is_lbs" "test_lbs" {
	  }`)
}

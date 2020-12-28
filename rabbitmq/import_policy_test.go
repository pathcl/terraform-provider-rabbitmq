package rabbitmq

import (
	"testing"

	rabbithole "github.com/pathcl/rabbit-hole/v2"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccPolicy_importBasic(t *testing.T) {
	resourceName := "rabbitmq_policy.test"
	var policy rabbithole.Policy

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyCheckDestroy(&policy),
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyConfig_basic,
				Check: testAccPolicyCheck(
					resourceName, &policy,
				),
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

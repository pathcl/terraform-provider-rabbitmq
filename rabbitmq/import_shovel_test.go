package rabbitmq

import (
	"testing"

	rabbithole "github.com/pathcl/rabbit-hole/v2"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccShovel_importBasic(t *testing.T) {

	resourceName := "rabbitmq_shovel.shovelTest"
	var shovel rabbithole.ShovelInfo

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccShovelCheckDestroy(&shovel),
		Steps: []resource.TestStep{
			{
				Config: testAccShovelConfig_basic,
				Check: testAccShovelCheck(
					resourceName, &shovel,
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

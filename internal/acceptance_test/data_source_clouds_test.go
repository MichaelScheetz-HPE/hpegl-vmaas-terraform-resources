// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package acceptancetest

import (
	"fmt"
	"testing"

	pkgutils "github.com/HewlettPackard/hpegl-vmaas-terraform-resources/pkg/utils"
	"github.com/spf13/viper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCloud(t *testing.T) {
	pkgutils.SkipAcc(t, "vmaas.datasource.cloud")

	resource.ParallelTest(t, resource.TestCase{
		IsUnitTest: false,
		PreCheck:   func() { testAccPreCheck(t) },
		Providers:  testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCloudConfig(),
				Check:  validateDataSourceID("data.hpegl_vmaas_cloud.cloud"),
			},
		},
	})
}

func testAccDataSourceCloudConfig() string {
	return providerStanza + fmt.Sprintf(`
	data "hpegl_vmaas_cloud" cloud {
		name = "%s"
	}
`, viper.GetString("vmaas.datasource.cloud.name"))
}

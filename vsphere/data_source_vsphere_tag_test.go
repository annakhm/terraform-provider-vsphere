package vsphere

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceVSphereTag(t *testing.T) {
	var tp *testing.T
	testAccDataSourceVSphereTagCases := []struct {
		name     string
		testCase resource.TestCase
	}{
		{
			"basic",
			resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(tp)
				},
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: testAccDataSourceVSphereTagConfig(),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(
								"data.vsphere_tag.terraform-test-tag-data",
								"name",
								testAccDataSourceVSphereTagConfigName,
							),
							resource.TestCheckResourceAttr(
								"data.vsphere_tag.terraform-test-tag-data",
								"description",
								testAccDataSourceVSphereTagConfigDescription,
							),
							resource.TestCheckResourceAttrPair(
								"data.vsphere_tag.terraform-test-tag-data", "id",
								"vsphere_tag.terraform-test-tag", "id",
							),
							resource.TestCheckResourceAttrPair(
								"data.vsphere_tag.terraform-test-tag-data", "category_id",
								"vsphere_tag_category.terraform-test-category", "id",
							),
						),
					},
				},
			},
		},
	}

	for _, tc := range testAccDataSourceVSphereTagCases {
		t.Run(tc.name, func(t *testing.T) {
			tp = t
			resource.Test(t, tc.testCase)
		})
	}
}

const testAccDataSourceVSphereTagConfigName = "terraform-test-tag"
const testAccDataSourceVSphereTagConfigDescription = "Managed by Terraform"

func testAccDataSourceVSphereTagConfig() string {
	return fmt.Sprintf(`
variable "tag_category_name" {
  default = "%s"
}

variable "tag_category_description" {
  default = "%s"
}

variable "tag_category_cardinality" {
  default = "%s"
}

variable "tag_category_associable_types" {
  default = [
    "%s",
  ]
}

variable "tag_name" {
  default = "%s"
}

variable "tag_description" {
  default = "%s"
}

resource "vsphere_tag_category" "terraform-test-category" {
  name        = "${var.tag_category_name}"
  description = "${var.tag_category_description}"
  cardinality = "${var.tag_category_cardinality}"

  associable_types = [
    "${var.tag_category_associable_types}",
  ]
}

resource "vsphere_tag" "terraform-test-tag" {
  name        = "${var.tag_name}"
  description = "${var.tag_description}"
  category_id = "${vsphere_tag_category.terraform-test-category.id}"
}

data "vsphere_tag" "terraform-test-tag-data" {
  name        = "${vsphere_tag.terraform-test-tag.name}"
  category_id = "${vsphere_tag.terraform-test-tag.category_id}"
}
`,
		testAccDataSourceVSphereTagCategoryConfigName,
		testAccDataSourceVSphereTagCategoryConfigDescription,
		testAccDataSourceVSphereTagCategoryConfigCardinality,
		testAccDataSourceVSphereTagCategoryConfigAssociableType,
		testAccDataSourceVSphereTagConfigName,
		testAccDataSourceVSphereTagConfigDescription,
	)
}

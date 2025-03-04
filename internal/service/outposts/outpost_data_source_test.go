// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package outposts_test

import (
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/service/outposts"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccOutpostsOutpostDataSource_id(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_outposts_outpost.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, outposts.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOutpostDataSourceConfig_id(),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckResourceAttrRegionalARNIgnoreRegionAndAccount(dataSourceName, "arn", "outposts", regexache.MustCompile(`outpost/op-.+$`).String()),
					resource.TestMatchResourceAttr(dataSourceName, "availability_zone", regexache.MustCompile(`^.+$`)),
					resource.TestMatchResourceAttr(dataSourceName, "availability_zone_id", regexache.MustCompile(`^.+$`)),
					resource.TestCheckResourceAttrSet(dataSourceName, "description"),
					resource.TestMatchResourceAttr(dataSourceName, "lifecycle_status", regexache.MustCompile(`^[ A-Za-z]+$`)),
					resource.TestMatchResourceAttr(dataSourceName, "id", regexache.MustCompile(`^op-.+$`)),
					resource.TestMatchResourceAttr(dataSourceName, "name", regexache.MustCompile(`^.+$`)),
					resource.TestMatchResourceAttr(dataSourceName, "owner_id", regexache.MustCompile(`\d{12}`)),
					acctest.CheckResourceAttrRegionalARNIgnoreRegionAndAccount(dataSourceName, "site_arn", "outposts", regexache.MustCompile(`site/os.+$`).String()),
					resource.TestCheckResourceAttrSet(dataSourceName, "site_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "supported_hardware_type"),
					resource.TestCheckResourceAttrSet(dataSourceName, "tags.%"),
				),
			},
		},
	})
}

func TestAccOutpostsOutpostDataSource_name(t *testing.T) {
	ctx := acctest.Context(t)
	sourceDataSourceName := "data.aws_outposts_outpost.source"
	dataSourceName := "data.aws_outposts_outpost.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, outposts.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOutpostDataSourceConfig_name(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "arn", sourceDataSourceName, "arn"),
					resource.TestCheckResourceAttrPair(dataSourceName, "availability_zone", sourceDataSourceName, "availability_zone"),
					resource.TestCheckResourceAttrPair(dataSourceName, "availability_zone_id", sourceDataSourceName, "availability_zone_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", sourceDataSourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", sourceDataSourceName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", sourceDataSourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "owner_id", sourceDataSourceName, "owner_id"),
				),
			},
		},
	})
}

func TestAccOutpostsOutpostDataSource_arn(t *testing.T) {
	ctx := acctest.Context(t)
	sourceDataSourceName := "data.aws_outposts_outpost.source"
	dataSourceName := "data.aws_outposts_outpost.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, outposts.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOutpostDataSourceConfig_arn(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "arn", sourceDataSourceName, "arn"),
					resource.TestCheckResourceAttrPair(dataSourceName, "availability_zone", sourceDataSourceName, "availability_zone"),
					resource.TestCheckResourceAttrPair(dataSourceName, "availability_zone_id", sourceDataSourceName, "availability_zone_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", sourceDataSourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", sourceDataSourceName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", sourceDataSourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "owner_id", sourceDataSourceName, "owner_id"),
				),
			},
		},
	})
}

func TestAccOutpostsOutpostDataSource_ownerID(t *testing.T) {
	ctx := acctest.Context(t)
	sourceDataSourceName := "data.aws_outposts_outpost.source"
	dataSourceName := "data.aws_outposts_outpost.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, outposts.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOutpostDataSourceConfig_ownerID(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "arn", sourceDataSourceName, "arn"),
					resource.TestCheckResourceAttrPair(dataSourceName, "availability_zone", sourceDataSourceName, "availability_zone"),
					resource.TestCheckResourceAttrPair(dataSourceName, "availability_zone_id", sourceDataSourceName, "availability_zone_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", sourceDataSourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", sourceDataSourceName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", sourceDataSourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "owner_id", sourceDataSourceName, "owner_id"),
				),
			},
		},
	})
}

func testAccOutpostDataSourceConfig_id() string {
	return `
data "aws_outposts_outposts" "test" {}

data "aws_outposts_outpost" "test" {
  id = tolist(data.aws_outposts_outposts.test.ids)[0]
}
`
}

func testAccOutpostDataSourceConfig_name() string {
	return `
data "aws_outposts_outposts" "test" {}

data "aws_outposts_outpost" "source" {
  id = tolist(data.aws_outposts_outposts.test.ids)[0]
}

data "aws_outposts_outpost" "test" {
  name = data.aws_outposts_outpost.source.name
}
`
}

func testAccOutpostDataSourceConfig_arn() string {
	return `
data "aws_outposts_outposts" "test" {}

data "aws_outposts_outpost" "source" {
  arn = tolist(data.aws_outposts_outposts.test.arns)[0]
}

data "aws_outposts_outpost" "test" {
  arn = data.aws_outposts_outpost.source.arn
}
`
}

func testAccOutpostDataSourceConfig_ownerID() string {
	return `
data "aws_outposts_outposts" "test" {}

data "aws_outposts_outpost" "source" {
  id = tolist(data.aws_outposts_outposts.test.ids)[0]
}

data "aws_outposts_outpost" "test" {
  owner_id = data.aws_outposts_outpost.source.owner_id
}
`
}

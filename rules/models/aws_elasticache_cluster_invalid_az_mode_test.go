// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsElastiCacheClusterInvalidAzModeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_elasticache_cluster" "foo" {
	az_mode = "multi-az"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsElastiCacheClusterInvalidAzModeRule(),
					Message: `"multi-az" is an invalid value as az_mode`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_elasticache_cluster" "foo" {
	az_mode = "cross-az"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsElastiCacheClusterInvalidAzModeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
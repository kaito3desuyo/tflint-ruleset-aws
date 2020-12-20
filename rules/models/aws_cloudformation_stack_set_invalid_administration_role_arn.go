// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudformationStackSetInvalidAdministrationRoleArnRule checks the pattern is valid
type AwsCloudformationStackSetInvalidAdministrationRoleArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudformationStackSetInvalidAdministrationRoleArnRule returns new rule with default attributes
func NewAwsCloudformationStackSetInvalidAdministrationRoleArnRule() *AwsCloudformationStackSetInvalidAdministrationRoleArnRule {
	return &AwsCloudformationStackSetInvalidAdministrationRoleArnRule{
		resourceType:  "aws_cloudformation_stack_set",
		attributeName: "administration_role_arn",
		max:           2048,
		min:           20,
	}
}

// Name returns the rule name
func (r *AwsCloudformationStackSetInvalidAdministrationRoleArnRule) Name() string {
	return "aws_cloudformation_stack_set_invalid_administration_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudformationStackSetInvalidAdministrationRoleArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudformationStackSetInvalidAdministrationRoleArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudformationStackSetInvalidAdministrationRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudformationStackSetInvalidAdministrationRoleArnRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"administration_role_arn must be 2048 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"administration_role_arn must be 20 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
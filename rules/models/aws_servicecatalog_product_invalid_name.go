// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsServicecatalogProductInvalidNameRule checks the pattern is valid
type AwsServicecatalogProductInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsServicecatalogProductInvalidNameRule returns new rule with default attributes
func NewAwsServicecatalogProductInvalidNameRule() *AwsServicecatalogProductInvalidNameRule {
	return &AwsServicecatalogProductInvalidNameRule{
		resourceType:  "aws_servicecatalog_product",
		attributeName: "name",
		max:           8191,
	}
}

// Name returns the rule name
func (r *AwsServicecatalogProductInvalidNameRule) Name() string {
	return "aws_servicecatalog_product_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServicecatalogProductInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServicecatalogProductInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServicecatalogProductInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServicecatalogProductInvalidNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"name must be 8191 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
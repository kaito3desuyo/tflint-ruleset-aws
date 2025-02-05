// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53ZoneInvalidCommentRule checks the pattern is valid
type AwsRoute53ZoneInvalidCommentRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsRoute53ZoneInvalidCommentRule returns new rule with default attributes
func NewAwsRoute53ZoneInvalidCommentRule() *AwsRoute53ZoneInvalidCommentRule {
	return &AwsRoute53ZoneInvalidCommentRule{
		resourceType:  "aws_route53_zone",
		attributeName: "comment",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsRoute53ZoneInvalidCommentRule) Name() string {
	return "aws_route53_zone_invalid_comment"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ZoneInvalidCommentRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53ZoneInvalidCommentRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ZoneInvalidCommentRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ZoneInvalidCommentRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"comment must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}

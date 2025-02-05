// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGlueSchemaInvalidSchemaDefinitionRule checks the pattern is valid
type AwsGlueSchemaInvalidSchemaDefinitionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsGlueSchemaInvalidSchemaDefinitionRule returns new rule with default attributes
func NewAwsGlueSchemaInvalidSchemaDefinitionRule() *AwsGlueSchemaInvalidSchemaDefinitionRule {
	return &AwsGlueSchemaInvalidSchemaDefinitionRule{
		resourceType:  "aws_glue_schema",
		attributeName: "schema_definition",
		max:           170000,
		min:           1,
		pattern:       regexp.MustCompile(`^.*\S.*$`),
	}
}

// Name returns the rule name
func (r *AwsGlueSchemaInvalidSchemaDefinitionRule) Name() string {
	return "aws_glue_schema_invalid_schema_definition"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGlueSchemaInvalidSchemaDefinitionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGlueSchemaInvalidSchemaDefinitionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGlueSchemaInvalidSchemaDefinitionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGlueSchemaInvalidSchemaDefinitionRule) Check(runner tflint.Runner) error {
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
					"schema_definition must be 170000 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"schema_definition must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^.*\S.*$`),
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

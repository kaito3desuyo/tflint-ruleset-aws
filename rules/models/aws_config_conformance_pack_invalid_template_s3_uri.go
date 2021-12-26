// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConfigConformancePackInvalidTemplateS3URIRule checks the pattern is valid
type AwsConfigConformancePackInvalidTemplateS3URIRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsConfigConformancePackInvalidTemplateS3URIRule returns new rule with default attributes
func NewAwsConfigConformancePackInvalidTemplateS3URIRule() *AwsConfigConformancePackInvalidTemplateS3URIRule {
	return &AwsConfigConformancePackInvalidTemplateS3URIRule{
		resourceType:  "aws_config_conformance_pack",
		attributeName: "template_s3_uri",
		max:           1024,
		min:           1,
		pattern:       regexp.MustCompile(`^s3://.*$`),
	}
}

// Name returns the rule name
func (r *AwsConfigConformancePackInvalidTemplateS3URIRule) Name() string {
	return "aws_config_conformance_pack_invalid_template_s3_uri"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigConformancePackInvalidTemplateS3URIRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigConformancePackInvalidTemplateS3URIRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigConformancePackInvalidTemplateS3URIRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigConformancePackInvalidTemplateS3URIRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"template_s3_uri must be 1024 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"template_s3_uri must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^s3://.*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
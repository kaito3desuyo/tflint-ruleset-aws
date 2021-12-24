// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAccountAlternateContactInvalidEmailAddressRule checks the pattern is valid
type AwsAccountAlternateContactInvalidEmailAddressRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsAccountAlternateContactInvalidEmailAddressRule returns new rule with default attributes
func NewAwsAccountAlternateContactInvalidEmailAddressRule() *AwsAccountAlternateContactInvalidEmailAddressRule {
	return &AwsAccountAlternateContactInvalidEmailAddressRule{
		resourceType:  "aws_account_alternate_contact",
		attributeName: "email_address",
		max:           64,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w+=,.-]+@[\w.-]+\.[\w]+$`),
	}
}

// Name returns the rule name
func (r *AwsAccountAlternateContactInvalidEmailAddressRule) Name() string {
	return "aws_account_alternate_contact_invalid_email_address"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAccountAlternateContactInvalidEmailAddressRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAccountAlternateContactInvalidEmailAddressRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAccountAlternateContactInvalidEmailAddressRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAccountAlternateContactInvalidEmailAddressRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"email_address must be 64 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"email_address must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\w+=,.-]+@[\w.-]+\.[\w]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCognitoUserPoolDomainInvalidDomainRule checks the pattern is valid
type AwsCognitoUserPoolDomainInvalidDomainRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCognitoUserPoolDomainInvalidDomainRule returns new rule with default attributes
func NewAwsCognitoUserPoolDomainInvalidDomainRule() *AwsCognitoUserPoolDomainInvalidDomainRule {
	return &AwsCognitoUserPoolDomainInvalidDomainRule{
		resourceType:  "aws_cognito_user_pool_domain",
		attributeName: "domain",
		max:           63,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-z0-9](?:[a-z0-9\-]{0,61}[a-z0-9])?$`),
	}
}

// Name returns the rule name
func (r *AwsCognitoUserPoolDomainInvalidDomainRule) Name() string {
	return "aws_cognito_user_pool_domain_invalid_domain"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoUserPoolDomainInvalidDomainRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCognitoUserPoolDomainInvalidDomainRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoUserPoolDomainInvalidDomainRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoUserPoolDomainInvalidDomainRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"domain must be 63 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"domain must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z0-9](?:[a-z0-9\-]{0,61}[a-z0-9])?$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
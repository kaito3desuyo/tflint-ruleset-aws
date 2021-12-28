// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsNetworkfirewallRuleGroupInvalidTypeRule checks the pattern is valid
type AwsNetworkfirewallRuleGroupInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsNetworkfirewallRuleGroupInvalidTypeRule returns new rule with default attributes
func NewAwsNetworkfirewallRuleGroupInvalidTypeRule() *AwsNetworkfirewallRuleGroupInvalidTypeRule {
	return &AwsNetworkfirewallRuleGroupInvalidTypeRule{
		resourceType:  "aws_networkfirewall_rule_group",
		attributeName: "type",
		enum: []string{
			"STATELESS",
			"STATEFUL",
		},
	}
}

// Name returns the rule name
func (r *AwsNetworkfirewallRuleGroupInvalidTypeRule) Name() string {
	return "aws_networkfirewall_rule_group_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsNetworkfirewallRuleGroupInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsNetworkfirewallRuleGroupInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsNetworkfirewallRuleGroupInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsNetworkfirewallRuleGroupInvalidTypeRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
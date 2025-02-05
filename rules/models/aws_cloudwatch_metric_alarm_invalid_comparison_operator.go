// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchMetricAlarmInvalidComparisonOperatorRule checks the pattern is valid
type AwsCloudwatchMetricAlarmInvalidComparisonOperatorRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCloudwatchMetricAlarmInvalidComparisonOperatorRule returns new rule with default attributes
func NewAwsCloudwatchMetricAlarmInvalidComparisonOperatorRule() *AwsCloudwatchMetricAlarmInvalidComparisonOperatorRule {
	return &AwsCloudwatchMetricAlarmInvalidComparisonOperatorRule{
		resourceType:  "aws_cloudwatch_metric_alarm",
		attributeName: "comparison_operator",
		enum: []string{
			"GreaterThanOrEqualToThreshold",
			"GreaterThanThreshold",
			"LessThanThreshold",
			"LessThanOrEqualToThreshold",
			"LessThanLowerOrGreaterThanUpperThreshold",
			"LessThanLowerThreshold",
			"GreaterThanUpperThreshold",
		},
	}
}

// Name returns the rule name
func (r *AwsCloudwatchMetricAlarmInvalidComparisonOperatorRule) Name() string {
	return "aws_cloudwatch_metric_alarm_invalid_comparison_operator"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchMetricAlarmInvalidComparisonOperatorRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchMetricAlarmInvalidComparisonOperatorRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchMetricAlarmInvalidComparisonOperatorRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchMetricAlarmInvalidComparisonOperatorRule) Check(runner tflint.Runner) error {
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
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as comparison_operator`, truncateLongMessage(val)),
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

// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule checks the pattern is valid
type AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsCloudwatchMetricAlarmInvalidExtendedStatisticRule returns new rule with default attributes
func NewAwsCloudwatchMetricAlarmInvalidExtendedStatisticRule() *AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule {
	return &AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule{
		resourceType:  "aws_cloudwatch_metric_alarm",
		attributeName: "extended_statistic",
		pattern:       regexp.MustCompile(`^p(\d{1,2}(\.\d{0,2})?|100)$`),
	}
}

// Name returns the rule name
func (r *AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule) Name() string {
	return "aws_cloudwatch_metric_alarm_invalid_extended_statistic"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule) Check(runner tflint.Runner) error {
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^p(\d{1,2}(\.\d{0,2})?|100)$`),
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

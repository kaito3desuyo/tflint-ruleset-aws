// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGlueMlTransformInvalidWorkerTypeRule checks the pattern is valid
type AwsGlueMlTransformInvalidWorkerTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsGlueMlTransformInvalidWorkerTypeRule returns new rule with default attributes
func NewAwsGlueMlTransformInvalidWorkerTypeRule() *AwsGlueMlTransformInvalidWorkerTypeRule {
	return &AwsGlueMlTransformInvalidWorkerTypeRule{
		resourceType:  "aws_glue_ml_transform",
		attributeName: "worker_type",
		enum: []string{
			"Standard",
			"G.1X",
			"G.2X",
		},
	}
}

// Name returns the rule name
func (r *AwsGlueMlTransformInvalidWorkerTypeRule) Name() string {
	return "aws_glue_ml_transform_invalid_worker_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGlueMlTransformInvalidWorkerTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGlueMlTransformInvalidWorkerTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGlueMlTransformInvalidWorkerTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGlueMlTransformInvalidWorkerTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as worker_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}

// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule checks the pattern is valid
type AwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule returns new rule with default attributes
func NewAwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule() *AwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule {
	return &AwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule{
		resourceType:  "aws_ec2_transit_gateway",
		attributeName: "default_route_table_association",
		enum: []string{
			"enable",
			"disable",
		},
	}
}

// Name returns the rule name
func (r *AwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule) Name() string {
	return "aws_ec2_transit_gateway_invalid_default_route_table_association"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as default_route_table_association`, truncateLongMessage(val)),
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

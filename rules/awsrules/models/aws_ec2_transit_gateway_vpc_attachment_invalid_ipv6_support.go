// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsEc2TransitGatewayVpcAttachmentInvalidIpv6SupportRule checks the pattern is valid
type AwsEc2TransitGatewayVpcAttachmentInvalidIpv6SupportRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEc2TransitGatewayVpcAttachmentInvalidIpv6SupportRule returns new rule with default attributes
func NewAwsEc2TransitGatewayVpcAttachmentInvalidIpv6SupportRule() *AwsEc2TransitGatewayVpcAttachmentInvalidIpv6SupportRule {
	return &AwsEc2TransitGatewayVpcAttachmentInvalidIpv6SupportRule{
		resourceType:  "aws_ec2_transit_gateway_vpc_attachment",
		attributeName: "ipv6_support",
		enum: []string{
			"enable",
			"disable",
		},
	}
}

// Name returns the rule name
func (r *AwsEc2TransitGatewayVpcAttachmentInvalidIpv6SupportRule) Name() string {
	return "aws_ec2_transit_gateway_vpc_attachment_invalid_ipv6_support"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEc2TransitGatewayVpcAttachmentInvalidIpv6SupportRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsEc2TransitGatewayVpcAttachmentInvalidIpv6SupportRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsEc2TransitGatewayVpcAttachmentInvalidIpv6SupportRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEc2TransitGatewayVpcAttachmentInvalidIpv6SupportRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`ipv6_support is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

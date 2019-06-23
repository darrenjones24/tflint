// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsRoute53ResolverRuleAssociationInvalidVpcIDRule checks the pattern is valid
type AwsRoute53ResolverRuleAssociationInvalidVpcIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsRoute53ResolverRuleAssociationInvalidVpcIDRule returns new rule with default attributes
func NewAwsRoute53ResolverRuleAssociationInvalidVpcIDRule() *AwsRoute53ResolverRuleAssociationInvalidVpcIDRule {
	return &AwsRoute53ResolverRuleAssociationInvalidVpcIDRule{
		resourceType:  "aws_route53_resolver_rule_association",
		attributeName: "vpc_id",
		max:           64,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsRoute53ResolverRuleAssociationInvalidVpcIDRule) Name() string {
	return "aws_route53_resolver_rule_association_invalid_vpc_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ResolverRuleAssociationInvalidVpcIDRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsRoute53ResolverRuleAssociationInvalidVpcIDRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ResolverRuleAssociationInvalidVpcIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ResolverRuleAssociationInvalidVpcIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"vpc_id must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"vpc_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

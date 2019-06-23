// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsRoute53ZoneInvalidNameRule checks the pattern is valid
type AwsRoute53ZoneInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsRoute53ZoneInvalidNameRule returns new rule with default attributes
func NewAwsRoute53ZoneInvalidNameRule() *AwsRoute53ZoneInvalidNameRule {
	return &AwsRoute53ZoneInvalidNameRule{
		resourceType:  "aws_route53_zone",
		attributeName: "name",
		max:           1024,
	}
}

// Name returns the rule name
func (r *AwsRoute53ZoneInvalidNameRule) Name() string {
	return "aws_route53_zone_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ZoneInvalidNameRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsRoute53ZoneInvalidNameRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ZoneInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ZoneInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

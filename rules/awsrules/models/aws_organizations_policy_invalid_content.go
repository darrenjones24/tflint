// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsOrganizationsPolicyInvalidContentRule checks the pattern is valid
type AwsOrganizationsPolicyInvalidContentRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsOrganizationsPolicyInvalidContentRule returns new rule with default attributes
func NewAwsOrganizationsPolicyInvalidContentRule() *AwsOrganizationsPolicyInvalidContentRule {
	return &AwsOrganizationsPolicyInvalidContentRule{
		resourceType:  "aws_organizations_policy",
		attributeName: "content",
		max:           1000000,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsOrganizationsPolicyInvalidContentRule) Name() string {
	return "aws_organizations_policy_invalid_content"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOrganizationsPolicyInvalidContentRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsOrganizationsPolicyInvalidContentRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsOrganizationsPolicyInvalidContentRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOrganizationsPolicyInvalidContentRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"content must be 1000000 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"content must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

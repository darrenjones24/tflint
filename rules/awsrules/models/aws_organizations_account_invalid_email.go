// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsOrganizationsAccountInvalidEmailRule checks the pattern is valid
type AwsOrganizationsAccountInvalidEmailRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsOrganizationsAccountInvalidEmailRule returns new rule with default attributes
func NewAwsOrganizationsAccountInvalidEmailRule() *AwsOrganizationsAccountInvalidEmailRule {
	return &AwsOrganizationsAccountInvalidEmailRule{
		resourceType:  "aws_organizations_account",
		attributeName: "email",
		max:           64,
		min:           6,
		pattern:       regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`),
	}
}

// Name returns the rule name
func (r *AwsOrganizationsAccountInvalidEmailRule) Name() string {
	return "aws_organizations_account_invalid_email"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOrganizationsAccountInvalidEmailRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsOrganizationsAccountInvalidEmailRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsOrganizationsAccountInvalidEmailRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOrganizationsAccountInvalidEmailRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"email must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"email must be 6 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`email does not match valid pattern ^[^\s@]+@[^\s@]+\.[^\s@]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

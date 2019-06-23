// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsLaunchTemplateInvalidDescriptionRule checks the pattern is valid
type AwsLaunchTemplateInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsLaunchTemplateInvalidDescriptionRule returns new rule with default attributes
func NewAwsLaunchTemplateInvalidDescriptionRule() *AwsLaunchTemplateInvalidDescriptionRule {
	return &AwsLaunchTemplateInvalidDescriptionRule{
		resourceType:  "aws_launch_template",
		attributeName: "description",
		max:           255,
	}
}

// Name returns the rule name
func (r *AwsLaunchTemplateInvalidDescriptionRule) Name() string {
	return "aws_launch_template_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLaunchTemplateInvalidDescriptionRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsLaunchTemplateInvalidDescriptionRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsLaunchTemplateInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLaunchTemplateInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

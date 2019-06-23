// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsXraySamplingRuleInvalidURLPathRule checks the pattern is valid
type AwsXraySamplingRuleInvalidURLPathRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsXraySamplingRuleInvalidURLPathRule returns new rule with default attributes
func NewAwsXraySamplingRuleInvalidURLPathRule() *AwsXraySamplingRuleInvalidURLPathRule {
	return &AwsXraySamplingRuleInvalidURLPathRule{
		resourceType:  "aws_xray_sampling_rule",
		attributeName: "url_path",
		max:           128,
	}
}

// Name returns the rule name
func (r *AwsXraySamplingRuleInvalidURLPathRule) Name() string {
	return "aws_xray_sampling_rule_invalid_url_path"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsXraySamplingRuleInvalidURLPathRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsXraySamplingRuleInvalidURLPathRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsXraySamplingRuleInvalidURLPathRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsXraySamplingRuleInvalidURLPathRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"url_path must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

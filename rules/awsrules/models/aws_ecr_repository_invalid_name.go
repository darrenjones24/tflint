// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsEcrRepositoryInvalidNameRule checks the pattern is valid
type AwsEcrRepositoryInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsEcrRepositoryInvalidNameRule returns new rule with default attributes
func NewAwsEcrRepositoryInvalidNameRule() *AwsEcrRepositoryInvalidNameRule {
	return &AwsEcrRepositoryInvalidNameRule{
		resourceType:  "aws_ecr_repository",
		attributeName: "name",
		max:           256,
		min:           2,
		pattern:       regexp.MustCompile(`^(?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*$`),
	}
}

// Name returns the rule name
func (r *AwsEcrRepositoryInvalidNameRule) Name() string {
	return "aws_ecr_repository_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEcrRepositoryInvalidNameRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsEcrRepositoryInvalidNameRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsEcrRepositoryInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEcrRepositoryInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 2 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`name does not match valid pattern ^(?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

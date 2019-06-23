// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsSecretsmanagerSecretInvalidNameRule checks the pattern is valid
type AwsSecretsmanagerSecretInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSecretsmanagerSecretInvalidNameRule returns new rule with default attributes
func NewAwsSecretsmanagerSecretInvalidNameRule() *AwsSecretsmanagerSecretInvalidNameRule {
	return &AwsSecretsmanagerSecretInvalidNameRule{
		resourceType:  "aws_secretsmanager_secret",
		attributeName: "name",
		max:           512,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSecretsmanagerSecretInvalidNameRule) Name() string {
	return "aws_secretsmanager_secret_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSecretsmanagerSecretInvalidNameRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsSecretsmanagerSecretInvalidNameRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsSecretsmanagerSecretInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSecretsmanagerSecretInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 512 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

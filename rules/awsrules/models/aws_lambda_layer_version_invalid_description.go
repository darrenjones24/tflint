// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsLambdaLayerVersionInvalidDescriptionRule checks the pattern is valid
type AwsLambdaLayerVersionInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsLambdaLayerVersionInvalidDescriptionRule returns new rule with default attributes
func NewAwsLambdaLayerVersionInvalidDescriptionRule() *AwsLambdaLayerVersionInvalidDescriptionRule {
	return &AwsLambdaLayerVersionInvalidDescriptionRule{
		resourceType:  "aws_lambda_layer_version",
		attributeName: "description",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsLambdaLayerVersionInvalidDescriptionRule) Name() string {
	return "aws_lambda_layer_version_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaLayerVersionInvalidDescriptionRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsLambdaLayerVersionInvalidDescriptionRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaLayerVersionInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaLayerVersionInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsStoragegatewayNfsFileShareInvalidSquashRule checks the pattern is valid
type AwsStoragegatewayNfsFileShareInvalidSquashRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewayNfsFileShareInvalidSquashRule returns new rule with default attributes
func NewAwsStoragegatewayNfsFileShareInvalidSquashRule() *AwsStoragegatewayNfsFileShareInvalidSquashRule {
	return &AwsStoragegatewayNfsFileShareInvalidSquashRule{
		resourceType:  "aws_storagegateway_nfs_file_share",
		attributeName: "squash",
		max:           15,
		min:           5,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayNfsFileShareInvalidSquashRule) Name() string {
	return "aws_storagegateway_nfs_file_share_invalid_squash"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayNfsFileShareInvalidSquashRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsStoragegatewayNfsFileShareInvalidSquashRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayNfsFileShareInvalidSquashRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayNfsFileShareInvalidSquashRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"squash must be 15 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"squash must be 5 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

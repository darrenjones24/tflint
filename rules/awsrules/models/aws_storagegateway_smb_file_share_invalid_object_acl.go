// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsStoragegatewaySmbFileShareInvalidObjectACLRule checks the pattern is valid
type AwsStoragegatewaySmbFileShareInvalidObjectACLRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsStoragegatewaySmbFileShareInvalidObjectACLRule returns new rule with default attributes
func NewAwsStoragegatewaySmbFileShareInvalidObjectACLRule() *AwsStoragegatewaySmbFileShareInvalidObjectACLRule {
	return &AwsStoragegatewaySmbFileShareInvalidObjectACLRule{
		resourceType:  "aws_storagegateway_smb_file_share",
		attributeName: "object_acl",
		enum: []string{
			"private",
			"public-read",
			"public-read-write",
			"authenticated-read",
			"bucket-owner-read",
			"bucket-owner-full-control",
			"aws-exec-read",
		},
	}
}

// Name returns the rule name
func (r *AwsStoragegatewaySmbFileShareInvalidObjectACLRule) Name() string {
	return "aws_storagegateway_smb_file_share_invalid_object_acl"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewaySmbFileShareInvalidObjectACLRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsStoragegatewaySmbFileShareInvalidObjectACLRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewaySmbFileShareInvalidObjectACLRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewaySmbFileShareInvalidObjectACLRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`object_acl is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

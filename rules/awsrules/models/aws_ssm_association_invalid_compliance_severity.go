// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsSsmAssociationInvalidComplianceSeverityRule checks the pattern is valid
type AwsSsmAssociationInvalidComplianceSeverityRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSsmAssociationInvalidComplianceSeverityRule returns new rule with default attributes
func NewAwsSsmAssociationInvalidComplianceSeverityRule() *AwsSsmAssociationInvalidComplianceSeverityRule {
	return &AwsSsmAssociationInvalidComplianceSeverityRule{
		resourceType:  "aws_ssm_association",
		attributeName: "compliance_severity",
		enum: []string{
			"CRITICAL",
			"HIGH",
			"MEDIUM",
			"LOW",
			"UNSPECIFIED",
		},
	}
}

// Name returns the rule name
func (r *AwsSsmAssociationInvalidComplianceSeverityRule) Name() string {
	return "aws_ssm_association_invalid_compliance_severity"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmAssociationInvalidComplianceSeverityRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsSsmAssociationInvalidComplianceSeverityRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmAssociationInvalidComplianceSeverityRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmAssociationInvalidComplianceSeverityRule) Check(runner *tflint.Runner) error {
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
					`compliance_severity is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsOrganizationsPolicyAttachmentInvalidTargetIDRule checks the pattern is valid
type AwsOrganizationsPolicyAttachmentInvalidTargetIDRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsOrganizationsPolicyAttachmentInvalidTargetIDRule returns new rule with default attributes
func NewAwsOrganizationsPolicyAttachmentInvalidTargetIDRule() *AwsOrganizationsPolicyAttachmentInvalidTargetIDRule {
	return &AwsOrganizationsPolicyAttachmentInvalidTargetIDRule{
		resourceType:  "aws_organizations_policy_attachment",
		attributeName: "target_id",
		pattern:       regexp.MustCompile(`^(r-[0-9a-z]{4,32})|(\d{12})|(ou-[0-9a-z]{4,32}-[a-z0-9]{8,32})$`),
	}
}

// Name returns the rule name
func (r *AwsOrganizationsPolicyAttachmentInvalidTargetIDRule) Name() string {
	return "aws_organizations_policy_attachment_invalid_target_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOrganizationsPolicyAttachmentInvalidTargetIDRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsOrganizationsPolicyAttachmentInvalidTargetIDRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsOrganizationsPolicyAttachmentInvalidTargetIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOrganizationsPolicyAttachmentInvalidTargetIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`target_id does not match valid pattern ^(r-[0-9a-z]{4,32})|(\d{12})|(ou-[0-9a-z]{4,32}-[a-z0-9]{8,32})$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

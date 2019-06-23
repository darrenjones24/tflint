// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsRoute53HealthCheckInvalidInsufficientDataHealthStatusRule checks the pattern is valid
type AwsRoute53HealthCheckInvalidInsufficientDataHealthStatusRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsRoute53HealthCheckInvalidInsufficientDataHealthStatusRule returns new rule with default attributes
func NewAwsRoute53HealthCheckInvalidInsufficientDataHealthStatusRule() *AwsRoute53HealthCheckInvalidInsufficientDataHealthStatusRule {
	return &AwsRoute53HealthCheckInvalidInsufficientDataHealthStatusRule{
		resourceType:  "aws_route53_health_check",
		attributeName: "insufficient_data_health_status",
		enum: []string{
			"Healthy",
			"Unhealthy",
			"LastKnownStatus",
		},
	}
}

// Name returns the rule name
func (r *AwsRoute53HealthCheckInvalidInsufficientDataHealthStatusRule) Name() string {
	return "aws_route53_health_check_invalid_insufficient_data_health_status"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53HealthCheckInvalidInsufficientDataHealthStatusRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsRoute53HealthCheckInvalidInsufficientDataHealthStatusRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53HealthCheckInvalidInsufficientDataHealthStatusRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53HealthCheckInvalidInsufficientDataHealthStatusRule) Check(runner *tflint.Runner) error {
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
					`insufficient_data_health_status is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

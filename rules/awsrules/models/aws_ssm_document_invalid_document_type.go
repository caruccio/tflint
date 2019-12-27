// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsSsmDocumentInvalidDocumentTypeRule checks the pattern is valid
type AwsSsmDocumentInvalidDocumentTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSsmDocumentInvalidDocumentTypeRule returns new rule with default attributes
func NewAwsSsmDocumentInvalidDocumentTypeRule() *AwsSsmDocumentInvalidDocumentTypeRule {
	return &AwsSsmDocumentInvalidDocumentTypeRule{
		resourceType:  "aws_ssm_document",
		attributeName: "document_type",
		enum: []string{
			"Command",
			"Policy",
			"Automation",
			"Session",
			"Package",
			"ApplicationConfiguration",
			"ApplicationConfigurationSchema",
			"DeploymentStrategy",
			"ChangeCalendar",
		},
	}
}

// Name returns the rule name
func (r *AwsSsmDocumentInvalidDocumentTypeRule) Name() string {
	return "aws_ssm_document_invalid_document_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmDocumentInvalidDocumentTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmDocumentInvalidDocumentTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmDocumentInvalidDocumentTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmDocumentInvalidDocumentTypeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

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
					`document_type is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

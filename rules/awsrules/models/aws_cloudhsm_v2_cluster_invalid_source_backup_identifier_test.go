// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform/configs"
	"github.com/hashicorp/terraform/configs/configload"
	"github.com/hashicorp/terraform/terraform"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

func Test_AwsCloudhsmV2ClusterInvalidSourceBackupIdentifierRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected issue.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudhsm_v2_cluster" "foo" {
	source_backup_identifier = "rtq2dwi2gq6"
}`,
			Expected: []*issue.Issue{
				{
					Detector: "aws_cloudhsm_v2_cluster_invalid_source_backup_identifier",
					Type:     "ERROR",
					Message:  `source_backup_identifier does not match valid pattern ^backup-[2-7a-zA-Z]{11,16}$`,
					Line:     3,
					File:     "resource.tf",
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudhsm_v2_cluster" "foo" {
	source_backup_identifier = "backup-rtq2dwi2gq6"
}`,
			Expected: []*issue.Issue{},
		},
	}

	dir, err := ioutil.TempDir("", "AwsCloudhsmV2ClusterInvalidSourceBackupIdentifierRule")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(currentDir)

	err = os.Chdir(dir)
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range cases {
		loader, err := configload.NewLoader(&configload.Config{})
		if err != nil {
			t.Fatal(err)
		}

		err = ioutil.WriteFile(dir+"/resource.tf", []byte(tc.Content), os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}

		mod, diags := loader.Parser().LoadConfigDir(".")
		if diags.HasErrors() {
			t.Fatal(diags)
		}
		cfg, tfdiags := configs.BuildConfig(mod, configs.DisabledModuleWalker)
		if tfdiags.HasErrors() {
			t.Fatal(tfdiags)
		}

		runner := tflint.NewRunner(tflint.EmptyConfig(), map[string]tflint.Annotations{}, cfg, map[string]*terraform.InputValue{})
		rule := NewAwsCloudhsmV2ClusterInvalidSourceBackupIdentifierRule()

		if err = rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		if !cmp.Equal(tc.Expected, runner.Issues) {
			t.Fatalf("Expected issues are not matched:\n %s\n", cmp.Diff(tc.Expected, runner.Issues))
		}
	}
}

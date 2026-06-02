package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestExecuteReturnsRootCommandErrors(t *testing.T) {
	rootCmd.SetArgs([]string{"unknown-command"})
	rootCmd.SetOut(new(bytes.Buffer))
	rootCmd.SetErr(new(bytes.Buffer))

	err := Execute()
	if err == nil {
		t.Fatal("expected Execute to return an error")
	}
}

func TestOptimizeRejectsInvalidPopulationBeforeLoadingInstance(t *testing.T) {
	_, _, err := executeRootForTest(
		"optimize", "ga",
		"--population", "0",
		"--instance", "missing.graph",
		"--results-dir", t.TempDir(),
	)
	if err == nil {
		t.Fatal("expected invalid population to fail")
	}
	if !strings.Contains(err.Error(), "invalid --population") {
		t.Fatalf("expected population validation error, got: %v", err)
	}
}

func TestServeWritesStatusToStderr(t *testing.T) {
	stdout, stderr, err := executeRootForTest("serve", "--addr", "bad addr")
	if err == nil {
		t.Fatal("expected invalid listen addr to fail")
	}
	if stdout != "" {
		t.Fatalf("expected stdout to remain empty, got: %q", stdout)
	}
	if !strings.Contains(stderr, "web server running") {
		t.Fatalf("expected server status on stderr, got: %q", stderr)
	}
}

func executeRootForTest(args ...string) (string, string, error) {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	rootCmd.SetArgs(args)
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err := rootCmd.Execute()

	rootCmd.SetArgs(nil)
	rootCmd.SetOut(nil)
	rootCmd.SetErr(nil)

	return stdout.String(), stderr.String(), err
}

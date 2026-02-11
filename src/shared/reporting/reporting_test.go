package reporting

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidateRunID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		runID string
		ok    bool
	}{
		{name: "simple", runID: "run_01-abc", ok: true},
		{name: "empty", runID: "", ok: false},
		{name: "with slash", runID: "../../x", ok: false},
		{name: "with backslash", runID: "..\\..\\x", ok: false},
		{name: "with dot", runID: "run.v1", ok: false},
		{name: "with space", runID: "run id", ok: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := ValidateRunID(tt.runID)
			if tt.ok && err != nil {
				t.Fatalf("expected valid run id, got error: %v", err)
			}
			if !tt.ok && err == nil {
				t.Fatal("expected error for invalid run id")
			}
		})
	}
}

func TestCheckExisting(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		policy       string
		create       []string
		wantSkip     bool
		wantError    bool
		errorContain string
	}{
		{
			name:     "skip sem artefatos",
			policy:   IfExistsSkip,
			create:   nil,
			wantSkip: false,
		},
		{
			name:     "skip com artefatos completos",
			policy:   IfExistsSkip,
			create:   []string{"summary", "evolution", "timing"},
			wantSkip: true,
		},
		{
			name:     "skip com artefatos incompletos",
			policy:   IfExistsSkip,
			create:   []string{"summary"},
			wantSkip: false,
		},
		{
			name:      "error com artefatos completos",
			policy:    IfExistsError,
			create:    []string{"summary", "evolution", "timing"},
			wantError: true,
		},
		{
			name:         "error com artefatos incompletos",
			policy:       IfExistsError,
			create:       []string{"summary"},
			wantError:    true,
			errorContain: "incomplete artifacts",
		},
		{
			name:     "overwrite com artefatos completos",
			policy:   IfExistsOverwrite,
			create:   []string{"summary", "evolution", "timing"},
			wantSkip: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			root := t.TempDir()
			paths := PathsForRun(root, "run_01")

			for _, artifact := range tt.create {
				switch artifact {
				case "summary":
					touchFile(t, paths.Summary)
				case "evolution":
					touchFile(t, paths.Evolution)
				case "timing":
					touchFile(t, paths.Timing)
				default:
					t.Fatalf("artefato desconhecido no teste: %s", artifact)
				}
			}

			skip, err := CheckExisting(paths, tt.policy)
			if tt.wantError {
				if err == nil {
					t.Fatal("era esperado erro, mas recebeu nil")
				}
				if tt.errorContain != "" && !strings.Contains(err.Error(), tt.errorContain) {
					t.Fatalf("erro inesperado: %v", err)
				}
				return
			}

			if err != nil {
				t.Fatalf("erro inesperado: %v", err)
			}

			if skip != tt.wantSkip {
				t.Fatalf("skip = %v, esperado %v", skip, tt.wantSkip)
			}
		})
	}
}

func touchFile(t *testing.T, path string) {
	t.Helper()

	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		t.Fatalf("erro ao criar diretório do arquivo %q: %v", path, err)
	}

	if err := os.WriteFile(path, []byte("{}\n"), 0o644); err != nil {
		t.Fatalf("erro ao criar arquivo %q: %v", path, err)
	}
}

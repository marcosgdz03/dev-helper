package core

import (
	"testing"
)

func TestNewExecuterDefaults(t *testing.T) {
	exec := NewExecutor("")
	if exec == nil {
		t.Fatal("NewExecutor(\"\") returned nil")
	}
}

func TestCheckCommandAvailable(t *testing.T) {
	exec := NewExecutor("")
	// On Windows, "powershell.exe" is guaranteed to exist
	if !exec.CheckCommand("powershell.exe") {
		t.Log("Warning: powershell.exe not found (this test is meaningless on non-Windows)")
	}
	// This one should always be false
	if exec.CheckCommand("nonexistent-binary-xyz-123") {
		t.Error("CheckCommand should return false for nonexistent binary")
	}
}

func TestArgStr(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{[]string{"a", "b", "c"}, "a b c"},
		{[]string{}, ""},
		{nil, ""},
		{[]string{"single"}, "single"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(st *testing.T) {
			got := argStr(tt.args)
			if got != tt.want {
				st.Errorf("argStr(%v) = %q, want %q", tt.args, got, tt.want)
			}
		})
	}
}

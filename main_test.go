package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

// Helper function to execute the command and capture its output.
func runCmd(args ...string) (string, error) {
	cmd := exec.Command("go", append([]string{"run", "main.go"}, args...)...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

func TestIPCombo(t *testing.T) {
	tests := []struct {
		ip       string
		expected []string
	}{
		{
			ip: "8.8.8.8",
			expected: []string{
				"8.8.8.8\n",
				"0x8.0x8.0x8.0x8\n",
				"010.010.010.010\n",
				"8.0x8.8.8\n",
				"8.8.0x8.8\n",
				"8.8.8.0x8\n",
				"8.010.8.8\n",
				"8.8.010.8\n",
				"8.8.8.010\n",
				"0x8.8.8.8\n",
				"0x8.0x8.8.8\n",
				"0x8.8.0x8.8\n",
				"0x8.8.8.0x8\n",
				"010.010.8.8\n",
				"8.0x8.8.0x8\n",
				"8.0x8.0x8.8\n",
				"134744072\n",
				"0x8080808\n",
				"01002004010\n",
				"8.8.2056\n",
				"8.0x8.2056\n",
				"8.010.2056\n",
				"8.526344\n",
				"0x8.526344\n",
				"010.526344\n",
				"010.010.2056\n",
				"010.010.0x808\n",
				"010.010.04010\n",
			},
		},
	}

	for _, test := range tests {
		output, err := runCmd(test.ip)
		if err != nil {
			t.Errorf("Failed to run command for IP %s: %v", test.ip, err)
		}
		for _, exp := range test.expected {
			if !strings.Contains(output, exp) {
				t.Errorf("For IP %s, expected output to contain %q, got %q", test.ip, exp, output)
			}
		}
	}
}

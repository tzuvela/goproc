package system

import (
	"testing"
)

func TestPrintCPUInfo(t *testing.T) {
	info, err := GetCPUInfo()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(info) == 0 {
		t.Error("Expected at least one CPU entry")
	}
}

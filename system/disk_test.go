package system

import (
	"testing"
)

func TestGetDiskInfo(t *testing.T) {
	info, err := GetDiskInfo()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if info.Total == 0 {
		t.Error("Expected non-zero total disk size")
	}
}

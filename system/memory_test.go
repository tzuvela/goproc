package system

import (
	"testing"
)

func TestGetMemoryInfo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "valid memory info"},
		//TODO add more mock test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats, err := GetMemoryInfo()
			if err != nil {
				t.Fatalf("GetMemoryInfo returned error: %v", err)
			}
			if stats.Total == 0 {
				t.Error("expected Total > 0")
			}
			if stats.UsedPercent < 0 || stats.UsedPercent > 100 {
				t.Errorf("UsedPercent out of range: %v", stats.UsedPercent)
			}
		})

	}
}

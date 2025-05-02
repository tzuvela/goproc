package system

import "testing"

func TestGetMemoryInfo(t *testing.T) {
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
}

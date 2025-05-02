package system

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

// GetMemoryInfo retrieves system‑wide memory statistics.
// It returns a *mem.VirtualMemoryStat and any error encountered.
func GetMemoryInfo() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory() // :contentReference[oaicite:0]{index=0}
}

// PrintMemoryInfo formats and prints memory statistics to stdout.
func PrintMemoryInfo() {
	stats, err := GetMemoryInfo()
	if err != nil {
		fmt.Println("Error getting memory info:", err)
		return
	}
	fmt.Println("Memory Info:")
	fmt.Printf("  Total: %.2f GB\n", float64(stats.Total)/1e9)
	fmt.Printf("  Used:  %.2f GB (%.2f%%)\n\n",
		float64(stats.Used)/1e9, stats.UsedPercent)
}

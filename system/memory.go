package system

import (
	"fmt"
	"runtime"
)

func PrintMemoryInfo() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println("Memory Info:")
	fmt.Printf("  Allocated: %.2f MB\n", float64(m.Alloc)/1024.0/1024.0)
	fmt.Printf("  Total Allocated: %.2f MB\n", float64(m.TotalAlloc)/1024.0/1024.0)
	fmt.Printf("  System Memory: %.2f MB\n", float64(m.Sys)/1024.0/1024.0)
	fmt.Println()
}

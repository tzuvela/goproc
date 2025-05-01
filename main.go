// Entry point for goproc CLI
package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error retrieving hostname:", err)
		return
	}

	fmt.Println("goproc - Simple System Info")
	fmt.Println("===========================")
	fmt.Printf("Hostname: %s\n", hostname)
	fmt.Printf("CPU Cores: %d\n", runtime.NumCPU())
	fmt.Println("Memory Info:")
	printMemoryStats()
}
func printMemoryStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory Allocated: %.2f MB\n", float64(m.Alloc)/1024.0/1024.0)
	fmt.Printf("Total Allocated: %.2f MB\n", float64(m.TotalAlloc)/1024.0/1024.0)
	fmt.Printf("System Memory Obtained: %.2f MB\n", float64(m.Sys)/1024.0/1024.0)
}

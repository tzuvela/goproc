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
}

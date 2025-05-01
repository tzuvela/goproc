package main

import (
	"fmt"

	"github.com/tzuvela/goproc/system"
)

func main() {
	fmt.Println("goproc - Simple System Info")
	fmt.Println("===========================")

	system.PrintHostname()
	system.PrintCPU()
	system.PrintMemory()
	system.PrintDiskUsage()
}

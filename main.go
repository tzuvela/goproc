package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tzuvela/goproc/system"
)

func main() {
	// Define flags
	showCPU := flag.Bool("cpu", false, "Show CPU info")
	showMem := flag.Bool("mem", false, "Show memory info")
	showDisk := flag.Bool("disk", false, "Show disk usage")
	showHost := flag.Bool("host", false, "Show OS and uptime info")
	showAll := flag.Bool("all", false, "Show all system info (default)")

	// Parse the flags
	flag.Parse()

	// Show app header
	fmt.Println("goproc - Simple System Info")
	fmt.Println("===========================")

	// Always show hostname
	system.PrintHostname()

	// If no flags are provided, or --all is passed, show everything
	if len(os.Args) == 1 || *showAll {
		system.PrintHostInfo()
		system.PrintCPUInfo()
		system.PrintMemoryInfo()
		system.PrintDiskUsage()
		return
	}

	// Show individual components based on flags
	if *showHost {
		system.PrintHostInfo()
	}
	if *showCPU {
		system.PrintCPUInfo()
	}
	if *showMem {
		system.PrintMemoryInfo()
	}
	if *showDisk {
		system.PrintDiskUsage()
	}
}

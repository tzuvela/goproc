package system

import (
	"fmt"
	"runtime"
)

func PrintCPUInfo() {
	fmt.Printf("CPU Cores: %d\n", runtime.NumCPU())
	fmt.Println()
}

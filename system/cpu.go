package system

import (
	"fmt"
	"runtime"
)

func PrintCPU() {
	fmt.Printf("CPU Cores: %d\n", runtime.NumCPU())
	fmt.Println()
}

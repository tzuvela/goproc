package system

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
)

func GetCPUInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()

}
func PrintCPUInfo() {
	infos, err := GetCPUInfo()
	if err != nil {
		fmt.Println("Error getting CPU info:", err)
		return
	}
	fmt.Printf("\nCPU Cores: %d\n", len(infos))
}

package system

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
)

func GetCPUInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()

}
func PrintCPUInfo() {
	//Physical core count

	cores, err := cpu.Counts(true)
	if err != nil {
		fmt.Println("Error getting CPU incoresfo:", err)
		return
	}
	fmt.Printf("\nCPU Cores: %d\n", cores)
}

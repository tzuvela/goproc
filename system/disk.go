package system

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
)

func GetDiskInfo() (*disk.UsageStat, error) {
	return disk.Usage("/")
}

func PrintDiskUsage() {
	usage, err := GetDiskInfo()
	if err != nil {
		fmt.Println("Error retrieving disk usage:", err.Error())
		return
	}

	println("Disk Usage (/):")
	fmt.Printf("  Total: %.2f GB\n", float64(usage.Total)/1e9)
	fmt.Printf("  Used:  %.2f GB\n", float64(usage.Used)/1e9)
	fmt.Printf("  Free: %.2f GB\n", float64(usage.Free)/1e9)
	fmt.Printf("  Used Percent: %.2f%%", usage.UsedPercent)
	fmt.Println()
}

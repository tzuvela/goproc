package system

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
)

func PrintDiskUsage() {
	usage, err := disk.Usage("/")
	if err != nil {
		fmt.Println("Error retrieving disk usage:", err)
		return
	}

	fmt.Println("Disk Usage (/):")
	fmt.Printf("  Total: %.2f GB\n", float64(usage.Total)/1e9)
	fmt.Printf("  Used:  %.2f GB\n", float64(usage.Used)/1e9)
	fmt.Printf("  Free:  %.2f GB\n", float64(usage.Free)/1e9)
	fmt.Printf("  Used Percent: %.2f%%\n", usage.UsedPercent)
	fmt.Println()
}

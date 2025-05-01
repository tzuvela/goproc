package system

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/host"
)

func PrintHostInfo() {
	info, err := host.Info()
	if err != nil {
		fmt.Println("Error retrieving host info:", err)
		return
	}

	fmt.Println("OS Info:")
	fmt.Printf("  OS: %s\n", info.OS)
	fmt.Printf("  Platform: %s %s\n", info.Platform, info.PlatformVersion)
	fmt.Printf("  Uptime: %s\n", formatUptime(info.Uptime))
	fmt.Println()
}

func formatUptime(seconds uint64) string {
	uptime := time.Duration(seconds) * time.Second
	return uptime.String()
}

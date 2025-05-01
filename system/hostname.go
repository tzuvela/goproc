package system

import (
	"fmt"
	"os"
)

func PrintHostname() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error retrieving hostname:", err)
		return
	}

	fmt.Printf("Hostname: %s\n", hostname)
	fmt.Println()
}

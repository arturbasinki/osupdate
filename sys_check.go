package main

import (
	"os"
	"strings"
)

// This a function that cheks what a linux distribution is running
func checkLinuxDistro() string {
	// Read the /etc/os-release file to determine the Linux distribution
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "Unknown Linux Distribution"
	}

	// Convert the data to a string and split it into lines
	lines := strings.SplitSeq(string(data), "\n")

	// Look for the line that starts with "ID="
	for line := range lines {
		if strings.HasPrefix(line, "ID=") {
			// Extract the ID value and return it
			return strings.Trim(strings.TrimPrefix(line, "ID="), "\"")
		}
	}

	return "Unknown Linux Distribution"
}

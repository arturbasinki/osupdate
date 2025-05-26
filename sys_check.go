package main

import (
	"os"
	"strings"
)

// This a function that cheks what a linux distribution is running
// and returns strig wit appropriate command to update and upgrade given distro
func checkLinuxDistro() (string, error) {
	// Read the /etc/os-release file to determine the Linux distribution
	distro := "Unknown Linux Distribution"
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return distro, err
	}

	// Convert the data to a string and split it into lines
	lines := strings.SplitSeq(string(data), "\n")

	// Look for the line that starts with "ID="
	for line := range lines {
		if strings.HasPrefix(line, "ID=") {
			// Extract the ID value and return it
			distro = strings.Trim(strings.TrimPrefix(line, "ID="), "\"")
			return distro, nil
		}
	}

	return distro, nil
}

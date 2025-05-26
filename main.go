package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	distro, err := checkLinuxDistro()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error detecting Linux distribution: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Detected Linux distribution: %s\n", distro)
	// Compose the command based on the detected distribution
	command := composeCommand(distro)
	if command == nil {
		fmt.Fprintf(os.Stderr, "Unsupported Linux distribution: %s\n", distro)
		os.Exit(1)
	}
	// Print the command to be executed
	fmt.Printf("Command to be executed: %s\n", command)
	// Print the message to the user
	fmt.Println("This script will update and upgrade your Linux system.")
	fmt.Println("Starting os update...")
	fmt.Println("You may be required to provide an administrator password (sudo).")
	fmt.Println("----------------------------------------------------------------")

	// Run appriopriate command to update the package list
	updateCmd := command
	// Set the standard output, error, and input for the command
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr
	updateCmd.Stdin = os.Stdin

	err = updateCmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during 'apt update': %v\n", err)
		os.Exit(1)
	}

	// Run "sudo apt dist-upgrade -y"
	upgradeCmd := exec.Command("sudo", "apt", "dist-upgrade", "-y")
	upgradeCmd.Stdout = os.Stdout
	upgradeCmd.Stderr = os.Stderr
	upgradeCmd.Stdin = os.Stdin

	err = upgradeCmd.Run()
	fmt.Println("---------------------------------------------------------------")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during 'apt dist-upgrade': %v\n", err)
		os.Exit(1)
	}

	// If there were no errors
	fmt.Println("System update and upgrade completed successfully.")
}

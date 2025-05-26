package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Checking Linux distro
	distro := checkLinuxDistro()
	if distro != "debian" && distro != "ubuntu" {
		fmt.Fprintf(os.Stderr, "This script is designed for Debian or Ubuntu-based systems. Detected: %s\n", distro)
		os.Exit(1)
	}
	fmt.Println("Starting os update...")
	fmt.Println("Running command: sudo apt update && sudo apt dist-upgrade -y")
	fmt.Println("You may be required to provide an administrator password (sudo).")
	fmt.Println("----------------------------------------------------------------")

	// Run "sudo apt update"
	updateCmd := exec.Command("sudo", "apt", "update")
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr
	updateCmd.Stdin = os.Stdin

	err := updateCmd.Run()
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

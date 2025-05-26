package main

import (
	"fmt"
	"os"
)

func main() {

	distro, err := checkLinuxDistro()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error detecting Linux distribution: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Detected Linux distribution: %s\n", distro)
	// Compose the command based on the detected distribution
	commands := composeCommand(distro)
	if commands == nil {
		fmt.Fprintf(os.Stderr, "Unsupported Linux distribution: %s\n", distro)
		os.Exit(1)
	}
	// Print the message to the user
	fmt.Println("This script will update and upgrade your Linux system.")
	fmt.Println("Starting os update...")
	fmt.Println("You may be required to provide an administrator password (sudo).")
	fmt.Println("----------------------------------------------------------------")

	for _, command := range commands {

		cmd := command
		// Set the standard output, error, and input for the command
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		err = cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error during %s: %v\n", cmd.String(), err)
			os.Exit(1)
		}
	}
	fmt.Println("----------------------------------------------------------------")

	// If there were no errors
	fmt.Println("System update and upgrade completed successfully.")
}

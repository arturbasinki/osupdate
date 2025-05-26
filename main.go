package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Starting os update...")
	fmt.Println("Running command: sudo apt update && sudo apt dist-upgrade -y")
	fmt.Println("You may be required to provide an administrator password (sudo).")
	fmt.Println("----------------------------------------------------------------")

	// Prepare the command `sudo apt update && sudo apt dist-upgrade -y`
	// We use "-y" to automatically confirm all apt questions.
	cmd := exec.Command("sudo", "apt", "update", "&&", "sudo", "apt", "dist-upgrade", "-y")

	// Connect Go's standard output, error, and input
	// to the appropriate streams of the command being run.
	// This allows you to display apt results directly in the terminal
	// and allows you to enter your sudo password if required.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Run the command
	err := cmd.Run()

	fmt.Println("---------------------------------------------------------------")

	// Check if there was an error while running the command

	if err != nil {
		// If cmd.Run() returns an error, it means there was a problem starting
		// the process (e.g., missing `sudo` or `apt`) or the process exited
		// with a non-zero error code.
		fmt.Fprintf(os.Stderr, "Error during update: %v\n", err)
		// Exit the program with an error code
		os.Exit(1)
	}

	// If there were no errors
	fmt.Println("System update completed successfully.")
}

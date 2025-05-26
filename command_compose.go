package main

import "os/exec"

func composeCommand(distro string) *exec.Cmd {
	// Compose the command based on the detected distribution
	switch distro {
	case "debian", "ubuntu":
		return exec.Command("sudo", "apt", "update")
	case "arch", "manjaro":
		return exec.Command("sudo", "pacman", "-Syu")
	case "opensuse":
		return exec.Command("sudo", "zypper", "refresh")
	case "fedora":
		return exec.Command("sudo", "dnf", "upgrade", "--refresh", "-y")
	case "centos", "rhel":
		return exec.Command("sudo", "yum", "update", "-y")
	default:
		return nil
	}
}

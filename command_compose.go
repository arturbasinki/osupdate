package main

import "os/exec"

func composeCommand(distro string) []*exec.Cmd {
	// Compose the command based on the detected distribution
	var commands []*exec.Cmd
	switch distro {
	case "debian", "ubuntu":
		commands = append(commands, exec.Command("sudo", "apt", "update"))
		commands = append(commands, exec.Command("sudo", "apt", "dist-upgrade", "-y"))

	case "arch", "manjaro":
		commands = append(commands, exec.Command("pacman", "-Syu"))
	case "opensuse":
		commands = append(commands, exec.Command("sudo", "zypper", "refresh"))
	case "fedora":
		commands = append(commands, exec.Command("sudo", "dnf", "upgrade", "--refresh", "-y"))
	case "centos", "rhel":
		commands = append(commands, exec.Command("sudo", "yum", "update", "-y"))
	default:
		return nil
	}
	return commands
}

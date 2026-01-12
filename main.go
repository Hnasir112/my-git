package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func saveChanges(message string) {
	fmt.Println("Adding files...")
	if err := runCommand("git", "add", "."); err != nil {
		fmt.Printf("Error adding files: %v\n", err)
		return
	}

	fmt.Printf("Commiting with message: %s\n", message)
	if err := runCommand("git", "commit", "-m", message); err != nil {
		fmt.Printf("Error committing changes: %v\n", err)
		return
	}

	fmt.Printf("Pushing to remote...")
	err := runCommand("git", "push")
	if err != nil {
		if err := runCommand("git", "push", "u", "origin", "main"); err != nil {
			fmt.Printf("Error pushing changes: %v\n", err)
			return
		}
	}

	fmt.Println("Done!")
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: mygit save \"commit message\"")
		os.Exit(1)
	}

	command := os.Args[1]
	if command == "save" {
		message := os.Args[2]
		saveChanges(message)
	} else {
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}

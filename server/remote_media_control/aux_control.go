package main

import (
	"log"
	"os"
	"os/exec"
)

func playPause() {
	pressAuxKey()
}

//todo: Make it configurable
func pressAuxKey() {
	// Command to run the Python script
	cmd := exec.Command("aux_control", "test.py")

	// Set the working directory if needed
	// cmd.Dir = "/path/to/your/python/script"

	// Capture the output (stdout and stderr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Default().Printf("Error: %v\n", err)
		return
	}
}
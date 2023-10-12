package media

import (
	"os"
	"os/exec"
)

func PlayPause() error {
	return pressAuxKey()
}

//todo: Make it configurable
func pressAuxKey() error {
	// Command to run the Python script
	cmd := exec.Command("python", "aux_control.py")

	// Set the working directory if needed
	// cmd.Dir = "/path/to/your/python/script"

	// Capture the output (stdout and stderr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	return cmd.Run()
}
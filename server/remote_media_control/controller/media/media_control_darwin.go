package media

import (
	"fmt"
	"os"
	"os/exec"
)

func PlayPause() error {
	return pressAuxKey()
}

//todo: Make it configurable
func pressAuxKey() error {
    path, err := os.Getwd() 
	if err != nil {
		return err
	}
	// Command to run the Python script
	cmd := exec.Command("python", fmt.Sprintf("%s/controller/media/aux_control.py", path))

	// Set the working directory if needed
	// cmd.Dir = "/path/to/your/python/script"

	// Capture the output (stdout and stderr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	return cmd.Run()
}
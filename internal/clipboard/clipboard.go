package clipboard

import (
	"os/exec"
)

// Copy copies the given text to the clipboard. It uses the pbcopy command on macOS.
func Copy(text string) error {
	copyCmd := exec.Command("pbcopy")
	in, err := copyCmd.StdinPipe()
	if err != nil {
		return err
	}
	if err := copyCmd.Start(); err != nil {
		return err
	}
	if _, err := in.Write([]byte(text)); err != nil {
		return err
	}
	if err := in.Close(); err != nil {
		return err
	}
	return nil
}

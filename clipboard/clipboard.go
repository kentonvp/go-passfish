package clipboard

import (
	"fmt"
	"os/exec"
)

func CopyToClipboard(text string) (bool, error) {
	fmt.Println("Copying password to clipboard...")
	copyCmd := exec.Command("pbcopy")
	in, err := copyCmd.StdinPipe()
	if err != nil {
		return false, err
	}
	if err := copyCmd.Start(); err != nil {
		return false, err
	}
	if _, err := in.Write([]byte(text)); err != nil {
		return false, err
	}
	if err := in.Close(); err != nil {
		return false, err
	}
	return true, nil
}

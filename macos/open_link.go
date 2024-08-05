package macos

import (
	"os/exec"
)

// open link in default browser
func OpenLink(link string) error {
	cmd := exec.Command("open", "-u", link)
	if err := cmd.Run(); err != nil {
		return err
	} else {
		return nil
	}
}

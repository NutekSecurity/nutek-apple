package macos

import (
	"github.com/nuteksecurity/nutek-apple/util"
)

// Install selected program using Homebrew
// or local package manager. Return true on successful
// installation, if program is already installed - false, or error
func Install(program string) (bool, error) {
	if !util.IsInstalled(program) {
		_, err := BrewCmd("install", program)
		if err != nil {
			return false, err
		}
		return true, nil
	} else {
		return false, nil
	}
}

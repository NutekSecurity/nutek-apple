package macos

import (
	"fmt"

	"github.com/nuteksecurity/nutek-apple/util"
)

// Install selected program using Homebrew
// or local package manager. Return true on successful
// installation, if program is already installed - false, or error
func Install(programs ...string) []error {
	var installList []string = []string{}
	var errorList []error = []error{}
	for _, program := range programs {
		// lists handling

		if !util.IsInstalled(program) {
			installList = append(installList, program)
		}
	}
	for _, program := range installList {
		_, err := BrewCmd("install", program)
		if err != nil {
			errorList = append(errorList, fmt.Errorf("error: %s, when installing program: %s", err, program))
		}
	}
	if len(errorList) > 0 {
		return errorList
	}
	return nil
}

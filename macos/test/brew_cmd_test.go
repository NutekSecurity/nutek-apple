package macos

import (
	"fmt"
	"strings"
	"testing"

	macos "github.com/nuteksecurity/nutek-apple/macos"
)

func TestBrewCmd(t *testing.T) {
	cliPrograms := macos.Cli()
	for _, cliProgram := range cliPrograms {
		cmdOutput, err := macos.BrewCmd("info", cliProgram)
		if err != nil {
			t.Error(err)
		}
		lookFor := fmt.Sprintf("==> %s:", cliProgram)
		if !strings.Contains(cmdOutput, lookFor) {
			t.Errorf("can't find %s", cliProgram)
		}
	}
	guiPrograms := macos.Gui()
	for _, guiProgram := range guiPrograms {
		cmdOutput, err := macos.BrewCmd("info", guiProgram)
		if err != nil {
			t.Error(err)
		}
		lookFor := fmt.Sprintf("==> %s:", guiProgram)
		if !strings.Contains(cmdOutput, lookFor) {
			t.Errorf("can't find %s", guiProgram)
		}
	}
}

package macos

import (
	"testing"

	"github.com/nuteksecurity/nutek-apple/macos"
	"github.com/nuteksecurity/nutek-apple/util"
)

func TestInstallOllama(t *testing.T) {
	err := macos.Install("ollama")
	if err != nil {
		t.Error(err)
	}
	if !util.IsInstalled("ollama") {
		t.Errorf("error: ollama did not installed on macOS")
	}
}

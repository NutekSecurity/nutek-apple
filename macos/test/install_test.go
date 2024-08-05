package macos

import (
	"testing"

	"github.com/nuteksecurity/nutek-apple/macos"
)

func TestInstallOllama(t *testing.T) {
	result, err := macos.Install("ollama")
	if err != nil {
		t.Error(err)
	}
	if !result {
		t.Errorf("already installed: %s", "ollama")
	}
}

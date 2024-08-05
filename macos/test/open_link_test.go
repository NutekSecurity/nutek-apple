package macos

import (
	"testing"

	"github.com/nuteksecurity/nutek-apple/macos"
)

// this test might fail without a default browser
func TestOpenLink(t *testing.T) {
	// Test that calling OpenLink with a valid link works
	link := "https://example.com"
	err := macos.OpenLink(link)
	if err != nil {
		t.Errorf("OpenLink(%q) returned error: %v", link, err)
	}
}

package macos

import (
	"bytes"
	"strings"
	"sync"
	"testing"

	"github.com/nuteksecurity/nutek-apple/util"
)

func TestIsInstalled(t *testing.T) {
	if !util.IsInstalled("pwd") {
		t.Errorf("not installed: %s", "pwd")
	}
}

func TestPrintOutput(t *testing.T) {
	// Create a mock output reader that returns some sample data
	// mu := &sync.Mutex{}
	util.Ch = make(chan string)
	wg := &sync.WaitGroup{}

	out := bytes.NewReader([]byte("Hello, World!"))

	wg.Add(1)
	go func() {
		util.PrintOutput(out)
		close(util.Ch) // close the channel when done
	}()

	var fullOutput strings.Builder
	for line := range util.Ch { // receive all output lines from the channel
		fullOutput.WriteString(line)
	}

	wg.Done()
	// Wait for the goroutine to finish
	wg.Wait()

	if fullOutput.String() != "Hello, World!" {
		t.Errorf("Expected %s, but have: %v on channel", "Hello, World!", fullOutput.String())
	}

	util.Ch = make(chan string)

	out = bytes.NewReader([]byte(""))

	wg.Add(1)
	go func() {
		util.PrintOutput(out)
		close(util.Ch) // close the channel when done
	}()

	var fullOutput2 strings.Builder
	for line := range util.Ch { // receive all output lines from the channel
		fullOutput2.WriteString(line)
	}

	wg.Done()
	// Wait for the goroutine to finish
	wg.Wait()

	if fullOutput2.String() != "" {
		t.Errorf("Expected %s, but have: %v on channel", "", fullOutput2.String())
	}
}

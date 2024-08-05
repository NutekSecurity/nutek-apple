package macos

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/nuteksecurity/nutek-apple/util"
)

// use Homebrew and add command line arguments
func BrewCmd(args ...string) (string, error) {
	cmd := "brew"
	// args := []string{"arg1", "arg2"}
	output, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		myError := fmt.Sprintf("Error: %s\n", err.Error())
		return myError, err
	}

	util.Ch = make(chan string)

	var wg sync.WaitGroup
	wg.Add(1) // add a WaitGroup to wait for the output goroutine

	go func() {
		util.PrintOutput(bytes.NewReader(output))
		close(util.Ch) // close the channel when done
	}()

	var fullOutput strings.Builder
	for line := range util.Ch { // receive all output lines from the channel
		fullOutput.WriteString(line + "\n")
	}
	wg.Done() // signal that the output goroutine has finished

	wg.Wait() // wait for the output goroutine to finish before exiting

	return fullOutput.String(), err
}

package util

import (
	"fmt"
	"io"
	"os/exec"
	"sync"
)

var mu sync.Mutex
var Ch chan string

// print from reader to standard output
func PrintOutput(out io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := out.Read(buf)
		if n == 0 || err != nil {
			return
		}
		output := string(buf[:n])
		fmt.Println(output)
		mu.Lock()
		Ch <- output
		mu.Unlock()
	}
}

// check if program is installed
func IsInstalled(program string) bool {
	cmd := "which"
	// args := []string{"arg1", "arg2"}
	output, err := exec.Command(cmd, program).CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return false
	}
	if string(output) != "" {
		return true
	} else {
		return false
	}
}

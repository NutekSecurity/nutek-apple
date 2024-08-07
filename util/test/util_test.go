package macos

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/nuteksecurity/nutek-apple/structs"
	"github.com/nuteksecurity/nutek-apple/util"
	"github.com/stretchr/testify/assert"
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

func TestWalkDir(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "walkdir-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a directory structure like this:
	//
	// walkdir-test/
	//   author1/theme1/list.json
	//   links/links.go
	//   author2/theme2/
	//     list.json
	//     subtheme1/
	//       subtheme1_list.json

	if err := os.MkdirAll(tempDir+"/author1/theme1", 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(tempDir+"/author1/theme1/list.json", []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	if err := os.MkdirAll(tempDir+"/author2/theme2", 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(tempDir+"/author2/theme2/list.json", []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	if err := os.MkdirAll(tempDir+"/links", 0755); err != nil {
		t.Fatal(err)
	}
	// Create the "links/links.go" file to skip this directory
	if err := os.WriteFile(tempDir+"/links/links.go", []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	// Call WalkDir and check that it prints the expected output
	walkedDir := tempDir + "/author2/theme2"
	err = func() (err error) {
		util.WalkDir(walkedDir)
		return err
	}()
	if err != nil {
		t.Fatal(err)
	}

	// Check that WalkDir doesn't print anything if we skip the current directory
	tempDir2, err := os.MkdirTemp("", "walkdir-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir2)

	if err := os.WriteFile(tempDir2+"/links.go", []byte{}, 0644); err != nil {
		t.Fatal(err)
	}
	walkedDir = tempDir2
	err = func() (err error) {
		util.WalkDir(walkedDir)
		return err
	}()
	if err != nil {
		t.Fatal(err)
	}
}

func TestWalkDirNoLinks(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "walkdir-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a directory structure like this:
	//
	// walkdir-test/
	//   author1/theme1/list.json

	if err := os.MkdirAll(tempDir+"/author1/theme1", 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(tempDir+"/author1/theme1/list.json", []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	walkedDir := tempDir
	err = func() (err error) {
		util.WalkDir(walkedDir)
		return err
	}()
	if err != nil {
		t.Fatal(err)
	}
}

func TestWalkDirMultipleDirectories(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "walkdir-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a directory structure like this:
	//
	// walkdir-test/
	//   author1/theme1/list.json
	//   author2/theme2/
	//     list.json

	if err := os.MkdirAll(tempDir+"/author1/theme1", 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(tempDir+"/author1/theme1/list.json", []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	if err := os.MkdirAll(tempDir+"/author2/theme2", 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(tempDir+"/author2/theme2/list.json", []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	walkedDir := tempDir
	err = func() (err error) {
		util.WalkDir(walkedDir)
		return err
	}()
	if err != nil {
		t.Fatal(err)
	}
}

func TestWalkDirNoJSON(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "walkdir-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a directory structure like this:
	//
	// walkdir-test/
	//   author1/theme1/

	if err := os.MkdirAll(tempDir+"/author1/theme1", 0755); err != nil {
		t.Fatal(err)
	}

	walkedDir := tempDir + "/author1/theme1"
	err = func() (err error) {
		util.WalkDir(walkedDir)
		return err
	}()
	if err != nil {
		t.Fatal(err)
	}
}

func TestWalkDirNoFiles(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "walkdir-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	walkedDir := tempDir
	err = func() (err error) {
		util.WalkDir(walkedDir)
		return err
	}()
	if err != nil {
		t.Fatal(err)
	}
}

func TestProjectRoot(t *testing.T) {
	// Test case 1: NUTEK_APPLE_ROOT is set and matches /var
	os.Setenv("NUTEK_APPLE_ROOT", "/var")
	root, err := util.GetProjectRoot()
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	expectedRoot := "/var"
	if root != expectedRoot {
		t.Errorf("expected %q, but got %q", expectedRoot, root)
	}

	// Test case 2: NUTEK_APPLE_ROOT is set and does not match /var
	os.Setenv("NUTEK_APPLE_ROOT", "/path/to/project")
	root, err = util.GetProjectRoot()
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	expectedRoot = "/path/to/project"
	if root != expectedRoot {
		t.Errorf("expected %q, but got %q", expectedRoot, root)
	}

	// Test case 3: NUTEK_APPLE_ROOT is not set
	os.Unsetenv("NUTEK_APPLE_ROOT")
	binPath, err := filepath.Abs(os.Args[0])
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	wdPath, err := os.Getwd()
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	unsetRoot, err := util.GetProjectRoot()
	if err != nil {
		t.Error(err)
	}
	if strings.Contains(binPath, unsetRoot) {
		t.Errorf("should not point to /var folder")
	}
	if !strings.Contains(wdPath, unsetRoot) {
		t.Errorf("%s should partially match %s", unsetRoot, wdPath)
	}

}

func TestListLinksInFile(t *testing.T) {
	result, err := util.ListLinksInFile("util", "test/example")
	if !assert.Nil(t, err) {
		t.Error(err)
	}
	expectedOutput := `This is the list you were looking for...

example.com: A simple link to example.com
  URL: https://example.com

golang.org: The official Go website
  URL: https://golang.org

github.com: The official GitHub website
  URL: https://github.com

Use "all" to list all bookmarks, "search" to look for author/theme/link...
`
	assert.Equal(t, expectedOutput, result)
}

func TestSearchLinksInDirs(t *testing.T) {
	rootDir := "util" // Update with your test data directory
	searchTerm := "ing"

	check1 := "seclists"

	check2 := "A list of useful payloads and bypass for Web Application Security and Pentest/CTF"

	check3 := "https://book.hacktricks.xyz"

	result, err := util.SearchLinksInDirs(rootDir, searchTerm)
	assert.Nil(t, err)

	assert.Contains(t, result, check1)
	assert.Contains(t, result, check2)
	assert.Contains(t, result, check3)
}

func TestSearchLinks(t *testing.T) {
	testData := structs.Links{
		Links: map[string]structs.Link{
			"link1":        {Description: "This is a test link", URL: "https://example.com/"},
			"another_link": {Description: "Another test link", URL: "https://www.google.com/"},
			"irrelevant":   {Description: "This has no match", URL: "https://invalid.com/"},
		},
	}

	keys, hits := util.GetSearchLinks(testData, "test")
	assert.Equal(t, []string{"link1", "another_link"}, keys)
	assert.Equal(t, []structs.Link{{Description: "This is a test link", URL: "https://example.com/"},
		{Description: "Another test link", URL: "https://www.google.com/"}}, hits)

	keys, hits = util.GetSearchLinks(testData, "another")
	assert.Equal(t, []string{"another_link"}, keys)
	assert.Equal(t, []structs.Link{{Description: "Another test link", URL: "https://www.google.com/"}}, hits)

	keys, hits = util.GetSearchLinks(testData, "no match")
	assert.Empty(t, keys)
	assert.Empty(t, hits)
}

func TestFieldsMatch(t *testing.T) {
	testCases := []struct {
		text     string
		query    string
		expected bool
	}{
		{"This is a test string", "test", true},
		{"Hello World!", "world", true},
		{"Go programming language", "programming", true},
		{"Another example sentence", "example", true},
		{"No match here", "python", false},
	}

	for _, tc := range testCases {
		result := util.GetFieldsMatch(tc.text, tc.query)
		assert.Equal(t, tc.expected, result, "Test case: %s, query: %s", tc.text, tc.query)
	}
}

func TestWalkLinksDir(t *testing.T) {
	linksChan := make(chan structs.Links)

	// Start a goroutine to walk the directories and send Links to the channel
	go func() {
		err := util.GetWalkLinksDir(linksChan, "util/test")
		if err != nil {
			t.Errorf("Error walking directories: %s", err)
		}
		close(linksChan)
	}()

	// Process links from the channel and search for matches
	for links := range linksChan {
		assert.NotEmpty(t, links.Links)
	}
}

func TestUpdate(t *testing.T) {
	err := util.Update([]string{"noupdate"}, func() error {
		return nil
	})
	assert.NoError(t, err, "have %s, when passing 'noupdate' argument", err)
	err = util.Update([]string{}, func() error {
		return nil
	})
	assert.NoError(t, err, "have %s, when blank parameters array and blank function", err)
	err = util.Update([]string{}, func() error {
		fmt.Println("Checking update...")
		return nil
	})
	assert.NoError(t, err, "have %s, when blank parameters array and printing from function", err)
	err = util.Update([]string{"some"}, func() error {
		fmt.Println("Checking update...")
		return nil
	})
	assert.NoError(t, err, "have %s, when some parameters array and printing from function", err)
}

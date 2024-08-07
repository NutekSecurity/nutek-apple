package util

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/nuteksecurity/nutek-apple/structs"
)

// def check_version_and_update
//   response = `git pull origin main --rebase`
//   puts response
//   if response.include?('error')
//     puts '❌ Error: Could not update from the repository.'
//     puts 'hint: You can run the script with --no-update to skip the update check.'
//     exit
//   elsif response.include? 'ahead'
//     puts 'hint: You can run the script with --no-update to skip the update check.'
//     on_your_own = get_yes_no_input "You're working on your own version of nutek-apple 🍎 Do you want to continue? (yes/no): "
//     return if on_your_own

//     puts 'Exit'
//     exit
//   elsif response.include?('Already up to date.')
//       puts '✅ Already up to date.'
//   elsif response.include?('Updating ')
//     puts '✅ Updated successfully. Please restart the script.'
//     exit
//   else
//     puts '❌ Error: Git command failed.'
//     puts 'hint: You can run the script with --no-update to skip the update check.'
//     exit
//   end
// end

// def update(args)
//   return if args.include?('--no-update')

//   puts 'Updating...'
//   # check if we're in a git repo
//   if File.directory?('.git')
//     # check if we're in the right repo
//     if `git remote get-url origin`.chomp == 'https://github.com/NutekSecurity/nutek-apple.git' || `git remote get-url upstream`.chomp == 'https://github.com/NutekSecurity/nutek-apple.git'
//       check_version_and_update
//     else
//       puts '❌ Error: Not in the right git repository, not updating.'
//       puts 'hint: You can try to run the script with --no-update to skip the update check.'
//       exit
//     end
//   else
//     puts '❌ Error: Not in a git repoository, not updating.'
//     puts 'hint: You can try to run the script with --no-update to skip the update check.'
//     exit
//   end
// end

func Update(argsSlice []string, runme func() error) error {
	for _, arg := range argsSlice {
		if arg == "noupdate" {
			return nil
		}
	}
	projectRoot, err := projectRoot()
	if err != nil {
		return err
	}
	paths, err := os.ReadDir(projectRoot)
	if err != nil {
		return err
	}
	for _, path := range paths {
		if path.Name() == ".git" && path.IsDir() {
			remoteOrigin := exec.Command("git", "remote", "get-url", "origin")
			remoteUpstream := exec.Command("git", "remote", "get-url", "upstream")
			orginOutput, err := remoteOrigin.Output()
			var gitRemoteErrors []string = []string{}
			if err != nil {
				if err.Error() == "error: exit status 2" && string(orginOutput) == "" {
					fmt.Println("git origin remote do not match Nutek Apple repository")
				} else {
					gitRemoteErrors = append(gitRemoteErrors, fmt.Sprintf("error: %s, when using git remote get-url origin output: %s", err, orginOutput))
				}
			}
			upstreamOutput, err := remoteUpstream.Output()
			if err != nil {
				if err.Error() == "error: exit status 2" && string(orginOutput) == "" {
					fmt.Println("git upstream remote do not match Nutek Apple repository")
				} else {
					gitRemoteErrors = append(gitRemoteErrors, fmt.Sprintf("error: %s, when using git remote get-url origin output: %s", err, upstreamOutput))
				}
			}
			if len(gitRemoteErrors) == 2 {
				return fmt.Errorf("error: %s and error: %s", gitRemoteErrors[0], gitRemoteErrors[1])
			}
			if string(orginOutput) == "https://github.com/NutekSecurity/nutek-apple.git" || string(upstreamOutput) == "https://github.com/NutekSecurity/nutek-apple.git" ||
				string(orginOutput) == "https://github.com/nuteksecurity/nutek-apple.git" || string(upstreamOutput) == "https://github.com/nuteksecurity/nutek-apple.git" {
				// check_version_and_update
			}
			runme()
			return nil
		}
	}
	return fmt.Errorf("error: did not update Nutek Apple git repository")
}

var mu sync.Mutex
var Ch chan string

// PrintOutput reads output from an io.Reader and prints it to the console.
//
// The function takes one parameter: the input `out` io.Reader, which provides output data.
// It uses a buffer to store chunks of output data, allowing for efficient reading and processing.
// Once the output is exhausted (i.e., Read returns 0 or an error), the function returns.
// Otherwise, it converts each chunk of output data into a string, prints it to the console using fmt.Println,
// and sends the output to a channel using mu.Lock() and Ch <- output.
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

// IsInstalled checks whether a specified program is installed on the system.
//
// The function takes one parameter: the input `program` string specifying the name of the program to check for
// installation.
// It uses the `exec.Command` function to execute the "which" command, which returns information about the location
// of executable files.
// If an error occurs during execution (e.g., if the "which" command is not found), the function prints an error
// message and returns false.
// Otherwise, it checks whether the output from the "which" command is non-empty. If it is, the program is
// installed, and the function returns true; otherwise, it returns false.
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

// WalkDir recursively walks through a directory tree, printing information about each file's author/theme.
//
// The function starts at the given `dir` and traverses down into its subdirectories. For each file found,
// it checks if the file is a JSON file (i.e., it ends with ".json") and prints out its path in the format
// "author\\theme/list".
//
// If the current directory being walked contains a file named "links.go", it skips that entire directory.
func WalkDir(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range files {
		path := dir + "/" + fi.Name()
		if path == "links/links.go" {
			continue
		}
		splitPath := strings.Split(path, "/")[1]
		fmt.Printf("Author\\theme: %s\n", splitPath)
		filesInFolder, _ := os.ReadDir(path)
		for _, file := range filesInFolder {
			if file.IsDir() {
				WalkDir(path)
			}
			fileName := strings.Split(file.Name(), ".json")[0]
			fmt.Printf("  List: %s/%s\n", splitPath, fileName)
		}
	}
}

// projectRoot returns the path to the project's root directory.
//
// The function first checks if the "NUTEK_APPLE_ROOT" environment variable is set. If it is,
// the value of that variable is returned as the project's root path.
//
// If the environment variable is not set, the function attempts to determine the project's
// root path by examining the binary's location and the current working directory. In this case,
// the project's root path is considered to be either the current working directory or the directory
// containing the binary (depending on whether the binary resides in a "var" directory).
func projectRoot() (string, error) {
	root := os.Getenv("NUTEK_APPLE_ROOT")
	if root == "" {
		binPath, err := filepath.Abs(os.Args[0])
		if err != nil {
			return "", err
		}
		wdPath, err := os.Getwd()
		if err != nil {
			return "", err
		}
		if strings.HasPrefix(binPath, "/var") {
			wdPathSplit := strings.Split(wdPath, "/")
			wdPath = "/"
			for _, path := range wdPathSplit {
				if path == "nutek-apple" {
					wdPath = filepath.Join(wdPath, path)
					return wdPath, nil
				}
				wdPath = filepath.Join(wdPath, path)
			}
		} else {
			binPathSplit := strings.Split(binPath, "/")
			binPath = "/"
			for _, path := range binPathSplit {
				if path == "nutek-apple" {
					binPath = filepath.Join(binPath, path)
					return binPath, nil
				}
				binPath = filepath.Join(binPath, path)
			}
		}
	}
	return root, nil
}

// return local function for testing purposes
func GetProjectRoot() (string, error) {
	return projectRoot()
}

// ListLinksInFile retrieves links from a JSON file matching the given search term.
//
// The function first constructs the path to the links file using the environment variables
// "HOME" and ".nutek-apple", as well as the search term. If the links file does not exist at that location,
// it falls back to searching for a directory named ".nutek-apple" in the project root, and constructs
// the path to the links file based on the current working directory.
//
// The function then attempts to read the contents of the links file as JSON data. If successful,
// it unmarshals the JSON into a `structs.Links` object and prints out the list of links in the format:
//
//	<key>: <description>
//	    URL: <url>
func ListLinksInFile(rootDir string, searchTerm string) (string, error) {
	linksFile := filepath.Join(os.Getenv("HOME"), ".nutek-apple", rootDir, searchTerm)

	if _, err := os.Stat(linksFile); os.IsNotExist(err) {
		linksFile = filepath.Join(rootDir, searchTerm)
		pwd, err := projectRoot()
		if err != nil {
			return "", err
		}
		pwdSplit := strings.Split(pwd, "/")
		var myPwd string = "/"
		for _, pathPart := range pwdSplit {
			if pathPart == "nutek-apple" {
				myPwd = filepath.Join(myPwd, pathPart)
				break
			}
			myPwd = filepath.Join(myPwd, pathPart)
		}
		linksFile = filepath.Join(myPwd, linksFile)
	}

	if strings.HasSuffix(linksFile, ".go") {
		return "", fmt.Errorf("error: this is not a .json file")
	}

	if !strings.HasSuffix(linksFile, ".json") {
		linksFile = linksFile + ".json"
	}

	data, err := os.ReadFile(linksFile)
	if err != nil {
		return "", fmt.Errorf("error reading file %s: %w", linksFile, err)
	}

	var links structs.Links
	err = json.Unmarshal(data, &links)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling JSON from file %s: %w", linksFile, err)
	}

	var result string = ""

	result = result + fmt.Sprintln("This is the list you were looking for...")
	result = result + fmt.Sprintln("")
	for key, link := range links.Links {
		result = result + fmt.Sprintf("%s: %s\n  URL: %s\n\n", key, link.Description, EscapeHTMLInUrl(link.URL))
	}
	result = result + fmt.Sprintln("Use \"all\" to list all bookmarks, \"search\" to look for author/theme/link...")

	return result, nil
}

// SearchLinksInDirs searches for links within all directories starting from the specified root directory that match
// the given search term.
//
// The function takes two parameters: the input `rootDir` string specifying the root directory to search in, and the
// `searchTerm` string specifying the term to search for.
// It creates a channel to receive links from the `walkLinksDir` function and initializes an empty slice to store
// matches.
// The function then starts a goroutine to walk through all directories using the `walkLinksDir` function, which
// sends loaded links over the channel.
// Once the walking process is complete (i.e., the channel is closed), the function iterates over all links received
// from the channel and searches for matches within each link using the `searchLinks` function.
// The function returns nil error.
func SearchLinksInDirs(rootDir string, searchTerm string) (string, error) {
	linksChan := make(chan structs.Links)
	hits := structs.Links{
		Links: make(map[string]structs.Link),
	}

	// Start a goroutine to walk the directories and send Links to the channel
	go func() {
		err := walkLinksDir(linksChan, rootDir)
		if err != nil {
			fmt.Println("Error walking directories:", err)
		}
		close(linksChan)
	}()

	// Process links from the channel and search for matches
	for links := range linksChan {
		newKeys, newLinks := searchLinks(links, searchTerm)
		for i, key := range newKeys {
			hits.Links[key] = newLinks[i]
		}
	}

	var result string = ""
	result = result + "Found this links:\n\n"
	for key, hit := range hits.Links {
		result = result + fmt.Sprintf("%s: %s\n  URL: %s\n", key, hit.Description, EscapeHTMLInUrl(hit.URL))
	}

	return result, nil
}

// searchLinks searches for matches to the given search term within the provided links.
//
// The function takes two parameters: the input `links` struct and the `searchTerm` string.
// It iterates over each link in the `links` struct, checks whether its description or key contains the search term
// using the `fieldsMatch` function, and appends a match to the `hits` slice if a match is found.
// The function then returns the completed `hits`slice.
func searchLinks(links structs.Links, searchTerm string) (keys []string, hits []structs.Link) {
	for key, link := range links.Links {
		matches := fieldsMatch(link.Description, searchTerm) ||
			strings.Contains(strings.ToLower(key), searchTerm)
		if matches {
			hits = append(hits, link)
			keys = append(keys, key)
		}
	}
	return keys, hits
}

func GetSearchLinks(links structs.Links, searchTerm string) ([]string, []structs.Link) {
	return searchLinks(links, searchTerm)
}

// fieldsMatch checks whether any words in the given string match the specified query.
//
// The function takes two parameters: the input string `s` and the query string `query`.
// It converts both strings to lowercase and splits them into individual words using the `strings.Fields` function.
// Then, it iterates over each word in the string and checks whether it contains the query string. If any match is
// found, the function immediately returns `true`. Otherwise, it returns `false`.
func fieldsMatch(s string, query string) bool {
	for _, word := range strings.Fields(strings.ToLower(s)) {
		if strings.Contains(word, query) {
			return true
		}
	}
	return false
}

func GetFieldsMatch(s string, query string) bool {
	return fieldsMatch(s, query)
}

// walkLinksDir walks through all files within the specified directory that match the format expected by this application.
//
// The function takes two parameters: the input `ch chan<- structs.Links` channel and the `rootDir` string
// specifying the root directory to search in.
// It constructs the path to the links directory using environment variables and checks whether it exists. If not,
// it adjusts the path to match the expected relative path.
// The function then uses the `filepath.Walk` function to recursively walk through all files within the directory.
// For each file found, it checks if it's a valid JSON file by unmarshaling its contents into a `structs.Links`
// object. If successful, it sends the loaded links over the provided channel.
func walkLinksDir(ch chan<- structs.Links, rootDir string) error {
	projectRoot, err := projectRoot()
	if err != nil {
		return err
	}
	linksDir := filepath.Join(projectRoot, rootDir)

	if _, err := os.Stat(linksDir); os.IsNotExist(err) {
		linksDir = filepath.Join(rootDir) // Adjust this to the correct relative path if needed
	}

	err = filepath.Walk(linksDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if it's a file and not a directory or hidden file
		if !info.IsDir() && filepath.Ext(path) != "" && strings.HasSuffix(path, ".json") {
			data, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("error reading file %s: %w", path, err)
			}

			var links structs.Links
			err = json.Unmarshal(data, &links)
			if err != nil {
				return fmt.Errorf("error unmarshaling JSON from file %s: %w", path, err)
			}

			// Send loaded links over the channel
			ch <- links
		}

		return nil
	})

	return err
}

func GetWalkLinksDir(ch chan<- structs.Links, rootDir string) error {
	return walkLinksDir(ch, rootDir)
}

func EscapeHTMLInUrl(unescaped string) string {
	escaped := html.EscapeString(unescaped)
	return escaped
}

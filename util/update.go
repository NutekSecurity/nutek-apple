package util

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func gitUpdate(upstream bool, verbose bool) error {
	if verbose {
		fmt.Println("saving current working directory")
	}
	wdPath, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to update when getting working dir, error: %s", err)
	}
	if verbose {
		fmt.Println("searching for project root of Nutek Apple")
	}
	rootPath, err := projectRoot()
	if err != nil {
		return fmt.Errorf("failed to update when getting project root, error: %s", err)
	}
	if verbose {
		fmt.Println("going to root of Nutek Apple project")
	}
	err = os.Chdir(rootPath)
	if err != nil {
		return fmt.Errorf("error: %s, when going to directory where the Nutek Apple is before update", err)
	}
	var pull *exec.Cmd
	if verbose {
		fmt.Println("pulling possible update from the GitHub")
	}
	if upstream {
		pull = exec.Command("git", "pull", "upstream", "main", "--rebase")
	} else {
		pull = exec.Command("git", "pull", "origin", "main", "--rebase")
	}
	if verbose {
		fmt.Println("changing directory back to where Nutek Apple command was issued")
	}
	err = os.Chdir(wdPath)
	if err != nil {
		return fmt.Errorf("error: %s, when going back to directory when the Nutek Apple was invoked after update", err)
	}

	// Create a pipe to capture both stdout and stderr.
	output, errGitPull := pull.CombinedOutput()

	// Handle cases with or without a newline character
	stdoutGitPull := string(output[:len(output)])
	stderrGitPull := "" // Initialize stderr as empty string

	// Split the output based on the first newline character (if present)
	if i := bytes.IndexByte(output, '\n'); i != -1 {
		stdoutGitPull = string(output[:i])
		stderrGitPull = string(output[i:])
	} else {
		stdoutGitPull = string(output) // If no newline, stdout is the entire output
	}

	// if verbose {
	fmt.Print(stdoutGitPull)
	fmt.Print(stderrGitPull)
	// }

	isInstalled := exec.Command("which", "nutek-apple")
	output, err = isInstalled.CombinedOutput()

	var isNotInstalled bool = false
	if err != nil {
		isNotInstalled = true
		fmt.Println("Nutek Apple is not installed as standalone command. Installing...")
	}

	// Handle cases with or without a newline character
	stdout := string(output[:len(output)])
	stderr := "" // Initialize stderr as empty string

	// Split the output based on the first newline character (if present)
	if i := bytes.IndexByte(output, '\n'); i != -1 {
		stdout = string(output[:i])
		stderr = string(output[i:])
	} else {
		stdout = string(output) // If no newline, stdout is the entire output
	}

	if verbose {
		fmt.Print(stdout)
		fmt.Print(stderr)
	}

	var isDifferentVersionInstalled bool = false
	if !isNotInstalled {
		versionCmd := exec.Command("nutek-apple", "--version")
		output, err := versionCmd.Output()
		if err != nil {
			return fmt.Errorf("error: %s, when trying to determine the version of product", err)
		}
		version := strings.Split(strings.TrimSpace(string(output)), " ")[2]
		if version != Version {
			isDifferentVersionInstalled = true
		}
	}

	withError := strings.Contains(stdoutGitPull, "error") || strings.Contains(stderrGitPull, "error") || strings.Contains(stderrGitPull, "You have unstaged changes") ||
		strings.Contains(stdoutGitPull, "You have unstaged changes")
	ahead := strings.Contains(stdoutGitPull, "ahead") || strings.Contains(stderrGitPull, "ahead")
	upToDate := strings.Contains(stdoutGitPull, "up to date.") || strings.Contains(stderrGitPull, "up to date.")
	updated := strings.Contains(stdoutGitPull, "Updating ") || strings.Contains(stderrGitPull, "Updating ") || isNotInstalled || isDifferentVersionInstalled
	rebase := strings.Contains(stdoutGitPull, "rebase") || strings.Contains(stderrGitPull, "rebase")

	if withError {
		fmt.Println("hint: if you dant want to perform the update, pass 'noupdate' argument")
		if strings.Contains(stderrGitPull, "You have unstaged changes") || strings.Contains(stdoutGitPull, "You have unstaged changes") {
			return fmt.Errorf("error: %s, you have to commit your changes to the repository.\nAny program list or bookmarks have to be commited to git repository first.\nUse 'git add filename' and 'git commit -m \"descriptive comment\"'", errGitPull)
		}
		return fmt.Errorf("error: git pull command failed with error: %s", errGitPull)
	} else if updated {
		fmt.Println("building Nutek Apple...")
		if verbose {
			fmt.Println("saving current working directory")
		}
		wdPath, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to update when getting working dir, error: %s", err)
		}
		if verbose {
			fmt.Println("searching for project root of Nutek Apple")
		}
		rootPath, err := projectRoot()
		if err != nil {
			return fmt.Errorf("failed to update when getting project root, error: %s", err)
		}
		if verbose {
			fmt.Println("going to root of Nutek Apple project")
		}
		err = os.Chdir(rootPath)
		if err != nil {
			return fmt.Errorf("error: %s, when going to directory where the Nutek Apple is before update", err)
		}
		if verbose {
			fmt.Println("building new version of Nutek Apple")
		}
		runme := exec.Command("go", "build")
		output, err := runme.CombinedOutput()

		// Handle cases with or without a newline character
		stdout := string(output[:len(output)])
		stderr := "" // Initialize stderr as empty string

		// Split the output based on the first newline character (if present)
		if i := bytes.IndexByte(output, '\n'); i != -1 {
			stdout = string(output[:i])
			stderr = string(output[i:])
		} else {
			stdout = string(output) // If no newline, stdout is the entire output
		}

		if verbose {
			fmt.Println(stdout)
			fmt.Println(stderr)
		}

		if err != nil {
			return fmt.Errorf("failed to update when building Nutek Apple, error: %s", err)
		}
		if verbose {
			fmt.Println("removing old symlink")
		}
		if !isNotInstalled {
			runme = exec.Command("rm", os.Getenv("HOMEBREW_PREFIX")+"/bin/nutek-apple")
			output, err = runme.CombinedOutput()

			// Handle cases with or without a newline character
			stdout = string(output[:len(output)])
			stderr = "" // Initialize stderr as empty string

			// Split the output based on the first newline character (if present)
			if i := bytes.IndexByte(output, '\n'); i != -1 {
				stdout = string(output[:i])
				stderr = string(output[i:])
			} else {
				stdout = string(output) // If no newline, stdout is the entire output
			}

			if verbose {
				fmt.Println(stdout)
				fmt.Println(stderr)
			}

			if err != nil {
				fmt.Printf("failed to update when removing old symlink to Nutek Apple, error: %s\n", err)
			}
		}
		if verbose {
			fmt.Println("symlinking nutek-apple to $HOMEBREW_PREFIX/bin/nutek-apple")
		}
		// macOS thing to do
		runme = exec.Command("ln", "-s", rootPath+"/nutek-apple", os.Getenv("HOMEBREW_PREFIX")+"/bin/nutek-apple")
		output, err = runme.CombinedOutput()

		// Handle cases with or without a newline character
		stdout = string(output[:len(output)])
		stderr = "" // Initialize stderr as empty string

		// Split the output based on the first newline character (if present)
		if i := bytes.IndexByte(output, '\n'); i != -1 {
			stdout = string(output[:i])
			stderr = string(output[i:])
		} else {
			stdout = string(output) // If no newline, stdout is the entire output
		}

		if verbose {
			fmt.Println(stdout)
			fmt.Println(stderr)
		}

		if err != nil {
			return fmt.Errorf("failed to update when symlinking Nutek Apple, error: %s", err)
		}
		if verbose {
			fmt.Println("changing directory back to where Nutek Apple command was issued")
		}
		err = os.Chdir(wdPath)
		if err != nil {
			return fmt.Errorf("error: %s, when going back to directory when the Nutek Apple was invoked after update", err)
		}
	} else if upToDate {
		if verbose {
			fmt.Println("Everything up to date!")
		}
	} else if ahead {
		fmt.Println("hint: if you dant want to perform the update, pass 'noupdate' argument")
		if verbose {
			fmt.Printf("You're working on your own version of nutek-apple 🍎 and youre ahead.\nCommit and create a pull request when you're ready.\n") // Prompt the user for input
		}

	} else if rebase {
		if verbose {
			fmt.Println("You have work to do. Commit and update rebased code. To restart from a new 'git pull origin main', or 'git pull upstream main'")
		}
	} else {
		return fmt.Errorf("error: git command '%s' failed", pull.String())
	}

	return nil
}

func Update(argsSlice []string, verbose bool, runme func() error) error {
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
			if verbose {
				fmt.Println("saving current working directory")
			}
			wdPath, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("failed to update when getting working dir, error: %s", err)
			}
			if verbose {
				fmt.Println("going to root of Nutek Apple project")
			}
			err = os.Chdir(projectRoot)
			if err != nil {
				return fmt.Errorf("error: %s, when going to directory where the Nutek Apple is before checking remote url", err)
			}
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
			if strings.Contains(string(orginOutput), "https://github.com/NutekSecurity/nutek-apple.git") ||
				strings.Contains(string(orginOutput), "https://github.com/nuteksecurity/nutek-apple.git") {
				upstream := false
				err := gitUpdate(upstream, verbose)
				if err != nil {
					return err
				}
			} else if strings.Contains(string(upstreamOutput), "https://github.com/nuteksecurity/nutek-apple.git") ||
				strings.Contains(string(upstreamOutput), "https://github.com/NutekSecurity/nutek-apple.git") {
				upstream := true
				err := gitUpdate(upstream, verbose)
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("error: did not found repository link in git")
			}
			if verbose {
				fmt.Println("changing directory back to where Nutek Apple command was issued")
			}
			err = os.Chdir(wdPath)
			if err != nil {
				return fmt.Errorf("error: %s, when going back to directory when the Nutek Apple was invoked after checking remote url", err)
			}
			if err := runme(); err != nil {
				return fmt.Errorf("error: cli app returns %s", err)
			}
			return nil
		}
	}
	return fmt.Errorf("error: not in a git repository")
}

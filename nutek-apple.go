package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/nuteksecurity/nutek-apple/macos"
	"github.com/nuteksecurity/nutek-apple/util"
	"github.com/urfave/cli/v2"
)

// nutek-apple start
func main() {
	app := cli.NewApp()
	app.Name = "nutek-apple"
	app.Usage = "Programs YOU want to use and links YOU want click!\nFind out more on https://github.com/NutekSecurity/nutek-apple\n"
	app.Version = util.Version
	app.Authors = []*cli.Author{
		{
			Name:  "Szymon Bronisław Błaszczyński",
			Email: "neosb@nuteksecurity.com",
		},
	}
	app.Copyright = "MIT License (c) 2024 Szymon Bronisław Błaszczyński"
	app.Commands = []*cli.Command{
		{
			Name:        "programs",
			Aliases:     []string{"p"},
			Usage:       "access programs list(s)",
			Description: "Show a list(s) of programs",
			Args:        true,
			ArgsUsage:   "[author/listname] [many] [lists]",
			Action: func(c *cli.Context) error {
				err := cli.ShowAppHelp(c) //, "programs")
				if err != nil {
					return fmt.Errorf("error %s, when printing help for program command", err)
				}
				return fmt.Errorf("not yet implemented")
			},
			Subcommands: []*cli.Command{
				{
					Name:        "install",
					Aliases:     []string{"i"},
					Usage:       "use it to install programs",
					Description: "This is the usual way to install a single or multiple programs, just how you do it in Homebrew, dnf, apt, Chocolatey...",
					Args:        true,
					ArgsUsage:   "['author/listname'] [program] [program]",
					Action: func(c *cli.Context) error {
						return fmt.Errorf("not yet implemented")
						programs := []string{}
						if c.NArg() > 0 {
							programs = c.Args().Slice()
						}
						macos.Install(programs...)
						return nil
					},
				},
				{
					Name:    "all",
					Aliases: []string{"a"},
					Usage:   "list all authors and lists",
					Action: func(cCtx *cli.Context) error {
						return fmt.Errorf("not yet implemented")
						fmt.Println("All authors and their lists:")
						fmt.Println("use nutek-apple info [program] to know more")
						return nil
					},
				},
				{
					Name:        "info",
					Usage:       "use it to know more about program",
					Description: "Get some description about a program from Homebrew (macOS, Linux), or local package manager (Linux)",
					ArgsUsage:   "[program name to check for information]",
					Args:        true,
					Action: func(c *cli.Context) error {
						programs := []string{}
						if c.NArg() > 0 {
							programs = c.Args().Slice()
						}
						for _, program := range programs {
							macos.BrewCmd("info", program)
						}
						return nil
					},
				},
			},
		},
		{
			Name:        "links",
			Aliases:     []string{"bookmarks"},
			Usage:       "use it to show useful places around the internet - add author\\theme/list_name to displat the link(s)",
			Description: "Show or open links around the internet. When adding more arguments, don't forget to stay in shape 'theme/links_list_name'",
			Args:        true,
			Action: func(c *cli.Context) error {
				if c.Args().Len() > 0 {
					links := c.Args().Slice()
					for _, link := range links {
						result, err := util.ListLinksInFile("links", link)
						if err != nil {
							return err
						}
						fmt.Print(result)
					}
					return nil
				}
				result, err := util.ListLinksInFile("links", "nuteksecurity/example")
				if err != nil {
					return err
				}
				fmt.Print(result)
				return nil
			},
			Subcommands: []*cli.Command{
				{
					Name:        "list",
					Aliases:     []string{"all"},
					Usage:       "use it to show all useful places around internet from author\\theme",
					Description: "Get list of authors/themes to use in your next search",
					Args:        true,
					Action: func(c *cli.Context) error {
						rootDir := "links" // Specify the root directory you're interested in.
						util.WalkDir(rootDir)

						return nil
					},
				},
				{
					Name:        "search",
					Usage:       "use it to find desired term in Bookmarks",
					Description: "Search for desired term across all links and display hits",
					Args:        true,
					ArgsUsage:   "['search term you are looking for']",
					Action: func(c *cli.Context) error {
						if c.Args().Len() >= 2 {
							return fmt.Errorf("error: too much search terms, use quotes to combine into one")
						}
						searchTerm := c.Args().Get(0)
						result, err := util.SearchLinksInDirs("links", searchTerm)
						if err != nil {
							return fmt.Errorf("error: %s, when searching links for %s", err, searchTerm)
						}
						fmt.Print(result)
						return nil
					},
				},
				{
					Name:        "open",
					Usage:       "to open in default browser",
					Description: "Simply open the link in your browser window",
					Args:        true,
					ArgsUsage:   "[nuteksecurity/example/example.com]",
					Action: func(c *cli.Context) error {
						macos.OpenLink("https://nuteksecurity.com")
						return fmt.Errorf("not yet implemented")
					},
				},
			},
		},
		{
			Name:        "readme",
			Aliases:     []string{"doc"},
			Usage:       "to pretty print Markdown",
			Description: "Print out formatted text using mdcat from Markdown files in this project",
			Args:        true,
			ArgsUsage:   "[path to file]",
			Action: func(c *cli.Context) error {
				root, err := util.GetProjectRoot()
				if err != nil {
					return fmt.Errorf("error: %s, when determining project root", err)
				}
				if util.IsInstalled("mdcat") {
					if c.Args().Len() > 0 {
						args := c.Args().Slice()
						for _, arg := range args {
							catCmd := exec.Command("mdcat", root+"/"+arg)
							output, err := catCmd.Output()
							if err != nil {
								return fmt.Errorf("error: %s, when executing 'mdcat %s'", err, arg)
							}
							fmt.Print(string(output))
						}
					} else {
						catCmd := exec.Command("mdcat", root+"/README.md")
						output, err := catCmd.Output()
						if err != nil {
							return fmt.Errorf("error: %s, when executing 'mdcat README.md'", err)
						}
						fmt.Print(string(output))
					}
				} else {
					return fmt.Errorf("error: mdcat not installed")
				}
				return nil
			},
		},
	}

	var verbose bool = false
	for _, arg := range os.Args[1:] {
		if arg == "verbose" {
			verbose = true
			break
		}
	}

	for _, arg := range os.Args[1:] {
		if arg == "noupdate" {
			var newArgs []string
			for _, str := range os.Args {
				if str != "noupdate" {
					newArgs = append(newArgs, str)
				}
			}
			if err := app.Run(newArgs); err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}
	}

	// first update the nutek repository, then run the app
	if err := util.Update(os.Args, verbose, func() error {
		if err := app.Run(os.Args); err != nil {
			log.Fatal(err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

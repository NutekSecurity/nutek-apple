package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nuteksecurity/nutek-apple/macos"
	"github.com/nuteksecurity/nutek-apple/util"
	"github.com/urfave/cli/v2"
)

// nutek-apple start
func main() {
	app := cli.NewApp()
	app.Name = "nutek-apple"
	app.Usage = "Programs YOU want to use and links YOU want click!\nFind out more on https://github.com/NutekSecurity/nutek-apple"
	app.Version = "3.0.0"
	app.Commands = []*cli.Command{
		{
			Name:        "install",
			Aliases:     []string{"i"},
			Usage:       "use it to install programs",
			Description: "This is the usual way to install a single or multiple programs, just how you do it in Homebrew, dnf, apt, Chocolatey...",
			Args:        true,
			Action: func(c *cli.Context) error {
				programs := []string{}
				if c.NArg() > 0 {
					programs = c.Args().Slice()
				}
				for _, program := range programs {
					// check here if system is macOS, Debian, Ubuntu, Fedora, RedHat, Arch, Alpine...
					macos.BrewCmd("info", program)
				}
				return nil
			},
			// change this to pull a list in form of nuteksecurity/list_name
			Subcommands: []*cli.Command{},
		},
		{
			Name:    "list",
			Aliases: []string{"all"},
			Usage:   "list all authors\\themes and lists inside",
			Args:    true,
			Action: func(cCtx *cli.Context) error {
				fmt.Println("All programs:")
				for _, program := range macos.Cli() {
					fmt.Println(program)
				}
				for _, program := range macos.Gui() {
					fmt.Println(program)
				}
				fmt.Println("use nutek-apple info [program] to know more")
				return nil
			},
			Subcommands: []*cli.Command{
				{
					Name:    "all",
					Aliases: []string{"a"},
					Usage:   "list all authors and lists inside",
					Action: func(cCtx *cli.Context) error {
						fmt.Println("All authors and their lists:")
						fmt.Println("use nutek-apple info [program] to know more")
						return nil
					},
				},
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
						fmt.Printf(result)
					}
					return nil
				}
				result, err := util.ListLinksInFile("links", "nuteksecurity/example")
				if err != nil {
					return err
				}
				fmt.Printf(result)
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
					Action: func(c *cli.Context) error {
						searchTerm := c.Args().Get(0)
						result, err := util.SearchLinksInDirs("links", searchTerm)
						if err != nil {
							return err
						}
						fmt.Printf(result)
						return nil
					},
				},
			},
		},
	}

	// first update the nutek repository, then run the app
	util.Update(os.Args, func() error {
		if err := app.Run(os.Args); err != nil {
			log.Fatal(err)
		}
		return nil
	})
}

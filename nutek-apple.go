package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nuteksecurity/nutek-apple/links"
	"github.com/nuteksecurity/nutek-apple/macos"
	"github.com/urfave/cli/v2"
)

// nutek-apple start
func main() {
	app := cli.NewApp()
	app.Name = "nutek-apple"
	app.Usage = "Nutek Security Platform\nfind more on https://nuteksecurity.com"
	app.Version = "3.0.0"
	app.Commands = []*cli.Command{
		{
			Name:        "install",
			Aliases:     []string{"i"},
			Usage:       "use it to install programs",
			Description: "This is how we describe describeit the function",
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
			Subcommands: []*cli.Command{
				{
					Name:    "all",
					Aliases: []string{"a"},
					Usage:   "install all programs",
					Action: func(cCtx *cli.Context) error {
						// install all programs
						return nil
					},
				},
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "list all programs",
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
							Name:  "all",
							Usage: "show all programs",
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
						},
						{
							Name:  "cli",
							Usage: "show all CLI programs",
							Action: func(cCtx *cli.Context) error {
								// show all CLI programs
								fmt.Println("CLI programs:")
								for _, program := range macos.Cli() {
									fmt.Println(program)
								}
								fmt.Println("use nutek-apple info [program] to know more")
								return nil
							},
						},
						{
							Name:  "gui",
							Usage: "show all GUI programs",
							Action: func(cCtx *cli.Context) error {
								// show all CLI programs
								fmt.Println("GUI programs:")
								for _, program := range macos.Gui() {
									fmt.Println(program)
								}
								fmt.Println("use nutek-apple info [program] to know more")
								return nil
							},
						},
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
			Usage:       "use it to show useful places around internet",
			Description: "Show or open links around the internet",
			Args:        true,
			Action: func(c *cli.Context) error {
				for _, link := range links.Links() {
					fmt.Println(link)
				}
				return nil
			},
			Subcommands: []*cli.Command{
				{
					Name:        "all",
					Usage:       "use it to show all useful places around internet",
					Description: "Get some description about a program from Homebrew (macOS, Linux), or local package manager (Linux)",
					Args:        true,
					Action: func(c *cli.Context) error {
						for _, link := range links.Links() {
							fmt.Println(link)
						}
						return nil
					},
				},
				{
					Name:        "seclists",
					Usage:       "use it to open SecLists in browser",
					Description: "SecLists is the security tester's companion. It's a collection of multiple types of lists used during security assessments, collected in one place. List types include usernames, passwords, URLs, sensitive data patterns, fuzzing payloads, web shells, and many more.",
					Args:        true,
					Action: func(c *cli.Context) error {
						macos.OpenLink(links.SecLists())
						return nil
					},
				},
				{
					Name:        "payloadsallthethings",
					Usage:       "use it to open PayloadAllTheThings in browser",
					Description: "A list of useful payloads and bypass for Web Application Security and Pentest/CTF",
					Args:        true,
					Action: func(c *cli.Context) error {
						macos.OpenLink(links.PayloadsAllTheThings())
						return nil
					},
				},
				{
					Name:        "hacktricks",
					Usage:       "use it to open HackTricks in browser",
					Description: "Give FREE access to EDUCATIONAL hacking resources to ALL Internet.",
					Args:        true,
					Action: func(c *cli.Context) error {
						macos.OpenLink(links.HackTricks())
						return nil
					},
				},
				{
					Name:        "exploitdb",
					Aliases:     []string{"exploit-db"},
					Usage:       "use it to open Exploit-DB in browser",
					Description: "Exploit Database",
					Args:        true,
					Action: func(c *cli.Context) error {
						macos.OpenLink(links.ExploitDb())
						return nil
					},
				},
				{
					Name:        "peass",
					Aliases:     []string{"peass-ng"},
					Usage:       "use it to open PEASS-ng in browser",
					Description: "Exploit Database",
					Args:        true,
					Action: func(c *cli.Context) error {
						macos.OpenLink(links.PEASS())
						return nil
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

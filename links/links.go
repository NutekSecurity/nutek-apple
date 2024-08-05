package links

// list of all links over the internet
func Links() []string {
	links := []string{
		"https://github.com/danielmiessler/SecLists",
		"https://github.com/swisskyrepo/PayloadsAllTheThings",
		"https://book.hacktricks.xyz",
		"https://www.exploit-db.com/",
		"https://github.com/peass-ng/PEASS-ng",
	}
	return links
}

// link to SecLists on GitHub
func SecLists() string {
	return "https://github.com/danielmiessler/SecLists"
}

// link to PayloadsAllTheThings on GitHub
func PayloadsAllTheThings() string {
	return "https://github.com/swisskyrepo/PayloadsAllTheThings"
}

// Link to Hack Tricks book
func HackTricks() string {
	return "https://book.hacktricks.xyz"
}

// link to Exploit-DB
func ExploitDb() string {
	return "https://www.exploit-db.com/"
}

// link to PEASS-ng on GitHub
func PEASS() string {
	return "https://github.com/peass-ng/PEASS-ng"
}

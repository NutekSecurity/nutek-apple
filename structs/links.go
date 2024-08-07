package structs

type Link struct {
	ShortName    string
	LongName     string
	URL          string
	Description  string
	Tags         []Tag
	Languages    []Language
	Technologies []Technology
	Author       []Author
	AddedBy      Author
}

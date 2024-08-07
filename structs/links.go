package structs

type Link struct {
	ShortName    string
	LongName     string
	URL          string
	Description  string
	Tags         []string
	Languages    []string
	Technologies []Technology
	Author       []Author
	AddedBy      Author
}

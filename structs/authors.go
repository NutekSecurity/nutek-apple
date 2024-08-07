package structs

type Author struct {
	ShortName    string
	LongName     string
	Description  string
	Contact      []string
	Origins      []string
	URL          []Link
	Technologies []Technology
	Languages    []Language
	Tags         []Tag
}

package structs

type Author struct {
	ShortName    string
	LongName     string
	Description  string
	Contact      []string
	Origins      []string
	Website      string
	Technologies []Technology
	Languages    []Language
}

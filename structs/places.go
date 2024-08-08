package structs

type Place struct {
	ShortName     string
	LongName      string
	Location      string
	Description   string
	Languages     []Language
	Contact       []string
	MaxPopulation int
	AddedBy       Author
	Tags          []Tag
	URL           string
	Technologies  []Technology
}

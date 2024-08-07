package structs

type License struct {
	ShortName   string
	LongName    string
	Description string
	FullText    string
	Tags        []Tag
	AddedBy     Author
	URL         []Link
}

package structs

type Language struct {
	ShortName string
	LongName  string
	Places    []Place
	Tags      []Tag
	AddedBy   Author
	URL       []Link
}

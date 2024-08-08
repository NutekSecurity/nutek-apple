package structs

type Author struct {
	ShortName            string
	LongName             string
	Description          string
	Contact              []string
	Origins              []string
	URL                  []Link
	Technologies         []Technology
	Programs             []Program
	ProgrammingLanguages []ProgrammingLanguage
	Languages            []Language
	Tags                 []Tag
}

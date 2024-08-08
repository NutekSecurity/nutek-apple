package structs

type ProgrammingLanguage struct {
	ShortName         string
	LongName          string
	Description       string
	Tags              []Tag
	AddedBy           Author
	URL               []Link
	PackageManagers   []PackageManager
	OtherWayToInstall []string
	Manual            []Link
}

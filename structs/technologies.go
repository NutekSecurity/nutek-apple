package structs

type Technology struct {
	ShortName       string
	LongName        string
	Description     string
	PackageManagers []PackageManager
	OS              []OS
	Tags            []Tag
	AddedBy         Author
	URL             []Link
	Manual          []Link
}

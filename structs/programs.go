package structs

type Program struct {
	ShortName       string
	LongName        string
	Description     string
	Tags            []string
	URL             []Link
	Manual          []Link
	PackageManagers []PackageManager
	License         License
	Languages       []string
	Technologies    []Technology
	Source          Link
	OS              []OS
	AddedBy         Author
}

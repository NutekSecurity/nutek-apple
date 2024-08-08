package structs

type OS struct {
	ShortName       string
	LongName        string
	Description     string
	PackageManagers []PackageManager
	Author          []Author
	AddedBy         Author
}

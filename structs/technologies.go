package structs

type Technology struct {
	ShortName       string
	LongName        string
	Description     string
	PackageManagers []PackageManager
	OS              []OS
}

type OS struct {
	ShortName       string
	LongName        string
	Description     string
	PackageManagers []PackageManager
	Author          []Author
}

package structs

type PrettyPrint interface {
	PrettyPrintOneList(list string)
	PrettyPrintAllListsInOneFolder(folder string)
}

type Install interface {
	Install(packageManager PackageManager, program string)
}

type Search interface {
	Search(searchTerm string)
	SearchByShortName(searchTerm string)
}

type Enumerate interface {
	ListFolders(folders ...string)
	ListOneFolder(folder string)
}

type DoesExist interface {
	DoExistByListName()
	DoExistByShortName()
	DoExistByFullListName()
	DoFolderExist()
}

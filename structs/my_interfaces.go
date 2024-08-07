package structs

type PrettyPrint interface {
	PrettyPrint()
}

type Install interface {
	Install()
}

type Search interface {
	Search()
	SearchName()
}

type List interface {
	ListAll()
	ListFolders()
	ListOneFolder()
}

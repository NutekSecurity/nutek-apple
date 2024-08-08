package structs

type AuthorList struct {
	ParentFolder   string
	FullAccessPath string
	Version        string
	Author         Author
	Elements       []Author
}

type LanguageList struct {
	ParentFolder   string
	FullAccessPath string
	Version        string
	Author         Author
	Elements       []Language
}

type LicenseList struct {
	ParentFolder   string
	FullAccessPath string
	Version        string
	Author         Author
	Elements       []License
}

type LinkList struct {
	ParentFolder   string
	FullAccessPath string
	Version        string
	Author         Author
	Elements       []Link
}

type OSList struct {
	ParentFolder   string
	FullAccessPath string
	Version        string
	Author         Author
	Elements       []OS
}

type PackageManagerList struct {
	ParentFolder   string
	FullAccessPath string
	Version        string
	Author         Author
	Elements       []PackageManager
}

type PlaceList struct {
	ParentFolder   string
	FullAccessPath string
	Version        string
	Author         Author
	Elements       []Place
}

type ProgrammingLanguageList struct {
	ParentFolder   string
	FullAccessPath string
	Version        string
	Author         Author
	Elements       []ProgrammingLanguage
}

type ProgramList struct {
	ParentFolder   string
	FullAccessPath string
	Version        string
	Author         Author
	Elements       []Program
}

type TagList struct {
	ParentFolder   string
	FullAccessPath string
	Version        string
	Author         Author
	Elements       []Tag
}

type TechnologyList struct {
	ParentFolder   string
	FullAccessPath string
	Version        string
	Author         Author
	Elements       []Technology
}

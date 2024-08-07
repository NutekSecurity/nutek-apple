package structs

type Program struct {
	ShortName       string
	LongName        string
	Description     string
	Tags            []string
	URL             string
	Manual          string
	PackageManagers []string
	License         string
	Languages       []string
	Technology      []Technology
}

type Programs struct {
	List    []Program
	Author  string
	Version string
}

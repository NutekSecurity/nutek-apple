package structs

type PackageManager struct {
	ShortName                 string
	LongName                  string
	Description               string
	OS                        []OS
	Tags                      []Tag
	AddedBy                   Author
	License                   License
	URL                       []Link
	Program                   Program
	Manual                    []Link
	InstallProgramInstruction string
	UpdateProgramInstruction  string
}

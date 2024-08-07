package structs

type Link struct {
	URL         string
	Description string
}

type Links struct {
	Links map[string]Link
}

package yaag

type Config struct {
	On bool

	BaseUrls map[string]*UrlValue

	DocTitle string
	DocPath  string
}

type UrlValue struct {
	Val string
}
package lookup

type LookupService interface {
	Lookup(urls []string) string
}

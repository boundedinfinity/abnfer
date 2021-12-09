package rule

type Element interface {
	HasPrefix(string) bool
}

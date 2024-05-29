package store

type Pagination[T any] interface {
	PageUrl(page int) string
	Pages() []int
	PageCount() int
	CurrentPage() int
	PerPage() int
	Items() []T
	Total() int
	NextPageUrl() string
	PreviousPageUrl() string
}

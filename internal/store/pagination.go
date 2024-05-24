package store

import "fmt"

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

type UserPagination struct {
	baseUrl      string
	curPage      int
	itemsPerPage int
	store        UserStore
}

func NewUserPagination(url string, store UserStore, pg int) UserPagination {
	return UserPagination{baseUrl: url,
		curPage:      pg,
		itemsPerPage: 5,
		store:        store,
	}
}

func (thing UserPagination) PageUrl(page int) string {
	return fmt.Sprintf("%s?page=%d", thing.baseUrl, page)
}

func (thing UserPagination) Pages() []int {
	var ret []int
	for i := 0; i < thing.PageCount(); i++ {
		ret = append(ret, i+1)
	}
	return ret
}

func (thing UserPagination) Items() []User {
	var ret []User
	k := (thing.curPage - 1) * thing.itemsPerPage
	for i, u := range thing.store.ListUsers() {
		if i >= k && i < k+thing.itemsPerPage {
			ret = append(ret, u)
		}
	}
	return ret
}

func (thing UserPagination) PageCount() int {
	k := thing.Total() / thing.itemsPerPage
	if k%thing.itemsPerPage != 0 {
		k += 1
	}
	return k
}

func (thing UserPagination) CurrentPage() int {
	return thing.curPage
}

func (thing UserPagination) PerPage() int {
	return thing.itemsPerPage
}

func (thing UserPagination) PreviousPageUrl() string {
	return fmt.Sprintf("%s?page=%d", thing.baseUrl, thing.curPage-1)
}

func (thing UserPagination) Total() int {
	return len(thing.store.ListUsers())
}

func (thing UserPagination) NextPageUrl() string {
	return fmt.Sprintf("%s?page=%d", thing.baseUrl, thing.curPage+1)
}

package store

import (
	"fmt"
)

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

type AbstractPagination[T any] struct {
	Pagination[T]
	BaseUrl      string
	CurPage      int
	ItemsPerPage int
	Child        Pagination[T]
}

func (thing AbstractPagination[T]) PageUrl(page int) string {
	return fmt.Sprintf("%s?page=%d", thing.BaseUrl, page)
}

func (thing AbstractPagination[T]) Pages() []int {
	var ret []int
	for i := 0; i < thing.PageCount(); i++ {
		ret = append(ret, i+1)
	}
	return ret
}

func (thing AbstractPagination[T]) PageCount() int {

	k := thing.Child.Total() / thing.ItemsPerPage
	if k%thing.ItemsPerPage != 0 {
		k += 1
	}
	return k
}

func (thing AbstractPagination[T]) CurrentPage() int {
	return thing.CurPage
}

func (thing AbstractPagination[T]) PerPage() int {
	return thing.ItemsPerPage
}

func (thing AbstractPagination[T]) PreviousPageUrl() string {
	return fmt.Sprintf("%s?page=%d", thing.BaseUrl, thing.CurPage-1)
}

// func (thing AbstractPagination) Total() int {
// 	return int(thing.store.GetUserCount())
// }

func (thing AbstractPagination[T]) NextPageUrl() string {
	return fmt.Sprintf("%s?page=%d", thing.BaseUrl, thing.CurPage+1)
}

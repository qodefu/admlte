package dbstore

import (
	"goth/internal/store"
	"time"
)

type ApptStore struct {
	appts []store.Appt // Slice of User structs to store user data.
}

func (thing ApptStore) ListAppts() []store.Appt {
	return thing.appts
}

// NewUserStore initializes and returns a new instance of UserStore.
// It pre-populates the store with a default user for demonstration or testing purposes.
func NewApptStore() *ApptStore {
	return &ApptStore{
		appts: []store.Appt{
			{
				Id:         1,
				ClientId:   1,
				ApptTime:   time.Now(),
				Status:     "good",
				Note:       "Nothing",
				TimeStamps: time.Now(),
			},
		},
	}
}

type ApptPagination struct {
	store store.ApptStore
}

func (thing ApptPagination) PageUrl(page int) string {
	return "#"
}

func (thing ApptPagination) Pages() []int {
	return []int{1}
}

func (thing ApptPagination) PageCount() int {
	return 1
}

func (thing ApptPagination) CurrentPage() int {
	return 1
}

func (thing ApptPagination) PerPage() int {
	return 5
}

func (thing ApptPagination) Items() []store.Appt {
	return thing.store.ListAppts()
}

func (thing ApptPagination) Total() int {
	return 1
}

func (thing ApptPagination) NextPageUrl() string {
	return "#"
}

func (thing ApptPagination) PreviousPageUrl() string {
	return "#"
}

func NewApptPagination(store store.ApptStore) store.Pagination[store.Appt] {
	return ApptPagination{
		store: store,
	}
}

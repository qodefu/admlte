package mockstore

import (
	"goth/internal/store"
	"goth/internal/store/models"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type ApptStore struct {
	appts []models.Appointment // Slice of User structs to store user data.
}

func (thing ApptStore) ListAppts() []models.Appointment {
	return thing.appts
}

// NewUserStore initializes and returns a new instance of UserStore.
// It pre-populates the store with a default user for demonstration or testing purposes.
func NewApptStore() *ApptStore {
	return &ApptStore{
		appts: []models.Appointment{
			{
				ID:       1,
				ClientID: pgtype.Int4{Int32: 1},
				ApptTime: pgtype.Timestamp{Time: time.Now()},
				Status:   pgtype.Text{String: "good"},
				Note:     pgtype.Text{String: "Nothing"},
				Created:  pgtype.Timestamp{Time: time.Now()},
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

func (thing ApptPagination) Items() []models.ListApptRow {
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

func NewApptPagination(store store.ApptStore) store.Pagination[models.ListApptRow] {
	return ApptPagination{
		store: store,
	}
}

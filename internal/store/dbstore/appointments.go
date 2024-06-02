package dbstore

import (
	"context"
	"goth/internal/store"
	"goth/internal/store/models"
)

type ApptStore struct {
	appts *models.Queries // Slice of User structs to store user data.
}

func (thing ApptStore) ListAppts() []models.Appointment {
	ret, _ := thing.appts.ListAppt(context.Background())
	return ret
}

// NewUserStore initializes and returns a new instance of UserStore.
// It pre-populates the store with a default user for demonstration or testing purposes.
func NewApptStore(queries *models.Queries) *ApptStore {
	return &ApptStore{
		appts: queries,
	}
}

type ApptPagination struct {
	store.AbstractPagination[models.Appointment]
	q *models.Queries
}

func NewApptPagination(url string, queries *models.Queries, pg int) store.Pagination[models.Appointment] {
	super := store.MkAbsPgtor[models.Appointment](5, pg, url)
	ret := ApptPagination{
		super,
		queries,
	}
	ret.Child = ret
	return ret
}

func (thing ApptPagination) Items() []models.Appointment {
	ret, _ := thing.q.ListAppt(context.Background())
	return ret
}

func (thing ApptPagination) Total() int {
	ret, _ := thing.q.GetAppointmentCount(context.Background())
	return int(ret)
}

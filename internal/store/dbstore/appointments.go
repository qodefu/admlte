package dbstore

import (
	"context"
	"goth/internal/store"
	"goth/internal/store/models"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type ApptStore struct {
	appts *models.Queries // Slice of User structs to store user data.
}

func (thing ApptStore) ListAppts() []models.ListApptRow {
	ret, _ := thing.appts.ListAppt(context.Background())
	return ret
}

func (thing ApptStore) CreateAppt(id int64, apptTime time.Time, status, note string) (models.Appointment, error) {
	return thing.appts.CreateAppt(context.Background(), models.CreateApptParams{
		ClientID: pgtype.Int8{Int64: id, Valid: true},
		ApptTime: pgtype.Timestamp{Time: apptTime, Valid: true},
		Status:   pgtype.Text{String: status, Valid: true},
		Note:     pgtype.Text{String: note, Valid: true},
	})
}

func (thing ApptStore) GetApptById(id int64) (models.GetAppointmentRow, error) {
	return thing.appts.GetAppointment(context.Background(), id)
}

// NewUserStore initializes and returns a new instance of UserStore.
// It pre-populates the store with a default user for demonstration or testing purposes.
func NewApptStore(queries *models.Queries) *ApptStore {
	return &ApptStore{
		appts: queries,
	}
}

type ApptPagination struct {
	store.AbstractPagination[models.ListApptRow]
	q *models.Queries
}

func NewApptPagination(url string, queries *models.Queries, pg int) store.Pagination[models.ListApptRow] {
	super := store.MkAbsPgtor[models.ListApptRow](5, pg, url)
	ret := ApptPagination{
		super,
		queries,
	}
	ret.Child = ret
	return ret
}

func (thing ApptPagination) Items() []models.ListApptRow {
	ret, _ := thing.q.ListAppt(context.Background())
	return ret
}

func (thing ApptPagination) Total() int {
	ret, _ := thing.q.GetAppointmentCount(context.Background())
	return int(ret)
}

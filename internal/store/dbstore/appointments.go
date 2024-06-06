package dbstore

import (
	"context"
	"goth/internal/config"
	"goth/internal/store"
	"goth/internal/store/models"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type ApptStore struct {
	appts *models.Queries // Slice of User structs to store user data.
}

func (thing ApptStore) ListAppts(offset, limit int) []models.ListApptRow {
	ret, _ := thing.appts.ListAppt(context.Background(), models.ListApptParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
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

func (thing ApptStore) UpdateAppt(id int64, clientId int64, apptTime time.Time, status string, note string) error {
	return thing.appts.UpdateAppt(context.Background(), models.UpdateApptParams{
		ID:       id,
		ClientID: pgtype.Int8{Int64: clientId, Valid: true},
		ApptTime: pgtype.Timestamp{Time: apptTime, Valid: true},
		Status:   pgtype.Text{String: status, Valid: true},
		Note:     pgtype.Text{String: note, Valid: true},
	})
}

func (thing ApptStore) GetApptById(id int64) (models.GetAppointmentRow, error) {
	return thing.appts.GetAppointment(context.Background(), id)
}

func (thing ApptStore) DeleteAppt(id int64) error {
	return thing.appts.DeleteAppt(context.Background(), id)
}

func (thing ApptStore) GetApptCount() (int64, error) {
	return thing.appts.GetAppointmentCount(context.Background())
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
	queries store.ApptStore
}

func NewApptPagination(queries store.ApptStore, pg int) store.Pagination[models.ListApptRow] {
	super := store.MkAbsPgtor[models.ListApptRow](5, pg, config.Routes().Admin.Appt.HX.List)
	ret := ApptPagination{
		super,
		queries,
	}
	ret.Child = ret
	return ret
}

func (thing ApptPagination) Items() []models.ListApptRow {
	ret := thing.queries.ListAppts(thing.Offset(), thing.ItemsPerPage)
	return ret
}

func (thing ApptPagination) Total() int {
	ret, _ := thing.queries.GetApptCount()
	return int(ret)
}

package appts

import (
	"encoding/json"
	"fmt"
	"goth/internal/store"
	"goth/internal/store/models"
	"goth/internal/templates"
	"goth/internal/utils"
	"goth/internal/validator"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func apptBadge(apptRow models.ListApptRow) string {
	if apptRow.Status.String == "SCHEDULED" {
		return "badge-primary"
	}
	return "badge-success"
}

type ApptFormValidation struct {
	ClientId validator.FormInput
	Date     validator.FormInput
	Time     validator.FormInput
	Note     validator.FormInput
	Status   validator.FormInput
}

func newValidation(cliIdStr, apptDate, apptTime, note, status string) ApptFormValidation {
	return ApptFormValidation{
		ClientId: validator.New("clientId", cliIdStr, &idGen, validator.NotEmpty("Client")),
		Date:     validator.New("appt_date", apptDate, &idGen, validator.NotEmpty("Appointment Date")),
		Time:     validator.New("appt_time", apptTime, &idGen, validator.NotEmpty("Appointment Time")),
		Note:     validator.New("note", note, &idGen),
		Status:   validator.New("status", status, &idGen, validator.NotEmpty("Status")),
	}
}

type ApptHandler struct {
	apptRepo store.ApptStore
	cliRepo  store.ClientStore
}

func NewApptsHandler(apptRepo store.ApptStore, cliRepo store.ClientStore) *ApptHandler {
	return &ApptHandler{apptRepo: apptRepo, cliRepo: cliRepo}
}

func (thing ApptHandler) CreateForm(w http.ResponseWriter, r *http.Request) error {
	clients := thing.cliRepo.ListClients()
	av := newValidation("", "", "", "", "")
	templates.Layout(ApptForm(clients, av), "Appt Form").Render(r.Context(), w)
	return nil
}

func (thing ApptHandler) SaveNew(w http.ResponseWriter, r *http.Request) error {

	cliIdStr := r.FormValue("clientId")
	var cliId int64
	_, err := fmt.Sscan(cliIdStr, &cliId)
	if err != nil {
		return err
	}
	apptDate := r.FormValue("appt_date")
	apptTime := r.FormValue("appt_time")
	apptStatus := r.FormValue("appt_status")
	note := r.FormValue("note")

	av := newValidation(cliIdStr, apptDate, apptTime, note, apptStatus)
	combineTime := apptDate + " " + apptTime
	validator.ValidateFields(&av)

	if !validator.ValidationOk(&av) {
		clients := thing.cliRepo.ListClients()
		ApptForm(clients, av).Render(r.Context(), w)
		// return errors.New("validation failed")
	} else {
		t, _ := time.Parse("01/02/2006 3:04 PM", combineTime)
		note := r.FormValue("appt_note")
		thing.apptRepo.CreateAppt(cliId, t, apptStatus, note)
		w.Header().Set("HX-Trigger", trigger(DeclMsg{
			Event:   "appointment-updated",
			Message: "appointment saved",
			Tags:    "success!",
		},
			DeclMsg{
				Event:   "appointment-messed",
				Message: "appointment messed",
				Tags:    "error!",
			}))
		// templates.Layout(ApptForm(appts), "Appt Form").Render(r.Context(), w)

	}
	return nil
}

type HXTrigger struct {
	eventName string
}

type DeclMsg struct {
	Event   string `json:"-"`
	Message string `json:"message"`
	Tags    string `json:"tags"`
	Data    any    `json:"data,omitempty"`
}

func trigger(events ...DeclMsg) string {
	var data = make(map[string][]DeclMsg)
	for _, e := range events {
		data[e.Event] = append([]DeclMsg{}, e)
	}
	bytes, _ := json.Marshal(data)
	return string(bytes)
}

func (thing ApptHandler) UpdateAppt(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	var apptId int64
	_, err := fmt.Sscan(idStr, &apptId)
	if err != nil {
		return err
	}
	clients := thing.cliRepo.ListClients()
	row, err := thing.apptRepo.GetApptById(apptId)
	if err != nil {
		return err
	}
	cliId := fmt.Sprint(row.ClientID.Int64)
	dateStr := utils.DateFormat(row.ApptTime.Time)
	timeStr := utils.TimeFormat(row.ApptTime.Time)
	av := newValidation(cliId, dateStr, timeStr, row.Note.String, row.Status.String)
	templates.Layout(ApptForm(clients, av), "Edit Form").Render(r.Context(), w)
	return nil
}

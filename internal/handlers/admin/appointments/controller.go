package appts

import (
	"goth/internal/store"
	"goth/internal/templates"
	"goth/internal/validator"
	"net/http"
	"strconv"
	"time"
)

type ApptFormValidation struct {
	ClientId validator.Validation
	Date     validator.Validation
	Time     validator.Validation
	Note     validator.Validation
	Status   validator.Validation
}

func newValidation(cliIdStr, apptDate, apptTime, note, status string) ApptFormValidation {
	return ApptFormValidation{
		ClientId: validator.New("clientId", cliIdStr, validator.NotEmpty("Client")),
		Date:     validator.New("appt_date", apptDate, validator.NotEmpty("Appointment Date")),
		Time:     validator.New("appt_time", apptTime, validator.NotEmpty("Appointment Time")),
		Note:     validator.New("note", note),
		Status:   validator.New("status", status, validator.NotEmpty("Status")),
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
	cliId, _ := strconv.Atoi(cliIdStr)
	apptDate := r.FormValue("appt_date")
	apptTime := r.FormValue("appt_time")
	apptStatus := r.FormValue("appt_status")
	note := r.FormValue("note")

	av := newValidation(cliIdStr, apptDate, apptTime, apptStatus, note)
	combineTime := apptDate + " " + apptTime
	validator.ValidateFields(&av)

	if !validator.ValidationOk(&av) {
		clients := thing.cliRepo.ListClients()
		templates.Layout(ApptForm(clients, av), "Appt Form").Render(r.Context(), w)
		// return errors.New("validation failed")
	} else {
		t, _ := time.Parse("01/02/2006 3:04 PM", combineTime)
		note := r.FormValue("appt_note")
		thing.apptRepo.CreateAppt(int32(cliId), t, apptStatus, note)
		// templates.Layout(ApptForm(appts), "Appt Form").Render(r.Context(), w)
		return nil

	}
}

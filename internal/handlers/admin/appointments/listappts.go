package appts

import (
	"fmt"
	"goth/internal/store"
	"goth/internal/templates"
	"net/http"
	"strconv"
	"time"
)

func init() {

}

type ApptHandler struct {
	apptRepo store.ApptStore
	cliRepo  store.ClientStore
}

func NewApptsHandler(apptRepo store.ApptStore, cliRepo store.ClientStore) *ApptHandler {
	return &ApptHandler{apptRepo: apptRepo, cliRepo: cliRepo}
}

func (thing ApptHandler) CreateForm(w http.ResponseWriter, r *http.Request) error {
	appts := thing.cliRepo.ListClients()
	templates.Layout(ApptForm(appts), "Appt Form").Render(r.Context(), w)
	return nil
}

func (thing ApptHandler) SaveNew(w http.ResponseWriter, r *http.Request) error {

	cliIdStr := r.FormValue("clientId")
	cliId, err := strconv.Atoi(cliIdStr)
	if err != nil {
		return err
	}
	apptDate := r.FormValue("appt_date")
	apptTime := r.FormValue("appt_time")
	combineTime := apptDate + " " + apptTime
	fmt.Println(combineTime)
	t, _ := time.Parse("01/02/2006 3:04 PM", combineTime)
	note := r.FormValue("appt_note")
	thing.apptRepo.CreateAppt(int32(cliId), t, "open", note)
	// templates.Layout(ApptForm(appts), "Appt Form").Render(r.Context(), w)
	return nil
}

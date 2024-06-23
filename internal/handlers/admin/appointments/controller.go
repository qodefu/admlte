package appts

import (
	"fmt"
	"goth/internal/components"
	"goth/internal/config"
	"goth/internal/middleware"
	"goth/internal/store"
	"goth/internal/store/dbstore"
	"goth/internal/store/models"
	"goth/internal/templates"
	"goth/internal/utils"
	"goth/internal/validator"
	"net/http"
	"strconv"
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
	Id       validator.FormInput
	ClientId validator.FormInput
	Date     validator.FormInput
	Time     validator.FormInput
	Note     validator.FormInput
	Status   validator.FormInput
}

func newEmptyForm() ApptFormValidation {
	return newValidation("", "", "", "", "", "")
}

func newValidation(idStr, cliIdStr, apptDate, apptTime, note, status string) ApptFormValidation {
	return ApptFormValidation{
		Id:       validator.New("id", idStr, &idGen),
		ClientId: validator.New("clientId", cliIdStr, &idGen, validator.NotEmpty("Client")),
		Date:     validator.New("appt_date", apptDate, &idGen, validator.NotEmpty("Appointment Date")),
		Time:     validator.New("appt_time", apptTime, &idGen, validator.NotEmpty("Appointment Time")),
		Note:     validator.New("note", note, &idGen),
		Status:   validator.New("appt_status", status, &idGen, validator.NotEmpty("Status")),
	}
}

type ComponentEnum int

const (
	COMP_LIST_APPT = iota
)

type apptHandler struct {
	apptRepo store.ApptStore
	cliRepo  store.ClientStore
	registry map[ComponentEnum]components.CreateComp
}

var handler *apptHandler

func NewApptsHandler(apptRepo store.ApptStore, cliRepo store.ClientStore) *apptHandler {
	if handler == nil {
		handler = &apptHandler{
			apptRepo: apptRepo,
			cliRepo:  cliRepo,
			registry: make(map[ComponentEnum]components.CreateComp)}
	}
	return handler
}

func (thing apptHandler) Component(enum ComponentEnum) components.CreateComp {
	return thing.registry[enum]
}

func (thing apptHandler) ListApptComp() (string, components.CreateComp) {
	compId := idGen.Id("list_appt_comp")
	ret := func(req middleware.RequestScope) components.RComp {
		return ListApptComp{
			compId:   compId,
			Page:     1,
			req:      req,
			apptRepo: thing.apptRepo,
		}
	}
	thing.registry[COMP_LIST_APPT] = ret
	return compId, ret
}

func (thing apptHandler) CreateForm(w http.ResponseWriter, r *http.Request) error {
	clients := thing.cliRepo.ListClients()
	av := newEmptyForm()
	templates.Layout(ApptForm(clients, av), "Appt Form").Render(r.Context(), w)
	return nil
}

func (thing apptHandler) SaveNew(req middleware.RequestScope) error {

	av := newEmptyForm()
	cliIdStr := req.Request().FormValue(av.ClientId.Key)
	var cliId int64
	_, err := fmt.Sscan(cliIdStr, &cliId)
	if err != nil {
		return err
	}
	apptDate := req.Request().FormValue(av.Date.Key)
	apptTime := req.Request().FormValue(av.Time.Key)
	apptStatus := req.Request().FormValue(av.Status.Key)
	note := req.Request().FormValue(av.Note.Key)

	combineTime := apptDate + " " + apptTime

	av.ClientId.Value = cliIdStr
	av.Date.Value = apptDate
	av.Time.Value = apptTime
	av.Note.Value = note
	av.Status.Value = apptStatus
	validator.ValidateFields(&av)

	if !validator.ValidationOk(&av) {
		clients := thing.cliRepo.ListClients()
		ApptForm(clients, av).Render(req.Context(), req.Response())
		// return errors.New("validation failed")
	} else {
		t, _ := time.Parse("01/02/2006 3:04 PM", combineTime)
		note := req.Request().FormValue("appt_note")

		apptIdStr := req.Request().FormValue("id")
		if len(apptIdStr) > 0 {
			//save existing
			var apptId int64
			_, err := fmt.Sscan(apptIdStr, &apptId)
			if err != nil {
				return err
			}
			thing.apptRepo.UpdateAppt(apptId, cliId, t, apptStatus, note)
			req.HxTrigger(middleware.TriggerMsg{
				Event:   "appointment-updated",
				Message: "appointment changes saved",
				Tags:    "success!",
			})
		} else {
			// create new appointment
			thing.apptRepo.CreateAppt(cliId, t, apptStatus, note)
			req.HxTrigger(middleware.TriggerMsg{
				Event:   "appointment-created",
				Message: "appointment created",
				Tags:    "success!",
			})

		}
		// templates.Layout(ApptForm(appts), "Appt Form").Render(r.Context(), w)

	}
	return nil
}

func (thing apptHandler) UpdateAppt(w http.ResponseWriter, r *http.Request) error {
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
	av := newValidation(idStr, cliId, dateStr, timeStr, row.Note.String, row.Status.String)
	templates.Layout(ApptForm(clients, av), "Edit Form").Render(r.Context(), w)
	return nil
}

type result[T any] struct {
	Val   T
	Error error
}

func (thing result[T]) HasError() bool {
	return thing.Error != nil
}

func parseInt64(str string) result[int64] {
	var val int64
	_, err := fmt.Sscan(str, &val)
	return result[int64]{val, err}
}

func (thing apptHandler) DeleteApptConfirm(req middleware.RequestScope) error {
	apptIdStr := chi.URLParam(req.Request(), "id")
	r := parseInt64(apptIdStr)
	if r.HasError() {
		return r.Error
	}
	req.HxTrigger(middleware.TriggerMsg{
		Event: "show-delete-confirmation",
		Data:  map[string]string{"deleteUrl": config.RouteTo(config.Routes().Admin.Appt.HX.DeleteAppt, "id", apptIdStr)},
	})
	return nil
}

func (thing apptHandler) ListAppt(req middleware.RequestScope) error {
	pageStr := req.QueryParam("page")
	page, _ := strconv.Atoi(pageStr)
	pgtor := dbstore.NewApptPagination(thing.apptRepo, page)
	// req := middleware.ReqScope(r.Context())
	// fmt.Printf("%p:\n", r)
	// fmt.Println()
	// fmt.Printf("%p\n", req.Request())
	ApptTableMain(pgtor).Render(req.Context(), req.Response())
	return nil
}
func (thing apptHandler) DeleteAppt(req middleware.RequestScope) error {
	apptIdStr := chi.URLParam(req.Request(), "id")
	r := parseInt64(apptIdStr)
	if r.HasError() {
		return r.Error
	}
	err := thing.apptRepo.DeleteAppt(r.Val)
	if err != nil {
		return err
	}
	// pgtor := dbstore.NewApptPagination(thing.apptRepo, 1)
	req.HxTrigger(middleware.TriggerMsg{
		Event:   "appointment-deleted",
		Message: "Appointment Deleted Successfully!",
	})
	rcomp := thing.Component(COMP_LIST_APPT)(req)
	// rcomp := thing.ListComponent(req)
	rcomp.Render()
	// ApptContent(pgtor).Render(req.Context(), req.Response())
	return nil
}

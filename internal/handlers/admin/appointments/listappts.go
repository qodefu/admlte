package appts

import (
	"goth/internal/store"
	"goth/internal/templates"
	"net/http"
)

func init() {

}

type ApptHandler struct {
	store store.ApptStore
}

func NewApptsHandler(st store.ApptStore) *ApptHandler {
	return &ApptHandler{store: st}
}

func (thing ApptHandler) CreateForm(w http.ResponseWriter, r *http.Request) error {
	templates.Layout(ApptForm(), "Appt Form").Render(r.Context(), w)
	return nil
}

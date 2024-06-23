package appts

import (
	"goth/internal/middleware"
	"goth/internal/store"
	"goth/internal/store/dbstore"
	"goth/internal/store/models"

	"github.com/a-h/templ"
)

type ListApptComp struct {
	Page      int
	compId    string                               `json:"-"`
	req       middleware.RequestScope              `json:"-"`
	apptRepo  store.ApptStore                      `json:"-"`
	paginator store.Pagination[models.ListApptRow] `json:"-"`
}

func (thing ListApptComp) Id() string {
	return thing.compId
}

func (thing ListApptComp) Content() templ.Component {

	// var schedCount, closedCount, totalCount int = 0, 0, 0
	thing.paginator = dbstore.NewApptPagination(thing.apptRepo, thing.Page)
	return ApptContent(thing)
}

func (thing ListApptComp) Render() error {
	thing.Content().Render(thing.req.Context(), thing.req.Response())
	return nil
}

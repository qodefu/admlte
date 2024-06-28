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
	SearchTxt string
	compId    string                               `json:"-"`
	req       middleware.RequestScope              `json:"-"`
	apptRepo  store.ApptStore                      `json:"-"`
	paginator store.Pagination[models.ListApptRow] `json:"-"`
}

func (thing *ListApptComp) DeleteConfirm(apptId int) {
	println("called!!!!!")
	thing.SearchTxt = "test delete confirm"
}

func (thing ListApptComp) GetSimple() {
	println("simple!!")
}

func (thing ListApptComp) Id() string {
	return thing.compId
}

func (thing ListApptComp) Content() templ.Component {

	// var schedCount, closedCount, totalCount int = 0, 0, 0
	thing.paginator = dbstore.NewApptPagination(thing.apptRepo, thing.Page)
	return apptListContent(thing)
}

func (thing ListApptComp) Render() error {
	thing.Content().Render(thing.req.Context(), thing.req.Response())
	return nil
}

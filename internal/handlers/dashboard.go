package handlers

import (
	"goth/internal/store/mockstore"
	"net/http"
)

type DashboardHandler struct {
	store mockstore.UserStore
}

func NewDashboardHandler(userRepo mockstore.UserStore) *DashboardHandler {
	return &DashboardHandler{
		userRepo,
	}
}

func (thing *DashboardHandler) PostCreateUser(w http.ResponseWriter, r *http.Request) {

}

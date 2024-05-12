package handlers

import (
	"goth/internal/store/dbstore"
	"net/http"
)

type DashboardHandler struct {
	store dbstore.UserStore
}

func NewDashboardHandler(userRepo dbstore.UserStore) *DashboardHandler {
	return &DashboardHandler{
		userRepo,
	}
}

func (thing *DashboardHandler) PostCreateUser(w http.ResponseWriter, r *http.Request) {

}

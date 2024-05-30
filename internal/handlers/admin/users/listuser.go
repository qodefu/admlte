package users

import (
	"goth/internal/store/dbstore"
	"goth/internal/templates/admin"
	"goth/internal/validator"
	v "goth/internal/validator"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ListUsers struct {
	userStore *dbstore.UserStore
}

func NewListUsersHandler(us *dbstore.UserStore) *ListUsers {
	return &ListUsers{us}
}

func (thing *ListUsers) HxAddUserModal(w http.ResponseWriter, r *http.Request) {

	// id := chi.URLParam(r, "id")

	w.Header().Set("HX-Trigger", "show-global-modal-form")
	uv := admin.UserValidations{
		v.New("name", "", nil),
		v.New("email", "", nil),
		v.New("password", "", nil),
		v.New("passwordConfirmation", "", nil),
	}
	admin.UserModalContent(uv, false).Render(r.Context(), w)

}

func (thing *ListUsers) HxEditUserModal(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	user, _ := thing.userStore.GetUser(email)

	uv := admin.UserValidations{
		v.New("name", user.Name, nil),
		v.New("email", user.Email, nil),
		v.New("password", "", nil),
		v.New("passwordConfirmation", "", nil),
	}

	w.Header().Set("HX-Trigger", "show-global-modal-form")
	admin.UserModalContent(uv, true).Render(r.Context(), w)
}

func (thing *ListUsers) HxDeleteUserModal(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	w.Header().Set("HX-Trigger", "show-global-modal-form")
	admin.DeleteModalContent(email).Render(r.Context(), w)
}

func (thing *ListUsers) HxCreateUser(w http.ResponseWriter, r *http.Request) {
	// first validate
	// validate fail, return form with error
	nameVal := r.FormValue("name")
	emailVal := r.FormValue("email")
	pwdVal := r.FormValue("password")
	pwdConfirm := r.FormValue("passwordConfirmation")

	validations := admin.UserValidations{
		Name:                 v.New("name", nameVal, v.NotEmpty("Name")),
		Email:                v.New("email", emailVal, v.NotEmpty("Email"), v.EmailFmt),
		Password:             v.New("password", pwdVal, v.NotEmpty("Password")),
		PasswordConfirmation: v.New("passwordConfirmation", pwdConfirm, v.NotEmpty("PasswordConfirmation"), v.PasswordMatch(pwdVal)),
	}
	validator.ValidateFields(&validations)

	// validation pass, create user, return empty form
	if validator.ValidationOk(&validations) {
		thing.userStore.CreateUser(nameVal, emailVal, pwdVal)
		w.Header().Set("HX-Trigger", `{"close-global-modal-form": [{"foo": 1, "message": "User Added", "tags": "Success!"}]}`)
	}

	admin.UserForm(validations, false).Render(r.Context(), w)
}

func (thing *ListUsers) existingUser(val string) v.VResult {

	_, err := thing.userStore.GetUser(val)
	if err != nil {
		return v.VResult{false, "email must exist "}

	}
	return v.VResult{true, ""}
}

func (thing *ListUsers) HxUpdateUser(w http.ResponseWriter, r *http.Request) {
	// first validate
	// validate fail, return form with error
	nameVal := r.FormValue("name")
	emailVal := r.FormValue("email")
	pwdVal := r.FormValue("password")
	pwdConfirm := r.FormValue("passwordConfirmation")

	validations := admin.UserValidations{
		Name:                 v.New("name", nameVal, v.NotEmpty("Name")),
		Email:                v.New("email", emailVal, v.NotEmpty("Email"), v.EmailFmt, thing.existingUser),
		Password:             v.New("password", pwdVal),
		PasswordConfirmation: v.New("passwordConfirmation", pwdConfirm, v.PasswordMatch(pwdVal)),
	}
	validator.ValidateFields(&validations)

	// validation pass, create user, return empty form
	if validator.ValidationOk(&validations) {
		thing.userStore.UpdateUser(nameVal, emailVal, pwdVal)
		w.Header().Set("HX-Trigger", `{"close-global-modal-form": [{"foo": 1, "message": "User `+emailVal+` Edited and saved", "tags": "Success!"}]}`)
	}

	admin.UserForm(validations, true).Render(r.Context(), w)
}

func (thing *ListUsers) HxDeleteUser(w http.ResponseWriter, r *http.Request) {
	emailVal := chi.URLParam(r, "email")
	thing.userStore.DeleteUser(emailVal)
	w.Header().Set("HX-Trigger", `{"close-global-modal-form": [{"foo": 1, "message": "User `+emailVal+` deleted", "tags": "Success!"}]}`)

}

func (thing *ListUsers) HxListUsers(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	pgNum, err := strconv.Atoi(page)
	if err != nil {
		panic(err)
	}
	paginator := dbstore.NewUserPagination("/admin/users/hx/list", thing.userStore, pgNum)
	admin.UserTableMain(paginator).Render(r.Context(), w)
}
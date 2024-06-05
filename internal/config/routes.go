package config

import (
	"log/slog"
	"strings"
)

type Admin struct {
	Base      string
	Dashboard Dash
	Appt      Appt
	Users     Users
}

type Dash struct {
	Base string
}
type Appt struct {
	Base       string
	Create     string
	SaveNew    string
	UpdateAppt string
}

type Users struct {
	Base string
	HX   usershx
}
type usershx struct {
	Create          string
	Update          string
	Delete          string
	List            string
	AddUserModal    string
	EditUserModal   string
	DeleteUserModal string
}

type route struct {
	Admin Admin
}

func Routes() route {
	return route{
		Admin: Admin{
			Base: "/admin",
			Dashboard: Dash{
				Base: "/admin/dashboard",
			},
			Appt: Appt{
				Base:       "/admin/appointments",
				Create:     "/admin/appointments/create",
				UpdateAppt: "/admin/appointments/{id}/edit",
				SaveNew:    "/admin/appointments/saveNew",
			},
			Users: Users{
				Base: "/admin/users",
				HX: usershx{
					List:            "/admin/users/hx/list",
					Create:          "/admin/users/hx/createUser",
					Update:          "/admin/users/hx/updateUser",
					Delete:          "/admin/users/hx/deleteUser/{id}",
					AddUserModal:    "/admin/users/hx/addUserModal",
					DeleteUserModal: "/admin/users/hx/deleteUserModal/{id}",
					EditUserModal:   "/admin/users/hx/editUserModal/{id}",
				},
			},
		},
	}

}

func RouteTo(url string, urlParams ...string) string {
	ret := url
	if len(urlParams)%2 != 0 {
		slog.Error("routing url param number mismatch", urlParams)
	}

	for i := 0; i < len(urlParams)/2; i += 1 {
		ret = strings.Replace(ret, "{"+urlParams[2*i]+"}", urlParams[2*i+1], 1)
	}

	return ret
}

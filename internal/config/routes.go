package config

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
	Base   string
	Create string
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
				Base:   "/admin/appointments",
				Create: "/admin/appointments/create",
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

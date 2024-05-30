// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package appts

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

// import m "goth/internal/middleware"
import v "goth/internal/validator"
import "fmt"
import "goth/internal/store"
import "strconv"
import "goth/internal/utils"

var idGen = utils.NewIdGen("appointments")

type ApptValidations struct {
	Name                 v.Validation
	Email                v.Validation
	Password             v.Validation
	PasswordConfirmation v.Validation
}

// templ modalHeader(edit bool) {
// 	<div class="modal-header">
// 		if edit {
// 			<h5 class="modal-title" id="exampleModalLabel">Edit Appointment</h5>
// 		} else {
// 			<h5 class="modal-title" id="exampleModalLabel">Add New Appointment</h5>
// 		}
// 		<button type="button" class="close" data-dismiss="modal" aria-label="Close">
// 			<span aria-hidden="true">&times;</span>
// 		</button>
// 	</div>
// }

// templ DeleteModalContent(email string) {
// 	<div class="modal-content" id="globalModalContent">
// 		<div class="modal-header">
// 			<h5 class="modal-title" id="exampleModalLabel">Delete Appt</h5>
// 			<button type="button" class="close" data-dismiss="modal" aria-label="Close">
// 				<span aria-hidden="true">&times;</span>
// 			</button>
// 		</div>
// 		<div class="modal-body">
// 			<h4>Are you sure you want to delte?</h4>
// 		</div>
// 		<div class="modal-footer">
// 			<button type="button" class="btn btn-secondary" data-dismiss="modal"><i class="fa fa-times mr-2"></i>Cancel</button>
// 			<button type="submit" class="btn btn-danger" hx-delete={ fmt.Sprintf("/admin/appts/hx/deleteAppt/%s", email) }><i class="fa fa-trash mr-2"></i>Delete Appt</button>
// 		</div>
// 	</div>
// }

//	templ ApptModalContent(uv ApptValidations, edit bool) {
//		<div class="modal-content" id="globalModalContent">
//			@modalHeader(edit)
//			@ApptForm(uv, edit)
//		</div>
//	}
func ApptTableMain(paginator store.Pagination[store.Appt]) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 1)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ApptTable(paginator).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 2)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(paginator.Total()))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/handlers/admin/appointments/appointment.templ`, Line: 62, Col: 43}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 3)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(paginator.PreviousPageUrl())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/handlers/admin/appointments/appointment.templ`, Line: 68, Col: 110}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 4)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs("<")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/handlers/admin/appointments/appointment.templ`, Line: 68, Col: 153}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 5)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, i := range paginator.Pages() {
			if paginator.CurrentPage() == i {
				templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 6)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(paginator.PageUrl(i))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/handlers/admin/appointments/appointment.templ`, Line: 73, Col: 105}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 7)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(i))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/handlers/admin/appointments/appointment.templ`, Line: 73, Col: 125}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 8)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 9)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(paginator.PageUrl(i))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/handlers/admin/appointments/appointment.templ`, Line: 76, Col: 127}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 10)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var8 string
				templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(i))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/handlers/admin/appointments/appointment.templ`, Line: 76, Col: 147}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 11)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 12)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(paginator.NextPageUrl())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/handlers/admin/appointments/appointment.templ`, Line: 85, Col: 106}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 13)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(">")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/handlers/admin/appointments/appointment.templ`, Line: 85, Col: 135}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 14)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func ApptContent(paginator store.Pagination[store.Appt]) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var11 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var11 == nil {
			templ_7745c5c3_Var11 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 15)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ApptTableMain(paginator).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 16)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

//	templ ApptForm(uvs ApptValidations, edit bool) {
//		<form hx-post={ m.If(edit, "/admin/appts/hx/updateAppt", "/admin/appts/hx/createAppt") }>
//			<div class="modal-body">
//				<div class="form-group">
//					<label for={ uvs.Name.Key + "_id" }>Name </label>
//					<input
//						type="text"
//						class={ "form-control" , m.If(uvs.Name.Result.Valid, "", "is-invalid") }
//						name={ uvs.Name.Key }
//						value={ uvs.Name.Value }
//						id={ uvs.Name.Key + "_id" }
//					/>
//					<div class="invalid-feedback">
//						{ uvs.Name.Result.ErrorMsg }
//					</div>
//				</div>
//				<div class="form-group">
//					<label for={ uvs.Email.Key + "_id" }>Email </label>
//					<input
//						type="text"
//						class={ "form-control" , m.If(uvs.Email.Result.Valid, "", "is-invalid") }
//						name={ uvs.Email.Key }
//						value={ uvs.Email.Value }
//						id={ uvs.Email.Key + "_id" }
//					/>
//					<div class="invalid-feedback">
//						{ uvs.Email.Result.ErrorMsg }
//					</div>
//				</div>
//				<div class="form-group">
//					<label for={ uvs.Password.Key + "_id" }>Password </label>
//					<input
//						type="text"
//						class={ "form-control" , m.If(uvs.Password.Result.Valid, "", "is-invalid") }
//						name={ uvs.Password.Key }
//						value={ uvs.Password.Value }
//						id={ uvs.Password.Key + "_id" }
//					/>
//					<div class="invalid-feedback">
//						{ uvs.Password.Result.ErrorMsg }
//					</div>
//				</div>
//				<div class="form-group">
//					<label for={ uvs.PasswordConfirmation.Key + "_id" }>Password Confirmation</label>
//					<input
//						type="text"
//						class={ "form-control" , m.If(uvs.PasswordConfirmation.Result.Valid, "", "is-invalid") }
//						name={ uvs.PasswordConfirmation.Key }
//						value={ uvs.PasswordConfirmation.Value }
//						id={ uvs.PasswordConfirmation.Key + "_id" }
//					/>
//					<div class="invalid-feedback">
//						{ uvs.PasswordConfirmation.Result.ErrorMsg }
//					</div>
//				</div>
//			</div>
//			<div class="modal-footer">
//				<button type="button" class="btn btn-secondary" data-dismiss="modal"><i class="fa fa-times mr-2"></i>Cancel</button>
//				if edit {
//					<button type="submit" class="btn btn-primary"><i class="fa fa-save mr-2"></i>Save Changes</button>
//				} else {
//					<button type="submit" class="btn btn-primary"><i class="fa fa-save mr-2"></i>Save</button>
//				}
//			</div>
//		</form>
//		<script>
//			htmx.on('close-global-modal-form', event => {
//			  	htmx.ajax(
//				'GET',
//				'/admin/appts/hx/list?page=1', {
//				  target: '#apptTableMain',
//				  swap:'outerHTML'
//				});
//			})
//		</script>
//	}
func ApptTable(paginator store.Pagination[store.Appt]) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var12 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var12 == nil {
			templ_7745c5c3_Var12 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 17)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i, appt := range paginator.Items() {
			templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 18)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", i+1))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/handlers/admin/appointments/appointment.templ`, Line: 236, Col: 45}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 19)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(appt.ApptTime.Format("2006/01/02"))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/handlers/admin/appointments/appointment.templ`, Line: 238, Col: 45}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 20)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var15 string
			templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(appt.ApptTime.Format("03:04:05"))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/handlers/admin/appointments/appointment.templ`, Line: 239, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 21)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var16 string
			templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(appt.Status)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/handlers/admin/appointments/appointment.templ`, Line: 240, Col: 22}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 22)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 23)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
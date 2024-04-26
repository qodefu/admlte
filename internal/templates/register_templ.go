// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

// 'RegisterPage' template generates the HTML for a user registration form.
// This form allows new users to create an account by providing their email and password.
func RegisterPage() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div hx-ext=\"response-targets\"><h1>Register an account</h1><form hx-post=\"/register\" hx-trigger=\"submit\" hx-target-401=\"#register-error\"><div id=\"register-error\"></div><div><label for=\"email\">Your email</label> <input type=\"email\" name=\"email\" id=\"email\" placeholder=\"name@company.com\" required=\"\"></div><div><label for=\"password\">Password</label> <input type=\"password\" name=\"password\" id=\"password\" placeholder=\"••••••••\" required=\"\"></div><button type=\"submit\">Register</button><p>Already have an account? <a href=\"/login\">Login</a></p></form></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

// 'RegisterSucces' (note the typo in 'Success') template provides feedback to the user after a successful registration.
// It includes a message and a link to the login page, allowing the user to proceed to sign in with their new account.
func RegisterSucces() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1>Registration successful</h1><p>Go to <a href=\"login\">login</a></p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

/*
Key Points:
    HTMX Integration: The registration form leverages HTMX for dynamic, asynchronous interactions. This allows for a smoother user experience by submitting the form data and handling responses without reloading the page.
    Form Design: The form includes input fields for the user's email and password, each marked as required. This ensures that the form cannot be submitted unless both fields are filled out. The placeholders provide hints to the user about the expected values.
    Error Handling: A designated area (<div id="register-error"></div>) is included to display any error messages that might arise during the registration process, such as when the email is already in use or doesn't meet validation criteria.
    Navigation Link: After successful registration, users are encouraged to log in through a direct link to the login page. This guides users through the next step in accessing their new account.
    Accessibility and UX: The form uses <label> elements associated with each input field via the for attribute, improving accessibility. Placeholders in the input fields provide additional guidance to users.

This register.templ file illustrates a clean and user-friendly approach to handling user registration in a web application, with thoughtful considerations for error handling, user guidance, and integration with modern web technologies like HTMX.
*/

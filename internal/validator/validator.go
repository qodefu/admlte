package validator

import (
	"reflect"
	"strings"
)

type VResult struct {
	Valid    bool
	ErrorMsg string
}
type Validator func(v string) VResult
type Validators []Validator
type Validation struct {
	Key        string
	Value      string
	Validators Validators
	Result     VResult
}

func New(key, val string, vtors ...Validator) Validation {
	return Validation{
		Key:        key,
		Value:      val,
		Validators: vtors,
		Result:     VResult{true, ""},
	}

}

func reflectInvoke(addr reflect.Value, name string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	m, _ := addr.Type().MethodByName(name)
	m.Func.Call([]reflect.Value{addr})
	// obj.MethodByName(name).Call(inputs)
}

func ValidateFields[T any](obj *T) {
	valPtr := reflect.ValueOf(obj)
	// fmt.Println(valPtr.Elem().NumField())
	for i := 0; i < valPtr.Elem().NumField(); i++ {
		f := valPtr.Elem().Field(i).Addr()
		if f.Type().AssignableTo(reflect.TypeOf(&Validation{})) {
			reflectInvoke(f, "Validate")
		}
	}
}

func ValidationOk[T any](obj *T) bool {
	ret := true
	hasValidation := false
	valPtr := reflect.ValueOf(obj)
	// fmt.Println(valPtr.Elem().NumField())
	for i := 0; i < valPtr.Elem().NumField(); i++ {
		f := valPtr.Elem().Field(i).Addr()
		if f.Type().AssignableTo(reflect.TypeOf(&Validation{})) {
			hasValidation = true
			ret = ret && f.Interface().(*Validation).Result.Valid
		}
	}
	return hasValidation && ret
}

// func Validate(validations ...Validation) bool {
// 	ok := true
// 	for _, vDesc := range validations {
// 		vDesc.Validate()
// 		if ok && !res.Valid {
// 			ok = false
// 		}
// 		vDesc.Result = res
// 	}
// 	return ok
// }

func (thing *Validation) Validate() {
	thing.Result = VResult{true, ""}
	for _, f := range thing.Validators {
		r := f(thing.Value)
		if !r.Valid {
			thing.Result = r
			return
		}
	}
}

func NotEmpty(msg string) Validator {
	return func(value string) VResult {
		if len(value) > 0 {
			return VResult{true, ""}
		}
		return VResult{false, "The " + msg + " field is required"}
	}
}

func EmailFmt(value string) VResult {
	if strings.Index(value, "@") > 0 {
		return VResult{true, ""}
	}
	return VResult{false, "A valid email address must be provided"}
}

func PasswordMatch(p1 string) Validator {
	return func(value string) VResult {
		if p1 == value {
			return VResult{true, ""}
		}
		return VResult{false, "Password must match"}
	}

}

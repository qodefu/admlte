package main

import (
	"encoding/json"
	"goth/internal/components"
	appts "goth/internal/handlers/admin/appointments"
	"reflect"
	"strings"
)

func reflectInvoke(addr reflect.Value, name string, args ...any) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	m, _ := addr.Type().MethodByName(name)
	// type coersion bc JSON parse force numeric to float64
	if len(args) > 0 {
		// println(len(args), m.Func.Type().NumIn())
		for i := 1; i < m.Func.Type().NumIn(); i++ {
			t1 := m.Func.Type().In(i).Name()
			t2 := reflect.TypeOf(args[i-1]).Name()
			if t1 != t2 && strings.Contains(t2, "float") {
				//float converstion to in
				inputs[i-1] = reflect.ValueOf(args[i-1]).Convert(reflect.TypeOf(0))
			}
		}
	}
	// fmt.Printf("%v: %v", m, args)
	return m.Func.Call(append([]reflect.Value{addr}, inputs...))
	// obj.MethodByName(name).Call(inputs)
}

func invoke[T any](thing components.RComp, callStr string) T {

	reflectThing := reflect.ValueOf(thing)
	xs := strings.FieldsFunc(callStr, func(r rune) bool {
		return r == '(' || r == ')' || r == ','
	})
	funcName := xs[0]
	xs = xs[1:]
	var args = make([]any, len(xs))
	for i, x := range xs {
		json.Unmarshal([]byte(x), &args[i])
		// switch args[i].(type) {
		// //support int only
		// case float64:
		// 	args[i], _ = strconv.Atoi(x)
		// }
	}

	ret := reflectInvoke(reflectThing, funcName, args...)
	return ret[0].Interface().(T)
}

func main() {
	var comp = appts.ListApptComp{
		Page:      8,
		SearchTxt: "hello",
	}
	// s := invoke[string](comp, `GetSearch(5, "hello")`)
	// fmt.Println("call result: ", s)
	// ret = reflectInvoke(compPtr, "Content")[0]
	// s := `GetSearch(5, "hello")`
	println(invoke[string](comp, "GetSimple"))
	println(invoke[string](comp, `GetWithArgs(8,"world", true)`))
	// args := strings.FieldsFunc(s, func(r rune) bool {
	// 	return r == '(' || r == ')' || r == ','
	// })
	// var i any
	// json.Unmarshal([]byte(args[2]), &i)
	// switch i.(type) {
	// case float32:
	// 	println("float")
	// case string:
	// 	println("string")
	// default:
	// 	println("unknown type")
	// }

}

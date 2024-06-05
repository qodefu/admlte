package main

import (
	"log/slog"
	"strings"
)

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

func main() {
	var data = make(map[string]any)
	data["ev1"] = 1
	data["foo"] = "zz"
	data["close-vent"] = append([]map[string]string{}, map[string]string{"a": "x", "b": "2"})
	// d0 := "06/04/2024"
	// t0 := "10:09 AM"
	// t, _ := time.Parse("01/02/2006 3:04 PM", d0+" "+t0)
	// fmt.Println(t)
	// fmt.Println(trigger(DeclMsg{"do-event", "User created", "Success", "zz"}, DeclMsg{"dont-event", "User created", "Success", "zz"}))

}

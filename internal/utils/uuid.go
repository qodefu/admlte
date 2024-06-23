package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log/slog"
	"time"
)

type IdGen struct {
	prefix string
	dict   map[string]string
}

var gens = make(map[string]IdGen)

func (thing IdGen) Id(elem string) string {
	key := thing.prefix + "-" + elem + "-"
	if _, ok := thing.dict[key]; !ok {
		thing.dict[key] = key + unique(8)
	}
	return thing.dict[key]
}
func (thing IdGen) IdRef(elem string) string {
	return "#" + thing.Id(elem)
}

func NewIdGen(module string) IdGen {
	if _, ok := gens[module]; ok {
		//error
		slog.Error("Duplicate Id Generator detected!!!")
	}
	gens[module] = IdGen{
		prefix: module,
		dict:   make(map[string]string),
	}

	return gens[module]
}

func unique(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	return hex.EncodeToString(b)
}

func DateFormat(t time.Time) string {
	return t.Format("01/02/2006")
}

func TimeFormat(t time.Time) string {
	return t.Format("03:04:05 PM")
}

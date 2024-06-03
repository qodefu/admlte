package main

import (
	"fmt"
	"time"
)

func main() {
	t, _ := time.Parse("01/02/2006 3:04 PM", "06/12/2024 1:11 PM")
	fmt.Println(t)

}

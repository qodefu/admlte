package main

import (
	"fmt"
	"time"
)

func main() {
	d0 := "06/04/2024"
	t0 := "10:09 AM"
	t, _ := time.Parse("01/02/2006 3:04 PM", d0+" "+t0)
	fmt.Println(t)

}

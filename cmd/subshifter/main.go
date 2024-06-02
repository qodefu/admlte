package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func tomilli(str string) int64 {
	t, err := time.Parse("15:04:05,999", strings.Trim(str, " "))
	if err != nil {
		log.Panic(err)
	}
	return t.AddDate(1970, 0, 0).UnixMilli()
}

func toStr(epoch int64) string {
	milli := epoch % 1000
	epoch /= 1000
	sec := epoch % 60
	epoch /= 60
	mn := epoch % 60
	hr := epoch / 60

	return fmt.Sprintf("%02d:%02d:%02d,%03d", hr, mn, sec, milli)
}

func shift(str string) string {
	slc := strings.Split(str, "-->")
	from, to := strings.Trim(slc[0], " "), strings.Trim(slc[1], " ")

	// fmt.Println(tomilli(from), tomilli(to))
	return toStr(tomilli(from)+3000) + " --> " + toStr(tomilli(to)+3000)

}
func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: subshift [srt file]")
	}

	fname := os.Args[1]
	fhandle, _ := os.Open(fname)
	scanner := bufio.NewScanner(fhandle)
	bldr := new(strings.Builder)
	writer := bufio.NewWriter(bldr)

	ln := 1
	// i := 4
	for scanner.Scan() {
		state := ln % 4
		if state == 2 {
			writer.WriteString(shift(scanner.Text()))
			writer.WriteString("\n")
		} else {
			writer.WriteString(scanner.Text())
			writer.WriteString("\n")
		}
		ln += 1
	}
	fhandle.Close()
	writer.Flush()
	fmt.Println(bldr.String())
	// now write srt file
	fhandle, err := os.Open(fname)
	if err != nil {
		log.Panic(err)
	}
	err = os.WriteFile(fname, []byte(bldr.String()), 0644)
	fmt.Println("done")
	if err != nil {
		log.Panic(err)
	}
	fhandle.Close()

}

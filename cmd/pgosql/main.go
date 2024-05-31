package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
)

func run(host, user, pwd, db, fpath string, port int) error {
	f, err := os.Open(fpath)
	if err != nil {
		return err
	}
	ctx := context.Background()
	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, pwd, host, port, db)
	conn, err := pgx.Connect(ctx, connectStr)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	bytes, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	contents := string(bytes)

	for _, sqlStmt := range strings.Split(contents, ";\n") {
		// skip empty stmt
		if strings.Trim(sqlStmt, " \n\t\r") == "" {
			continue
		}
		r, err := conn.Exec(context.Background(), sqlStmt)
		if err == nil {
			fmt.Println(r.String(), "Stmt Executed")
			if r.Delete() {
				fmt.Print("Deleted", r.RowsAffected())
			}
			if r.Insert() {
				fmt.Print("Inserted", r.RowsAffected())
			}
			if r.Select() {
				fmt.Print("Selected", r.RowsAffected())
			}
			if r.Update() {
				fmt.Print("Updated", r.RowsAffected())
			}

		}
	}
	return nil

}
func main() {
	var user = flag.String("user", "", "user")
	var pwd = flag.String("password", "", "password")
	var host = flag.String("host", "", "host")
	var port = flag.Int("port", 5432, "port")
	var db = flag.String("db", "", "database")
	var file = flag.String("file", "", "sql file")
	flag.Parse()
	fmt.Println("loading script: ", *file)
	if err := run(*host, *user, *pwd, *db, *file, *port); err != nil {
		log.Fatal(err)
	}
}

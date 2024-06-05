package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"goth/internal/store/models"
	"log"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ctx context.Context
var conn *pgxpool.Pool
var queries *models.Queries

type Errors struct {
	errors []error
}

func (thing Errors) Error() string {
	e := errors.Join(thing.errors...)
	return e.Error()
}

func fakeUsers() error {
	_, err := conn.Exec(ctx, "DELETE FROM users;")
	if err != nil {
		return err
	}
	for i := 0; i < 10; i++ {
		// create an user
		_, err := queries.CreateUser(ctx, models.CreateUserParams{
			Name:     gofakeit.Name(),
			Email:    pgtype.Text{String: gofakeit.Email(), Valid: true},
			Password: pgtype.Text{String: gofakeit.Password(true, true, true, false, false, 8), Valid: true},
		})
		if err != nil {
			return err
		}
		// log.Println(insertedAuthor)

	}
	return nil

}

func fakeClis() ([]int, error) {
	_, err := conn.Exec(ctx, "DELETE FROM clients;")
	if err != nil {
		return nil, err
	}
	var clids []int
	for i := 0; i < 10; i++ {
		// create an user
		client, err := queries.CreateClient(ctx, gofakeit.Name())
		if err != nil {
			return nil, err
		}
		clids = append(clids, int(client.ID))

	}
	return clids, nil
}

func fakeAppts(clids []int) error {
	_, err := conn.Exec(ctx, "DELETE FROM appointments;")
	if err != nil {
		return err
	}
	stat := []string{"CLOSED", "SCHEDULED"}
	for i := 0; i < 10; i++ {
		// create an user
		_, err := queries.CreateAppt(ctx, models.CreateApptParams{
			ClientID: pgtype.Int8{Int64: int64(gofakeit.RandomInt(clids)), Valid: true},
			ApptTime: pgtype.Timestamp{Time: gofakeit.Date(), Valid: true},
			Note:     pgtype.Text{String: gofakeit.Paragraph(2, 1, 35, " "), Valid: true},
			Status:   pgtype.Text{String: gofakeit.RandomString(stat), Valid: true},
		})
		if err != nil {
			return err
		}
		// log.Println(insertedAuthor)

	}
	return nil
}

var err error

func run(host, user, pwd, db string, port int) error {
	ctx = context.Background()
	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, pwd, host, port, db)
	conn, err = pgxpool.New(ctx, connectStr)
	if err != nil {
		return err
	}
	defer conn.Close()

	queries = models.New(conn)

	gofakeit.Seed(42)
	usrerr := fakeUsers()
	clids, clierr := fakeClis()
	appterr := fakeAppts(clids)

	// log.Println(authors)
	// // get the author we just inserted
	// fetchedAuthor, err := queries.GetUser(ctx, insertedAuthor.ID)
	// if err != nil {
	// 	return err
	// }

	// // prints true
	// log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return errors.Join(usrerr, clierr, appterr)
}

func main() {
	var user = flag.String("user", "", "user")
	var pwd = flag.String("password", "", "password")
	var host = flag.String("host", "", "host")
	var port = flag.Int("port", 5432, "port")
	var db = flag.String("db", "", "database")
	flag.Parse()
	if err := run(*host, *user, *pwd, *db, *port); err != nil {
		log.Fatal(err)
	}
}

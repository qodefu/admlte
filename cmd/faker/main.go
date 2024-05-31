package main

import (
	"context"
	"fmt"
	"goth/internal/store/models"
	"log"

	"github.com/jackc/pgx/v5"
)

func run() error {
	ctx := context.Background()
	connectStr := "postgres://app1:app1@lovelace:9000/app1"
	conn, err := pgx.Connect(ctx, connectStr)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := models.New(conn)

	// list all authors
	authors, err := queries.ListUsers(ctx)
	if err != nil {
		return err
	}
	for _, u := range authors {
		fmt.Println(u)
	}
	// log.Println(authors)

	// create an author
	// insertedAuthor, err := queries.CreateUser(ctx, models.CreateUserParams{
	// 	Name:     "kai",
	// 	Email:    pgtype.Text{String: "foo@bar.com", Valid: true},
	// 	Password: pgtype.Text{String: "pwd", Valid: true},
	// })
	// if err != nil {
	// 	return err
	// }
	// log.Println(insertedAuthor)

	// // get the author we just inserted
	// fetchedAuthor, err := queries.GetUser(ctx, insertedAuthor.ID)
	// if err != nil {
	// 	return err
	// }

	// // prints true
	// log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

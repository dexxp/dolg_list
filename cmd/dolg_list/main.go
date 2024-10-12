package main

import (
	"fmt"
	"context"
	"os"
	"github.com/jackc/pgx/v5"
	"net/http"
)

func main() {
	dsn := "postgres://postgres:postgres@postgres:5432/postgres"
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), `
		create table if not exists dolgs (
			id serial primary key,
			subject text,
			completed boolean	
		);
	`)

	if err != nil {
		fmt.Println("Unable to create table dolg ", err.Error())
		os.Exit(1)
	}

	
	q := `
		insert into dolgs (subject, completed) values ($1, $2)
		returning id
	`

	var id int
	if err := conn.QueryRow(context.Background(), q, "Электротехника", false).Scan(&id); err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Created dolg. ID = ", id)
	http.ListenAndServe(":8080", nil)
}

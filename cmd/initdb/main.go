package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	println("Init Database")

	postgresUrl := "postgres://alex:qwerty@localhost:5532/mydatabase"

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, postgresUrl)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close(context.Background())

	if err := conn.Ping(context.Background()); err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < 50; i++ {
		username := fmt.Sprintf("user-%03d", i)
		password := fmt.Sprintf("pass-%03d", i)

		id, err := InsertNewAccount(conn, username, password)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("New user ID:", id.String())
		}
	}

	// bcrypt.CompareHashAndPassword()

}

func InsertNewAccount(conn *pgx.Conn, username string, password string) (*uuid.UUID, error) {
	const qInsertAccount string = `
		INSERT INTO public.account
			(username, password)
			VALUES (LOWER(BTRIM($1)), $2)
		RETURNING id
		;
	`
	newID := &uuid.UUID{}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	} else {
		row := conn.QueryRow(context.Background(), qInsertAccount, username, encryptedPassword)
		if err := row.Scan(newID); err != nil {
			return nil, err
		} else {
			return newID, nil
		}
	}
}

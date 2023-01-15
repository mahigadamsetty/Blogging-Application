package database

import (
	"blog1/ent"
	"blog1/ent/migrate"
	"context"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var EntClient *ent.Client

// Hash hashes the given password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func init() {

	client, err := ent.Open("postgres", "postgresql://root:secret@localhost:5432/blog?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		panic(err)
	}

	EntClient = client

	mail := "mahigadamsetty@gmail.com"
	password := "Blog@1234"

	hashedPassword, err := Hash(password)
	if err != nil {
		panic(err)
	}

	// Seeding the admin's credentials
	err = CreateUser(mail, string(hashedPassword))
	if err != nil {
		panic(err)
	}
}

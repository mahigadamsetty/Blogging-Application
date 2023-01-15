package database

import (
	"blog1/ent"
	"blog1/ent/user"
	"context"
	"time"
)

// CreateUser creates user with the given mail and password.
func CreateUser(mail, password string) (*ent.User, error) {
	user, err := EntClient.User.Create().
		SetMail(mail).
		SetPassword(password).
		SetCreatedAt(time.Now()).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByMail gets user by mail.
func GetUserByMail(mail string) (*ent.User, error) {
	user, err := EntClient.User.Query().Where(user.Mail(mail)).Only(context.Background())
	if err != nil {
		return nil, err
	}

	return user, nil

}

package handler

import (
	"blog1/auth"
	"blog1/database"
	"errors"
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

func (u *User) validateLoginCredentials() error {
	if u.Password == "" {
		return errors.New("Required Password")
	}
	if u.Mail == "" {
		return errors.New("Required Email")
	}
	if err := checkmail.ValidateFormat(u.Mail); err != nil {
		return errors.New("Invalid Email")
	}
	return nil
}

// VerifyPassword compares the hashed password with given password.
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// HandleLogin is a handler on the login api.
func HandleLogin(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		ResponseJSON(c, http.StatusBadRequest, 400, "not able to parse the input data", nil)
		return
	}

	if err := user.validateLoginCredentials(); err != nil {
		ResponseJSON(c, http.StatusBadRequest, 400, err.Error(), nil)
		return
	}

	entUser, err := database.GetUserByMail(user.Mail)
	if err != nil {
		ResponseJSON(c, http.StatusInternalServerError, 500, err.Error(), nil)
		return
	}

	err = VerifyPassword(entUser.Password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		ResponseJSON(c, http.StatusInternalServerError, 500, err.Error(), nil)
		return
	}

	tokenString, err := auth.GenerateJWT(user.Mail, "Admin")
	if err != nil {
		ResponseJSON(c, http.StatusInternalServerError, 500, err.Error(), nil)
		return
	}

	token := Token{
		Role:        "Admin",
		Email:       user.Mail,
		TokenString: tokenString,
	}

	ResponseJSON(c, http.StatusOK, 200, "", token)

}

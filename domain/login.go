package domain

import (
	"database/sql"
	"errors"
	"go/token"
	"log"
	"time"
)

type Login struct {
	Username   string         `db:"username"`
	CustomerId sql.NullString `db:"customer_id"`
	Account    sql.NullString `db:"account_id"`
	Role       string         `db:"role"`
}

func (l Login) GenerateToken() (*string, error) {
	var claims jwt.MapClaims
	if l.Account.Valid && l.CustomerId.Valid {
		claims = l.claimsForUser()
	}else {
		claims = l.claimsForAdmin()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenAsString, err := token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		log.Println("Failed while signing token: " + err.Error())
		return nil, errors.New(text:"cannot generate token")
	}
	return &signedTokenAsString, nil
}

func (l Login) claimsForUser() jwt.MapClaims {
	account := string.Split(l.Account.String, sep:",")
	return jwt.MapClaims{
		"customer_id": l.CustomerId.String,
		"role": l.Role,
		"username": l.Username,
		"account": account,
		"exp": time.Now().Add(TOCKEN_DURATION).Unix()
	}
}

func (l Login) claimsForAdmin() jwt.MapClaims {
	return jwt.MapClaims{
		"role": l.Role,
		"username": l.Username,
		"exp": time.Now().Add(TOCKEN_DURATION).Unix()
	}
}

type AuthRepository interface {
	FindBy(string, string) (*Login, error)
}

package model

import (
	"database/sql"
	"encoding/hex"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/scrypt"
)

const salt = "lGSLyiFmTtKzycS6tsXgutflsKh2CGvQuHpgGXUEnTWltcVswi6"

func UserLoginByUsernameAndPassword(db *sql.DB, username string, password string) (user User, err error) {
	// GoDoc says "The recommended parameters for interactive logins as of 2009 are N=16384, r=8, p=1"
	// https://godoc.org/golang.org/x/crypto/scrypt
	k, err := scrypt.Key([]byte(username+password), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		return
	}

	h := hex.EncodeToString(k[:])
	q := "SELECT id, username FROM webvuln_user WHERE username = '" + username + "' AND password = '" + h + "'"
	row := db.QueryRow(q)
	user, err = ScanUser(row)
	if err != nil {
		return
	}

	return
}

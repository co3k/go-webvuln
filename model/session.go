package model

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"time"
)

func GenerateSession(db *sql.DB, userId int) (sessionId string, err error) {
	sessionId = strconv.FormatInt(time.Now().Unix(), 10)
	q := "INSERT INTO webvuln_session (session_id, user_id) VALUES (?, ?)"
	_, err = db.Exec(q, sessionId, userId)
	if err != nil {
		sessionId = ""
		return
	}

	return
}

func FindUserBySessionId(db *sql.DB, sessionId string) (user User, err error) {
	sessionRow := db.QueryRow("SELECT * FROM webvuln_session WHERE session_id = ?", sessionId)
	session, err := ScanSession(sessionRow)
	if err != nil {
		return
	}

	userRow := db.QueryRow("SELECT id, username FROM webvuln_user WHERE id = ?", session.UserId)
	user, err = ScanUser(userRow)
	if err != nil {
		return
	}

	return
}

func ClearUserSession(db *sql.DB, userId int) (err error) {
	q := "DELETE FROM webvuln_session WHERE user_id = ?"
	_, err = db.Exec(q, userId)

	return
}

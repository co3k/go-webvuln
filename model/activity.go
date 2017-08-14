package model

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func CreateActivity(db *sql.DB, u User, body string) (err error) {
	q := "INSERT INTO webvuln_activity (user_id, body, created_at) VALUES (?, ?, ?)"
	_, err = db.Exec(q, u.Id, body, time.Now().Format("2006-01-02 03:04:05"))

	return
}

func CountActivities(db *sql.DB) (result int, err error) {
	q := "SELECT COUNT(*) FROM webvuln_activity"
	err = db.QueryRow(q).Scan(&result)
	if err != nil {
		return
	}

	return
}

func GetActivities(db *sql.DB, size int, offset int) (result []Activity, err error) {
	q := "SELECT a.id, u.id, u.username, a.body, a.created_at FROM webvuln_activity a LEFT JOIN webvuln_user u ON user_id = u.id ORDER BY created_at DESC LIMIT ?, ?"
	rows, err := db.Query(q, offset, size)
	if err != nil {
		return
	}

	result, err = ScanActivitys(rows)
	if err != nil {
		return
	}

	return
}

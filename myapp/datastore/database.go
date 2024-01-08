package datastore

import (
	"database/sql"
	"log"
	"myapp/models"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Assuming you've set up your database and have the connection details ready
var db *sql.DB

func init() {
	var err error
	// Connect to the database (replace with your connection details)
	db, err = sql.Open("postgres", "user=youruser dbname=yourdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

// CreateEntry inserts a new entry into the database
func CreateEntry(entry models.Entry) error {
	// SQL query to insert a new entry
	_, err := db.Exec("INSERT INTO entries (status, counteragent, sport, live_status, amount, market, selection, odd, notes, tournament, date_created, date_finish) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
		entry.Status, entry.Counteragent, entry.Sport, entry.LiveStatus, entry.Amount, entry.Market, entry.Selection, entry.ODD, entry.Notes, entry.Tournament, entry.DateCreated, entry.DateFinish)
	return err
}

// GetAllEntries retrieves all entries from the database
func GetAllEntries() ([]models.Entry, error) {
	rows, err := db.Query("SELECT id, status, counteragent, sport, live_status, amount, market, selection, odd, notes, tournament, date_created, date_finish FROM entries")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []models.Entry
	for rows.Next() {
		var e models.Entry
		if err := rows.Scan(&e.ID, &e.Status, &e.Counteragent, &e.Sport, &e.LiveStatus, &e.Amount, &e.Market, &e.Selection, &e.ODD, &e.Notes, &e.Tournament, &e.DateCreated, &e.DateFinish); err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	return entries, nil
}

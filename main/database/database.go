package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite3 driver
)

// Database struct
type Database struct {
	db *sql.DB
}

// NewDatabase creates a new Database instance
func NewDatabase() *Database {
	db, err := sql.Open("sqlite3", "chat.db")
	if err != nil {
		log.Fatal(err)
	}

	return &Database{db: db}
}

// Close closes the database connection
func (d *Database) Close() {
	if err := d.db.Close(); err != nil {
		log.Fatal(err)
	}
}

// CreateUsersTable creates the users table
func (d *Database) CreateUsersTable() {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			first_name TEXT NOT NULL,
			last_name TEXT NOT NULL,
			profile_picture TEXT,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		);
	`
	if _, err := d.db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

// CreateMessagesTable creates the messages table
func (d *Database) CreateMessagesTable() {
	query := `
		CREATE TABLE IF NOT EXISTS messages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			sender TEXT NOT NULL,
			room_id TEXT NOT NULL,
			content TEXT NOT NULL,
			type TEXT NOT NULL,
			time TEXT NOT NULL
		);
	`
	if _, err := d.db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

// InsertUser inserts a new user into the users table
func (d *Database) InsertUser(username, firstName, lastName, email, password string) {
	query := `
		INSERT INTO users (username, first_name, last_name, email, password)
		VALUES (?, ?, ?, ?, ?);
	`
	if _, err := d.db.Exec(query, username, firstName, lastName, email, password); err != nil {
		log.Fatal(err)
	}
}

// InsertMessage inserts a new message into the messages table
func (d *Database) InsertMessage(sender, roomID, content, messageType, time string) {
	query := `
		INSERT INTO messages (sender, room_id, content, type, time)
		VALUES (?, ?, ?, ?, ?);
	`
	if _, err := d.db.Exec(query, sender, roomID, content, messageType, time); err != nil {
		log.Fatal(err)
	}
}

// QueryUsers executes a query on the users table and returns the result
func (d *Database) QueryUsers(query string, args ...interface{}) *sql.Rows {
	rows, err := d.db.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

// QueryMessages executes a query on the messages table and returns the result
func (d *Database) QueryMessages(query string, args ...interface{}) *sql.Rows {
	rows, err := d.db.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

// UpdateUser updates a user in the users table
func (d *Database) UpdateUser(query string, args ...interface{}) {
	if _, err := d.db.Exec(query, args...); err != nil {
		log.Fatal(err)
	}
}

// UpdateMessage updates a message in the messages table
func (d *Database) UpdateMessage(query string, args ...interface{}) {
	if _, err := d.db.Exec(query, args...); err != nil {
		log.Fatal(err)
	}
}

// DeleteUser deletes a user from the users table
func (d *Database) DeleteUser(query string, args ...interface{}) {
	if _, err := d.db.Exec(query, args...); err != nil {
		log.Fatal(err)
	}
}

// DeleteMessage deletes a message from the messages table
func (d *Database) DeleteMessage(query string, args ...interface{}) {
	if _, err := d.db.Exec(query, args...); err != nil {
		log.Fatal(err)
	}
}

// DropUsersTable drops the users table
func (d *Database) DropUsersTable() {
	query := "DROP TABLE IF EXISTS users;"
	if _, err := d.db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

// DropMessagesTable drops the messages table
func (d *Database) DropMessagesTable() {
	query := "DROP TABLE IF EXISTS messages;"
	if _, err := d.db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

// Migrate migrates the database schema
func (d *Database) Migrate() {
	d.CreateUsersTable()
	d.CreateMessagesTable()
}

// Rollback rolls back the database schema
func (d *Database) Rollback() {
	d.DropUsersTable()
	d.DropMessagesTable()
}

// Seed seeds the database with initial data
func (d *Database) Seed() {
	d.InsertUser("admin", "Admin", "User", "admin@example.com", "password")
}

// Reset resets the database schema and seeds it with initial data
func (d *Database) Reset() {
	d.Rollback()
	d.Migrate()
	d.Seed()
}

// Truncate truncates the database tables
func (d *Database) Truncate() {
	d.DeleteUser("DELETE FROM users;")
	d.DeleteMessage("DELETE FROM messages;")
}

// Drop drops the database
func (d *Database) Drop() {
	if err := os.Remove("chat.db"); err != nil {
		log.Fatal(err)
	}
}

// GetDatabase returns the Database instance
func (d *Database) GetDatabase() *Database {
	return d
}

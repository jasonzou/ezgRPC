package server

import (
	"database/sql"
	"errors"
	"log"
	"time"

	api "github.com/jasonzou/ezgRPC/src/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"

	// needed for SQLite driver
	_ "github.com/mattn/go-sqlite3"
)

const create string = `
  CREATE TABLE IF NOT EXISTS entries (
  id INTEGER NOT NULL PRIMARY KEY,
  time DATETIME NOT NULL,
	prefix TEXT,
  suffix TEXT,
  title TEXT NOT NULL,
  url TEXT NOT NULL,
  config TEXT NOT NULL);
    `
const file string = "entries.db"

type Entries struct {
	db *sql.DB
}

func NewEntries() (*Entries, error) {
	log.Print("new entries")
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	log.Print("new entries 22")
	if _, err := db.Exec(create); err != nil {
		return nil, err
	}
	log.Print("new entries 223332")
	return &Entries{
		db: db,
	}, nil
}
func (c *Entries) Insert(entry *api.Entry) (int, error) {
	res, err := c.db.Exec("INSERT INTO entries VALUES(NULL,?,?,?,?,?,?);",
		entry.Time.AsTime(), entry.Prefix, entry.Suffix, entry.Title, entry.Url, entry.Config)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	log.Printf("Added %v as %d", entry, id)
	return int(id), nil
}

var ErrIDNotFound = errors.New("Id not found")

func (c *Entries) Retrieve(id int) (*api.Entry, error) {
	log.Printf("Getting %d", id)

	// Query DB row based on ID
	row := c.db.QueryRow("SELECT id, time, prefix, suffix, title, url, config FROM entries WHERE id=?", id)

	// Parse row into Interval struct
	entry := api.Entry{}
	var err error
	var time time.Time
	if err = row.Scan(&entry.Id, &time, &entry.Prefix, &entry.Suffix, &entry.Title, &entry.Url, &entry.Config); err == sql.ErrNoRows {
		log.Printf("Id not found")
		return &api.Entry{}, ErrIDNotFound
	}
	entry.Time = timestamppb.New(time)
	return &entry, err
}

func (c *Entries) List(offset int) ([]*api.Entry, error) {
	log.Printf("Getting list from offset %d\n", offset)

	// Query DB row based on ID
	rows, err := c.db.Query("SELECT * FROM entries WHERE ID > ? ORDER BY id DESC LIMIT 100", offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := []*api.Entry{}
	for rows.Next() {
		i := api.Entry{}
		var time time.Time
		err = rows.Scan(&i.Id, &time, &i.Prefix, &i.Suffix, &i.Title, &i.Url, &i.Config)
		if err != nil {
			return nil, err
		}
		i.Time = timestamppb.New(time)
		data = append(data, &i)
	}
	return data, nil
}

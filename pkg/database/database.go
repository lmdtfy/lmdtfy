package database

import (
	"errors"
	"log"
	"time"

	r "github.com/dancannon/gorethink"
)

const (
	testAddress = "localhost:28015"
	testDBName  = "test_lmdtfy"
)

var (
	// DB
	sess   *r.Session
	dbName string

	// Errors
	ErrNotFound   = errors.New("record not found")
	ErrValidation = errors.New("validation error")

	tables = []string{"users", "repos", "commits", "builds", "settings"}
)

// type BaseStore struct {
// 	//Table
// }
//
// func (s BaseStore) Table(name string) *Table {
// 	return &Table{r.Table(name)}
// }

type Store struct {
	// *BaseStore
}

func (s Store) Table(name string) *Table {
	return &Table{r.Table(name)}
}

// Repos stores all user repositories.
func (s Store) Repos() *Table {
	return s.Table("repos")
}

type Table struct {
	r.Term
}

func (t *Table) Run() {
	t.Term.Run(sess)
}
func (t *Table) RunWrite() (r.WriteResponse, error) {
	return t.Term.RunWrite(sess)
}

// Builds stores all builds for a an app.
func (s *Store) Builds() r.Term {
	return r.Table("builds")
}

// Commits stores all commits for a repo/app.
func (s *Store) Commits() r.Term {
	return r.Table("commits")
}

// Settings stores all app (lmdtfy) settings.
func (s *Store) Settings() r.Term {
	return r.Table("settings")
}

// Users stores all app users.
func (s *Store) Users() r.Term {
	return r.Table("users")
}

//
// type Term struct {
// 	r.Term
// }
//
// func (t *Term) Run() (*r.Cursor, error) {
// 	return t.Term.Run(sess)
// }

// Connect establishes connection with rethinkDB
func Connect(address, database string) error {

	dbName = database
	var err error
	sess, err = r.Connect(r.ConnectOpts{
		Address:     address,
		Database:    dbName,
		MaxIdle:     10,
		IdleTimeout: time.Second * 10,
	})
	if err != nil {
		return err
	}

	return nil
}

// SetupDB will be used to bootstrap the DB
func SetupDB(testData bool) {
	log.Println("SetupDB: Start")
	createDatabase()
	createTables()
	createIndexes()

	// if testData {
	// 	cleanTables()
	// 	insertTestData()
	// }

	log.Println("SetupDB: Done")
}

package store

import (
	"log"
	"time"

	r "github.com/dancannon/gorethink"
)

var (
	sess   *r.Session
	dbName string
)

// Store represents a connection to the DB and allows you to run queries.
type Store struct {
}

// RunWrite will run a query for the current session.
func (s Store) RunWrite(term r.Term) (r.WriteResponse, error) {
	writeRes, err := term.RunWrite(sess)
	if err != nil {
		log.Println("DB error: ", err, "; In:", term)
		return r.WriteResponse{}, err
	}

	log.Printf("Query: %s, Res: %+v", term, writeRes)
	return writeRes, nil
}

// All will run a query and return the results scanned into an interface.
func (s Store) Run(term r.Term) (*r.Cursor, error) {
	cursor, err := term.Run(sess)
	if err != nil {
		log.Println("DB error: ", err, "; In:", term)
		return nil, err
	}

	log.Printf("Query: %s", term)
	return cursor, nil
}

// All will run a query and return the results scanned into an interface.
func (s Store) All(i interface{}, term r.Term) error {
	cursor, err := s.Run(term)
	if err != nil {
		return err
	}

	err = cursor.All(i)
	if err != nil {
		log.Println("DB error: ", err, "; In:", term)
		return err
	}

	return nil
}

// Collection represents a table in rethinkdb.
type Collection struct {
	r.Term
}

// Repos stores all user repositories.
func Repos() Collection {
	return Collection{r.Table("repos")}
}

// Apps stores all apps for a uder.
func Apps() Collection {
	return Collection{r.Table("apps")}
}

// Builds stores all builds for a an app.
func Builds() Collection {
	return Collection{r.Table("builds")}
}

// Commits stores all commits for a repo/app.
func Commits() Collection {
	return Collection{r.Table("commits")}
}

// Settings stores all app (lmdtfy) settings.
func Settings() Collection {
	return Collection{r.Table("settings")}
}

// Users stores all app users.
func Users() Collection {
	return Collection{r.Table("users")}
}

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

package store

import (
	"time"

	r "github.com/dancannon/gorethink"
)

var (
	sess   *r.Session
	dbName string
)

type Store struct {
	Session      *r.Session
	DatabaseName string
}

// RunWrite will run a query for the current session.
func (s *Store) RunWrite(term r.Term) (r.WriteResponse, error) {
	return term.RunWrite(s.Session)
}

// Collection represents a table in rethinkdb.
type Collection struct {
	r.Term
}

// Repos stores all user repositories.
func Repos() Collection {
	return Collection{r.Table("repos")}
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

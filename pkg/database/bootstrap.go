package database

import r "github.com/dancannon/gorethink"

func createDatabase() {
	r.DbCreate(dbName).RunWrite(sess)
}

func createTables() {
	for _, name := range tables {
		r.Db(dbName).TableCreate(name).Run(sess)
	}
}

func cleanTables() {
	for _, name := range tables {
		r.Table(name).Delete().Run(sess)
	}
}

func createIndexes() {
	r.Db(dbName).Table("users").IndexCreate("email").Run(sess)
}

func insertTestData() {

}

func insertTestUsers() {
	//
	// u, _ := m.NewUser("test@test.com", "somePassword", "Admin")
	// Users.Store(u)
}

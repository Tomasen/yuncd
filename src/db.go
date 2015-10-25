package crane

import (
	"database/sql"
	"os"
	gorp  "gopkg.in/gorp.v1"
	log "github.com/Sirupsen/logrus"
)

var _dbmap = initDB()

func initDB() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("mysql", os.Getenv("SQL_DSN"))
	if err != nil {
		log.Fatal("sql.Open failed")
	}

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

	// add a table, setting the table name to 'posts'
	dbmap.AddTableWithName(Host{}, "hosts")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal("Create tables failed")
	}

	return dbmap
}

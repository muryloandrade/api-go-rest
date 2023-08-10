package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	DBONE *sql.DB
	DBTWO *sql.DB
)

func DbTwoConnect() {
	var (
		host1     = "172.20.120.180"
		port1     = 5432
		user1     = "root"
		password1 = "root"
		host2     = "172.20.120.180"
		port2     = 5433
		user2     = "root"
		password2 = "root"
		dbname    = "root"
	)

	rwPrimary := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host1, port1, user1, password1, dbname)
	readOnlyReplica := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host2, port2, user2, password2, dbname)

	// Open database for primary
	DBONE, err := sql.Open("postgres", rwPrimary)
	if err != nil {
		log.Print("go error when connecting to the primary DB")
	}

	// Open database for read-only replica
	DBTWO, err = sql.Open("postgres", readOnlyReplica)
	if err != nil {
		log.Print("go error when connecting to the read-only replica DB")
	}

	// Example: Query from read-only replica
	rows, err := DBTWO.Query("SELECT * FROM users")

	if err != nil {
		log.Print("error querying read-only replica:", err)
	} else {
		defer rows.Close()
		// Process rows
	}

	// Reusing the 'rows' variable for a query on the primary DB
	rows, err = DBONE.Query("SELECT * FROM personalidades")
	if err != nil {
		log.Print("error querying primary DB:", err)
	} else {
		defer rows.Close()
		// Process rows
	}
}

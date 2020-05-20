package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const DB_PATH string = "postgresql://lhincapie0@localhost:26257/restapi?sslmode=disable"

func AddDomain(database *sql.DB, host string, sslGrade string) {

	if _, err := database.Exec(
		"INSERT INTO domains (host, ssllab) VALUES ( '" + host + "', '" + sslGrade + "')"); err != nil {
		log.Fatal(err)
	}
}

func GetDomains(database *sql.DB) []string {
	var hosts []string
	var host string

	rows, err := database.Query("SELECT host FROM domains;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&host); err != nil {
			log.Fatal(err)
		}
		hosts = append(hosts, host)
	}
	return hosts

}

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", DB_PATH)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	fmt.Println("database started")

	return db

}

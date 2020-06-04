package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	ds "github.com/lhincapie0/go-restAPI/API/dataStructure"

	_ "github.com/lib/pq"
)

const DB_PATH string = "postgresql://lhincapie0@localhost:26257/restapi?sslmode=disable"

func AddDomain(database *sql.DB, host string, sslGrade string) {

	if _, err := database.Exec(
		"INSERT INTO domains (host, ssllab) VALUES ( '" + host + "', '" + sslGrade + "')"); err != nil {
		log.Fatal(err)
	}
}

func AddDomainInfo(database *sql.DB, host string, domain ds.DomainInfo) {
	b, _ := json.Marshal(domain.Servers)
	servers := string(b)
	fmt.Println(b)
	if _, err := database.Exec(
		"INSERT INTO domains (hostName,host,servers, ssllab,is_down) VALUES ( '" + host + "', '" + domain.Title + "', '" + servers + "', '" + domain.SslGrade + "' , " + strconv.FormatBool(domain.IsDown) + ")"); err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(servers), &servers)
	fmt.Println(servers)

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

func addColums(database *sql.DB) {

	_, err := database.Exec("ALTER TABLE domains ADD COLUMN hostName string;")
	if err != nil {
		fmt.Println(err)
	}

}

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", DB_PATH)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	fmt.Println("database started")
	addColums(db)

	return db

}

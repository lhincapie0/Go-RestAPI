package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	ds "github.com/lhincapie0/go-restAPI/API/dataStructure"

	_ "github.com/lib/pq"
)

const DB_PATH string = "postgresql://lhincapie0@localhost:26257/restapi?sslmode=disable"

func AddDomainInfo(database *sql.DB, host string, domain ds.DomainInfo) {
	b, _ := json.Marshal(domain.Servers)
	servers := string(b)

	//ERROR DETECTED FOR TITLES WITH 'S
	if strings.Contains(domain.Title, "'") {
		domain.Title = strings.Replace(domain.Title, "'", "", 1)
	}

	values := "'" + host + "', '" + servers + "','" + strconv.FormatBool(domain.SeversChanged) + "','" + domain.SslGrade + "','" + domain.PreviousSslGrade + "','" + domain.Logo + "','" + domain.Title + "','" + strconv.FormatBool(domain.IsDown) + "'"
	fmt.Println("VALUES")
	fmt.Println(values)

	if _, err := database.Exec(
		"INSERT INTO domains (host,servers, servers_changed, ssl_grade,previous_ssl_grade,logo,title,is_down) VALUES (" + values + ")"); err != nil {
		log.Fatal(err)
		fmt.Println("ERROR ERROR")
	}

	//TO SEND INFORMATION LATER
	//	json.Unmarshal([]byte(servers), &servers)

}

func GetDomains(database *sql.DB) []ds.DomainHistoryElement {
	//	var host string

	rows, err := database.Query("SELECT * FROM domains;")
	if err != nil {
		log.Fatal("err")
		fmt.Println("error 1")
		return nil
	}
	defer rows.Close()

	domains := make([]ds.DomainHistoryElement, 0)

	for rows.Next() {
		domain := ds.DomainHistoryElement{}

		err := rows.Scan(&domain.Host, &domain.Info.Servers, &domain.Info.SeversChanged, &domain.Info.SslGrade, &domain.Info.PreviousSslGrade, &domain.Info.Logo, &domain.Info.Title, &domain.Info.IsDown)
		if err != nil {
			log.Fatal(err)

			return nil
		}

		domains = append(domains, domain)

	}

	return domains

}

func deleteRows(database *sql.DB) {
	//	var host string

	_, err := database.Query("DELETE FROM domains;")
	if err != nil {
		log.Fatal("err")
		fmt.Println("error 1")
	}
	fmt.Println("ELIMINA")
}

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", DB_PATH)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	fmt.Println("database started")
	//deleteRows(db)

	return db

}

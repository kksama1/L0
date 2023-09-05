package postgre

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type pgInfo struct {
	host     string
	port     int
	database string
	username string
	password string
}

// createConnection performs connection to db.
func createConnection() *sql.DB {

	ci := pgInfo{host: "localhost", port: 5432, database: "testing", username: "admin_pg", password: "root"}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		ci.host, ci.port, ci.username, ci.password, ci.database)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

// SendData inserting data to the "data" table.
func SendData(order_uid, orders_f string) {
	db := createConnection()
	defer db.Close()
	db.QueryRow("INSERT INTO data (order_uid, orders_f) values ($1, $2)", order_uid, orders_f)

}

// GetData gets data from "data" table and returns it as map.
func GetData(cashedData map[string]string) {
	var (
		orderUid string
		OrdersF  string
	)
	db := createConnection()
	defer db.Close()
	rows, err := db.Query(
		"SELECT * FROM data;")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&orderUid, &OrdersF)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		cashedData[orderUid] = OrdersF

	}
	defer rows.Close()
}

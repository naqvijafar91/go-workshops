package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3" // go driver for sql lite
)

// SQL Lite is used in browser caches , building demo apps ,IPhone call history etc
//sqllite is installed by default in Linux or else run in terminal
//sudo apt-get install sqlite3 libsqlite3-dev
//sqlite3

func main() {
	driverName := "sqlite3"
	dataSourceName := "./mydatabase.db"
	database, _ := sql.Open(driverName, dataSourceName)
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("DELETE FROM people")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	var id int
	var firstname string
	var lastname string
	loadDummyData(statement, firstname, lastname)
	rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}
}

//Helper function to insert 10 random names
func loadDummyData(statement *sql.Stmt, firstname, lastname string) {
	// using math.rand package
	// Need different seed on each run, most used seed is the current time, converted to int64 by UnixNano
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		firstname = randomString(5)
		lastname = randomString(5)
		fmt.Printf("Inserting FirstName %s Lastname %s \n", firstname, lastname)
		statement.Exec(firstname, lastname)
	}

}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90)) // 65-90 are uppercase letters
	}
	return string(bytes)
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min) //rand.Intn is core function to get a random int in [0,n]
}

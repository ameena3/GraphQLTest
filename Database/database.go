package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	server   = "0.0.0.0"
	port     = 1433
	user     = "sa"
	password = "Anubh@v0162"
	database = "master"
)

//ConnectToDb ... Connects to the database for you.
func (d *Data) ConnectToDb() error {
	// Connect to database
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	fmt.Printf("Connected!\n")
	d.db = conn
	return err

}

// GetUsers ... gets the suers from the database
func (d *Data) GetUsers() (int, []User, error) {
	defer d.db.Close()
	users := []User{}
	tsql := fmt.Sprintf("SELECT Id, First_name FROM Users;")
	rows, err := d.db.Query(tsql)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return -1, nil, err
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		user := User{}
		var id int
		var firstname string
		err := rows.Scan(&id, &firstname)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return -1, nil, err
		}
		user.ID = id
		user.FirstName = firstname
		users = append(users, user)
		fmt.Printf("ID: %d, Name: %s \n", id, firstname)
		count++
	}
	return count, users, err
}

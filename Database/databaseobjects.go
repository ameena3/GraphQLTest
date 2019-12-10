package database

import "database/sql"

//Data ... holds the database context.
type Data struct {
	db *sql.DB
}

//User ... holds the user object
type User struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	APIKey    string
	CreatedAt string
}

//For running test docker container
//docker run -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=Anubh@v0162' -p 1433:1433 -d mcr.microsoft.com/mssql/server

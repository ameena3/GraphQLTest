package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	server   = "192.168.4.23"
	instance = "FNMS_PROD01"
	port     = 1433
	user     = "dbameena"
	database = "FNMP"
)

//ConnectToDb ... Connects to the database for you.
func (d *Data) ConnectToDb(password string) error {
	// Connect to database
	// connString := fmt.Sprintf("server=%s;instance=%s;user id=%s;password=%s;port=%d;database=%s;",
	// 	server, instance, user, password, port, database)
	uristr := &url.URL{
		Scheme: "sqlserver",
		Host:   server,
		Path:   instance,
		User:   url.UserPassword(user, password),
	}
	dbconnstr := uristr.String()
	conn, err := sql.Open("mssql", dbconnstr)
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

func (d *Data) GetComplianceComputerByComplianceComputerID(ComplianceComputerID int) (int, ComplianceComputer, error) {
	defer d.db.Close()
	cc := ComplianceComputer{}
	// // tsql1 := fmt.Sprintf("")
	// // contextrows, err := d.db.Query(tsql1)
	// contextrows.Close()
	err := d.db.Ping()
	checkerr(err)
	fmt.Println("Ping successfull")
	tsql := fmt.Sprintf("USE FNMP ; SELECT ComplianceComputerID, ComputerName, AssetID, InventoryAgent FROM ComplianceComputer_MT WHERE ComplianceComputerID = %d AND TenantID = 11", ComplianceComputerID)
	fmt.Println("The following query is being run : " + tsql)
	rows, err := d.db.Query(tsql)
	checkerr(err)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return -1, cc, err
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		var ComplianceComputerID int
		var AssetID sql.NullInt32
		var ComputerName string
		var InventoryAgent sql.NullString
		err := rows.Scan(&ComplianceComputerID, &ComputerName, &AssetID, &InventoryAgent)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return -1, cc, err
		}
		cc.ComplianceComputerID = ComplianceComputerID
		cc.ComputerName = ComputerName
		cc.AssetID = AssetID
		cc.InventoryAgent = InventoryAgent
		count++
	}
	return count, cc, err

}

func (d *Data) GetListOfComplianceComputer() (int, []ComplianceComputer, error) {
	defer d.db.Close()
	cca := []ComplianceComputer{}
	// // tsql1 := fmt.Sprintf("")
	// // contextrows, err := d.db.Query(tsql1)
	// contextrows.Close()
	err := d.db.Ping()
	checkerr(err)
	fmt.Println("Ping successfull")
	tsql := fmt.Sprintf("USE FNMP ; SELECT ComplianceComputerID, ComputerName, AssetID, InventoryAgent FROM ComplianceComputer_MT WHERE TenantID = 11")
	fmt.Println("The following query is being run : " + tsql)
	rows, err := d.db.Query(tsql)
	checkerr(err)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return -1, cca, err
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		cc := ComplianceComputer{}
		var ComplianceComputerID int
		var AssetID sql.NullInt32
		var ComputerName string
		var InventoryAgent sql.NullString
		err := rows.Scan(&ComplianceComputerID, &ComputerName, &AssetID, &InventoryAgent)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return -1, cca, err
		}
		cc.ComplianceComputerID = ComplianceComputerID
		cc.ComputerName = ComputerName
		cc.AssetID = AssetID
		cc.InventoryAgent = InventoryAgent
		cca = append(cca, cc)
		count++
	}
	return count, cca, err

}
func checkerr(err error) {
	if err != nil {
		log.Printf(err.Error())
	}
}

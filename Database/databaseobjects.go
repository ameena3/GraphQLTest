package database

import (
	"database/sql"
	"time"
)

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

//ComplianceComputer ... database object from the compliance database.
type ComplianceComputer struct {
	ComplianceComputerID                    int
	ComplianceComputerTypeID                int
	IsComplianceComputerTypeIDFromInventory bool
	ComputerName                            string
	ComplianceDomainID                      int
	ComplianceComputerStatusID              int
	ComplianceComputerRoleID                int
	ComplianceComputerInventorySourceTypeID int
	AssetID                                 sql.NullInt32
	OperatingSystem                         string
	ServicePack                             string
	NumberOfProcessors                      int
	NumberOfProcessorsDefault               int
	ProcessorType                           string
	ProcessorTypeDefault                    string
	MaxClockSpeed                           int
	MaxClockSpeedDefault                    int
	TotalMemory                             int64
	ChassisTypeID                           int
	AssignedChassisTypeID                   int
	NumberOfHardDrives                      int
	TotalDiskSpace                          int64
	NumberOfNetworkCards                    int
	NumberOfDisplayAdapters                 int
	IPAddress                               string
	MACAddress                              string
	Manufacturer                            string
	ModelNo                                 string
	ModelNoDefault                          string
	SerialNo                                string
	ComplianceUserID                        int
	AssignedUserID                          int
	CalculatedUserID                        int
	LocationID                              string
	BusinessUnitID                          string
	CostCenterID                            string
	CategoryID                              string
	InventoryDate                           time.Time
	HardwareInventoryDate                   time.Time
	ServicesInventoryDate                   time.Time
	UpdatedUser                             string
	UpdatedDate                             time.Time
	CreationUser                            string
	CreationDate                            time.Time
	InventoryAgent                          sql.NullString
	NumberOfCores                           int
	NumberOfCoresDefault                    int
	NumberOfSockets                         int
	NumberOfSocketsDefault                  int
	AssetComplianceStatusID                 int
	PartialNumberOfProcessors               int64
	PartialNumberOfProcessorsDefault        int64
	UntrustedSerialNo                       bool
	ILMTAgentID                             int64
	FNMPComputerUID                         string
	UUID                                    string
	HostIdentifyingNumber                   string
	HostType                                string
	NumberOfLogicalProcessors               int
	NumberOfLogicalProcessorsDefault        int
	PrimaryComplianceUserID                 int
	MDScheduleGeneratedDate                 time.Time
	MDScheduleContainsPVUScan               bool
	HostID                                  string
	FirmwareSerialNumber                    string
	MachineID                               string
	CloudServiceProviderID                  int
}

//For running test docker container
//docker run -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=Anubh@v0162' -p 1433:1433 -d mcr.microsoft.com/mssql/server

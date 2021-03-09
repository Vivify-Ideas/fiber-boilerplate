package database

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/vivify-ideas/fiber_boilerplate/config"
	"github.com/vivify-ideas/fiber_boilerplate/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// DriverType - database driver type
type DriverType string

const (
	sqliteType    DriverType = "sqlite"
	mysqlType     DriverType = "mysql"
	postgresType  DriverType = "postgres"
	sqlserverType DriverType = "sqlserver"
)

// IsValid - check if database driver is valid
func (dt DriverType) IsValid() error {
	switch dt {
	case sqliteType, mysqlType, postgresType, sqlserverType:
		return nil
	}
	return errors.New("Invalid Database Driver Type")
}

// DBConfig definition
type DBConfig struct {
	Driver   string
	Host     string
	Username string
	Password string
	Port     int
	Database string
}

// Database struct
type Database struct {
	*gorm.DB
}

var db *Database

// Init -> init databse or return the instance
func Init() *Database {
	if db != nil {
		return db
	}

	log.Println("Database instantiated")

	config := config.App.Env
	port, _ := strconv.ParseInt(config["DB_PORT"], 0, 32)

	driverType := DriverType(config["DB_DRIVER"])
	if err := driverType.IsValid(); err != nil {
		log.Fatal(err)
	}

	var err error
	db, err = GetDB(&DBConfig{
		Driver:   string(driverType),
		Host:     config["DB_HOST"],
		Username: config["DB_USERNAME"],
		Password: config["DB_PASSWORD"],
		Port:     int(port),
		Database: config["DB_DATABASE"],
	})

	// Auto-migrate database models
	if err != nil {
		log.Println("failed to connect to database:", err.Error())
	} else {
		if db == nil {
			log.Println("failed to connect to database: db variable is nil")
		} else {
			err = db.AutoMigrate(&models.User{}, &models.File{})
			if err != nil {
				log.Println("failed to automigrate user model:", err.Error())
				return nil
			}
			// TODO: automigrate???
		}
	}

	return db
}

// GetDB - initialize database
func GetDB(dbconfig *DBConfig) (*Database, error) {
	var db *gorm.DB
	var err error

	switch strings.ToLower(dbconfig.Driver) {
	case string(mysqlType):
		dsn := dbconfig.Username + ":" + dbconfig.Password + "@tcp(" + dbconfig.Host + ":" + strconv.Itoa(dbconfig.Port) + ")/" + dbconfig.Database + "?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=UTC"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		break
	case "postgresql", "postgres":
		dsn := "user=" + dbconfig.Username + " password=" + dbconfig.Password + " dbname=" + dbconfig.Database + " host=" + dbconfig.Host + " port=" + strconv.Itoa(dbconfig.Port) + " TimeZone=UTC"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		break
	case "sqlserver", "mssql":
		dsn := "sqlserver://" + dbconfig.Username + ":" + dbconfig.Password + "@" + dbconfig.Host + ":" + strconv.Itoa(dbconfig.Port) + "?database=" + dbconfig.Database
		db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		break
	case string(sqliteType):
		db, err = gorm.Open(sqlite.Open(dbconfig.Host), &gorm.Config{})
		break
	}
	return &Database{db}, err
}

package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BaseDBService struct {
	DB *gorm.DB
}

func (ds *BaseDBService) Connect(host string, user string, password string, dbname string, port string) error {
	var err error

	// TODO: sslmode
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
	Dbdriver := postgres.Open(dsn)

	ds.DB, err = gorm.Open(Dbdriver, &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		return fmt.Errorf("Error: %s", err.Error())
	}

	fmt.Printf("Connected to the %s database", Dbdriver)
	return nil
}

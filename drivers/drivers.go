package drivers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectCloudSQLPostgres() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "cloudsqlpostgres",
		DSN:        "host=project:region:instance user=postgres dbname=postgres password=password sslmode=disable",
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

	}
	return db, nil
}

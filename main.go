package main

import (
	"fmt"

	"github.com/zaynalabuev/gorm-gin/config"
	"github.com/zaynalabuev/gorm-gin/models"
	"github.com/zaynalabuev/gorm-gin/routers"
)

func main() {

	config.ConnecctPostgres()
	fmt.Println("подключено", config.DB != nil)
	config.DB.AutoMigrate(&models.Student{}, &models.Group{})
	routers.Routers()
}

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zaynalabuev/gorm-gin/handlers"
)

func Routers() {
	r := gin.Default()
	r.GET("/ping", handlers.PingPong)
	r.POST("/students", handlers.CreateStudent)
	r.GET("/students/:id", handlers.GetByID)
	r.PATCH("/students/:id", handlers.UpgradeStudentById)
	r.DELETE("/students/:id", handlers.DeleteById)
	r.GET("/students", handlers.GetAllStudents)
	r.POST("/groups", handlers.CreateGroup)
	r.GET("/groups/:id", handlers.GetGroupsByID)
	r.PATCH("/groups/:id", handlers.UpgradeGroupsById)
	r.DELETE("/groups/:id", handlers.DeleteGroupsById)
	r.Run()
}

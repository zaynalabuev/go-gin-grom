package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zaynalabuev/gorm-gin/config"
	"github.com/zaynalabuev/gorm-gin/models"
)

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": "ошибка"})
		return
	}
	if err := config.DB.Create(&student).Error; err != nil {
		c.JSON(500, gin.H{"error": "error"})
		return
	}
	c.JSON(201, student)
}

func Pingpong(c *gin.Context) {

	c.JSON(200, gin.H{"message": "pong"})
}
func GetByID(c *gin.Context) {
	var students models.Student
	id := c.Param("id")
	if err := config.DB.First(&students, id).Error; err != nil {
		c.JSON(400, gin.H{"error": "error"})
		return
	}
	c.JSON(200, students)
}
func UpgradeStudentById(c *gin.Context) {
	var students models.Student
	id := c.Param("id")
	if err := config.DB.First(&students, id).Error; err != nil {
		c.JSON(400, &students)
		return
	}
	var l models.Student
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(500, gin.H{"error": "error"})
		return
	}
	config.DB.Model(&students).Updates(&l)
	c.JSON(200, students)

}
func DeleteById(c *gin.Context) {
	var students models.Student
	id := c.Param("id")
	if err := config.DB.Delete(&students, id).Error; err != nil {
		c.JSON(404, err)
		return
	}
	c.JSON(200, gin.H{"message": students})
}
func GetAllStudents(c *gin.Context) {
	var students []models.Student
	if err := config.DB.Find(&students).Error; err != nil {
		c.JSON(404, err)
		return
	}
	c.JSON(200, gin.H{"students": students})
}

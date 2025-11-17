package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zaynalabuev/gorm-gin/config"
	"github.com/zaynalabuev/gorm-gin/models"
)

func CreateGroup(c *gin.Context) {
	var groups models.Group
	if err := c.ShouldBindJSON(&groups); err != nil {
		c.JSON(400, gin.H{"error": "error"})
		return
	}
	if err := config.DB.Create(&groups).Error; err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, groups)

}
func GetGroupsByID(c *gin.Context) {
	var groups models.Group
	id := c.Param("id")
	if err := config.DB.First(&groups, id).Error; err != nil {
		c.JSON(400, gin.H{"error": "error"})
		return
	}
	c.JSON(200, groups)
}
func UpgradeGroupsById(c *gin.Context) {
	var groups models.Group
	id := c.Param("id")
	if err := config.DB.First(&groups, id).Error; err != nil {
		c.JSON(400, &groups)
		return
	}
	var l models.Group
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(500, gin.H{"error": "error"})
		return
	}
	config.DB.Model(&groups).Updates(&l)
	c.JSON(200, groups)

}
func DeleteGroupsById(c *gin.Context) {
	var groups models.Group
	id := c.Param("id")
	if err := config.DB.Delete(&groups, id).Error; err != nil {
		c.JSON(404, err)
		return
	}
	c.JSON(200, gin.H{"message": groups})
}
func GetAllGroups(c *gin.Context) {
	var groups []models.Group
	if err := config.DB.Find(&groups).Error; err != nil {
		c.JSON(404, err)
		return
	}
	c.JSON(200, gin.H{"students": groups})
}

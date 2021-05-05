package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sam-samid/go-refview-activities-module/models"
	"gorm.io/gorm"
)

type ActivityInput struct {
	id   int    `json:"id"`
	Name string `json:"name"`
	Abbr string `json:"abbr"`
}

func FindAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var activities []models.Activity
	db.Find(&activities)
	// Check if returns RecordNotFound error
	err := db.Find(&activities).Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": errors.Is(err, gorm.ErrRecordNotFound),
		})
		c.Abort()
		return
	}

	if activities == nil {
		c.JSON(404, gin.H{
			"message": "Activities Not Found",
			"data":    activities,
		})
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{
			"status": "Activities Found",
			"data":   activities,
		})
	}
	c.Abort()
}

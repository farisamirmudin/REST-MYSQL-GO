package controllers

import (
	"github.com/farisamirmudin/go-crud/initializers"
	"github.com/farisamirmudin/go-crud/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// Get data from url
	var body struct {
		Email   string
		Bio     string
		Country string
	}
	c.Bind(&body)
	// Create a post
	user := models.User{Email: body.Email, Bio: body.Bio, Country: body.Country}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Send it
	c.JSON(200, gin.H{
		"user": user,
	})
}

func GetUsers(c *gin.Context) {
	// Get all users
	var users []models.User
	result := initializers.DB.Find(&users)
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Send it
	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetUser(c *gin.Context) {
	// get id from url
	id := c.Param("id")
	// get user based on id
	var user models.User
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		c.Status(400)
		return
	}

	// send it
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UpdateUser(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// get data from body
	var body struct {
		Email   string
		Bio     string
		Country string
	}
	c.Bind(&body)

	// get user based on id
	var user models.User
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		c.Status(400)
		return
	}

	// update it
	initializers.DB.Model(&user).Updates(models.User{Email: body.Email, Bio: body.Bio, Country: body.Country})

	// send it
	c.JSON(200, gin.H{
		"user": user,
	})
}
func DeleteUser(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// delete user based on id
	initializers.DB.Delete(&models.User{}, id)

	// send it
	c.Status(200)
}

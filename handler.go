package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func getUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&user)
	c.JSON(http.StatusOK, user)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func createDevice(c *gin.Context) {
	var device Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&device)
	c.JSON(http.StatusCreated, device)
}

func getDevices(c *gin.Context) {
	var devices []Device
	db.Find(&devices)
	c.JSON(http.StatusOK, devices)
}

func getDevice(c *gin.Context) {
	id := c.Param("id")
	var device Device
	if err := db.First(&device, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}
	c.JSON(http.StatusOK, device)
}

func updateDevice(c *gin.Context) {
	id := c.Param("id")
	var device Device
	if err := db.First(&device, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&device)
	c.JSON(http.StatusOK, device)
}

func deleteDevice(c *gin.Context) {
	id := c.Param("id")
	var device Device
	if err := db.First(&device, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}
	db.Delete(&device)
	c.JSON(http.StatusOK, gin.H{"message": "Device deleted"})
}

func createAuthorization(c *gin.Context) {
	var auth Authorization
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&auth)
	c.JSON(http.StatusCreated, auth)
}

func getAuthorizations(c *gin.Context) {
	var auths []Authorization
	db.Preload("User").Preload("Device").Find(&auths)
	c.JSON(http.StatusOK, auths)
}

func getAuthorization(c *gin.Context) {
	id := c.Param("id")
	var auth Authorization
	if err := db.Preload("User").Preload("Device").First(&auth, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Authorization not found"})
		return
	}
	c.JSON(http.StatusOK, auth)
}

func deleteAuthorization(c *gin.Context) {
	id := c.Param("id")
	var auth Authorization
	if err := db.First(&auth, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Authorization not found"})
		return
	}
	db.Delete(&auth)
	c.JSON(http.StatusOK, gin.H{"message": "Authorization deleted"})
}

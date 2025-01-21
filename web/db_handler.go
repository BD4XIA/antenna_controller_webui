package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
	var users []User
	DB.Find(&users)

	for i := range len(users) {
		users[i].Password = ""
	}

	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func CreateDevice(c *gin.Context) {
	var device Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&device)
	c.JSON(http.StatusCreated, device)
}

func GetDevices(c *gin.Context) {
	var devices []Device
	DB.Find(&devices)
	c.JSON(http.StatusOK, devices)
}

func GetDevice(c *gin.Context) {
	id := c.Param("id")
	var device Device
	if err := DB.First(&device, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}
	c.JSON(http.StatusOK, device)
}

func UpdateDevice(c *gin.Context) {
	id := c.Param("id")
	var device Device
	if err := DB.First(&device, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Save(&device)
	c.JSON(http.StatusOK, device)
}

func DeleteDevice(c *gin.Context) {
	id := c.Param("id")
	var device Device
	if err := DB.First(&device, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}
	DB.Delete(&device)
	c.JSON(http.StatusOK, gin.H{"message": "Device deleted"})
}

func CreateAuthorization(c *gin.Context) {
	var auth Authorization
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&auth)
	c.JSON(http.StatusCreated, auth)
}

func GetAuthorizations(c *gin.Context) {
	var auths []Authorization
	DB.Preload("User").Preload("Device").Find(&auths)
	c.JSON(http.StatusOK, auths)
}

func GetAuthorization(c *gin.Context) {
	id := c.Param("id")
	var auth Authorization
	if err := DB.Preload("User").Preload("Device").First(&auth, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Authorization not found"})
		return
	}
	c.JSON(http.StatusOK, auth)
}

func DeleteAuthorization(c *gin.Context) {
	id := c.Param("id")
	var auth Authorization
	if err := DB.First(&auth, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Authorization not found"})
		return
	}
	DB.Delete(&auth)
	c.JSON(http.StatusOK, gin.H{"message": "Authorization deleted"})
}

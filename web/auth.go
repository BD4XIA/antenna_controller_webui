package web

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var realm = "Basic realm=" + strconv.Quote("Authorization Required")

func AdminAuth(c *gin.Context) {
	username, passwd, ok := c.Request.BasicAuth()
	c.Set("username", username)

	if !ok {
		c.Header("WWW-Authenticate", realm)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	h := md5.Sum([]byte(passwd))
	passwd = hex.EncodeToString(h[:])
	log.Printf("%s", passwd)
	h = md5.Sum([]byte(passwd + "x"))
	passwd = hex.EncodeToString(h[:])

	admin := User{UserName: username}
	result := DB.First(&admin)
	if result.Error != nil || admin.Password != passwd || admin.UserType != AdminUser {
		c.Header("WWW-Authenticate", realm)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func UserAuth(c *gin.Context) {
	username, passwd, ok := c.Request.BasicAuth()
	c.Set("username", username)

	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	h := md5.Sum([]byte(passwd))
	passwd = hex.EncodeToString(h[:])
	log.Printf("%s", passwd)
	h = md5.Sum([]byte(passwd + "x"))
	passwd = hex.EncodeToString(h[:])

	user := User{UserName: username}
	result := DB.First(&user)
	if result.Error != nil || user.Password != passwd {
		c.Header("WWW-Authenticate", realm)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

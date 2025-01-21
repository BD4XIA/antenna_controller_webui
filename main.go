package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BD4XIA/antenna_controller_webui/web"
	"github.com/gin-gonic/gin"
)

// // go:embed web
// var webFS embed.FS

func init() {
	web.DBInit("data.db")

	admin := web.User{
		UserType: web.AdminUser,
	}
	result := web.DB.First(&admin)
	if result.Error == nil {
		return
	}

	passwd := "admin"
	h := md5.Sum([]byte(passwd))
	passwd = hex.EncodeToString(h[:])
	h = md5.Sum([]byte(passwd + "x"))
	passwd = hex.EncodeToString(h[:])

	admin = web.User{
		UserName: "admin",
		UserType: web.AdminUser,
		Password: passwd,
	}
	result = web.DB.Create(&admin)
	if result.Error != nil {
		panic(result.Error)
	}
}

func main() {
	r := gin.Default()

	userRoutes := r.Group("/users", web.AdminAuth)
	{
		userRoutes.POST("/", web.CreateUser)
		userRoutes.GET("/", web.GetUsers)
		userRoutes.GET("/:id", web.GetUser)
		userRoutes.PUT("/:id", web.UpdateUser)
		userRoutes.DELETE("/:id", web.DeleteUser)
	}

	authRoutes := r.Group("/authorizations", web.AdminAuth)
	{
		authRoutes.POST("/", web.CreateAuthorization)
		authRoutes.GET("/", web.GetAuthorizations)
		authRoutes.GET("/:id", web.GetAuthorization)
		authRoutes.DELETE("/:id", web.DeleteAuthorization)
	}

	deviceRoutes := r.Group("/devices", web.UserAuth)
	{
		deviceRoutes.POST("/", web.CreateDevice)
		deviceRoutes.GET("/", web.GetDevices)
		deviceRoutes.GET("/:id", web.GetDevice)
		deviceRoutes.PUT("/:id", web.UpdateDevice)
		deviceRoutes.DELETE("/:id", web.DeleteDevice)
	}

	r.Static("/css", "./static/css")
	r.Static("/fonts", "./static/fonts")
	r.Static("/icons", "./static/icons")
	r.Static("/js", "./static/js")

	r.StaticFile("/", "./static/index.html")
	r.StaticFile("/device", "./static/device.html")
	r.StaticFile("/user", "./static/user.html")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("start")
	go server.ListenAndServe()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-sig
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	server.Shutdown(ctx)
	log.Println("shutdown")
}

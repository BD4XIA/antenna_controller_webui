package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// // go:embed web
// var webFS embed.FS

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Connected to database")

	err = db.AutoMigrate(&User{}, &Device{}, &Authorization{})
	if err != nil {
		log.Fatal("Failed to migrate tables:", err)
	}

	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", createUser)
		userRoutes.GET("/", getUsers)
		userRoutes.GET("/:id", getUser)
		userRoutes.PUT("/:id", updateUser)
		userRoutes.DELETE("/:id", deleteUser)
	}

	deviceRoutes := r.Group("/devices")
	{
		deviceRoutes.POST("/", createDevice)
		deviceRoutes.GET("/", getDevices)
		deviceRoutes.GET("/:id", getDevice)
		deviceRoutes.PUT("/:id", updateDevice)
		deviceRoutes.DELETE("/:id", deleteDevice)
	}

	authRoutes := r.Group("/authorizations")
	{
		authRoutes.POST("/", createAuthorization)
		authRoutes.GET("/", getAuthorizations)
		authRoutes.GET("/:id", getAuthorization)
		authRoutes.DELETE("/:id", deleteAuthorization)
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

	fmt.Println("start")
	go server.ListenAndServe()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-sig
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	server.Shutdown(ctx)
	fmt.Println("shutdown")

	// sw := NewSwitch("192.168.1.150", 80)

	// err := sw.EN(11)
	// if err != nil {
	// 	panic(err)
	// }

	// status, err := sw.Query()
	// if err != nil {
	// 	panic(err)
	// }

	// if status == nil {
	// 	return
	// }

	// data := status[0]
	// var m map[string]string
	// err = json.Unmarshal(data, &m)
	// if err != nil {
	// 	panic(err)
	// }

	// for k, v := range m {
	// 	fmt.Printf("%s %s\n", k, v)
	// }

	// rt := internal.NewRotator("192.168.1.130", 80)
	// angles, err := rt.Query()
	// if err != nil {
	// 	panic(err)
	// }
	// for i, angle := range angles {
	// 	fmt.Printf("%d %s\n", i+1, angle)
	// }
	// setup, err := rt.Status()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s %s\n", setup[0].G1000, setup[1].Stu)
	// err = rt.EN(2)
	// if err != nil {
	// 	panic(err)
	// }
}

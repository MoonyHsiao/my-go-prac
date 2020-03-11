package main

import (
	"fmt"

	"github.com/MoonyHsiao/my-go-prac/router/middleware"

	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("err")

	Port := ":18086"
	// Port := ":8080"
	route := gin.Default()

	fmt.Printf(Port)

	v1 := route.Group("/api")

	v1.Use(middleware.TimeoutMiddleware(time.Second * 2))
	{
		v1.GET("/short", middleware.TimedHandler(time.Second))
		v1.GET("/long", middleware.TimedHandler(time.Second*5))

	}

	route.Run(Port)
}

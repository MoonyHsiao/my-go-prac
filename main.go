package main

import (
	"github.com/MoonyHsiao/my-go-prac/endpoints/alexa"
	"github.com/gin-gonic/gin"
)

func main() {
	Port := ":18086"
	router := gin.Default()
	v1 := router.Group("/clawer")
	{
		v1.GET("", alexa.GetTop)
		v1.GET("/country", alexa.GetCountry)
	}
	router.Run(Port)
}

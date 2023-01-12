package main

import (
	"fmt"
	"log"
//	"time"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Print("get ping request")
		select {
		case <-c.Request.Context().Done():
			log.Print("request cancel")
			return
		default:
		}

		//time.Sleep(10*time.Second)
		log.Print("response pong")

		c.JSON(200,gin.H{
			"message": "pong",
			"code": 0,
		})
	}

}

func main() {
	r := gin.Default()

	// return some thing
	r.GET("/ping", ping())

	r.GET("/user/:name/age", func(c *gin.Context) {
		name := c.Param("name")
		fmt.Println("username",name)

		c.String(http.StatusOK, "Hello %s", name)
	})

	r.Run(":8082") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

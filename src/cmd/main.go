package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line_notify/src/client/line"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
}

func handler(ctx *gin.Context) {
	line.SendMessage("test message")
	ctx.JSON(200, "Hello world")
}

func main() {
	port := os.Getenv("PORT")

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		handler(ctx)
	})

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting server at Port %s", port)
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
}

func handler(ctx *gin.Context) {
	token := os.Getenv("LINE_NOTIFY_TOKEN")
	fmt.Println(token)
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

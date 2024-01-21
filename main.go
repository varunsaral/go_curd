package main

import (
	"go_curd/initializers"
	"os"
	"github.com/gin-gonic/gin"

)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDb()
		
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, varun !")
	})
	PORT := os.Getenv("PORT")
	r.Run("localhost:"+PORT) // listen and serve on 0.0.0.0:8080
}

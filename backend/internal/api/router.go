package api

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/users", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "Hello World",
				})
			})
		}
	}

	return r
}

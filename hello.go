package main

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"nome":  "Osmanito Torres",
		"email": "osmanitobrito@gmail.com",
		"site":  "www.osmanito.com",
		"fone":  "(63)991127568",
		"zap":   "(63)991127568",
	})
}

func ShowNome(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"Nome: ": nome,
	})
}

func main() {
	router := gin.Default()
	router.GET("/", Index)
	router.GET("/:nome", ShowNome)
	router.Run()
}

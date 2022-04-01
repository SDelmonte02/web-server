package main

import "github.com/gin-gonic/gin"

func main() {

	// Creo el router con gin
	router := gin.Default()

	// Captura solicitud GET "hola"
	router.GET("/hola", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola mundo desde Gin!",
		})
	})

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run(":8080")

}

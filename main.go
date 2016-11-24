package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"servidor/funciones"
)


func main (){
	router := gin.Default()

	router.GET("/users/", funciones.GetClientes)
	router.GET("/users/:id", funciones.GetCliente)
	router.GET("/welcome", funciones.Saludar)
	router.POST("/users/add", funciones.AddCliente)
	router.Run(":8080")
}




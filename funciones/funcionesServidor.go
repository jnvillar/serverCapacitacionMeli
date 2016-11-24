package funciones

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"servidor/model"

	"servidor/domain"
)

func Saludar(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func GetClientes(c *gin.Context) {
	clientes := model.GetClientes()
	c.JSON(200,  clientes)
}

func GetCliente(c *gin.Context) {
	cliente := model.GetCliente(c.Params.ByName("id"))
	c.JSON(200,cliente)
}

func AddCliente(c *gin.Context) {
	cliente := domain.Cliente{}
	c.Bind(&cliente)
	model.AddCliente(cliente)
	c.JSON(200,cliente)
}
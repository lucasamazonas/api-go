package routes

import (
	"api-go/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()

	pacientes := router.Group("/pacientes")
	{
		pacientes.GET("/", controllers.Listar)
		pacientes.GET("/:id", controllers.ListarPaciente)
		pacientes.POST("/", controllers.Criar)
		pacientes.PUT("/:id", controllers.Editar)
		pacientes.DELETE("/:id", controllers.Excluir)
	}

	router.Run(":8888")
}

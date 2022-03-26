package controllers

import (
	database "api-go/db"
	"api-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Listar(c *gin.Context) {
	var pacientes []models.Paciente
	database.DB.Find(&pacientes)
	c.JSON(http.StatusOK, pacientes)
}

func ListarPaciente(c *gin.Context) {
	var paciente models.Paciente
	id := c.Param("id")
	database.DB.Find(&paciente, id)
	c.JSON(http.StatusOK, paciente)
}

func Criar(c *gin.Context) {
	var paciente = models.Paciente{}

	if err := c.ShouldBind(&paciente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	if err := models.ValidadorPaciente(&paciente); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&paciente)
	c.JSON(http.StatusCreated, gin.H{"mensagem": "paciente criado com sucesso"})
}

func Editar(c *gin.Context) {
	var paciente models.Paciente
	id := c.Param("id")

	database.DB.Find(&paciente, id)

	if err := c.ShouldBindJSON(&paciente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	if err := models.ValidadorPaciente(&paciente); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&paciente).UpdateColumns(paciente)
	c.JSON(http.StatusOK, gin.H{"mensagem": "paciente atualizado com sucesso."})
}

func Excluir(c *gin.Context) {
	var paciente models.Paciente
	id := c.Param("id")
	database.DB.Delete(&paciente, id)
	c.JSON(http.StatusOK, gin.H{"mensagem": "Aluno excluido com sucesso."})
}

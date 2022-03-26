package database

import (
	"api-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConecaoCOmBancoDeDados() {
	stringDeConecao := "host=localhost user=postgres password=postgres dbname=api_go port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(stringDeConecao))
	if err != nil {
		log.Panic("Erro ao conectar ao Banco de Dados")
	}
	DB.AutoMigrate(&models.Paciente{})
}

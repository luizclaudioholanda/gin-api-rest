package database

import (
	"log"

	"github.com/luizclaudioholanda/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBancoDeDados() {

	stringConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao))

	if err != nil {
		log.Panic("Erro ao conectar no banco de dados.")
	}

	DB.AutoMigrate(&models.Aluno{})
}

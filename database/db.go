package database

import (
	"log"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	enderecoDB := os.Getenv("DB_HOST")
	usuarioDB := os.Getenv("DB_USER")
	senhaDB := os.Getenv("DB_PASSWORD")
	nomeDB := os.Getenv("DB_NAME")
	portaDB := os.Getenv("DB_PORT")
	stringDeConexao := "host=" + enderecoDB + " user=" + usuarioDB + " password=" + senhaDB + " dbname=" + nomeDB + " port=" + portaDB + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	_ = DB.AutoMigrate(&models.Aluno{})
}

package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const DRIVER = "mysql"

const PROTOCOL = "tcp(db_mysql)" //descomentar quando usar docker
const DBCONF = "charset=utf8mb4&parseTime=True&loc=Local"

func Connect() *gorm.DB {
	CONNECT := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + PROTOCOL + "/" + os.Getenv("DB_NAME") + "?" + DBCONF
	// CONNECT := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_PROTOCOL") + "/" + os.Getenv("DB_NAME") + "?" + DBCONF
	db, err := gorm.Open(DRIVER, CONNECT)

	if err != nil {
		fmt.Println("Não foi possível conectar ao banco de dados!")
		panic(err.Error())
	}
	// fmt.Println("Banco conectado com sucesso, na porta 3306")

	return db
}

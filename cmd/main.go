package main

import (
	"movies-crud/api"
	db "movies-crud/scripts"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	dbConn, err := gorm.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@/filmes_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Falha ao conectar na base")
	}
	defer dbConn.Close()

	db.Migrate(dbConn)

	r := gin.Default()

	api.SetupRoutes(r, dbConn)

	r.Run()
}

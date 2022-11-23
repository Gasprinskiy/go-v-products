package main

import (
	"jun2/cur/config"
	"jun2/cur/external/ginapi"
	"jun2/cur/rimport"
	"jun2/cur/uimport"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"*"},
	// }))

	r.Use(cors.Default())

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	if conf == nil {
		log.Fatalln("нет конфига")
	}

	db, err := sqlx.Open("postgres", conf.DbConn())
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	ri := rimport.NewRepositoryImports(db)
	ui := uimport.NewUsecaseImports(ri)
	api := ginapi.NewGinApi(conf, r, ui)

	api.StartServer()
}

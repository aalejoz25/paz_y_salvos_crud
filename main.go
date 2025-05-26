package main

import (
	"os" // <- nuevo

	_ "github.com/aalejoz25/paz_y_salvos_crud/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
)

func main() {
	orm.Debug = true

	err := orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("sqlconn"))
	if err != nil {
		panic("Error registrando la BD: " + err.Error())
	}
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "x-requested-with", "content-type", "accept", "origin", "authorization", "x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// ===== Lectura del puerto desde la variable de entorno =====
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback para desarrollo local
	}
	// Ejecuta Beego en ":<PORT>"
	beego.Run(":" + port)
}

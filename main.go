package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	"event-diary/internal/api"

	_ "event-diary/docs"
)

// @title MonitoLara - Diário de Crise
// @description Ferramenta simplificada para registro de crises de Lara

// @host localhost:8080
// @BasePath /diario
func main() {
	log := zap.NewExample().Sugar()
	api, err := api.NewHandler(log)
	if err != nil {
		log.Error("api.NewHandler", err)
	}

	r := gin.Default()

	r.Static("/bar.html", "./bar.html")
	r.LoadHTMLFiles("bar.html")
	g := r.Group("diario")

	g.GET("/registro", api.AddRecord)
	g.GET("/relatorio", api.Report)
	g.GET("/registros", api.AllRecords)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run()
}

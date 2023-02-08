package main

import (
	"net/http"

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
	r.GET("/relatorio", func(c *gin.Context) {
		// file, err := os.ReadFile("bar.html")
		// if err != nil {
		// 	c.JSON(http.StatusNotFound, web.APIError{ErrorCode: http.StatusNotFound, ErrorMessage: "relatório não encontrado; experimente fazer alguns registros"})
		// }
		c.HTML(
			http.StatusOK,
			"bar.html",
			gin.H{
				"content": "Esse é um relatório de crises",
				"title":   "Relatório de crises",
				"url":     "/bar.html",
			},
		)
	})
	g.GET("/registros", api.AllRecords)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run()
}

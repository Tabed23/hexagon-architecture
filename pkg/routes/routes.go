package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tabed23/hexagon-architecture/pkg/rest"
	"github.com/tabed23/hexagon-architecture/pkg/service"
	"github.com/tabed23/hexagon-architecture/pkg/storage"
)

func Routes(endpoints *rest.Rest) *gin.Engine {
	r := gin.Default()
	r.GET("/candy", endpoints.GetCandy)
	r.POST("/candy", endpoints.PostCandy)
	r.PUT("/candy/:id", endpoints.PuteCandy)
	r.DELETE("/candy/:id", endpoints.DeleteCandy)
	return r
}

func Init() {
	fmt.Println("Starting server")
	db, err := storage.SetupStorage()
	if err != nil {
		log.Fatal(err)
	}
	svc := service.NewService(db)
	app := rest.NewRest(svc)

	http.ListenAndServe(":8080", Routes(app))
}

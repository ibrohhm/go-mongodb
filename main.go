package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-mongo/connection"
	"github.com/go-mongo/handler"
	"github.com/go-mongo/repository"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}

	conn := connection.MongoDB()
	pRepo := repository.NewProductRepository(conn)
	pHandler := handler.NewProductHandler(pRepo)

	router := httprouter.New()
	router.GET("/healtzh", pHandler.Healthz)
	router.GET("/products", pHandler.GetAll)
	router.GET("/products/:id", pHandler.Get)
	router.POST("/products", pHandler.Insert)
	router.PUT("/products/:id", pHandler.Update)
	router.DELETE("/products/:id", pHandler.Delete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("listen at port " + port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		panic(fmt.Sprintf("%s: %s", "Failed to listen and server", err))
	}
}

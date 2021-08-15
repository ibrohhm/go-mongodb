package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-mongo/connection"
	"github.com/go-mongo/handler"
	m "github.com/go-mongo/middleware"
	"github.com/go-mongo/repository"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn := connection.MongoDB(ctx)
	defer conn.Disconnect(ctx)

	pRepo := repository.NewProductRepository(conn)
	pHandler := handler.NewProductHandler(pRepo)

	router := httprouter.New()
	router.GET("/healtzh", Healthz)
	router.GET("/products", m.Middleware(pHandler.GetAll))
	router.GET("/products/:id", m.Middleware(pHandler.Get))
	router.POST("/products", m.Middleware(pHandler.Insert))
	router.PUT("/products/:id", m.Middleware(pHandler.Update))
	router.DELETE("/products/:id", m.Middleware(pHandler.Delete))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("listen at port " + port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		panic(fmt.Sprintf("%s: %s", "Failed to listen and server", err))
	}
}

func Healthz(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("ok")
}

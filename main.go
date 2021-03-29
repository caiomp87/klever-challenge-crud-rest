package main

import (
	"backend/config"
	"backend/controllers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()

	router := mux.NewRouter()
	router.HandleFunc("/api/cryptos", controllers.FindAll).Methods(http.MethodGet)
	router.HandleFunc("/api/cryptos/{id}", controllers.FindById).Methods(http.MethodGet)
	router.HandleFunc("/api/cryptos", controllers.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/cryptos", controllers.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/cryptos/{id}", controllers.Delete).Methods(http.MethodDelete)

	router.HandleFunc("/api/cryptos/likes/add/{id}", controllers.AddLike).Methods(http.MethodPatch)
	router.HandleFunc("/api/cryptos/likes/remove/{id}", controllers.RemoveLike).Methods(http.MethodPatch)

	router.HandleFunc("/api/cryptos/dislikes/add/{id}", controllers.AddDislike).Methods(http.MethodPatch)
	router.HandleFunc("/api/cryptos/dislikes/remove/{id}", controllers.RemoveDislike).Methods(http.MethodPatch)

	apiPort := fmt.Sprintf(":%s", os.Getenv("API_PORT"))

	fmt.Printf("API listening on port %s", apiPort)
	log.Fatal(http.ListenAndServe(apiPort, router))
}

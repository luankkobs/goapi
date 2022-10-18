package routes

import (
	"github.com/gorilla/mux"
	"github.com/luankkobs/goapi/controllers"
	"github.com/luankkobs/goapi/middleware"
	"log"
	"net/http"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.GiveAPersonality).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.CreateNewPersonality).Methods("POST")
	r.HandleFunc("/api/personalities/{id}", controllers.DeleteAPersonality).Methods("DELETE")
	r.HandleFunc("/api/personalities/{id}", controllers.EditAPersonality).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}

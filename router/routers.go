package router

import (
	"linkingsign/handlers"
	"linkingsign/services"
	"net/http"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	baseUrl := "/api"
	version := "v1"
	baseUrl = baseUrl + "/" + version

	TopicRouter(router, baseUrl)
	WordRouter(router, baseUrl)
	UploadRouter(router, baseUrl)

	return router
}

func TopicRouter(router *mux.Router, baseUrl string) {
	topicURL := baseUrl + "/topic"
	topicRouter := router.PathPrefix(topicURL).Subrouter()
	topicRouter.HandleFunc("", handlers.GetAllTopic).Methods(http.MethodGet)
	topicRouter.HandleFunc("/{id}", handlers.GetTopic).Methods(http.MethodGet)
	topicRouter.HandleFunc("", handlers.CreateTopic).Methods(http.MethodPost)
	topicRouter.HandleFunc("/{id}", handlers.UpdateTopic).Methods(http.MethodPut)
	topicRouter.HandleFunc("/{id}", handlers.DeleteTopic).Methods(http.MethodDelete)
}

func WordRouter(router *mux.Router, baseUrl string) {
	wordURL := baseUrl + "/word"
	wordRouter := router.PathPrefix(wordURL).Subrouter()
	wordRouter.HandleFunc("", handlers.GetAllWord).Methods(http.MethodGet)
	wordRouter.HandleFunc("/{id}", handlers.GetWord).Methods(http.MethodGet)
	wordRouter.HandleFunc("", handlers.CreateWord).Methods(http.MethodPost)
	wordRouter.HandleFunc("/{id}", handlers.UpdateWord).Methods(http.MethodPut)
	wordRouter.HandleFunc("/{id}", handlers.DeleteWord).Methods(http.MethodDelete)
}

func UploadRouter(router *mux.Router, baseUrl string) {
	uploadURL := baseUrl + "/upload"
	uploadRouter := router.PathPrefix(uploadURL).Subrouter()
	uploadRouter.HandleFunc("", services.UploadFile).Methods(http.MethodPost)
}

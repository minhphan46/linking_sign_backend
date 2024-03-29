package router

import (
	"io"
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

	HomeRouter(router, baseUrl)
	TopicRouter(router, baseUrl)
	WordRouter(router, baseUrl)
	UploadRouter(router, baseUrl)

	return router
}

func HomeRouter(router *mux.Router, baseUrl string) {
	router.HandleFunc(baseUrl, BaseRouter).Methods(http.MethodGet)
}

func BaseRouter(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to LinkingSign API")
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
	wordRouter.Path("").Queries("topic_id", "{topic_id}").HandlerFunc(handlers.GetAllWordByTopicId).Methods(http.MethodGet)
	wordRouter.HandleFunc("", handlers.GetAllWord).Methods(http.MethodGet)
	wordRouter.HandleFunc("/{id}", handlers.GetWord).Methods(http.MethodGet)
	wordRouter.HandleFunc("", handlers.CreateWord).Methods(http.MethodPost)
	wordRouter.HandleFunc("/{id}", handlers.UpdateWord).Methods(http.MethodPut)
	wordRouter.HandleFunc("/{id}", handlers.DeleteWord).Methods(http.MethodDelete)
}

func UploadRouter(router *mux.Router, baseUrl string) {
	uploadURL := baseUrl + "/upload"
	uploadRouter := router.PathPrefix(uploadURL).Subrouter()
	uploadRouter.HandleFunc("", services.UploadCloud).Methods(http.MethodPost)
}

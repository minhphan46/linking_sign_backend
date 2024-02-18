package router

import (
	"linkingsign/handlers"

	"github.com/gorilla/mux"

	"net/http"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/topic/{id}", handlers.GetTopic).Methods(http.MethodGet)
	router.HandleFunc("/api/topic", handlers.GetAllTopic).Methods(http.MethodGet)
	router.HandleFunc("/api/topic", handlers.CreateTopic).Methods(http.MethodPost)
	router.HandleFunc("/api/topic/{id}", handlers.UpdateTopic).Methods(http.MethodPut)
	router.HandleFunc("/api/topic/{id}", handlers.DeleteTopic).Methods(http.MethodDelete)

	return router
}

// func (server *Server) routes() {
// 	api := server.app.Group("/api")

// 	api.Route("/profile", routes.ProfileRoute)
// 	api.Route("/users", routes.UserRoute)
// 	api.Route("/skills", routes.SkillsRoute)
// 	api.Route("/companies", routes.CompanyRoute)
// 	api.Route("/jobs", routes.JobsRoute)
// 	api.Route("/experiences", routes.ExperienceRoute)
// 	api.Route("/applyrequests", routes.ApplyRequestsRoute)
// }

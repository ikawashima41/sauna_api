package controllers

import "saunaApi/api/middlewares"

func (s *Server) initializeRoutes() {

	// TOP
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Top)).Methods("GET")

	// Login
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	// Users
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUserList)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	// Saunas
	s.Router.HandleFunc("/saunas", middlewares.SetMiddlewareJSON(s.CreateSauna)).Methods("POST")
	s.Router.HandleFunc("/saunas", middlewares.SetMiddlewareJSON(s.GetSaunaList)).Methods("GET")
	s.Router.HandleFunc("/saunas/{id}", middlewares.SetMiddlewareJSON(s.GetSauna)).Methods("GET")
	s.Router.HandleFunc("/saunas/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateSauna))).Methods("PUT")
	s.Router.HandleFunc("/saunas/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteSauna)).Methods("DELETE")
}

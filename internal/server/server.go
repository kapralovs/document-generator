package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) Run() error {
	return http.ListenAndServe(":8080", s.router)
}

func New() *server {
	return &server{
		router: mux.NewRouter(),
	}
}

func (s *server) initRouter() {
	s.router.HandleFunc("/gendoc", s.genDoc())
}

func (s *server) genDoc() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

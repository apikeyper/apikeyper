package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/health", s.HealthHandler)

	// Workspace
	r.Post("/workspace", s.CreateWorkspaceHandler)

	// RootKey FIXME: Add JWT auth
	r.Post("/rootKey", s.CreateRootKeyHandler)

	// Protected routes
	// Api
	r.Post("/api", Auth(s.Db, s.CreateApiHandler))

	// Api Key
	r.Post("/apiKey", Auth(s.Db, s.CreateApiKeyHandler))

	return r
}

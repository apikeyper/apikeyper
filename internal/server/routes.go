package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/health", s.HealthHandler)

	// Workspace
	r.Post("/workspace", s.CreateWorkspaceHandler)
	// r.Get("/workspace/{workspace_id}", s.FetchWorkspaceHandler)
	// r.Get("/workspace/list", s.ListWorkspacesHandler)
	// r.Put("/workspace/{workspace_id}", s.UpdateWorkspaceHandler)
	// r.Delete("/workspace/{workspace_id}", s.DeleteWorkspaceHandler)

	// RootKey FIXME: Add JWT auth
	r.Post("/rootKey", s.CreateRootKeyHandler)
	// r.Get("/rootKey/{root_key}", s.FetchRootKeyHandler)
	// r.Get("/rootKey/list", s.ListRootKeysHandler)
	// r.Put("/rootKey/{root_key}", s.UpdateRootKeyHandler)
	// r.Delete("/rootKey/{root_key}", s.DeleteRootKeyHandler)

	// Protected routes
	// Api
	r.Post("/api", Auth(s.Db, s.CreateApiHandler))
	r.Get("/api/{api_id}", Auth(s.Db, s.FetchApiHandler))
	r.Get("/api/list", Auth(s.Db, s.ListApsiHandler))
	// r.Put("/api/{api_id}", Auth(s.Db, s.UpdateApiHandler))
	// r.Delete("/api/{api_id}", Auth(s.Db, s.DeleteApiHandler))

	// Api Key
	r.Post("/apiKey", Auth(s.Db, s.CreateApiKeyHandler))
	r.Post("/apiKey/verify", Auth(s.Db, s.VerifyApiKeyHandler))
	r.Put("/apiKey/revoke", Auth(s.Db, s.RevokeApiKeyHandler))
	// r.Get("/apiKey/{api_key}", Auth(s.Db, s.FetchApiKeyHandler)
	// r.Get("/apiKey/list", Auth(s.Db, s.ListApiKeysHandler)
	// r.Put("/apiKey/{api_key}", Auth(s.Db, s.UpdateApiKeyHandler)
	// r.Delete("/apiKey/{api_key}", Auth(s.Db, s.DeleteApiKeyHandler)
	r.Get("/apiKey/{api_key_id}/usage", Auth(s.Db, s.FetchApiKeyUsageHandler))

	return r
}

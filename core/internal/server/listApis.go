package server

import (
	"net/http"
)

func (s *Server) ListApsiHandler(w http.ResponseWriter, r *http.Request) {
	// Get rootKey from context
	session := r.Context().Value("session").(Session)

	apis, err := s.Db.ListApis(session.WorkspaceId)

	if err != nil {
		encode(w, r, http.StatusInternalServerError, "Failed to retrieve apis for workspace")
		return
	}

	var respBody []FetchApiResponse

	for _, api := range *apis {
		respBody = append(respBody, FetchApiResponse{
			ApiId:       api.ID,
			WorkspaceId: api.WorkspaceId,
			ApiName:     api.ApiName,
			CreatedAt:   api.CreatedAt,
			UpdatedAt:   api.UpdatedAt,
		})
	}

	encode(w, r, http.StatusOK, respBody)
}

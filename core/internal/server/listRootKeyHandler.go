package server

import (
	"net/http"
)

func (s *Server) ListRootKeysHandler(w http.ResponseWriter, r *http.Request) {

	// Get rootKey from context
	session := r.Context().Value("session").(Session)

	rootKey, err := s.Db.ListRootKeysForWorkspace(session.WorkspaceId)

	if err != nil {
		encode(w, r, http.StatusInternalServerError, "Failed to retrieve rootKey for workspace")
		return
	}

	var respBody []FetchRootKeyResponse

	for _, rk := range *rootKey {
		respBody = append(respBody, FetchRootKeyResponse{
			RootKeyId:   rk.ID,
			WorkspaceId: rk.WorkspaceId,
			RootKeyName: rk.RootKeyName,
			CreatedAt:   rk.CreatedAt,
			UpdatedAt:   rk.UpdatedAt,
		})
	}

	encode(w, r, http.StatusOK, respBody)
}

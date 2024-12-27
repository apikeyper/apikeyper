package server

import (
	"net/http"
)

func (s *Server) FetchAllApiKeysActivityHandler(w http.ResponseWriter, r *http.Request) {
	// Get rootKey from context
	session := r.Context().Value("session").(Session)

	apiKeyUsageRecords, err := s.Db.ListAllApiKeysActivity(session.WorkspaceId)

	if err != nil {
		encode(w, r, http.StatusInternalServerError, "Failed to retrieve apis for workspace")
		return
	}

	var respBody []ApiKeysActivityResponse

	for _, record := range *apiKeyUsageRecords {
		respBody = append(respBody, ApiKeysActivityResponse{
			KeyId:     record.ApiKeyId,
			KeyName:   *record.KeyName,
			ApiId:     record.ApiId,
			Usage:     record.Usage,
			Timestamp: record.Timestamp,
		})
	}

	encode(w, r, http.StatusCreated, respBody)
}

package schemas

type CreateApiKeyRequest struct {
	ApiId  string   `json:"apiId"`
	Name   string   `json:"name,omitempty"`
	Prefix string   `json:"prefix,omitempty"`
	Roles  []string `json:"roles,omitempty"`
}

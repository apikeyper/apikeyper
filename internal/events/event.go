package events

type Event string

const (
	API_CREATED            Event = "api.created"
	API_UPDATED            Event = "api.updated"
	API_DELETED            Event = "api.deleted"
	API_KEY_CREATED        Event = "api.key.created"
	API_KEY_DELETED        Event = "api.key.deleted"
	API_KEY_VERIFY_SUCCESS Event = "api.key.verify.success"
	API_KEY_VERIFY_FAILED  Event = "api.key.verify.failed"
	API_KEY_RATE_LIMITED   Event = "api.key.rate.limited"
	API_KEY_REVOKED        Event = "api.key.revoked"
)

type EventData struct {
	WorkspaceId string `json:"workspace_id"`
	ApiKeyId    string `json:"api_key_id"`
	ApiId       string `json:"api_id"`
	EventTime   string `json:"event_time"`
}

type EventPayload struct {
	EventType Event     `json:"event_type"`
	Data      EventData `json:"data"`
}

package heapanalytics

// Event represents an individual heap event that can be sent to the heap API
type Event struct {
	AppID    string `json:"app_id"`
	Identity string `json:"identity"`
	Event    string `json:"event,omitempty"`
	// properties is an interface as it could be a string or int
	Properties map[string]interface{} `json:"properties,omitempty"`
}

func NewEvent(appID string, identity, event string, properties map[string]interface{}) *Event {
	return &Event{appID, identity, event, properties}
}

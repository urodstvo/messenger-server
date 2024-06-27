package types

type WebSocketMessage struct {
	EventName string `json:"eventName"`
	Data      any    `json:"data"`
	CreatedAt string `json:"createdAt"`
}

type NewMessage struct {
	Text string `json:"text"`
}

type EditMessage struct {
	MessageID uint32 `json:"messageId"`
	Text      string `json:"text"`
}

type ReadMessage struct {
	MessageID uint32 `json:"messageId"`
}

type DeleteMessage struct {
	MessageID uint32 `json:"messageId"`
}

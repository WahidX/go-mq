package entities

type Message struct {
	Header    map[string]any
	Data      []byte
	Timestamp int64
	Topic     string
}

type MessageRow struct {
	ID        int64 // auto incremented pk
	Data      []byte
	Timestamp int64
	Topic     string
}

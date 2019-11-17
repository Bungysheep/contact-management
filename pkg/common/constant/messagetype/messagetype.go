package messagetype

// MessageType type
type MessageType int

const (
	// Information message type
	Information MessageType = iota

	// Warning message type
	Warning

	// Error message type
	Error
)

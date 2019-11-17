package message

import "github.com/bungysheep/contact-management/pkg/common/constant/messagetype"

import "fmt"

// IMessage interface
type IMessage interface {
	IsError() bool
	Error() string
}

// Message type
type Message struct {
	code        string
	messageType messagetype.MessageType
	definition  string
	args        []interface{}
}

// NewMessage creates new message
func NewMessage(code string, messageType messagetype.MessageType, definition string, args ...interface{}) IMessage {
	return &Message{
		code:        code,
		messageType: messageType,
		definition:  definition,
		args:        args,
	}
}

// IsError whether is error or not
func (m *Message) IsError() bool {
	return m.messageType == messagetype.Error
}

func (m *Message) Error() string {
	return fmt.Sprintf("[%s] %s", m.code, fmt.Sprintf(m.definition, m.args))
}

package message

import (
	"fmt"

	"github.com/bungysheep/contact-management/pkg/common/constant/messagetype"
)

// IMessage interface
type IMessage interface {
	Code() string
	Type() messagetype.MessageType
	IsError() bool
	IsWarning() bool
	ShortDescription() string
	LongDescription() string
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

// GetMessage returns message
func GetMessage(err IMessage) string {
	if err == nil {
		return ""
	}
	return err.LongDescription()
}

// Code is message code
func (m *Message) Code() string {
	return m.code
}

// Type is message type
func (m *Message) Type() messagetype.MessageType {
	return m.messageType
}

// IsError whether is error or not
func (m *Message) IsError() bool {
	return m.messageType == messagetype.Error
}

// IsWarning whether is warning or not
func (m *Message) IsWarning() bool {
	return m.messageType == messagetype.Warning
}

// ShortDescription is message short description
func (m *Message) ShortDescription() string {
	return fmt.Sprintf("%s", fmt.Sprintf(m.definition, m.args...))
}

// LongDescription is message long description
func (m *Message) LongDescription() string {
	return fmt.Sprintf("[%s] %s", m.code, fmt.Sprintf(m.definition, m.args...))
}

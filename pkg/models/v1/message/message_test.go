package message

import (
	"testing"

	"github.com/bungysheep/contact-management/pkg/common/constant/messagetype"
)

func TestCreateMessage(t *testing.T) {
	msg := NewMessage("CNT0001", messagetype.Error, "%s does not exist", "Contact System")

	if msg == nil {
		t.Fatalf("Expect message is not nil")
	}

	if msg.Code() != "CNT0001" {
		t.Errorf("Expect message code %v, but got %v", "CNT0001", msg.Code())
	}

	if msg.Type() != messagetype.Error {
		t.Errorf("Expect message type %v, but got %v", messagetype.Error, msg.Type())
	}

	if !msg.IsError() {
		t.Errorf("Expect message is an error")
	}

	if msg.IsWarning() {
		t.Errorf("Expect message is not a warning")
	}

	if msg.ShortDescription() != "Contact System does not exist" {
		t.Errorf("Expect message short description %v, but got %v", "Contact System does not exist", msg.ShortDescription())
	}

	if msg.LongDescription() != "[CNT0001] Contact System does not exist" {
		t.Errorf("Expect message long description %v, but got %v", "[CNT0001] Contact System does not exist", msg.LongDescription())
	}
}

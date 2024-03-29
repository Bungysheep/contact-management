package contactcommunicationmethodfield

import (
	"testing"
)

func TestCreateContactCommunicationMethodField(t *testing.T) {
	contactCommMethodField := NewContactCommunicationMethodField()

	if contactCommMethodField == nil {
		t.Fatalf("Expect contact communication method field is not nil")
	}

	contactCommMethodField.ContactSystemCode = "CNTSYS001"
	contactCommMethodField.ContactID = 1
	contactCommMethodField.ContactCommunicationMethodID = 1
	contactCommMethodField.FieldCode = "EMAIL_ADDRESS"
	contactCommMethodField.FieldValue = "test@gmail.com"

	if contactCommMethodField.GetContactSystemCode() != "CNTSYS001" {
		t.Errorf("Expect contact system code %v, but got %v", "CNTSYS001", contactCommMethodField.GetContactSystemCode())
	}

	if contactCommMethodField.GetContactID() != 1 {
		t.Errorf("Expect contact id %v, but got %v", 1, contactCommMethodField.GetContactID())
	}

	if contactCommMethodField.GetContactCommunicationMethodID() != 1 {
		t.Errorf("Expect contact communication method id %v, but got %v", 1, contactCommMethodField.GetContactCommunicationMethodID())
	}

	if contactCommMethodField.GetFieldCode() != "EMAIL_ADDRESS" {
		t.Errorf("Expect field code %v, but got %v", "EMAIL_ADDRESS", contactCommMethodField.GetFieldCode())
	}

	if contactCommMethodField.GetFieldValue() != "test@gmail.com" {
		t.Errorf("Expect field value %v, but got %v", "test@gmail.com", contactCommMethodField.GetFieldValue())
	}
}

func TestValidate(t *testing.T) {
	contactCommMethodField := NewContactCommunicationMethodField()

	if contactCommMethodField == nil {
		t.Fatalf("Expect contact communication method field is not nil")
	}

	contactCommMethodField.ContactSystemCode = "CNTSYS001"
	contactCommMethodField.ContactID = 1
	contactCommMethodField.ContactCommunicationMethodID = 1
	contactCommMethodField.FieldCode = "EMAIL_ADDRESS"
	contactCommMethodField.FieldValue = "test@gmail.com"

	if err := contactCommMethodField.DoValidate(); err != nil {
		t.Fatalf("Expect error is nil, but got %v", err)
	}
}

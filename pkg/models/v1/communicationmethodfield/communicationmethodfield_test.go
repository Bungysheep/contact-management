package communicationmethodfield

import (
	"testing"
	"time"
)

func TestCreateCommunicationMethodField(t *testing.T) {
	commMethodField := NewCommunicationMethodField()

	if commMethodField == nil {
		t.Fatalf("Expect communication method field is not nil")
	}

	timeNow := time.Now()

	commMethodField.ContactSystemCode = "CNTSYS001"
	commMethodField.CommunicationMethodCode = "EMAIL"
	commMethodField.FieldCode = "EMAIL_ADDRESS"
	commMethodField.Caption = "Email Address"
	commMethodField.Sequence = 1
	commMethodField.Audit.CreatedAt = timeNow
	commMethodField.Audit.ModifiedAt = timeNow
	commMethodField.Audit.Vers = 1

	if commMethodField.GetContactSystemCode() != "CNTSYS001" {
		t.Errorf("Expect contact system code %v, but got %v", "CNTSYS001", commMethodField.GetContactSystemCode())
	}

	if commMethodField.GetCommunicationMethodCode() != "EMAIL" {
		t.Errorf("Expect communication method code %v, but got %v", "EMAIL", commMethodField.GetCommunicationMethodCode())
	}

	if commMethodField.GetFieldCode() != "EMAIL_ADDRESS" {
		t.Errorf("Expect field code %v, but got %v", "EMAIL_ADDRESS", commMethodField.GetFieldCode())
	}

	if commMethodField.GetCaption() != "Email Address" {
		t.Errorf("Expect caption %v, but got %v", "Email Address", commMethodField.GetCaption())
	}

	if commMethodField.GetSequence() != 1 {
		t.Errorf("Expect sequence %v, but got %v", 1, commMethodField.GetSequence())
	}

	if commMethodField.GetAudit().GetCreatedAt() != timeNow {
		t.Errorf("Expect created at %v, but got %v", timeNow, commMethodField.GetAudit().GetCreatedAt())
	}

	if commMethodField.GetAudit().GetModifiedAt() != timeNow {
		t.Errorf("Expect modified at %v, but got %v", timeNow, commMethodField.GetAudit().GetModifiedAt())
	}

	if commMethodField.GetAudit().GetVers() != 1 {
		t.Errorf("Expect vers %v, but got %v", 1, commMethodField.GetAudit().GetVers())
	}
}

func TestValidate(t *testing.T) {
	commMethodField := NewCommunicationMethodField()

	if commMethodField == nil {
		t.Fatalf("Expect communication method field is not nil")
	}

	timeNow := time.Now()

	commMethodField.ContactSystemCode = "CNTSYS001"
	commMethodField.CommunicationMethodCode = "EMAIL"
	commMethodField.FieldCode = "EMAIL_ADDRESS"
	commMethodField.Caption = "Email Address"
	commMethodField.Sequence = 1
	commMethodField.Audit.CreatedAt = timeNow
	commMethodField.Audit.ModifiedAt = timeNow
	commMethodField.Audit.Vers = 1

	if err := commMethodField.DoValidate(); err != nil {
		t.Fatalf("Expect error is nil, but got %v", err)
	}
}

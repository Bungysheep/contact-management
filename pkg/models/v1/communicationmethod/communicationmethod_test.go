package communicationmethod

import (
	"testing"
	"time"
)

func TestCreateCommunicationMethod(t *testing.T) {
	commMethod := NewCommunicationMethod()

	if commMethod == nil {
		t.Fatalf("Expect communication method is not nil")
	}

	timeNow := time.Now()

	commMethod.ContactSystemCode = "CNTSYS001"
	commMethod.CommunicationMethodCode = "EMAIL"
	commMethod.Description = "Email"
	commMethod.Details = "Email"
	commMethod.Status = "A"
	commMethod.FormatField = "[EMAIL_ADDRESS]"
	commMethod.Audit.CreatedAt = timeNow
	commMethod.Audit.ModifiedAt = timeNow
	commMethod.Audit.Vers = 1

	if commMethod.GetContactSystemCode() != "CNTSYS001" {
		t.Errorf("Expect contact system code %v, but got %v", "CNTSYS001", commMethod.GetContactSystemCode())
	}

	if commMethod.GetCommunicationMethodCode() != "EMAIL" {
		t.Errorf("Expect communication method code %v, but got %v", "EMAIL", commMethod.GetCommunicationMethodCode())
	}

	if commMethod.GetDescription() != "Email" {
		t.Errorf("Expect description %v, but got %v", "Email", commMethod.GetDescription())
	}

	if commMethod.GetDetails() != "Email" {
		t.Errorf("Expect details %v, but got %v", "Email", commMethod.GetDetails())
	}

	if commMethod.GetStatus() != "A" {
		t.Errorf("Expect status %v, but got %v", "A", commMethod.GetStatus())
	}

	if commMethod.GetFormatField() != "[EMAIL_ADDRESS]" {
		t.Errorf("Expect format field %v, but got %v", "[EMAIL_ADDRESS]", commMethod.GetFormatField())
	}

	if commMethod.GetAudit().GetCreatedAt() != timeNow {
		t.Errorf("Expect created at %v, but got %v", timeNow, commMethod.GetAudit().GetCreatedAt())
	}

	if commMethod.GetAudit().GetModifiedAt() != timeNow {
		t.Errorf("Expect modified at %v, but got %v", timeNow, commMethod.GetAudit().GetModifiedAt())
	}

	if commMethod.GetAudit().GetVers() != 1 {
		t.Errorf("Expect vers %v, but got %v", 1, commMethod.GetAudit().GetVers())
	}
}

func TestValidate(t *testing.T) {
	commMethod := NewCommunicationMethod()

	if commMethod == nil {
		t.Fatalf("Expect communication method is not nil")
	}

	timeNow := time.Now()

	commMethod.ContactSystemCode = "CNTSYS001"
	commMethod.CommunicationMethodCode = "EMAIL"
	commMethod.Description = "Email"
	commMethod.Details = "Email"
	commMethod.Status = "A"
	commMethod.FormatField = "[EMAIL_ADDRESS]"
	commMethod.Audit.CreatedAt = timeNow
	commMethod.Audit.ModifiedAt = timeNow
	commMethod.Audit.Vers = 1

	if !commMethod.DoValidate() {
		t.Fatalf("Expect TRUE")
	}
}

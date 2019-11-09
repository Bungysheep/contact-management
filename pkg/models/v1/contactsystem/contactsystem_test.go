package contactsystem

import (
	"testing"
	"time"
)

func TestCreateContactSystem(t *testing.T) {
	contactSystem := NewContactSystem()

	if contactSystem == nil {
		t.Fatalf("Expect contact system is not nil")
	}

	timeNow := time.Now()

	contactSystem.ContactSystemCode = "CNTSYS001"
	contactSystem.Description = "Contact System 1"
	contactSystem.Details = "Contact System 1"
	contactSystem.Status = "A"
	contactSystem.Audit.CreatedAt = timeNow
	contactSystem.Audit.ModifiedAt = timeNow
	contactSystem.Audit.Vers = 1

	if contactSystem.GetContactSystemCode() != "CNTSYS001" {
		t.Errorf("Expect contact system code %v, but got %v", "CNTSYS001", contactSystem.GetContactSystemCode())
	}

	if contactSystem.GetDescription() != "Contact System 1" {
		t.Errorf("Expect description %v, but got %v", "Contact System 1", contactSystem.GetDescription())
	}

	if contactSystem.GetDetails() != "Contact System 1" {
		t.Errorf("Expect details %v, but got %v", "Contact System 1", contactSystem.GetDetails())
	}

	if contactSystem.GetStatus() != "A" {
		t.Errorf("Expect status %v, but got %v", "A", contactSystem.GetStatus())
	}

	if contactSystem.GetAudit().GetCreatedAt() != timeNow {
		t.Errorf("Expect craeted at %v, but got %v", timeNow, contactSystem.GetAudit().GetCreatedAt())
	}

	if contactSystem.GetAudit().GetModifiedAt() != timeNow {
		t.Errorf("Expect modified at %v, but got %v", timeNow, contactSystem.GetAudit().GetModifiedAt())
	}

	if contactSystem.GetAudit().GetVers() != 1 {
		t.Errorf("Expect vers %v, but got %v", 1, contactSystem.GetAudit().GetVers())
	}
}

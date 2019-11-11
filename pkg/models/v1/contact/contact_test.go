package contact

import (
	"testing"
	"time"
)

func TestCreateContact(t *testing.T) {
	contact := NewContact()

	if contact == nil {
		t.Fatalf("Expect contact is not nil")
	}

	timeNow := time.Now()

	contact.ContactSystemCode = "CNTSYS001"
	contact.ContactID = 1
	contact.FirstName = "James"
	contact.LastName = "Embongbulan"
	contact.Status = "A"
	contact.Audit.CreatedAt = timeNow
	contact.Audit.ModifiedAt = timeNow
	contact.Audit.Vers = 1

	if contact.GetContactSystemCode() != "CNTSYS001" {
		t.Errorf("Expect contact system code %v, but got %v", "CNTSYS001", contact.GetContactSystemCode())
	}

	if contact.GetContactID() != 1 {
		t.Errorf("Expect description %v, but got %v", 1, contact.GetContactID())
	}

	if contact.GetFirstName() != "James" {
		t.Errorf("Expect first name %v, but got %v", "James", contact.GetFirstName())
	}

	if contact.GetLastName() != "Embongbulan" {
		t.Errorf("Expect last name %v, but got %v", "Embongbulan", contact.GetLastName())
	}

	if contact.GetStatus() != "A" {
		t.Errorf("Expect status %v, but got %v", "A", contact.GetStatus())
	}

	if contact.GetAudit().GetCreatedAt() != timeNow {
		t.Errorf("Expect craeted at %v, but got %v", timeNow, contact.GetAudit().GetCreatedAt())
	}

	if contact.GetAudit().GetModifiedAt() != timeNow {
		t.Errorf("Expect modified at %v, but got %v", timeNow, contact.GetAudit().GetModifiedAt())
	}

	if contact.GetAudit().GetVers() != 1 {
		t.Errorf("Expect vers %v, but got %v", 1, contact.GetAudit().GetVers())
	}
}

func TestValidate(t *testing.T) {
	contact := NewContact()

	if contact == nil {
		t.Fatalf("Expect contact is not nil")
	}

	timeNow := time.Now()

	contact.ContactSystemCode = "CNTSYS001"
	contact.ContactID = 1
	contact.FirstName = "James"
	contact.LastName = "Embongbulan"
	contact.Status = "A"
	contact.Audit.CreatedAt = timeNow
	contact.Audit.ModifiedAt = timeNow
	contact.Audit.Vers = 1

	if !contact.DoValidate() {
		t.Fatalf("Expect TRUE")
	}
}

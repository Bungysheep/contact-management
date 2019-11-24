package contactcommunicationmethod

import (
	"testing"
	"time"

	"github.com/bungysheep/contact-management/pkg/models/v1/contactcommunicationmethodfield"
)

func TestCreateContactCommunicationMethod(t *testing.T) {
	contactCommMethod := NewContactCommunicationMethod()

	if contactCommMethod == nil {
		t.Fatalf("Expect contact communication method is not nil")
	}

	timeNow := time.Now()

	contactCommMethod.ContactSystemCode = "CNTSYS001"
	contactCommMethod.ContactID = 1
	contactCommMethod.ContactCommunicationMethodID = 1
	contactCommMethod.CommunicationMethodCode = "EMAIL"
	contactCommMethod.CommunicationMethodLabelCode = "HOME"
	contactCommMethod.CommunicationMethodLabelCaption = "Home"
	contactCommMethod.FormatValue = "test@gmail.com"
	contactCommMethod.Status = "A"
	contactCommMethod.IsDefault = true

	contactCommMethodField := contactcommunicationmethodfield.NewContactCommunicationMethodField()
	contactCommMethodField.ContactSystemCode = "CNTSYS001"
	contactCommMethodField.ContactID = 1
	contactCommMethodField.ContactCommunicationMethodID = 1
	contactCommMethodField.FieldCode = "EMAIL_ADDRESS"
	contactCommMethodField.FieldValue = "test@gmail.com"
	contactCommMethod.ContactCommunicationMethodField = append(contactCommMethod.ContactCommunicationMethodField, contactCommMethodField)

	contactCommMethod.Audit.CreatedAt = timeNow
	contactCommMethod.Audit.ModifiedAt = timeNow
	contactCommMethod.Audit.Vers = 1

	if contactCommMethod.GetContactSystemCode() != "CNTSYS001" {
		t.Errorf("Expect contact system code %v, but got %v", "CNTSYS001", contactCommMethod.GetContactSystemCode())
	}

	if contactCommMethod.GetContactID() != 1 {
		t.Errorf("Expect contact id %v, but got %v", 1, contactCommMethod.GetContactID())
	}

	if contactCommMethod.GetContactCommunicationMethodID() != 1 {
		t.Errorf("Expect contact communication method id %v, but got %v", 1, contactCommMethod.GetContactCommunicationMethodID())
	}

	if contactCommMethod.GetCommunicationMethodCode() != "EMAIL" {
		t.Errorf("Expect communication method code %v, but got %v", "EMAIL", contactCommMethod.GetCommunicationMethodCode())
	}

	if contactCommMethod.GetCommunicationMethodLabelCode() != "HOME" {
		t.Errorf("Expect communication method label code %v, but got %v", "HOME", contactCommMethod.GetCommunicationMethodLabelCode())
	}

	if contactCommMethod.GetCommunicationMethodLabelCaption() != "Home" {
		t.Errorf("Expect communication method label caption %v, but got %v", "Home", contactCommMethod.GetCommunicationMethodLabelCaption())
	}

	if contactCommMethod.GetFormatValue() != "test@gmail.com" {
		t.Errorf("Expect format value %v, but got %v", "test@gmail.com", contactCommMethod.GetFormatValue())
	}

	if contactCommMethod.GetStatus() != "A" {
		t.Errorf("Expect status %v, but got %v", "A", contactCommMethod.GetStatus())
	}

	if contactCommMethod.GetIsDefault() != true {
		t.Errorf("Expect is default %v, but got %v", true, contactCommMethod.GetIsDefault())
	}

	for i := range contactCommMethod.GetContactCommunicationMethodField() {
		if contactCommMethod.GetContactCommunicationMethodField()[i].GetFieldCode() != "EMAIL_ADDRESS" {
			t.Errorf("Expect field code %d: %v, but got %v", i, "EMAIL_ADDRESS", contactCommMethod.GetContactCommunicationMethodField()[i].GetFieldCode())
		}

		if contactCommMethod.GetContactCommunicationMethodField()[i].GetFieldValue() != "test@gmail.com" {
			t.Errorf("Expect field value %d: %v, but got %v", i, "test@gmail.com", contactCommMethod.GetContactCommunicationMethodField()[i].GetFieldValue())
		}
	}

	if contactCommMethod.GetAudit().GetCreatedAt() != timeNow {
		t.Errorf("Expect craeted at %v, but got %v", timeNow, contactCommMethod.GetAudit().GetCreatedAt())
	}

	if contactCommMethod.GetAudit().GetModifiedAt() != timeNow {
		t.Errorf("Expect modified at %v, but got %v", timeNow, contactCommMethod.GetAudit().GetModifiedAt())
	}

	if contactCommMethod.GetAudit().GetVers() != 1 {
		t.Errorf("Expect vers %v, but got %v", 1, contactCommMethod.GetAudit().GetVers())
	}
}

func TestValidate(t *testing.T) {
	contactCommMethod := NewContactCommunicationMethod()

	if contactCommMethod == nil {
		t.Fatalf("Expect contact communication method is not nil")
	}

	timeNow := time.Now()

	contactCommMethod.ContactSystemCode = "CNTSYS001"
	contactCommMethod.ContactID = 1
	contactCommMethod.ContactCommunicationMethodID = 1
	contactCommMethod.CommunicationMethodCode = "EMAIL"
	contactCommMethod.CommunicationMethodLabelCode = "HOME"
	contactCommMethod.CommunicationMethodLabelCaption = "Home"
	contactCommMethod.FormatValue = "test@gmail.com"
	contactCommMethod.Status = "A"
	contactCommMethod.IsDefault = true

	contactCommMethodField := contactcommunicationmethodfield.NewContactCommunicationMethodField()
	contactCommMethodField.ContactSystemCode = "CNTSYS001"
	contactCommMethodField.ContactID = 1
	contactCommMethodField.ContactCommunicationMethodID = 1
	contactCommMethodField.FieldCode = "EMAIL_ADDRESS"
	contactCommMethodField.FieldValue = "test@gmail.com"
	contactCommMethod.ContactCommunicationMethodField = append(contactCommMethod.ContactCommunicationMethodField, contactCommMethodField)

	contactCommMethod.Audit.CreatedAt = timeNow
	contactCommMethod.Audit.ModifiedAt = timeNow
	contactCommMethod.Audit.Vers = 1

	if err := contactCommMethod.DoValidate(); err != nil {
		t.Fatalf("Expect error is nil, but got %v", err)
	}
}

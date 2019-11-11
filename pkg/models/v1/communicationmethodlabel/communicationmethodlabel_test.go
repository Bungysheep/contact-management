package communicationmethodlabel

import "testing"

func TestCreateCommunicationMethodLabel(t *testing.T) {
	commMethodLabel := NewCommunicationMethodLabel()

	if commMethodLabel == nil {
		t.Fatalf("Expect communication method label is not nil")
	}

	commMethodLabel.ContactSystemCode = "CNTSYS001"
	commMethodLabel.CommunicationMethodCode = "EMAIL"
	commMethodLabel.CommunicationMethodLabelCode = "HOME"
	commMethodLabel.Caption = "Home"

	if commMethodLabel.GetContactSystemCode() != "CNTSYS001" {
		t.Errorf("Expect contact system code %v, but got %v", "CNTSYS001", commMethodLabel.GetContactSystemCode())
	}

	if commMethodLabel.GetCommunicationMethodCode() != "EMAIL" {
		t.Errorf("Expect communication method code %v, but got %v", "EMAIL", commMethodLabel.GetCommunicationMethodCode())
	}

	if commMethodLabel.GetCommunicationMethodLabelCode() != "HOME" {
		t.Errorf("Expect communication method label code %v, but got %v", "HOME", commMethodLabel.GetCommunicationMethodLabelCode())
	}

	if commMethodLabel.GetCaption() != "Home" {
		t.Errorf("Expect caption %v, but got %v", "Home", commMethodLabel.GetCaption())
	}
}

func TestValidate(t *testing.T) {
	commMethodLabel := NewCommunicationMethodLabel()

	if commMethodLabel == nil {
		t.Fatalf("Expect communication method label is not nil")
	}

	commMethodLabel.ContactSystemCode = "CNTSYS001"
	commMethodLabel.CommunicationMethodCode = "EMAIL"
	commMethodLabel.CommunicationMethodLabelCode = "HOME"
	commMethodLabel.Caption = "Home"

	if err := commMethodLabel.DoValidate(); err != nil {
		t.Fatalf("Expect error is nil, but got %v", err)
	}
}

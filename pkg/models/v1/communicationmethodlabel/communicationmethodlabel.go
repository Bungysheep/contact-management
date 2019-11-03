package communicationmethodlabel

// CommunicationMethodLabel model
type CommunicationMethodLabel struct {
	ContactSystemCode            string
	CommunicationMethodCode      string
	CommunicationMethodLabelCode string
	Caption                      string
}

// NewCommunicationMethodLabel creates Communication Method Label
func NewCommunicationMethodLabel() *CommunicationMethodLabel {
	return &CommunicationMethodLabel{}
}

// GetContactSystemCode returns Contact System Code
func (cml *CommunicationMethodLabel) GetContactSystemCode() string {
	return cml.ContactSystemCode
}

// GetCommunicationMethodCode returns Communication Method Code
func (cml *CommunicationMethodLabel) GetCommunicationMethodCode() string {
	return cml.CommunicationMethodCode
}

// GetCommunicationMethodLabelCode returns Communication Method Label Code
func (cml *CommunicationMethodLabel) GetCommunicationMethodLabelCode() string {
	return cml.CommunicationMethodLabelCode
}

// GetCaption returns Caption
func (cml *CommunicationMethodLabel) GetCaption() string {
	return cml.Caption
}

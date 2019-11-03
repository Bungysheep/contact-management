package contactcommunicationmethodfield

// ContactCommunicationMethodField model
type ContactCommunicationMethodField struct {
	ContactSystemCode            string
	ContactID                    int64
	ContactCommunicationMethodID int64
	CommunicationMethodCode      string
	FieldCode                    string
	FieldValue                   string
}

// NewContactCommunicationMethodField creates Contact Communication Method Field
func NewContactCommunicationMethodField() *ContactCommunicationMethodField {
	return &ContactCommunicationMethodField{}
}

// GetContactSystemCode returns Contact System Code
func (ccmf *ContactCommunicationMethodField) GetContactSystemCode() string {
	return ccmf.ContactSystemCode
}

// GetContactID returns Contact ID
func (ccmf *ContactCommunicationMethodField) GetContactID() int64 {
	return ccmf.ContactID
}

// GetContactCommunicationMethodID returns Contact Communication Method ID
func (ccmf *ContactCommunicationMethodField) GetContactCommunicationMethodID() int64 {
	return ccmf.ContactCommunicationMethodID
}

// GetCommunicationMethodCode returns Communication Method Code
func (ccmf *ContactCommunicationMethodField) GetCommunicationMethodCode() string {
	return ccmf.CommunicationMethodCode
}

// GetFieldCode returns Field Code
func (ccmf *ContactCommunicationMethodField) GetFieldCode() string {
	return ccmf.FieldCode
}

// GetFieldValue returns Field Value
func (ccmf *ContactCommunicationMethodField) GetFieldValue() string {
	return ccmf.FieldValue
}

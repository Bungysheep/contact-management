package contactcommunicationmethod

import (
	"context"

	auditapi "github.com/bungysheep/contact-management/pkg/api/v1/audit"
	contactcommunicationmethodapi "github.com/bungysheep/contact-management/pkg/api/v1/contactcommunicationmethod"
	contactcommunicationmethodmodel "github.com/bungysheep/contact-management/pkg/models/v1/contactcommunicationmethod"
	contactcommunicationmethodservice "github.com/bungysheep/contact-management/pkg/service/v1/contactcommunicationmethod"
	"github.com/golang/protobuf/ptypes"
)

type contactcommunicationmethodServiceServer struct {
	svc contactcommunicationmethodservice.IContactCommunicationMethodService
}

// NewContactCommunicationMethodServiceServer - Contact Communication Method service server implementation
func NewContactCommunicationMethodServiceServer(svc contactcommunicationmethodservice.IContactCommunicationMethodService) contactcommunicationmethodapi.ContactCommunicationMethodServiceServer {
	return &contactcommunicationmethodServiceServer{svc: svc}
}

func (cmm *contactcommunicationmethodServiceServer) DoRead(ctx context.Context, req *contactcommunicationmethodapi.DoReadContactCommunicationMethodRequest) (*contactcommunicationmethodapi.DoReadContactCommunicationMethodResponse, error) {
	result, err := cmm.svc.DoRead(ctx, req.GetContactSystemCode(), req.GetContactId(), req.GetContactCommunicationMethodId())

	return &contactcommunicationmethodapi.DoReadContactCommunicationMethodResponse{ContactCommunicationMethod: contactCommunicationMethodModelToAPI(result)}, err
}

func (cmm *contactcommunicationmethodServiceServer) DoReadAll(ctx context.Context, req *contactcommunicationmethodapi.DoReadAllContactCommunicationMethodRequest) (*contactcommunicationmethodapi.DoReadAllContactCommunicationMethodResponse, error) {
	result, err := cmm.svc.DoReadAll(ctx, req.GetContactSystemCode(), req.GetContactId())

	resp := make([]*contactcommunicationmethodapi.ContactCommunicationMethod, 0)

	for _, item := range result {
		resp = append(resp, contactCommunicationMethodModelToAPI(item))
	}

	return &contactcommunicationmethodapi.DoReadAllContactCommunicationMethodResponse{ContactCommunicationMethod: resp}, err
}

func (cmm *contactcommunicationmethodServiceServer) DoSave(ctx context.Context, req *contactcommunicationmethodapi.DoSaveContactCommunicationMethodRequest) (*contactcommunicationmethodapi.DoSaveContactCommunicationMethodResponse, error) {
	err := cmm.svc.DoSave(ctx, contactCommunicationMethodAPIToModel(req.GetContactCommunicationMethod()))

	return &contactcommunicationmethodapi.DoSaveContactCommunicationMethodResponse{Result: err == nil}, err
}

func (cmm *contactcommunicationmethodServiceServer) DoDelete(ctx context.Context, req *contactcommunicationmethodapi.DoDeleteContactCommunicationMethodRequest) (*contactcommunicationmethodapi.DoDeleteContactCommunicationMethodResponse, error) {
	err := cmm.svc.DoDelete(ctx, req.GetContactSystemCode(), req.GetContactId(), req.GetContactCommunicationMethodId())

	return &contactcommunicationmethodapi.DoDeleteContactCommunicationMethodResponse{Result: err == nil}, err
}

func contactCommunicationMethodModelToAPI(dataModel *contactcommunicationmethodmodel.ContactCommunicationMethod) *contactcommunicationmethodapi.ContactCommunicationMethod {
	contactCommMethod := &contactcommunicationmethodapi.ContactCommunicationMethod{Audit: &auditapi.Audit{}}
	contactCommMethod.ContactSystemCode = dataModel.GetContactSystemCode()
	contactCommMethod.ContactId = dataModel.GetContactID()
	contactCommMethod.ContactCommunicationMethodId = dataModel.GetContactCommunicationMethodID()
	contactCommMethod.CommunicationMethodCode = dataModel.GetCommunicationMethodCode()
	contactCommMethod.CommunicationMethodLabelCode = dataModel.GetCommunicationMethodLabelCode()
	contactCommMethod.CommunicationMethodLabelCaption = dataModel.GetCommunicationMethodLabelCaption()
	contactCommMethod.FormatValue = dataModel.GetFormatValue()
	contactCommMethod.Status = dataModel.GetStatus()
	contactCommMethod.IsDefault = dataModel.GetIsDefault()
	contactCommMethod.GetAudit().CreatedAt, _ = ptypes.TimestampProto(dataModel.GetAudit().GetCreatedAt())
	contactCommMethod.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(dataModel.GetAudit().GetModifiedAt())
	contactCommMethod.GetAudit().Vers = dataModel.GetAudit().GetVers()
	return contactCommMethod
}

func contactCommunicationMethodAPIToModel(data *contactcommunicationmethodapi.ContactCommunicationMethod) *contactcommunicationmethodmodel.ContactCommunicationMethod {
	contactCommMethod := contactcommunicationmethodmodel.NewContactCommunicationMethod()
	contactCommMethod.ContactSystemCode = data.GetContactSystemCode()
	contactCommMethod.ContactID = data.GetContactId()
	contactCommMethod.ContactCommunicationMethodID = data.GetContactCommunicationMethodId()
	contactCommMethod.CommunicationMethodCode = data.GetCommunicationMethodCode()
	contactCommMethod.CommunicationMethodLabelCode = data.GetCommunicationMethodLabelCode()
	contactCommMethod.CommunicationMethodLabelCaption = data.GetCommunicationMethodLabelCaption()
	contactCommMethod.FormatValue = data.GetFormatValue()
	contactCommMethod.Status = data.GetStatus()
	contactCommMethod.IsDefault = data.GetIsDefault()
	contactCommMethod.GetAudit().CreatedAt, _ = ptypes.Timestamp(data.GetAudit().GetCreatedAt())
	contactCommMethod.GetAudit().ModifiedAt, _ = ptypes.Timestamp(data.GetAudit().GetModifiedAt())
	contactCommMethod.GetAudit().Vers = data.GetAudit().GetVers()
	return contactCommMethod
}

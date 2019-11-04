package contactcommunicationmethod

import (
	"context"

	auditapi "github.com/bungysheep/contact-management/pkg/api/v1/audit"
	contactcommunicationmethodapi "github.com/bungysheep/contact-management/pkg/api/v1/contactcommunicationmethod"
	contactcommunicationmethodmodel "github.com/bungysheep/contact-management/pkg/models/v1/contactcommunicationmethod"
	contactcommunicationmethodrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contactcommunicationmethod"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type contactcommunicationmethodService struct {
	repo contactcommunicationmethodrepository.IContactCommunicationMethodRepository
}

// NewContactCommunicationMethodService - Contact Communication Method service implementation
func NewContactCommunicationMethodService(repo contactcommunicationmethodrepository.IContactCommunicationMethodRepository) contactcommunicationmethodapi.ContactCommunicationMethodServiceServer {
	return &contactcommunicationmethodService{repo: repo}
}

func (cmm *contactcommunicationmethodService) DoRead(ctx context.Context, req *contactcommunicationmethodapi.DoReadContactCommunicationMethodRequest) (*contactcommunicationmethodapi.DoReadContactCommunicationMethodResponse, error) {
	result, err := cmm.repo.DoRead(ctx, req.GetContactSystemCode(), req.GetContactId(), req.GetContactCommunicationMethodId())

	return &contactcommunicationmethodapi.DoReadContactCommunicationMethodResponse{ContactCommunicationMethod: contactCommunicationMethodModelToAPI(result)}, err
}

func (cmm *contactcommunicationmethodService) DoReadAll(ctx context.Context, req *contactcommunicationmethodapi.DoReadAllContactCommunicationMethodRequest) (*contactcommunicationmethodapi.DoReadAllContactCommunicationMethodResponse, error) {
	result, err := cmm.repo.DoReadAll(ctx, req.GetContactSystemCode(), req.GetContactId())

	resp := make([]*contactcommunicationmethodapi.ContactCommunicationMethod, 0)

	for _, item := range result {
		resp = append(resp, contactCommunicationMethodModelToAPI(item))
	}

	return &contactcommunicationmethodapi.DoReadAllContactCommunicationMethodResponse{ContactCommunicationMethod: resp}, err
}

func (cmm *contactcommunicationmethodService) DoSave(ctx context.Context, req *contactcommunicationmethodapi.DoSaveContactCommunicationMethodRequest) (*contactcommunicationmethodapi.DoSaveContactCommunicationMethodResponse, error) {
	res, err := doUpdate(ctx, cmm.repo, req)
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return doInsert(ctx, cmm.repo, req)
			}
		}
	}

	return res, err
}

func (cmm *contactcommunicationmethodService) DoDelete(ctx context.Context, req *contactcommunicationmethodapi.DoDeleteContactCommunicationMethodRequest) (*contactcommunicationmethodapi.DoDeleteContactCommunicationMethodResponse, error) {
	err := cmm.repo.DoDelete(ctx, req.GetContactSystemCode(), req.GetContactId(), req.GetContactCommunicationMethodId())

	return &contactcommunicationmethodapi.DoDeleteContactCommunicationMethodResponse{Result: err == nil}, err
}

func doInsert(ctx context.Context, repo contactcommunicationmethodrepository.IContactCommunicationMethodRepository, req *contactcommunicationmethodapi.DoSaveContactCommunicationMethodRequest) (*contactcommunicationmethodapi.DoSaveContactCommunicationMethodResponse, error) {
	err := repo.DoInsert(ctx, contactCommunicationMethodAPIToModel(req.GetContactCommunicationMethod()))

	return &contactcommunicationmethodapi.DoSaveContactCommunicationMethodResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo contactcommunicationmethodrepository.IContactCommunicationMethodRepository, req *contactcommunicationmethodapi.DoSaveContactCommunicationMethodRequest) (*contactcommunicationmethodapi.DoSaveContactCommunicationMethodResponse, error) {
	err := repo.DoUpdate(ctx, contactCommunicationMethodAPIToModel(req.GetContactCommunicationMethod()))

	return &contactcommunicationmethodapi.DoSaveContactCommunicationMethodResponse{Result: err == nil}, err
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
	contactCommMethod.IsDefault = data.GetIsDefault()
	contactCommMethod.GetAudit().CreatedAt, _ = ptypes.Timestamp(data.GetAudit().GetCreatedAt())
	contactCommMethod.GetAudit().ModifiedAt, _ = ptypes.Timestamp(data.GetAudit().GetModifiedAt())
	contactCommMethod.GetAudit().Vers = data.GetAudit().GetVers()
	return contactCommMethod
}

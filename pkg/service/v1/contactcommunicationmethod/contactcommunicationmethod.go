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

	resp := &contactcommunicationmethodapi.ContactCommunicationMethod{Audit: &auditapi.Audit{}}
	resp.ContactSystemCode = result.GetContactSystemCode()
	resp.ContactId = result.GetContactID()
	resp.ContactCommunicationMethodId = result.GetContactCommunicationMethodID()
	resp.CommunicationMethodCode = result.GetCommunicationMethodCode()
	resp.CommunicationMethodLabelCode = result.GetCommunicationMethodLabelCode()
	resp.CommunicationMethodLabelCaption = result.GetCommunicationMethodLabelCaption()
	resp.FormatValue = result.GetFormatValue()
	resp.IsDefault = result.GetIsDefault()
	resp.GetAudit().CreatedAt, _ = ptypes.TimestampProto(result.GetAudit().GetCreatedAt())
	resp.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(result.GetAudit().GetModifiedAt())
	resp.GetAudit().Vers = result.GetAudit().GetVers()

	return &contactcommunicationmethodapi.DoReadContactCommunicationMethodResponse{ContactCommunicationMethod: resp}, err
}

func (cmm *contactcommunicationmethodService) DoReadAll(ctx context.Context, req *contactcommunicationmethodapi.DoReadAllContactCommunicationMethodRequest) (*contactcommunicationmethodapi.DoReadAllContactCommunicationMethodResponse, error) {
	result, err := cmm.repo.DoReadAll(ctx, req.GetContactSystemCode(), req.GetContactId())

	resp := make([]*contactcommunicationmethodapi.ContactCommunicationMethod, 0)

	for _, item := range result {
		contactCommMethod := &contactcommunicationmethodapi.ContactCommunicationMethod{Audit: &auditapi.Audit{}}
		contactCommMethod.ContactSystemCode = item.GetContactSystemCode()
		contactCommMethod.ContactId = item.GetContactID()
		contactCommMethod.ContactCommunicationMethodId = item.GetContactCommunicationMethodID()
		contactCommMethod.CommunicationMethodCode = item.GetCommunicationMethodCode()
		contactCommMethod.CommunicationMethodLabelCode = item.GetCommunicationMethodLabelCode()
		contactCommMethod.CommunicationMethodLabelCaption = item.GetCommunicationMethodLabelCaption()
		contactCommMethod.FormatValue = item.GetFormatValue()
		contactCommMethod.IsDefault = item.GetIsDefault()
		contactCommMethod.GetAudit().CreatedAt, _ = ptypes.TimestampProto(item.GetAudit().GetCreatedAt())
		contactCommMethod.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(item.GetAudit().GetModifiedAt())
		contactCommMethod.GetAudit().Vers = item.GetAudit().GetVers()

		resp = append(resp, contactCommMethod)
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
	contactCommMethod := contactcommunicationmethodmodel.NewContactCommunicationMethod()
	contactCommMethod.ContactSystemCode = req.GetContactCommunicationMethod().GetContactSystemCode()
	contactCommMethod.ContactID = req.GetContactCommunicationMethod().GetContactId()
	contactCommMethod.ContactCommunicationMethodID = req.GetContactCommunicationMethod().GetContactCommunicationMethodId()
	contactCommMethod.CommunicationMethodCode = req.GetContactCommunicationMethod().GetCommunicationMethodCode()
	contactCommMethod.CommunicationMethodLabelCode = req.GetContactCommunicationMethod().GetCommunicationMethodLabelCode()
	contactCommMethod.CommunicationMethodLabelCaption = req.GetContactCommunicationMethod().GetCommunicationMethodLabelCaption()
	contactCommMethod.FormatValue = req.GetContactCommunicationMethod().GetFormatValue()
	contactCommMethod.IsDefault = req.GetContactCommunicationMethod().GetIsDefault()
	contactCommMethod.GetAudit().CreatedAt, _ = ptypes.Timestamp(req.GetContactCommunicationMethod().GetAudit().GetCreatedAt())
	contactCommMethod.GetAudit().ModifiedAt, _ = ptypes.Timestamp(req.GetContactCommunicationMethod().GetAudit().GetModifiedAt())
	contactCommMethod.GetAudit().Vers = req.GetContactCommunicationMethod().GetAudit().GetVers()

	err := repo.DoInsert(ctx, contactCommMethod)

	return &contactcommunicationmethodapi.DoSaveContactCommunicationMethodResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo contactcommunicationmethodrepository.IContactCommunicationMethodRepository, req *contactcommunicationmethodapi.DoSaveContactCommunicationMethodRequest) (*contactcommunicationmethodapi.DoSaveContactCommunicationMethodResponse, error) {
	contactCommMethod := contactcommunicationmethodmodel.NewContactCommunicationMethod()
	contactCommMethod.ContactSystemCode = req.GetContactCommunicationMethod().GetContactSystemCode()
	contactCommMethod.ContactID = req.GetContactCommunicationMethod().GetContactId()
	contactCommMethod.ContactCommunicationMethodID = req.GetContactCommunicationMethod().GetContactCommunicationMethodId()
	contactCommMethod.CommunicationMethodCode = req.GetContactCommunicationMethod().GetCommunicationMethodCode()
	contactCommMethod.CommunicationMethodLabelCode = req.GetContactCommunicationMethod().GetCommunicationMethodLabelCode()
	contactCommMethod.CommunicationMethodLabelCaption = req.GetContactCommunicationMethod().GetCommunicationMethodLabelCaption()
	contactCommMethod.FormatValue = req.GetContactCommunicationMethod().GetFormatValue()
	contactCommMethod.IsDefault = req.GetContactCommunicationMethod().GetIsDefault()
	contactCommMethod.GetAudit().CreatedAt, _ = ptypes.Timestamp(req.GetContactCommunicationMethod().GetAudit().GetCreatedAt())
	contactCommMethod.GetAudit().ModifiedAt, _ = ptypes.Timestamp(req.GetContactCommunicationMethod().GetAudit().GetModifiedAt())
	contactCommMethod.GetAudit().Vers = req.GetContactCommunicationMethod().GetAudit().GetVers()

	err := repo.DoUpdate(ctx, contactCommMethod)

	return &contactcommunicationmethodapi.DoSaveContactCommunicationMethodResponse{Result: err == nil}, err
}

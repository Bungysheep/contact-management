package contactsystem

import (
	"context"

	auditapi "github.com/bungysheep/contact-management/pkg/api/v1/audit"
	contactsystemapi "github.com/bungysheep/contact-management/pkg/api/v1/contactsystem"
	contactsystemmodel "github.com/bungysheep/contact-management/pkg/models/v1/contactsystem"
	contactsystemrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contactsystem"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type contactSystemService struct {
	repo contactsystemrepository.IContactSystemRepository
}

// NewContactSystemService - Contact System service implementation
func NewContactSystemService(repo contactsystemrepository.IContactSystemRepository) contactsystemapi.ContactSystemServiceServer {
	return &contactSystemService{repo: repo}
}

func (cntsys *contactSystemService) DoRead(ctx context.Context, req *contactsystemapi.DoReadContactSystemRequest) (*contactsystemapi.DoReadContactSystemResponse, error) {
	result, err := cntsys.repo.DoRead(ctx, req.GetContactSystemCode())

	resp := &contactsystemapi.ContactSystem{Audit: &auditapi.Audit{}}
	resp.ContactSystemCode = result.GetContactSystemCode()
	resp.Description = result.GetDescription()
	resp.Details = result.GetDetails()
	resp.Status = result.GetStatus()
	resp.GetAudit().CreatedAt, _ = ptypes.TimestampProto(result.GetAudit().GetCreatedAt())
	resp.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(result.GetAudit().GetModifiedAt())
	resp.GetAudit().Vers = result.GetAudit().GetVers()

	return &contactsystemapi.DoReadContactSystemResponse{ContactSystem: resp}, err
}

func (cntsys *contactSystemService) DoReadAll(ctx context.Context, req *contactsystemapi.DoReadAllContactSystemRequest) (*contactsystemapi.DoReadAllContactSystemResponse, error) {
	result, err := cntsys.repo.DoReadAll(ctx)

	resp := make([]*contactsystemapi.ContactSystem, 0)

	for _, item := range result {
		contactSystem := &contactsystemapi.ContactSystem{Audit: &auditapi.Audit{}}
		contactSystem.ContactSystemCode = item.GetContactSystemCode()
		contactSystem.Description = item.GetDescription()
		contactSystem.Details = item.GetDetails()
		contactSystem.Status = item.GetStatus()
		contactSystem.GetAudit().CreatedAt, _ = ptypes.TimestampProto(item.GetAudit().GetCreatedAt())
		contactSystem.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(item.GetAudit().GetModifiedAt())
		contactSystem.GetAudit().Vers = item.GetAudit().GetVers()

		resp = append(resp, contactSystem)
	}

	return &contactsystemapi.DoReadAllContactSystemResponse{ContactSystems: resp}, err
}

func (cntsys *contactSystemService) DoSave(ctx context.Context, req *contactsystemapi.DoSaveContactSystemRequest) (*contactsystemapi.DoSaveContactSystemResponse, error) {
	res, err := doUpdate(ctx, cntsys.repo, req)
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return doInsert(ctx, cntsys.repo, req)
			}
		}
	}

	return res, err
}

func (cntsys *contactSystemService) DoDelete(ctx context.Context, req *contactsystemapi.DoDeleteContactSystemRequest) (*contactsystemapi.DoDeleteContactSystemResponse, error) {
	err := cntsys.repo.DoDelete(ctx, req.GetContactSystemCode())

	return &contactsystemapi.DoDeleteContactSystemResponse{Result: err == nil}, err
}

func doInsert(ctx context.Context, repo contactsystemrepository.IContactSystemRepository, req *contactsystemapi.DoSaveContactSystemRequest) (*contactsystemapi.DoSaveContactSystemResponse, error) {
	contactSystem := contactsystemmodel.NewContactSystem()
	contactSystem.ContactSystemCode = req.GetContactSystem().GetContactSystemCode()
	contactSystem.Description = req.GetContactSystem().GetDescription()
	contactSystem.Details = req.GetContactSystem().GetDetails()
	contactSystem.Status = req.GetContactSystem().GetStatus()
	contactSystem.GetAudit().CreatedAt, _ = ptypes.Timestamp(req.GetContactSystem().GetAudit().GetCreatedAt())
	contactSystem.GetAudit().ModifiedAt, _ = ptypes.Timestamp(req.GetContactSystem().GetAudit().GetModifiedAt())
	contactSystem.GetAudit().Vers = req.GetContactSystem().GetAudit().GetVers()

	err := repo.DoInsert(ctx, contactSystem)

	return &contactsystemapi.DoSaveContactSystemResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo contactsystemrepository.IContactSystemRepository, req *contactsystemapi.DoSaveContactSystemRequest) (*contactsystemapi.DoSaveContactSystemResponse, error) {
	contactSystem := contactsystemmodel.NewContactSystem()
	contactSystem.ContactSystemCode = req.GetContactSystem().GetContactSystemCode()
	contactSystem.Description = req.GetContactSystem().GetDescription()
	contactSystem.Details = req.GetContactSystem().GetDetails()
	contactSystem.Status = req.GetContactSystem().GetStatus()
	contactSystem.GetAudit().CreatedAt, _ = ptypes.Timestamp(req.GetContactSystem().GetAudit().GetCreatedAt())
	contactSystem.GetAudit().ModifiedAt, _ = ptypes.Timestamp(req.GetContactSystem().GetAudit().GetModifiedAt())
	contactSystem.GetAudit().Vers = req.GetContactSystem().GetAudit().GetVers()

	err := repo.DoUpdate(ctx, contactSystem)

	return &contactsystemapi.DoSaveContactSystemResponse{Result: err == nil}, err
}

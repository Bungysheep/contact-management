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

	return &contactsystemapi.DoReadContactSystemResponse{ContactSystem: contactSystemModelToAPI(result)}, err
}

func (cntsys *contactSystemService) DoReadAll(ctx context.Context, req *contactsystemapi.DoReadAllContactSystemRequest) (*contactsystemapi.DoReadAllContactSystemResponse, error) {
	result, err := cntsys.repo.DoReadAll(ctx)

	resp := make([]*contactsystemapi.ContactSystem, 0)

	for _, item := range result {
		resp = append(resp, contactSystemModelToAPI(item))
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
	err := repo.DoInsert(ctx, contactSystemAPIToModel(req.GetContactSystem()))

	return &contactsystemapi.DoSaveContactSystemResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo contactsystemrepository.IContactSystemRepository, req *contactsystemapi.DoSaveContactSystemRequest) (*contactsystemapi.DoSaveContactSystemResponse, error) {
	err := repo.DoUpdate(ctx, contactSystemAPIToModel(req.GetContactSystem()))

	return &contactsystemapi.DoSaveContactSystemResponse{Result: err == nil}, err
}

func contactSystemModelToAPI(dataModel *contactsystemmodel.ContactSystem) *contactsystemapi.ContactSystem {
	contactSystem := &contactsystemapi.ContactSystem{Audit: &auditapi.Audit{}}
	contactSystem.ContactSystemCode = dataModel.GetContactSystemCode()
	contactSystem.Description = dataModel.GetDescription()
	contactSystem.Details = dataModel.GetDetails()
	contactSystem.Status = dataModel.GetStatus()
	contactSystem.GetAudit().CreatedAt, _ = ptypes.TimestampProto(dataModel.GetAudit().GetCreatedAt())
	contactSystem.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(dataModel.GetAudit().GetModifiedAt())
	contactSystem.GetAudit().Vers = dataModel.GetAudit().GetVers()
	return contactSystem
}

func contactSystemAPIToModel(data *contactsystemapi.ContactSystem) *contactsystemmodel.ContactSystem {
	contactSystem := contactsystemmodel.NewContactSystem()
	contactSystem.ContactSystemCode = data.GetContactSystemCode()
	contactSystem.Description = data.GetDescription()
	contactSystem.Details = data.GetDetails()
	contactSystem.Status = data.GetStatus()
	contactSystem.GetAudit().CreatedAt, _ = ptypes.Timestamp(data.GetAudit().GetCreatedAt())
	contactSystem.GetAudit().ModifiedAt, _ = ptypes.Timestamp(data.GetAudit().GetModifiedAt())
	contactSystem.GetAudit().Vers = data.GetAudit().GetVers()
	return contactSystem
}

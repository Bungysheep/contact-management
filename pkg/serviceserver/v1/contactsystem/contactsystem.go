package contactsystem

import (
	"context"

	auditapi "github.com/bungysheep/contact-management/pkg/api/v1/audit"
	contactsystemapi "github.com/bungysheep/contact-management/pkg/api/v1/contactsystem"
	contactsystemmodel "github.com/bungysheep/contact-management/pkg/models/v1/contactsystem"
	contactsystemservice "github.com/bungysheep/contact-management/pkg/service/v1/contactsystem"
	"github.com/golang/protobuf/ptypes"
)

type contactSystemServiceServer struct {
	svc contactsystemservice.IContactSystemService
}

// NewContactSystemServiceServer - Contact System service server implementation
func NewContactSystemServiceServer(svc contactsystemservice.IContactSystemService) contactsystemapi.ContactSystemServiceServer {
	return &contactSystemServiceServer{svc: svc}
}

func (cntsys *contactSystemServiceServer) DoRead(ctx context.Context, req *contactsystemapi.DoReadContactSystemRequest) (*contactsystemapi.DoReadContactSystemResponse, error) {
	result, err := cntsys.svc.DoRead(ctx, req.GetContactSystemCode())

	return &contactsystemapi.DoReadContactSystemResponse{ContactSystem: contactSystemModelToAPI(result)}, err
}

func (cntsys *contactSystemServiceServer) DoReadAll(ctx context.Context, req *contactsystemapi.DoReadAllContactSystemRequest) (*contactsystemapi.DoReadAllContactSystemResponse, error) {
	result, err := cntsys.svc.DoReadAll(ctx)

	resp := make([]*contactsystemapi.ContactSystem, 0)

	for _, item := range result {
		resp = append(resp, contactSystemModelToAPI(item))
	}

	return &contactsystemapi.DoReadAllContactSystemResponse{ContactSystems: resp}, err
}

func (cntsys *contactSystemServiceServer) DoSave(ctx context.Context, req *contactsystemapi.DoSaveContactSystemRequest) (*contactsystemapi.DoSaveContactSystemResponse, error) {
	err := cntsys.svc.DoSave(ctx, contactSystemAPIToModel(req.GetContactSystem()))

	return &contactsystemapi.DoSaveContactSystemResponse{Result: err == nil}, err
}

func (cntsys *contactSystemServiceServer) DoDelete(ctx context.Context, req *contactsystemapi.DoDeleteContactSystemRequest) (*contactsystemapi.DoDeleteContactSystemResponse, error) {
	err := cntsys.svc.DoDelete(ctx, req.GetContactSystemCode())

	return &contactsystemapi.DoDeleteContactSystemResponse{Result: err == nil}, err
}

func contactSystemModelToAPI(dataModel *contactsystemmodel.ContactSystem) *contactsystemapi.ContactSystem {
	if dataModel == nil {
		return nil
	}

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
	if data == nil {
		return nil
	}

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

package contactsystem

import (
	"context"
	"os"
	"testing"
	"time"

	auditapi "github.com/bungysheep/contact-management/pkg/api/v1/audit"
	contactsystemapi "github.com/bungysheep/contact-management/pkg/api/v1/contactsystem"
	"github.com/bungysheep/contact-management/pkg/common/message"
	auditmodel "github.com/bungysheep/contact-management/pkg/models/v1/audit"
	contactsystemmodel "github.com/bungysheep/contact-management/pkg/models/v1/contactsystem"
	"github.com/bungysheep/contact-management/pkg/repository/v1/contactsystem/mock_contactsystem"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ctx  context.Context
	data []*contactsystemmodel.ContactSystem
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	tmNow := time.Now().In(time.UTC)

	data = append(data, &contactsystemmodel.ContactSystem{
		ContactSystemCode: "CNTSYS001",
		Description:       "Contact System 1",
		Details:           "Contact System 1",
		Status:            "A",
		Audit: &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		},
	}, &contactsystemmodel.ContactSystem{
		ContactSystemCode: "CNTSYS002",
		Description:       "Contact System 2",
		Details:           "Contact System 2",
		Status:            "A",
		Audit: &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		},
	}, &contactsystemmodel.ContactSystem{
		ContactSystemCode: "CNTSYS003",
		Description:       "Contact System 3",
		Details:           "Contact System 3",
		Status:            "A",
		Audit: &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		},
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestContactSystemService(t *testing.T) {
	t.Run("DoRead Contact System", doRead(ctx, data[0]))

	t.Run("DoReadAll Contact System", doReadAll(ctx, data[0]))

	t.Run("DoSave new Contact System", doSaveNew(ctx, data[0]))

	t.Run("DoSave existing Contact System", doSaveExisting(ctx, data[0]))

	t.Run("DoDelete Contact System", doDelete(ctx, data[0]))
}

func doRead(ctx context.Context, input *contactsystemmodel.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_contactsystem.NewMockIContactSystemRepository(ctl)

		repo.EXPECT().DoRead(ctx, input.GetContactSystemCode()).Return(input, nil)

		svc := NewContactSystemService(repo)

		resp, err := svc.DoRead(ctx, &contactsystemapi.DoReadContactSystemRequest{ContactSystemCode: input.GetContactSystemCode()})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if resp.GetContactSystem() == nil {
			t.Errorf("Expect contact system is not nil")
		}

		if resp.GetContactSystem().GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetContactSystem().GetContactSystemCode())
		}

		if resp.GetContactSystem().GetDescription() != input.GetDescription() {
			t.Errorf("Expect description %s, but got %s", input.GetDescription(), resp.GetContactSystem().GetDescription())
		}

		if resp.GetContactSystem().GetDetails() != input.GetDetails() {
			t.Errorf("Expect details %s, but got %s", input.GetDetails(), resp.GetContactSystem().GetDetails())
		}

		if resp.GetContactSystem().GetStatus() != input.GetStatus() {
			t.Errorf("Expect status %s, but got %s", input.GetStatus(), resp.GetContactSystem().GetStatus())
		}
	}
}

func doReadAll(ctx context.Context, input *contactsystemmodel.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_contactsystem.NewMockIContactSystemRepository(ctl)

		repo.EXPECT().DoReadAll(ctx).Return(data, nil)

		svc := NewContactSystemService(repo)

		resp, err := svc.DoReadAll(ctx, &contactsystemapi.DoReadAllContactSystemRequest{})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if resp.GetContactSystems() == nil {
			t.Errorf("Expect contact system is not nil")
		}

		if len(resp.GetContactSystems()) < 3 {
			t.Errorf("Expect there are contact systems retrieved")
		}

		if resp.GetContactSystems()[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetContactSystems()[0].GetContactSystemCode())
		}

		if resp.GetContactSystems()[0].GetDescription() != input.GetDescription() {
			t.Errorf("Expect description %s, but got %s", input.GetDescription(), resp.GetContactSystems()[0].GetDescription())
		}

		if resp.GetContactSystems()[0].GetDetails() != input.GetDetails() {
			t.Errorf("Expect details %s, but got %s", input.GetDetails(), resp.GetContactSystems()[0].GetDetails())
		}

		if resp.GetContactSystems()[0].GetStatus() != input.GetStatus() {
			t.Errorf("Expect status %s, but got %s", input.GetStatus(), resp.GetContactSystems()[0].GetStatus())
		}
	}
}

func doSaveNew(ctx context.Context, input *contactsystemmodel.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_contactsystem.NewMockIContactSystemRepository(ctl)

		contactSystem := &contactsystemapi.ContactSystem{Audit: &auditapi.Audit{}}
		contactSystem.ContactSystemCode = input.GetContactSystemCode()
		contactSystem.Description = input.GetDescription()
		contactSystem.Details = input.GetDetails()
		contactSystem.Status = input.GetStatus()
		contactSystem.GetAudit().CreatedAt, _ = ptypes.TimestampProto(input.GetAudit().GetCreatedAt())
		contactSystem.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(input.GetAudit().GetModifiedAt())
		contactSystem.GetAudit().Vers = input.GetAudit().GetVers()

		repo.EXPECT().DoUpdate(ctx, input).Return(status.Errorf(codes.NotFound, message.DoesNotExist("Contact System")))

		repo.EXPECT().DoInsert(ctx, input).Return(nil)

		svc := NewContactSystemService(repo)

		resp, err := svc.DoSave(ctx, &contactsystemapi.DoSaveContactSystemRequest{ContactSystem: contactSystem})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

func doSaveExisting(ctx context.Context, input *contactsystemmodel.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_contactsystem.NewMockIContactSystemRepository(ctl)

		contactSystem := &contactsystemapi.ContactSystem{Audit: &auditapi.Audit{}}
		contactSystem.ContactSystemCode = input.GetContactSystemCode()
		contactSystem.Description = input.GetDescription()
		contactSystem.Details = input.GetDetails()
		contactSystem.Status = input.GetStatus()
		contactSystem.GetAudit().CreatedAt, _ = ptypes.TimestampProto(input.GetAudit().GetCreatedAt())
		contactSystem.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(input.GetAudit().GetModifiedAt())
		contactSystem.GetAudit().Vers = input.GetAudit().GetVers()

		repo.EXPECT().DoUpdate(ctx, input).Return(nil)

		svc := NewContactSystemService(repo)

		resp, err := svc.DoSave(ctx, &contactsystemapi.DoSaveContactSystemRequest{ContactSystem: contactSystem})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

func doDelete(ctx context.Context, input *contactsystemmodel.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_contactsystem.NewMockIContactSystemRepository(ctl)

		repo.EXPECT().DoDelete(ctx, input.GetContactSystemCode()).Return(nil)

		svc := NewContactSystemService(repo)

		resp, err := svc.DoDelete(ctx, &contactsystemapi.DoDeleteContactSystemRequest{ContactSystemCode: input.GetContactSystemCode()})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

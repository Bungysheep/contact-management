package contact

import (
	"context"
	"os"
	"testing"
	"time"

	auditapi "github.com/bungysheep/contact-management/pkg/api/v1/audit"
	contactapi "github.com/bungysheep/contact-management/pkg/api/v1/contact"
	auditmodel "github.com/bungysheep/contact-management/pkg/models/v1/audit"
	contactmodel "github.com/bungysheep/contact-management/pkg/models/v1/contact"
	"github.com/bungysheep/contact-management/pkg/service/v1/contact/mock_contact"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
)

var (
	ctx  context.Context
	data []*contactmodel.Contact
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	tmNow := time.Now().In(time.UTC)

	data = append(data, &contactmodel.Contact{
		ContactSystemCode: "CNTSYS001",
		ContactID:         1,
		FirstName:         "James",
		LastName:          "Embongbulan",
		Status:            "A",
		Audit: &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		},
	}, &contactmodel.Contact{
		ContactSystemCode: "CNTSYS001",
		ContactID:         2,
		FirstName:         "Rindi",
		LastName:          "Allorerung",
		Status:            "A",
		Audit: &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		},
	}, &contactmodel.Contact{
		ContactSystemCode: "CNTSYS001",
		ContactID:         3,
		FirstName:         "Marvel",
		LastName:          "Embongbulan",
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

func TestContactService(t *testing.T) {
	t.Run("DoRead Communication Method", doRead(ctx, data[0]))

	t.Run("DoReadAll Communication Method", doReadAll(ctx, data[0]))

	t.Run("DoSave Communication Method", doSave(ctx, data[0]))

	t.Run("DoDelete Communication Method", doDelete(ctx, data[0]))
}

func doRead(ctx context.Context, input *contactmodel.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		svc := mock_contact.NewMockIContactService(ctl)

		svc.EXPECT().DoRead(ctx, input.GetContactSystemCode(), input.GetContactID()).Return(input, nil)

		svcServer := NewContactServiceServer(svc)

		resp, err := svcServer.DoRead(ctx, &contactapi.DoReadContactRequest{ContactSystemCode: input.GetContactSystemCode(), ContactId: input.GetContactID()})
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}

		if resp.GetContact() == nil {
			t.Errorf("Expect contact is not nil")
		}

		if resp.GetContact().GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetContact().GetContactSystemCode())
		}

		if resp.GetContact().GetContactId() != input.GetContactID() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactID(), resp.GetContact().GetContactId())
		}

		if resp.GetContact().GetFirstName() != input.GetFirstName() {
			t.Errorf("Expect firstname %s, but got %s", input.GetFirstName(), resp.GetContact().GetFirstName())
		}

		if resp.GetContact().GetLastName() != input.GetLastName() {
			t.Errorf("Expect lastname %s, but got %s", input.GetLastName(), resp.GetContact().GetLastName())
		}

		if resp.GetContact().GetStatus() != input.GetStatus() {
			t.Errorf("Expect status %s, but got %s", input.GetStatus(), resp.GetContact().GetStatus())
		}
	}
}

func doReadAll(ctx context.Context, input *contactmodel.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		svc := mock_contact.NewMockIContactService(ctl)

		svc.EXPECT().DoReadAll(ctx, input.GetContactSystemCode()).Return(data, nil)

		svcServer := NewContactServiceServer(svc)

		resp, err := svcServer.DoReadAll(ctx, &contactapi.DoReadAllContactRequest{ContactSystemCode: input.GetContactSystemCode()})
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}

		if resp.GetContact() == nil {
			t.Errorf("Expect contact is not nil")
		}

		if len(resp.GetContact()) < 3 {
			t.Errorf("Expect there are contacts retrieved")
		}

		if resp.GetContact()[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetContact()[0].GetContactSystemCode())
		}

		if resp.GetContact()[0].GetContactId() != input.GetContactID() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactID(), resp.GetContact()[0].GetContactId())
		}

		if resp.GetContact()[0].GetFirstName() != input.GetFirstName() {
			t.Errorf("Expect firstname %s, but got %s", input.GetFirstName(), resp.GetContact()[0].GetFirstName())
		}

		if resp.GetContact()[0].GetLastName() != input.GetLastName() {
			t.Errorf("Expect lastname %s, but got %s", input.GetLastName(), resp.GetContact()[0].GetLastName())
		}

		if resp.GetContact()[0].GetStatus() != input.GetStatus() {
			t.Errorf("Expect status %s, but got %s", input.GetStatus(), resp.GetContact()[0].GetStatus())
		}
	}
}

func doSave(ctx context.Context, input *contactmodel.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		svc := mock_contact.NewMockIContactService(ctl)

		contact := &contactapi.Contact{Audit: &auditapi.Audit{}}
		contact.ContactSystemCode = input.GetContactSystemCode()
		contact.ContactId = input.GetContactID()
		contact.FirstName = input.GetFirstName()
		contact.LastName = input.GetLastName()
		contact.Status = input.GetStatus()
		contact.GetAudit().CreatedAt, _ = ptypes.TimestampProto(input.GetAudit().GetCreatedAt())
		contact.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(input.GetAudit().GetModifiedAt())
		contact.GetAudit().Vers = input.GetAudit().GetVers()

		svc.EXPECT().DoSave(ctx, input).Return(nil)

		svcServer := NewContactServiceServer(svc)

		resp, err := svcServer.DoSave(ctx, &contactapi.DoSaveContactRequest{Contact: contact})
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

func doDelete(ctx context.Context, input *contactmodel.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		svc := mock_contact.NewMockIContactService(ctl)

		svc.EXPECT().DoDelete(ctx, input.GetContactSystemCode(), input.GetContactID()).Return(nil)

		svcServer := NewContactServiceServer(svc)

		resp, err := svcServer.DoDelete(ctx, &contactapi.DoDeleteContactRequest{ContactSystemCode: input.GetContactSystemCode(), ContactId: input.GetContactID()})
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

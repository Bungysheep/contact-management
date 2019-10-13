package contact

import (
	"context"
	"os"
	"testing"

	"github.com/bungysheep/contact-management/pkg/api/v1/contact"
	"github.com/bungysheep/contact-management/pkg/common/message"
	"github.com/bungysheep/contact-management/pkg/repository/v1/contact/mock_contact"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ctx  context.Context
	data []*contact.Contact
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	data = append(data, &contact.Contact{
		ContactSystemCode: "CNTSYS001",
		ContactId:         1,
		FirstName:         "James",
		LastName:          "Embongbulan",
		Status:            "A",
	}, &contact.Contact{
		ContactSystemCode: "CNTSYS001",
		ContactId:         2,
		FirstName:         "Rindi",
		LastName:          "Allorerung",
		Status:            "A",
	}, &contact.Contact{
		ContactSystemCode: "CNTSYS001",
		ContactId:         3,
		FirstName:         "Marvel",
		LastName:          "Embongbulan",
		Status:            "A",
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestContactService(t *testing.T) {
	t.Run("DoRead Communication Method", doRead(ctx, data[0]))

	t.Run("DoReadAll Communication Method", doReadAll(ctx, data[0]))

	t.Run("DoSave new Communication Method", doSaveNew(ctx, data[0]))

	t.Run("DoSave existing Communication Method", doSaveExisting(ctx, data[0]))

	t.Run("DoDelete Communication Method", doDelete(ctx, data[0]))
}

func doRead(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_contact.NewMockIContactRepository(ctl)

		repo.EXPECT().DoRead(ctx, input.GetContactSystemCode(), input.GetContactId()).Return(input, nil)

		svc := NewContactService(repo)

		resp, err := svc.DoRead(ctx, &contact.DoReadRequest{ContactSystemCode: input.GetContactSystemCode(), ContactId: input.GetContactId()})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if resp == nil {
			t.Errorf("Expect contact is not nil")
		}

		if resp.GetContact().GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetContact().GetContactSystemCode())
		}

		if resp.GetContact().GetContactId() != input.GetContactId() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactId(), resp.GetContact().GetContactId())
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

func doReadAll(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_contact.NewMockIContactRepository(ctl)

		repo.EXPECT().DoReadAll(ctx).Return(data, nil)

		svc := NewContactService(repo)

		resp, err := svc.DoReadAll(ctx, &contact.DoReadAllRequest{})
		if err != nil {
			t.Errorf("Expect error is nil")
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

		if resp.GetContact()[0].GetContactId() != input.GetContactId() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactId(), resp.GetContact()[0].GetContactId())
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

func doSaveNew(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_contact.NewMockIContactRepository(ctl)

		repo.EXPECT().DoUpdate(ctx, input).Return(status.Errorf(codes.NotFound, message.DoesNotExist("Communication Method")))

		repo.EXPECT().DoInsert(ctx, input).Return(nil)

		svc := NewContactService(repo)

		resp, err := svc.DoSave(ctx, &contact.DoSaveRequest{Contact: input})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

func doSaveExisting(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_contact.NewMockIContactRepository(ctl)

		repo.EXPECT().DoUpdate(ctx, input).Return(nil)

		svc := NewContactService(repo)

		resp, err := svc.DoSave(ctx, &contact.DoSaveRequest{Contact: input})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

func doDelete(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_contact.NewMockIContactRepository(ctl)

		repo.EXPECT().DoDelete(ctx, input.GetContactSystemCode(), input.GetContactId()).Return(nil)

		svc := NewContactService(repo)

		resp, err := svc.DoDelete(ctx, &contact.DoDeleteRequest{ContactSystemCode: input.GetContactSystemCode(), ContactId: input.GetContactId()})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

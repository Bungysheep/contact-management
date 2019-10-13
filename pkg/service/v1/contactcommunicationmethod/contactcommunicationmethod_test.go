package contactcommunicationmethod

import (
	"context"
	"os"
	"testing"

	"github.com/bungysheep/contact-management/pkg/api/v1/contactcommunicationmethod"
	"github.com/bungysheep/contact-management/pkg/common/message"
	"github.com/bungysheep/contact-management/pkg/repository/v1/contactcommunicationmethod/mock_contactcommunicationmethod"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ctx  context.Context
	data []*contactcommunicationmethod.ContactCommunicationMethod
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	data = append(data, &contactcommunicationmethod.ContactCommunicationMethod{
		ContactSystemCode:            "CNTSYS001",
		ContactId:                    1,
		ContactCommunicationMethodId: 1,
		CommunicationMethodCode:      "EMAIL",
	}, &contactcommunicationmethod.ContactCommunicationMethod{
		ContactSystemCode:            "CNTSYS001",
		ContactId:                    1,
		ContactCommunicationMethodId: 2,
		CommunicationMethodCode:      "MOBILE",
	}, &contactcommunicationmethod.ContactCommunicationMethod{
		ContactSystemCode:            "CNTSYS001",
		ContactId:                    1,
		ContactCommunicationMethodId: 3,
		CommunicationMethodCode:      "FAX",
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestContactCommunicationMethodService(t *testing.T) {
	t.Run("DoRead Contact Communication Method", doRead(ctx, data[0]))

	t.Run("DoReadAll Contact Communication Method", doReadAll(ctx, data[0]))

	t.Run("DoSave new Contact Communication Method", doSaveNew(ctx, data[0]))

	t.Run("DoSave existing Contact Communication Method", doSaveExisting(ctx, data[0]))

	t.Run("DoDelete Contact Communication Method", doDelete(ctx, data[0]))
}

func doRead(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_contactcommunicationmethod.NewMockIContactCommunicationMethodRepository(ctl)

		repo.EXPECT().DoRead(ctx, input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId()).Return(input, nil)

		svc := NewContactCommunicationMethodService(repo)

		resp, err := svc.DoRead(ctx, &contactcommunicationmethod.DoReadRequest{ContactSystemCode: input.GetContactSystemCode(), ContactId: input.GetContactId(), ContactCommunicationMethodId: input.GetContactCommunicationMethodId()})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if resp == nil {
			t.Errorf("Expect communication method is not nil")
		}

		if resp.GetCommunicationMethod().GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetCommunicationMethod().GetContactSystemCode())
		}

		if resp.GetCommunicationMethod().GetContactId() != input.GetContactId() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactId(), resp.GetCommunicationMethod().GetContactId())
		}

		if resp.GetCommunicationMethod().GetContactCommunicationMethodId() != input.GetContactCommunicationMethodId() {
			t.Errorf("Expect contact communication method id %d, but got %d", input.GetContactCommunicationMethodId(), resp.GetCommunicationMethod().GetContactCommunicationMethodId())
		}

		if resp.GetCommunicationMethod().GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp.GetCommunicationMethod().GetCommunicationMethodCode())
		}
	}
}

func doReadAll(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_contactcommunicationmethod.NewMockIContactCommunicationMethodRepository(ctl)

		repo.EXPECT().DoReadAll(ctx).Return(data, nil)

		svc := NewContactCommunicationMethodService(repo)

		resp, err := svc.DoReadAll(ctx, &contactcommunicationmethod.DoReadAllRequest{})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if resp.GetCommunicationMethod() == nil {
			t.Errorf("Expect communication method is not nil")
		}

		if len(resp.GetCommunicationMethod()) < 3 {
			t.Errorf("Expect there are communication methods retrieved")
		}

		if resp.GetCommunicationMethod()[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetCommunicationMethod()[0].GetContactSystemCode())
		}

		if resp.GetCommunicationMethod()[0].GetContactId() != input.GetContactId() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactId(), resp.GetCommunicationMethod()[0].GetContactId())
		}

		if resp.GetCommunicationMethod()[0].GetContactCommunicationMethodId() != input.GetContactCommunicationMethodId() {
			t.Errorf("Expect contact communication method id %d, but got %d", input.GetContactCommunicationMethodId(), resp.GetCommunicationMethod()[0].GetContactCommunicationMethodId())
		}

		if resp.GetCommunicationMethod()[0].GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp.GetCommunicationMethod()[0].GetCommunicationMethodCode())
		}
	}
}

func doSaveNew(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_contactcommunicationmethod.NewMockIContactCommunicationMethodRepository(ctl)

		repo.EXPECT().DoUpdate(ctx, input).Return(status.Errorf(codes.NotFound, message.DoesNotExist("Contact Communication Method")))

		repo.EXPECT().DoInsert(ctx, input).Return(nil)

		svc := NewContactCommunicationMethodService(repo)

		resp, err := svc.DoSave(ctx, &contactcommunicationmethod.DoSaveRequest{CommunicationMethod: input})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

func doSaveExisting(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_contactcommunicationmethod.NewMockIContactCommunicationMethodRepository(ctl)

		repo.EXPECT().DoUpdate(ctx, input).Return(nil)

		svc := NewContactCommunicationMethodService(repo)

		resp, err := svc.DoSave(ctx, &contactcommunicationmethod.DoSaveRequest{CommunicationMethod: input})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

func doDelete(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_contactcommunicationmethod.NewMockIContactCommunicationMethodRepository(ctl)

		repo.EXPECT().DoDelete(ctx, input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId()).Return(nil)

		svc := NewContactCommunicationMethodService(repo)

		resp, err := svc.DoDelete(ctx, &contactcommunicationmethod.DoDeleteRequest{ContactSystemCode: input.GetContactSystemCode(), ContactId: input.GetContactId(), ContactCommunicationMethodId: input.GetContactCommunicationMethodId()})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

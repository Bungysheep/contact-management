package communicationmethod

import (
	"context"
	"os"
	"testing"

	"github.com/bungysheep/contact-management/pkg/api/v1/communicationmethod"
	"github.com/bungysheep/contact-management/pkg/common/message"
	"github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethod/mock_communicationmethod"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ctx  context.Context
	data []*communicationmethod.CommunicationMethod
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	data = append(data, &communicationmethod.CommunicationMethod{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "EMAIL",
		Description:             "Email",
		Details:                 "Email",
		Status:                  "A",
		FormatField:             "[EMAIL_ADDRESS]",
	}, &communicationmethod.CommunicationMethod{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "MOBILE",
		Description:             "Mobile",
		Details:                 "Mobile",
		Status:                  "A",
		FormatField:             "[MOBILE_NO]",
	}, &communicationmethod.CommunicationMethod{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "FAX",
		Description:             "Fax",
		Details:                 "Fax",
		Status:                  "A",
		FormatField:             "[FAX_NO]",
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestCommunicationMethodService(t *testing.T) {
	t.Run("DoRead Communication Method", doRead(ctx, data[0]))

	t.Run("DoReadAll Communication Method", doReadAll(ctx, data[0]))

	t.Run("DoSave new Communication Method", doSaveNew(ctx, data[0]))

	t.Run("DoSave existing Communication Method", doSaveExisting(ctx, data[0]))

	t.Run("DoDelete Communication Method", doDelete(ctx, data[0]))
}

func doRead(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_communicationmethod.NewMockICommunicationMethodRepository(ctl)

		repo.EXPECT().DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode()).Return(input, nil)

		svc := NewCommunicationMethodService(repo)

		resp, err := svc.DoRead(ctx, &communicationmethod.DoReadRequest{ContactSystemCode: input.GetContactSystemCode(), CommunicationMethodCode: input.GetCommunicationMethodCode()})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if resp.GetCommunicationMethod() == nil {
			t.Errorf("Expect communication method is not nil")
		}

		if resp.GetCommunicationMethod().GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetCommunicationMethod().GetContactSystemCode())
		}

		if resp.GetCommunicationMethod().GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp.GetCommunicationMethod().GetCommunicationMethodCode())
		}

		if resp.GetCommunicationMethod().GetDescription() != input.GetDescription() {
			t.Errorf("Expect description %s, but got %s", input.GetDescription(), resp.GetCommunicationMethod().GetDescription())
		}

		if resp.GetCommunicationMethod().GetDetails() != input.GetDetails() {
			t.Errorf("Expect details %s, but got %s", input.GetDetails(), resp.GetCommunicationMethod().GetDetails())
		}

		if resp.GetCommunicationMethod().GetStatus() != input.GetStatus() {
			t.Errorf("Expect status %s, but got %s", input.GetStatus(), resp.GetCommunicationMethod().GetStatus())
		}

		if resp.GetCommunicationMethod().GetFormatField() != input.GetFormatField() {
			t.Errorf("Expect format field %s, but got %s", input.GetFormatField(), resp.GetCommunicationMethod().GetFormatField())
		}
	}
}

func doReadAll(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_communicationmethod.NewMockICommunicationMethodRepository(ctl)

		repo.EXPECT().DoReadAll(ctx, input.GetContactSystemCode()).Return(data, nil)

		svc := NewCommunicationMethodService(repo)

		resp, err := svc.DoReadAll(ctx, &communicationmethod.DoReadAllRequest{ContactSystemCode: input.GetContactSystemCode()})
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

		if resp.GetCommunicationMethod()[0].GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp.GetCommunicationMethod()[0].GetCommunicationMethodCode())
		}

		if resp.GetCommunicationMethod()[0].GetDescription() != input.GetDescription() {
			t.Errorf("Expect description %s, but got %s", input.GetDescription(), resp.GetCommunicationMethod()[0].GetDescription())
		}

		if resp.GetCommunicationMethod()[0].GetDetails() != input.GetDetails() {
			t.Errorf("Expect details %s, but got %s", input.GetDetails(), resp.GetCommunicationMethod()[0].GetDetails())
		}

		if resp.GetCommunicationMethod()[0].GetStatus() != input.GetStatus() {
			t.Errorf("Expect status %s, but got %s", input.GetStatus(), resp.GetCommunicationMethod()[0].GetStatus())
		}

		if resp.GetCommunicationMethod()[0].GetFormatField() != input.GetFormatField() {
			t.Errorf("Expect format field %s, but got %s", input.GetFormatField(), resp.GetCommunicationMethod()[0].GetFormatField())
		}
	}
}

func doSaveNew(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_communicationmethod.NewMockICommunicationMethodRepository(ctl)

		repo.EXPECT().DoUpdate(ctx, input).Return(status.Errorf(codes.NotFound, message.DoesNotExist("Communication Method")))

		repo.EXPECT().DoInsert(ctx, input).Return(nil)

		svc := NewCommunicationMethodService(repo)

		resp, err := svc.DoSave(ctx, &communicationmethod.DoSaveRequest{CommunicationMethod: input})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

func doSaveExisting(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_communicationmethod.NewMockICommunicationMethodRepository(ctl)

		repo.EXPECT().DoUpdate(ctx, input).Return(nil)

		svc := NewCommunicationMethodService(repo)

		resp, err := svc.DoSave(ctx, &communicationmethod.DoSaveRequest{CommunicationMethod: input})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

func doDelete(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_communicationmethod.NewMockICommunicationMethodRepository(ctl)

		repo.EXPECT().DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode()).Return(nil)

		svc := NewCommunicationMethodService(repo)

		resp, err := svc.DoDelete(ctx, &communicationmethod.DoDeleteRequest{ContactSystemCode: input.GetContactSystemCode(), CommunicationMethodCode: input.GetCommunicationMethodCode()})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

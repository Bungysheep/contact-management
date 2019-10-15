package communicationmethodfield

import (
	"context"
	"os"
	"testing"

	"github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodfield"
	"github.com/bungysheep/contact-management/pkg/common/message"
	"github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethodfield/mock_communicationmethodfield"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ctx  context.Context
	data []*communicationmethodfield.CommunicationMethodField
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	data = append(data, &communicationmethodfield.CommunicationMethodField{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "EMAIL",
		FieldCode:               "EMAIL_ADDRESS",
		Caption:                 "Email Address",
		Sequence:                1,
	}, &communicationmethodfield.CommunicationMethodField{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "MOBILE",
		FieldCode:               "MOBILE_NO",
		Caption:                 "Mobile No",
		Sequence:                1,
	}, &communicationmethodfield.CommunicationMethodField{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "FAX",
		FieldCode:               "FAX_NO",
		Caption:                 "Fax No",
		Sequence:                1,
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestCommunicationMethodFieldService(t *testing.T) {
	t.Run("DoRead Communication Method Field", doRead(ctx, data[0]))

	t.Run("DoReadAll Communication Method Field", doReadAll(ctx, data[0]))

	t.Run("DoSave new Communication Method Field", doSaveNew(ctx, data[0]))

	t.Run("DoSave existing Communication Method Field", doSaveExisting(ctx, data[0]))

	t.Run("DoDelete Communication Method Field", doDelete(ctx, data[0]))
}

func doRead(ctx context.Context, input *communicationmethodfield.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_communicationmethodfield.NewMockICommunicationMethodFieldRepository(ctl)

		repo.EXPECT().DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode()).Return(input, nil)

		svc := NewCommunicationMethodFieldService(repo)

		resp, err := svc.DoRead(ctx, &communicationmethodfield.DoReadRequest{ContactSystemCode: input.GetContactSystemCode(), CommunicationMethodCode: input.GetCommunicationMethodCode(), FieldCode: input.GetFieldCode()})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if resp.GetCommunicationMethodField() == nil {
			t.Errorf("Expect communication method field is not nil")
		}

		if resp.GetCommunicationMethodField().GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetCommunicationMethodField().GetContactSystemCode())
		}

		if resp.GetCommunicationMethodField().GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp.GetCommunicationMethodField().GetCommunicationMethodCode())
		}

		if resp.GetCommunicationMethodField().GetFieldCode() != input.GetFieldCode() {
			t.Errorf("Expect field code %s, but got %s", input.GetFieldCode(), resp.GetCommunicationMethodField().GetFieldCode())
		}

		if resp.GetCommunicationMethodField().GetCaption() != input.GetCaption() {
			t.Errorf("Expect caption %s, but got %s", input.GetCaption(), resp.GetCommunicationMethodField().GetCaption())
		}

		if resp.GetCommunicationMethodField().GetSequence() != input.GetSequence() {
			t.Errorf("Expect sequence %d, but got %d", input.GetSequence(), resp.GetCommunicationMethodField().GetSequence())
		}
	}
}

func doReadAll(ctx context.Context, input *communicationmethodfield.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_communicationmethodfield.NewMockICommunicationMethodFieldRepository(ctl)

		repo.EXPECT().DoReadAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode()).Return(data, nil)

		svc := NewCommunicationMethodFieldService(repo)

		resp, err := svc.DoReadAll(ctx, &communicationmethodfield.DoReadAllRequest{ContactSystemCode: input.GetContactSystemCode(), CommunicationMethodCode: input.GetCommunicationMethodCode()})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if resp.GetCommunicationMethodField() == nil {
			t.Errorf("Expect communication method field is not nil")
		}

		if len(resp.GetCommunicationMethodField()) < 3 {
			t.Errorf("Expect there are communication method fields retrieved")
		}

		if resp.GetCommunicationMethodField()[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetCommunicationMethodField()[0].GetContactSystemCode())
		}

		if resp.GetCommunicationMethodField()[0].GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp.GetCommunicationMethodField()[0].GetCommunicationMethodCode())
		}

		if resp.GetCommunicationMethodField()[0].GetFieldCode() != input.GetFieldCode() {
			t.Errorf("Expect field code %s, but got %s", input.GetFieldCode(), resp.GetCommunicationMethodField()[0].GetFieldCode())
		}

		if resp.GetCommunicationMethodField()[0].GetCaption() != input.GetCaption() {
			t.Errorf("Expect caption %s, but got %s", input.GetCaption(), resp.GetCommunicationMethodField()[0].GetCaption())
		}

		if resp.GetCommunicationMethodField()[0].GetSequence() != input.GetSequence() {
			t.Errorf("Expect sequence %d, but got %d", input.GetSequence(), resp.GetCommunicationMethodField()[0].GetSequence())
		}
	}
}

func doSaveNew(ctx context.Context, input *communicationmethodfield.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_communicationmethodfield.NewMockICommunicationMethodFieldRepository(ctl)

		repo.EXPECT().DoUpdate(ctx, input).Return(status.Errorf(codes.NotFound, message.DoesNotExist("Communication Method Field")))

		repo.EXPECT().DoInsert(ctx, input).Return(nil)

		svc := NewCommunicationMethodFieldService(repo)

		resp, err := svc.DoSave(ctx, &communicationmethodfield.DoSaveRequest{CommunicationMethodField: input})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

func doSaveExisting(ctx context.Context, input *communicationmethodfield.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_communicationmethodfield.NewMockICommunicationMethodFieldRepository(ctl)

		repo.EXPECT().DoUpdate(ctx, input).Return(nil)

		svc := NewCommunicationMethodFieldService(repo)

		resp, err := svc.DoSave(ctx, &communicationmethodfield.DoSaveRequest{CommunicationMethodField: input})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

func doDelete(ctx context.Context, input *communicationmethodfield.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_communicationmethodfield.NewMockICommunicationMethodFieldRepository(ctl)

		repo.EXPECT().DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode()).Return(nil)

		svc := NewCommunicationMethodFieldService(repo)

		resp, err := svc.DoDelete(ctx, &communicationmethodfield.DoDeleteRequest{ContactSystemCode: input.GetContactSystemCode(), CommunicationMethodCode: input.GetCommunicationMethodCode(), FieldCode: input.GetFieldCode()})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

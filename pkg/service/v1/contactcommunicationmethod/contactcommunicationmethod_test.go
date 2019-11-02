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
		ContactSystemCode:               "CNTSYS001",
		ContactId:                       1,
		ContactCommunicationMethodId:    1,
		CommunicationMethodCode:         "EMAIL",
		CommunicationMethodLabelCode:    "HOME",
		CommunicationMethodLabelCaption: "Home",
		FormatValue:                     "test@gmail.com",
		IsDefault:                       true,
	}, &contactcommunicationmethod.ContactCommunicationMethod{
		ContactSystemCode:               "CNTSYS001",
		ContactId:                       1,
		ContactCommunicationMethodId:    2,
		CommunicationMethodCode:         "MOBILE",
		CommunicationMethodLabelCode:    "WORK",
		CommunicationMethodLabelCaption: "Work",
		FormatValue:                     "62-81234567890",
		IsDefault:                       true,
	}, &contactcommunicationmethod.ContactCommunicationMethod{
		ContactSystemCode:               "CNTSYS001",
		ContactId:                       1,
		ContactCommunicationMethodId:    3,
		CommunicationMethodCode:         "FAX",
		CommunicationMethodLabelCode:    "SCHOOL",
		CommunicationMethodLabelCaption: "School",
		FormatValue:                     "62-2471234567",
		IsDefault:                       true,
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

		resp, err := svc.DoRead(ctx, &contactcommunicationmethod.DoReadContactCommunicationMethodRequest{ContactSystemCode: input.GetContactSystemCode(), ContactId: input.GetContactId(), ContactCommunicationMethodId: input.GetContactCommunicationMethodId()})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if resp.GetContactCommunicationMethod() == nil {
			t.Errorf("Expect contact communication method is not nil")
		}

		if resp.GetContactCommunicationMethod().GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetContactCommunicationMethod().GetContactSystemCode())
		}

		if resp.GetContactCommunicationMethod().GetContactId() != input.GetContactId() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactId(), resp.GetContactCommunicationMethod().GetContactId())
		}

		if resp.GetContactCommunicationMethod().GetContactCommunicationMethodId() != input.GetContactCommunicationMethodId() {
			t.Errorf("Expect contact communication method id %d, but got %d", input.GetContactCommunicationMethodId(), resp.GetContactCommunicationMethod().GetContactCommunicationMethodId())
		}

		if resp.GetContactCommunicationMethod().GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp.GetContactCommunicationMethod().GetCommunicationMethodCode())
		}

		if resp.GetContactCommunicationMethod().GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp.GetContactCommunicationMethod().GetCommunicationMethodCode())
		}

		if resp.GetContactCommunicationMethod().GetCommunicationMethodLabelCode() != input.GetCommunicationMethodLabelCode() {
			t.Errorf("Expect communication method label code %s, but got %s", input.GetCommunicationMethodLabelCode(), resp.GetContactCommunicationMethod().GetCommunicationMethodLabelCode())
		}

		if resp.GetContactCommunicationMethod().GetCommunicationMethodLabelCaption() != input.GetCommunicationMethodLabelCaption() {
			t.Errorf("Expect communication method label caption %s, but got %s", input.GetCommunicationMethodLabelCaption(), resp.GetContactCommunicationMethod().GetCommunicationMethodLabelCaption())
		}

		if resp.GetContactCommunicationMethod().GetFormatValue() != input.GetFormatValue() {
			t.Errorf("Expect format value %s, but got %s", input.GetFormatValue(), resp.GetContactCommunicationMethod().GetFormatValue())
		}

		if resp.GetContactCommunicationMethod().GetIsDefault() != input.GetIsDefault() {
			t.Errorf("Expect default %v, but got %v", input.GetIsDefault(), resp.GetContactCommunicationMethod().GetIsDefault())
		}
	}
}

func doReadAll(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_contactcommunicationmethod.NewMockIContactCommunicationMethodRepository(ctl)

		repo.EXPECT().DoReadAll(ctx, input.GetContactSystemCode(), input.GetContactId()).Return(data, nil)

		svc := NewContactCommunicationMethodService(repo)

		resp, err := svc.DoReadAll(ctx, &contactcommunicationmethod.DoReadAllContactCommunicationMethodRequest{ContactSystemCode: input.GetContactSystemCode(), ContactId: input.GetContactId()})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if resp.GetContactCommunicationMethod() == nil {
			t.Errorf("Expect contact communication method is not nil")
		}

		if len(resp.GetContactCommunicationMethod()) < 3 {
			t.Errorf("Expect there are contact communication methods retrieved")
		}

		if resp.GetContactCommunicationMethod()[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetContactCommunicationMethod()[0].GetContactSystemCode())
		}

		if resp.GetContactCommunicationMethod()[0].GetContactId() != input.GetContactId() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactId(), resp.GetContactCommunicationMethod()[0].GetContactId())
		}

		if resp.GetContactCommunicationMethod()[0].GetContactCommunicationMethodId() != input.GetContactCommunicationMethodId() {
			t.Errorf("Expect contact communication method id %d, but got %d", input.GetContactCommunicationMethodId(), resp.GetContactCommunicationMethod()[0].GetContactCommunicationMethodId())
		}

		if resp.GetContactCommunicationMethod()[0].GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp.GetContactCommunicationMethod()[0].GetCommunicationMethodCode())
		}

		if resp.GetContactCommunicationMethod()[0].GetCommunicationMethodLabelCode() != input.GetCommunicationMethodLabelCode() {
			t.Errorf("Expect communication method label code %s, but got %s", input.GetCommunicationMethodLabelCode(), resp.GetContactCommunicationMethod()[0].GetCommunicationMethodLabelCode())
		}

		if resp.GetContactCommunicationMethod()[0].GetCommunicationMethodLabelCaption() != input.GetCommunicationMethodLabelCaption() {
			t.Errorf("Expect communication method label caption %s, but got %s", input.GetCommunicationMethodLabelCaption(), resp.GetContactCommunicationMethod()[0].GetCommunicationMethodLabelCaption())
		}

		if resp.GetContactCommunicationMethod()[0].GetFormatValue() != input.GetFormatValue() {
			t.Errorf("Expect format value %s, but got %s", input.GetFormatValue(), resp.GetContactCommunicationMethod()[0].GetFormatValue())
		}

		if resp.GetContactCommunicationMethod()[0].GetIsDefault() != input.GetIsDefault() {
			t.Errorf("Expect default %v, but got %v", input.GetIsDefault(), resp.GetContactCommunicationMethod()[0].GetIsDefault())
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

		resp, err := svc.DoSave(ctx, &contactcommunicationmethod.DoSaveContactCommunicationMethodRequest{ContactCommunicationMethod: input})
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

		resp, err := svc.DoSave(ctx, &contactcommunicationmethod.DoSaveContactCommunicationMethodRequest{ContactCommunicationMethod: input})
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

		resp, err := svc.DoDelete(ctx, &contactcommunicationmethod.DoDeleteContactCommunicationMethodRequest{ContactSystemCode: input.GetContactSystemCode(), ContactId: input.GetContactId(), ContactCommunicationMethodId: input.GetContactCommunicationMethodId()})
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

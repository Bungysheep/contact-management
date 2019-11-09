package contactcommunicationmethod

import (
	"context"
	"os"
	"testing"
	"time"

	auditapi "github.com/bungysheep/contact-management/pkg/api/v1/audit"
	contactcommunicationmethodapi "github.com/bungysheep/contact-management/pkg/api/v1/contactcommunicationmethod"
	auditmodel "github.com/bungysheep/contact-management/pkg/models/v1/audit"
	contactcommunicationmethodmodel "github.com/bungysheep/contact-management/pkg/models/v1/contactcommunicationmethod"
	contactcommunicationmethodfieldmodel "github.com/bungysheep/contact-management/pkg/models/v1/contactcommunicationmethodfield"
	"github.com/bungysheep/contact-management/pkg/service/v1/contactcommunicationmethod/mock_contactcommunicationmethod"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
)

var (
	ctx  context.Context
	data []*contactcommunicationmethodmodel.ContactCommunicationMethod
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	tmNow := time.Now().In(time.UTC)

	data = append(data, &contactcommunicationmethodmodel.ContactCommunicationMethod{
		ContactSystemCode:               "CNTSYS001",
		ContactID:                       1,
		ContactCommunicationMethodID:    1,
		CommunicationMethodCode:         "EMAIL",
		CommunicationMethodLabelCode:    "HOME",
		CommunicationMethodLabelCaption: "Home",
		FormatValue:                     "test@gmail.com",
		IsDefault:                       true,
		ContactCommunicationMethodField: make([]*contactcommunicationmethodfieldmodel.ContactCommunicationMethodField, 0),
		Audit: &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		},
	}, &contactcommunicationmethodmodel.ContactCommunicationMethod{
		ContactSystemCode:               "CNTSYS001",
		ContactID:                       1,
		ContactCommunicationMethodID:    2,
		CommunicationMethodCode:         "MOBILE",
		CommunicationMethodLabelCode:    "WORK",
		CommunicationMethodLabelCaption: "Work",
		FormatValue:                     "62-81234567890",
		IsDefault:                       true,
		ContactCommunicationMethodField: make([]*contactcommunicationmethodfieldmodel.ContactCommunicationMethodField, 0),
		Audit: &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		},
	}, &contactcommunicationmethodmodel.ContactCommunicationMethod{
		ContactSystemCode:               "CNTSYS001",
		ContactID:                       1,
		ContactCommunicationMethodID:    3,
		CommunicationMethodCode:         "FAX",
		CommunicationMethodLabelCode:    "SCHOOL",
		CommunicationMethodLabelCaption: "School",
		FormatValue:                     "62-2471234567",
		IsDefault:                       true,
		ContactCommunicationMethodField: make([]*contactcommunicationmethodfieldmodel.ContactCommunicationMethodField, 0),
		Audit: &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		},
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestContactCommunicationMethodService(t *testing.T) {
	t.Run("DoRead Contact Communication Method", doRead(ctx, data[0]))

	t.Run("DoReadAll Contact Communication Method", doReadAll(ctx, data[0]))

	t.Run("DoSave Contact Communication Method", doSave(ctx, data[0]))

	t.Run("DoDelete Contact Communication Method", doDelete(ctx, data[0]))
}

func doRead(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		svc := mock_contactcommunicationmethod.NewMockIContactCommunicationMethodService(ctl)

		svc.EXPECT().DoRead(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).Return(input, nil)

		svcServer := NewContactCommunicationMethodServiceServer(svc)

		resp, err := svcServer.DoRead(ctx, &contactcommunicationmethodapi.DoReadContactCommunicationMethodRequest{ContactSystemCode: input.GetContactSystemCode(), ContactId: input.GetContactID(), ContactCommunicationMethodId: input.GetContactCommunicationMethodID()})
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}

		if resp.GetContactCommunicationMethod() == nil {
			t.Fatalf("Expect contact communication method is not nil")
		}

		if resp.GetContactCommunicationMethod().GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetContactCommunicationMethod().GetContactSystemCode())
		}

		if resp.GetContactCommunicationMethod().GetContactId() != input.GetContactID() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactID(), resp.GetContactCommunicationMethod().GetContactId())
		}

		if resp.GetContactCommunicationMethod().GetContactCommunicationMethodId() != input.GetContactCommunicationMethodID() {
			t.Errorf("Expect contact communication method id %d, but got %d", input.GetContactCommunicationMethodID(), resp.GetContactCommunicationMethod().GetContactCommunicationMethodId())
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

func doReadAll(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		svc := mock_contactcommunicationmethod.NewMockIContactCommunicationMethodService(ctl)

		svc.EXPECT().DoReadAll(ctx, input.GetContactSystemCode(), input.GetContactID()).Return(data, nil)

		svcServer := NewContactCommunicationMethodServiceServer(svc)

		resp, err := svcServer.DoReadAll(ctx, &contactcommunicationmethodapi.DoReadAllContactCommunicationMethodRequest{ContactSystemCode: input.GetContactSystemCode(), ContactId: input.GetContactID()})
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}

		if resp.GetContactCommunicationMethod() == nil {
			t.Fatalf("Expect contact communication method is not nil")
		}

		if len(resp.GetContactCommunicationMethod()) < 3 {
			t.Errorf("Expect there are contact communication methods retrieved")
		}

		if resp.GetContactCommunicationMethod()[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetContactCommunicationMethod()[0].GetContactSystemCode())
		}

		if resp.GetContactCommunicationMethod()[0].GetContactId() != input.GetContactID() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactID(), resp.GetContactCommunicationMethod()[0].GetContactId())
		}

		if resp.GetContactCommunicationMethod()[0].GetContactCommunicationMethodId() != input.GetContactCommunicationMethodID() {
			t.Errorf("Expect contact communication method id %d, but got %d", input.GetContactCommunicationMethodID(), resp.GetContactCommunicationMethod()[0].GetContactCommunicationMethodId())
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

func doSave(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		svc := mock_contactcommunicationmethod.NewMockIContactCommunicationMethodService(ctl)

		contactCommMethod := &contactcommunicationmethodapi.ContactCommunicationMethod{Audit: &auditapi.Audit{}}
		contactCommMethod.ContactSystemCode = input.GetContactSystemCode()
		contactCommMethod.ContactId = input.GetContactID()
		contactCommMethod.ContactCommunicationMethodId = input.GetContactCommunicationMethodID()
		contactCommMethod.CommunicationMethodCode = input.GetCommunicationMethodCode()
		contactCommMethod.CommunicationMethodLabelCode = input.GetCommunicationMethodLabelCode()
		contactCommMethod.CommunicationMethodLabelCaption = input.GetCommunicationMethodLabelCaption()
		contactCommMethod.FormatValue = input.GetFormatValue()
		contactCommMethod.IsDefault = input.GetIsDefault()
		contactCommMethod.GetAudit().CreatedAt, _ = ptypes.TimestampProto(input.GetAudit().GetCreatedAt())
		contactCommMethod.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(input.GetAudit().GetModifiedAt())
		contactCommMethod.GetAudit().Vers = input.GetAudit().GetVers()

		svc.EXPECT().DoSave(ctx, input).Return(nil)

		svcServer := NewContactCommunicationMethodServiceServer(svc)

		resp, err := svcServer.DoSave(ctx, &contactcommunicationmethodapi.DoSaveContactCommunicationMethodRequest{ContactCommunicationMethod: contactCommMethod})
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

func doDelete(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		svc := mock_contactcommunicationmethod.NewMockIContactCommunicationMethodService(ctl)

		svc.EXPECT().DoDelete(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).Return(nil)

		svcServer := NewContactCommunicationMethodServiceServer(svc)

		resp, err := svcServer.DoDelete(ctx, &contactcommunicationmethodapi.DoDeleteContactCommunicationMethodRequest{ContactSystemCode: input.GetContactSystemCode(), ContactId: input.GetContactID(), ContactCommunicationMethodId: input.GetContactCommunicationMethodID()})
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

package communicationmethodfield

import (
	"context"
	"os"
	"testing"
	"time"

	auditapi "github.com/bungysheep/contact-management/pkg/api/v1/audit"
	communicationmethodfieldapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodfield"
	auditmodel "github.com/bungysheep/contact-management/pkg/models/v1/audit"
	communicationmethodfieldmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethodfield"
	"github.com/bungysheep/contact-management/pkg/service/v1/communicationmethodfield/mock_communicationmethodfield"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
)

var (
	ctx  context.Context
	data []*communicationmethodfieldmodel.CommunicationMethodField
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	tmNow := time.Now().In(time.UTC)

	data = append(data, &communicationmethodfieldmodel.CommunicationMethodField{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "EMAIL",
		FieldCode:               "EMAIL_ADDRESS",
		Caption:                 "Email Address",
		Sequence:                1,
		Audit: &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		},
	}, &communicationmethodfieldmodel.CommunicationMethodField{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "MOBILE",
		FieldCode:               "MOBILE_NO",
		Caption:                 "Mobile No",
		Sequence:                1,
		Audit: &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		},
	}, &communicationmethodfieldmodel.CommunicationMethodField{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "FAX",
		FieldCode:               "FAX_NO",
		Caption:                 "Fax No",
		Sequence:                1,
		Audit: &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		},
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestCommunicationMethodFieldService(t *testing.T) {
	t.Run("DoRead Communication Method Field", doRead(ctx, data[0]))

	t.Run("DoReadAll Communication Method Field", doReadAll(ctx, data[0]))

	t.Run("DoSave Communication Method Field", doSave(ctx, data[0]))

	t.Run("DoDelete Communication Method Field", doDelete(ctx, data[0]))
}

func doRead(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		svc := mock_communicationmethodfield.NewMockICommunicationMethodFieldService(ctl)

		svc.EXPECT().DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode()).Return(input, nil)

		svcServer := NewCommunicationMethodFieldServiceServer(svc)

		resp, err := svcServer.DoRead(ctx, &communicationmethodfieldapi.DoReadCommunicationMethodFieldRequest{ContactSystemCode: input.GetContactSystemCode(), CommunicationMethodCode: input.GetCommunicationMethodCode(), FieldCode: input.GetFieldCode()})
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
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

func doReadAll(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		svc := mock_communicationmethodfield.NewMockICommunicationMethodFieldService(ctl)

		svc.EXPECT().DoReadAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode()).Return(data, nil)

		svcServer := NewCommunicationMethodFieldServiceServer(svc)

		resp, err := svcServer.DoReadAll(ctx, &communicationmethodfieldapi.DoReadAllCommunicationMethodFieldRequest{ContactSystemCode: input.GetContactSystemCode(), CommunicationMethodCode: input.GetCommunicationMethodCode()})
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
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

func doSave(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		svc := mock_communicationmethodfield.NewMockICommunicationMethodFieldService(ctl)

		communicationMethodField := &communicationmethodfieldapi.CommunicationMethodField{Audit: &auditapi.Audit{}}
		communicationMethodField.ContactSystemCode = input.GetContactSystemCode()
		communicationMethodField.CommunicationMethodCode = input.GetCommunicationMethodCode()
		communicationMethodField.FieldCode = input.GetFieldCode()
		communicationMethodField.Caption = input.GetCaption()
		communicationMethodField.Sequence = input.GetSequence()
		communicationMethodField.GetAudit().CreatedAt, _ = ptypes.TimestampProto(input.GetAudit().GetCreatedAt())
		communicationMethodField.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(input.GetAudit().GetModifiedAt())
		communicationMethodField.GetAudit().Vers = input.GetAudit().GetVers()

		svc.EXPECT().DoSave(ctx, input).Return(nil)

		svcServer := NewCommunicationMethodFieldServiceServer(svc)

		resp, err := svcServer.DoSave(ctx, &communicationmethodfieldapi.DoSaveCommunicationMethodFieldRequest{CommunicationMethodField: communicationMethodField})
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

func doDelete(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		svc := mock_communicationmethodfield.NewMockICommunicationMethodFieldService(ctl)

		svc.EXPECT().DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode()).Return(nil)

		svcServer := NewCommunicationMethodFieldServiceServer(svc)

		resp, err := svcServer.DoDelete(ctx, &communicationmethodfieldapi.DoDeleteCommunicationMethodFieldRequest{ContactSystemCode: input.GetContactSystemCode(), CommunicationMethodCode: input.GetCommunicationMethodCode(), FieldCode: input.GetFieldCode()})
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

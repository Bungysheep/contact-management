package communicationmethod

import (
	"context"
	"os"
	"testing"
	"time"

	auditapi "github.com/bungysheep/contact-management/pkg/api/v1/audit"
	communicationmethodapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethod"
	auditmodel "github.com/bungysheep/contact-management/pkg/models/v1/audit"
	communicationmethodmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethod"
	"github.com/bungysheep/contact-management/pkg/service/v1/communicationmethod/mock_communicationmethod"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
)

var (
	ctx  context.Context
	data []*communicationmethodmodel.CommunicationMethod
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	tmNow := time.Now().In(time.UTC)

	data = append(data, &communicationmethodmodel.CommunicationMethod{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "EMAIL",
		Description:             "Email",
		Details:                 "Email",
		Status:                  "A",
		FormatField:             "[EMAIL_ADDRESS]",
		Audit: &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		},
	}, &communicationmethodmodel.CommunicationMethod{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "MOBILE",
		Description:             "Mobile",
		Details:                 "Mobile",
		Status:                  "A",
		FormatField:             "[MOBILE_NO]",
		Audit: &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		},
	}, &communicationmethodmodel.CommunicationMethod{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "FAX",
		Description:             "Fax",
		Details:                 "Fax",
		Status:                  "A",
		FormatField:             "[FAX_NO]",
		Audit: &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		},
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestCommunicationMethodService(t *testing.T) {
	t.Run("DoRead Communication Method", doRead(ctx, data[0]))

	t.Run("DoReadAll Communication Method", doReadAll(ctx, data[0]))

	t.Run("DoSave Communication Method", doSave(ctx, data[0]))

	t.Run("DoDelete Communication Method", doDelete(ctx, data[0]))
}

func doRead(ctx context.Context, input *communicationmethodmodel.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		svc := mock_communicationmethod.NewMockICommunicationMethodService(ctl)

		svc.EXPECT().DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode()).Return(input, nil)

		svcServer := NewCommunicationMethodServiceServer(svc)

		resp, err := svcServer.DoRead(ctx, &communicationmethodapi.DoReadCommunicationMethodRequest{ContactSystemCode: input.GetContactSystemCode(), CommunicationMethodCode: input.GetCommunicationMethodCode()})
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
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

func doReadAll(ctx context.Context, input *communicationmethodmodel.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		svc := mock_communicationmethod.NewMockICommunicationMethodService(ctl)

		svc.EXPECT().DoReadAll(ctx, input.GetContactSystemCode()).Return(data, nil)

		svcServer := NewCommunicationMethodServiceServer(svc)

		resp, err := svcServer.DoReadAll(ctx, &communicationmethodapi.DoReadAllCommunicationMethodRequest{ContactSystemCode: input.GetContactSystemCode()})
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
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

func doSave(ctx context.Context, input *communicationmethodmodel.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		svc := mock_communicationmethod.NewMockICommunicationMethodService(ctl)

		communicationMethod := &communicationmethodapi.CommunicationMethod{Audit: &auditapi.Audit{}}
		communicationMethod.ContactSystemCode = input.GetContactSystemCode()
		communicationMethod.CommunicationMethodCode = input.GetCommunicationMethodCode()
		communicationMethod.Description = input.GetDescription()
		communicationMethod.Details = input.GetDetails()
		communicationMethod.Status = input.GetStatus()
		communicationMethod.FormatField = input.GetFormatField()
		communicationMethod.GetAudit().CreatedAt, _ = ptypes.TimestampProto(input.GetAudit().GetCreatedAt())
		communicationMethod.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(input.GetAudit().GetModifiedAt())
		communicationMethod.GetAudit().Vers = input.GetAudit().GetVers()

		svc.EXPECT().DoSave(ctx, input).Return(nil)

		svcServer := NewCommunicationMethodServiceServer(svc)

		resp, err := svcServer.DoSave(ctx, &communicationmethodapi.DoSaveCommunicationMethodRequest{CommunicationMethod: communicationMethod})
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

func doDelete(ctx context.Context, input *communicationmethodmodel.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		svc := mock_communicationmethod.NewMockICommunicationMethodService(ctl)

		svc.EXPECT().DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode()).Return(nil)

		svcServer := NewCommunicationMethodServiceServer(svc)

		resp, err := svcServer.DoDelete(ctx, &communicationmethodapi.DoDeleteCommunicationMethodRequest{ContactSystemCode: input.GetContactSystemCode(), CommunicationMethodCode: input.GetCommunicationMethodCode()})
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

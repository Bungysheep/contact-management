package communicationmethodlabel

import (
	"context"
	"os"
	"testing"

	communicationmethodlabelapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodlabel"
	communicationmethodlabelmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethodlabel"
	"github.com/bungysheep/contact-management/pkg/service/v1/communicationmethodlabel/mock_communicationmethodlabel"
	"github.com/golang/mock/gomock"
)

var (
	ctx  context.Context
	data []*communicationmethodlabelmodel.CommunicationMethodLabel
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	data = append(data, &communicationmethodlabelmodel.CommunicationMethodLabel{
		ContactSystemCode:            "CNTSYS001",
		CommunicationMethodCode:      "EMAIL",
		CommunicationMethodLabelCode: "HOME",
		Caption:                      "Home",
	}, &communicationmethodlabelmodel.CommunicationMethodLabel{
		ContactSystemCode:            "CNTSYS001",
		CommunicationMethodCode:      "MOBILE",
		CommunicationMethodLabelCode: "WORK",
		Caption:                      "Work",
	}, &communicationmethodlabelmodel.CommunicationMethodLabel{
		ContactSystemCode:            "CNTSYS001",
		CommunicationMethodCode:      "FAX",
		CommunicationMethodLabelCode: "SCHOOL",
		Caption:                      "School",
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestCommunicationMethodLabelService(t *testing.T) {
	t.Run("DoRead Communication Method Label", doRead(ctx, data[0]))

	t.Run("DoReadAll Communication Method Label", doReadAll(ctx, data[0]))

	t.Run("DoSave Communication Method Label", doSave(ctx, data[0]))

	t.Run("DoDelete Communication Method Label", doDelete(ctx, data[0]))
}

func doRead(ctx context.Context, input *communicationmethodlabelmodel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_communicationmethodlabel.NewMockICommunicationMethodLabelService(ctl)

		repo.EXPECT().DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode()).Return(input, nil)

		svc := NewCommunicationMethodLabelServiceServer(repo)

		resp, err := svc.DoRead(ctx, &communicationmethodlabelapi.DoReadCommunicationMethodLabelRequest{ContactSystemCode: input.GetContactSystemCode(), CommunicationMethodCode: input.GetCommunicationMethodCode(), CommunicationMethodLabelCode: input.GetCommunicationMethodLabelCode()})
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}

		if resp.GetCommunicationMethodLabel() == nil {
			t.Fatalf("Expect communication method is not nil")
		}

		if resp.GetCommunicationMethodLabel().GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetCommunicationMethodLabel().GetContactSystemCode())
		}

		if resp.GetCommunicationMethodLabel().GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp.GetCommunicationMethodLabel().GetCommunicationMethodCode())
		}

		if resp.GetCommunicationMethodLabel().GetCommunicationMethodLabelCode() != input.GetCommunicationMethodLabelCode() {
			t.Errorf("Expect communication method label code %s, but got %s", input.GetCommunicationMethodLabelCode(), resp.GetCommunicationMethodLabel().GetCommunicationMethodLabelCode())
		}

		if resp.GetCommunicationMethodLabel().GetCaption() != input.GetCaption() {
			t.Errorf("Expect caption %s, but got %s", input.GetCaption(), resp.GetCommunicationMethodLabel().GetCaption())
		}
	}
}

func doReadAll(ctx context.Context, input *communicationmethodlabelmodel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_communicationmethodlabel.NewMockICommunicationMethodLabelService(ctl)

		repo.EXPECT().DoReadAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode()).Return(data, nil)

		svc := NewCommunicationMethodLabelServiceServer(repo)

		resp, err := svc.DoReadAll(ctx, &communicationmethodlabelapi.DoReadAllCommunicationMethodLabelRequest{ContactSystemCode: input.GetContactSystemCode(), CommunicationMethodCode: input.GetCommunicationMethodCode()})
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}

		if resp.GetCommunicationMethodLabel() == nil {
			t.Fatalf("Expect communication method is not nil")
		}

		if len(resp.GetCommunicationMethodLabel()) < 3 {
			t.Errorf("Expect there are communication methods retrieved")
		}

		if resp.GetCommunicationMethodLabel()[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetCommunicationMethodLabel()[0].GetContactSystemCode())
		}

		if resp.GetCommunicationMethodLabel()[0].GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp.GetCommunicationMethodLabel()[0].GetCommunicationMethodCode())
		}

		if resp.GetCommunicationMethodLabel()[0].GetCommunicationMethodLabelCode() != input.GetCommunicationMethodLabelCode() {
			t.Errorf("Expect communication method label %s, but got %s", input.GetCommunicationMethodLabelCode(), resp.GetCommunicationMethodLabel()[0].GetCommunicationMethodLabelCode())
		}

		if resp.GetCommunicationMethodLabel()[0].GetCaption() != input.GetCaption() {
			t.Errorf("Expect caption %s, but got %s", input.GetCaption(), resp.GetCommunicationMethodLabel()[0].GetCaption())
		}
	}
}

func doSave(ctx context.Context, input *communicationmethodlabelmodel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_communicationmethodlabel.NewMockICommunicationMethodLabelService(ctl)

		communicationMethodLabel := &communicationmethodlabelapi.CommunicationMethodLabel{}
		communicationMethodLabel.ContactSystemCode = input.GetContactSystemCode()
		communicationMethodLabel.CommunicationMethodCode = input.GetCommunicationMethodCode()
		communicationMethodLabel.CommunicationMethodLabelCode = input.GetCommunicationMethodLabelCode()
		communicationMethodLabel.Caption = input.GetCaption()

		repo.EXPECT().DoSave(ctx, input).Return(nil)

		svc := NewCommunicationMethodLabelServiceServer(repo)

		resp, err := svc.DoSave(ctx, &communicationmethodlabelapi.DoSaveCommunicationMethodLabelRequest{CommunicationMethodLabel: communicationMethodLabel})
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

func doDelete(ctx context.Context, input *communicationmethodlabelmodel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_communicationmethodlabel.NewMockICommunicationMethodLabelService(ctl)

		repo.EXPECT().DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode()).Return(nil)

		svc := NewCommunicationMethodLabelServiceServer(repo)

		resp, err := svc.DoDelete(ctx, &communicationmethodlabelapi.DoDeleteCommunicationMethodLabelRequest{ContactSystemCode: input.GetContactSystemCode(), CommunicationMethodCode: input.GetCommunicationMethodCode(), CommunicationMethodLabelCode: input.GetCommunicationMethodLabelCode()})
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}

		if !resp.GetResult() {
			t.Errorf("Expect the result is successful")
		}
	}
}

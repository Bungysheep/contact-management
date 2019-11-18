package communicationmethodlabel

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	communicationmethodlabelmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethodlabel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ctx  context.Context
	svc  ICommunicationMethodLabelService
	db   *sql.DB
	mock sqlmock.Sqlmock
	data []*communicationmethodlabelmodel.CommunicationMethodLabel
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	db, mock, _ = sqlmock.New()
	defer db.Close()

	svc = NewCommunicationMethodLabelService(db)

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
		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "communication_method_label_code", "caption"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetCaption())

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, communication_method_label_code, caption 
			FROM communication_method_label`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode()).WillReturnRows(rows)

		resp, err := svc.DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode())
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}

		if resp == nil {
			t.Fatalf("Expect communication method is not nil")
		}

		if resp.GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetContactSystemCode())
		}

		if resp.GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp.GetCommunicationMethodCode())
		}

		if resp.GetCommunicationMethodLabelCode() != input.GetCommunicationMethodLabelCode() {
			t.Errorf("Expect communication method label code %s, but got %s", input.GetCommunicationMethodLabelCode(), resp.GetCommunicationMethodLabelCode())
		}

		if resp.GetCaption() != input.GetCaption() {
			t.Errorf("Expect caption %s, but got %s", input.GetCaption(), resp.GetCaption())
		}
	}
}

func doReadAll(ctx context.Context, input *communicationmethodlabelmodel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "communication_method_label_code", "caption"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetCommunicationMethodCode(), data[0].GetCommunicationMethodLabelCode(), data[0].GetCaption()).
			AddRow(data[1].GetContactSystemCode(), data[1].GetCommunicationMethodCode(), data[1].GetCommunicationMethodLabelCode(), data[1].GetCaption()).
			AddRow(data[2].GetContactSystemCode(), data[2].GetCommunicationMethodCode(), data[2].GetCommunicationMethodLabelCode(), data[2].GetCaption())

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, communication_method_label_code, caption 
			FROM communication_method_label`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(rows)

		resp, err := svc.DoReadAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}

		if resp == nil {
			t.Fatalf("Expect communication method is not nil")
		}

		if len(resp) < 3 {
			t.Errorf("Expect there are communication methods retrieved")
		}

		if resp[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp[0].GetContactSystemCode())
		}

		if resp[0].GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp[0].GetCommunicationMethodCode())
		}

		if resp[0].GetCommunicationMethodLabelCode() != input.GetCommunicationMethodLabelCode() {
			t.Errorf("Expect communication method label %s, but got %s", input.GetCommunicationMethodLabelCode(), resp[0].GetCommunicationMethodLabelCode())
		}

		if resp[0].GetCaption() != input.GetCaption() {
			t.Errorf("Expect caption %s, but got %s", input.GetCaption(), resp[0].GetCaption())
		}
	}
}

func doSave(ctx context.Context, input *communicationmethodlabelmodel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoSave fail validation Communication Method Label", doSaveFailValidation(ctx))

		t.Run("DoSave invalid Communication Method", doSaveInvalidCommunicationMethod(ctx, input))

		t.Run("DoSave new Communication Method Label", doSaveNew(ctx, input))

		t.Run("DoSave existing Communication Method Label", doSaveExisting(ctx, input))
	}
}

func doSaveFailValidation(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		input := &communicationmethodlabelmodel.CommunicationMethodLabel{
			ContactSystemCode:            "CNTSYS000000000001",
			CommunicationMethodCode:      "EMAIL",
			CommunicationMethodLabelCode: "HOME",
			Caption:                      "Home",
		}

		err := svc.DoSave(ctx, input)
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a Unknown error, but got %s", s.Code())
				}
			}
		}
	}
}

func doSaveInvalidCommunicationMethod(ctx context.Context, input *communicationmethodlabelmodel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"})

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, description, details, status, format_field, 
				created_at, modified_at, vers 
			FROM communication_method`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(rows)

		err := svc.DoSave(ctx, input)
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.NotFound {
					t.Fatalf("Expect a NotFound error, but got %s", s.Code())
				}
			}
		}
	}
}

func doSaveNew(ctx context.Context, input *communicationmethodlabelmodel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), "", "", "", "", tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, description, details, status, format_field, 
				created_at, modified_at, vers 
			FROM communication_method`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(rows)

		expUpdQuery := mock.ExpectPrepare("UPDATE communication_method_label").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetCaption()).WillReturnResult(sqlmock.NewResult(0, 0))

		expInsQuery := mock.ExpectPrepare("INSERT INTO communication_method_label").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetCaption()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoSave(ctx, input)
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}
	}
}

func doSaveExisting(ctx context.Context, input *communicationmethodlabelmodel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), "", "", "", "", tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, description, details, status, format_field, 
				created_at, modified_at, vers 
			FROM communication_method`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(rows)

		expUpdQuery := mock.ExpectPrepare("UPDATE communication_method_label").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetCaption()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoSave(ctx, input)
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}
	}
}

func doDelete(ctx context.Context, input *communicationmethodlabelmodel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM communication_method_label").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode())
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}
	}
}

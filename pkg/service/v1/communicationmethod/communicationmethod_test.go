package communicationmethod

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	auditmodel "github.com/bungysheep/contact-management/pkg/models/v1/audit"
	communicationmethodmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethod"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ctx  context.Context
	svc  ICommunicationMethodService
	db   *sql.DB
	mock sqlmock.Sqlmock
	data []*communicationmethodmodel.CommunicationMethod
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	db, mock, _ = sqlmock.New()
	defer db.Close()

	svc = NewCommunicationMethodService(db)

	data = append(data, &communicationmethodmodel.CommunicationMethod{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "EMAIL",
		Description:             "Email",
		Details:                 "Email",
		Status:                  "A",
		FormatField:             "[EMAIL_ADDRESS]",
	}, &communicationmethodmodel.CommunicationMethod{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "MOBILE",
		Description:             "Mobile",
		Details:                 "Mobile",
		Status:                  "A",
		FormatField:             "[MOBILE_NO]",
	}, &communicationmethodmodel.CommunicationMethod{
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

	t.Run("DoSave Communication Method", doSave(ctx, data[0]))

	t.Run("DoDelete Communication Method", doDelete(ctx, data[0]))
}

func doRead(ctx context.Context, input *communicationmethodmodel.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), input.GetFormatField(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, description, details, status, format_field, 
				created_at, modified_at, vers 
			FROM communication_method`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(rows)

		resp, err := svc.DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
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

		if resp.GetDescription() != input.GetDescription() {
			t.Errorf("Expect description %s, but got %s", input.GetDescription(), resp.GetDescription())
		}

		if resp.GetDetails() != input.GetDetails() {
			t.Errorf("Expect details %s, but got %s", input.GetDetails(), resp.GetDetails())
		}

		if resp.GetStatus() != input.GetStatus() {
			t.Errorf("Expect status %s, but got %s", input.GetStatus(), resp.GetStatus())
		}

		if resp.GetFormatField() != input.GetFormatField() {
			t.Errorf("Expect format field %s, but got %s", input.GetFormatField(), resp.GetFormatField())
		}
	}
}

func doReadAll(ctx context.Context, input *communicationmethodmodel.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetCommunicationMethodCode(), data[0].GetDescription(), data[0].GetDetails(), data[0].GetStatus(), data[0].GetFormatField(), tmNow, tmNow, 1).
			AddRow(data[1].GetContactSystemCode(), data[1].GetCommunicationMethodCode(), data[1].GetDescription(), data[1].GetDetails(), data[1].GetStatus(), data[1].GetFormatField(), tmNow, tmNow, 1).
			AddRow(data[2].GetContactSystemCode(), data[2].GetCommunicationMethodCode(), data[2].GetDescription(), data[2].GetDetails(), data[2].GetStatus(), data[2].GetFormatField(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, description, details, status, format_field, 
				created_at, modified_at, vers 
			FROM communication_method`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(rows)

		resp, err := svc.DoReadAll(ctx, input.GetContactSystemCode())
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

		if resp[0].GetDescription() != input.GetDescription() {
			t.Errorf("Expect description %s, but got %s", input.GetDescription(), resp[0].GetDescription())
		}

		if resp[0].GetDetails() != input.GetDetails() {
			t.Errorf("Expect details %s, but got %s", input.GetDetails(), resp[0].GetDetails())
		}

		if resp[0].GetStatus() != input.GetStatus() {
			t.Errorf("Expect status %s, but got %s", input.GetStatus(), resp[0].GetStatus())
		}

		if resp[0].GetFormatField() != input.GetFormatField() {
			t.Errorf("Expect format field %s, but got %s", input.GetFormatField(), resp[0].GetFormatField())
		}
	}
}

func doSave(ctx context.Context, input *communicationmethodmodel.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoSave fail validation Communication Method", doSaveFailValidation(ctx))

		t.Run("DoSave invalid Contact System", doSaveInvalidContactSystem(ctx, input))

		t.Run("DoSave new Communication Method", doSaveNew(ctx, input))

		t.Run("DoSave existing Communication Method", doSaveExisting(ctx, input))
	}
}

func doSaveFailValidation(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		input := &communicationmethodmodel.CommunicationMethod{
			ContactSystemCode:       "CNTSYS001",
			CommunicationMethodCode: "EMAIL",
			Description:             "Email",
			Details:                 "Email",
			Status:                  "X",
			FormatField:             "[EMAIL_ADDRESS]",
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

func doSaveInvalidContactSystem(ctx context.Context, input *communicationmethodmodel.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "description", "details", "status", "created_at", "modified_at", "vers"})

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, description, details, status, 
				created_at, modified_at, vers 
			FROM contact_system`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(rows)

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

func doSaveNew(ctx context.Context, input *communicationmethodmodel.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		}

		rows := sqlmock.NewRows([]string{"contact_system_code", "description", "details", "status", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), "", "", "", tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, description, details, status, created_at, modified_at, vers FROM contact_system").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(rows)

		expUpdQuery := mock.ExpectPrepare("UPDATE communication_method").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), input.GetFormatField(), tmNow).WillReturnResult(sqlmock.NewResult(0, 0))

		expInsQuery := mock.ExpectPrepare("INSERT INTO communication_method").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), input.GetFormatField(), tmNow, tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoSave(ctx, input)
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}
	}
}

func doSaveExisting(ctx context.Context, input *communicationmethodmodel.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       2,
		}

		rows := sqlmock.NewRows([]string{"contact_system_code", "description", "details", "status", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), "", "", "", tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, description, details, status, 
				created_at, modified_at, vers 
			FROM contact_system`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(rows)

		expUpdQuery := mock.ExpectPrepare("UPDATE communication_method").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), input.GetFormatField(), tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoSave(ctx, input)
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}
	}
}

func doDelete(ctx context.Context, input *communicationmethodmodel.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoDelete fail Communication Method Field", doDeleteFailCommunicationMethodField(ctx, input))

		t.Run("DoDelete fail Communication Method Label", doDeleteFailCommunicationMethodLabel(ctx, input))

		t.Run("DoDelete existing Communication Method", doDeleteExisting(ctx, input))
	}
}

func doDeleteFailCommunicationMethodField(ctx context.Context, input *communicationmethodmodel.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expCommMethodFieldQuery := mock.ExpectPrepare("DELETE FROM communication_method_field").ExpectExec()
		expCommMethodFieldQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnError(fmt.Errorf("Delete all communication method fields failed"))

		err := svc.DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
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

func doDeleteFailCommunicationMethodLabel(ctx context.Context, input *communicationmethodmodel.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expCommMethodFieldQuery := mock.ExpectPrepare("DELETE FROM communication_method_field").ExpectExec()
		expCommMethodFieldQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnResult(sqlmock.NewResult(0, 1))

		expCommMethodLabelQuery := mock.ExpectPrepare("DELETE FROM communication_method_label").ExpectExec()
		expCommMethodLabelQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnError(fmt.Errorf("Delete all communication method labels failed"))

		err := svc.DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
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

func doDeleteExisting(ctx context.Context, input *communicationmethodmodel.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expCommMethodFieldQuery := mock.ExpectPrepare("DELETE FROM communication_method_field").ExpectExec()
		expCommMethodFieldQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnResult(sqlmock.NewResult(0, 1))

		expCommMethodLabelQuery := mock.ExpectPrepare("DELETE FROM communication_method_label").ExpectExec()
		expCommMethodLabelQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnResult(sqlmock.NewResult(0, 1))

		expCommMethodQuery := mock.ExpectPrepare("DELETE FROM communication_method").ExpectExec()
		expCommMethodQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}
	}
}

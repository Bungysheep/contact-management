package communicationmethodfield

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	auditmodel "github.com/bungysheep/contact-management/pkg/models/v1/audit"
	communicationmethodfieldmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethodfield"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ctx  context.Context
	svc  ICommunicationMethodFieldService
	db   *sql.DB
	mock sqlmock.Sqlmock
	data []*communicationmethodfieldmodel.CommunicationMethodField
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	db, mock, _ = sqlmock.New()
	defer db.Close()

	svc = NewCommunicationMethodFieldService(db)

	data = append(data, &communicationmethodfieldmodel.CommunicationMethodField{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "EMAIL",
		FieldCode:               "EMAIL_ADDRESS",
		Caption:                 "Email Address",
		Sequence:                1,
	}, &communicationmethodfieldmodel.CommunicationMethodField{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "MOBILE",
		FieldCode:               "MOBILE_NO",
		Caption:                 "Mobile No",
		Sequence:                1,
	}, &communicationmethodfieldmodel.CommunicationMethodField{
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

	t.Run("DoSave Communication Method Field", doSave(ctx, data[0]))

	t.Run("DoDelete Communication Method Field", doDelete(ctx, data[0]))
}

func doRead(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "field_code", "caption", "sequence", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode(), input.GetCaption(), input.GetSequence(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, field_code, caption, sequence, 
				created_at, modified_at, vers 
			FROM communication_method_field`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode()).WillReturnRows(rows)

		resp, err := svc.DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode())
		if err != nil {
			t.Fatalf("Expect error is nil, but got %s", err)
		}

		if resp == nil {
			t.Fatalf("Expect communication method field is not nil")
		}

		if resp.GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetContactSystemCode())
		}

		if resp.GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp.GetCommunicationMethodCode())
		}

		if resp.GetFieldCode() != input.GetFieldCode() {
			t.Errorf("Expect field code %s, but got %s", input.GetFieldCode(), resp.GetFieldCode())
		}

		if resp.GetCaption() != input.GetCaption() {
			t.Errorf("Expect caption %s, but got %s", input.GetCaption(), resp.GetCaption())
		}

		if resp.GetSequence() != input.GetSequence() {
			t.Errorf("Expect sequence %d, but got %d", input.GetSequence(), resp.GetSequence())
		}
	}
}

func doReadAll(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "field_code", "caption", "sequence", "created_at", "modified_at", "vers"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetCommunicationMethodCode(), data[0].GetFieldCode(), data[0].GetCaption(), data[0].GetSequence(), tmNow, tmNow, 1).
			AddRow(data[1].GetContactSystemCode(), data[1].GetCommunicationMethodCode(), data[1].GetFieldCode(), data[1].GetCaption(), data[1].GetSequence(), tmNow, tmNow, 1).
			AddRow(data[2].GetContactSystemCode(), data[2].GetCommunicationMethodCode(), data[2].GetFieldCode(), data[2].GetCaption(), data[2].GetSequence(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, field_code, caption, sequence, 
				created_at, modified_at, vers 
			FROM communication_method_field`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(rows)

		resp, err := svc.DoReadAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
		if err != nil {
			t.Fatalf("Expect error is nil, but got %s", err)
		}

		if resp == nil {
			t.Fatalf("Expect communication method field is not nil")
		}

		if len(resp) < 3 {
			t.Errorf("Expect there are communication method fields retrieved")
		}

		if resp[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp[0].GetContactSystemCode())
		}

		if resp[0].GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp[0].GetCommunicationMethodCode())
		}

		if resp[0].GetFieldCode() != input.GetFieldCode() {
			t.Errorf("Expect field code %s, but got %s", input.GetFieldCode(), resp[0].GetFieldCode())
		}

		if resp[0].GetCaption() != input.GetCaption() {
			t.Errorf("Expect caption %s, but got %s", input.GetCaption(), resp[0].GetCaption())
		}

		if resp[0].GetSequence() != input.GetSequence() {
			t.Errorf("Expect sequence %d, but got %d", input.GetSequence(), resp[0].GetSequence())
		}
	}
}

func doSave(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoSave invalid Communication Method", doSaveInvalidCommunicationMethod(ctx, input))

		t.Run("DoSave new Communication Method Field", doSaveNew(ctx, input))

		t.Run("DoSave existing Communication Method Field", doSaveExisting(ctx, input))
	}
}

func doSaveInvalidCommunicationMethod(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
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

func doSaveNew(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		}

		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), "", "", "", "", tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, description, details, status, format_field, 
				created_at, modified_at, vers 
			FROM communication_method`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(rows)

		expUpdQuery := mock.ExpectPrepare("UPDATE communication_method_field").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode(), input.GetCaption(), input.GetSequence(), tmNow).WillReturnResult(sqlmock.NewResult(0, 0))

		expInsQuery := mock.ExpectPrepare("INSERT INTO communication_method_field").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode(), input.GetCaption(), input.GetSequence(), tmNow, tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoSave(ctx, input)
		if err != nil {
			t.Fatalf("Expect error is nil, but got %s", err)
		}
	}
}

func doSaveExisting(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       2,
		}

		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), "", "", "", "", tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, description, details, status, format_field, 
				created_at, modified_at, vers 
			FROM communication_method`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(rows)

		expUpdQuery := mock.ExpectPrepare("UPDATE communication_method_field").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode(), input.GetCaption(), input.GetSequence(), tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoSave(ctx, input)
		if err != nil {
			t.Fatalf("Expect error is nil, but got %s", err)
		}
	}
}

func doDelete(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM communication_method_field").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode())
		if err != nil {
			t.Fatalf("Expect error is nil, but got %s", err)
		}
	}
}

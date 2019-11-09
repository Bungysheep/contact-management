package communicationmethodfield

import (
	"context"
	"database/sql"
	"fmt"
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
	repo ICommunicationMethodFieldRepository
	db   *sql.DB
	mock sqlmock.Sqlmock
	data []*communicationmethodfieldmodel.CommunicationMethodField
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	db, mock, _ = sqlmock.New()
	defer db.Close()

	repo = NewCommunicationMethodFieldRepository(db)

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

func TestCommunicationMethodFieldRepository(t *testing.T) {
	t.Run("DoRead Communication Method Field", doRead(ctx))

	t.Run("DoReadAll Communication Method Field", doReadAll(ctx))

	t.Run("DoSave Communication Method Field", doSave(ctx))

	t.Run("DoDelete Communication Method Field", doDelete(ctx))

	t.Run("DoDeleteAll Communication Method Field", doDeleteAll(ctx))
}

func doRead(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoRead fail", doReadFailCommunicationMethodField(ctx, data[0]))

		t.Run("DoRead unexisting", doReadUnexistingCommunicationMethodField(ctx, data[0]))

		t.Run("DoRead row error", doReadRowErrorCommunicationMethodField(ctx, data[0]))

		t.Run("DoRead existing", doReadExistingCommunicationMethodField(ctx, data[0]))
	}
}

func doReadAll(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoReadAll fail", doReadAllFailCommunicationMethodFields(ctx, data[0]))

		t.Run("DoReadAll unexisting", doReadAllUnexistingCommunicationMethodFields(ctx, data[0]))

		t.Run("DoReadAll row error", doReadAllRowErrorCommunicationMethodFields(ctx, data[0]))

		t.Run("DoReadAll existing", doReadAllExistingCommunicationMethodFields(ctx, data[0]))
	}
}

func doSave(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoSave new fail", doSaveNewFailCommunicationMethodField(ctx, data[0]))

		t.Run("DoSave new", doSaveNewCommunicationMethodField(ctx, data[0]))

		t.Run("DoSave existing fail", doSaveExistingFailCommunicationMethodField(ctx, data[0]))

		t.Run("DoSave existing", doSaveExistingCommunicationMethodField(ctx, data[0]))
	}
}

func doDelete(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoDelete fail", doDeleteFailCommunicationMethodField(ctx, data[0]))

		t.Run("DoDelete unexisting", doDeleteUnexistingCommunicationMethodField(ctx, data[0]))

		t.Run("DoDelete existing", doDeleteExistingCommunicationMethodField(ctx, data[0]))
	}
}

func doDeleteAll(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoDeleteAll fail", doDeleteAllFailCommunicationMethodField(ctx, data[0]))

		t.Run("DoDeleteAll unexisting", doDeleteAllUnexistingCommunicationMethodField(ctx, data[0]))

		t.Run("DoDeleteAll existing", doDeleteAllExistingCommunicationMethodField(ctx, data[0]))
	}
}

func doReadFailCommunicationMethodField(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, field_code, caption, sequence, 
				created_at, modified_at, vers 
			FROM communication_method_field`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode()).WillReturnError(fmt.Errorf("DoRead communication method field failed"))

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a Unknown error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}

		if res != nil {
			t.Errorf("Expect communication method field is nil")
		}
	}
}

func doReadUnexistingCommunicationMethodField(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "field_code", "caption", "sequence", "created_at", "modified_at", "vers"})

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, field_code, caption, sequence, 
				created_at, modified_at, vers 
			FROM communication_method_field`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.NotFound {
					t.Fatalf("Expect a NotFound error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}

		if res != nil {
			t.Errorf("Expect communication method field is nil")
		}
	}
}

func doReadRowErrorCommunicationMethodField(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "field_code", "caption", "sequence", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode(), input.GetCaption(), input.GetSequence(), tmNow, tmNow, 1).
			RowError(0, fmt.Errorf("DoRead row error"))

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, field_code, caption, sequence, 
				created_at, modified_at, vers 
			FROM communication_method_field`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a Unknown error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}

		if res != nil {
			t.Errorf("Expect communication method field is nil")
		}
	}
}

func doReadExistingCommunicationMethodField(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "field_code", "caption", "sequence", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode(), input.GetCaption(), input.GetSequence(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, field_code, caption, sequence, 
				created_at, modified_at, vers 
			FROM communication_method_field`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode())
		if err != nil {
			t.Fatalf("Failed to read communication method field: %v", err)
		}

		if res == nil {
			t.Fatalf("Expect communication method field is not nil")
		}

		if res.GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), res.GetContactSystemCode())
		}

		if res.GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method field code %s, but got %s", input.GetCommunicationMethodCode(), res.GetCommunicationMethodCode())
		}

		if res.GetFieldCode() != input.GetFieldCode() {
			t.Errorf("Expect field code %s, but got %s", input.GetFieldCode(), res.GetFieldCode())
		}

		if res.GetCaption() != input.GetCaption() {
			t.Errorf("Expect caption %s, but got %s", input.GetCaption(), res.GetCaption())
		}

		if res.GetSequence() != input.GetSequence() {
			t.Errorf("Expect sequence %d, but got %d", input.GetSequence(), res.GetSequence())
		}
	}
}

func doReadAllFailCommunicationMethodFields(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, field_code, caption, sequence, 
				created_at, modified_at, vers 
			FROM communication_method_field`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnError(fmt.Errorf("DoReadAll communication method field failed"))

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a Unknown error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}

		if res == nil {
			t.Fatalf("Expect communication method fields is not nil")
		}

		if len(res) != 0 {
			t.Errorf("Expect response is nil")
		}
	}
}

func doReadAllUnexistingCommunicationMethodFields(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "field_code", "caption", "sequence", "created_at", "modified_at", "vers"})

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, field_code, caption, sequence, 
				created_at, modified_at, vers 
			FROM communication_method_field`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(rows)

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.NotFound {
					t.Fatalf("Expect a NotFound error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}

		if res == nil {
			t.Fatalf("Expect communication method fields is not nil")
		}

		if len(res) != 0 {
			t.Errorf("Expect no communication method fields retrieved")
		}
	}
}

func doReadAllRowErrorCommunicationMethodFields(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "field_code", "caption", "sequence", "created_at", "modified_at", "vers"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetCommunicationMethodCode(), data[0].GetFieldCode(), data[0].GetCaption(), data[0].GetSequence(), tmNow, tmNow, 1).
			RowError(0, fmt.Errorf("DoReadAll row error")).
			AddRow(data[1].GetContactSystemCode(), data[1].GetCommunicationMethodCode(), data[1].GetFieldCode(), data[1].GetCaption(), data[1].GetSequence(), tmNow, tmNow, 1).
			AddRow(data[2].GetContactSystemCode(), data[2].GetCommunicationMethodCode(), data[2].GetFieldCode(), data[2].GetCaption(), data[2].GetSequence(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, field_code, caption, sequence, 
				created_at, modified_at, vers 
			FROM communication_method_field`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(rows)

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a Unknown error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}

		if res == nil {
			t.Fatalf("Expect communication method fields is not nil")
		}

		if len(res) != 0 {
			t.Errorf("Expect response is nil")
		}
	}
}

func doReadAllExistingCommunicationMethodFields(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
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

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
		if err != nil {
			t.Fatalf("Failed to read all communication method fields: %v", err)
		}

		if res == nil {
			t.Fatalf("Expect communication method fields is not nil")
		}

		if len(res) < 3 {
			t.Errorf("Expect there are communication method fields retrieved")
		}

		if res[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), res[0].GetContactSystemCode())
		}

		if res[0].GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), res[0].GetCommunicationMethodCode())
		}

		if res[0].GetFieldCode() != input.GetFieldCode() {
			t.Errorf("Expect field code %s, but got %s", input.GetFieldCode(), res[0].GetFieldCode())
		}

		if res[0].GetCaption() != input.GetCaption() {
			t.Errorf("Expect caption %s, but got %s", input.GetCaption(), res[0].GetCaption())
		}

		if res[0].GetSequence() != input.GetSequence() {
			t.Errorf("Expect sequence %d, but got %d", input.GetSequence(), res[0].GetSequence())
		}
	}
}

func doSaveNewFailCommunicationMethodField(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		}

		expInsQuery := mock.ExpectPrepare("INSERT INTO communication_method_field").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode(), input.GetCaption(), input.GetSequence(), tmNow, tmNow).WillReturnError(fmt.Errorf("DoInsert communication method field failed"))

		err := repo.DoInsert(ctx, input)
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a Unknown error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doSaveNewCommunicationMethodField(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		}

		expInsQuery := mock.ExpectPrepare("INSERT INTO communication_method_field").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode(), input.GetCaption(), input.GetSequence(), tmNow, tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoInsert(ctx, input)
		if err != nil {
			t.Fatalf("Failed to save communication method field: %v", err)
		}
	}
}

func doSaveExistingFailCommunicationMethodField(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       2,
		}

		expUpdQuery := mock.ExpectPrepare("UPDATE communication_method_field").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode(), input.GetCaption(), input.GetSequence(), tmNow).WillReturnError(fmt.Errorf("DoUpdate communication method field failed"))

		err := repo.DoUpdate(ctx, input)
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a Unknown error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doSaveExistingCommunicationMethodField(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       2,
		}

		expUpdQuery := mock.ExpectPrepare("UPDATE communication_method_field").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode(), input.GetCaption(), input.GetSequence(), tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoUpdate(ctx, input)
		if err != nil {
			t.Fatalf("Failed to save communication method field: %v", err)
		}
	}
}

func doDeleteFailCommunicationMethodField(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM communication_method_field").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode()).WillReturnError(fmt.Errorf("Delete communication method field failed"))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a Unknown error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doDeleteUnexistingCommunicationMethodField(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM communication_method_field").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode()).WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.NotFound {
					t.Fatalf("Expect a NotFound error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doDeleteExistingCommunicationMethodField(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM communication_method_field").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetFieldCode())
		if err != nil {
			t.Fatalf("Failed to delete communication method field: %v", err)
		}
	}
}

func doDeleteAllFailCommunicationMethodField(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM communication_method_field").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnError(fmt.Errorf("Delete all communication method fields failed"))

		err := repo.DoDeleteAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a Unknown error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doDeleteAllUnexistingCommunicationMethodField(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM communication_method_field").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.DoDeleteAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
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

func doDeleteAllExistingCommunicationMethodField(ctx context.Context, input *communicationmethodfieldmodel.CommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM communication_method_field").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoDeleteAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
		if err != nil {
			t.Fatalf("Failed to delete all communication method fields: %v", err)
		}
	}
}

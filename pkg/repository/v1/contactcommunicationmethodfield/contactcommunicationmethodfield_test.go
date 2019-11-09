package contactcommunicationmethodfield

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	contactcommunicationmethodfieldmodel "github.com/bungysheep/contact-management/pkg/models/v1/contactcommunicationmethodfield"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ctx  context.Context
	repo IContactCommunicationMethodFieldRepository
	db   *sql.DB
	mock sqlmock.Sqlmock
	data []*contactcommunicationmethodfieldmodel.ContactCommunicationMethodField
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	db, mock, _ = sqlmock.New()
	defer db.Close()

	repo = NewContactCommunicationMethodFieldRepository(db)

	data = append(data, &contactcommunicationmethodfieldmodel.ContactCommunicationMethodField{
		ContactSystemCode:            "CNTSYS001",
		ContactID:                    1,
		ContactCommunicationMethodID: 1,
		FieldCode:                    "EMAIL_ADDRESS",
		FieldValue:                   "test@gmail.com",
	}, &contactcommunicationmethodfieldmodel.ContactCommunicationMethodField{
		ContactSystemCode:            "CNTSYS001",
		ContactID:                    1,
		ContactCommunicationMethodID: 2,
		FieldCode:                    "MOBILE_NO",
		FieldValue:                   "62-81234567890",
	}, &contactcommunicationmethodfieldmodel.ContactCommunicationMethodField{
		ContactSystemCode:            "CNTSYS001",
		ContactID:                    1,
		ContactCommunicationMethodID: 3,
		FieldCode:                    "FAX_NO",
		FieldValue:                   "62-2471234567",
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestContactCommunicationMethodRepository(t *testing.T) {
	t.Run("DoRead Contact Communication Method Field", doRead(ctx))

	t.Run("DoSave Contact Communication Method Field", doSave(ctx))

	t.Run("DoDelete Contact Communication Method Field", doDelete(ctx))
}

func doRead(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoRead fail", doReadFailContactCommunicationMethodFields(ctx, data[0]))

		t.Run("DoRead unexisting", doReadUnexistingContactCommunicationMethodFields(ctx, data[0]))

		t.Run("DoRead row error", doReadRowErrorContactCommunicationMethodFields(ctx, data[0]))

		t.Run("DoRead existing", doReadExistingContactCommunicationMethodFields(ctx, data[0]))
	}
}

func doSave(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoSave new fail", doSaveNewFailContactCommunicationMethodField(ctx, data[0]))

		t.Run("DoSave new", doSaveNewContactCommunicationMethodField(ctx, data[0]))

		t.Run("DoSave existing fail", doSaveExistingFailContactCommunicationMethodField(ctx, data[0]))

		t.Run("DoSave existing", doSaveExistingContactCommunicationMethodField(ctx, data[0]))
	}
}

func doDelete(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoDelete fail", doDeleteFailContactCommunicationMethodField(ctx, data[0]))

		t.Run("DoDelete unexisting", doDeleteUnexistingContactCommunicationMethodField(ctx, data[0]))

		t.Run("DoDelete existing", doDeleteExistingContactCommunicationMethodField(ctx, data[0]))
	}
}

func doReadFailContactCommunicationMethodFields(ctx context.Context, input *contactcommunicationmethodfieldmodel.ContactCommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, contact_id, contact_communication_method_id, field_code, field_value 
			FROM contact_communication_method_field`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnError(fmt.Errorf("DoRead contact communication method field failed"))

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID())
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
			t.Fatalf("Expect contact communication method fields is not nil")
		}

		if len(res) != 0 {
			t.Errorf("Expect response is nil")
		}
	}
}

func doReadUnexistingContactCommunicationMethodFields(ctx context.Context, input *contactcommunicationmethodfieldmodel.ContactCommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "field_code", "field_value"})

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, contact_id, contact_communication_method_id, field_code, field_value 
			FROM contact_communication_method_field`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID())
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
			t.Fatalf("Expect contact communication method fields is not nil")
		}

		if len(res) != 0 {
			t.Errorf("Expect no contact communication methods retrieved")
		}
	}
}

func doReadRowErrorContactCommunicationMethodFields(ctx context.Context, input *contactcommunicationmethodfieldmodel.ContactCommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "field_code", "field_value"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetContactID(), data[0].GetContactCommunicationMethodID(), data[0].GetFieldCode(), data[0].GetFieldValue()).
			RowError(0, fmt.Errorf("DoReadAll row error")).
			AddRow(data[1].GetContactSystemCode(), data[1].GetContactID(), data[1].GetContactCommunicationMethodID(), data[1].GetFieldCode(), data[1].GetFieldValue()).
			AddRow(data[2].GetContactSystemCode(), data[2].GetContactID(), data[2].GetContactCommunicationMethodID(), data[2].GetFieldCode(), data[2].GetFieldValue())

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, contact_id, contact_communication_method_id, field_code, field_value 
			FROM contact_communication_method_field`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID())
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
			t.Fatalf("Expect contact communication methods is not nil")
		}

		if len(res) != 0 {
			t.Errorf("Expect response is nil")
		}
	}
}

func doReadExistingContactCommunicationMethodFields(ctx context.Context, input *contactcommunicationmethodfieldmodel.ContactCommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "field_code", "field_value"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetContactID(), data[0].GetContactCommunicationMethodID(), data[0].GetFieldCode(), data[0].GetFieldValue()).
			AddRow(data[1].GetContactSystemCode(), data[1].GetContactID(), data[1].GetContactCommunicationMethodID(), data[1].GetFieldCode(), data[1].GetFieldValue()).
			AddRow(data[2].GetContactSystemCode(), data[2].GetContactID(), data[2].GetContactCommunicationMethodID(), data[2].GetFieldCode(), data[2].GetFieldValue())

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, contact_id, contact_communication_method_id, field_code, field_value 
				FROM contact_communication_method_field`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID())
		if err != nil {
			t.Fatalf("Failed to read all contact communication method fields: %v", err)
		}

		if res == nil {
			t.Fatalf("Expect contact communication method fields is not nil")
		}

		if len(res) < 3 {
			t.Errorf("Expect there are contact communication method fields retrieved")
		}

		if res[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), res[0].GetContactSystemCode())
		}

		if res[0].GetContactID() != input.GetContactID() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactID(), res[0].GetContactID())
		}

		if res[0].GetContactCommunicationMethodID() != input.GetContactCommunicationMethodID() {
			t.Errorf("Expect contact communication method id %d, but got %d", input.GetContactCommunicationMethodID(), res[0].GetContactCommunicationMethodID())
		}

		if res[0].GetFieldCode() != input.GetFieldCode() {
			t.Errorf("Expect field code %s, but got %s", input.GetFieldCode(), res[0].GetFieldCode())
		}

		if res[0].GetFieldValue() != input.GetFieldValue() {
			t.Errorf("Expect field value %s, but got %s", input.GetFieldValue(), res[0].GetFieldValue())
		}
	}
}

func doSaveNewFailContactCommunicationMethodField(ctx context.Context, input *contactcommunicationmethodfieldmodel.ContactCommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		expInsQuery := mock.ExpectPrepare("INSERT INTO contact_communication_method_field").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetFieldCode(), input.GetFieldValue()).WillReturnError(fmt.Errorf("DoInsert contact communication method field failed"))

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

func doSaveNewContactCommunicationMethodField(ctx context.Context, input *contactcommunicationmethodfieldmodel.ContactCommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		expInsQuery := mock.ExpectPrepare("INSERT INTO contact_communication_method_field").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetFieldCode(), input.GetFieldValue()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoInsert(ctx, input)
		if err != nil {
			t.Fatalf("Failed to save contact communication method field: %v", err)
		}
	}
}

func doSaveExistingFailContactCommunicationMethodField(ctx context.Context, input *contactcommunicationmethodfieldmodel.ContactCommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		expUpdQuery := mock.ExpectPrepare("UPDATE contact_communication_method_field").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetFieldCode(), input.GetFieldValue()).WillReturnError(fmt.Errorf("DoUpdate contact communication method field failed"))

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

func doSaveExistingContactCommunicationMethodField(ctx context.Context, input *contactcommunicationmethodfieldmodel.ContactCommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		expUpdQuery := mock.ExpectPrepare("UPDATE contact_communication_method_field").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetFieldCode(), input.GetFieldValue()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoUpdate(ctx, input)
		if err != nil {
			t.Fatalf("Failed to save contact communication method: %v", err)
		}
	}
}

func doDeleteFailContactCommunicationMethodField(ctx context.Context, input *contactcommunicationmethodfieldmodel.ContactCommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM contact_communication_method_field").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnError(fmt.Errorf("Delete all contact communication method fields failed"))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID())
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

func doDeleteUnexistingContactCommunicationMethodField(ctx context.Context, input *contactcommunicationmethodfieldmodel.ContactCommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM contact_communication_method_field").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID())
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

func doDeleteExistingContactCommunicationMethodField(ctx context.Context, input *contactcommunicationmethodfieldmodel.ContactCommunicationMethodField) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM contact_communication_method_field").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID())
		if err != nil {
			t.Fatalf("Failed to delete all contact communication method fields: %v", err)
		}
	}
}

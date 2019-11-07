package contactcommunicationmethod

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	auditmodel "github.com/bungysheep/contact-management/pkg/models/v1/audit"
	contactcommunicationmethodmodel "github.com/bungysheep/contact-management/pkg/models/v1/contactcommunicationmethod"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ctx  context.Context
	repo IContactCommunicationMethodRepository
	db   *sql.DB
	mock sqlmock.Sqlmock
	data []*contactcommunicationmethodmodel.ContactCommunicationMethod
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	db, mock, _ = sqlmock.New()
	defer db.Close()

	repo = NewContactCommunicationMethodRepository(db)

	data = append(data, &contactcommunicationmethodmodel.ContactCommunicationMethod{
		ContactSystemCode:               "CNTSYS001",
		ContactID:                       1,
		ContactCommunicationMethodID:    1,
		CommunicationMethodCode:         "EMAIL",
		CommunicationMethodLabelCode:    "HOME",
		CommunicationMethodLabelCaption: "Home",
		FormatValue:                     "test@gmail.com",
		IsDefault:                       true,
	}, &contactcommunicationmethodmodel.ContactCommunicationMethod{
		ContactSystemCode:               "CNTSYS001",
		ContactID:                       1,
		ContactCommunicationMethodID:    2,
		CommunicationMethodCode:         "MOBILE",
		CommunicationMethodLabelCode:    "WORK",
		CommunicationMethodLabelCaption: "Work",
		FormatValue:                     "62-81234567890",
		IsDefault:                       false,
	}, &contactcommunicationmethodmodel.ContactCommunicationMethod{
		ContactSystemCode:               "CNTSYS001",
		ContactID:                       1,
		ContactCommunicationMethodID:    3,
		CommunicationMethodCode:         "FAX",
		CommunicationMethodLabelCode:    "SCHOOL",
		CommunicationMethodLabelCaption: "School",
		FormatValue:                     "62-2471234567",
		IsDefault:                       false,
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestContactCommunicationMethodRepository(t *testing.T) {
	t.Run("DoRead Contact Communication Method", doRead(ctx))

	t.Run("DoReadAll Contact Communication Method", doReadAll(ctx))

	t.Run("DoSave Contact Communication Method", doSave(ctx))

	t.Run("DoDelete Contact Communication Method", doDelete(ctx))

	t.Run("DoDeleteAll Contact Communication Method", doDeleteAll(ctx))
}

func doRead(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoRead fail", doReadFailContactCommunicationMethod(ctx, data[0]))

		t.Run("DoRead unexisting", doReadUnexistingContactCommunicationMethod(ctx, data[0]))

		t.Run("DoRead row error", doReadRowErrorContactCommunicationMethod(ctx, data[0]))

		t.Run("DoRead existing", doReadExistingContactCommunicationMethod(ctx, data[0]))
	}
}

func doReadAll(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoReadAll fail", doReadAllFailContactCommunicationMethods(ctx, data[0]))

		t.Run("DoReadAll unexisting", doReadAllUnexistingContactCommunicationMethods(ctx, data[0]))

		t.Run("DoReadAll row error", doReadAllRowErrorContactCommunicationMethods(ctx, data[0]))

		t.Run("DoReadAll existing", doReadAllExistingContactCommunicationMethods(ctx, data[0]))
	}
}

func doSave(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoSave new fail", doSaveNewFailContactCommunicationMethod(ctx, data[0]))

		t.Run("DoSave new", doSaveNewContactCommunicationMethod(ctx, data[0]))

		t.Run("DoSave existing fail", doSaveExistingFailContactCommunicationMethod(ctx, data[0]))

		t.Run("DoSave existing", doSaveExistingContactCommunicationMethod(ctx, data[0]))
	}
}

func doDelete(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoDelete fail", doDeleteFailContactCommunicationMethod(ctx, data[1]))

		t.Run("DoDelete default", doDeleteDefaultContactCommunicationMethod(ctx, data[0]))

		t.Run("DoDelete unexisting", doDeleteUnexistingContactCommunicationMethod(ctx, data[1]))

		t.Run("DoDelete existing", doDeleteExistingContactCommunicationMethod(ctx, data[1]))
	}
}

func doDeleteAll(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoDeleteAll fail", doDeleteAllFailContactCommunicationMethod(ctx, data[0]))

		t.Run("DoDeleteAll unexisting", doDeleteAllUnexistingContactCommunicationMethod(ctx, data[0]))

		t.Run("DoDeleteAll existing", doDeleteAllExistingContactCommunicationMethod(ctx, data[0]))
	}
}

func doReadFailContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, communication_method_label_code, format_value, is_default, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnError(fmt.Errorf("DoRead contact communication method failed"))

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

		if res != nil {
			t.Errorf("Expect contact communication method is nil")
		}
	}
}

func doReadUnexistingContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "communication_method_label_code", "format_value", "is_default", "created_at", "modified_at", "vers"})

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, communication_method_label_code, format_value, is_default, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
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

		if res != nil {
			t.Errorf("Expect contact communication method is nil")
		}
	}
}

func doReadRowErrorContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "communication_method_label_code", "format_value", "is_default", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetFormatValue(), input.GetIsDefault(), tmNow, tmNow, 1).
			RowError(0, fmt.Errorf("DoRead row error"))

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, communication_method_label_code, format_value, is_default, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
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

		if res != nil {
			t.Errorf("Expect contact communication method is nil")
		}
	}
}

func doReadExistingContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "communication_method_label_code", "format_value", "is_default", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetFormatValue(), input.GetIsDefault(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, communication_method_label_code, format_value, is_default, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID())
		if err != nil {
			t.Errorf("Failed to read contact communication method: %v", err)
		}

		if res == nil {
			t.Errorf("Expect contact communication method is not nil")
		}

		if res.GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact contact system code %s, but got %s", input.GetContactSystemCode(), res.GetContactSystemCode())
		}

		if res.GetContactID() != input.GetContactID() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactID(), res.GetContactID())
		}

		if res.GetContactCommunicationMethodID() != input.GetContactCommunicationMethodID() {
			t.Errorf("Expect contact communication method id %d, but got %d", input.GetContactCommunicationMethodID(), res.GetContactCommunicationMethodID())
		}

		if res.GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), res.GetCommunicationMethodCode())
		}

		if res.GetCommunicationMethodLabelCode() != input.GetCommunicationMethodLabelCode() {
			t.Errorf("Expect communication method label code %s, but got %s", input.GetCommunicationMethodLabelCode(), res.GetCommunicationMethodLabelCode())
		}

		if res.GetFormatValue() != input.GetFormatValue() {
			t.Errorf("Expect format value %s, but got %s", input.GetFormatValue(), res.GetFormatValue())
		}

		if res.GetIsDefault() != input.GetIsDefault() {
			t.Errorf("Expect default %v, but got %v", input.GetIsDefault(), res.GetIsDefault())
		}
	}
}

func doReadAllFailContactCommunicationMethods(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, communication_method_label_code, format_value, is_default, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnError(fmt.Errorf("DoReadAll contact communication method failed"))

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode(), input.GetContactID())
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

		if len(res) != 0 {
			t.Errorf("Expect response is nil")
		}
	}
}

func doReadAllUnexistingContactCommunicationMethods(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "communication_method_label_code", "format_value", "is_default", "created_at", "modified_at", "vers"})

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, communication_method_label_code, format_value, is_default, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnRows(rows)

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode(), input.GetContactID())
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
			t.Errorf("Expect contact communication methods is not nil")
		}

		if len(res) != 0 {
			t.Errorf("Expect no contact communication methods retrieved")
		}
	}
}

func doReadAllRowErrorContactCommunicationMethods(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "communication_method_label_code", "format_value", "is_default", "created_at", "modified_at", "vers"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetContactID(), data[0].GetContactCommunicationMethodID(), data[0].GetCommunicationMethodCode(), data[0].GetCommunicationMethodLabelCode(), data[0].GetFormatValue(), data[0].GetIsDefault(), tmNow, tmNow, 1).
			RowError(0, fmt.Errorf("DoReadAll row error")).
			AddRow(data[1].GetContactSystemCode(), data[1].GetContactID(), data[1].GetContactCommunicationMethodID(), data[1].GetCommunicationMethodCode(), data[1].GetCommunicationMethodLabelCode(), data[1].GetFormatValue(), data[1].GetIsDefault(), tmNow, tmNow, 1).
			AddRow(data[2].GetContactSystemCode(), data[2].GetContactID(), data[2].GetContactCommunicationMethodID(), data[2].GetCommunicationMethodCode(), data[2].GetCommunicationMethodLabelCode(), data[2].GetFormatValue(), data[2].GetIsDefault(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, communication_method_label_code, format_value, is_default, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnRows(rows)

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode(), input.GetContactID())
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

		if len(res) != 0 {
			t.Errorf("Expect response is nil")
		}
	}
}

func doReadAllExistingContactCommunicationMethods(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "communication_method_label_code", "format_value", "is_default", "created_at", "modified_at", "vers"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetContactID(), data[0].GetContactCommunicationMethodID(), data[0].GetCommunicationMethodCode(), data[0].GetCommunicationMethodLabelCode(), data[0].GetFormatValue(), data[0].GetIsDefault(), tmNow, tmNow, 1).
			AddRow(data[1].GetContactSystemCode(), data[1].GetContactID(), data[1].GetContactCommunicationMethodID(), data[1].GetCommunicationMethodCode(), data[1].GetCommunicationMethodLabelCode(), data[1].GetFormatValue(), data[1].GetIsDefault(), tmNow, tmNow, 1).
			AddRow(data[2].GetContactSystemCode(), data[2].GetContactID(), data[2].GetContactCommunicationMethodID(), data[2].GetCommunicationMethodCode(), data[2].GetCommunicationMethodLabelCode(), data[2].GetFormatValue(), data[2].GetIsDefault(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, communication_method_label_code, format_value, is_default, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnRows(rows)

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode(), input.GetContactID())
		if err != nil {
			t.Errorf("Failed to read all contact communication methods: %v", err)
		}

		if res == nil {
			t.Errorf("Expect contact communication methods is not nil")
		}

		if len(res) < 3 {
			t.Errorf("Expect there are contact communication methods retrieved")
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

		if res[0].GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), res[0].GetCommunicationMethodCode())
		}

		if res[0].GetCommunicationMethodLabelCode() != input.GetCommunicationMethodLabelCode() {
			t.Errorf("Expect communication method label code %s, but got %s", input.GetCommunicationMethodLabelCode(), res[0].GetCommunicationMethodLabelCode())
		}

		if res[0].GetFormatValue() != input.GetFormatValue() {
			t.Errorf("Expect format value %s, but got %s", input.GetFormatValue(), res[0].GetFormatValue())
		}

		if res[0].GetIsDefault() != input.GetIsDefault() {
			t.Errorf("Expect default %v, but got %v", input.GetIsDefault(), res[0].GetIsDefault())
		}
	}
}

func doSaveNewFailContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		}

		expInsQuery := mock.ExpectPrepare("INSERT INTO contact_communication_method").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetFormatValue(), input.GetIsDefault(), tmNow, tmNow).WillReturnError(fmt.Errorf("DoInsert contact communication method failed"))

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

func doSaveNewContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		}

		expInsQuery := mock.ExpectPrepare("INSERT INTO contact_communication_method").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetFormatValue(), input.GetIsDefault(), tmNow, tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoInsert(ctx, input)
		if err != nil {
			t.Errorf("Failed to save contact communication method: %v", err)
		}
	}
}

func doSaveExistingFailContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       2,
		}

		expUpdQuery := mock.ExpectPrepare("UPDATE contact_communication_method").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetFormatValue(), input.GetIsDefault(), tmNow).WillReturnError(fmt.Errorf("DoUpdate contact communication method failed"))

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

func doSaveExistingContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       2,
		}

		expUpdQuery := mock.ExpectPrepare("UPDATE contact_communication_method").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetFormatValue(), input.GetIsDefault(), tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoUpdate(ctx, input)
		if err != nil {
			t.Errorf("Failed to save contact communication method: %v", err)
		}
	}
}

func doDeleteFailContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "communication_method_label_code", "format_value", "is_default", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetFormatValue(), input.GetIsDefault(), tmNow, tmNow, 1)

		expDefCommMethodQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, communication_method_label_code, format_value, is_default, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
		expDefCommMethodQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnRows(rows)

		expQuery := mock.ExpectPrepare("DELETE FROM contact_communication_method").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnError(fmt.Errorf("Delete contact communication method failed"))

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

func doDeleteDefaultContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "communication_method_label_code", "format_value", "is_default", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetFormatValue(), input.GetIsDefault(), tmNow, tmNow, 1)

		expDefCommMethodQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, communication_method_label_code, format_value, is_default, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
		expDefCommMethodQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnRows(rows)

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

func doDeleteUnexistingContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "communication_method_label_code", "format_value", "is_default", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetFormatValue(), input.GetIsDefault(), tmNow, tmNow, 1)

		expDefCommMethodQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, communication_method_label_code, format_value, is_default, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
		expDefCommMethodQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnRows(rows)

		expQuery := mock.ExpectPrepare("DELETE FROM contact_communication_method").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID())
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

func doDeleteExistingContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "communication_method_label_code", "format_value", "is_default", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetFormatValue(), input.GetIsDefault(), tmNow, tmNow, 1)

		expDefCommMethodQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, communication_method_label_code, format_value, is_default, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
		expDefCommMethodQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnRows(rows)

		expQuery := mock.ExpectPrepare("DELETE FROM contact_communication_method").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID())
		if err != nil {
			t.Errorf("Failed to delete contact communication method: %v", err)
		}
	}
}

func doDeleteAllFailContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM contact_communication_method").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnError(fmt.Errorf("Delete all contact communication methods failed"))

		err := repo.DoDeleteAll(ctx, input.GetContactSystemCode(), input.GetContactID())
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

func doDeleteAllUnexistingContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM contact_communication_method").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.DoDeleteAll(ctx, input.GetContactSystemCode(), input.GetContactID())
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

func doDeleteAllExistingContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM contact_communication_method").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoDeleteAll(ctx, input.GetContactSystemCode(), input.GetContactID())
		if err != nil {
			t.Errorf("Failed to delete all contact communication methods: %v", err)
		}
	}
}

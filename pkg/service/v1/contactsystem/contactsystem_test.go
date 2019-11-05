package contactsystem

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	auditmodel "github.com/bungysheep/contact-management/pkg/models/v1/audit"
	contactsystemmodel "github.com/bungysheep/contact-management/pkg/models/v1/contactsystem"
)

var (
	ctx  context.Context
	svc  IContactSystemService
	db   *sql.DB
	mock sqlmock.Sqlmock
	data []*contactsystemmodel.ContactSystem
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	db, mock, _ = sqlmock.New()
	defer db.Close()

	svc = NewContactSystemService(db)

	data = append(data, &contactsystemmodel.ContactSystem{
		ContactSystemCode: "CNTSYS001",
		Description:       "Contact System 1",
		Details:           "Contact System 1",
		Status:            "A",
	}, &contactsystemmodel.ContactSystem{
		ContactSystemCode: "CNTSYS002",
		Description:       "Contact System 2",
		Details:           "Contact System 2",
		Status:            "A",
	}, &contactsystemmodel.ContactSystem{
		ContactSystemCode: "CNTSYS003",
		Description:       "Contact System 3",
		Details:           "Contact System 3",
		Status:            "A",
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestContactSystemService(t *testing.T) {
	t.Run("DoRead Contact System", doRead(ctx, data[0]))

	t.Run("DoReadAll Contact System", doReadAll(ctx, data[0]))

	t.Run("DoSave new Contact System", doSaveNew(ctx, data[0]))

	t.Run("DoSave existing Contact System", doSaveExisting(ctx, data[0]))

	t.Run("DoDelete Contact System", doDelete(ctx, data[0]))
}

func doRead(ctx context.Context, input *contactsystemmodel.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "description", "details", "status", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, description, details, status, created_at, modified_at, vers FROM contact_system").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(rows)

		resp, err := svc.DoRead(ctx, input.GetContactSystemCode())
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}

		if resp == nil {
			t.Errorf("Expect contact system is not nil")
		}

		if resp.GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetContactSystemCode())
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
	}
}

func doReadAll(ctx context.Context, input *contactsystemmodel.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "description", "details", "status", "created_at", "modified_at", "vers"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetDescription(), data[0].GetDetails(), data[0].GetStatus(), tmNow, tmNow, 1).
			AddRow(data[1].GetContactSystemCode(), data[1].GetDescription(), data[1].GetDetails(), data[1].GetStatus(), tmNow, tmNow, 1).
			AddRow(data[2].GetContactSystemCode(), data[2].GetDescription(), data[2].GetDetails(), data[2].GetStatus(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, description, details, status, created_at, modified_at, vers FROM contact_system").ExpectQuery()
		expQuery.WillReturnRows(rows)

		resp, err := svc.DoReadAll(ctx)
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}

		if resp == nil {
			t.Errorf("Expect contact system is not nil")
		}

		if len(resp) < 3 {
			t.Errorf("Expect there are contact systems retrieved")
		}

		if resp[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp[0].GetContactSystemCode())
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
	}
}

func doSaveNew(ctx context.Context, input *contactsystemmodel.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		}

		expUpdQuery := mock.ExpectPrepare("UPDATE contact_system").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), tmNow).WillReturnResult(sqlmock.NewResult(0, 0))

		expInsQuery := mock.ExpectPrepare("INSERT INTO contact_system").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), tmNow, tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoSave(ctx, input)
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}
	}
}

func doSaveExisting(ctx context.Context, input *contactsystemmodel.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       2,
		}

		expUpdQuery := mock.ExpectPrepare("UPDATE contact_system").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoSave(ctx, input)
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}
	}
}

func doDelete(ctx context.Context, input *contactsystemmodel.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		expCommMethodQuery := mock.ExpectPrepare("SELECT 1 FROM communication_method").ExpectQuery()
		expCommMethodQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(sqlmock.NewRows([]string{"exists"}))

		expContactQuery := mock.ExpectPrepare("SELECT 1 FROM contact").ExpectQuery()
		expContactQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(sqlmock.NewRows([]string{"exists"}))

		expQuery := mock.ExpectPrepare("DELETE FROM contact_system").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoDelete(ctx, input.GetContactSystemCode())
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}
	}
}

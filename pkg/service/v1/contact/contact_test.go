package contact

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	auditmodel "github.com/bungysheep/contact-management/pkg/models/v1/audit"
	contactmodel "github.com/bungysheep/contact-management/pkg/models/v1/contact"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ctx  context.Context
	svc  IContactService
	db   *sql.DB
	mock sqlmock.Sqlmock
	data []*contactmodel.Contact
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	db, mock, _ = sqlmock.New()
	defer db.Close()

	svc = NewContactService(db)

	data = append(data, &contactmodel.Contact{
		ContactSystemCode: "CNTSYS001",
		ContactID:         1,
		FirstName:         "James",
		LastName:          "Embongbulan",
		Status:            "A",
	}, &contactmodel.Contact{
		ContactSystemCode: "CNTSYS001",
		ContactID:         2,
		FirstName:         "Rindi",
		LastName:          "Allorerung",
		Status:            "A",
	}, &contactmodel.Contact{
		ContactSystemCode: "CNTSYS001",
		ContactID:         3,
		FirstName:         "Marvel",
		LastName:          "Embongbulan",
		Status:            "A",
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestContactService(t *testing.T) {
	t.Run("DoRead Communication Method", doRead(ctx, data[0]))

	t.Run("DoReadAll Communication Method", doReadAll(ctx, data[0]))

	t.Run("DoSave Communication Method", doSave(ctx, data[0]))

	t.Run("DoDelete Communication Method", doDelete(ctx, data[0]))
}

func doRead(ctx context.Context, input *contactmodel.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "first_name", "last_name", "status", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), input.GetFirstName(), input.GetLastName(), input.GetStatus(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, first_name, last_name, status, created_at, modified_at, vers FROM contact").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnRows(rows)

		resp, err := svc.DoRead(ctx, input.GetContactSystemCode(), input.GetContactID())
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}

		if resp == nil {
			t.Errorf("Expect contact is not nil")
		}

		if resp.GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetContactSystemCode())
		}

		if resp.GetContactID() != input.GetContactID() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactID(), resp.GetContactID())
		}

		if resp.GetFirstName() != input.GetFirstName() {
			t.Errorf("Expect firstname %s, but got %s", input.GetFirstName(), resp.GetFirstName())
		}

		if resp.GetLastName() != input.GetLastName() {
			t.Errorf("Expect lastname %s, but got %s", input.GetLastName(), resp.GetLastName())
		}

		if resp.GetStatus() != input.GetStatus() {
			t.Errorf("Expect status %s, but got %s", input.GetStatus(), resp.GetStatus())
		}
	}
}

func doReadAll(ctx context.Context, input *contactmodel.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "first_name", "last_name", "status", "created_at", "modified_at", "vers"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetContactID(), data[0].GetFirstName(), data[0].GetLastName(), data[0].GetStatus(), tmNow, tmNow, 1).
			AddRow(data[1].GetContactSystemCode(), data[1].GetContactID(), data[1].GetFirstName(), data[1].GetLastName(), data[1].GetStatus(), tmNow, tmNow, 1).
			AddRow(data[2].GetContactSystemCode(), data[2].GetContactID(), data[2].GetFirstName(), data[2].GetLastName(), data[2].GetStatus(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, first_name, last_name, status, created_at, modified_at, vers FROM contact").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(rows)

		resp, err := svc.DoReadAll(ctx, input.GetContactSystemCode())
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}

		if resp == nil {
			t.Errorf("Expect contact is not nil")
		}

		if len(resp) < 3 {
			t.Errorf("Expect there are contacts retrieved")
		}

		if resp[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp[0].GetContactSystemCode())
		}

		if resp[0].GetContactID() != input.GetContactID() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactID(), resp[0].GetContactID())
		}

		if resp[0].GetFirstName() != input.GetFirstName() {
			t.Errorf("Expect firstname %s, but got %s", input.GetFirstName(), resp[0].GetFirstName())
		}

		if resp[0].GetLastName() != input.GetLastName() {
			t.Errorf("Expect lastname %s, but got %s", input.GetLastName(), resp[0].GetLastName())
		}

		if resp[0].GetStatus() != input.GetStatus() {
			t.Errorf("Expect status %s, but got %s", input.GetStatus(), resp[0].GetStatus())
		}
	}
}

func doSave(ctx context.Context, input *contactmodel.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoSave invalid Contact System", doSaveInvalidContactSystem(ctx, input))

		t.Run("DoSave new Contact", doSaveNew(ctx, input))

		t.Run("DoSave existing Contact", doSaveExisting(ctx, input))
	}
}

func doSaveInvalidContactSystem(ctx context.Context, input *contactmodel.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "description", "details", "status", "created_at", "modified_at", "vers"})

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, description, details, status, created_at, modified_at, vers FROM contact_system").ExpectQuery()
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

func doSaveNew(ctx context.Context, input *contactmodel.Contact) func(t *testing.T) {
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

		expUpdQuery := mock.ExpectPrepare("UPDATE contact").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetFirstName(), input.GetLastName(), input.GetStatus(), tmNow).WillReturnResult(sqlmock.NewResult(0, 0))

		expInsQuery := mock.ExpectPrepare("INSERT INTO contact").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetFirstName(), input.GetLastName(), input.GetStatus(), tmNow, tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoSave(ctx, input)
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}
	}
}

func doSaveExisting(ctx context.Context, input *contactmodel.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       2,
		}

		rows := sqlmock.NewRows([]string{"contact_system_code", "description", "details", "status", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), "", "", "", tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, description, details, status, created_at, modified_at, vers FROM contact_system").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(rows)

		expUpdQuery := mock.ExpectPrepare("UPDATE contact").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetFirstName(), input.GetLastName(), input.GetStatus(), tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoSave(ctx, input)
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}
	}
}

func doDelete(ctx context.Context, input *contactmodel.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoSave fail Contact Communication Method", doDeleteFailContactCommunicationMethod(ctx, input))

		t.Run("DoSave existing Contact", doDeleteExisting(ctx, input))
	}
}

func doDeleteFailContactCommunicationMethod(ctx context.Context, input *contactmodel.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM contact_communication_method").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnError(fmt.Errorf("Delete all contact communication methods failed"))

		err := svc.DoDelete(ctx, input.GetContactSystemCode(), input.GetContactID())
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

func doDeleteExisting(ctx context.Context, input *contactmodel.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		expContactCommMethodQuery := mock.ExpectPrepare("DELETE FROM contact_communication_method").ExpectExec()
		expContactCommMethodQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnResult(sqlmock.NewResult(0, 1))

		expContactQuery := mock.ExpectPrepare("DELETE FROM contact").ExpectExec()
		expContactQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoDelete(ctx, input.GetContactSystemCode(), input.GetContactID())
		if err != nil {
			t.Errorf("Expect error is nil, but got %v", err)
		}
	}
}

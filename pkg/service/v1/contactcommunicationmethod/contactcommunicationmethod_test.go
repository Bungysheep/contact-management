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
	contactcommunicationmethodfieldmodel "github.com/bungysheep/contact-management/pkg/models/v1/contactcommunicationmethodfield"
)

var (
	ctx  context.Context
	svc  IContactCommunicationMethodService
	db   *sql.DB
	mock sqlmock.Sqlmock
	data []*contactcommunicationmethodmodel.ContactCommunicationMethod
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	db, mock, _ = sqlmock.New()
	defer db.Close()

	svc = NewContactCommunicationMethodService(db)

	data = append(data, &contactcommunicationmethodmodel.ContactCommunicationMethod{
		ContactSystemCode:               "CNTSYS001",
		ContactID:                       1,
		ContactCommunicationMethodID:    1,
		CommunicationMethodCode:         "EMAIL",
		CommunicationMethodLabelCode:    "HOME",
		CommunicationMethodLabelCaption: "Home",
		FormatValue:                     "test@gmail.com",
		Status:                          "A",
		IsDefault:                       true,
		ContactCommunicationMethodField: []*contactcommunicationmethodfieldmodel.ContactCommunicationMethodField{
			&contactcommunicationmethodfieldmodel.ContactCommunicationMethodField{
				ContactSystemCode:            "CNTSYS001",
				ContactID:                    1,
				ContactCommunicationMethodID: 1,
				FieldCode:                    "EMAIL_ADDRESS",
				FieldValue:                   "test@gmail.com",
			},
		},
	}, &contactcommunicationmethodmodel.ContactCommunicationMethod{
		ContactSystemCode:               "CNTSYS001",
		ContactID:                       1,
		ContactCommunicationMethodID:    2,
		CommunicationMethodCode:         "MOBILE",
		CommunicationMethodLabelCode:    "WORK",
		CommunicationMethodLabelCaption: "Work",
		FormatValue:                     "62-81234567890",
		Status:                          "A",
		IsDefault:                       false,
		ContactCommunicationMethodField: []*contactcommunicationmethodfieldmodel.ContactCommunicationMethodField{
			&contactcommunicationmethodfieldmodel.ContactCommunicationMethodField{
				ContactSystemCode:            "CNTSYS001",
				ContactID:                    1,
				ContactCommunicationMethodID: 2,
				FieldCode:                    "MOBILE_NO",
				FieldValue:                   "62-81234567890",
			},
		},
	}, &contactcommunicationmethodmodel.ContactCommunicationMethod{
		ContactSystemCode:               "CNTSYS001",
		ContactID:                       1,
		ContactCommunicationMethodID:    3,
		CommunicationMethodCode:         "FAX",
		CommunicationMethodLabelCode:    "SCHOOL",
		CommunicationMethodLabelCaption: "School",
		FormatValue:                     "62-2471234567",
		Status:                          "A",
		IsDefault:                       false,
		ContactCommunicationMethodField: []*contactcommunicationmethodfieldmodel.ContactCommunicationMethodField{
			&contactcommunicationmethodfieldmodel.ContactCommunicationMethodField{
				ContactSystemCode:            "CNTSYS001",
				ContactID:                    1,
				ContactCommunicationMethodID: 3,
				FieldCode:                    "FAX_NO",
				FieldValue:                   "62-2471234567",
			},
		},
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestContactCommunicationMethodService(t *testing.T) {
	t.Run("DoRead Contact Communication Method", doRead(ctx, data[0]))

	t.Run("DoReadAll Contact Communication Method", doReadAll(ctx, data[0]))

	t.Run("DoSave Contact Communication Method", doSave(ctx, data[0]))

	t.Run("DoDelete Contact Communication Method", doDelete(ctx, data[1]))
}

func doRead(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoRead fail Contact Communication Method", doReadFail(ctx, data[0]))

		t.Run("DoRead fail Contact Communication Method Field", doReadFailField(ctx, data[0]))

		t.Run("DoRead existing Contact Communication Method", doReadExisting(ctx, data[0]))
	}
}

func doReadFail(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare(
			`SELECT ccm.contact_system_code, ccm.contact_id, ccm.contact_communication_method_id, 
				ccm.communication_method_code, ccm.communication_method_label_code, cml.caption, ccm.format_value, ccm.status, ccm.is_default, 
				ccm.created_at, ccm.modified_at, ccm.vers 
			FROM contact_communication_method ccm
			INNER JOIN communication_method_label cml`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnError(fmt.Errorf("DoRead contact communication method failed"))

		resp, err := svc.DoRead(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID())
		if err == nil {
			t.Errorf("Expect error is not nil")
		}

		if resp != nil {
			t.Errorf("Expect contact communication method is nil")
		}
	}
}

func doReadFailField(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "communication_method_label_code", "communication_method_label_code", "format_value", "status", "is_default", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetCommunicationMethodLabelCaption(), input.GetFormatValue(), input.GetStatus(), input.GetIsDefault(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare(
			`SELECT ccm.contact_system_code, ccm.contact_id, ccm.contact_communication_method_id, 
				ccm.communication_method_code, ccm.communication_method_label_code, cml.caption, ccm.format_value, ccm.status, ccm.is_default, 
				ccm.created_at, ccm.modified_at, ccm.vers 
			FROM contact_communication_method ccm
			INNER JOIN communication_method_label cml`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnRows(rows)

		expQuery = mock.ExpectPrepare(
			`SELECT contact_system_code, contact_id, contact_communication_method_id, field_code, field_value 
			FROM contact_communication_method_field`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnError(fmt.Errorf("DoRead contact communication method field failed"))

		resp, err := svc.DoRead(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID())
		if err == nil {
			t.Errorf("Expect error is not nil")
		}

		if resp == nil {
			t.Fatalf("Expect contact communication method is not nil")
		}

		if resp.GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetContactSystemCode())
		}

		if resp.GetContactID() != input.GetContactID() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactID(), resp.GetContactID())
		}

		if resp.GetContactCommunicationMethodID() != input.GetContactCommunicationMethodID() {
			t.Errorf("Expect contact communication method id %d, but got %d", input.GetContactCommunicationMethodID(), resp.GetContactCommunicationMethodID())
		}

		if resp.GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp.GetCommunicationMethodCode())
		}

		if resp.GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp.GetCommunicationMethodCode())
		}

		if resp.GetCommunicationMethodLabelCode() != input.GetCommunicationMethodLabelCode() {
			t.Errorf("Expect communication method label code %s, but got %s", input.GetCommunicationMethodLabelCode(), resp.GetCommunicationMethodLabelCode())
		}

		if resp.GetCommunicationMethodLabelCaption() != input.GetCommunicationMethodLabelCaption() {
			t.Errorf("Expect communication method label caption %s, but got %s", input.GetCommunicationMethodLabelCaption(), resp.GetCommunicationMethodLabelCaption())
		}

		if resp.GetFormatValue() != input.GetFormatValue() {
			t.Errorf("Expect format value %s, but got %s", input.GetFormatValue(), resp.GetFormatValue())
		}

		if resp.GetStatus() != input.GetStatus() {
			t.Errorf("Expect status %s, but got %s", input.GetStatus(), resp.GetStatus())
		}

		if resp.GetIsDefault() != input.GetIsDefault() {
			t.Errorf("Expect default %v, but got %v", input.GetIsDefault(), resp.GetIsDefault())
		}

		if resp.GetContactCommunicationMethodField() == nil {
			t.Fatalf("Expect contact communication method fields is not nil")
		}

		if len(resp.GetContactCommunicationMethodField()) != 0 {
			t.Errorf("Expect no contact communication method fields")
		}
	}
}

func doReadExisting(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "communication_method_label_code", "communication_method_label_code", "format_value", "status", "is_default", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetCommunicationMethodLabelCaption(), input.GetFormatValue(), input.GetStatus(), input.GetIsDefault(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare(
			`SELECT ccm.contact_system_code, ccm.contact_id, ccm.contact_communication_method_id, 
				ccm.communication_method_code, ccm.communication_method_label_code, cml.caption, ccm.format_value, ccm.status, ccm.is_default, 
				ccm.created_at, ccm.modified_at, ccm.vers 
			FROM contact_communication_method ccm
			INNER JOIN communication_method_label cml`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnRows(rows)

		rows = sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "field_code", "field_value"}).
			AddRow(input.GetContactCommunicationMethodField()[0].GetContactSystemCode(), input.GetContactCommunicationMethodField()[0].GetContactID(), input.GetContactCommunicationMethodField()[0].GetContactCommunicationMethodID(), input.GetContactCommunicationMethodField()[0].GetFieldCode(), input.GetContactCommunicationMethodField()[0].GetFieldValue())

		expQuery = mock.ExpectPrepare(
			`SELECT contact_system_code, contact_id, contact_communication_method_id, field_code, field_value 
				FROM contact_communication_method_field`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnRows(rows)

		resp, err := svc.DoRead(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID())
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}

		if resp == nil {
			t.Fatalf("Expect contact communication method is not nil")
		}

		if resp.GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp.GetContactSystemCode())
		}

		if resp.GetContactID() != input.GetContactID() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactID(), resp.GetContactID())
		}

		if resp.GetContactCommunicationMethodID() != input.GetContactCommunicationMethodID() {
			t.Errorf("Expect contact communication method id %d, but got %d", input.GetContactCommunicationMethodID(), resp.GetContactCommunicationMethodID())
		}

		if resp.GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp.GetCommunicationMethodCode())
		}

		if resp.GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp.GetCommunicationMethodCode())
		}

		if resp.GetCommunicationMethodLabelCode() != input.GetCommunicationMethodLabelCode() {
			t.Errorf("Expect communication method label code %s, but got %s", input.GetCommunicationMethodLabelCode(), resp.GetCommunicationMethodLabelCode())
		}

		if resp.GetCommunicationMethodLabelCaption() != input.GetCommunicationMethodLabelCaption() {
			t.Errorf("Expect communication method label caption %s, but got %s", input.GetCommunicationMethodLabelCaption(), resp.GetCommunicationMethodLabelCaption())
		}

		if resp.GetFormatValue() != input.GetFormatValue() {
			t.Errorf("Expect format value %s, but got %s", input.GetFormatValue(), resp.GetFormatValue())
		}

		if resp.GetIsDefault() != input.GetIsDefault() {
			t.Errorf("Expect default %v, but got %v", input.GetIsDefault(), resp.GetIsDefault())
		}

		if resp.GetStatus() != input.GetStatus() {
			t.Errorf("Expect status %s, but got %s", input.GetStatus(), resp.GetStatus())
		}

		if resp.GetContactCommunicationMethodField() == nil {
			t.Fatalf("Expect contact communication method fields is not nil")
		}

		if len(resp.GetContactCommunicationMethodField()) < 1 {
			t.Errorf("Expect there are contact communication method fields retrieved")
		}

		if resp.GetContactCommunicationMethodField()[0].GetContactSystemCode() != input.GetContactCommunicationMethodField()[0].GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactCommunicationMethodField()[0].GetContactSystemCode(), resp.GetContactCommunicationMethodField()[0].GetContactSystemCode())
		}

		if resp.GetContactCommunicationMethodField()[0].GetContactID() != input.GetContactCommunicationMethodField()[0].GetContactID() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactCommunicationMethodField()[0].GetContactID(), resp.GetContactCommunicationMethodField()[0].GetContactID())
		}

		if resp.GetContactCommunicationMethodField()[0].GetContactCommunicationMethodID() != input.GetContactCommunicationMethodID() {
			t.Errorf("Expect contact communication method id %d, but got %d", input.GetContactCommunicationMethodID(), resp.GetContactCommunicationMethodField()[0].GetContactCommunicationMethodID())
		}

		if resp.GetContactCommunicationMethodField()[0].GetFieldCode() != input.GetContactCommunicationMethodField()[0].GetFieldCode() {
			t.Errorf("Expect field code %s, but got %s", input.GetContactCommunicationMethodField()[0].GetFieldCode(), resp.GetContactCommunicationMethodField()[0].GetFieldCode())
		}

		if resp.GetContactCommunicationMethodField()[0].GetFieldValue() != input.GetContactCommunicationMethodField()[0].GetFieldValue() {
			t.Errorf("Expect field value %s, but got %s", input.GetContactCommunicationMethodField()[0].GetFieldValue(), resp.GetContactCommunicationMethodField()[0].GetFieldValue())
		}
	}
}

func doReadAll(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoReadAll fail Contact Communication Method", doReadAllFail(ctx, data[0]))

		t.Run("DoReadAll fail Contact Communication Method Field", doReadAllFailField(ctx, data[0]))

		t.Run("DoReadAll exising Contact Communication Method", doReadAllExisting(ctx, data[0]))
	}
}

func doReadAllFail(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare(
			`SELECT ccm.contact_system_code, ccm.contact_id, ccm.contact_communication_method_id, 
				ccm.communication_method_code, ccm.communication_method_label_code, cml.caption, ccm.format_value, ccm.status, ccm.is_default, 
				ccm.created_at, ccm.modified_at, ccm.vers 
			FROM contact_communication_method ccm
			INNER JOIN communication_method_label cml`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnError(fmt.Errorf("DoReadAll contact communication method failed"))

		resp, err := svc.DoReadAll(ctx, input.GetContactSystemCode(), input.GetContactID())
		if err == nil {
			t.Errorf("Expect error is not nil")
		}

		if resp == nil {
			t.Fatalf("Expect contact communication method is not nil")
		}

		if len(resp) != 0 {
			t.Errorf("Expect response is nil")
		}
	}
}

func doReadAllFailField(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "communication_method_label_code", "communication_method_label_code", "format_value", "status", "is_default", "created_at", "modified_at", "vers"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetContactID(), data[0].GetContactCommunicationMethodID(), data[0].GetCommunicationMethodCode(), data[0].GetCommunicationMethodLabelCode(), data[0].GetCommunicationMethodLabelCaption(), data[0].GetFormatValue(), data[0].GetStatus(), data[0].GetIsDefault(), tmNow, tmNow, 1).
			AddRow(data[1].GetContactSystemCode(), data[1].GetContactID(), data[1].GetContactCommunicationMethodID(), data[1].GetCommunicationMethodCode(), data[1].GetCommunicationMethodLabelCode(), data[1].GetCommunicationMethodLabelCaption(), data[1].GetFormatValue(), data[1].GetStatus(), data[1].GetIsDefault(), tmNow, tmNow, 1).
			AddRow(data[2].GetContactSystemCode(), data[2].GetContactID(), data[2].GetContactCommunicationMethodID(), data[2].GetCommunicationMethodCode(), data[2].GetCommunicationMethodLabelCode(), data[2].GetCommunicationMethodLabelCaption(), data[2].GetFormatValue(), data[2].GetStatus(), data[2].GetIsDefault(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare(
			`SELECT ccm.contact_system_code, ccm.contact_id, ccm.contact_communication_method_id, 
				ccm.communication_method_code, ccm.communication_method_label_code, cml.caption, ccm.format_value, ccm.status, ccm.is_default, 
				ccm.created_at, ccm.modified_at, ccm.vers 
			FROM contact_communication_method ccm
			INNER JOIN communication_method_label cml`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnRows(rows)

		expQuery = mock.ExpectPrepare(
			`SELECT contact_system_code, contact_id, contact_communication_method_id, field_code, field_value 
			FROM contact_communication_method_field`).ExpectQuery()
		expQuery.WithArgs(data[0].GetContactSystemCode(), data[0].GetContactID(), data[0].GetContactCommunicationMethodID()).WillReturnError(fmt.Errorf("DoRead contact communication method field failed"))

		resp, err := svc.DoReadAll(ctx, input.GetContactSystemCode(), input.GetContactID())
		if err == nil {
			t.Errorf("Expect error is not nil")
		}

		if resp == nil {
			t.Fatalf("Expect contact communication method is not nil")
		}

		if len(resp) < 3 {
			t.Errorf("Expect there are contact communication methods retrieved")
		}

		if resp[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp[0].GetContactSystemCode())
		}

		if resp[0].GetContactID() != input.GetContactID() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactID(), resp[0].GetContactID())
		}

		if resp[0].GetContactCommunicationMethodID() != input.GetContactCommunicationMethodID() {
			t.Errorf("Expect contact communication method id %d, but got %d", input.GetContactCommunicationMethodID(), resp[0].GetContactCommunicationMethodID())
		}

		if resp[0].GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp[0].GetCommunicationMethodCode())
		}

		if resp[0].GetCommunicationMethodLabelCode() != input.GetCommunicationMethodLabelCode() {
			t.Errorf("Expect communication method label code %s, but got %s", input.GetCommunicationMethodLabelCode(), resp[0].GetCommunicationMethodLabelCode())
		}

		if resp[0].GetCommunicationMethodLabelCaption() != input.GetCommunicationMethodLabelCaption() {
			t.Errorf("Expect communication method label caption %s, but got %s", input.GetCommunicationMethodLabelCaption(), resp[0].GetCommunicationMethodLabelCaption())
		}

		if resp[0].GetFormatValue() != input.GetFormatValue() {
			t.Errorf("Expect format value %s, but got %s", input.GetFormatValue(), resp[0].GetFormatValue())
		}

		if resp[0].GetStatus() != input.GetStatus() {
			t.Errorf("Expect status %s, but got %s", input.GetStatus(), resp[0].GetStatus())
		}

		if resp[0].GetIsDefault() != input.GetIsDefault() {
			t.Errorf("Expect default %v, but got %v", input.GetIsDefault(), resp[0].GetIsDefault())
		}

		if resp[0].GetContactCommunicationMethodField() == nil {
			t.Fatalf("Expect contact communication method fields is not nil")
		}

		if len(resp[0].GetContactCommunicationMethodField()) != 0 {
			t.Errorf("Expect no contact communication method fields")
		}
	}
}

func doReadAllExisting(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "communication_method_label_code", "communication_method_label_code", "format_value", "status", "is_default", "created_at", "modified_at", "vers"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetContactID(), data[0].GetContactCommunicationMethodID(), data[0].GetCommunicationMethodCode(), data[0].GetCommunicationMethodLabelCode(), data[0].GetCommunicationMethodLabelCaption(), data[0].GetFormatValue(), data[0].GetStatus(), data[0].GetIsDefault(), tmNow, tmNow, 1).
			AddRow(data[1].GetContactSystemCode(), data[1].GetContactID(), data[1].GetContactCommunicationMethodID(), data[1].GetCommunicationMethodCode(), data[1].GetCommunicationMethodLabelCode(), data[1].GetCommunicationMethodLabelCaption(), data[1].GetFormatValue(), data[1].GetStatus(), data[1].GetIsDefault(), tmNow, tmNow, 1).
			AddRow(data[2].GetContactSystemCode(), data[2].GetContactID(), data[2].GetContactCommunicationMethodID(), data[2].GetCommunicationMethodCode(), data[2].GetCommunicationMethodLabelCode(), data[2].GetCommunicationMethodLabelCaption(), data[2].GetFormatValue(), data[2].GetStatus(), data[2].GetIsDefault(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare(
			`SELECT ccm.contact_system_code, ccm.contact_id, ccm.contact_communication_method_id, 
				ccm.communication_method_code, ccm.communication_method_label_code, cml.caption, ccm.format_value, ccm.status, ccm.is_default, 
				ccm.created_at, ccm.modified_at, ccm.vers 
			FROM contact_communication_method ccm
			INNER JOIN communication_method_label cml`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnRows(rows)

		for _, item := range data {
			rows = sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "field_code", "field_value"}).
				AddRow(item.GetContactCommunicationMethodField()[0].GetContactSystemCode(), item.GetContactCommunicationMethodField()[0].GetContactID(), item.GetContactCommunicationMethodField()[0].GetContactCommunicationMethodID(), item.GetContactCommunicationMethodField()[0].GetFieldCode(), item.GetContactCommunicationMethodField()[0].GetFieldValue())

			expQuery = mock.ExpectPrepare(
				`SELECT contact_system_code, contact_id, contact_communication_method_id, field_code, field_value 
				FROM contact_communication_method_field`).ExpectQuery()
			expQuery.WithArgs(item.GetContactSystemCode(), item.GetContactID(), item.GetContactCommunicationMethodID()).WillReturnRows(rows)
		}

		resp, err := svc.DoReadAll(ctx, input.GetContactSystemCode(), input.GetContactID())
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}

		if resp == nil {
			t.Fatalf("Expect contact communication method is not nil")
		}

		if len(resp) < 3 {
			t.Errorf("Expect there are contact communication methods retrieved")
		}

		if resp[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), resp[0].GetContactSystemCode())
		}

		if resp[0].GetContactID() != input.GetContactID() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactID(), resp[0].GetContactID())
		}

		if resp[0].GetContactCommunicationMethodID() != input.GetContactCommunicationMethodID() {
			t.Errorf("Expect contact communication method id %d, but got %d", input.GetContactCommunicationMethodID(), resp[0].GetContactCommunicationMethodID())
		}

		if resp[0].GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), resp[0].GetCommunicationMethodCode())
		}

		if resp[0].GetCommunicationMethodLabelCode() != input.GetCommunicationMethodLabelCode() {
			t.Errorf("Expect communication method label code %s, but got %s", input.GetCommunicationMethodLabelCode(), resp[0].GetCommunicationMethodLabelCode())
		}

		if resp[0].GetCommunicationMethodLabelCaption() != input.GetCommunicationMethodLabelCaption() {
			t.Errorf("Expect communication method label caption %s, but got %s", input.GetCommunicationMethodLabelCaption(), resp[0].GetCommunicationMethodLabelCaption())
		}

		if resp[0].GetFormatValue() != input.GetFormatValue() {
			t.Errorf("Expect format value %s, but got %s", input.GetFormatValue(), resp[0].GetFormatValue())
		}

		if resp[0].GetStatus() != input.GetStatus() {
			t.Errorf("Expect status %s, but got %s", input.GetStatus(), resp[0].GetStatus())
		}

		if resp[0].GetIsDefault() != input.GetIsDefault() {
			t.Errorf("Expect default %v, but got %v", input.GetIsDefault(), resp[0].GetIsDefault())
		}

		if resp[0].GetContactCommunicationMethodField() == nil {
			t.Fatalf("Expect contact communication method fields is not nil")
		}

		if len(resp[0].GetContactCommunicationMethodField()) < 1 {
			t.Errorf("Expect there are contact communication method fields retrieved")
		}

		if resp[0].GetContactCommunicationMethodField()[0].GetContactSystemCode() != input.GetContactCommunicationMethodField()[0].GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactCommunicationMethodField()[0].GetContactSystemCode(), resp[0].GetContactCommunicationMethodField()[0].GetContactSystemCode())
		}

		if resp[0].GetContactCommunicationMethodField()[0].GetContactID() != input.GetContactCommunicationMethodField()[0].GetContactID() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactCommunicationMethodField()[0].GetContactID(), resp[0].GetContactCommunicationMethodField()[0].GetContactID())
		}

		if resp[0].GetContactCommunicationMethodField()[0].GetContactCommunicationMethodID() != input.GetContactCommunicationMethodID() {
			t.Errorf("Expect contact communication method id %d, but got %d", input.GetContactCommunicationMethodID(), resp[0].GetContactCommunicationMethodField()[0].GetContactCommunicationMethodID())
		}

		if resp[0].GetContactCommunicationMethodField()[0].GetFieldCode() != input.GetContactCommunicationMethodField()[0].GetFieldCode() {
			t.Errorf("Expect field code %s, but got %s", input.GetContactCommunicationMethodField()[0].GetFieldCode(), resp[0].GetContactCommunicationMethodField()[0].GetFieldCode())
		}

		if resp[0].GetContactCommunicationMethodField()[0].GetFieldValue() != input.GetContactCommunicationMethodField()[0].GetFieldValue() {
			t.Errorf("Expect field value %s, but got %s", input.GetContactCommunicationMethodField()[0].GetFieldValue(), resp[0].GetContactCommunicationMethodField()[0].GetFieldValue())
		}
	}
}

func doSave(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoSave fail validation Contact Communication Method", doSaveFailValidation(ctx))

		t.Run("DoSave invalid Contact", doSaveInvalidContact(ctx, input))

		t.Run("DoSave invalid Communication Method", doSaveInvalidCommunicationMethod(ctx, input))

		t.Run("DoSave invalid Communication Method Label", doSaveInvalidCommunicationMethodLabel(ctx, input))

		t.Run("DoSave fail new Contact Communication Method", doSaveFailNew(ctx, input))

		t.Run("DoSave fail new Contact Communication Method Field", doSaveFailNewField(ctx, input))

		t.Run("DoSave fail new Contact Communication Method", doSaveNew(ctx, input))

		t.Run("DoSave existing Contact Communication Method with existing Field", doSaveExisting(ctx, input))

		t.Run("DoSave existing Contact Communication Method with new Field", doSaveExistingWithNewField(ctx, input))
	}
}

func doSaveFailValidation(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		input := &contactcommunicationmethodmodel.ContactCommunicationMethod{
			ContactSystemCode:               "CNTSYS000000000001",
			ContactID:                       1,
			ContactCommunicationMethodID:    1,
			CommunicationMethodCode:         "EMAIL",
			CommunicationMethodLabelCode:    "HOME",
			CommunicationMethodLabelCaption: "Home",
			FormatValue:                     "test@gmail.com",
			Status:                          "A",
			IsDefault:                       true,
			ContactCommunicationMethodField: []*contactcommunicationmethodfieldmodel.ContactCommunicationMethodField{
				&contactcommunicationmethodfieldmodel.ContactCommunicationMethodField{
					ContactSystemCode:            "CNTSYS001",
					ContactID:                    1,
					ContactCommunicationMethodID: 1,
					FieldCode:                    "EMAIL_ADDRESS",
					FieldValue:                   "test@gmail.com",
				},
			},
		}

		err := svc.DoSave(ctx, input)
		if err == nil {
			t.Fatalf("Expect error is not nil")
		}
	}
}

func doSaveInvalidContact(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "first_name", "last_name", "status", "created_at", "modified_at", "vers"})

		expQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, contact_id, first_name, last_name, status, 
				created_at, modified_at, vers 
			FROM contact`).ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnRows(rows)

		err := svc.DoSave(ctx, input)
		if err == nil {
			t.Fatalf("Expect error is not nil")
		}
	}
}

func doSaveInvalidCommunicationMethod(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		contactRows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "first_name", "last_name", "status", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), "", "", "", tmNow, tmNow, 1)

		expContactQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, contact_id, first_name, last_name, status, 
				created_at, modified_at, vers 
			FROM contact`).ExpectQuery()
		expContactQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnRows(contactRows)

		commMethodRows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"})

		expCommMethodQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, description, details, status, format_field, 
				created_at, modified_at, vers 
			FROM communication_method`).ExpectQuery()
		expCommMethodQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(commMethodRows)

		err := svc.DoSave(ctx, input)
		if err == nil {
			t.Fatalf("Expect error is not nil")
		}
	}
}

func doSaveInvalidCommunicationMethodLabel(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		contactRows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "first_name", "last_name", "status", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), "", "", "", tmNow, tmNow, 1)

		expContactQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, contact_id, first_name, last_name, status, 
				created_at, modified_at, vers 
			FROM contact`).ExpectQuery()
		expContactQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnRows(contactRows)

		commMethodRows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), "", "", "", "", tmNow, tmNow, 1)

		expCommMethodQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, description, details, status, format_field, 
				created_at, modified_at, vers 
			FROM communication_method`).ExpectQuery()
		expCommMethodQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(commMethodRows)

		commMethodLabelRows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "communication_method_label_code", "caption"})

		expCommMethodLabelQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, communication_method_label_code, caption 
			FROM communication_method_label`).ExpectQuery()
		expCommMethodLabelQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode()).WillReturnRows(commMethodLabelRows)

		err := svc.DoSave(ctx, input)
		if err == nil {
			t.Fatalf("Expect error is not nil")
		}
	}
}

func doSaveFailNew(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		}

		contactRows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "first_name", "last_name", "status", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), "", "", "", tmNow, tmNow, 1)

		expContactQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, contact_id, first_name, last_name, status, 
				created_at, modified_at, vers 
			FROM contact`).ExpectQuery()
		expContactQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnRows(contactRows)

		commMethodRows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), "", "", "", "", tmNow, tmNow, 1)

		expCommMethodQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, description, details, status, format_field, 
				created_at, modified_at, vers 
			FROM communication_method`).ExpectQuery()
		expCommMethodQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(commMethodRows)

		commMethodLabelRows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "communication_method_label_code", "caption"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), "")

		expCommMethodLabelQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, communication_method_label_code, caption 
			FROM communication_method_label`).ExpectQuery()
		expCommMethodLabelQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode()).WillReturnRows(commMethodLabelRows)

		expUpdQuery := mock.ExpectPrepare("UPDATE contact_communication_method").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetFormatValue(), input.GetStatus(), input.GetIsDefault(), tmNow).WillReturnResult(sqlmock.NewResult(0, 0))

		expInsQuery := mock.ExpectPrepare("INSERT INTO contact_communication_method").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetFormatValue(), input.GetStatus(), input.GetIsDefault(), tmNow, tmNow).WillReturnError(fmt.Errorf("DoInsert contact communication method failed"))

		err := svc.DoSave(ctx, input)
		if err == nil {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doSaveFailNewField(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		}

		contactRows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "first_name", "last_name", "status", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), "", "", "", tmNow, tmNow, 1)

		expContactQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, contact_id, first_name, last_name, status, 
				created_at, modified_at, vers 
			FROM contact`).ExpectQuery()
		expContactQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnRows(contactRows)

		commMethodRows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), "", "", "", "", tmNow, tmNow, 1)

		expCommMethodQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, description, details, status, format_field, 
				created_at, modified_at, vers 
			FROM communication_method`).ExpectQuery()
		expCommMethodQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(commMethodRows)

		commMethodLabelRows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "communication_method_label_code", "caption"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), "")

		expCommMethodLabelQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, communication_method_label_code, caption 
			FROM communication_method_label`).ExpectQuery()
		expCommMethodLabelQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode()).WillReturnRows(commMethodLabelRows)

		expUpdQuery := mock.ExpectPrepare("UPDATE contact_communication_method").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetFormatValue(), input.GetStatus(), input.GetIsDefault(), tmNow).WillReturnResult(sqlmock.NewResult(0, 0))

		expInsQuery := mock.ExpectPrepare("INSERT INTO contact_communication_method").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetFormatValue(), input.GetStatus(), input.GetIsDefault(), tmNow, tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		expInsQuery = mock.ExpectPrepare("INSERT INTO contact_communication_method_field").ExpectExec()
		expInsQuery.WithArgs(input.GetContactCommunicationMethodField()[0].GetContactSystemCode(), input.GetContactCommunicationMethodField()[0].GetContactID(), input.GetContactCommunicationMethodField()[0].GetContactCommunicationMethodID(), input.GetContactCommunicationMethodField()[0].GetFieldCode(), input.GetContactCommunicationMethodField()[0].GetFieldValue()).WillReturnError(fmt.Errorf("DoInsert contact communication method field failed"))

		err := svc.DoSave(ctx, input)
		if err == nil {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doSaveNew(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		}

		contactRows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "first_name", "last_name", "status", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), "", "", "", tmNow, tmNow, 1)

		expContactQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, contact_id, first_name, last_name, status, 
				created_at, modified_at, vers 
			FROM contact`).ExpectQuery()
		expContactQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnRows(contactRows)

		commMethodRows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), "", "", "", "", tmNow, tmNow, 1)

		expCommMethodQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, description, details, status, format_field, 
				created_at, modified_at, vers 
			FROM communication_method`).ExpectQuery()
		expCommMethodQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(commMethodRows)

		commMethodLabelRows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "communication_method_label_code", "caption"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), "")

		expCommMethodLabelQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, communication_method_label_code, caption 
			FROM communication_method_label`).ExpectQuery()
		expCommMethodLabelQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode()).WillReturnRows(commMethodLabelRows)

		expUpdQuery := mock.ExpectPrepare("UPDATE contact_communication_method").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetFormatValue(), input.GetStatus(), input.GetIsDefault(), tmNow).WillReturnResult(sqlmock.NewResult(0, 0))

		expInsQuery := mock.ExpectPrepare("INSERT INTO contact_communication_method").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetFormatValue(), input.GetStatus(), input.GetIsDefault(), tmNow, tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		expInsQuery = mock.ExpectPrepare("INSERT INTO contact_communication_method_field").ExpectExec()
		expInsQuery.WithArgs(input.GetContactCommunicationMethodField()[0].GetContactSystemCode(), input.GetContactCommunicationMethodField()[0].GetContactID(), input.GetContactCommunicationMethodField()[0].GetContactCommunicationMethodID(), input.GetContactCommunicationMethodField()[0].GetFieldCode(), input.GetContactCommunicationMethodField()[0].GetFieldValue()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoSave(ctx, input)
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}
	}
}

func doSaveExisting(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		}

		contactRows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "first_name", "last_name", "status", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), "", "", "", tmNow, tmNow, 1)

		expContactQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, contact_id, first_name, last_name, status, 
			created_at, modified_at, vers 
		FROM contact`).ExpectQuery()
		expContactQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnRows(contactRows)

		commMethodRows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), "", "", "", "", tmNow, tmNow, 1)

		expCommMethodQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, description, details, status, format_field, 
			created_at, modified_at, vers 
		FROM communication_method`).ExpectQuery()
		expCommMethodQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(commMethodRows)

		commMethodLabelRows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "communication_method_label_code", "caption"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), "")

		expCommMethodLabelQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, communication_method_label_code, caption 
		FROM communication_method_label`).ExpectQuery()
		expCommMethodLabelQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode()).WillReturnRows(commMethodLabelRows)

		expUpdQuery := mock.ExpectPrepare("UPDATE contact_communication_method").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetFormatValue(), input.GetStatus(), input.GetIsDefault(), tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		expUpdQuery = mock.ExpectPrepare("UPDATE contact_communication_method_field").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactCommunicationMethodField()[0].GetContactSystemCode(), input.GetContactCommunicationMethodField()[0].GetContactID(), input.GetContactCommunicationMethodField()[0].GetContactCommunicationMethodID(), input.GetContactCommunicationMethodField()[0].GetFieldCode(), input.GetContactCommunicationMethodField()[0].GetFieldValue()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoSave(ctx, input)
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}
	}
}

func doSaveExistingWithNewField(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		input.Audit = &auditmodel.Audit{
			CreatedAt:  tmNow,
			ModifiedAt: tmNow,
			Vers:       1,
		}

		contactRows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "first_name", "last_name", "status", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), "", "", "", tmNow, tmNow, 1)

		expContactQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, contact_id, first_name, last_name, status, 
			created_at, modified_at, vers 
		FROM contact`).ExpectQuery()
		expContactQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID()).WillReturnRows(contactRows)

		commMethodRows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), "", "", "", "", tmNow, tmNow, 1)

		expCommMethodQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, description, details, status, format_field, 
			created_at, modified_at, vers 
		FROM communication_method`).ExpectQuery()
		expCommMethodQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(commMethodRows)

		commMethodLabelRows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "communication_method_label_code", "caption"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), "")

		expCommMethodLabelQuery := mock.ExpectPrepare(
			`SELECT contact_system_code, communication_method_code, communication_method_label_code, caption 
		FROM communication_method_label`).ExpectQuery()
		expCommMethodLabelQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode()).WillReturnRows(commMethodLabelRows)

		expUpdQuery := mock.ExpectPrepare("UPDATE contact_communication_method").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetFormatValue(), input.GetStatus(), input.GetIsDefault(), tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		expUpdQuery = mock.ExpectPrepare("UPDATE contact_communication_method_field").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactCommunicationMethodField()[0].GetContactSystemCode(), input.GetContactCommunicationMethodField()[0].GetContactID(), input.GetContactCommunicationMethodField()[0].GetContactCommunicationMethodID(), input.GetContactCommunicationMethodField()[0].GetFieldCode(), input.GetContactCommunicationMethodField()[0].GetFieldValue()).WillReturnResult(sqlmock.NewResult(0, 0))

		expInsQuery := mock.ExpectPrepare("INSERT INTO contact_communication_method_field").ExpectExec()
		expInsQuery.WithArgs(input.GetContactCommunicationMethodField()[0].GetContactSystemCode(), input.GetContactCommunicationMethodField()[0].GetContactID(), input.GetContactCommunicationMethodField()[0].GetContactCommunicationMethodID(), input.GetContactCommunicationMethodField()[0].GetFieldCode(), input.GetContactCommunicationMethodField()[0].GetFieldValue()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoSave(ctx, input)
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}
	}
}

func doDelete(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoDelete fail default Contact Communication Method", doDeleteFailDefault(ctx, input))

		t.Run("DoDelete default Contact Communication Method", doDeleteDefault(ctx, input))

		t.Run("DoDelete fail Contact Communication Method Field", doDeleteFailField(ctx, input))

		t.Run("DoDelete existing Contact Communication Method", doDeleteExisting(ctx, input))
	}
}

func doDeleteFailDefault(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expDefCommMethodQuery := mock.ExpectPrepare(
			`SELECT ccm.contact_system_code, ccm.contact_id, ccm.contact_communication_method_id, 
				ccm.communication_method_code, ccm.communication_method_label_code, cml.caption, ccm.format_value, ccm.status, ccm.is_default, 
				ccm.created_at, ccm.modified_at, ccm.vers 
			FROM contact_communication_method ccm
			INNER JOIN communication_method_label cml`).ExpectQuery()
		expDefCommMethodQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnError(fmt.Errorf("DoRead contact communication method failed"))

		err := svc.DoDelete(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID())
		if err == nil {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doDeleteDefault(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "communication_method_label_code", "communication_method_label_code", "format_value", "status", "is_default", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetCommunicationMethodLabelCaption(), input.GetFormatValue(), "A", true, tmNow, tmNow, 1)

		expDefCommMethodQuery := mock.ExpectPrepare(
			`SELECT ccm.contact_system_code, ccm.contact_id, ccm.contact_communication_method_id, 
				ccm.communication_method_code, ccm.communication_method_label_code, cml.caption, ccm.format_value, ccm.status, ccm.is_default, 
				ccm.created_at, ccm.modified_at, ccm.vers 
			FROM contact_communication_method ccm
			INNER JOIN communication_method_label cml`).ExpectQuery()
		expDefCommMethodQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnRows(rows)

		err := svc.DoDelete(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID())
		if err == nil {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doDeleteFailField(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "communication_method_label_code", "communication_method_label_code", "format_value", "status", "is_default", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetCommunicationMethodLabelCaption(), input.GetFormatValue(), input.GetStatus(), input.GetIsDefault(), tmNow, tmNow, 1)

		expDefCommMethodQuery := mock.ExpectPrepare(
			`SELECT ccm.contact_system_code, ccm.contact_id, ccm.contact_communication_method_id, 
				ccm.communication_method_code, ccm.communication_method_label_code, cml.caption, ccm.format_value, ccm.status, ccm.is_default, 
				ccm.created_at, ccm.modified_at, ccm.vers 
			FROM contact_communication_method ccm
			INNER JOIN communication_method_label cml`).ExpectQuery()
		expDefCommMethodQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnRows(rows)

		expQuery := mock.ExpectPrepare("DELETE FROM contact_communication_method_field").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnError(fmt.Errorf("Delete all contact communication method fields failed"))

		err := svc.DoDelete(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID())
		if err == nil {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doDeleteExisting(ctx context.Context, input *contactcommunicationmethodmodel.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "communication_method_label_code", "communication_method_label_code", "format_value", "status", "is_default", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetCommunicationMethodLabelCaption(), input.GetFormatValue(), input.GetStatus(), input.GetIsDefault(), tmNow, tmNow, 1)

		expDefCommMethodQuery := mock.ExpectPrepare(
			`SELECT ccm.contact_system_code, ccm.contact_id, ccm.contact_communication_method_id, 
				ccm.communication_method_code, ccm.communication_method_label_code, cml.caption, ccm.format_value, ccm.status, ccm.is_default, 
				ccm.created_at, ccm.modified_at, ccm.vers 
			FROM contact_communication_method ccm
			INNER JOIN communication_method_label cml`).ExpectQuery()
		expDefCommMethodQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnRows(rows)

		expQuery := mock.ExpectPrepare("DELETE FROM contact_communication_method_field").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnResult(sqlmock.NewResult(0, 1))

		expQuery = mock.ExpectPrepare("DELETE FROM contact_communication_method").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := svc.DoDelete(ctx, input.GetContactSystemCode(), input.GetContactID(), input.GetContactCommunicationMethodID())
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}
	}
}

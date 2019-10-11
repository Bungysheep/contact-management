package contactsystem

import (
	"context"
	"os"
	"testing"

	"github.com/bungysheep/contact-management/pkg/api/v1/contactsystem"
)

var (
	ctx context.Context
	svc contactsystem.ContactSystemServiceServer
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	svc = NewContactSystemService()

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestContactSystemService(t *testing.T) {
	t.Run("DoRead Contact System", doRead(ctx))

	t.Run("DoReadAll Contact System", doRead(ctx))

	t.Run("DoSave Contact System", doRead(ctx))

	t.Run("DoDelete Contact System", doRead(ctx))
}

func doRead(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoRead(ctx, &contactsystem.DoReadRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}

func doReadAll(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoReadAll(ctx, &contactsystem.DoReadAllRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}

func doSave(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoSave(ctx, &contactsystem.DoSaveRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}

func doDelete(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoDelete(ctx, &contactsystem.DoDeleteRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}

package contact

import (
	"context"
	"os"
	"testing"

	"github.com/bungysheep/contact-management/pkg/api/v1/contact"
)

var (
	ctx context.Context
	svc contact.ContactServiceServer
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	svc = NewContactService()

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestContactService(t *testing.T) {
	t.Run("DoRead Contact", doRead(ctx))

	t.Run("DoReadAll Contact", doRead(ctx))

	t.Run("DoSave Contact", doRead(ctx))

	t.Run("DoDelete Contact", doRead(ctx))
}

func doRead(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoRead(ctx, &contact.DoReadRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}

func doReadAll(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoReadAll(ctx, &contact.DoReadAllRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}

func doSave(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoSave(ctx, &contact.DoSaveRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}

func doDelete(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoDelete(ctx, &contact.DoDeleteRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}

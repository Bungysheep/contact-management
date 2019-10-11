package communicationmethod

import (
	"context"
	"os"
	"testing"

	"github.com/bungysheep/contact-management/pkg/api/v1/communicationmethod"
)

var (
	ctx context.Context
	svc communicationmethod.CommunicationMethodServiceServer
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	svc = NewCommunicationMethodService()

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestCommunicationMethodService(t *testing.T) {
	t.Run("DoRead Communication Method", doRead(ctx))

	t.Run("DoReadAll Communication Method", doRead(ctx))

	t.Run("DoSave Communication Method", doRead(ctx))

	t.Run("DoDelete Communication Method", doRead(ctx))
}

func doRead(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoRead(ctx, &communicationmethod.DoReadRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}

func doReadAll(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoReadAll(ctx, &communicationmethod.DoReadAllRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}

func doSave(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoSave(ctx, &communicationmethod.DoSaveRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}

func doDelete(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoDelete(ctx, &communicationmethod.DoDeleteRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}

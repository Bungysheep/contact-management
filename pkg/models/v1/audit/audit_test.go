package audit

import (
	"testing"
	"time"
)

func TestCreateAudit(t *testing.T) {
	audit := NewAudit()

	if audit == nil {
		t.Fatalf("Expect audit is not nil")
	}

	timeNow := time.Now()

	audit.CreatedAt = timeNow
	audit.ModifiedAt = timeNow
	audit.Vers = 1

	if audit.GetCreatedAt() != timeNow {
		t.Errorf("Expect craeted at %v, but got %v", timeNow, audit.GetCreatedAt())
	}

	if audit.GetModifiedAt() != timeNow {
		t.Errorf("Expect modified at %v, but got %v", timeNow, audit.GetModifiedAt())
	}

	if audit.GetVers() != 1 {
		t.Errorf("Expect vers %v, but got %v", 1, audit.GetVers())
	}
}

func TestValidate(t *testing.T) {
	audit := NewAudit()

	if audit == nil {
		t.Fatalf("Expect audit is not nil")
	}

	timeNow := time.Now()

	audit.CreatedAt = timeNow
	audit.ModifiedAt = timeNow
	audit.Vers = 1

	if err := audit.DoValidate(); err != nil {
		t.Fatalf("Expect error is nil, but got %v", err)
	}
}

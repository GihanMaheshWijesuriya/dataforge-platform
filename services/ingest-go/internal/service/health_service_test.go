package service

import "testing"

func TestHealthService_Status(t *testing.T) {
	svc := NewHealthService()
	if got := svc.Status(); got != "ok" {
		t.Fatalf("expected ok, got %s", got)
	}
}

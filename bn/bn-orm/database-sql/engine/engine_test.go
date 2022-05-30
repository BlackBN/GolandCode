package engine

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestNewEngine(t *testing.T) {
	t.Helper()
	engine, err := New("mysql", "test:test@tcp(127.0.0.1:3306)/test")
	if err != nil {
		t.Fatal("Failed to connect", err)
	}
	engine.Close()
}

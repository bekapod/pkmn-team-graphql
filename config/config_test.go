package config

import (
	"os"
	"testing"

	"github.com/go-test/deep"
)

func TestNew(t *testing.T) {
	os.Setenv("APP_HOST", "localhost")
	os.Setenv("PORT", "1234")
	os.Setenv("DATABASE_URL", "postgres://user:pass@host:1234/database-name?sslmode=disable")
	os.Setenv("TRACING", "1")

	exp := Config{
		AppHost:     "localhost",
		Port:        "1234",
		DatabaseUrl: "postgres://user:pass@host:1234/database-name?sslmode=disable",
		Tracing:     true,
	}

	got := New()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}

	os.Clearenv()
}

func TestConfig_DatabaseName(t *testing.T) {
	cfg := Config{
		AppHost:     "localhost",
		Port:        "1234",
		DatabaseUrl: "postgres://user:pass@host:1234/database-name?sslmode=disable",
		Tracing:     false,
	}

	exp := "database-name"
	if got := cfg.DatabaseName(); exp != got {
		t.Errorf("expected '%s' but got '%s' instead", exp, got)
	}
}

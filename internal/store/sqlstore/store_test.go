package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

//TestMain ...
func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=recumbent_stone_test user=postgres password=12345 sslmode=disable"
	}

	os.Exit(m.Run())
}

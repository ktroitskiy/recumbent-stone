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
		databaseURL = "host=localhost dbname=recumbent_stone_test user=postgres"
	}

	os.Exit(m.Run())
}

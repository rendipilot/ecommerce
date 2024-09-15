package dbTest

import (
	"e-commerce-synapsis/database"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestPgOpenConnection(t *testing.T) {
	os.Setenv("DB_HOST_POSTGRES", "localhost")
	os.Setenv("DB_PORT_POSTGRES", "5432")
	os.Setenv("DB_USER_POSTGRES", "rendy")
	os.Setenv("DB_PW_POSTGRES", "rendy")
	os.Setenv("DB_NAME_POSTGRES", "rendy")
	os.Setenv("DB_APP_NAME", "rendy")

	// Buat mock database connection
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	conn := database.PgOpenConnection()

	require.NotNil(t, conn)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("There were unmet expectations: %s", err)
	}
}
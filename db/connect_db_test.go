package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testDsn = "host=localhost user=postgres password=keranglaut26 dbname=todo_app_test port=5432 sslmode=disable"
const invalidDsn = "host=localhost user=postgres password=kdpkadw dbname=ab port=9301 sslmode=disable"

func TestDBConnection(t *testing.T) {
	t.Run("Should estabilish a database connection succesfully", func(t *testing.T) {
		conn, err := NewDBConnection(testDsn)
		
		assert.NoError(t, err, "expected no error")
		assert.NotNil(t, conn, "expected Conn_DB not nil")
		
		if conn != nil {
			assert.NotNil(t, conn.DB, "expected *sql.BD not nil")
			err := conn.DB.Close()
			assert.NoError(t, err, "expected no error when closing db conn")
		}
	})


	t.Run("Database connnection should error for invalid dsn", func(t *testing.T) {
		conn, err := NewDBConnection(invalidDsn)

		assert.Error(t, err, "expected an error with an invalid dsn")
		assert.Nil(t, conn, "expected conn to be nil on connected failure")
	})
}
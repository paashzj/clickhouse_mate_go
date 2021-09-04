package metrics

import (
	"database/sql"
	"testing"

	_ "github.com/mailru/go-clickhouse"
)

func TestMetrics(t *testing.T) {
	connect, err := sql.Open("clickhouse", "http://localhost:8123/default")
	if err != nil {
		t.Error(err)
	}
	err = connect.Ping()
	if err != nil {
		t.Error(err)
	}
	rows, err := connect.Query("SELECT * FROM system.metrics")
	if err != nil {
		t.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			metrics string
			value int
			description string
		)
		err := rows.Scan(&metrics, &value, &description)
		if err != nil {
			t.Error(err)
		}
	}
}

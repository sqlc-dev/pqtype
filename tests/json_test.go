package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sqlc-dev/pqtype"
)

func TestJSONRawMessage(t *testing.T) {
	for _, payload := range []string{
		`{}`,
		`[]`,
		`1`,
		`1.2`,
		`"a"`,
		`true`,
		`false`,
		`{"foo": "bar"}`,
	} {
		payload = payload
		t.Run(payload, func(t *testing.T) {
			var n pqtype.NullRawMessage
			if err := db.QueryRow(fmt.Sprintf(`SELECT '%s'::json`, payload)).Scan(&n); err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(true, n.Valid); diff != "" {
				t.Errorf("json valid mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(string(json.RawMessage(payload)), string(n.RawMessage)); diff != "" {
				t.Errorf("json mismatch (-want +got):\n%s", diff)
			}
			if _, err := db.Exec(`SELECT json_typeof($1)`, n); err != nil {
				t.Fatal(err)
			}
			if err := db.QueryRow(fmt.Sprintf(`SELECT '%s'::jsonb`, payload)).Scan(&n); err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(true, n.Valid); diff != "" {
				t.Errorf("jsonb valid mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(string(json.RawMessage(payload)), string(n.RawMessage)); diff != "" {
				t.Errorf("jsonb mismatch (-want +got):\n%s", diff)
			}
			if _, err := db.Exec(`SELECT jsonb_typeof($1)`, n); err != nil {
				t.Fatal(err)
			}
		})
		t.Run("/stdlib/"+payload, func(t *testing.T) {
			var n json.RawMessage
			if err := db.QueryRow(fmt.Sprintf(`SELECT '%s'::json`, payload)).Scan(&n); err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(string(json.RawMessage(payload)), string(n)); diff != "" {
				t.Errorf("json mismatch (-want +got):\n%s", diff)
			}
			if _, err := db.Exec(`SELECT json_typeof($1)`, n); err != nil {
				t.Fatal(err)
			}
			if err := db.QueryRow(fmt.Sprintf(`SELECT '%s'::jsonb`, payload)).Scan(&n); err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(string(json.RawMessage(payload)), string(n)); diff != "" {
				t.Errorf("jsonb mismatch (-want +got):\n%s", diff)
			}
			if _, err := db.Exec(`SELECT jsonb_typeof($1)`, n); err != nil {
				t.Fatal(err)
			}
		})
	}
	t.Run("NULL", func(t *testing.T) {
		var n pqtype.NullRawMessage
		if err := db.QueryRow(`SELECT NULL::json`).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(false, n.Valid); diff != "" {
			t.Errorf("valid mismatch (-want +got):\n%s", diff)
		}
		if _, err := db.Exec(`SELECT json_typeof($1)`, n); err != nil {
			t.Fatal(err)
		}
		if err := db.QueryRow(`SELECT NULL::jsonb`).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(false, n.Valid); diff != "" {
			t.Errorf("valid mismatch (-want +got):\n%s", diff)
		}
		if _, err := db.Exec(`SELECT jsonb_typeof($1)`, n); err != nil {
			t.Fatal(err)
		}
	})
}

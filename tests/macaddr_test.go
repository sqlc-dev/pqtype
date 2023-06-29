package tests

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sqlc-dev/pqtype"
)

// https://www.postgresql.org/docs/current/datatype-net-types.html#DATATYPE-MACADDR
func TestMacaddr(t *testing.T) {
	for _, addr := range []struct {
		input  string
		output string
	}{
		{
			"08:00:2b:01:02:03",
			"08:00:2b:01:02:03",
		},
		{
			"08-00-2b-01-02-03",
			"08:00:2b:01:02:03",
		},
		{
			"08002b:010203",
			"08:00:2b:01:02:03",
		},
		{
			"08002b-010203",
			"08:00:2b:01:02:03",
		},
		{
			"0800.2b01.0203",
			"08:00:2b:01:02:03",
		},
		{
			"0800-2b01-0203",
			"08:00:2b:01:02:03",
		},
		{
			"08002b010203",
			"08:00:2b:01:02:03",
		},
	} {
		addr = addr
		t.Run(addr.input, func(t *testing.T) {
			var cidr pqtype.Macaddr
			if err := db.QueryRow(fmt.Sprintf(`SELECT '%s'::macaddr`, addr.input)).Scan(&cidr); err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(true, cidr.Valid); diff != "" {
				t.Errorf("valid mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(addr.output, cidr.Addr.String()); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			if _, err := db.Exec(`SELECT trunc($1::macaddr)`, cidr); err != nil {
				t.Fatal(err)
			}
		})
	}
	t.Run("NULL", func(t *testing.T) {
		var cidr pqtype.Macaddr
		if err := db.QueryRow(fmt.Sprintf(`SELECT NULL::macaddr`)).Scan(&cidr); err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(false, cidr.Valid); diff != "" {
			t.Errorf("valid mismatch (-want +got):\n%s", diff)
		}
		if _, err := db.Exec(`SELECT trunc($1::macaddr)`, cidr); err != nil {
			t.Fatal(err)
		}
	})
}

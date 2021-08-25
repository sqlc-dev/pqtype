package tests

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tabbed/pqtype"
)

func TestInet(t *testing.T) {
	for _, addr := range []string{
		"0.0.0.0/32",
		"127.0.0.1/8",
		"12.34.56.65/32",
		"192.168.1.16/24",
		"255.0.0.0/8",
		"255.255.255.255/32",
		"10.0.0.1/32",
		"2607:f8b0:4009:80b::200e/128",
		"::1/64",
		"::/0",
		"::1/128",
		"2607:f8b0:4009:80b::200e/64",
	} {
		addr = addr
		t.Run(addr, func(t *testing.T) {
			var ip pqtype.Inet
			if err := db.QueryRow(fmt.Sprintf(`SELECT '%s'::inet`, addr)).Scan(&ip); err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(true, ip.Valid); diff != "" {
				t.Errorf("valid mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(addr, ip.IPNet.String()); diff != "" {
				t.Errorf("IPNet mismatch (-want +got):\n%s", diff)
			}
			if _, err := db.Exec(`SELECT $1`, ip); err != nil {
				t.Fatal(err)
			}
		})
	}
	t.Run("NULL", func(t *testing.T) {
		var ip pqtype.Inet
		if err := db.QueryRow(`SELECT NULL::inet`).Scan(&ip); err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(false, ip.Valid); diff != "" {
			t.Errorf("valid mismatch (-want +got):\n%s", diff)
		}
		if _, err := db.Exec(`SELECT $1`, ip); err != nil {
			t.Fatal(err)
		}
	})
}

package tests

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sqlc-dev/pqtype"
)

func TestCIDR(t *testing.T) {
	// https://www.postgresql.org/docs/current/datatype-net-types.html#DATATYPE-NET-CIDR-TABLE
	for _, addr := range []struct {
		input  string
		output string
	}{
		{
			"192.168.100.128/25",
			"192.168.100.128/25",
		},
		{
			"192.168/24",
			"192.168.0.0/24",
		},
		{
			"192.168/25",
			"192.168.0.0/25",
		},
		{
			"192.168.1",
			"192.168.1.0/24",
		},
		{
			"192.168",
			"192.168.0.0/24",
		},
		{
			"128.1",
			"128.1.0.0/16",
		},
		{
			"128",
			"128.0.0.0/16",
		},
		{
			"128.1.2",
			"128.1.2.0/24",
		},
		{
			"10.1.2",
			"10.1.2.0/24",
		},
		{
			"10.1",
			"10.1.0.0/16",
		},
		{
			"10",
			"10.0.0.0/8",
		},
		{
			"10.1.2.3/32",
			"10.1.2.3/32",
		},
		{
			"2001:4f8:3:ba::/64",
			"2001:4f8:3:ba::/64",
		},
		{
			"2001:4f8:3:ba:2e0:81ff:fe22:d1f1/128",
			"2001:4f8:3:ba:2e0:81ff:fe22:d1f1/128",
		},
	} {
		addr = addr
		t.Run(addr.input, func(t *testing.T) {
			var cidr pqtype.CIDR
			if err := db.QueryRow(fmt.Sprintf(`SELECT '%s'::cidr`, addr.input)).Scan(&cidr); err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(true, cidr.Valid); diff != "" {
				t.Errorf("valid mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(addr.output, cidr.IPNet.String()); diff != "" {
				t.Errorf("IPNet mismatch (-want +got):\n%s", diff)
			}
		})
	}
	t.Run("NULL", func(t *testing.T) {
		var cidr pqtype.CIDR
		if err := db.QueryRow(`SELECT NULL::cidr`).Scan(&cidr); err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(false, cidr.Valid); diff != "" {
			t.Errorf("valid mismatch (-want +got):\n%s", diff)
		}
	})
}

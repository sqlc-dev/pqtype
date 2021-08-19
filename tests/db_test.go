package tests

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	// Build the connection string from the environment
	params := []string{}
	if host := os.Getenv("PG_HOST"); host != "" {
		params = append(params, fmt.Sprintf("host=%s", host))
	}
	if user := os.Getenv("PG_USER"); user != "" {
		params = append(params, fmt.Sprintf("user=%s", user))
	}
	if pass := os.Getenv("PG_PASSWORD"); pass != "" {
		params = append(params, fmt.Sprintf("password=%s", pass))
	}
	if port := os.Getenv("PG_PORT"); port != "" {
		params = append(params, fmt.Sprintf("port=%s", port))
	}
	if name := os.Getenv("PG_DATABASE"); name != "" {
		params = append(params, fmt.Sprintf("dbname=%s", name))
	} else {
		params = append(params, "dbname=pqtype")
	}
	db, err = sql.Open("postgres", strings.Join(append(params, "sslmode=disable"), " "))
	if err != nil {
		panic(err)
	}
}

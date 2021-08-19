module github.com/tabbed/pqtype/tests

go 1.17

require (
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgtype v1.8.1 // indirect
	github.com/lib/pq v1.10.2 // indirect
	github.com/tabbed/pqtype v1.0.0
)

require github.com/google/go-cmp v0.5.6 // indirect

replace github.com/tabbed/pqtype => ../

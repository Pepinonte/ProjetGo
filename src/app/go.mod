module app

go 1.21.6

replace app/foldermanagement => ./foldermanagement

replace app/filesmanagement => ./filesmanagement

replace app/database => ./database

require (
	app/database v0.0.0-00010101000000-000000000000
	app/filesmanagement v0.0.0-00010101000000-000000000000
	app/foldermanagement v0.0.0-00010101000000-000000000000
)

require github.com/go-sql-driver/mysql v1.7.1 // indirect

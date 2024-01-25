module app

go 1.21.6

replace app/foldermanagement => ./foldermanagement

replace app/filesmanagement => ./filesmanagement

require (
	app/filesmanagement v0.0.0-00010101000000-000000000000
	app/foldermanagement v0.0.0-00010101000000-000000000000
)

module github.com/wintermonth2298/feedback-service

go 1.18

require (
	github.com/Masterminds/squirrel v1.5.3
	github.com/jackc/pgx/v4 v4.16.1
	github.com/jmoiron/sqlx v1.3.5
	github.com/wintermonth2298/feedback-service/api/feedback-service v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.48.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/net v0.0.0-20220708220712-1185a9018129 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220719170305-83ca9fad585f // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace github.com/wintermonth2298/feedback-service/api/feedback-service => ./api/feedback-service

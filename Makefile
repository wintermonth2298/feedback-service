protogen:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		./api/feedback-service/feedback-service.proto

goose_up:
	goose -dir migrations postgres "user=winte password=123 dbname=xployalty_test sslmode=disable" up

goose_down:
	goose -dir migrations postgres "user=winte password=123 dbname=xployalty_test sslmode=disable" down

goose_redo:
	goose -dir migrations postgres "user=winte password=123 dbname=xployalty_test sslmode=disable" redo

goose_status:
	goose -dir migrations postgres "user=winte password=123 dbname=xployalty_test sslmode=disable" status

